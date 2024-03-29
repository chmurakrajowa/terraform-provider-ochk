// Code generated by go-swagger; DO NOT EDIT.

package n_a_t_rules

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

// NewNatRuleDeleteUsingDELETEParams creates a new NatRuleDeleteUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNatRuleDeleteUsingDELETEParams() *NatRuleDeleteUsingDELETEParams {
	return &NatRuleDeleteUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNatRuleDeleteUsingDELETEParamsWithTimeout creates a new NatRuleDeleteUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewNatRuleDeleteUsingDELETEParamsWithTimeout(timeout time.Duration) *NatRuleDeleteUsingDELETEParams {
	return &NatRuleDeleteUsingDELETEParams{
		timeout: timeout,
	}
}

// NewNatRuleDeleteUsingDELETEParamsWithContext creates a new NatRuleDeleteUsingDELETEParams object
// with the ability to set a context for a request.
func NewNatRuleDeleteUsingDELETEParamsWithContext(ctx context.Context) *NatRuleDeleteUsingDELETEParams {
	return &NatRuleDeleteUsingDELETEParams{
		Context: ctx,
	}
}

// NewNatRuleDeleteUsingDELETEParamsWithHTTPClient creates a new NatRuleDeleteUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewNatRuleDeleteUsingDELETEParamsWithHTTPClient(client *http.Client) *NatRuleDeleteUsingDELETEParams {
	return &NatRuleDeleteUsingDELETEParams{
		HTTPClient: client,
	}
}

/*
NatRuleDeleteUsingDELETEParams contains all the parameters to send to the API endpoint

	for the nat rule delete using d e l e t e operation.

	Typically these are written to a http.Request.
*/
type NatRuleDeleteUsingDELETEParams struct {

	/* RuleID.

	   ruleId
	*/
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nat rule delete using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NatRuleDeleteUsingDELETEParams) WithDefaults() *NatRuleDeleteUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nat rule delete using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NatRuleDeleteUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the nat rule delete using d e l e t e params
func (o *NatRuleDeleteUsingDELETEParams) WithTimeout(timeout time.Duration) *NatRuleDeleteUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nat rule delete using d e l e t e params
func (o *NatRuleDeleteUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nat rule delete using d e l e t e params
func (o *NatRuleDeleteUsingDELETEParams) WithContext(ctx context.Context) *NatRuleDeleteUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nat rule delete using d e l e t e params
func (o *NatRuleDeleteUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nat rule delete using d e l e t e params
func (o *NatRuleDeleteUsingDELETEParams) WithHTTPClient(client *http.Client) *NatRuleDeleteUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nat rule delete using d e l e t e params
func (o *NatRuleDeleteUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRuleID adds the ruleID to the nat rule delete using d e l e t e params
func (o *NatRuleDeleteUsingDELETEParams) WithRuleID(ruleID string) *NatRuleDeleteUsingDELETEParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the nat rule delete using d e l e t e params
func (o *NatRuleDeleteUsingDELETEParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *NatRuleDeleteUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ruleId
	if err := r.SetPathParam("ruleId", o.RuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
