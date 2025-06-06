// Code generated by go-swagger; DO NOT EDIT.

package deployments

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

// NewGetDeploymentsParams creates a new GetDeploymentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeploymentsParams() *GetDeploymentsParams {
	return &GetDeploymentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentsParamsWithTimeout creates a new GetDeploymentsParams object
// with the ability to set a timeout on a request.
func NewGetDeploymentsParamsWithTimeout(timeout time.Duration) *GetDeploymentsParams {
	return &GetDeploymentsParams{
		timeout: timeout,
	}
}

// NewGetDeploymentsParamsWithContext creates a new GetDeploymentsParams object
// with the ability to set a context for a request.
func NewGetDeploymentsParamsWithContext(ctx context.Context) *GetDeploymentsParams {
	return &GetDeploymentsParams{
		Context: ctx,
	}
}

// NewGetDeploymentsParamsWithHTTPClient creates a new GetDeploymentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeploymentsParamsWithHTTPClient(client *http.Client) *GetDeploymentsParams {
	return &GetDeploymentsParams{
		HTTPClient: client,
	}
}

/*
GetDeploymentsParams contains all the parameters to send to the API endpoint

	for the get deployments operation.

	Typically these are written to a http.Request.
*/
type GetDeploymentsParams struct {

	// DisplayName.
	DisplayName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get deployments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentsParams) WithDefaults() *GetDeploymentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get deployments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get deployments params
func (o *GetDeploymentsParams) WithTimeout(timeout time.Duration) *GetDeploymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployments params
func (o *GetDeploymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployments params
func (o *GetDeploymentsParams) WithContext(ctx context.Context) *GetDeploymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployments params
func (o *GetDeploymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployments params
func (o *GetDeploymentsParams) WithHTTPClient(client *http.Client) *GetDeploymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployments params
func (o *GetDeploymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisplayName adds the displayName to the get deployments params
func (o *GetDeploymentsParams) WithDisplayName(displayName *string) *GetDeploymentsParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the get deployments params
func (o *GetDeploymentsParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
