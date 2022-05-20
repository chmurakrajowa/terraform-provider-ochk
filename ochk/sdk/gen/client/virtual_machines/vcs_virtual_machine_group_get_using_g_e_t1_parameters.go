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

// NewVcsVirtualMachineGroupGetUsingGET1Params creates a new VcsVirtualMachineGroupGetUsingGET1Params object
// with the default values initialized.
func NewVcsVirtualMachineGroupGetUsingGET1Params() *VcsVirtualMachineGroupGetUsingGET1Params {
	var ()
	return &VcsVirtualMachineGroupGetUsingGET1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewVcsVirtualMachineGroupGetUsingGET1ParamsWithTimeout creates a new VcsVirtualMachineGroupGetUsingGET1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewVcsVirtualMachineGroupGetUsingGET1ParamsWithTimeout(timeout time.Duration) *VcsVirtualMachineGroupGetUsingGET1Params {
	var ()
	return &VcsVirtualMachineGroupGetUsingGET1Params{

		timeout: timeout,
	}
}

// NewVcsVirtualMachineGroupGetUsingGET1ParamsWithContext creates a new VcsVirtualMachineGroupGetUsingGET1Params object
// with the default values initialized, and the ability to set a context for a request
func NewVcsVirtualMachineGroupGetUsingGET1ParamsWithContext(ctx context.Context) *VcsVirtualMachineGroupGetUsingGET1Params {
	var ()
	return &VcsVirtualMachineGroupGetUsingGET1Params{

		Context: ctx,
	}
}

// NewVcsVirtualMachineGroupGetUsingGET1ParamsWithHTTPClient creates a new VcsVirtualMachineGroupGetUsingGET1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVcsVirtualMachineGroupGetUsingGET1ParamsWithHTTPClient(client *http.Client) *VcsVirtualMachineGroupGetUsingGET1Params {
	var ()
	return &VcsVirtualMachineGroupGetUsingGET1Params{
		HTTPClient: client,
	}
}

/*VcsVirtualMachineGroupGetUsingGET1Params contains all the parameters to send to the API endpoint
for the vcs virtual machine group get using g e t 1 operation typically these are written to a http.Request
*/
type VcsVirtualMachineGroupGetUsingGET1Params struct {

	/*VirtualMachineID
	  virtualMachineId

	*/
	VirtualMachineID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the vcs virtual machine group get using g e t 1 params
func (o *VcsVirtualMachineGroupGetUsingGET1Params) WithTimeout(timeout time.Duration) *VcsVirtualMachineGroupGetUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vcs virtual machine group get using g e t 1 params
func (o *VcsVirtualMachineGroupGetUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vcs virtual machine group get using g e t 1 params
func (o *VcsVirtualMachineGroupGetUsingGET1Params) WithContext(ctx context.Context) *VcsVirtualMachineGroupGetUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vcs virtual machine group get using g e t 1 params
func (o *VcsVirtualMachineGroupGetUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vcs virtual machine group get using g e t 1 params
func (o *VcsVirtualMachineGroupGetUsingGET1Params) WithHTTPClient(client *http.Client) *VcsVirtualMachineGroupGetUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vcs virtual machine group get using g e t 1 params
func (o *VcsVirtualMachineGroupGetUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVirtualMachineID adds the virtualMachineID to the vcs virtual machine group get using g e t 1 params
func (o *VcsVirtualMachineGroupGetUsingGET1Params) WithVirtualMachineID(virtualMachineID string) *VcsVirtualMachineGroupGetUsingGET1Params {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the vcs virtual machine group get using g e t 1 params
func (o *VcsVirtualMachineGroupGetUsingGET1Params) SetVirtualMachineID(virtualMachineID string) {
	o.VirtualMachineID = virtualMachineID
}

// WriteToRequest writes these params to a swagger request
func (o *VcsVirtualMachineGroupGetUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
