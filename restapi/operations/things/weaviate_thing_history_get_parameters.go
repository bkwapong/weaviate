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
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWeaviateThingHistoryGetParams creates a new WeaviateThingHistoryGetParams object
// no default values defined in spec.
func NewWeaviateThingHistoryGetParams() WeaviateThingHistoryGetParams {

	return WeaviateThingHistoryGetParams{}
}

// WeaviateThingHistoryGetParams contains all the bound params for the weaviate thing history get operation
// typically these are obtained from a http.Request
//
// swagger:parameters weaviate.thing.history.get
type WeaviateThingHistoryGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Unique ID of the Thing.
	  Required: true
	  In: path
	*/
	ThingID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewWeaviateThingHistoryGetParams() beforehand.
func (o *WeaviateThingHistoryGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rThingID, rhkThingID, _ := route.Params.GetOK("thingId")
	if err := o.bindThingID(rThingID, rhkThingID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindThingID binds and validates parameter ThingID from path.
func (o *WeaviateThingHistoryGetParams) bindThingID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *WeaviateThingHistoryGetParams) validateThingID(formats strfmt.Registry) error {

	if err := validate.FormatOf("thingId", "path", "uuid", o.ThingID.String(), formats); err != nil {
		return err
	}
	return nil
}
