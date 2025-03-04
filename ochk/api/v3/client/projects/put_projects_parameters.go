// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewPutProjectsParams creates a new PutProjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutProjectsParams() *PutProjectsParams {
	return &PutProjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutProjectsParamsWithTimeout creates a new PutProjectsParams object
// with the ability to set a timeout on a request.
func NewPutProjectsParamsWithTimeout(timeout time.Duration) *PutProjectsParams {
	return &PutProjectsParams{
		timeout: timeout,
	}
}

// NewPutProjectsParamsWithContext creates a new PutProjectsParams object
// with the ability to set a context for a request.
func NewPutProjectsParamsWithContext(ctx context.Context) *PutProjectsParams {
	return &PutProjectsParams{
		Context: ctx,
	}
}

// NewPutProjectsParamsWithHTTPClient creates a new PutProjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutProjectsParamsWithHTTPClient(client *http.Client) *PutProjectsParams {
	return &PutProjectsParams{
		HTTPClient: client,
	}
}

/*
PutProjectsParams contains all the parameters to send to the API endpoint

	for the put projects operation.

	Typically these are written to a http.Request.
*/
type PutProjectsParams struct {

	// Body.
	Body *models.ProjectInstance

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutProjectsParams) WithDefaults() *PutProjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutProjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put projects params
func (o *PutProjectsParams) WithTimeout(timeout time.Duration) *PutProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put projects params
func (o *PutProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put projects params
func (o *PutProjectsParams) WithContext(ctx context.Context) *PutProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put projects params
func (o *PutProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put projects params
func (o *PutProjectsParams) WithHTTPClient(client *http.Client) *PutProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put projects params
func (o *PutProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put projects params
func (o *PutProjectsParams) WithBody(body *models.ProjectInstance) *PutProjectsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put projects params
func (o *PutProjectsParams) SetBody(body *models.ProjectInstance) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
