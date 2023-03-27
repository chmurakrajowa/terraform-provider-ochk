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
)

// NewKeyDeleteUsingDELETEParams creates a new KeyDeleteUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKeyDeleteUsingDELETEParams() *KeyDeleteUsingDELETEParams {
	return &KeyDeleteUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKeyDeleteUsingDELETEParamsWithTimeout creates a new KeyDeleteUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewKeyDeleteUsingDELETEParamsWithTimeout(timeout time.Duration) *KeyDeleteUsingDELETEParams {
	return &KeyDeleteUsingDELETEParams{
		timeout: timeout,
	}
}

// NewKeyDeleteUsingDELETEParamsWithContext creates a new KeyDeleteUsingDELETEParams object
// with the ability to set a context for a request.
func NewKeyDeleteUsingDELETEParamsWithContext(ctx context.Context) *KeyDeleteUsingDELETEParams {
	return &KeyDeleteUsingDELETEParams{
		Context: ctx,
	}
}

// NewKeyDeleteUsingDELETEParamsWithHTTPClient creates a new KeyDeleteUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewKeyDeleteUsingDELETEParamsWithHTTPClient(client *http.Client) *KeyDeleteUsingDELETEParams {
	return &KeyDeleteUsingDELETEParams{
		HTTPClient: client,
	}
}

/*
KeyDeleteUsingDELETEParams contains all the parameters to send to the API endpoint

	for the key delete using d e l e t e operation.

	Typically these are written to a http.Request.
*/
type KeyDeleteUsingDELETEParams struct {

	/* ID.

	   id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the key delete using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KeyDeleteUsingDELETEParams) WithDefaults() *KeyDeleteUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the key delete using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KeyDeleteUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the key delete using d e l e t e params
func (o *KeyDeleteUsingDELETEParams) WithTimeout(timeout time.Duration) *KeyDeleteUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the key delete using d e l e t e params
func (o *KeyDeleteUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the key delete using d e l e t e params
func (o *KeyDeleteUsingDELETEParams) WithContext(ctx context.Context) *KeyDeleteUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the key delete using d e l e t e params
func (o *KeyDeleteUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the key delete using d e l e t e params
func (o *KeyDeleteUsingDELETEParams) WithHTTPClient(client *http.Client) *KeyDeleteUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the key delete using d e l e t e params
func (o *KeyDeleteUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the key delete using d e l e t e params
func (o *KeyDeleteUsingDELETEParams) WithID(id string) *KeyDeleteUsingDELETEParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the key delete using d e l e t e params
func (o *KeyDeleteUsingDELETEParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *KeyDeleteUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
