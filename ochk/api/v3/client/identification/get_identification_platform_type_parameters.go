// Code generated by go-swagger; DO NOT EDIT.

package identification

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

// NewGetIdentificationPlatformTypeParams creates a new GetIdentificationPlatformTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIdentificationPlatformTypeParams() *GetIdentificationPlatformTypeParams {
	return &GetIdentificationPlatformTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIdentificationPlatformTypeParamsWithTimeout creates a new GetIdentificationPlatformTypeParams object
// with the ability to set a timeout on a request.
func NewGetIdentificationPlatformTypeParamsWithTimeout(timeout time.Duration) *GetIdentificationPlatformTypeParams {
	return &GetIdentificationPlatformTypeParams{
		timeout: timeout,
	}
}

// NewGetIdentificationPlatformTypeParamsWithContext creates a new GetIdentificationPlatformTypeParams object
// with the ability to set a context for a request.
func NewGetIdentificationPlatformTypeParamsWithContext(ctx context.Context) *GetIdentificationPlatformTypeParams {
	return &GetIdentificationPlatformTypeParams{
		Context: ctx,
	}
}

// NewGetIdentificationPlatformTypeParamsWithHTTPClient creates a new GetIdentificationPlatformTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIdentificationPlatformTypeParamsWithHTTPClient(client *http.Client) *GetIdentificationPlatformTypeParams {
	return &GetIdentificationPlatformTypeParams{
		HTTPClient: client,
	}
}

/*
GetIdentificationPlatformTypeParams contains all the parameters to send to the API endpoint

	for the get identification platform type operation.

	Typically these are written to a http.Request.
*/
type GetIdentificationPlatformTypeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get identification platform type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIdentificationPlatformTypeParams) WithDefaults() *GetIdentificationPlatformTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get identification platform type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIdentificationPlatformTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get identification platform type params
func (o *GetIdentificationPlatformTypeParams) WithTimeout(timeout time.Duration) *GetIdentificationPlatformTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get identification platform type params
func (o *GetIdentificationPlatformTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get identification platform type params
func (o *GetIdentificationPlatformTypeParams) WithContext(ctx context.Context) *GetIdentificationPlatformTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get identification platform type params
func (o *GetIdentificationPlatformTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get identification platform type params
func (o *GetIdentificationPlatformTypeParams) WithHTTPClient(client *http.Client) *GetIdentificationPlatformTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get identification platform type params
func (o *GetIdentificationPlatformTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetIdentificationPlatformTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
