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

// NewVcsVirtualMachineResetUsingPOSTParams creates a new VcsVirtualMachineResetUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVcsVirtualMachineResetUsingPOSTParams() *VcsVirtualMachineResetUsingPOSTParams {
	return &VcsVirtualMachineResetUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVcsVirtualMachineResetUsingPOSTParamsWithTimeout creates a new VcsVirtualMachineResetUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewVcsVirtualMachineResetUsingPOSTParamsWithTimeout(timeout time.Duration) *VcsVirtualMachineResetUsingPOSTParams {
	return &VcsVirtualMachineResetUsingPOSTParams{
		timeout: timeout,
	}
}

// NewVcsVirtualMachineResetUsingPOSTParamsWithContext creates a new VcsVirtualMachineResetUsingPOSTParams object
// with the ability to set a context for a request.
func NewVcsVirtualMachineResetUsingPOSTParamsWithContext(ctx context.Context) *VcsVirtualMachineResetUsingPOSTParams {
	return &VcsVirtualMachineResetUsingPOSTParams{
		Context: ctx,
	}
}

// NewVcsVirtualMachineResetUsingPOSTParamsWithHTTPClient creates a new VcsVirtualMachineResetUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewVcsVirtualMachineResetUsingPOSTParamsWithHTTPClient(client *http.Client) *VcsVirtualMachineResetUsingPOSTParams {
	return &VcsVirtualMachineResetUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
VcsVirtualMachineResetUsingPOSTParams contains all the parameters to send to the API endpoint

	for the vcs virtual machine reset using p o s t operation.

	Typically these are written to a http.Request.
*/
type VcsVirtualMachineResetUsingPOSTParams struct {

	/* VirtualMachineID.

	   virtualMachineId
	*/
	VirtualMachineID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vcs virtual machine reset using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VcsVirtualMachineResetUsingPOSTParams) WithDefaults() *VcsVirtualMachineResetUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vcs virtual machine reset using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VcsVirtualMachineResetUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the vcs virtual machine reset using p o s t params
func (o *VcsVirtualMachineResetUsingPOSTParams) WithTimeout(timeout time.Duration) *VcsVirtualMachineResetUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vcs virtual machine reset using p o s t params
func (o *VcsVirtualMachineResetUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vcs virtual machine reset using p o s t params
func (o *VcsVirtualMachineResetUsingPOSTParams) WithContext(ctx context.Context) *VcsVirtualMachineResetUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vcs virtual machine reset using p o s t params
func (o *VcsVirtualMachineResetUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vcs virtual machine reset using p o s t params
func (o *VcsVirtualMachineResetUsingPOSTParams) WithHTTPClient(client *http.Client) *VcsVirtualMachineResetUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vcs virtual machine reset using p o s t params
func (o *VcsVirtualMachineResetUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVirtualMachineID adds the virtualMachineID to the vcs virtual machine reset using p o s t params
func (o *VcsVirtualMachineResetUsingPOSTParams) WithVirtualMachineID(virtualMachineID string) *VcsVirtualMachineResetUsingPOSTParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the vcs virtual machine reset using p o s t params
func (o *VcsVirtualMachineResetUsingPOSTParams) SetVirtualMachineID(virtualMachineID string) {
	o.VirtualMachineID = virtualMachineID
}

// WriteToRequest writes these params to a swagger request
func (o *VcsVirtualMachineResetUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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