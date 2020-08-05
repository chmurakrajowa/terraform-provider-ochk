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

// NewVcsVirtualMachineListUsingGET1Params creates a new VcsVirtualMachineListUsingGET1Params object
// with the default values initialized.
func NewVcsVirtualMachineListUsingGET1Params() *VcsVirtualMachineListUsingGET1Params {

	return &VcsVirtualMachineListUsingGET1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewVcsVirtualMachineListUsingGET1ParamsWithTimeout creates a new VcsVirtualMachineListUsingGET1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewVcsVirtualMachineListUsingGET1ParamsWithTimeout(timeout time.Duration) *VcsVirtualMachineListUsingGET1Params {

	return &VcsVirtualMachineListUsingGET1Params{

		timeout: timeout,
	}
}

// NewVcsVirtualMachineListUsingGET1ParamsWithContext creates a new VcsVirtualMachineListUsingGET1Params object
// with the default values initialized, and the ability to set a context for a request
func NewVcsVirtualMachineListUsingGET1ParamsWithContext(ctx context.Context) *VcsVirtualMachineListUsingGET1Params {

	return &VcsVirtualMachineListUsingGET1Params{

		Context: ctx,
	}
}

// NewVcsVirtualMachineListUsingGET1ParamsWithHTTPClient creates a new VcsVirtualMachineListUsingGET1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVcsVirtualMachineListUsingGET1ParamsWithHTTPClient(client *http.Client) *VcsVirtualMachineListUsingGET1Params {

	return &VcsVirtualMachineListUsingGET1Params{
		HTTPClient: client,
	}
}

/*VcsVirtualMachineListUsingGET1Params contains all the parameters to send to the API endpoint
for the vcs virtual machine list using g e t 1 operation typically these are written to a http.Request
*/
type VcsVirtualMachineListUsingGET1Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the vcs virtual machine list using g e t 1 params
func (o *VcsVirtualMachineListUsingGET1Params) WithTimeout(timeout time.Duration) *VcsVirtualMachineListUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vcs virtual machine list using g e t 1 params
func (o *VcsVirtualMachineListUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vcs virtual machine list using g e t 1 params
func (o *VcsVirtualMachineListUsingGET1Params) WithContext(ctx context.Context) *VcsVirtualMachineListUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vcs virtual machine list using g e t 1 params
func (o *VcsVirtualMachineListUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vcs virtual machine list using g e t 1 params
func (o *VcsVirtualMachineListUsingGET1Params) WithHTTPClient(client *http.Client) *VcsVirtualMachineListUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vcs virtual machine list using g e t 1 params
func (o *VcsVirtualMachineListUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *VcsVirtualMachineListUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
