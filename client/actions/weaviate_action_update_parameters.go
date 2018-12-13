// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// NewWeaviateActionUpdateParams creates a new WeaviateActionUpdateParams object
// with the default values initialized.
func NewWeaviateActionUpdateParams() *WeaviateActionUpdateParams {
	var ()
	return &WeaviateActionUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateActionUpdateParamsWithTimeout creates a new WeaviateActionUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateActionUpdateParamsWithTimeout(timeout time.Duration) *WeaviateActionUpdateParams {
	var ()
	return &WeaviateActionUpdateParams{

		timeout: timeout,
	}
}

// NewWeaviateActionUpdateParamsWithContext creates a new WeaviateActionUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateActionUpdateParamsWithContext(ctx context.Context) *WeaviateActionUpdateParams {
	var ()
	return &WeaviateActionUpdateParams{

		Context: ctx,
	}
}

// NewWeaviateActionUpdateParamsWithHTTPClient creates a new WeaviateActionUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateActionUpdateParamsWithHTTPClient(client *http.Client) *WeaviateActionUpdateParams {
	var ()
	return &WeaviateActionUpdateParams{
		HTTPClient: client,
	}
}

/*WeaviateActionUpdateParams contains all the parameters to send to the API endpoint
for the weaviate action update operation typically these are written to a http.Request
*/
type WeaviateActionUpdateParams struct {

	/*ActionID
	  Unique ID of the Action.

	*/
	ActionID strfmt.UUID
	/*Body*/
	Body *models.ActionUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate action update params
func (o *WeaviateActionUpdateParams) WithTimeout(timeout time.Duration) *WeaviateActionUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate action update params
func (o *WeaviateActionUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate action update params
func (o *WeaviateActionUpdateParams) WithContext(ctx context.Context) *WeaviateActionUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate action update params
func (o *WeaviateActionUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate action update params
func (o *WeaviateActionUpdateParams) WithHTTPClient(client *http.Client) *WeaviateActionUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate action update params
func (o *WeaviateActionUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the weaviate action update params
func (o *WeaviateActionUpdateParams) WithActionID(actionID strfmt.UUID) *WeaviateActionUpdateParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the weaviate action update params
func (o *WeaviateActionUpdateParams) SetActionID(actionID strfmt.UUID) {
	o.ActionID = actionID
}

// WithBody adds the body to the weaviate action update params
func (o *WeaviateActionUpdateParams) WithBody(body *models.ActionUpdate) *WeaviateActionUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the weaviate action update params
func (o *WeaviateActionUpdateParams) SetBody(body *models.ActionUpdate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateActionUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param actionId
	if err := r.SetPathParam("actionId", o.ActionID.String()); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
