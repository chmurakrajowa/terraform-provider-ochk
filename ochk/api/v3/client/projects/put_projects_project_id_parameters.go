// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewPutProjectsProjectIDParams creates a new PutProjectsProjectIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutProjectsProjectIDParams() *PutProjectsProjectIDParams {
	return &PutProjectsProjectIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutProjectsProjectIDParamsWithTimeout creates a new PutProjectsProjectIDParams object
// with the ability to set a timeout on a request.
func NewPutProjectsProjectIDParamsWithTimeout(timeout time.Duration) *PutProjectsProjectIDParams {
	return &PutProjectsProjectIDParams{
		timeout: timeout,
	}
}

// NewPutProjectsProjectIDParamsWithContext creates a new PutProjectsProjectIDParams object
// with the ability to set a context for a request.
func NewPutProjectsProjectIDParamsWithContext(ctx context.Context) *PutProjectsProjectIDParams {
	return &PutProjectsProjectIDParams{
		Context: ctx,
	}
}

// NewPutProjectsProjectIDParamsWithHTTPClient creates a new PutProjectsProjectIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutProjectsProjectIDParamsWithHTTPClient(client *http.Client) *PutProjectsProjectIDParams {
	return &PutProjectsProjectIDParams{
		HTTPClient: client,
	}
}

/*
PutProjectsProjectIDParams contains all the parameters to send to the API endpoint

	for the put projects project ID operation.

	Typically these are written to a http.Request.
*/
type PutProjectsProjectIDParams struct {

	// Body.
	Body *models.ProjectInstance

	// ProjectID.
	//
	// Format: uuid
	ProjectID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put projects project ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutProjectsProjectIDParams) WithDefaults() *PutProjectsProjectIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put projects project ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutProjectsProjectIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put projects project ID params
func (o *PutProjectsProjectIDParams) WithTimeout(timeout time.Duration) *PutProjectsProjectIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put projects project ID params
func (o *PutProjectsProjectIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put projects project ID params
func (o *PutProjectsProjectIDParams) WithContext(ctx context.Context) *PutProjectsProjectIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put projects project ID params
func (o *PutProjectsProjectIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put projects project ID params
func (o *PutProjectsProjectIDParams) WithHTTPClient(client *http.Client) *PutProjectsProjectIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put projects project ID params
func (o *PutProjectsProjectIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put projects project ID params
func (o *PutProjectsProjectIDParams) WithBody(body *models.ProjectInstance) *PutProjectsProjectIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put projects project ID params
func (o *PutProjectsProjectIDParams) SetBody(body *models.ProjectInstance) {
	o.Body = body
}

// WithProjectID adds the projectID to the put projects project ID params
func (o *PutProjectsProjectIDParams) WithProjectID(projectID strfmt.UUID) *PutProjectsProjectIDParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the put projects project ID params
func (o *PutProjectsProjectIDParams) SetProjectID(projectID strfmt.UUID) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *PutProjectsProjectIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param projectId
	if err := r.SetPathParam("projectId", o.ProjectID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
