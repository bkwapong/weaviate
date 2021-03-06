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

package network_get

import (
	"fmt"
	"regexp"

	"github.com/creativesoftwarefdn/weaviate/graphqlapi/network/common"
	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/graphql-go/graphql"
)

// Params ties a SubQuery and a single instance
// together
type Params struct {
	SubQuery       common.SubQuery
	TargetInstance string
}

type Resolver interface {
	ProxyGetInstance(info Params) (*models.GraphQLResponse, error)
}

// FiltersAndResolver is a helper tuple to bubble data through the resolvers.
type FiltersAndResolver struct {
	Resolver Resolver
}

func NetworkGetInstanceResolve(p graphql.ResolveParams) (interface{}, error) {
	filterAndResolver, ok := p.Source.(FiltersAndResolver)
	if !ok {
		return nil, fmt.Errorf("expected source to be a FilterAndResolver, but was \n%#v",
			p.Source)
	}

	resolver := filterAndResolver.Resolver
	astLoc := p.Info.FieldASTs[0].GetLoc()
	rawSubQuery := astLoc.Source.Body[astLoc.Start:astLoc.End]
	subQueryWithoutInstance, err := replaceInstanceName(p.Info.FieldName, rawSubQuery)
	if err != nil {
		return nil, fmt.Errorf("could not replace instance name in sub-query: %s", err)
	}

	params := Params{
		SubQuery:       common.ParseSubQuery(subQueryWithoutInstance),
		TargetInstance: p.Info.FieldName,
	}

	graphQLResponse, err := resolver.ProxyGetInstance(params)
	if err != nil {
		return nil, fmt.Errorf("could not proxy to remote instance: %s", err)
	}

	local, ok := graphQLResponse.Data["Local"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected response.data.Local to be map[string]interface{}, but response was %#v",
			graphQLResponse.Data["Local"])
	}

	return local["Get"], nil
}

func replaceInstanceName(instanceName string, query []byte) ([]byte, error) {
	r, err := regexp.Compile(fmt.Sprintf(`^%s\s*`, instanceName))
	if err != nil {
		return []byte{}, err
	}

	return r.ReplaceAll(query, []byte("Get ")), nil
}
