// Code generated by go-swagger; DO NOT EDIT.

package nat_rule

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

// NewDeleteNetworkNatRulesRuleIDParams creates a new DeleteNetworkNatRulesRuleIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNetworkNatRulesRuleIDParams() *DeleteNetworkNatRulesRuleIDParams {
	return &DeleteNetworkNatRulesRuleIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNetworkNatRulesRuleIDParamsWithTimeout creates a new DeleteNetworkNatRulesRuleIDParams object
// with the ability to set a timeout on a request.
func NewDeleteNetworkNatRulesRuleIDParamsWithTimeout(timeout time.Duration) *DeleteNetworkNatRulesRuleIDParams {
	return &DeleteNetworkNatRulesRuleIDParams{
		timeout: timeout,
	}
}

// NewDeleteNetworkNatRulesRuleIDParamsWithContext creates a new DeleteNetworkNatRulesRuleIDParams object
// with the ability to set a context for a request.
func NewDeleteNetworkNatRulesRuleIDParamsWithContext(ctx context.Context) *DeleteNetworkNatRulesRuleIDParams {
	return &DeleteNetworkNatRulesRuleIDParams{
		Context: ctx,
	}
}

// NewDeleteNetworkNatRulesRuleIDParamsWithHTTPClient creates a new DeleteNetworkNatRulesRuleIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNetworkNatRulesRuleIDParamsWithHTTPClient(client *http.Client) *DeleteNetworkNatRulesRuleIDParams {
	return &DeleteNetworkNatRulesRuleIDParams{
		HTTPClient: client,
	}
}

/*
DeleteNetworkNatRulesRuleIDParams contains all the parameters to send to the API endpoint

	for the delete network nat rules rule ID operation.

	Typically these are written to a http.Request.
*/
type DeleteNetworkNatRulesRuleIDParams struct {

	// RuleID.
	//
	// Format: uuid
	RuleID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete network nat rules rule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkNatRulesRuleIDParams) WithDefaults() *DeleteNetworkNatRulesRuleIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete network nat rules rule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkNatRulesRuleIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete network nat rules rule ID params
func (o *DeleteNetworkNatRulesRuleIDParams) WithTimeout(timeout time.Duration) *DeleteNetworkNatRulesRuleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete network nat rules rule ID params
func (o *DeleteNetworkNatRulesRuleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete network nat rules rule ID params
func (o *DeleteNetworkNatRulesRuleIDParams) WithContext(ctx context.Context) *DeleteNetworkNatRulesRuleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete network nat rules rule ID params
func (o *DeleteNetworkNatRulesRuleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete network nat rules rule ID params
func (o *DeleteNetworkNatRulesRuleIDParams) WithHTTPClient(client *http.Client) *DeleteNetworkNatRulesRuleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete network nat rules rule ID params
func (o *DeleteNetworkNatRulesRuleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRuleID adds the ruleID to the delete network nat rules rule ID params
func (o *DeleteNetworkNatRulesRuleIDParams) WithRuleID(ruleID strfmt.UUID) *DeleteNetworkNatRulesRuleIDParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the delete network nat rules rule ID params
func (o *DeleteNetworkNatRulesRuleIDParams) SetRuleID(ruleID strfmt.UUID) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNetworkNatRulesRuleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ruleId
	if err := r.SetPathParam("ruleId", o.RuleID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
