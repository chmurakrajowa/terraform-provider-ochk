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

// NewDeploymentListUsingGETParams creates a new DeploymentListUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeploymentListUsingGETParams() *DeploymentListUsingGETParams {
	return &DeploymentListUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeploymentListUsingGETParamsWithTimeout creates a new DeploymentListUsingGETParams object
// with the ability to set a timeout on a request.
func NewDeploymentListUsingGETParamsWithTimeout(timeout time.Duration) *DeploymentListUsingGETParams {
	return &DeploymentListUsingGETParams{
		timeout: timeout,
	}
}

// NewDeploymentListUsingGETParamsWithContext creates a new DeploymentListUsingGETParams object
// with the ability to set a context for a request.
func NewDeploymentListUsingGETParamsWithContext(ctx context.Context) *DeploymentListUsingGETParams {
	return &DeploymentListUsingGETParams{
		Context: ctx,
	}
}

// NewDeploymentListUsingGETParamsWithHTTPClient creates a new DeploymentListUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeploymentListUsingGETParamsWithHTTPClient(client *http.Client) *DeploymentListUsingGETParams {
	return &DeploymentListUsingGETParams{
		HTTPClient: client,
	}
}

/*
DeploymentListUsingGETParams contains all the parameters to send to the API endpoint

	for the deployment list using g e t operation.

	Typically these are written to a http.Request.
*/
type DeploymentListUsingGETParams struct {

	/* DisplayName.

	   displayName
	*/
	DisplayName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the deployment list using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeploymentListUsingGETParams) WithDefaults() *DeploymentListUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the deployment list using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeploymentListUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the deployment list using g e t params
func (o *DeploymentListUsingGETParams) WithTimeout(timeout time.Duration) *DeploymentListUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deployment list using g e t params
func (o *DeploymentListUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deployment list using g e t params
func (o *DeploymentListUsingGETParams) WithContext(ctx context.Context) *DeploymentListUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deployment list using g e t params
func (o *DeploymentListUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deployment list using g e t params
func (o *DeploymentListUsingGETParams) WithHTTPClient(client *http.Client) *DeploymentListUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deployment list using g e t params
func (o *DeploymentListUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisplayName adds the displayName to the deployment list using g e t params
func (o *DeploymentListUsingGETParams) WithDisplayName(displayName *string) *DeploymentListUsingGETParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the deployment list using g e t params
func (o *DeploymentListUsingGETParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WriteToRequest writes these params to a swagger request
func (o *DeploymentListUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
