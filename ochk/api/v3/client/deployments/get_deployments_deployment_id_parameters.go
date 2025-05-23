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

// NewGetDeploymentsDeploymentIDParams creates a new GetDeploymentsDeploymentIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeploymentsDeploymentIDParams() *GetDeploymentsDeploymentIDParams {
	return &GetDeploymentsDeploymentIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentsDeploymentIDParamsWithTimeout creates a new GetDeploymentsDeploymentIDParams object
// with the ability to set a timeout on a request.
func NewGetDeploymentsDeploymentIDParamsWithTimeout(timeout time.Duration) *GetDeploymentsDeploymentIDParams {
	return &GetDeploymentsDeploymentIDParams{
		timeout: timeout,
	}
}

// NewGetDeploymentsDeploymentIDParamsWithContext creates a new GetDeploymentsDeploymentIDParams object
// with the ability to set a context for a request.
func NewGetDeploymentsDeploymentIDParamsWithContext(ctx context.Context) *GetDeploymentsDeploymentIDParams {
	return &GetDeploymentsDeploymentIDParams{
		Context: ctx,
	}
}

// NewGetDeploymentsDeploymentIDParamsWithHTTPClient creates a new GetDeploymentsDeploymentIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeploymentsDeploymentIDParamsWithHTTPClient(client *http.Client) *GetDeploymentsDeploymentIDParams {
	return &GetDeploymentsDeploymentIDParams{
		HTTPClient: client,
	}
}

/*
GetDeploymentsDeploymentIDParams contains all the parameters to send to the API endpoint

	for the get deployments deployment ID operation.

	Typically these are written to a http.Request.
*/
type GetDeploymentsDeploymentIDParams struct {

	// DeploymentID.
	//
	// Format: uuid
	DeploymentID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get deployments deployment ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentsDeploymentIDParams) WithDefaults() *GetDeploymentsDeploymentIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get deployments deployment ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentsDeploymentIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get deployments deployment ID params
func (o *GetDeploymentsDeploymentIDParams) WithTimeout(timeout time.Duration) *GetDeploymentsDeploymentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployments deployment ID params
func (o *GetDeploymentsDeploymentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployments deployment ID params
func (o *GetDeploymentsDeploymentIDParams) WithContext(ctx context.Context) *GetDeploymentsDeploymentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployments deployment ID params
func (o *GetDeploymentsDeploymentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployments deployment ID params
func (o *GetDeploymentsDeploymentIDParams) WithHTTPClient(client *http.Client) *GetDeploymentsDeploymentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployments deployment ID params
func (o *GetDeploymentsDeploymentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeploymentID adds the deploymentID to the get deployments deployment ID params
func (o *GetDeploymentsDeploymentIDParams) WithDeploymentID(deploymentID strfmt.UUID) *GetDeploymentsDeploymentIDParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the get deployments deployment ID params
func (o *GetDeploymentsDeploymentIDParams) SetDeploymentID(deploymentID strfmt.UUID) {
	o.DeploymentID = deploymentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentsDeploymentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deploymentId
	if err := r.SetPathParam("deploymentId", o.DeploymentID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
