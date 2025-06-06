// Code generated by go-swagger; DO NOT EDIT.

package floating_ip

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

// NewGetNetworkFloatingIpsParams creates a new GetNetworkFloatingIpsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkFloatingIpsParams() *GetNetworkFloatingIpsParams {
	return &GetNetworkFloatingIpsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkFloatingIpsParamsWithTimeout creates a new GetNetworkFloatingIpsParams object
// with the ability to set a timeout on a request.
func NewGetNetworkFloatingIpsParamsWithTimeout(timeout time.Duration) *GetNetworkFloatingIpsParams {
	return &GetNetworkFloatingIpsParams{
		timeout: timeout,
	}
}

// NewGetNetworkFloatingIpsParamsWithContext creates a new GetNetworkFloatingIpsParams object
// with the ability to set a context for a request.
func NewGetNetworkFloatingIpsParamsWithContext(ctx context.Context) *GetNetworkFloatingIpsParams {
	return &GetNetworkFloatingIpsParams{
		Context: ctx,
	}
}

// NewGetNetworkFloatingIpsParamsWithHTTPClient creates a new GetNetworkFloatingIpsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkFloatingIpsParamsWithHTTPClient(client *http.Client) *GetNetworkFloatingIpsParams {
	return &GetNetworkFloatingIpsParams{
		HTTPClient: client,
	}
}

/*
GetNetworkFloatingIpsParams contains all the parameters to send to the API endpoint

	for the get network floating ips operation.

	Typically these are written to a http.Request.
*/
type GetNetworkFloatingIpsParams struct {

	// Name.
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network floating ips params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkFloatingIpsParams) WithDefaults() *GetNetworkFloatingIpsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network floating ips params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkFloatingIpsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network floating ips params
func (o *GetNetworkFloatingIpsParams) WithTimeout(timeout time.Duration) *GetNetworkFloatingIpsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network floating ips params
func (o *GetNetworkFloatingIpsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network floating ips params
func (o *GetNetworkFloatingIpsParams) WithContext(ctx context.Context) *GetNetworkFloatingIpsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network floating ips params
func (o *GetNetworkFloatingIpsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network floating ips params
func (o *GetNetworkFloatingIpsParams) WithHTTPClient(client *http.Client) *GetNetworkFloatingIpsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network floating ips params
func (o *GetNetworkFloatingIpsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get network floating ips params
func (o *GetNetworkFloatingIpsParams) WithName(name *string) *GetNetworkFloatingIpsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get network floating ips params
func (o *GetNetworkFloatingIpsParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkFloatingIpsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
