// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

// NewDeleteBillingAccountsAccountIDParams creates a new DeleteBillingAccountsAccountIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBillingAccountsAccountIDParams() *DeleteBillingAccountsAccountIDParams {
	return &DeleteBillingAccountsAccountIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBillingAccountsAccountIDParamsWithTimeout creates a new DeleteBillingAccountsAccountIDParams object
// with the ability to set a timeout on a request.
func NewDeleteBillingAccountsAccountIDParamsWithTimeout(timeout time.Duration) *DeleteBillingAccountsAccountIDParams {
	return &DeleteBillingAccountsAccountIDParams{
		timeout: timeout,
	}
}

// NewDeleteBillingAccountsAccountIDParamsWithContext creates a new DeleteBillingAccountsAccountIDParams object
// with the ability to set a context for a request.
func NewDeleteBillingAccountsAccountIDParamsWithContext(ctx context.Context) *DeleteBillingAccountsAccountIDParams {
	return &DeleteBillingAccountsAccountIDParams{
		Context: ctx,
	}
}

// NewDeleteBillingAccountsAccountIDParamsWithHTTPClient creates a new DeleteBillingAccountsAccountIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBillingAccountsAccountIDParamsWithHTTPClient(client *http.Client) *DeleteBillingAccountsAccountIDParams {
	return &DeleteBillingAccountsAccountIDParams{
		HTTPClient: client,
	}
}

/*
DeleteBillingAccountsAccountIDParams contains all the parameters to send to the API endpoint

	for the delete billing accounts account ID operation.

	Typically these are written to a http.Request.
*/
type DeleteBillingAccountsAccountIDParams struct {

	// AccountID.
	//
	// Format: uuid
	AccountID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete billing accounts account ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBillingAccountsAccountIDParams) WithDefaults() *DeleteBillingAccountsAccountIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete billing accounts account ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBillingAccountsAccountIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete billing accounts account ID params
func (o *DeleteBillingAccountsAccountIDParams) WithTimeout(timeout time.Duration) *DeleteBillingAccountsAccountIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete billing accounts account ID params
func (o *DeleteBillingAccountsAccountIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete billing accounts account ID params
func (o *DeleteBillingAccountsAccountIDParams) WithContext(ctx context.Context) *DeleteBillingAccountsAccountIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete billing accounts account ID params
func (o *DeleteBillingAccountsAccountIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete billing accounts account ID params
func (o *DeleteBillingAccountsAccountIDParams) WithHTTPClient(client *http.Client) *DeleteBillingAccountsAccountIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete billing accounts account ID params
func (o *DeleteBillingAccountsAccountIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the delete billing accounts account ID params
func (o *DeleteBillingAccountsAccountIDParams) WithAccountID(accountID strfmt.UUID) *DeleteBillingAccountsAccountIDParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the delete billing accounts account ID params
func (o *DeleteBillingAccountsAccountIDParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBillingAccountsAccountIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
