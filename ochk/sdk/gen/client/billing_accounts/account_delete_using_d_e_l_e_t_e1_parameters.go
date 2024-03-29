// Code generated by go-swagger; DO NOT EDIT.

package billing_accounts

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

// NewAccountDeleteUsingDELETE1Params creates a new AccountDeleteUsingDELETE1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountDeleteUsingDELETE1Params() *AccountDeleteUsingDELETE1Params {
	return &AccountDeleteUsingDELETE1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountDeleteUsingDELETE1ParamsWithTimeout creates a new AccountDeleteUsingDELETE1Params object
// with the ability to set a timeout on a request.
func NewAccountDeleteUsingDELETE1ParamsWithTimeout(timeout time.Duration) *AccountDeleteUsingDELETE1Params {
	return &AccountDeleteUsingDELETE1Params{
		timeout: timeout,
	}
}

// NewAccountDeleteUsingDELETE1ParamsWithContext creates a new AccountDeleteUsingDELETE1Params object
// with the ability to set a context for a request.
func NewAccountDeleteUsingDELETE1ParamsWithContext(ctx context.Context) *AccountDeleteUsingDELETE1Params {
	return &AccountDeleteUsingDELETE1Params{
		Context: ctx,
	}
}

// NewAccountDeleteUsingDELETE1ParamsWithHTTPClient creates a new AccountDeleteUsingDELETE1Params object
// with the ability to set a custom HTTPClient for a request.
func NewAccountDeleteUsingDELETE1ParamsWithHTTPClient(client *http.Client) *AccountDeleteUsingDELETE1Params {
	return &AccountDeleteUsingDELETE1Params{
		HTTPClient: client,
	}
}

/*
AccountDeleteUsingDELETE1Params contains all the parameters to send to the API endpoint

	for the account delete using d e l e t e 1 operation.

	Typically these are written to a http.Request.
*/
type AccountDeleteUsingDELETE1Params struct {

	/* AlarmDefinitionID.

	   alarmDefinitionId
	*/
	AlarmDefinitionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the account delete using d e l e t e 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountDeleteUsingDELETE1Params) WithDefaults() *AccountDeleteUsingDELETE1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the account delete using d e l e t e 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountDeleteUsingDELETE1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the account delete using d e l e t e 1 params
func (o *AccountDeleteUsingDELETE1Params) WithTimeout(timeout time.Duration) *AccountDeleteUsingDELETE1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account delete using d e l e t e 1 params
func (o *AccountDeleteUsingDELETE1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account delete using d e l e t e 1 params
func (o *AccountDeleteUsingDELETE1Params) WithContext(ctx context.Context) *AccountDeleteUsingDELETE1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account delete using d e l e t e 1 params
func (o *AccountDeleteUsingDELETE1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account delete using d e l e t e 1 params
func (o *AccountDeleteUsingDELETE1Params) WithHTTPClient(client *http.Client) *AccountDeleteUsingDELETE1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account delete using d e l e t e 1 params
func (o *AccountDeleteUsingDELETE1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlarmDefinitionID adds the alarmDefinitionID to the account delete using d e l e t e 1 params
func (o *AccountDeleteUsingDELETE1Params) WithAlarmDefinitionID(alarmDefinitionID string) *AccountDeleteUsingDELETE1Params {
	o.SetAlarmDefinitionID(alarmDefinitionID)
	return o
}

// SetAlarmDefinitionID adds the alarmDefinitionId to the account delete using d e l e t e 1 params
func (o *AccountDeleteUsingDELETE1Params) SetAlarmDefinitionID(alarmDefinitionID string) {
	o.AlarmDefinitionID = alarmDefinitionID
}

// WriteToRequest writes these params to a swagger request
func (o *AccountDeleteUsingDELETE1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param alarmDefinitionId
	if err := r.SetPathParam("alarmDefinitionId", o.AlarmDefinitionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
