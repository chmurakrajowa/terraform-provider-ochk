// Code generated by go-swagger; DO NOT EDIT.

package firewall_rules_e_w

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

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// NewDfwRuleCreateUsingPUTParams creates a new DfwRuleCreateUsingPUTParams object
// with the default values initialized.
func NewDfwRuleCreateUsingPUTParams() *DfwRuleCreateUsingPUTParams {
	var ()
	return &DfwRuleCreateUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDfwRuleCreateUsingPUTParamsWithTimeout creates a new DfwRuleCreateUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDfwRuleCreateUsingPUTParamsWithTimeout(timeout time.Duration) *DfwRuleCreateUsingPUTParams {
	var ()
	return &DfwRuleCreateUsingPUTParams{

		timeout: timeout,
	}
}

// NewDfwRuleCreateUsingPUTParamsWithContext creates a new DfwRuleCreateUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewDfwRuleCreateUsingPUTParamsWithContext(ctx context.Context) *DfwRuleCreateUsingPUTParams {
	var ()
	return &DfwRuleCreateUsingPUTParams{

		Context: ctx,
	}
}

// NewDfwRuleCreateUsingPUTParamsWithHTTPClient creates a new DfwRuleCreateUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDfwRuleCreateUsingPUTParamsWithHTTPClient(client *http.Client) *DfwRuleCreateUsingPUTParams {
	var ()
	return &DfwRuleCreateUsingPUTParams{
		HTTPClient: client,
	}
}

/*DfwRuleCreateUsingPUTParams contains all the parameters to send to the API endpoint
for the dfw rule create using p u t operation typically these are written to a http.Request
*/
type DfwRuleCreateUsingPUTParams struct {

	/*DfwRule
	  dfwRule

	*/
	DfwRule *models.DFWRule
	/*RouterID
	  routerId

	*/
	RouterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dfw rule create using p u t params
func (o *DfwRuleCreateUsingPUTParams) WithTimeout(timeout time.Duration) *DfwRuleCreateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dfw rule create using p u t params
func (o *DfwRuleCreateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dfw rule create using p u t params
func (o *DfwRuleCreateUsingPUTParams) WithContext(ctx context.Context) *DfwRuleCreateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dfw rule create using p u t params
func (o *DfwRuleCreateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dfw rule create using p u t params
func (o *DfwRuleCreateUsingPUTParams) WithHTTPClient(client *http.Client) *DfwRuleCreateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dfw rule create using p u t params
func (o *DfwRuleCreateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDfwRule adds the dfwRule to the dfw rule create using p u t params
func (o *DfwRuleCreateUsingPUTParams) WithDfwRule(dfwRule *models.DFWRule) *DfwRuleCreateUsingPUTParams {
	o.SetDfwRule(dfwRule)
	return o
}

// SetDfwRule adds the dfwRule to the dfw rule create using p u t params
func (o *DfwRuleCreateUsingPUTParams) SetDfwRule(dfwRule *models.DFWRule) {
	o.DfwRule = dfwRule
}

// WithRouterID adds the routerID to the dfw rule create using p u t params
func (o *DfwRuleCreateUsingPUTParams) WithRouterID(routerID string) *DfwRuleCreateUsingPUTParams {
	o.SetRouterID(routerID)
	return o
}

// SetRouterID adds the routerId to the dfw rule create using p u t params
func (o *DfwRuleCreateUsingPUTParams) SetRouterID(routerID string) {
	o.RouterID = routerID
}

// WriteToRequest writes these params to a swagger request
func (o *DfwRuleCreateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DfwRule != nil {
		if err := r.SetBodyParam(o.DfwRule); err != nil {
			return err
		}
	}

	// path param routerId
	if err := r.SetPathParam("routerId", o.RouterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
