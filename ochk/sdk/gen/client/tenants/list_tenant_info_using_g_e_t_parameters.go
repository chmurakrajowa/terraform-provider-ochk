// Code generated by go-swagger; DO NOT EDIT.

package tenants

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

// NewListTenantInfoUsingGETParams creates a new ListTenantInfoUsingGETParams object
// with the default values initialized.
func NewListTenantInfoUsingGETParams() *ListTenantInfoUsingGETParams {
	var ()
	return &ListTenantInfoUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListTenantInfoUsingGETParamsWithTimeout creates a new ListTenantInfoUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListTenantInfoUsingGETParamsWithTimeout(timeout time.Duration) *ListTenantInfoUsingGETParams {
	var ()
	return &ListTenantInfoUsingGETParams{

		timeout: timeout,
	}
}

// NewListTenantInfoUsingGETParamsWithContext creates a new ListTenantInfoUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewListTenantInfoUsingGETParamsWithContext(ctx context.Context) *ListTenantInfoUsingGETParams {
	var ()
	return &ListTenantInfoUsingGETParams{

		Context: ctx,
	}
}

// NewListTenantInfoUsingGETParamsWithHTTPClient creates a new ListTenantInfoUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListTenantInfoUsingGETParamsWithHTTPClient(client *http.Client) *ListTenantInfoUsingGETParams {
	var ()
	return &ListTenantInfoUsingGETParams{
		HTTPClient: client,
	}
}

/*ListTenantInfoUsingGETParams contains all the parameters to send to the API endpoint
for the list tenant info using g e t operation typically these are written to a http.Request
*/
type ListTenantInfoUsingGETParams struct {

	/*TenantName
	  tenantName

	*/
	TenantName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list tenant info using g e t params
func (o *ListTenantInfoUsingGETParams) WithTimeout(timeout time.Duration) *ListTenantInfoUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list tenant info using g e t params
func (o *ListTenantInfoUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list tenant info using g e t params
func (o *ListTenantInfoUsingGETParams) WithContext(ctx context.Context) *ListTenantInfoUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list tenant info using g e t params
func (o *ListTenantInfoUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list tenant info using g e t params
func (o *ListTenantInfoUsingGETParams) WithHTTPClient(client *http.Client) *ListTenantInfoUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list tenant info using g e t params
func (o *ListTenantInfoUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenantName adds the tenantName to the list tenant info using g e t params
func (o *ListTenantInfoUsingGETParams) WithTenantName(tenantName *string) *ListTenantInfoUsingGETParams {
	o.SetTenantName(tenantName)
	return o
}

// SetTenantName adds the tenantName to the list tenant info using g e t params
func (o *ListTenantInfoUsingGETParams) SetTenantName(tenantName *string) {
	o.TenantName = tenantName
}

// WriteToRequest writes these params to a swagger request
func (o *ListTenantInfoUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.TenantName != nil {

		// query param tenantName
		var qrTenantName string
		if o.TenantName != nil {
			qrTenantName = *o.TenantName
		}
		qTenantName := qrTenantName
		if qTenantName != "" {
			if err := r.SetQueryParam("tenantName", qTenantName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
