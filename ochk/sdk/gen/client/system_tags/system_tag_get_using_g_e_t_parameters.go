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

// NewSystemTagGetUsingGETParams creates a new SystemTagGetUsingGETParams object
// with the default values initialized.
func NewSystemTagGetUsingGETParams() *SystemTagGetUsingGETParams {
	var ()
	return &SystemTagGetUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSystemTagGetUsingGETParamsWithTimeout creates a new SystemTagGetUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSystemTagGetUsingGETParamsWithTimeout(timeout time.Duration) *SystemTagGetUsingGETParams {
	var ()
	return &SystemTagGetUsingGETParams{

		timeout: timeout,
	}
}

// NewSystemTagGetUsingGETParamsWithContext creates a new SystemTagGetUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewSystemTagGetUsingGETParamsWithContext(ctx context.Context) *SystemTagGetUsingGETParams {
	var ()
	return &SystemTagGetUsingGETParams{

		Context: ctx,
	}
}

// NewSystemTagGetUsingGETParamsWithHTTPClient creates a new SystemTagGetUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSystemTagGetUsingGETParamsWithHTTPClient(client *http.Client) *SystemTagGetUsingGETParams {
	var ()
	return &SystemTagGetUsingGETParams{
		HTTPClient: client,
	}
}

/*SystemTagGetUsingGETParams contains all the parameters to send to the API endpoint
for the system tag get using g e t operation typically these are written to a http.Request
*/
type SystemTagGetUsingGETParams struct {

	/*SystemTagID
	  systemTagId

	*/
	SystemTagID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the system tag get using g e t params
func (o *SystemTagGetUsingGETParams) WithTimeout(timeout time.Duration) *SystemTagGetUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system tag get using g e t params
func (o *SystemTagGetUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system tag get using g e t params
func (o *SystemTagGetUsingGETParams) WithContext(ctx context.Context) *SystemTagGetUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system tag get using g e t params
func (o *SystemTagGetUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system tag get using g e t params
func (o *SystemTagGetUsingGETParams) WithHTTPClient(client *http.Client) *SystemTagGetUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system tag get using g e t params
func (o *SystemTagGetUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSystemTagID adds the systemTagID to the system tag get using g e t params
func (o *SystemTagGetUsingGETParams) WithSystemTagID(systemTagID int32) *SystemTagGetUsingGETParams {
	o.SetSystemTagID(systemTagID)
	return o
}

// SetSystemTagID adds the systemTagId to the system tag get using g e t params
func (o *SystemTagGetUsingGETParams) SetSystemTagID(systemTagID int32) {
	o.SystemTagID = systemTagID
}

// WriteToRequest writes these params to a swagger request
func (o *SystemTagGetUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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