// Code generated by go-swagger; DO NOT EDIT.

package custom_services

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
)

// NewCustomServiceDeleteUsingDELETEParams creates a new CustomServiceDeleteUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomServiceDeleteUsingDELETEParams() *CustomServiceDeleteUsingDELETEParams {
	return &CustomServiceDeleteUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomServiceDeleteUsingDELETEParamsWithTimeout creates a new CustomServiceDeleteUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewCustomServiceDeleteUsingDELETEParamsWithTimeout(timeout time.Duration) *CustomServiceDeleteUsingDELETEParams {
	return &CustomServiceDeleteUsingDELETEParams{
		timeout: timeout,
	}
}

// NewCustomServiceDeleteUsingDELETEParamsWithContext creates a new CustomServiceDeleteUsingDELETEParams object
// with the ability to set a context for a request.
func NewCustomServiceDeleteUsingDELETEParamsWithContext(ctx context.Context) *CustomServiceDeleteUsingDELETEParams {
	return &CustomServiceDeleteUsingDELETEParams{
		Context: ctx,
	}
}

// NewCustomServiceDeleteUsingDELETEParamsWithHTTPClient creates a new CustomServiceDeleteUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomServiceDeleteUsingDELETEParamsWithHTTPClient(client *http.Client) *CustomServiceDeleteUsingDELETEParams {
	return &CustomServiceDeleteUsingDELETEParams{
		HTTPClient: client,
	}
}

/*
CustomServiceDeleteUsingDELETEParams contains all the parameters to send to the API endpoint

	for the custom service delete using d e l e t e operation.

	Typically these are written to a http.Request.
*/
type CustomServiceDeleteUsingDELETEParams struct {

	/* Force.

	   force
	*/
	Force *bool

	/* ServiceID.

	   serviceId
	*/
	ServiceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the custom service delete using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomServiceDeleteUsingDELETEParams) WithDefaults() *CustomServiceDeleteUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the custom service delete using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomServiceDeleteUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the custom service delete using d e l e t e params
func (o *CustomServiceDeleteUsingDELETEParams) WithTimeout(timeout time.Duration) *CustomServiceDeleteUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom service delete using d e l e t e params
func (o *CustomServiceDeleteUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom service delete using d e l e t e params
func (o *CustomServiceDeleteUsingDELETEParams) WithContext(ctx context.Context) *CustomServiceDeleteUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom service delete using d e l e t e params
func (o *CustomServiceDeleteUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom service delete using d e l e t e params
func (o *CustomServiceDeleteUsingDELETEParams) WithHTTPClient(client *http.Client) *CustomServiceDeleteUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom service delete using d e l e t e params
func (o *CustomServiceDeleteUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the custom service delete using d e l e t e params
func (o *CustomServiceDeleteUsingDELETEParams) WithForce(force *bool) *CustomServiceDeleteUsingDELETEParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the custom service delete using d e l e t e params
func (o *CustomServiceDeleteUsingDELETEParams) SetForce(force *bool) {
	o.Force = force
}

// WithServiceID adds the serviceID to the custom service delete using d e l e t e params
func (o *CustomServiceDeleteUsingDELETEParams) WithServiceID(serviceID string) *CustomServiceDeleteUsingDELETEParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the custom service delete using d e l e t e params
func (o *CustomServiceDeleteUsingDELETEParams) SetServiceID(serviceID string) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomServiceDeleteUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
	}

	// path param serviceId
	if err := r.SetPathParam("serviceId", o.ServiceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
