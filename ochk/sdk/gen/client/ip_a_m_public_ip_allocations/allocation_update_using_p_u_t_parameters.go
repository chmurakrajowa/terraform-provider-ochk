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

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// NewAllocationUpdateUsingPUTParams creates a new AllocationUpdateUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAllocationUpdateUsingPUTParams() *AllocationUpdateUsingPUTParams {
	return &AllocationUpdateUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAllocationUpdateUsingPUTParamsWithTimeout creates a new AllocationUpdateUsingPUTParams object
// with the ability to set a timeout on a request.
func NewAllocationUpdateUsingPUTParamsWithTimeout(timeout time.Duration) *AllocationUpdateUsingPUTParams {
	return &AllocationUpdateUsingPUTParams{
		timeout: timeout,
	}
}

// NewAllocationUpdateUsingPUTParamsWithContext creates a new AllocationUpdateUsingPUTParams object
// with the ability to set a context for a request.
func NewAllocationUpdateUsingPUTParamsWithContext(ctx context.Context) *AllocationUpdateUsingPUTParams {
	return &AllocationUpdateUsingPUTParams{
		Context: ctx,
	}
}

// NewAllocationUpdateUsingPUTParamsWithHTTPClient creates a new AllocationUpdateUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewAllocationUpdateUsingPUTParamsWithHTTPClient(client *http.Client) *AllocationUpdateUsingPUTParams {
	return &AllocationUpdateUsingPUTParams{
		HTTPClient: client,
	}
}

/* AllocationUpdateUsingPUTParams contains all the parameters to send to the API endpoint
   for the allocation update using p u t operation.

   Typically these are written to a http.Request.
*/
type AllocationUpdateUsingPUTParams struct {

	/* AllocationID.

	   allocationId

	   Format: int32
	*/
	AllocationID int32

	/* PublicIPAllocation.

	   publicIpAllocation
	*/
	PublicIPAllocation *models.PublicIPAllocation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the allocation update using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllocationUpdateUsingPUTParams) WithDefaults() *AllocationUpdateUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the allocation update using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllocationUpdateUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the allocation update using p u t params
func (o *AllocationUpdateUsingPUTParams) WithTimeout(timeout time.Duration) *AllocationUpdateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the allocation update using p u t params
func (o *AllocationUpdateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the allocation update using p u t params
func (o *AllocationUpdateUsingPUTParams) WithContext(ctx context.Context) *AllocationUpdateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the allocation update using p u t params
func (o *AllocationUpdateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the allocation update using p u t params
func (o *AllocationUpdateUsingPUTParams) WithHTTPClient(client *http.Client) *AllocationUpdateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the allocation update using p u t params
func (o *AllocationUpdateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllocationID adds the allocationID to the allocation update using p u t params
func (o *AllocationUpdateUsingPUTParams) WithAllocationID(allocationID int32) *AllocationUpdateUsingPUTParams {
	o.SetAllocationID(allocationID)
	return o
}

// SetAllocationID adds the allocationId to the allocation update using p u t params
func (o *AllocationUpdateUsingPUTParams) SetAllocationID(allocationID int32) {
	o.AllocationID = allocationID
}

// WithPublicIPAllocation adds the publicIPAllocation to the allocation update using p u t params
func (o *AllocationUpdateUsingPUTParams) WithPublicIPAllocation(publicIPAllocation *models.PublicIPAllocation) *AllocationUpdateUsingPUTParams {
	o.SetPublicIPAllocation(publicIPAllocation)
	return o
}

// SetPublicIPAllocation adds the publicIpAllocation to the allocation update using p u t params
func (o *AllocationUpdateUsingPUTParams) SetPublicIPAllocation(publicIPAllocation *models.PublicIPAllocation) {
	o.PublicIPAllocation = publicIPAllocation
}

// WriteToRequest writes these params to a swagger request
func (o *AllocationUpdateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param allocationId
	if err := r.SetPathParam("allocationId", swag.FormatInt32(o.AllocationID)); err != nil {
		return err
	}
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
