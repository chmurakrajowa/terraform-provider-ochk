// Code generated by go-swagger; DO NOT EDIT.

package firewall_rules_s_n

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

// NewGfwRuleGetUsingGETParams creates a new GfwRuleGetUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGfwRuleGetUsingGETParams() *GfwRuleGetUsingGETParams {
	return &GfwRuleGetUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGfwRuleGetUsingGETParamsWithTimeout creates a new GfwRuleGetUsingGETParams object
// with the ability to set a timeout on a request.
func NewGfwRuleGetUsingGETParamsWithTimeout(timeout time.Duration) *GfwRuleGetUsingGETParams {
	return &GfwRuleGetUsingGETParams{
		timeout: timeout,
	}
}

// NewGfwRuleGetUsingGETParamsWithContext creates a new GfwRuleGetUsingGETParams object
// with the ability to set a context for a request.
func NewGfwRuleGetUsingGETParamsWithContext(ctx context.Context) *GfwRuleGetUsingGETParams {
	return &GfwRuleGetUsingGETParams{
		Context: ctx,
	}
}

// NewGfwRuleGetUsingGETParamsWithHTTPClient creates a new GfwRuleGetUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGfwRuleGetUsingGETParamsWithHTTPClient(client *http.Client) *GfwRuleGetUsingGETParams {
	return &GfwRuleGetUsingGETParams{
		HTTPClient: client,
	}
}

/* GfwRuleGetUsingGETParams contains all the parameters to send to the API endpoint
   for the gfw rule get using g e t operation.

   Typically these are written to a http.Request.
*/
type GfwRuleGetUsingGETParams struct {

	/* RouterID.

	   routerId
	*/
	RouterID string

	/* RuleID.

	   ruleId
	*/
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gfw rule get using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GfwRuleGetUsingGETParams) WithDefaults() *GfwRuleGetUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gfw rule get using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GfwRuleGetUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the gfw rule get using g e t params
func (o *GfwRuleGetUsingGETParams) WithTimeout(timeout time.Duration) *GfwRuleGetUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gfw rule get using g e t params
func (o *GfwRuleGetUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gfw rule get using g e t params
func (o *GfwRuleGetUsingGETParams) WithContext(ctx context.Context) *GfwRuleGetUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gfw rule get using g e t params
func (o *GfwRuleGetUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gfw rule get using g e t params
func (o *GfwRuleGetUsingGETParams) WithHTTPClient(client *http.Client) *GfwRuleGetUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gfw rule get using g e t params
func (o *GfwRuleGetUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRouterID adds the routerID to the gfw rule get using g e t params
func (o *GfwRuleGetUsingGETParams) WithRouterID(routerID string) *GfwRuleGetUsingGETParams {
	o.SetRouterID(routerID)
	return o
}

// SetRouterID adds the routerId to the gfw rule get using g e t params
func (o *GfwRuleGetUsingGETParams) SetRouterID(routerID string) {
	o.RouterID = routerID
}

// WithRuleID adds the ruleID to the gfw rule get using g e t params
func (o *GfwRuleGetUsingGETParams) WithRuleID(ruleID string) *GfwRuleGetUsingGETParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the gfw rule get using g e t params
func (o *GfwRuleGetUsingGETParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *GfwRuleGetUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param routerId
	if err := r.SetPathParam("routerId", o.RouterID); err != nil {
		return err
	}

	// path param ruleId
	if err := r.SetPathParam("ruleId", o.RuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
