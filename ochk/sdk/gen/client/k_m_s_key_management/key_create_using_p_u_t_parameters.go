// Code generated by go-swagger; DO NOT EDIT.

package k_m_s_key_management

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

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// NewKeyCreateUsingPUTParams creates a new KeyCreateUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKeyCreateUsingPUTParams() *KeyCreateUsingPUTParams {
	return &KeyCreateUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKeyCreateUsingPUTParamsWithTimeout creates a new KeyCreateUsingPUTParams object
// with the ability to set a timeout on a request.
func NewKeyCreateUsingPUTParamsWithTimeout(timeout time.Duration) *KeyCreateUsingPUTParams {
	return &KeyCreateUsingPUTParams{
		timeout: timeout,
	}
}

// NewKeyCreateUsingPUTParamsWithContext creates a new KeyCreateUsingPUTParams object
// with the ability to set a context for a request.
func NewKeyCreateUsingPUTParamsWithContext(ctx context.Context) *KeyCreateUsingPUTParams {
	return &KeyCreateUsingPUTParams{
		Context: ctx,
	}
}

// NewKeyCreateUsingPUTParamsWithHTTPClient creates a new KeyCreateUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewKeyCreateUsingPUTParamsWithHTTPClient(client *http.Client) *KeyCreateUsingPUTParams {
	return &KeyCreateUsingPUTParams{
		HTTPClient: client,
	}
}

/*
KeyCreateUsingPUTParams contains all the parameters to send to the API endpoint

	for the key create using p u t operation.

	Typically these are written to a http.Request.
*/
type KeyCreateUsingPUTParams struct {

	/* KeyInstance.

	   keyInstance
	*/
	KeyInstance *models.KeyInstance

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the key create using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KeyCreateUsingPUTParams) WithDefaults() *KeyCreateUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the key create using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KeyCreateUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the key create using p u t params
func (o *KeyCreateUsingPUTParams) WithTimeout(timeout time.Duration) *KeyCreateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the key create using p u t params
func (o *KeyCreateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the key create using p u t params
func (o *KeyCreateUsingPUTParams) WithContext(ctx context.Context) *KeyCreateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the key create using p u t params
func (o *KeyCreateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the key create using p u t params
func (o *KeyCreateUsingPUTParams) WithHTTPClient(client *http.Client) *KeyCreateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the key create using p u t params
func (o *KeyCreateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeyInstance adds the keyInstance to the key create using p u t params
func (o *KeyCreateUsingPUTParams) WithKeyInstance(keyInstance *models.KeyInstance) *KeyCreateUsingPUTParams {
	o.SetKeyInstance(keyInstance)
	return o
}

// SetKeyInstance adds the keyInstance to the key create using p u t params
func (o *KeyCreateUsingPUTParams) SetKeyInstance(keyInstance *models.KeyInstance) {
	o.KeyInstance = keyInstance
}

// WriteToRequest writes these params to a swagger request
func (o *KeyCreateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.KeyInstance != nil {
		if err := r.SetBodyParam(o.KeyInstance); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
