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

// NewGatewayPolicyListUsingGETParams creates a new GatewayPolicyListUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGatewayPolicyListUsingGETParams() *GatewayPolicyListUsingGETParams {
	return &GatewayPolicyListUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGatewayPolicyListUsingGETParamsWithTimeout creates a new GatewayPolicyListUsingGETParams object
// with the ability to set a timeout on a request.
func NewGatewayPolicyListUsingGETParamsWithTimeout(timeout time.Duration) *GatewayPolicyListUsingGETParams {
	return &GatewayPolicyListUsingGETParams{
		timeout: timeout,
	}
}

// NewGatewayPolicyListUsingGETParamsWithContext creates a new GatewayPolicyListUsingGETParams object
// with the ability to set a context for a request.
func NewGatewayPolicyListUsingGETParamsWithContext(ctx context.Context) *GatewayPolicyListUsingGETParams {
	return &GatewayPolicyListUsingGETParams{
		Context: ctx,
	}
}

// NewGatewayPolicyListUsingGETParamsWithHTTPClient creates a new GatewayPolicyListUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGatewayPolicyListUsingGETParamsWithHTTPClient(client *http.Client) *GatewayPolicyListUsingGETParams {
	return &GatewayPolicyListUsingGETParams{
		HTTPClient: client,
	}
}

/* GatewayPolicyListUsingGETParams contains all the parameters to send to the API endpoint
   for the gateway policy list using g e t operation.

   Typically these are written to a http.Request.
*/
type GatewayPolicyListUsingGETParams struct {

	/* DisplayName.

	   displayName
	*/
	DisplayName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gateway policy list using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GatewayPolicyListUsingGETParams) WithDefaults() *GatewayPolicyListUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gateway policy list using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GatewayPolicyListUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the gateway policy list using g e t params
func (o *GatewayPolicyListUsingGETParams) WithTimeout(timeout time.Duration) *GatewayPolicyListUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gateway policy list using g e t params
func (o *GatewayPolicyListUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gateway policy list using g e t params
func (o *GatewayPolicyListUsingGETParams) WithContext(ctx context.Context) *GatewayPolicyListUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gateway policy list using g e t params
func (o *GatewayPolicyListUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gateway policy list using g e t params
func (o *GatewayPolicyListUsingGETParams) WithHTTPClient(client *http.Client) *GatewayPolicyListUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gateway policy list using g e t params
func (o *GatewayPolicyListUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisplayName adds the displayName to the gateway policy list using g e t params
func (o *GatewayPolicyListUsingGETParams) WithDisplayName(displayName *string) *GatewayPolicyListUsingGETParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the gateway policy list using g e t params
func (o *GatewayPolicyListUsingGETParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WriteToRequest writes these params to a swagger request
func (o *GatewayPolicyListUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DisplayName != nil {

		// query param displayName
		var qrDisplayName string

		if o.DisplayName != nil {
			qrDisplayName = *o.DisplayName
		}
		qDisplayName := qrDisplayName
		if qDisplayName != "" {

			if err := r.SetQueryParam("displayName", qDisplayName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
