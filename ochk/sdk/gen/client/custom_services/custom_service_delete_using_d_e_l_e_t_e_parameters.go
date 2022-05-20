// Code generated by go-swagger; DO NOT EDIT.

package custom_services

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

// NewCustomServiceDeleteUsingDELETEParams creates a new CustomServiceDeleteUsingDELETEParams object
// with the default values initialized.
func NewCustomServiceDeleteUsingDELETEParams() *CustomServiceDeleteUsingDELETEParams {
	var ()
	return &CustomServiceDeleteUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomServiceDeleteUsingDELETEParamsWithTimeout creates a new CustomServiceDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomServiceDeleteUsingDELETEParamsWithTimeout(timeout time.Duration) *CustomServiceDeleteUsingDELETEParams {
	var ()
	return &CustomServiceDeleteUsingDELETEParams{

		timeout: timeout,
	}
}

// NewCustomServiceDeleteUsingDELETEParamsWithContext creates a new CustomServiceDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomServiceDeleteUsingDELETEParamsWithContext(ctx context.Context) *CustomServiceDeleteUsingDELETEParams {
	var ()
	return &CustomServiceDeleteUsingDELETEParams{

		Context: ctx,
	}
}

// NewCustomServiceDeleteUsingDELETEParamsWithHTTPClient creates a new CustomServiceDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomServiceDeleteUsingDELETEParamsWithHTTPClient(client *http.Client) *CustomServiceDeleteUsingDELETEParams {
	var ()
	return &CustomServiceDeleteUsingDELETEParams{
		HTTPClient: client,
	}
}

/*CustomServiceDeleteUsingDELETEParams contains all the parameters to send to the API endpoint
for the custom service delete using d e l e t e operation typically these are written to a http.Request
*/
type CustomServiceDeleteUsingDELETEParams struct {

	/*ServiceID
	  serviceId

	*/
	ServiceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom service delete using d e l e t e params
func (o *CustomServiceDeleteUsingDELETEParams) WithTimeout(timeout time.Duration) *CustomServiceDeleteUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom service delete using d e l e t e params
func (o *CustomServiceDeleteUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom service delete using d e l e t e params
func (o *CustomServiceDeleteUsingDELETEParams) WithContext(ctx context.Context) *CustomServiceDeleteUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom service delete using d e l e t e params
func (o *CustomServiceDeleteUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom service delete using d e l e t e params
func (o *CustomServiceDeleteUsingDELETEParams) WithHTTPClient(client *http.Client) *CustomServiceDeleteUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom service delete using d e l e t e params
func (o *CustomServiceDeleteUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceID adds the serviceID to the custom service delete using d e l e t e params
func (o *CustomServiceDeleteUsingDELETEParams) WithServiceID(serviceID string) *CustomServiceDeleteUsingDELETEParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the custom service delete using d e l e t e params
func (o *CustomServiceDeleteUsingDELETEParams) SetServiceID(serviceID string) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomServiceDeleteUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serviceId
	if err := r.SetPathParam("serviceId", o.ServiceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
