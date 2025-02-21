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

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
)

// NewPutVcsVirtualMachinesVirtualMachineIDSnapshotsParams creates a new PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutVcsVirtualMachinesVirtualMachineIDSnapshotsParams() *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams {
	return &PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutVcsVirtualMachinesVirtualMachineIDSnapshotsParamsWithTimeout creates a new PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams object
// with the ability to set a timeout on a request.
func NewPutVcsVirtualMachinesVirtualMachineIDSnapshotsParamsWithTimeout(timeout time.Duration) *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams {
	return &PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams{
		timeout: timeout,
	}
}

// NewPutVcsVirtualMachinesVirtualMachineIDSnapshotsParamsWithContext creates a new PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams object
// with the ability to set a context for a request.
func NewPutVcsVirtualMachinesVirtualMachineIDSnapshotsParamsWithContext(ctx context.Context) *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams {
	return &PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams{
		Context: ctx,
	}
}

// NewPutVcsVirtualMachinesVirtualMachineIDSnapshotsParamsWithHTTPClient creates a new PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutVcsVirtualMachinesVirtualMachineIDSnapshotsParamsWithHTTPClient(client *http.Client) *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams {
	return &PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams{
		HTTPClient: client,
	}
}

/*
PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams contains all the parameters to send to the API endpoint

	for the put vcs virtual machines virtual machine ID snapshots operation.

	Typically these are written to a http.Request.
*/
type PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams struct {

	// Body.
	Body *models.SnapshotInstance

	// VirtualMachineID.
	//
	// Format: uuid
	VirtualMachineID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put vcs virtual machines virtual machine ID snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams) WithDefaults() *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put vcs virtual machines virtual machine ID snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put vcs virtual machines virtual machine ID snapshots params
func (o *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams) WithTimeout(timeout time.Duration) *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put vcs virtual machines virtual machine ID snapshots params
func (o *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put vcs virtual machines virtual machine ID snapshots params
func (o *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams) WithContext(ctx context.Context) *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put vcs virtual machines virtual machine ID snapshots params
func (o *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put vcs virtual machines virtual machine ID snapshots params
func (o *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams) WithHTTPClient(client *http.Client) *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put vcs virtual machines virtual machine ID snapshots params
func (o *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put vcs virtual machines virtual machine ID snapshots params
func (o *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams) WithBody(body *models.SnapshotInstance) *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put vcs virtual machines virtual machine ID snapshots params
func (o *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams) SetBody(body *models.SnapshotInstance) {
	o.Body = body
}

// WithVirtualMachineID adds the virtualMachineID to the put vcs virtual machines virtual machine ID snapshots params
func (o *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams) WithVirtualMachineID(virtualMachineID strfmt.UUID) *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the put vcs virtual machines virtual machine ID snapshots params
func (o *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams) SetVirtualMachineID(virtualMachineID strfmt.UUID) {
	o.VirtualMachineID = virtualMachineID
}

// WriteToRequest writes these params to a swagger request
func (o *PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
