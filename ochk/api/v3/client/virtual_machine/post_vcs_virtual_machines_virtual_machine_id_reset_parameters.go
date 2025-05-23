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

// NewPostVcsVirtualMachinesVirtualMachineIDResetParams creates a new PostVcsVirtualMachinesVirtualMachineIDResetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostVcsVirtualMachinesVirtualMachineIDResetParams() *PostVcsVirtualMachinesVirtualMachineIDResetParams {
	return &PostVcsVirtualMachinesVirtualMachineIDResetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostVcsVirtualMachinesVirtualMachineIDResetParamsWithTimeout creates a new PostVcsVirtualMachinesVirtualMachineIDResetParams object
// with the ability to set a timeout on a request.
func NewPostVcsVirtualMachinesVirtualMachineIDResetParamsWithTimeout(timeout time.Duration) *PostVcsVirtualMachinesVirtualMachineIDResetParams {
	return &PostVcsVirtualMachinesVirtualMachineIDResetParams{
		timeout: timeout,
	}
}

// NewPostVcsVirtualMachinesVirtualMachineIDResetParamsWithContext creates a new PostVcsVirtualMachinesVirtualMachineIDResetParams object
// with the ability to set a context for a request.
func NewPostVcsVirtualMachinesVirtualMachineIDResetParamsWithContext(ctx context.Context) *PostVcsVirtualMachinesVirtualMachineIDResetParams {
	return &PostVcsVirtualMachinesVirtualMachineIDResetParams{
		Context: ctx,
	}
}

// NewPostVcsVirtualMachinesVirtualMachineIDResetParamsWithHTTPClient creates a new PostVcsVirtualMachinesVirtualMachineIDResetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostVcsVirtualMachinesVirtualMachineIDResetParamsWithHTTPClient(client *http.Client) *PostVcsVirtualMachinesVirtualMachineIDResetParams {
	return &PostVcsVirtualMachinesVirtualMachineIDResetParams{
		HTTPClient: client,
	}
}

/*
PostVcsVirtualMachinesVirtualMachineIDResetParams contains all the parameters to send to the API endpoint

	for the post vcs virtual machines virtual machine ID reset operation.

	Typically these are written to a http.Request.
*/
type PostVcsVirtualMachinesVirtualMachineIDResetParams struct {

	// VirtualMachineID.
	//
	// Format: uuid
	VirtualMachineID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post vcs virtual machines virtual machine ID reset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostVcsVirtualMachinesVirtualMachineIDResetParams) WithDefaults() *PostVcsVirtualMachinesVirtualMachineIDResetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post vcs virtual machines virtual machine ID reset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostVcsVirtualMachinesVirtualMachineIDResetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post vcs virtual machines virtual machine ID reset params
func (o *PostVcsVirtualMachinesVirtualMachineIDResetParams) WithTimeout(timeout time.Duration) *PostVcsVirtualMachinesVirtualMachineIDResetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post vcs virtual machines virtual machine ID reset params
func (o *PostVcsVirtualMachinesVirtualMachineIDResetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post vcs virtual machines virtual machine ID reset params
func (o *PostVcsVirtualMachinesVirtualMachineIDResetParams) WithContext(ctx context.Context) *PostVcsVirtualMachinesVirtualMachineIDResetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post vcs virtual machines virtual machine ID reset params
func (o *PostVcsVirtualMachinesVirtualMachineIDResetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post vcs virtual machines virtual machine ID reset params
func (o *PostVcsVirtualMachinesVirtualMachineIDResetParams) WithHTTPClient(client *http.Client) *PostVcsVirtualMachinesVirtualMachineIDResetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post vcs virtual machines virtual machine ID reset params
func (o *PostVcsVirtualMachinesVirtualMachineIDResetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVirtualMachineID adds the virtualMachineID to the post vcs virtual machines virtual machine ID reset params
func (o *PostVcsVirtualMachinesVirtualMachineIDResetParams) WithVirtualMachineID(virtualMachineID strfmt.UUID) *PostVcsVirtualMachinesVirtualMachineIDResetParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the post vcs virtual machines virtual machine ID reset params
func (o *PostVcsVirtualMachinesVirtualMachineIDResetParams) SetVirtualMachineID(virtualMachineID strfmt.UUID) {
	o.VirtualMachineID = virtualMachineID
}

// WriteToRequest writes these params to a swagger request
func (o *PostVcsVirtualMachinesVirtualMachineIDResetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
