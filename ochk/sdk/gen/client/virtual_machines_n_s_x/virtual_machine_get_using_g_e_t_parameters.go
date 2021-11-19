// Code generated by go-swagger; DO NOT EDIT.

package virtual_machines_n_s_x

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

// NewVirtualMachineGetUsingGETParams creates a new VirtualMachineGetUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualMachineGetUsingGETParams() *VirtualMachineGetUsingGETParams {
	return &VirtualMachineGetUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualMachineGetUsingGETParamsWithTimeout creates a new VirtualMachineGetUsingGETParams object
// with the ability to set a timeout on a request.
func NewVirtualMachineGetUsingGETParamsWithTimeout(timeout time.Duration) *VirtualMachineGetUsingGETParams {
	return &VirtualMachineGetUsingGETParams{
		timeout: timeout,
	}
}

// NewVirtualMachineGetUsingGETParamsWithContext creates a new VirtualMachineGetUsingGETParams object
// with the ability to set a context for a request.
func NewVirtualMachineGetUsingGETParamsWithContext(ctx context.Context) *VirtualMachineGetUsingGETParams {
	return &VirtualMachineGetUsingGETParams{
		Context: ctx,
	}
}

// NewVirtualMachineGetUsingGETParamsWithHTTPClient creates a new VirtualMachineGetUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualMachineGetUsingGETParamsWithHTTPClient(client *http.Client) *VirtualMachineGetUsingGETParams {
	return &VirtualMachineGetUsingGETParams{
		HTTPClient: client,
	}
}

/* VirtualMachineGetUsingGETParams contains all the parameters to send to the API endpoint
   for the virtual machine get using g e t operation.

   Typically these are written to a http.Request.
*/
type VirtualMachineGetUsingGETParams struct {

	/* VirtualMachineID.

	   virtualMachineId
	*/
	VirtualMachineID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtual machine get using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualMachineGetUsingGETParams) WithDefaults() *VirtualMachineGetUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtual machine get using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualMachineGetUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtual machine get using g e t params
func (o *VirtualMachineGetUsingGETParams) WithTimeout(timeout time.Duration) *VirtualMachineGetUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtual machine get using g e t params
func (o *VirtualMachineGetUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtual machine get using g e t params
func (o *VirtualMachineGetUsingGETParams) WithContext(ctx context.Context) *VirtualMachineGetUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtual machine get using g e t params
func (o *VirtualMachineGetUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtual machine get using g e t params
func (o *VirtualMachineGetUsingGETParams) WithHTTPClient(client *http.Client) *VirtualMachineGetUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtual machine get using g e t params
func (o *VirtualMachineGetUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVirtualMachineID adds the virtualMachineID to the virtual machine get using g e t params
func (o *VirtualMachineGetUsingGETParams) WithVirtualMachineID(virtualMachineID string) *VirtualMachineGetUsingGETParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the virtual machine get using g e t params
func (o *VirtualMachineGetUsingGETParams) SetVirtualMachineID(virtualMachineID string) {
	o.VirtualMachineID = virtualMachineID
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualMachineGetUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
