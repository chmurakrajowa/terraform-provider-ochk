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

// NewVcsVirtualMachineCreateUsingPUTParams creates a new VcsVirtualMachineCreateUsingPUTParams object
// with the default values initialized.
func NewVcsVirtualMachineCreateUsingPUTParams() *VcsVirtualMachineCreateUsingPUTParams {
	var ()
	return &VcsVirtualMachineCreateUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVcsVirtualMachineCreateUsingPUTParamsWithTimeout creates a new VcsVirtualMachineCreateUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVcsVirtualMachineCreateUsingPUTParamsWithTimeout(timeout time.Duration) *VcsVirtualMachineCreateUsingPUTParams {
	var ()
	return &VcsVirtualMachineCreateUsingPUTParams{

		timeout: timeout,
	}
}

// NewVcsVirtualMachineCreateUsingPUTParamsWithContext creates a new VcsVirtualMachineCreateUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewVcsVirtualMachineCreateUsingPUTParamsWithContext(ctx context.Context) *VcsVirtualMachineCreateUsingPUTParams {
	var ()
	return &VcsVirtualMachineCreateUsingPUTParams{

		Context: ctx,
	}
}

// NewVcsVirtualMachineCreateUsingPUTParamsWithHTTPClient creates a new VcsVirtualMachineCreateUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVcsVirtualMachineCreateUsingPUTParamsWithHTTPClient(client *http.Client) *VcsVirtualMachineCreateUsingPUTParams {
	var ()
	return &VcsVirtualMachineCreateUsingPUTParams{
		HTTPClient: client,
	}
}

/*VcsVirtualMachineCreateUsingPUTParams contains all the parameters to send to the API endpoint
for the vcs virtual machine create using p u t operation typically these are written to a http.Request
*/
type VcsVirtualMachineCreateUsingPUTParams struct {

	/*VirtualMachine
	  virtualMachine

	*/
	VirtualMachine *models.VcsVirtualMachineInstance

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the vcs virtual machine create using p u t params
func (o *VcsVirtualMachineCreateUsingPUTParams) WithTimeout(timeout time.Duration) *VcsVirtualMachineCreateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vcs virtual machine create using p u t params
func (o *VcsVirtualMachineCreateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vcs virtual machine create using p u t params
func (o *VcsVirtualMachineCreateUsingPUTParams) WithContext(ctx context.Context) *VcsVirtualMachineCreateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vcs virtual machine create using p u t params
func (o *VcsVirtualMachineCreateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vcs virtual machine create using p u t params
func (o *VcsVirtualMachineCreateUsingPUTParams) WithHTTPClient(client *http.Client) *VcsVirtualMachineCreateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vcs virtual machine create using p u t params
func (o *VcsVirtualMachineCreateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVirtualMachine adds the virtualMachine to the vcs virtual machine create using p u t params
func (o *VcsVirtualMachineCreateUsingPUTParams) WithVirtualMachine(virtualMachine *models.VcsVirtualMachineInstance) *VcsVirtualMachineCreateUsingPUTParams {
	o.SetVirtualMachine(virtualMachine)
	return o
}

// SetVirtualMachine adds the virtualMachine to the vcs virtual machine create using p u t params
func (o *VcsVirtualMachineCreateUsingPUTParams) SetVirtualMachine(virtualMachine *models.VcsVirtualMachineInstance) {
	o.VirtualMachine = virtualMachine
}

// WriteToRequest writes these params to a swagger request
func (o *VcsVirtualMachineCreateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.VirtualMachine != nil {
		if err := r.SetBodyParam(o.VirtualMachine); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
