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

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// NewVcsVirtualMachineUpdateUsingPUTParams creates a new VcsVirtualMachineUpdateUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVcsVirtualMachineUpdateUsingPUTParams() *VcsVirtualMachineUpdateUsingPUTParams {
	return &VcsVirtualMachineUpdateUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVcsVirtualMachineUpdateUsingPUTParamsWithTimeout creates a new VcsVirtualMachineUpdateUsingPUTParams object
// with the ability to set a timeout on a request.
func NewVcsVirtualMachineUpdateUsingPUTParamsWithTimeout(timeout time.Duration) *VcsVirtualMachineUpdateUsingPUTParams {
	return &VcsVirtualMachineUpdateUsingPUTParams{
		timeout: timeout,
	}
}

// NewVcsVirtualMachineUpdateUsingPUTParamsWithContext creates a new VcsVirtualMachineUpdateUsingPUTParams object
// with the ability to set a context for a request.
func NewVcsVirtualMachineUpdateUsingPUTParamsWithContext(ctx context.Context) *VcsVirtualMachineUpdateUsingPUTParams {
	return &VcsVirtualMachineUpdateUsingPUTParams{
		Context: ctx,
	}
}

// NewVcsVirtualMachineUpdateUsingPUTParamsWithHTTPClient creates a new VcsVirtualMachineUpdateUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewVcsVirtualMachineUpdateUsingPUTParamsWithHTTPClient(client *http.Client) *VcsVirtualMachineUpdateUsingPUTParams {
	return &VcsVirtualMachineUpdateUsingPUTParams{
		HTTPClient: client,
	}
}

/*
VcsVirtualMachineUpdateUsingPUTParams contains all the parameters to send to the API endpoint

	for the vcs virtual machine update using p u t operation.

	Typically these are written to a http.Request.
*/
type VcsVirtualMachineUpdateUsingPUTParams struct {

	/* VirtualMachine.

	   virtualMachine
	*/
	VirtualMachine *models.VcsVirtualMachineInstance

	/* VirtualMachineID.

	   virtualMachineId
	*/
	VirtualMachineID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vcs virtual machine update using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VcsVirtualMachineUpdateUsingPUTParams) WithDefaults() *VcsVirtualMachineUpdateUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vcs virtual machine update using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VcsVirtualMachineUpdateUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the vcs virtual machine update using p u t params
func (o *VcsVirtualMachineUpdateUsingPUTParams) WithTimeout(timeout time.Duration) *VcsVirtualMachineUpdateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vcs virtual machine update using p u t params
func (o *VcsVirtualMachineUpdateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vcs virtual machine update using p u t params
func (o *VcsVirtualMachineUpdateUsingPUTParams) WithContext(ctx context.Context) *VcsVirtualMachineUpdateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vcs virtual machine update using p u t params
func (o *VcsVirtualMachineUpdateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vcs virtual machine update using p u t params
func (o *VcsVirtualMachineUpdateUsingPUTParams) WithHTTPClient(client *http.Client) *VcsVirtualMachineUpdateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vcs virtual machine update using p u t params
func (o *VcsVirtualMachineUpdateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVirtualMachine adds the virtualMachine to the vcs virtual machine update using p u t params
func (o *VcsVirtualMachineUpdateUsingPUTParams) WithVirtualMachine(virtualMachine *models.VcsVirtualMachineInstance) *VcsVirtualMachineUpdateUsingPUTParams {
	o.SetVirtualMachine(virtualMachine)
	return o
}

// SetVirtualMachine adds the virtualMachine to the vcs virtual machine update using p u t params
func (o *VcsVirtualMachineUpdateUsingPUTParams) SetVirtualMachine(virtualMachine *models.VcsVirtualMachineInstance) {
	o.VirtualMachine = virtualMachine
}

// WithVirtualMachineID adds the virtualMachineID to the vcs virtual machine update using p u t params
func (o *VcsVirtualMachineUpdateUsingPUTParams) WithVirtualMachineID(virtualMachineID string) *VcsVirtualMachineUpdateUsingPUTParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the vcs virtual machine update using p u t params
func (o *VcsVirtualMachineUpdateUsingPUTParams) SetVirtualMachineID(virtualMachineID string) {
	o.VirtualMachineID = virtualMachineID
}

// WriteToRequest writes these params to a swagger request
func (o *VcsVirtualMachineUpdateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.VirtualMachine != nil {
		if err := r.SetBodyParam(o.VirtualMachine); err != nil {
			return err
		}
	}

	// path param virtualMachineId
	if err := r.SetPathParam("virtualMachineId", o.VirtualMachineID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
