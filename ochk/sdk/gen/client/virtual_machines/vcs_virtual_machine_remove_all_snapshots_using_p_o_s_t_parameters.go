// Code generated by go-swagger; DO NOT EDIT.

package virtual_machines

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

// NewVcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams creates a new VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams() *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams {
	return &VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVcsVirtualMachineRemoveAllSnapshotsUsingPOSTParamsWithTimeout creates a new VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewVcsVirtualMachineRemoveAllSnapshotsUsingPOSTParamsWithTimeout(timeout time.Duration) *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams {
	return &VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams{
		timeout: timeout,
	}
}

// NewVcsVirtualMachineRemoveAllSnapshotsUsingPOSTParamsWithContext creates a new VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams object
// with the ability to set a context for a request.
func NewVcsVirtualMachineRemoveAllSnapshotsUsingPOSTParamsWithContext(ctx context.Context) *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams {
	return &VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams{
		Context: ctx,
	}
}

// NewVcsVirtualMachineRemoveAllSnapshotsUsingPOSTParamsWithHTTPClient creates a new VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewVcsVirtualMachineRemoveAllSnapshotsUsingPOSTParamsWithHTTPClient(client *http.Client) *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams {
	return &VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams contains all the parameters to send to the API endpoint

	for the vcs virtual machine remove all snapshots using p o s t operation.

	Typically these are written to a http.Request.
*/
type VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams struct {

	/* VirtualMachineID.

	   virtualMachineId
	*/
	VirtualMachineID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vcs virtual machine remove all snapshots using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams) WithDefaults() *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vcs virtual machine remove all snapshots using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the vcs virtual machine remove all snapshots using p o s t params
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams) WithTimeout(timeout time.Duration) *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vcs virtual machine remove all snapshots using p o s t params
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vcs virtual machine remove all snapshots using p o s t params
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams) WithContext(ctx context.Context) *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vcs virtual machine remove all snapshots using p o s t params
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vcs virtual machine remove all snapshots using p o s t params
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams) WithHTTPClient(client *http.Client) *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vcs virtual machine remove all snapshots using p o s t params
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVirtualMachineID adds the virtualMachineID to the vcs virtual machine remove all snapshots using p o s t params
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams) WithVirtualMachineID(virtualMachineID string) *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the vcs virtual machine remove all snapshots using p o s t params
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams) SetVirtualMachineID(virtualMachineID string) {
	o.VirtualMachineID = virtualMachineID
}

// WriteToRequest writes these params to a swagger request
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param virtualMachineId
	if err := r.SetPathParam("virtualMachineId", o.VirtualMachineID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}