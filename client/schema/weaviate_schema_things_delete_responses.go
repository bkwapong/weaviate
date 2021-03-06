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
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateSchemaThingsDeleteReader is a Reader for the WeaviateSchemaThingsDelete structure.
type WeaviateSchemaThingsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateSchemaThingsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateSchemaThingsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewWeaviateSchemaThingsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewWeaviateSchemaThingsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateSchemaThingsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateSchemaThingsDeleteOK creates a WeaviateSchemaThingsDeleteOK with default headers values
func NewWeaviateSchemaThingsDeleteOK() *WeaviateSchemaThingsDeleteOK {
	return &WeaviateSchemaThingsDeleteOK{}
}

/*WeaviateSchemaThingsDeleteOK handles this case with default header values.

Removed the Thing class from the ontology.
*/
type WeaviateSchemaThingsDeleteOK struct {
}

func (o *WeaviateSchemaThingsDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /schema/things/{className}][%d] weaviateSchemaThingsDeleteOK ", 200)
}

func (o *WeaviateSchemaThingsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaThingsDeleteBadRequest creates a WeaviateSchemaThingsDeleteBadRequest with default headers values
func NewWeaviateSchemaThingsDeleteBadRequest() *WeaviateSchemaThingsDeleteBadRequest {
	return &WeaviateSchemaThingsDeleteBadRequest{}
}

/*WeaviateSchemaThingsDeleteBadRequest handles this case with default header values.

Could not delete the Thing class.
*/
type WeaviateSchemaThingsDeleteBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaThingsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /schema/things/{className}][%d] weaviateSchemaThingsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *WeaviateSchemaThingsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateSchemaThingsDeleteUnauthorized creates a WeaviateSchemaThingsDeleteUnauthorized with default headers values
func NewWeaviateSchemaThingsDeleteUnauthorized() *WeaviateSchemaThingsDeleteUnauthorized {
	return &WeaviateSchemaThingsDeleteUnauthorized{}
}

/*WeaviateSchemaThingsDeleteUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateSchemaThingsDeleteUnauthorized struct {
}

func (o *WeaviateSchemaThingsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /schema/things/{className}][%d] weaviateSchemaThingsDeleteUnauthorized ", 401)
}

func (o *WeaviateSchemaThingsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaThingsDeleteInternalServerError creates a WeaviateSchemaThingsDeleteInternalServerError with default headers values
func NewWeaviateSchemaThingsDeleteInternalServerError() *WeaviateSchemaThingsDeleteInternalServerError {
	return &WeaviateSchemaThingsDeleteInternalServerError{}
}

/*WeaviateSchemaThingsDeleteInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateSchemaThingsDeleteInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaThingsDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /schema/things/{className}][%d] weaviateSchemaThingsDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateSchemaThingsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
