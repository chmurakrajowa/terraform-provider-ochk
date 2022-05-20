// Code generated by go-swagger; DO NOT EDIT.

package active_directory_users

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

// NewGetADUserUsingGETParams creates a new GetADUserUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetADUserUsingGETParams() *GetADUserUsingGETParams {
	return &GetADUserUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetADUserUsingGETParamsWithTimeout creates a new GetADUserUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetADUserUsingGETParamsWithTimeout(timeout time.Duration) *GetADUserUsingGETParams {
	return &GetADUserUsingGETParams{
		timeout: timeout,
	}
}

// NewGetADUserUsingGETParamsWithContext creates a new GetADUserUsingGETParams object
// with the ability to set a context for a request.
func NewGetADUserUsingGETParamsWithContext(ctx context.Context) *GetADUserUsingGETParams {
	return &GetADUserUsingGETParams{
		Context: ctx,
	}
}

// NewGetADUserUsingGETParamsWithHTTPClient creates a new GetADUserUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetADUserUsingGETParamsWithHTTPClient(client *http.Client) *GetADUserUsingGETParams {
	return &GetADUserUsingGETParams{
		HTTPClient: client,
	}
}

/* GetADUserUsingGETParams contains all the parameters to send to the API endpoint
   for the get a d user using g e t operation.

   Typically these are written to a http.Request.
*/
type GetADUserUsingGETParams struct {

	/* SamAccountName.

	   samAccountName
	*/
	SamAccountName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get a d user using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetADUserUsingGETParams) WithDefaults() *GetADUserUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get a d user using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetADUserUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get a d user using g e t params
func (o *GetADUserUsingGETParams) WithTimeout(timeout time.Duration) *GetADUserUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get a d user using g e t params
func (o *GetADUserUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get a d user using g e t params
func (o *GetADUserUsingGETParams) WithContext(ctx context.Context) *GetADUserUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get a d user using g e t params
func (o *GetADUserUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get a d user using g e t params
func (o *GetADUserUsingGETParams) WithHTTPClient(client *http.Client) *GetADUserUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get a d user using g e t params
func (o *GetADUserUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSamAccountName adds the samAccountName to the get a d user using g e t params
func (o *GetADUserUsingGETParams) WithSamAccountName(samAccountName string) *GetADUserUsingGETParams {
	o.SetSamAccountName(samAccountName)
	return o
}

// SetSamAccountName adds the samAccountName to the get a d user using g e t params
func (o *GetADUserUsingGETParams) SetSamAccountName(samAccountName string) {
	o.SamAccountName = samAccountName
}

// WriteToRequest writes these params to a swagger request
func (o *GetADUserUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param samAccountName
	if err := r.SetPathParam("samAccountName", o.SamAccountName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
