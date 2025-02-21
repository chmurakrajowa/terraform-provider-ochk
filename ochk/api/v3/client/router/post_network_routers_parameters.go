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

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
)

// NewPostNetworkRoutersParams creates a new PostNetworkRoutersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostNetworkRoutersParams() *PostNetworkRoutersParams {
	return &PostNetworkRoutersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostNetworkRoutersParamsWithTimeout creates a new PostNetworkRoutersParams object
// with the ability to set a timeout on a request.
func NewPostNetworkRoutersParamsWithTimeout(timeout time.Duration) *PostNetworkRoutersParams {
	return &PostNetworkRoutersParams{
		timeout: timeout,
	}
}

// NewPostNetworkRoutersParamsWithContext creates a new PostNetworkRoutersParams object
// with the ability to set a context for a request.
func NewPostNetworkRoutersParamsWithContext(ctx context.Context) *PostNetworkRoutersParams {
	return &PostNetworkRoutersParams{
		Context: ctx,
	}
}

// NewPostNetworkRoutersParamsWithHTTPClient creates a new PostNetworkRoutersParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostNetworkRoutersParamsWithHTTPClient(client *http.Client) *PostNetworkRoutersParams {
	return &PostNetworkRoutersParams{
		HTTPClient: client,
	}
}

/*
PostNetworkRoutersParams contains all the parameters to send to the API endpoint

	for the post network routers operation.

	Typically these are written to a http.Request.
*/
type PostNetworkRoutersParams struct {

	// Body.
	Body *models.RouterInstance

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post network routers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostNetworkRoutersParams) WithDefaults() *PostNetworkRoutersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post network routers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostNetworkRoutersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post network routers params
func (o *PostNetworkRoutersParams) WithTimeout(timeout time.Duration) *PostNetworkRoutersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post network routers params
func (o *PostNetworkRoutersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post network routers params
func (o *PostNetworkRoutersParams) WithContext(ctx context.Context) *PostNetworkRoutersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post network routers params
func (o *PostNetworkRoutersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post network routers params
func (o *PostNetworkRoutersParams) WithHTTPClient(client *http.Client) *PostNetworkRoutersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post network routers params
func (o *PostNetworkRoutersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post network routers params
func (o *PostNetworkRoutersParams) WithBody(body *models.RouterInstance) *PostNetworkRoutersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post network routers params
func (o *PostNetworkRoutersParams) SetBody(body *models.RouterInstance) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostNetworkRoutersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
