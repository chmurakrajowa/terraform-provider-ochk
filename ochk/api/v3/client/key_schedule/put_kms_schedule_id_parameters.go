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

// NewPutKmsScheduleIDParams creates a new PutKmsScheduleIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutKmsScheduleIDParams() *PutKmsScheduleIDParams {
	return &PutKmsScheduleIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutKmsScheduleIDParamsWithTimeout creates a new PutKmsScheduleIDParams object
// with the ability to set a timeout on a request.
func NewPutKmsScheduleIDParamsWithTimeout(timeout time.Duration) *PutKmsScheduleIDParams {
	return &PutKmsScheduleIDParams{
		timeout: timeout,
	}
}

// NewPutKmsScheduleIDParamsWithContext creates a new PutKmsScheduleIDParams object
// with the ability to set a context for a request.
func NewPutKmsScheduleIDParamsWithContext(ctx context.Context) *PutKmsScheduleIDParams {
	return &PutKmsScheduleIDParams{
		Context: ctx,
	}
}

// NewPutKmsScheduleIDParamsWithHTTPClient creates a new PutKmsScheduleIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutKmsScheduleIDParamsWithHTTPClient(client *http.Client) *PutKmsScheduleIDParams {
	return &PutKmsScheduleIDParams{
		HTTPClient: client,
	}
}

/*
PutKmsScheduleIDParams contains all the parameters to send to the API endpoint

	for the put kms schedule ID operation.

	Typically these are written to a http.Request.
*/
type PutKmsScheduleIDParams struct {

	// Body.
	Body *models.KeyRotationSchedule

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put kms schedule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutKmsScheduleIDParams) WithDefaults() *PutKmsScheduleIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put kms schedule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutKmsScheduleIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put kms schedule ID params
func (o *PutKmsScheduleIDParams) WithTimeout(timeout time.Duration) *PutKmsScheduleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put kms schedule ID params
func (o *PutKmsScheduleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put kms schedule ID params
func (o *PutKmsScheduleIDParams) WithContext(ctx context.Context) *PutKmsScheduleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put kms schedule ID params
func (o *PutKmsScheduleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put kms schedule ID params
func (o *PutKmsScheduleIDParams) WithHTTPClient(client *http.Client) *PutKmsScheduleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put kms schedule ID params
func (o *PutKmsScheduleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put kms schedule ID params
func (o *PutKmsScheduleIDParams) WithBody(body *models.KeyRotationSchedule) *PutKmsScheduleIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put kms schedule ID params
func (o *PutKmsScheduleIDParams) SetBody(body *models.KeyRotationSchedule) {
	o.Body = body
}

// WithID adds the id to the put kms schedule ID params
func (o *PutKmsScheduleIDParams) WithID(id string) *PutKmsScheduleIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put kms schedule ID params
func (o *PutKmsScheduleIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutKmsScheduleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
