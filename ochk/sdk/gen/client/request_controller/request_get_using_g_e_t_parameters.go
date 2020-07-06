// Code generated by go-swagger; DO NOT EDIT.

package request_controller

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

// NewRequestGetUsingGETParams creates a new RequestGetUsingGETParams object
// with the default values initialized.
func NewRequestGetUsingGETParams() *RequestGetUsingGETParams {
	var ()
	return &RequestGetUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRequestGetUsingGETParamsWithTimeout creates a new RequestGetUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRequestGetUsingGETParamsWithTimeout(timeout time.Duration) *RequestGetUsingGETParams {
	var ()
	return &RequestGetUsingGETParams{

		timeout: timeout,
	}
}

// NewRequestGetUsingGETParamsWithContext creates a new RequestGetUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewRequestGetUsingGETParamsWithContext(ctx context.Context) *RequestGetUsingGETParams {
	var ()
	return &RequestGetUsingGETParams{

		Context: ctx,
	}
}

// NewRequestGetUsingGETParamsWithHTTPClient creates a new RequestGetUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRequestGetUsingGETParamsWithHTTPClient(client *http.Client) *RequestGetUsingGETParams {
	var ()
	return &RequestGetUsingGETParams{
		HTTPClient: client,
	}
}

/*RequestGetUsingGETParams contains all the parameters to send to the API endpoint
for the request get using g e t operation typically these are written to a http.Request
*/
type RequestGetUsingGETParams struct {

	/*RequestID
	  requestId

	*/
	RequestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the request get using g e t params
func (o *RequestGetUsingGETParams) WithTimeout(timeout time.Duration) *RequestGetUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the request get using g e t params
func (o *RequestGetUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the request get using g e t params
func (o *RequestGetUsingGETParams) WithContext(ctx context.Context) *RequestGetUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the request get using g e t params
func (o *RequestGetUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the request get using g e t params
func (o *RequestGetUsingGETParams) WithHTTPClient(client *http.Client) *RequestGetUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the request get using g e t params
func (o *RequestGetUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestID adds the requestID to the request get using g e t params
func (o *RequestGetUsingGETParams) WithRequestID(requestID string) *RequestGetUsingGETParams {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the request get using g e t params
func (o *RequestGetUsingGETParams) SetRequestID(requestID string) {
	o.RequestID = requestID
}

// WriteToRequest writes these params to a swagger request
func (o *RequestGetUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param requestId
	if err := r.SetPathParam("requestId", o.RequestID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
