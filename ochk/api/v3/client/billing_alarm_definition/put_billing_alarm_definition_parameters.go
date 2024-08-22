// Code generated by go-swagger; DO NOT EDIT.

package billing_alarm_definition

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

// NewPutBillingAlarmDefinitionParams creates a new PutBillingAlarmDefinitionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutBillingAlarmDefinitionParams() *PutBillingAlarmDefinitionParams {
	return &PutBillingAlarmDefinitionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutBillingAlarmDefinitionParamsWithTimeout creates a new PutBillingAlarmDefinitionParams object
// with the ability to set a timeout on a request.
func NewPutBillingAlarmDefinitionParamsWithTimeout(timeout time.Duration) *PutBillingAlarmDefinitionParams {
	return &PutBillingAlarmDefinitionParams{
		timeout: timeout,
	}
}

// NewPutBillingAlarmDefinitionParamsWithContext creates a new PutBillingAlarmDefinitionParams object
// with the ability to set a context for a request.
func NewPutBillingAlarmDefinitionParamsWithContext(ctx context.Context) *PutBillingAlarmDefinitionParams {
	return &PutBillingAlarmDefinitionParams{
		Context: ctx,
	}
}

// NewPutBillingAlarmDefinitionParamsWithHTTPClient creates a new PutBillingAlarmDefinitionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutBillingAlarmDefinitionParamsWithHTTPClient(client *http.Client) *PutBillingAlarmDefinitionParams {
	return &PutBillingAlarmDefinitionParams{
		HTTPClient: client,
	}
}

/*
PutBillingAlarmDefinitionParams contains all the parameters to send to the API endpoint

	for the put billing alarm definition operation.

	Typically these are written to a http.Request.
*/
type PutBillingAlarmDefinitionParams struct {

	// Body.
	Body *models.BillingAlarmDefinition

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put billing alarm definition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutBillingAlarmDefinitionParams) WithDefaults() *PutBillingAlarmDefinitionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put billing alarm definition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutBillingAlarmDefinitionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put billing alarm definition params
func (o *PutBillingAlarmDefinitionParams) WithTimeout(timeout time.Duration) *PutBillingAlarmDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put billing alarm definition params
func (o *PutBillingAlarmDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put billing alarm definition params
func (o *PutBillingAlarmDefinitionParams) WithContext(ctx context.Context) *PutBillingAlarmDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put billing alarm definition params
func (o *PutBillingAlarmDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put billing alarm definition params
func (o *PutBillingAlarmDefinitionParams) WithHTTPClient(client *http.Client) *PutBillingAlarmDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put billing alarm definition params
func (o *PutBillingAlarmDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put billing alarm definition params
func (o *PutBillingAlarmDefinitionParams) WithBody(body *models.BillingAlarmDefinition) *PutBillingAlarmDefinitionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put billing alarm definition params
func (o *PutBillingAlarmDefinitionParams) SetBody(body *models.BillingAlarmDefinition) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutBillingAlarmDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
