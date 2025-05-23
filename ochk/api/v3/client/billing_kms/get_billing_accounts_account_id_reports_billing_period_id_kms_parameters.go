// Code generated by go-swagger; DO NOT EDIT.

package billing_kms

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

// NewGetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams creates a new GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams() *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams {
	return &GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBillingAccountsAccountIDReportsBillingPeriodIDKmsParamsWithTimeout creates a new GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams object
// with the ability to set a timeout on a request.
func NewGetBillingAccountsAccountIDReportsBillingPeriodIDKmsParamsWithTimeout(timeout time.Duration) *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams {
	return &GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams{
		timeout: timeout,
	}
}

// NewGetBillingAccountsAccountIDReportsBillingPeriodIDKmsParamsWithContext creates a new GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams object
// with the ability to set a context for a request.
func NewGetBillingAccountsAccountIDReportsBillingPeriodIDKmsParamsWithContext(ctx context.Context) *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams {
	return &GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams{
		Context: ctx,
	}
}

// NewGetBillingAccountsAccountIDReportsBillingPeriodIDKmsParamsWithHTTPClient creates a new GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBillingAccountsAccountIDReportsBillingPeriodIDKmsParamsWithHTTPClient(client *http.Client) *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams {
	return &GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams{
		HTTPClient: client,
	}
}

/*
GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams contains all the parameters to send to the API endpoint

	for the get billing accounts account ID reports billing period ID kms operation.

	Typically these are written to a http.Request.
*/
type GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams struct {

	// AccountID.
	//
	// Format: uuid
	AccountID strfmt.UUID

	// BillingPeriodID.
	//
	// Format: int32
	BillingPeriodID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get billing accounts account ID reports billing period ID kms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams) WithDefaults() *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get billing accounts account ID reports billing period ID kms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get billing accounts account ID reports billing period ID kms params
func (o *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams) WithTimeout(timeout time.Duration) *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get billing accounts account ID reports billing period ID kms params
func (o *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get billing accounts account ID reports billing period ID kms params
func (o *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams) WithContext(ctx context.Context) *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get billing accounts account ID reports billing period ID kms params
func (o *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get billing accounts account ID reports billing period ID kms params
func (o *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams) WithHTTPClient(client *http.Client) *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get billing accounts account ID reports billing period ID kms params
func (o *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get billing accounts account ID reports billing period ID kms params
func (o *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams) WithAccountID(accountID strfmt.UUID) *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get billing accounts account ID reports billing period ID kms params
func (o *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithBillingPeriodID adds the billingPeriodID to the get billing accounts account ID reports billing period ID kms params
func (o *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams) WithBillingPeriodID(billingPeriodID int32) *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams {
	o.SetBillingPeriodID(billingPeriodID)
	return o
}

// SetBillingPeriodID adds the billingPeriodId to the get billing accounts account ID reports billing period ID kms params
func (o *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams) SetBillingPeriodID(billingPeriodID int32) {
	o.BillingPeriodID = billingPeriodID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBillingAccountsAccountIDReportsBillingPeriodIDKmsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
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
