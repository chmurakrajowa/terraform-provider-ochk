// Code generated by go-swagger; DO NOT EDIT.

package resource_pools

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

// NewVcsResourcePoolsListUsingGETParams creates a new VcsResourcePoolsListUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVcsResourcePoolsListUsingGETParams() *VcsResourcePoolsListUsingGETParams {
	return &VcsResourcePoolsListUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVcsResourcePoolsListUsingGETParamsWithTimeout creates a new VcsResourcePoolsListUsingGETParams object
// with the ability to set a timeout on a request.
func NewVcsResourcePoolsListUsingGETParamsWithTimeout(timeout time.Duration) *VcsResourcePoolsListUsingGETParams {
	return &VcsResourcePoolsListUsingGETParams{
		timeout: timeout,
	}
}

// NewVcsResourcePoolsListUsingGETParamsWithContext creates a new VcsResourcePoolsListUsingGETParams object
// with the ability to set a context for a request.
func NewVcsResourcePoolsListUsingGETParamsWithContext(ctx context.Context) *VcsResourcePoolsListUsingGETParams {
	return &VcsResourcePoolsListUsingGETParams{
		Context: ctx,
	}
}

// NewVcsResourcePoolsListUsingGETParamsWithHTTPClient creates a new VcsResourcePoolsListUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewVcsResourcePoolsListUsingGETParamsWithHTTPClient(client *http.Client) *VcsResourcePoolsListUsingGETParams {
	return &VcsResourcePoolsListUsingGETParams{
		HTTPClient: client,
	}
}

/* VcsResourcePoolsListUsingGETParams contains all the parameters to send to the API endpoint
   for the vcs resource pools list using g e t operation.

   Typically these are written to a http.Request.
*/
type VcsResourcePoolsListUsingGETParams struct {

	/* ResourcePoolName.

	   resourcePoolName
	*/
	ResourcePoolName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vcs resource pools list using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VcsResourcePoolsListUsingGETParams) WithDefaults() *VcsResourcePoolsListUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vcs resource pools list using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VcsResourcePoolsListUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the vcs resource pools list using g e t params
func (o *VcsResourcePoolsListUsingGETParams) WithTimeout(timeout time.Duration) *VcsResourcePoolsListUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vcs resource pools list using g e t params
func (o *VcsResourcePoolsListUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vcs resource pools list using g e t params
func (o *VcsResourcePoolsListUsingGETParams) WithContext(ctx context.Context) *VcsResourcePoolsListUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vcs resource pools list using g e t params
func (o *VcsResourcePoolsListUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vcs resource pools list using g e t params
func (o *VcsResourcePoolsListUsingGETParams) WithHTTPClient(client *http.Client) *VcsResourcePoolsListUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vcs resource pools list using g e t params
func (o *VcsResourcePoolsListUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourcePoolName adds the resourcePoolName to the vcs resource pools list using g e t params
func (o *VcsResourcePoolsListUsingGETParams) WithResourcePoolName(resourcePoolName *string) *VcsResourcePoolsListUsingGETParams {
	o.SetResourcePoolName(resourcePoolName)
	return o
}

// SetResourcePoolName adds the resourcePoolName to the vcs resource pools list using g e t params
func (o *VcsResourcePoolsListUsingGETParams) SetResourcePoolName(resourcePoolName *string) {
	o.ResourcePoolName = resourcePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *VcsResourcePoolsListUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ResourcePoolName != nil {

		// query param resourcePoolName
		var qrResourcePoolName string

		if o.ResourcePoolName != nil {
			qrResourcePoolName = *o.ResourcePoolName
		}
		qResourcePoolName := qrResourcePoolName
		if qResourcePoolName != "" {

			if err := r.SetQueryParam("resourcePoolName", qResourcePoolName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
