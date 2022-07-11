// Code generated by go-swagger; DO NOT EDIT.

package ip_a_m_public_ip_allocations

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

// NewAllocationDeleteUsingDELETEParams creates a new AllocationDeleteUsingDELETEParams object
// with the default values initialized.
func NewAllocationDeleteUsingDELETEParams() *AllocationDeleteUsingDELETEParams {
	var ()
	return &AllocationDeleteUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAllocationDeleteUsingDELETEParamsWithTimeout creates a new AllocationDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAllocationDeleteUsingDELETEParamsWithTimeout(timeout time.Duration) *AllocationDeleteUsingDELETEParams {
	var ()
	return &AllocationDeleteUsingDELETEParams{

		timeout: timeout,
	}
}

// NewAllocationDeleteUsingDELETEParamsWithContext creates a new AllocationDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewAllocationDeleteUsingDELETEParamsWithContext(ctx context.Context) *AllocationDeleteUsingDELETEParams {
	var ()
	return &AllocationDeleteUsingDELETEParams{

		Context: ctx,
	}
}

// NewAllocationDeleteUsingDELETEParamsWithHTTPClient creates a new AllocationDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAllocationDeleteUsingDELETEParamsWithHTTPClient(client *http.Client) *AllocationDeleteUsingDELETEParams {
	var ()
	return &AllocationDeleteUsingDELETEParams{
		HTTPClient: client,
	}
}

/*AllocationDeleteUsingDELETEParams contains all the parameters to send to the API endpoint
for the allocation delete using d e l e t e operation typically these are written to a http.Request
*/
type AllocationDeleteUsingDELETEParams struct {

	/*AllocationID
	  allocationId

	*/
	AllocationID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the allocation delete using d e l e t e params
func (o *AllocationDeleteUsingDELETEParams) WithTimeout(timeout time.Duration) *AllocationDeleteUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the allocation delete using d e l e t e params
func (o *AllocationDeleteUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the allocation delete using d e l e t e params
func (o *AllocationDeleteUsingDELETEParams) WithContext(ctx context.Context) *AllocationDeleteUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the allocation delete using d e l e t e params
func (o *AllocationDeleteUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the allocation delete using d e l e t e params
func (o *AllocationDeleteUsingDELETEParams) WithHTTPClient(client *http.Client) *AllocationDeleteUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the allocation delete using d e l e t e params
func (o *AllocationDeleteUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllocationID adds the allocationID to the allocation delete using d e l e t e params
func (o *AllocationDeleteUsingDELETEParams) WithAllocationID(allocationID int32) *AllocationDeleteUsingDELETEParams {
	o.SetAllocationID(allocationID)
	return o
}

// SetAllocationID adds the allocationId to the allocation delete using d e l e t e params
func (o *AllocationDeleteUsingDELETEParams) SetAllocationID(allocationID int32) {
	o.AllocationID = allocationID
}

// WriteToRequest writes these params to a swagger request
func (o *AllocationDeleteUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param allocationId
	if err := r.SetPathParam("allocationId", swag.FormatInt32(o.AllocationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}