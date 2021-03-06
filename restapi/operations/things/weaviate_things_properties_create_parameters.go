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
 */ // Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// NewWeaviateThingsPropertiesCreateParams creates a new WeaviateThingsPropertiesCreateParams object
// no default values defined in spec.
func NewWeaviateThingsPropertiesCreateParams() WeaviateThingsPropertiesCreateParams {

	return WeaviateThingsPropertiesCreateParams{}
}

// WeaviateThingsPropertiesCreateParams contains all the bound params for the weaviate things properties create operation
// typically these are obtained from a http.Request
//
// swagger:parameters weaviate.things.properties.create
type WeaviateThingsPropertiesCreateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Body *models.SingleRef
	/*Unique name of the property related to the Thing.
	  Required: true
	  In: path
	*/
	PropertyName string
	/*Unique ID of the Thing.
	  Required: true
	  In: path
	*/
	ThingID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewWeaviateThingsPropertiesCreateParams() beforehand.
func (o *WeaviateThingsPropertiesCreateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.SingleRef
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body"))
	}
	rPropertyName, rhkPropertyName, _ := route.Params.GetOK("propertyName")
	if err := o.bindPropertyName(rPropertyName, rhkPropertyName, route.Formats); err != nil {
		res = append(res, err)
	}

	rThingID, rhkThingID, _ := route.Params.GetOK("thingId")
	if err := o.bindThingID(rThingID, rhkThingID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPropertyName binds and validates parameter PropertyName from path.
func (o *WeaviateThingsPropertiesCreateParams) bindPropertyName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.PropertyName = raw

	return nil
}

// bindThingID binds and validates parameter ThingID from path.
func (o *WeaviateThingsPropertiesCreateParams) bindThingID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("thingId", "path", "strfmt.UUID", raw)
	}
	o.ThingID = *(value.(*strfmt.UUID))

	if err := o.validateThingID(formats); err != nil {
		return err
	}

	return nil
}

// validateThingID carries on validations for parameter ThingID
func (o *WeaviateThingsPropertiesCreateParams) validateThingID(formats strfmt.Registry) error {

	if err := validate.FormatOf("thingId", "path", "uuid", o.ThingID.String(), formats); err != nil {
		return err
	}
	return nil
}
