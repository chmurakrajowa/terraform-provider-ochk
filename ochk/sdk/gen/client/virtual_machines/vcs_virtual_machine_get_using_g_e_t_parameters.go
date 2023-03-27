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

// NewVcsVirtualMachineGetUsingGETParams creates a new VcsVirtualMachineGetUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVcsVirtualMachineGetUsingGETParams() *VcsVirtualMachineGetUsingGETParams {
	return &VcsVirtualMachineGetUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVcsVirtualMachineGetUsingGETParamsWithTimeout creates a new VcsVirtualMachineGetUsingGETParams object
// with the ability to set a timeout on a request.
func NewVcsVirtualMachineGetUsingGETParamsWithTimeout(timeout time.Duration) *VcsVirtualMachineGetUsingGETParams {
	return &VcsVirtualMachineGetUsingGETParams{
		timeout: timeout,
	}
}

// NewVcsVirtualMachineGetUsingGETParamsWithContext creates a new VcsVirtualMachineGetUsingGETParams object
// with the ability to set a context for a request.
func NewVcsVirtualMachineGetUsingGETParamsWithContext(ctx context.Context) *VcsVirtualMachineGetUsingGETParams {
	return &VcsVirtualMachineGetUsingGETParams{
		Context: ctx,
	}
}

// NewVcsVirtualMachineGetUsingGETParamsWithHTTPClient creates a new VcsVirtualMachineGetUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewVcsVirtualMachineGetUsingGETParamsWithHTTPClient(client *http.Client) *VcsVirtualMachineGetUsingGETParams {
	return &VcsVirtualMachineGetUsingGETParams{
		HTTPClient: client,
	}
}

/*
VcsVirtualMachineGetUsingGETParams contains all the parameters to send to the API endpoint

	for the vcs virtual machine get using g e t operation.

	Typically these are written to a http.Request.
*/
type VcsVirtualMachineGetUsingGETParams struct {

	/* VirtualMachineID.

	   virtualMachineId
	*/
	VirtualMachineID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vcs virtual machine get using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VcsVirtualMachineGetUsingGETParams) WithDefaults() *VcsVirtualMachineGetUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vcs virtual machine get using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VcsVirtualMachineGetUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the vcs virtual machine get using g e t params
func (o *VcsVirtualMachineGetUsingGETParams) WithTimeout(timeout time.Duration) *VcsVirtualMachineGetUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vcs virtual machine get using g e t params
func (o *VcsVirtualMachineGetUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vcs virtual machine get using g e t params
func (o *VcsVirtualMachineGetUsingGETParams) WithContext(ctx context.Context) *VcsVirtualMachineGetUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vcs virtual machine get using g e t params
func (o *VcsVirtualMachineGetUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vcs virtual machine get using g e t params
func (o *VcsVirtualMachineGetUsingGETParams) WithHTTPClient(client *http.Client) *VcsVirtualMachineGetUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vcs virtual machine get using g e t params
func (o *VcsVirtualMachineGetUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVirtualMachineID adds the virtualMachineID to the vcs virtual machine get using g e t params
func (o *VcsVirtualMachineGetUsingGETParams) WithVirtualMachineID(virtualMachineID string) *VcsVirtualMachineGetUsingGETParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the vcs virtual machine get using g e t params
func (o *VcsVirtualMachineGetUsingGETParams) SetVirtualMachineID(virtualMachineID string) {
	o.VirtualMachineID = virtualMachineID
}

// WriteToRequest writes these params to a swagger request
func (o *VcsVirtualMachineGetUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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