// Code generated by go-swagger; DO NOT EDIT.

package gateway_policies

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

// NewGatewayPolicyGetUsingGETParams creates a new GatewayPolicyGetUsingGETParams object
// with the default values initialized.
func NewGatewayPolicyGetUsingGETParams() *GatewayPolicyGetUsingGETParams {
	var ()
	return &GatewayPolicyGetUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGatewayPolicyGetUsingGETParamsWithTimeout creates a new GatewayPolicyGetUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGatewayPolicyGetUsingGETParamsWithTimeout(timeout time.Duration) *GatewayPolicyGetUsingGETParams {
	var ()
	return &GatewayPolicyGetUsingGETParams{

		timeout: timeout,
	}
}

// NewGatewayPolicyGetUsingGETParamsWithContext creates a new GatewayPolicyGetUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGatewayPolicyGetUsingGETParamsWithContext(ctx context.Context) *GatewayPolicyGetUsingGETParams {
	var ()
	return &GatewayPolicyGetUsingGETParams{

		Context: ctx,
	}
}

// NewGatewayPolicyGetUsingGETParamsWithHTTPClient creates a new GatewayPolicyGetUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGatewayPolicyGetUsingGETParamsWithHTTPClient(client *http.Client) *GatewayPolicyGetUsingGETParams {
	var ()
	return &GatewayPolicyGetUsingGETParams{
		HTTPClient: client,
	}
}

/*GatewayPolicyGetUsingGETParams contains all the parameters to send to the API endpoint
for the gateway policy get using g e t operation typically these are written to a http.Request
*/
type GatewayPolicyGetUsingGETParams struct {

	/*GatewayPolicyID
	  gatewayPolicyId

	*/
	GatewayPolicyID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the gateway policy get using g e t params
func (o *GatewayPolicyGetUsingGETParams) WithTimeout(timeout time.Duration) *GatewayPolicyGetUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gateway policy get using g e t params
func (o *GatewayPolicyGetUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gateway policy get using g e t params
func (o *GatewayPolicyGetUsingGETParams) WithContext(ctx context.Context) *GatewayPolicyGetUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gateway policy get using g e t params
func (o *GatewayPolicyGetUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gateway policy get using g e t params
func (o *GatewayPolicyGetUsingGETParams) WithHTTPClient(client *http.Client) *GatewayPolicyGetUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gateway policy get using g e t params
func (o *GatewayPolicyGetUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayPolicyID adds the gatewayPolicyID to the gateway policy get using g e t params
func (o *GatewayPolicyGetUsingGETParams) WithGatewayPolicyID(gatewayPolicyID string) *GatewayPolicyGetUsingGETParams {
	o.SetGatewayPolicyID(gatewayPolicyID)
	return o
}

// SetGatewayPolicyID adds the gatewayPolicyId to the gateway policy get using g e t params
func (o *GatewayPolicyGetUsingGETParams) SetGatewayPolicyID(gatewayPolicyID string) {
	o.GatewayPolicyID = gatewayPolicyID
}

// WriteToRequest writes these params to a swagger request
func (o *GatewayPolicyGetUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gatewayPolicyId
	if err := r.SetPathParam("gatewayPolicyId", o.GatewayPolicyID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
