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

// NewPostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams creates a new PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams() *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams {
	return &PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParamsWithTimeout creates a new PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams object
// with the ability to set a timeout on a request.
func NewPostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParamsWithTimeout(timeout time.Duration) *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams {
	return &PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams{
		timeout: timeout,
	}
}

// NewPostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParamsWithContext creates a new PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams object
// with the ability to set a context for a request.
func NewPostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParamsWithContext(ctx context.Context) *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams {
	return &PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams{
		Context: ctx,
	}
}

// NewPostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParamsWithHTTPClient creates a new PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParamsWithHTTPClient(client *http.Client) *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams {
	return &PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams{
		HTTPClient: client,
	}
}

/*
PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams contains all the parameters to send to the API endpoint

	for the post vcs virtual machines virtual machine ID remove all snapshots operation.

	Typically these are written to a http.Request.
*/
type PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams struct {

	// VirtualMachineID.
	//
	// Format: uuid
	VirtualMachineID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post vcs virtual machines virtual machine ID remove all snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams) WithDefaults() *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post vcs virtual machines virtual machine ID remove all snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post vcs virtual machines virtual machine ID remove all snapshots params
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams) WithTimeout(timeout time.Duration) *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post vcs virtual machines virtual machine ID remove all snapshots params
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post vcs virtual machines virtual machine ID remove all snapshots params
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams) WithContext(ctx context.Context) *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post vcs virtual machines virtual machine ID remove all snapshots params
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post vcs virtual machines virtual machine ID remove all snapshots params
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams) WithHTTPClient(client *http.Client) *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post vcs virtual machines virtual machine ID remove all snapshots params
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVirtualMachineID adds the virtualMachineID to the post vcs virtual machines virtual machine ID remove all snapshots params
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams) WithVirtualMachineID(virtualMachineID strfmt.UUID) *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the post vcs virtual machines virtual machine ID remove all snapshots params
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams) SetVirtualMachineID(virtualMachineID strfmt.UUID) {
	o.VirtualMachineID = virtualMachineID
}

// WriteToRequest writes these params to a swagger request
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param virtualMachineId
	if err := r.SetPathParam("virtualMachineId", o.VirtualMachineID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
