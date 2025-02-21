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

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
)

// NewPutNetworkFloatingIpsFloatingIPIDPortForwardingsParams creates a new PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutNetworkFloatingIpsFloatingIPIDPortForwardingsParams() *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams {
	return &PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutNetworkFloatingIpsFloatingIPIDPortForwardingsParamsWithTimeout creates a new PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams object
// with the ability to set a timeout on a request.
func NewPutNetworkFloatingIpsFloatingIPIDPortForwardingsParamsWithTimeout(timeout time.Duration) *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams {
	return &PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams{
		timeout: timeout,
	}
}

// NewPutNetworkFloatingIpsFloatingIPIDPortForwardingsParamsWithContext creates a new PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams object
// with the ability to set a context for a request.
func NewPutNetworkFloatingIpsFloatingIPIDPortForwardingsParamsWithContext(ctx context.Context) *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams {
	return &PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams{
		Context: ctx,
	}
}

// NewPutNetworkFloatingIpsFloatingIPIDPortForwardingsParamsWithHTTPClient creates a new PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutNetworkFloatingIpsFloatingIPIDPortForwardingsParamsWithHTTPClient(client *http.Client) *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams {
	return &PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams{
		HTTPClient: client,
	}
}

/*
PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams contains all the parameters to send to the API endpoint

	for the put network floating ips floating IP ID port forwardings operation.

	Typically these are written to a http.Request.
*/
type PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams struct {

	// Body.
	Body *models.PortForwarding

	// FloatingIPID.
	//
	// Format: uuid
	FloatingIPID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put network floating ips floating IP ID port forwardings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams) WithDefaults() *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put network floating ips floating IP ID port forwardings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put network floating ips floating IP ID port forwardings params
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams) WithTimeout(timeout time.Duration) *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put network floating ips floating IP ID port forwardings params
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put network floating ips floating IP ID port forwardings params
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams) WithContext(ctx context.Context) *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put network floating ips floating IP ID port forwardings params
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put network floating ips floating IP ID port forwardings params
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams) WithHTTPClient(client *http.Client) *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put network floating ips floating IP ID port forwardings params
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put network floating ips floating IP ID port forwardings params
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams) WithBody(body *models.PortForwarding) *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put network floating ips floating IP ID port forwardings params
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams) SetBody(body *models.PortForwarding) {
	o.Body = body
}

// WithFloatingIPID adds the floatingIPID to the put network floating ips floating IP ID port forwardings params
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams) WithFloatingIPID(floatingIPID strfmt.UUID) *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams {
	o.SetFloatingIPID(floatingIPID)
	return o
}

// SetFloatingIPID adds the floatingIpId to the put network floating ips floating IP ID port forwardings params
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams) SetFloatingIPID(floatingIPID strfmt.UUID) {
	o.FloatingIPID = floatingIPID
}

// WriteToRequest writes these params to a swagger request
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param floatingIpId
	if err := r.SetPathParam("floatingIpId", o.FloatingIPID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
