// Code generated by go-swagger; DO NOT EDIT.

package custom_services

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

// NewCustomServiceCreateUsingPUTParams creates a new CustomServiceCreateUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomServiceCreateUsingPUTParams() *CustomServiceCreateUsingPUTParams {
	return &CustomServiceCreateUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomServiceCreateUsingPUTParamsWithTimeout creates a new CustomServiceCreateUsingPUTParams object
// with the ability to set a timeout on a request.
func NewCustomServiceCreateUsingPUTParamsWithTimeout(timeout time.Duration) *CustomServiceCreateUsingPUTParams {
	return &CustomServiceCreateUsingPUTParams{
		timeout: timeout,
	}
}

// NewCustomServiceCreateUsingPUTParamsWithContext creates a new CustomServiceCreateUsingPUTParams object
// with the ability to set a context for a request.
func NewCustomServiceCreateUsingPUTParamsWithContext(ctx context.Context) *CustomServiceCreateUsingPUTParams {
	return &CustomServiceCreateUsingPUTParams{
		Context: ctx,
	}
}

// NewCustomServiceCreateUsingPUTParamsWithHTTPClient creates a new CustomServiceCreateUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomServiceCreateUsingPUTParamsWithHTTPClient(client *http.Client) *CustomServiceCreateUsingPUTParams {
	return &CustomServiceCreateUsingPUTParams{
		HTTPClient: client,
	}
}

/* CustomServiceCreateUsingPUTParams contains all the parameters to send to the API endpoint
   for the custom service create using p u t operation.

   Typically these are written to a http.Request.
*/
type CustomServiceCreateUsingPUTParams struct {

	/* CustomServiceInstance.

	   customServiceInstance
	*/
	CustomServiceInstance *models.CustomServiceInstance

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the custom service create using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomServiceCreateUsingPUTParams) WithDefaults() *CustomServiceCreateUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the custom service create using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomServiceCreateUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the custom service create using p u t params
func (o *CustomServiceCreateUsingPUTParams) WithTimeout(timeout time.Duration) *CustomServiceCreateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom service create using p u t params
func (o *CustomServiceCreateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom service create using p u t params
func (o *CustomServiceCreateUsingPUTParams) WithContext(ctx context.Context) *CustomServiceCreateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom service create using p u t params
func (o *CustomServiceCreateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom service create using p u t params
func (o *CustomServiceCreateUsingPUTParams) WithHTTPClient(client *http.Client) *CustomServiceCreateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom service create using p u t params
func (o *CustomServiceCreateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomServiceInstance adds the customServiceInstance to the custom service create using p u t params
func (o *CustomServiceCreateUsingPUTParams) WithCustomServiceInstance(customServiceInstance *models.CustomServiceInstance) *CustomServiceCreateUsingPUTParams {
	o.SetCustomServiceInstance(customServiceInstance)
	return o
}

// SetCustomServiceInstance adds the customServiceInstance to the custom service create using p u t params
func (o *CustomServiceCreateUsingPUTParams) SetCustomServiceInstance(customServiceInstance *models.CustomServiceInstance) {
	o.CustomServiceInstance = customServiceInstance
}

// WriteToRequest writes these params to a swagger request
func (o *CustomServiceCreateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CustomServiceInstance != nil {
		if err := r.SetBodyParam(o.CustomServiceInstance); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
