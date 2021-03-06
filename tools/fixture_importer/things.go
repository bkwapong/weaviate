/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */
package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/creativesoftwarefdn/weaviate/client/things"
	"github.com/creativesoftwarefdn/weaviate/models"
	spew "github.com/davecgh/go-spew/spew"
	"github.com/go-openapi/strfmt"
)

func createThings() {
	for _, thing := range demoDataset.Things {
		className := thing["class"].(string)
		classKinds[className] = "Thing"
		uuid := thing["uuid"].(string)

		properties := map[string]interface{}{}

		for key, value := range thing {
			if key == "class" || key == "uuid" {
				continue
			}

			switch ref := value.(type) {
			case map[string]interface{}: // a single object implies a single reference
				var location string
				location, ok := ref["location"].(string)
				if !ok {
					location = ""
				}
				thingFixups = append(thingFixups, fixupAddRef{
					fromId:       uuid,
					fromProperty: key,
					toClass:      ref["class"].(string),
					toId:         ref["uuid"].(string),
					location:     location,
				})
			case []interface{}: // a list of objects implies multiple references
				multiFixUps := []fixupAddRef{}
				for _, singleRef := range ref {
					singleRefMap, ok := singleRef.(map[string]interface{})
					if !ok {
						panic(fmt.Sprintf("have []interface{}, but items is not a ref, instead have %#v", singleRef))
					}

					var location string
					location, ok = singleRefMap["location"].(string)
					if !ok {
						location = ""
					}
					toClass, ok := singleRefMap["class"].(string)
					if !ok {
						toClass = ""
					}
					multiFixUps = append(thingFixups, fixupAddRef{
						fromId:       uuid,
						fromProperty: key,
						toClass:      toClass,
						toId:         singleRefMap["uuid"].(string),
						location:     location,
					})
				}
				thingManyFixups = append(thingManyFixups, multiFixUps)
			default: // everything else must be a primitive
				class := findClass(schema.Things, className)
				property := findProperty(class, key)
				if len(property.AtDataType) != 1 {
					panic(fmt.Sprintf("Only one datatype supported for import. Failed in thing %s.%s with @dataTypes %#v on value %t",
						className, property.Name, property.AtDataType, value))
				}
				dataType := property.AtDataType[0]

				switch dataType {
				case "string", "date":
					properties[key] = value
				case "int":
					switch typedValue := value.(type) {
					case string:
						v, err := strconv.ParseInt(typedValue, 10, 64)
						if err != nil {
							panic(err)
						}
						properties[key] = v
					case float64:
						properties[key] = int(typedValue)
					default:
						panic("Unexpected type")
					}
				case "number":
					properties[key] = value.(float64)
				case "boolean":
					properties[key] = value.(bool)
				default:
					panic(fmt.Sprintf("No such datatype supported: %s", dataType))
				}
			}
		}

		t := models.ThingCreate{
			AtContext: "http://example.org",
			AtClass:   className,
			Schema:    properties,
		}

		thing := assertCreateThing(&t)
		idMap[uuid] = string(thing.ThingID) // Store mapping of ID's
		fmt.Printf("Created Thing %s\n", thing.ThingID)
	}
}

func fixupThings() {
	fmt.Printf("Checking if all things that need a patch are created.\n")
	for {
		allExist := true

		for _, fixup := range thingFixups {
			if fixup.location != "" {
				// it's a network ref, we can't do any validation
				continue
			}

			if !checkThingExists(idMap[fixup.fromId]) {
				allExist = false
				fmt.Printf("From does not exist! %v\n", idMap[fixup.fromId])
			}
			if classKinds[fixup.toClass] == "Action" {
				if !checkActionExists(idMap[fixup.toId]) {
					allExist = false
					fmt.Printf("To action does not exist! %v\n", idMap[fixup.toId])
					break
				}
			} else { // assume it's a thing.
				if !checkThingExists(idMap[fixup.toId]) {
					allExist = false
					fmt.Printf("To thing does not exist! %v\n", idMap[fixup.toId])
					break
				}
			}
		}

		if allExist {
			fmt.Printf("Everything that needs to be patched exists!\n")
			break
		} else {
			fmt.Printf("Not everything that needs to be patched exists\n")

			var waitSecondsUntilSettled time.Duration = 2 * time.Second
			fmt.Printf("Waiting for %v to settle\n", waitSecondsUntilSettled)
			time.Sleep(waitSecondsUntilSettled)
			continue
		}
	}

	// Now fix up refs
	op := "add"
	for _, fixup := range thingFixups {
		var patch *models.PatchDocument
		path := fmt.Sprintf("/schema/%s", fixup.fromProperty)

		if fixup.location == "" {
			// is local ref
			_, ok := classKinds[fixup.toClass]
			if !ok {
				panic(fmt.Sprintf("Unknown class '%s'", fixup.toClass))
			}

			patch = &models.PatchDocument{
				Op:   &op,
				Path: &path,
				Value: map[string]interface{}{
					"$cref": fmt.Sprintf("weaviate://localhost/things/%s", idMap[fixup.toId]),
				},
			}
		} else {
			// is network ref
			patch = &models.PatchDocument{
				Op:   &op,
				Path: &path,
				Value: map[string]interface{}{
					"$cref": fmt.Sprintf("weaviate://%s/things/%s", fixup.location, fixup.toId),
				},
			}
		}

		assertPatchThing(idMap[fixup.fromId], patch)
		fmt.Printf("Patched thing %s\n", idMap[fixup.fromId])
	}

	for _, fixups := range thingManyFixups {
		var patch *models.PatchDocument
		path := fmt.Sprintf("/schema/%s", fixups[0].fromProperty)

		patch = &models.PatchDocument{
			Op:    &op,
			Path:  &path,
			Value: []map[string]interface{}{},
		}

		for _, fixup := range fixups {
			if fixup.location == "" {
				// is local ref
				_, ok := classKinds[fixup.toClass]
				if !ok {
					panic(fmt.Sprintf("Unknown class '%s'", fixup.toClass))
				}

				patch.Value = append(patch.Value.([]map[string]interface{}), map[string]interface{}{
					"$cref": fmt.Sprintf("weaviate://localhost/things/%s", idMap[fixup.toId]),
				})
			} else {
				patch.Value = append(patch.Value.([]map[string]interface{}), map[string]interface{}{
					"$cref": fmt.Sprintf("weaviate://%s/things/%s", fixup.location, fixup.toId),
				})
			}
		}

		assertPatchThing(idMap[fixups[0].fromId], patch)
		fmt.Printf("Patched thing %s\n", idMap[fixups[0].fromId])
	}
}

func checkThingExists(id string) bool {
	params := things.NewWeaviateThingsGetParams().WithThingID(strfmt.UUID(id))
	resp, err := client.Things.WeaviateThingsGet(params)

	if err != nil {
		switch v := err.(type) {
		case *things.WeaviateThingsGetNotFound:
			return false
		default:
			panic(fmt.Sprintf("Can't create thing %#v, because %#v", resp, spew.Sdump(v)))
		}
	}

	return true
}

func assertCreateThing(t *models.ThingCreate) *models.ThingGetResponse {
	params := things.NewWeaviateThingsCreateParams().WithBody(things.WeaviateThingsCreateBody{Thing: t})

	resp, _, err := client.Things.WeaviateThingsCreate(params)

	if err != nil {
		switch v := err.(type) {
		case *things.WeaviateThingsCreateUnprocessableEntity:
			panic(fmt.Sprintf("Can't create thing %#v, because %s", t, joinErrorMessages(v.Payload)))
		default:
			panic(fmt.Sprintf("Can't create thing %#v, because %#v", t, spew.Sdump(err)))
		}
	}

	return resp.Payload
}

func assertUpdateThing(id string, update *models.ThingUpdate) *models.ThingGetResponse {
	params := things.NewWeaviateThingsUpdateParams().WithBody(update).WithThingID(strfmt.UUID(id))

	resp, err := client.Things.WeaviateThingsUpdate(params)

	if err != nil {
		switch v := err.(type) {
		case *things.WeaviateThingsUpdateNotFound:
			panic(fmt.Sprintf("Can't patch thing with %s, because thing cannot be found", spew.Sdump(update)))
		case *things.WeaviateThingsUpdateUnprocessableEntity:
			panic(fmt.Sprintf("Can't patch thing, because %s (patch: %#v)", joinErrorMessages(v.Payload), *update))
		default:
			_ = v
			panic(fmt.Sprintf("Can't patch thing with %#v, because %#v", update, spew.Sdump(err)))
		}
	}

	return resp.Payload
}

func assertPatchThing(id string, p *models.PatchDocument) *models.ThingGetResponse {
	params := things.NewWeaviateThingsPatchParams().WithBody([]*models.PatchDocument{p}).WithThingID(strfmt.UUID(id))

	resp, _, err := client.Things.WeaviateThingsPatch(params)

	if err != nil {
		switch v := err.(type) {
		case *things.WeaviateThingsPatchNotFound:
			panic(fmt.Sprintf("Can't patch thing with %s, because thing cannot be found", spew.Sdump(p)))
		case *things.WeaviateThingsPatchUnprocessableEntity:
			panic(fmt.Sprintf("Can't patch thing, because %s", joinErrorMessages(v.Payload)))
		default:
			_ = v
			panic(fmt.Sprintf("Can't patch thing with %#v, because %#v", p, spew.Sdump(err)))
		}
	}

	return resp.Payload
}
