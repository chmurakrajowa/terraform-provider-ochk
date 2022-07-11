// Code generated by go-swagger; DO NOT EDIT.

package system_tags

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
	"github.com/go-openapi/swag"
)

// NewSystemTagDeleteUsingDELETEParams creates a new SystemTagDeleteUsingDELETEParams object
// with the default values initialized.
func NewSystemTagDeleteUsingDELETEParams() *SystemTagDeleteUsingDELETEParams {
	var ()
	return &SystemTagDeleteUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSystemTagDeleteUsingDELETEParamsWithTimeout creates a new SystemTagDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSystemTagDeleteUsingDELETEParamsWithTimeout(timeout time.Duration) *SystemTagDeleteUsingDELETEParams {
	var ()
	return &SystemTagDeleteUsingDELETEParams{

		timeout: timeout,
	}
}

// NewSystemTagDeleteUsingDELETEParamsWithContext creates a new SystemTagDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewSystemTagDeleteUsingDELETEParamsWithContext(ctx context.Context) *SystemTagDeleteUsingDELETEParams {
	var ()
	return &SystemTagDeleteUsingDELETEParams{

		Context: ctx,
	}
}

// NewSystemTagDeleteUsingDELETEParamsWithHTTPClient creates a new SystemTagDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSystemTagDeleteUsingDELETEParamsWithHTTPClient(client *http.Client) *SystemTagDeleteUsingDELETEParams {
	var ()
	return &SystemTagDeleteUsingDELETEParams{
		HTTPClient: client,
	}
}

/*SystemTagDeleteUsingDELETEParams contains all the parameters to send to the API endpoint
for the system tag delete using d e l e t e operation typically these are written to a http.Request
*/
type SystemTagDeleteUsingDELETEParams struct {

	/*SystemTagID
	  systemTagId

	*/
	SystemTagID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the system tag delete using d e l e t e params
func (o *SystemTagDeleteUsingDELETEParams) WithTimeout(timeout time.Duration) *SystemTagDeleteUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system tag delete using d e l e t e params
func (o *SystemTagDeleteUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system tag delete using d e l e t e params
func (o *SystemTagDeleteUsingDELETEParams) WithContext(ctx context.Context) *SystemTagDeleteUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system tag delete using d e l e t e params
func (o *SystemTagDeleteUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system tag delete using d e l e t e params
func (o *SystemTagDeleteUsingDELETEParams) WithHTTPClient(client *http.Client) *SystemTagDeleteUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system tag delete using d e l e t e params
func (o *SystemTagDeleteUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSystemTagID adds the systemTagID to the system tag delete using d e l e t e params
func (o *SystemTagDeleteUsingDELETEParams) WithSystemTagID(systemTagID int32) *SystemTagDeleteUsingDELETEParams {
	o.SetSystemTagID(systemTagID)
	return o
}

// SetSystemTagID adds the systemTagId to the system tag delete using d e l e t e params
func (o *SystemTagDeleteUsingDELETEParams) SetSystemTagID(systemTagID int32) {
	o.SystemTagID = systemTagID
}

// WriteToRequest writes these params to a swagger request
func (o *SystemTagDeleteUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param systemTagId
	if err := r.SetPathParam("systemTagId", swag.FormatInt32(o.SystemTagID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
