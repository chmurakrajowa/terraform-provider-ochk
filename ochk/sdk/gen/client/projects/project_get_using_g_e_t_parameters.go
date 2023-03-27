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
)

// NewProjectGetUsingGETParams creates a new ProjectGetUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectGetUsingGETParams() *ProjectGetUsingGETParams {
	return &ProjectGetUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectGetUsingGETParamsWithTimeout creates a new ProjectGetUsingGETParams object
// with the ability to set a timeout on a request.
func NewProjectGetUsingGETParamsWithTimeout(timeout time.Duration) *ProjectGetUsingGETParams {
	return &ProjectGetUsingGETParams{
		timeout: timeout,
	}
}

// NewProjectGetUsingGETParamsWithContext creates a new ProjectGetUsingGETParams object
// with the ability to set a context for a request.
func NewProjectGetUsingGETParamsWithContext(ctx context.Context) *ProjectGetUsingGETParams {
	return &ProjectGetUsingGETParams{
		Context: ctx,
	}
}

// NewProjectGetUsingGETParamsWithHTTPClient creates a new ProjectGetUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectGetUsingGETParamsWithHTTPClient(client *http.Client) *ProjectGetUsingGETParams {
	return &ProjectGetUsingGETParams{
		HTTPClient: client,
	}
}

/*
ProjectGetUsingGETParams contains all the parameters to send to the API endpoint

	for the project get using g e t operation.

	Typically these are written to a http.Request.
*/
type ProjectGetUsingGETParams struct {

	/* ProjectID.

	   projectId
	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the project get using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectGetUsingGETParams) WithDefaults() *ProjectGetUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the project get using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectGetUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the project get using g e t params
func (o *ProjectGetUsingGETParams) WithTimeout(timeout time.Duration) *ProjectGetUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project get using g e t params
func (o *ProjectGetUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project get using g e t params
func (o *ProjectGetUsingGETParams) WithContext(ctx context.Context) *ProjectGetUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project get using g e t params
func (o *ProjectGetUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project get using g e t params
func (o *ProjectGetUsingGETParams) WithHTTPClient(client *http.Client) *ProjectGetUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project get using g e t params
func (o *ProjectGetUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the project get using g e t params
func (o *ProjectGetUsingGETParams) WithProjectID(projectID string) *ProjectGetUsingGETParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the project get using g e t params
func (o *ProjectGetUsingGETParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectGetUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param projectId
	if err := r.SetPathParam("projectId", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
