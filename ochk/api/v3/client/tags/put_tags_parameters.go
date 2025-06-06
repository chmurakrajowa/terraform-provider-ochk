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

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
)

// NewPutTagsParams creates a new PutTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutTagsParams() *PutTagsParams {
	return &PutTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutTagsParamsWithTimeout creates a new PutTagsParams object
// with the ability to set a timeout on a request.
func NewPutTagsParamsWithTimeout(timeout time.Duration) *PutTagsParams {
	return &PutTagsParams{
		timeout: timeout,
	}
}

// NewPutTagsParamsWithContext creates a new PutTagsParams object
// with the ability to set a context for a request.
func NewPutTagsParamsWithContext(ctx context.Context) *PutTagsParams {
	return &PutTagsParams{
		Context: ctx,
	}
}

// NewPutTagsParamsWithHTTPClient creates a new PutTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutTagsParamsWithHTTPClient(client *http.Client) *PutTagsParams {
	return &PutTagsParams{
		HTTPClient: client,
	}
}

/*
PutTagsParams contains all the parameters to send to the API endpoint

	for the put tags operation.

	Typically these are written to a http.Request.
*/
type PutTagsParams struct {

	// Body.
	Body *models.Tag

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutTagsParams) WithDefaults() *PutTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put tags params
func (o *PutTagsParams) WithTimeout(timeout time.Duration) *PutTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put tags params
func (o *PutTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put tags params
func (o *PutTagsParams) WithContext(ctx context.Context) *PutTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put tags params
func (o *PutTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put tags params
func (o *PutTagsParams) WithHTTPClient(client *http.Client) *PutTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put tags params
func (o *PutTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put tags params
func (o *PutTagsParams) WithBody(body *models.Tag) *PutTagsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put tags params
func (o *PutTagsParams) SetBody(body *models.Tag) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
