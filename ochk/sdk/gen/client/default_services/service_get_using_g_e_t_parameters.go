// Code generated by go-swagger; DO NOT EDIT.

package default_services

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
)

// NewServiceGetUsingGETParams creates a new ServiceGetUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServiceGetUsingGETParams() *ServiceGetUsingGETParams {
	return &ServiceGetUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServiceGetUsingGETParamsWithTimeout creates a new ServiceGetUsingGETParams object
// with the ability to set a timeout on a request.
func NewServiceGetUsingGETParamsWithTimeout(timeout time.Duration) *ServiceGetUsingGETParams {
	return &ServiceGetUsingGETParams{
		timeout: timeout,
	}
}

// NewServiceGetUsingGETParamsWithContext creates a new ServiceGetUsingGETParams object
// with the ability to set a context for a request.
func NewServiceGetUsingGETParamsWithContext(ctx context.Context) *ServiceGetUsingGETParams {
	return &ServiceGetUsingGETParams{
		Context: ctx,
	}
}

// NewServiceGetUsingGETParamsWithHTTPClient creates a new ServiceGetUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewServiceGetUsingGETParamsWithHTTPClient(client *http.Client) *ServiceGetUsingGETParams {
	return &ServiceGetUsingGETParams{
		HTTPClient: client,
	}
}

/*
ServiceGetUsingGETParams contains all the parameters to send to the API endpoint

	for the service get using g e t operation.

	Typically these are written to a http.Request.
*/
type ServiceGetUsingGETParams struct {

	/* ServiceID.

	   serviceId
	*/
	ServiceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service get using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceGetUsingGETParams) WithDefaults() *ServiceGetUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service get using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceGetUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service get using g e t params
func (o *ServiceGetUsingGETParams) WithTimeout(timeout time.Duration) *ServiceGetUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service get using g e t params
func (o *ServiceGetUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service get using g e t params
func (o *ServiceGetUsingGETParams) WithContext(ctx context.Context) *ServiceGetUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service get using g e t params
func (o *ServiceGetUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service get using g e t params
func (o *ServiceGetUsingGETParams) WithHTTPClient(client *http.Client) *ServiceGetUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service get using g e t params
func (o *ServiceGetUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceID adds the serviceID to the service get using g e t params
func (o *ServiceGetUsingGETParams) WithServiceID(serviceID string) *ServiceGetUsingGETParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the service get using g e t params
func (o *ServiceGetUsingGETParams) SetServiceID(serviceID string) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceGetUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serviceId
	if err := r.SetPathParam("serviceId", o.ServiceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
