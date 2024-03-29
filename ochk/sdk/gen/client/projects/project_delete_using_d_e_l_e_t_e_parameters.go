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

// NewProjectDeleteUsingDELETEParams creates a new ProjectDeleteUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectDeleteUsingDELETEParams() *ProjectDeleteUsingDELETEParams {
	return &ProjectDeleteUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectDeleteUsingDELETEParamsWithTimeout creates a new ProjectDeleteUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewProjectDeleteUsingDELETEParamsWithTimeout(timeout time.Duration) *ProjectDeleteUsingDELETEParams {
	return &ProjectDeleteUsingDELETEParams{
		timeout: timeout,
	}
}

// NewProjectDeleteUsingDELETEParamsWithContext creates a new ProjectDeleteUsingDELETEParams object
// with the ability to set a context for a request.
func NewProjectDeleteUsingDELETEParamsWithContext(ctx context.Context) *ProjectDeleteUsingDELETEParams {
	return &ProjectDeleteUsingDELETEParams{
		Context: ctx,
	}
}

// NewProjectDeleteUsingDELETEParamsWithHTTPClient creates a new ProjectDeleteUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectDeleteUsingDELETEParamsWithHTTPClient(client *http.Client) *ProjectDeleteUsingDELETEParams {
	return &ProjectDeleteUsingDELETEParams{
		HTTPClient: client,
	}
}

/*
ProjectDeleteUsingDELETEParams contains all the parameters to send to the API endpoint

	for the project delete using d e l e t e operation.

	Typically these are written to a http.Request.
*/
type ProjectDeleteUsingDELETEParams struct {

	/* ProjectID.

	   projectId
	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the project delete using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectDeleteUsingDELETEParams) WithDefaults() *ProjectDeleteUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the project delete using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectDeleteUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the project delete using d e l e t e params
func (o *ProjectDeleteUsingDELETEParams) WithTimeout(timeout time.Duration) *ProjectDeleteUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project delete using d e l e t e params
func (o *ProjectDeleteUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project delete using d e l e t e params
func (o *ProjectDeleteUsingDELETEParams) WithContext(ctx context.Context) *ProjectDeleteUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project delete using d e l e t e params
func (o *ProjectDeleteUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project delete using d e l e t e params
func (o *ProjectDeleteUsingDELETEParams) WithHTTPClient(client *http.Client) *ProjectDeleteUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project delete using d e l e t e params
func (o *ProjectDeleteUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the project delete using d e l e t e params
func (o *ProjectDeleteUsingDELETEParams) WithProjectID(projectID string) *ProjectDeleteUsingDELETEParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the project delete using d e l e t e params
func (o *ProjectDeleteUsingDELETEParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectDeleteUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
