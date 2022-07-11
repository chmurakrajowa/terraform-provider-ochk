// Code generated by go-swagger; DO NOT EDIT.

package local_groups

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

// NewLocalGroupCreateUsingPUTParams creates a new LocalGroupCreateUsingPUTParams object
// with the default values initialized.
func NewLocalGroupCreateUsingPUTParams() *LocalGroupCreateUsingPUTParams {
	var ()
	return &LocalGroupCreateUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLocalGroupCreateUsingPUTParamsWithTimeout creates a new LocalGroupCreateUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLocalGroupCreateUsingPUTParamsWithTimeout(timeout time.Duration) *LocalGroupCreateUsingPUTParams {
	var ()
	return &LocalGroupCreateUsingPUTParams{

		timeout: timeout,
	}
}

// NewLocalGroupCreateUsingPUTParamsWithContext creates a new LocalGroupCreateUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewLocalGroupCreateUsingPUTParamsWithContext(ctx context.Context) *LocalGroupCreateUsingPUTParams {
	var ()
	return &LocalGroupCreateUsingPUTParams{

		Context: ctx,
	}
}

// NewLocalGroupCreateUsingPUTParamsWithHTTPClient creates a new LocalGroupCreateUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLocalGroupCreateUsingPUTParamsWithHTTPClient(client *http.Client) *LocalGroupCreateUsingPUTParams {
	var ()
	return &LocalGroupCreateUsingPUTParams{
		HTTPClient: client,
	}
}

/*LocalGroupCreateUsingPUTParams contains all the parameters to send to the API endpoint
for the local group create using p u t operation typically these are written to a http.Request
*/
type LocalGroupCreateUsingPUTParams struct {

	/*LocalGroup
	  localGroup

	*/
	LocalGroup *models.LocalGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the local group create using p u t params
func (o *LocalGroupCreateUsingPUTParams) WithTimeout(timeout time.Duration) *LocalGroupCreateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the local group create using p u t params
func (o *LocalGroupCreateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the local group create using p u t params
func (o *LocalGroupCreateUsingPUTParams) WithContext(ctx context.Context) *LocalGroupCreateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the local group create using p u t params
func (o *LocalGroupCreateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the local group create using p u t params
func (o *LocalGroupCreateUsingPUTParams) WithHTTPClient(client *http.Client) *LocalGroupCreateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the local group create using p u t params
func (o *LocalGroupCreateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocalGroup adds the localGroup to the local group create using p u t params
func (o *LocalGroupCreateUsingPUTParams) WithLocalGroup(localGroup *models.LocalGroup) *LocalGroupCreateUsingPUTParams {
	o.SetLocalGroup(localGroup)
	return o
}

// SetLocalGroup adds the localGroup to the local group create using p u t params
func (o *LocalGroupCreateUsingPUTParams) SetLocalGroup(localGroup *models.LocalGroup) {
	o.LocalGroup = localGroup
}

// WriteToRequest writes these params to a swagger request
func (o *LocalGroupCreateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LocalGroup != nil {
		if err := r.SetBodyParam(o.LocalGroup); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}