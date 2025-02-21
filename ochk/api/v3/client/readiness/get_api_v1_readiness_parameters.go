// Code generated by go-swagger; DO NOT EDIT.

package readiness

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

// NewGetAPIV1ReadinessParams creates a new GetAPIV1ReadinessParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV1ReadinessParams() *GetAPIV1ReadinessParams {
	return &GetAPIV1ReadinessParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV1ReadinessParamsWithTimeout creates a new GetAPIV1ReadinessParams object
// with the ability to set a timeout on a request.
func NewGetAPIV1ReadinessParamsWithTimeout(timeout time.Duration) *GetAPIV1ReadinessParams {
	return &GetAPIV1ReadinessParams{
		timeout: timeout,
	}
}

// NewGetAPIV1ReadinessParamsWithContext creates a new GetAPIV1ReadinessParams object
// with the ability to set a context for a request.
func NewGetAPIV1ReadinessParamsWithContext(ctx context.Context) *GetAPIV1ReadinessParams {
	return &GetAPIV1ReadinessParams{
		Context: ctx,
	}
}

// NewGetAPIV1ReadinessParamsWithHTTPClient creates a new GetAPIV1ReadinessParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV1ReadinessParamsWithHTTPClient(client *http.Client) *GetAPIV1ReadinessParams {
	return &GetAPIV1ReadinessParams{
		HTTPClient: client,
	}
}

/*
GetAPIV1ReadinessParams contains all the parameters to send to the API endpoint

	for the get API v1 readiness operation.

	Typically these are written to a http.Request.
*/
type GetAPIV1ReadinessParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API v1 readiness params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV1ReadinessParams) WithDefaults() *GetAPIV1ReadinessParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v1 readiness params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV1ReadinessParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API v1 readiness params
func (o *GetAPIV1ReadinessParams) WithTimeout(timeout time.Duration) *GetAPIV1ReadinessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v1 readiness params
func (o *GetAPIV1ReadinessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v1 readiness params
func (o *GetAPIV1ReadinessParams) WithContext(ctx context.Context) *GetAPIV1ReadinessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v1 readiness params
func (o *GetAPIV1ReadinessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v1 readiness params
func (o *GetAPIV1ReadinessParams) WithHTTPClient(client *http.Client) *GetAPIV1ReadinessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v1 readiness params
func (o *GetAPIV1ReadinessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV1ReadinessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
