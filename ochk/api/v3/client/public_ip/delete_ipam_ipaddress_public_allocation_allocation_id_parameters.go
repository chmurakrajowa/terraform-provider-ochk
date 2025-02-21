// Code generated by go-swagger; DO NOT EDIT.

package public_ip

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

// NewDeleteIpamIpaddressPublicAllocationAllocationIDParams creates a new DeleteIpamIpaddressPublicAllocationAllocationIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteIpamIpaddressPublicAllocationAllocationIDParams() *DeleteIpamIpaddressPublicAllocationAllocationIDParams {
	return &DeleteIpamIpaddressPublicAllocationAllocationIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIpamIpaddressPublicAllocationAllocationIDParamsWithTimeout creates a new DeleteIpamIpaddressPublicAllocationAllocationIDParams object
// with the ability to set a timeout on a request.
func NewDeleteIpamIpaddressPublicAllocationAllocationIDParamsWithTimeout(timeout time.Duration) *DeleteIpamIpaddressPublicAllocationAllocationIDParams {
	return &DeleteIpamIpaddressPublicAllocationAllocationIDParams{
		timeout: timeout,
	}
}

// NewDeleteIpamIpaddressPublicAllocationAllocationIDParamsWithContext creates a new DeleteIpamIpaddressPublicAllocationAllocationIDParams object
// with the ability to set a context for a request.
func NewDeleteIpamIpaddressPublicAllocationAllocationIDParamsWithContext(ctx context.Context) *DeleteIpamIpaddressPublicAllocationAllocationIDParams {
	return &DeleteIpamIpaddressPublicAllocationAllocationIDParams{
		Context: ctx,
	}
}

// NewDeleteIpamIpaddressPublicAllocationAllocationIDParamsWithHTTPClient creates a new DeleteIpamIpaddressPublicAllocationAllocationIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteIpamIpaddressPublicAllocationAllocationIDParamsWithHTTPClient(client *http.Client) *DeleteIpamIpaddressPublicAllocationAllocationIDParams {
	return &DeleteIpamIpaddressPublicAllocationAllocationIDParams{
		HTTPClient: client,
	}
}

/*
DeleteIpamIpaddressPublicAllocationAllocationIDParams contains all the parameters to send to the API endpoint

	for the delete ipam ipaddress public allocation allocation ID operation.

	Typically these are written to a http.Request.
*/
type DeleteIpamIpaddressPublicAllocationAllocationIDParams struct {

	// AllocationID.
	//
	// Format: int32
	AllocationID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete ipam ipaddress public allocation allocation ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIpamIpaddressPublicAllocationAllocationIDParams) WithDefaults() *DeleteIpamIpaddressPublicAllocationAllocationIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete ipam ipaddress public allocation allocation ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIpamIpaddressPublicAllocationAllocationIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete ipam ipaddress public allocation allocation ID params
func (o *DeleteIpamIpaddressPublicAllocationAllocationIDParams) WithTimeout(timeout time.Duration) *DeleteIpamIpaddressPublicAllocationAllocationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete ipam ipaddress public allocation allocation ID params
func (o *DeleteIpamIpaddressPublicAllocationAllocationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete ipam ipaddress public allocation allocation ID params
func (o *DeleteIpamIpaddressPublicAllocationAllocationIDParams) WithContext(ctx context.Context) *DeleteIpamIpaddressPublicAllocationAllocationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete ipam ipaddress public allocation allocation ID params
func (o *DeleteIpamIpaddressPublicAllocationAllocationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete ipam ipaddress public allocation allocation ID params
func (o *DeleteIpamIpaddressPublicAllocationAllocationIDParams) WithHTTPClient(client *http.Client) *DeleteIpamIpaddressPublicAllocationAllocationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete ipam ipaddress public allocation allocation ID params
func (o *DeleteIpamIpaddressPublicAllocationAllocationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllocationID adds the allocationID to the delete ipam ipaddress public allocation allocation ID params
func (o *DeleteIpamIpaddressPublicAllocationAllocationIDParams) WithAllocationID(allocationID int32) *DeleteIpamIpaddressPublicAllocationAllocationIDParams {
	o.SetAllocationID(allocationID)
	return o
}

// SetAllocationID adds the allocationId to the delete ipam ipaddress public allocation allocation ID params
func (o *DeleteIpamIpaddressPublicAllocationAllocationIDParams) SetAllocationID(allocationID int32) {
	o.AllocationID = allocationID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIpamIpaddressPublicAllocationAllocationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
