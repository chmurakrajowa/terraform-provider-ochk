// Code generated by go-swagger; DO NOT EDIT.

package port_forwarding

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

// NewGetNetworkFloatingIpsFloatingIPIDPortForwardingsParams creates a new GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkFloatingIpsFloatingIPIDPortForwardingsParams() *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams {
	return &GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkFloatingIpsFloatingIPIDPortForwardingsParamsWithTimeout creates a new GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams object
// with the ability to set a timeout on a request.
func NewGetNetworkFloatingIpsFloatingIPIDPortForwardingsParamsWithTimeout(timeout time.Duration) *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams {
	return &GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams{
		timeout: timeout,
	}
}

// NewGetNetworkFloatingIpsFloatingIPIDPortForwardingsParamsWithContext creates a new GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams object
// with the ability to set a context for a request.
func NewGetNetworkFloatingIpsFloatingIPIDPortForwardingsParamsWithContext(ctx context.Context) *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams {
	return &GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams{
		Context: ctx,
	}
}

// NewGetNetworkFloatingIpsFloatingIPIDPortForwardingsParamsWithHTTPClient creates a new GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkFloatingIpsFloatingIPIDPortForwardingsParamsWithHTTPClient(client *http.Client) *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams {
	return &GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams{
		HTTPClient: client,
	}
}

/*
GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams contains all the parameters to send to the API endpoint

	for the get network floating ips floating IP ID port forwardings operation.

	Typically these are written to a http.Request.
*/
type GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams struct {

	// FloatingIPID.
	//
	// Format: uuid
	FloatingIPID strfmt.UUID

	// Name.
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network floating ips floating IP ID port forwardings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams) WithDefaults() *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network floating ips floating IP ID port forwardings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network floating ips floating IP ID port forwardings params
func (o *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams) WithTimeout(timeout time.Duration) *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network floating ips floating IP ID port forwardings params
func (o *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network floating ips floating IP ID port forwardings params
func (o *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams) WithContext(ctx context.Context) *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network floating ips floating IP ID port forwardings params
func (o *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network floating ips floating IP ID port forwardings params
func (o *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams) WithHTTPClient(client *http.Client) *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network floating ips floating IP ID port forwardings params
func (o *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFloatingIPID adds the floatingIPID to the get network floating ips floating IP ID port forwardings params
func (o *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams) WithFloatingIPID(floatingIPID strfmt.UUID) *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams {
	o.SetFloatingIPID(floatingIPID)
	return o
}

// SetFloatingIPID adds the floatingIpId to the get network floating ips floating IP ID port forwardings params
func (o *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams) SetFloatingIPID(floatingIPID strfmt.UUID) {
	o.FloatingIPID = floatingIPID
}

// WithName adds the name to the get network floating ips floating IP ID port forwardings params
func (o *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams) WithName(name *string) *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get network floating ips floating IP ID port forwardings params
func (o *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param floatingIpId
	if err := r.SetPathParam("floatingIpId", o.FloatingIPID.String()); err != nil {
		return err
	}

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
