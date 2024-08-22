// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
)

// NewPutTagsTagIDParams creates a new PutTagsTagIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutTagsTagIDParams() *PutTagsTagIDParams {
	return &PutTagsTagIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutTagsTagIDParamsWithTimeout creates a new PutTagsTagIDParams object
// with the ability to set a timeout on a request.
func NewPutTagsTagIDParamsWithTimeout(timeout time.Duration) *PutTagsTagIDParams {
	return &PutTagsTagIDParams{
		timeout: timeout,
	}
}

// NewPutTagsTagIDParamsWithContext creates a new PutTagsTagIDParams object
// with the ability to set a context for a request.
func NewPutTagsTagIDParamsWithContext(ctx context.Context) *PutTagsTagIDParams {
	return &PutTagsTagIDParams{
		Context: ctx,
	}
}

// NewPutTagsTagIDParamsWithHTTPClient creates a new PutTagsTagIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutTagsTagIDParamsWithHTTPClient(client *http.Client) *PutTagsTagIDParams {
	return &PutTagsTagIDParams{
		HTTPClient: client,
	}
}

/*
PutTagsTagIDParams contains all the parameters to send to the API endpoint

	for the put tags tag ID operation.

	Typically these are written to a http.Request.
*/
type PutTagsTagIDParams struct {

	// Body.
	Body *models.Tag

	// TagID.
	//
	// Format: int32
	TagID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put tags tag ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutTagsTagIDParams) WithDefaults() *PutTagsTagIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put tags tag ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutTagsTagIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put tags tag ID params
func (o *PutTagsTagIDParams) WithTimeout(timeout time.Duration) *PutTagsTagIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put tags tag ID params
func (o *PutTagsTagIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put tags tag ID params
func (o *PutTagsTagIDParams) WithContext(ctx context.Context) *PutTagsTagIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put tags tag ID params
func (o *PutTagsTagIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put tags tag ID params
func (o *PutTagsTagIDParams) WithHTTPClient(client *http.Client) *PutTagsTagIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put tags tag ID params
func (o *PutTagsTagIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put tags tag ID params
func (o *PutTagsTagIDParams) WithBody(body *models.Tag) *PutTagsTagIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put tags tag ID params
func (o *PutTagsTagIDParams) SetBody(body *models.Tag) {
	o.Body = body
}

// WithTagID adds the tagID to the put tags tag ID params
func (o *PutTagsTagIDParams) WithTagID(tagID int32) *PutTagsTagIDParams {
	o.SetTagID(tagID)
	return o
}

// SetTagID adds the tagId to the put tags tag ID params
func (o *PutTagsTagIDParams) SetTagID(tagID int32) {
	o.TagID = tagID
}

// WriteToRequest writes these params to a swagger request
func (o *PutTagsTagIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param tagId
	if err := r.SetPathParam("tagId", swag.FormatInt32(o.TagID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
