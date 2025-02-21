// Code generated by go-swagger; DO NOT EDIT.

package router

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

// NewGetNetworkRoutersParams creates a new GetNetworkRoutersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkRoutersParams() *GetNetworkRoutersParams {
	return &GetNetworkRoutersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkRoutersParamsWithTimeout creates a new GetNetworkRoutersParams object
// with the ability to set a timeout on a request.
func NewGetNetworkRoutersParamsWithTimeout(timeout time.Duration) *GetNetworkRoutersParams {
	return &GetNetworkRoutersParams{
		timeout: timeout,
	}
}

// NewGetNetworkRoutersParamsWithContext creates a new GetNetworkRoutersParams object
// with the ability to set a context for a request.
func NewGetNetworkRoutersParamsWithContext(ctx context.Context) *GetNetworkRoutersParams {
	return &GetNetworkRoutersParams{
		Context: ctx,
	}
}

// NewGetNetworkRoutersParamsWithHTTPClient creates a new GetNetworkRoutersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkRoutersParamsWithHTTPClient(client *http.Client) *GetNetworkRoutersParams {
	return &GetNetworkRoutersParams{
		HTTPClient: client,
	}
}

/*
GetNetworkRoutersParams contains all the parameters to send to the API endpoint

	for the get network routers operation.

	Typically these are written to a http.Request.
*/
type GetNetworkRoutersParams struct {

	// DisplayName.
	DisplayName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network routers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkRoutersParams) WithDefaults() *GetNetworkRoutersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network routers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkRoutersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network routers params
func (o *GetNetworkRoutersParams) WithTimeout(timeout time.Duration) *GetNetworkRoutersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network routers params
func (o *GetNetworkRoutersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network routers params
func (o *GetNetworkRoutersParams) WithContext(ctx context.Context) *GetNetworkRoutersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network routers params
func (o *GetNetworkRoutersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network routers params
func (o *GetNetworkRoutersParams) WithHTTPClient(client *http.Client) *GetNetworkRoutersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network routers params
func (o *GetNetworkRoutersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisplayName adds the displayName to the get network routers params
func (o *GetNetworkRoutersParams) WithDisplayName(displayName *string) *GetNetworkRoutersParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the get network routers params
func (o *GetNetworkRoutersParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkRoutersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DisplayName != nil {

		// query param displayName
		var qrDisplayName string

		if o.DisplayName != nil {
			qrDisplayName = *o.DisplayName
		}
		qDisplayName := qrDisplayName
		if qDisplayName != "" {

			if err := r.SetQueryParam("displayName", qDisplayName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
