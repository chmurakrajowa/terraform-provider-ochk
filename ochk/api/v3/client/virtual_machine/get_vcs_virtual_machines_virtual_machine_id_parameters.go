// Code generated by go-swagger; DO NOT EDIT.

package virtual_machine

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

// NewGetVcsVirtualMachinesVirtualMachineIDParams creates a new GetVcsVirtualMachinesVirtualMachineIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVcsVirtualMachinesVirtualMachineIDParams() *GetVcsVirtualMachinesVirtualMachineIDParams {
	return &GetVcsVirtualMachinesVirtualMachineIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVcsVirtualMachinesVirtualMachineIDParamsWithTimeout creates a new GetVcsVirtualMachinesVirtualMachineIDParams object
// with the ability to set a timeout on a request.
func NewGetVcsVirtualMachinesVirtualMachineIDParamsWithTimeout(timeout time.Duration) *GetVcsVirtualMachinesVirtualMachineIDParams {
	return &GetVcsVirtualMachinesVirtualMachineIDParams{
		timeout: timeout,
	}
}

// NewGetVcsVirtualMachinesVirtualMachineIDParamsWithContext creates a new GetVcsVirtualMachinesVirtualMachineIDParams object
// with the ability to set a context for a request.
func NewGetVcsVirtualMachinesVirtualMachineIDParamsWithContext(ctx context.Context) *GetVcsVirtualMachinesVirtualMachineIDParams {
	return &GetVcsVirtualMachinesVirtualMachineIDParams{
		Context: ctx,
	}
}

// NewGetVcsVirtualMachinesVirtualMachineIDParamsWithHTTPClient creates a new GetVcsVirtualMachinesVirtualMachineIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVcsVirtualMachinesVirtualMachineIDParamsWithHTTPClient(client *http.Client) *GetVcsVirtualMachinesVirtualMachineIDParams {
	return &GetVcsVirtualMachinesVirtualMachineIDParams{
		HTTPClient: client,
	}
}

/*
GetVcsVirtualMachinesVirtualMachineIDParams contains all the parameters to send to the API endpoint

	for the get vcs virtual machines virtual machine ID operation.

	Typically these are written to a http.Request.
*/
type GetVcsVirtualMachinesVirtualMachineIDParams struct {

	// VirtualMachineID.
	//
	// Format: uuid
	VirtualMachineID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vcs virtual machines virtual machine ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVcsVirtualMachinesVirtualMachineIDParams) WithDefaults() *GetVcsVirtualMachinesVirtualMachineIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vcs virtual machines virtual machine ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVcsVirtualMachinesVirtualMachineIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vcs virtual machines virtual machine ID params
func (o *GetVcsVirtualMachinesVirtualMachineIDParams) WithTimeout(timeout time.Duration) *GetVcsVirtualMachinesVirtualMachineIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vcs virtual machines virtual machine ID params
func (o *GetVcsVirtualMachinesVirtualMachineIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vcs virtual machines virtual machine ID params
func (o *GetVcsVirtualMachinesVirtualMachineIDParams) WithContext(ctx context.Context) *GetVcsVirtualMachinesVirtualMachineIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vcs virtual machines virtual machine ID params
func (o *GetVcsVirtualMachinesVirtualMachineIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vcs virtual machines virtual machine ID params
func (o *GetVcsVirtualMachinesVirtualMachineIDParams) WithHTTPClient(client *http.Client) *GetVcsVirtualMachinesVirtualMachineIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vcs virtual machines virtual machine ID params
func (o *GetVcsVirtualMachinesVirtualMachineIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVirtualMachineID adds the virtualMachineID to the get vcs virtual machines virtual machine ID params
func (o *GetVcsVirtualMachinesVirtualMachineIDParams) WithVirtualMachineID(virtualMachineID strfmt.UUID) *GetVcsVirtualMachinesVirtualMachineIDParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the get vcs virtual machines virtual machine ID params
func (o *GetVcsVirtualMachinesVirtualMachineIDParams) SetVirtualMachineID(virtualMachineID strfmt.UUID) {
	o.VirtualMachineID = virtualMachineID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVcsVirtualMachinesVirtualMachineIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
