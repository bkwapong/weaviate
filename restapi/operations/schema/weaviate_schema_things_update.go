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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateSchemaThingsUpdateHandlerFunc turns a function with the right signature into a weaviate schema things update handler
type WeaviateSchemaThingsUpdateHandlerFunc func(WeaviateSchemaThingsUpdateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateSchemaThingsUpdateHandlerFunc) Handle(params WeaviateSchemaThingsUpdateParams) middleware.Responder {
	return fn(params)
}

// WeaviateSchemaThingsUpdateHandler interface for that can handle valid weaviate schema things update params
type WeaviateSchemaThingsUpdateHandler interface {
	Handle(WeaviateSchemaThingsUpdateParams) middleware.Responder
}

// NewWeaviateSchemaThingsUpdate creates a new http.Handler for the weaviate schema things update operation
func NewWeaviateSchemaThingsUpdate(ctx *middleware.Context, handler WeaviateSchemaThingsUpdateHandler) *WeaviateSchemaThingsUpdate {
	return &WeaviateSchemaThingsUpdate{Context: ctx, Handler: handler}
}

/*WeaviateSchemaThingsUpdate swagger:route PUT /schema/things/{className} schema weaviateSchemaThingsUpdate

Rename, or replace the keywords of the Thing.

*/
type WeaviateSchemaThingsUpdate struct {
	Context *middleware.Context
	Handler WeaviateSchemaThingsUpdateHandler
}

func (o *WeaviateSchemaThingsUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateSchemaThingsUpdateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// WeaviateSchemaThingsUpdateBody weaviate schema things update body
// swagger:model WeaviateSchemaThingsUpdateBody
type WeaviateSchemaThingsUpdateBody struct {

	// keywords
	Keywords models.SemanticSchemaKeywords `json:"keywords,omitempty"`

	// The new name of the Thing.
	NewName string `json:"newName,omitempty"`
}

// Validate validates this weaviate schema things update body
func (o *WeaviateSchemaThingsUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateKeywords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WeaviateSchemaThingsUpdateBody) validateKeywords(formats strfmt.Registry) error {

	if swag.IsZero(o.Keywords) { // not required
		return nil
	}

	if err := o.Keywords.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "keywords")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *WeaviateSchemaThingsUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WeaviateSchemaThingsUpdateBody) UnmarshalBinary(b []byte) error {
	var res WeaviateSchemaThingsUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
