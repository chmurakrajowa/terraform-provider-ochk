// Code generated by go-swagger; DO NOT EDIT.

package key_schedule

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

// NewPutKmsScheduleParams creates a new PutKmsScheduleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutKmsScheduleParams() *PutKmsScheduleParams {
	return &PutKmsScheduleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutKmsScheduleParamsWithTimeout creates a new PutKmsScheduleParams object
// with the ability to set a timeout on a request.
func NewPutKmsScheduleParamsWithTimeout(timeout time.Duration) *PutKmsScheduleParams {
	return &PutKmsScheduleParams{
		timeout: timeout,
	}
}

// NewPutKmsScheduleParamsWithContext creates a new PutKmsScheduleParams object
// with the ability to set a context for a request.
func NewPutKmsScheduleParamsWithContext(ctx context.Context) *PutKmsScheduleParams {
	return &PutKmsScheduleParams{
		Context: ctx,
	}
}

// NewPutKmsScheduleParamsWithHTTPClient creates a new PutKmsScheduleParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutKmsScheduleParamsWithHTTPClient(client *http.Client) *PutKmsScheduleParams {
	return &PutKmsScheduleParams{
		HTTPClient: client,
	}
}

/*
PutKmsScheduleParams contains all the parameters to send to the API endpoint

	for the put kms schedule operation.

	Typically these are written to a http.Request.
*/
type PutKmsScheduleParams struct {

	// Body.
	Body *models.KeyRotationSchedule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put kms schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutKmsScheduleParams) WithDefaults() *PutKmsScheduleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put kms schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutKmsScheduleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put kms schedule params
func (o *PutKmsScheduleParams) WithTimeout(timeout time.Duration) *PutKmsScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put kms schedule params
func (o *PutKmsScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put kms schedule params
func (o *PutKmsScheduleParams) WithContext(ctx context.Context) *PutKmsScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put kms schedule params
func (o *PutKmsScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put kms schedule params
func (o *PutKmsScheduleParams) WithHTTPClient(client *http.Client) *PutKmsScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put kms schedule params
func (o *PutKmsScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put kms schedule params
func (o *PutKmsScheduleParams) WithBody(body *models.KeyRotationSchedule) *PutKmsScheduleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put kms schedule params
func (o *PutKmsScheduleParams) SetBody(body *models.KeyRotationSchedule) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutKmsScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
