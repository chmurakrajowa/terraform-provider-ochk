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

// NewBillingTagGetUsingGETParams creates a new BillingTagGetUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBillingTagGetUsingGETParams() *BillingTagGetUsingGETParams {
	return &BillingTagGetUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBillingTagGetUsingGETParamsWithTimeout creates a new BillingTagGetUsingGETParams object
// with the ability to set a timeout on a request.
func NewBillingTagGetUsingGETParamsWithTimeout(timeout time.Duration) *BillingTagGetUsingGETParams {
	return &BillingTagGetUsingGETParams{
		timeout: timeout,
	}
}

// NewBillingTagGetUsingGETParamsWithContext creates a new BillingTagGetUsingGETParams object
// with the ability to set a context for a request.
func NewBillingTagGetUsingGETParamsWithContext(ctx context.Context) *BillingTagGetUsingGETParams {
	return &BillingTagGetUsingGETParams{
		Context: ctx,
	}
}

// NewBillingTagGetUsingGETParamsWithHTTPClient creates a new BillingTagGetUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewBillingTagGetUsingGETParamsWithHTTPClient(client *http.Client) *BillingTagGetUsingGETParams {
	return &BillingTagGetUsingGETParams{
		HTTPClient: client,
	}
}

/* BillingTagGetUsingGETParams contains all the parameters to send to the API endpoint
   for the billing tag get using g e t operation.

   Typically these are written to a http.Request.
*/
type BillingTagGetUsingGETParams struct {

	/* BillingTagID.

	   billingTagId

	   Format: int32
	*/
	BillingTagID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the billing tag get using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BillingTagGetUsingGETParams) WithDefaults() *BillingTagGetUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the billing tag get using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BillingTagGetUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the billing tag get using g e t params
func (o *BillingTagGetUsingGETParams) WithTimeout(timeout time.Duration) *BillingTagGetUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billing tag get using g e t params
func (o *BillingTagGetUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billing tag get using g e t params
func (o *BillingTagGetUsingGETParams) WithContext(ctx context.Context) *BillingTagGetUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billing tag get using g e t params
func (o *BillingTagGetUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billing tag get using g e t params
func (o *BillingTagGetUsingGETParams) WithHTTPClient(client *http.Client) *BillingTagGetUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billing tag get using g e t params
func (o *BillingTagGetUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBillingTagID adds the billingTagID to the billing tag get using g e t params
func (o *BillingTagGetUsingGETParams) WithBillingTagID(billingTagID int32) *BillingTagGetUsingGETParams {
	o.SetBillingTagID(billingTagID)
	return o
}

// SetBillingTagID adds the billingTagId to the billing tag get using g e t params
func (o *BillingTagGetUsingGETParams) SetBillingTagID(billingTagID int32) {
	o.BillingTagID = billingTagID
}

// WriteToRequest writes these params to a swagger request
func (o *BillingTagGetUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
