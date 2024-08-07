// Code generated by go-swagger; DO NOT EDIT.

package billing_public_ip_report

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
	"github.com/go-openapi/swag"
)

// NewGeneratePIPBillingUsingGETParams creates a new GeneratePIPBillingUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGeneratePIPBillingUsingGETParams() *GeneratePIPBillingUsingGETParams {
	return &GeneratePIPBillingUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGeneratePIPBillingUsingGETParamsWithTimeout creates a new GeneratePIPBillingUsingGETParams object
// with the ability to set a timeout on a request.
func NewGeneratePIPBillingUsingGETParamsWithTimeout(timeout time.Duration) *GeneratePIPBillingUsingGETParams {
	return &GeneratePIPBillingUsingGETParams{
		timeout: timeout,
	}
}

// NewGeneratePIPBillingUsingGETParamsWithContext creates a new GeneratePIPBillingUsingGETParams object
// with the ability to set a context for a request.
func NewGeneratePIPBillingUsingGETParamsWithContext(ctx context.Context) *GeneratePIPBillingUsingGETParams {
	return &GeneratePIPBillingUsingGETParams{
		Context: ctx,
	}
}

// NewGeneratePIPBillingUsingGETParamsWithHTTPClient creates a new GeneratePIPBillingUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGeneratePIPBillingUsingGETParamsWithHTTPClient(client *http.Client) *GeneratePIPBillingUsingGETParams {
	return &GeneratePIPBillingUsingGETParams{
		HTTPClient: client,
	}
}

/*
GeneratePIPBillingUsingGETParams contains all the parameters to send to the API endpoint

	for the generate p IP billing using g e t operation.

	Typically these are written to a http.Request.
*/
type GeneratePIPBillingUsingGETParams struct {

	/* AccountID.

	   accountId
	*/
	AccountID string

	/* BillingPeriodID.

	   billingPeriodId

	   Format: int32
	*/
	BillingPeriodID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the generate p IP billing using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GeneratePIPBillingUsingGETParams) WithDefaults() *GeneratePIPBillingUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the generate p IP billing using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GeneratePIPBillingUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the generate p IP billing using g e t params
func (o *GeneratePIPBillingUsingGETParams) WithTimeout(timeout time.Duration) *GeneratePIPBillingUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generate p IP billing using g e t params
func (o *GeneratePIPBillingUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generate p IP billing using g e t params
func (o *GeneratePIPBillingUsingGETParams) WithContext(ctx context.Context) *GeneratePIPBillingUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generate p IP billing using g e t params
func (o *GeneratePIPBillingUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generate p IP billing using g e t params
func (o *GeneratePIPBillingUsingGETParams) WithHTTPClient(client *http.Client) *GeneratePIPBillingUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generate p IP billing using g e t params
func (o *GeneratePIPBillingUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the generate p IP billing using g e t params
func (o *GeneratePIPBillingUsingGETParams) WithAccountID(accountID string) *GeneratePIPBillingUsingGETParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the generate p IP billing using g e t params
func (o *GeneratePIPBillingUsingGETParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithBillingPeriodID adds the billingPeriodID to the generate p IP billing using g e t params
func (o *GeneratePIPBillingUsingGETParams) WithBillingPeriodID(billingPeriodID int32) *GeneratePIPBillingUsingGETParams {
	o.SetBillingPeriodID(billingPeriodID)
	return o
}

// SetBillingPeriodID adds the billingPeriodId to the generate p IP billing using g e t params
func (o *GeneratePIPBillingUsingGETParams) SetBillingPeriodID(billingPeriodID int32) {
	o.BillingPeriodID = billingPeriodID
}

// WriteToRequest writes these params to a swagger request
func (o *GeneratePIPBillingUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID); err != nil {
		return err
	}

	// path param billingPeriodId
	if err := r.SetPathParam("billingPeriodId", swag.FormatInt32(o.BillingPeriodID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
