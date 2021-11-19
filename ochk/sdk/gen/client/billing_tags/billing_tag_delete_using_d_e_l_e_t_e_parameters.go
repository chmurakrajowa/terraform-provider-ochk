// Code generated by go-swagger; DO NOT EDIT.

package billing_tags

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

// NewBillingTagDeleteUsingDELETEParams creates a new BillingTagDeleteUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBillingTagDeleteUsingDELETEParams() *BillingTagDeleteUsingDELETEParams {
	return &BillingTagDeleteUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBillingTagDeleteUsingDELETEParamsWithTimeout creates a new BillingTagDeleteUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewBillingTagDeleteUsingDELETEParamsWithTimeout(timeout time.Duration) *BillingTagDeleteUsingDELETEParams {
	return &BillingTagDeleteUsingDELETEParams{
		timeout: timeout,
	}
}

// NewBillingTagDeleteUsingDELETEParamsWithContext creates a new BillingTagDeleteUsingDELETEParams object
// with the ability to set a context for a request.
func NewBillingTagDeleteUsingDELETEParamsWithContext(ctx context.Context) *BillingTagDeleteUsingDELETEParams {
	return &BillingTagDeleteUsingDELETEParams{
		Context: ctx,
	}
}

// NewBillingTagDeleteUsingDELETEParamsWithHTTPClient creates a new BillingTagDeleteUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewBillingTagDeleteUsingDELETEParamsWithHTTPClient(client *http.Client) *BillingTagDeleteUsingDELETEParams {
	return &BillingTagDeleteUsingDELETEParams{
		HTTPClient: client,
	}
}

/* BillingTagDeleteUsingDELETEParams contains all the parameters to send to the API endpoint
   for the billing tag delete using d e l e t e operation.

   Typically these are written to a http.Request.
*/
type BillingTagDeleteUsingDELETEParams struct {

	/* BillingTagID.

	   billingTagId

	   Format: int32
	*/
	BillingTagID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the billing tag delete using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BillingTagDeleteUsingDELETEParams) WithDefaults() *BillingTagDeleteUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the billing tag delete using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BillingTagDeleteUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the billing tag delete using d e l e t e params
func (o *BillingTagDeleteUsingDELETEParams) WithTimeout(timeout time.Duration) *BillingTagDeleteUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billing tag delete using d e l e t e params
func (o *BillingTagDeleteUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billing tag delete using d e l e t e params
func (o *BillingTagDeleteUsingDELETEParams) WithContext(ctx context.Context) *BillingTagDeleteUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billing tag delete using d e l e t e params
func (o *BillingTagDeleteUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billing tag delete using d e l e t e params
func (o *BillingTagDeleteUsingDELETEParams) WithHTTPClient(client *http.Client) *BillingTagDeleteUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billing tag delete using d e l e t e params
func (o *BillingTagDeleteUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBillingTagID adds the billingTagID to the billing tag delete using d e l e t e params
func (o *BillingTagDeleteUsingDELETEParams) WithBillingTagID(billingTagID int32) *BillingTagDeleteUsingDELETEParams {
	o.SetBillingTagID(billingTagID)
	return o
}

// SetBillingTagID adds the billingTagId to the billing tag delete using d e l e t e params
func (o *BillingTagDeleteUsingDELETEParams) SetBillingTagID(billingTagID int32) {
	o.BillingTagID = billingTagID
}

// WriteToRequest writes these params to a swagger request
func (o *BillingTagDeleteUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param billingTagId
	if err := r.SetPathParam("billingTagId", swag.FormatInt32(o.BillingTagID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
