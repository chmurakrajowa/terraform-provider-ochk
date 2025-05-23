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
)

// NewDeleteNetworkCustomServicesServiceIDParams creates a new DeleteNetworkCustomServicesServiceIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNetworkCustomServicesServiceIDParams() *DeleteNetworkCustomServicesServiceIDParams {
	return &DeleteNetworkCustomServicesServiceIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNetworkCustomServicesServiceIDParamsWithTimeout creates a new DeleteNetworkCustomServicesServiceIDParams object
// with the ability to set a timeout on a request.
func NewDeleteNetworkCustomServicesServiceIDParamsWithTimeout(timeout time.Duration) *DeleteNetworkCustomServicesServiceIDParams {
	return &DeleteNetworkCustomServicesServiceIDParams{
		timeout: timeout,
	}
}

// NewDeleteNetworkCustomServicesServiceIDParamsWithContext creates a new DeleteNetworkCustomServicesServiceIDParams object
// with the ability to set a context for a request.
func NewDeleteNetworkCustomServicesServiceIDParamsWithContext(ctx context.Context) *DeleteNetworkCustomServicesServiceIDParams {
	return &DeleteNetworkCustomServicesServiceIDParams{
		Context: ctx,
	}
}

// NewDeleteNetworkCustomServicesServiceIDParamsWithHTTPClient creates a new DeleteNetworkCustomServicesServiceIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNetworkCustomServicesServiceIDParamsWithHTTPClient(client *http.Client) *DeleteNetworkCustomServicesServiceIDParams {
	return &DeleteNetworkCustomServicesServiceIDParams{
		HTTPClient: client,
	}
}

/*
DeleteNetworkCustomServicesServiceIDParams contains all the parameters to send to the API endpoint

	for the delete network custom services service ID operation.

	Typically these are written to a http.Request.
*/
type DeleteNetworkCustomServicesServiceIDParams struct {

	// ServiceID.
	//
	// Format: uuid
	ServiceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete network custom services service ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkCustomServicesServiceIDParams) WithDefaults() *DeleteNetworkCustomServicesServiceIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete network custom services service ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkCustomServicesServiceIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete network custom services service ID params
func (o *DeleteNetworkCustomServicesServiceIDParams) WithTimeout(timeout time.Duration) *DeleteNetworkCustomServicesServiceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete network custom services service ID params
func (o *DeleteNetworkCustomServicesServiceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete network custom services service ID params
func (o *DeleteNetworkCustomServicesServiceIDParams) WithContext(ctx context.Context) *DeleteNetworkCustomServicesServiceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete network custom services service ID params
func (o *DeleteNetworkCustomServicesServiceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete network custom services service ID params
func (o *DeleteNetworkCustomServicesServiceIDParams) WithHTTPClient(client *http.Client) *DeleteNetworkCustomServicesServiceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete network custom services service ID params
func (o *DeleteNetworkCustomServicesServiceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceID adds the serviceID to the delete network custom services service ID params
func (o *DeleteNetworkCustomServicesServiceIDParams) WithServiceID(serviceID strfmt.UUID) *DeleteNetworkCustomServicesServiceIDParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the delete network custom services service ID params
func (o *DeleteNetworkCustomServicesServiceIDParams) SetServiceID(serviceID strfmt.UUID) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNetworkCustomServicesServiceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serviceId
	if err := r.SetPathParam("serviceId", o.ServiceID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
