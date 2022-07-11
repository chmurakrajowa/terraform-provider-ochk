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

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// NewAllocationCreateUsingPUTParams creates a new AllocationCreateUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAllocationCreateUsingPUTParams() *AllocationCreateUsingPUTParams {
	return &AllocationCreateUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAllocationCreateUsingPUTParamsWithTimeout creates a new AllocationCreateUsingPUTParams object
// with the ability to set a timeout on a request.
func NewAllocationCreateUsingPUTParamsWithTimeout(timeout time.Duration) *AllocationCreateUsingPUTParams {
	return &AllocationCreateUsingPUTParams{
		timeout: timeout,
	}
}

// NewAllocationCreateUsingPUTParamsWithContext creates a new AllocationCreateUsingPUTParams object
// with the ability to set a context for a request.
func NewAllocationCreateUsingPUTParamsWithContext(ctx context.Context) *AllocationCreateUsingPUTParams {
	return &AllocationCreateUsingPUTParams{
		Context: ctx,
	}
}

// NewAllocationCreateUsingPUTParamsWithHTTPClient creates a new AllocationCreateUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewAllocationCreateUsingPUTParamsWithHTTPClient(client *http.Client) *AllocationCreateUsingPUTParams {
	return &AllocationCreateUsingPUTParams{
		HTTPClient: client,
	}
}

/* AllocationCreateUsingPUTParams contains all the parameters to send to the API endpoint
   for the allocation create using p u t operation.

   Typically these are written to a http.Request.
*/
type AllocationCreateUsingPUTParams struct {

	/* PublicIPAllocation.

	   publicIpAllocation
	*/
	PublicIPAllocation *models.PublicIPAllocation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the allocation create using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllocationCreateUsingPUTParams) WithDefaults() *AllocationCreateUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the allocation create using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllocationCreateUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the allocation create using p u t params
func (o *AllocationCreateUsingPUTParams) WithTimeout(timeout time.Duration) *AllocationCreateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the allocation create using p u t params
func (o *AllocationCreateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the allocation create using p u t params
func (o *AllocationCreateUsingPUTParams) WithContext(ctx context.Context) *AllocationCreateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the allocation create using p u t params
func (o *AllocationCreateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the allocation create using p u t params
func (o *AllocationCreateUsingPUTParams) WithHTTPClient(client *http.Client) *AllocationCreateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the allocation create using p u t params
func (o *AllocationCreateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPublicIPAllocation adds the publicIPAllocation to the allocation create using p u t params
func (o *AllocationCreateUsingPUTParams) WithPublicIPAllocation(publicIPAllocation *models.PublicIPAllocation) *AllocationCreateUsingPUTParams {
	o.SetPublicIPAllocation(publicIPAllocation)
	return o
}

// SetPublicIPAllocation adds the publicIpAllocation to the allocation create using p u t params
func (o *AllocationCreateUsingPUTParams) SetPublicIPAllocation(publicIPAllocation *models.PublicIPAllocation) {
	o.PublicIPAllocation = publicIPAllocation
}

// WriteToRequest writes these params to a swagger request
func (o *AllocationCreateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PublicIPAllocation != nil {
		if err := r.SetBodyParam(o.PublicIPAllocation); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
