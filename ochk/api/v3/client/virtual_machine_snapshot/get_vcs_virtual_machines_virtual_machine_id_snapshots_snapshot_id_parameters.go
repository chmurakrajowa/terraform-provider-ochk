// Code generated by go-swagger; DO NOT EDIT.

package virtual_machine_snapshot

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

// NewGetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams creates a new GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams() *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams {
	return &GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParamsWithTimeout creates a new GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams object
// with the ability to set a timeout on a request.
func NewGetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParamsWithTimeout(timeout time.Duration) *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams {
	return &GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams{
		timeout: timeout,
	}
}

// NewGetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParamsWithContext creates a new GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams object
// with the ability to set a context for a request.
func NewGetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParamsWithContext(ctx context.Context) *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams {
	return &GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams{
		Context: ctx,
	}
}

// NewGetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParamsWithHTTPClient creates a new GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParamsWithHTTPClient(client *http.Client) *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams {
	return &GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams{
		HTTPClient: client,
	}
}

/*
GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams contains all the parameters to send to the API endpoint

	for the get vcs virtual machines virtual machine ID snapshots snapshot ID operation.

	Typically these are written to a http.Request.
*/
type GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams struct {

	// SnapshotID.
	//
	// Format: uuid
	SnapshotID strfmt.UUID

	// VirtualMachineID.
	//
	// Format: uuid
	VirtualMachineID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vcs virtual machines virtual machine ID snapshots snapshot ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams) WithDefaults() *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vcs virtual machines virtual machine ID snapshots snapshot ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vcs virtual machines virtual machine ID snapshots snapshot ID params
func (o *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams) WithTimeout(timeout time.Duration) *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vcs virtual machines virtual machine ID snapshots snapshot ID params
func (o *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vcs virtual machines virtual machine ID snapshots snapshot ID params
func (o *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams) WithContext(ctx context.Context) *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vcs virtual machines virtual machine ID snapshots snapshot ID params
func (o *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vcs virtual machines virtual machine ID snapshots snapshot ID params
func (o *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams) WithHTTPClient(client *http.Client) *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vcs virtual machines virtual machine ID snapshots snapshot ID params
func (o *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSnapshotID adds the snapshotID to the get vcs virtual machines virtual machine ID snapshots snapshot ID params
func (o *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams) WithSnapshotID(snapshotID strfmt.UUID) *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams {
	o.SetSnapshotID(snapshotID)
	return o
}

// SetSnapshotID adds the snapshotId to the get vcs virtual machines virtual machine ID snapshots snapshot ID params
func (o *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams) SetSnapshotID(snapshotID strfmt.UUID) {
	o.SnapshotID = snapshotID
}

// WithVirtualMachineID adds the virtualMachineID to the get vcs virtual machines virtual machine ID snapshots snapshot ID params
func (o *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams) WithVirtualMachineID(virtualMachineID strfmt.UUID) *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the get vcs virtual machines virtual machine ID snapshots snapshot ID params
func (o *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams) SetVirtualMachineID(virtualMachineID strfmt.UUID) {
	o.VirtualMachineID = virtualMachineID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param snapshotId
	if err := r.SetPathParam("snapshotId", o.SnapshotID.String()); err != nil {
		return err
	}

	// path param virtualMachineId
	if err := r.SetPathParam("virtualMachineId", o.VirtualMachineID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
