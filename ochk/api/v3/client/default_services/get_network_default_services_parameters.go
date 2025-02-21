// Code generated by go-swagger; DO NOT EDIT.

package default_services

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

// NewGetNetworkDefaultServicesParams creates a new GetNetworkDefaultServicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkDefaultServicesParams() *GetNetworkDefaultServicesParams {
	return &GetNetworkDefaultServicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkDefaultServicesParamsWithTimeout creates a new GetNetworkDefaultServicesParams object
// with the ability to set a timeout on a request.
func NewGetNetworkDefaultServicesParamsWithTimeout(timeout time.Duration) *GetNetworkDefaultServicesParams {
	return &GetNetworkDefaultServicesParams{
		timeout: timeout,
	}
}

// NewGetNetworkDefaultServicesParamsWithContext creates a new GetNetworkDefaultServicesParams object
// with the ability to set a context for a request.
func NewGetNetworkDefaultServicesParamsWithContext(ctx context.Context) *GetNetworkDefaultServicesParams {
	return &GetNetworkDefaultServicesParams{
		Context: ctx,
	}
}

// NewGetNetworkDefaultServicesParamsWithHTTPClient creates a new GetNetworkDefaultServicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkDefaultServicesParamsWithHTTPClient(client *http.Client) *GetNetworkDefaultServicesParams {
	return &GetNetworkDefaultServicesParams{
		HTTPClient: client,
	}
}

/*
GetNetworkDefaultServicesParams contains all the parameters to send to the API endpoint

	for the get network default services operation.

	Typically these are written to a http.Request.
*/
type GetNetworkDefaultServicesParams struct {

	// DisplayName.
	DisplayName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network default services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkDefaultServicesParams) WithDefaults() *GetNetworkDefaultServicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network default services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkDefaultServicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network default services params
func (o *GetNetworkDefaultServicesParams) WithTimeout(timeout time.Duration) *GetNetworkDefaultServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network default services params
func (o *GetNetworkDefaultServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network default services params
func (o *GetNetworkDefaultServicesParams) WithContext(ctx context.Context) *GetNetworkDefaultServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network default services params
func (o *GetNetworkDefaultServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network default services params
func (o *GetNetworkDefaultServicesParams) WithHTTPClient(client *http.Client) *GetNetworkDefaultServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network default services params
func (o *GetNetworkDefaultServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisplayName adds the displayName to the get network default services params
func (o *GetNetworkDefaultServicesParams) WithDisplayName(displayName *string) *GetNetworkDefaultServicesParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the get network default services params
func (o *GetNetworkDefaultServicesParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkDefaultServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
