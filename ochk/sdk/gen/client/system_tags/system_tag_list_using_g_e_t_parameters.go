// Code generated by go-swagger; DO NOT EDIT.

package system_tags

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

// NewSystemTagListUsingGETParams creates a new SystemTagListUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSystemTagListUsingGETParams() *SystemTagListUsingGETParams {
	return &SystemTagListUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSystemTagListUsingGETParamsWithTimeout creates a new SystemTagListUsingGETParams object
// with the ability to set a timeout on a request.
func NewSystemTagListUsingGETParamsWithTimeout(timeout time.Duration) *SystemTagListUsingGETParams {
	return &SystemTagListUsingGETParams{
		timeout: timeout,
	}
}

// NewSystemTagListUsingGETParamsWithContext creates a new SystemTagListUsingGETParams object
// with the ability to set a context for a request.
func NewSystemTagListUsingGETParamsWithContext(ctx context.Context) *SystemTagListUsingGETParams {
	return &SystemTagListUsingGETParams{
		Context: ctx,
	}
}

// NewSystemTagListUsingGETParamsWithHTTPClient creates a new SystemTagListUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewSystemTagListUsingGETParamsWithHTTPClient(client *http.Client) *SystemTagListUsingGETParams {
	return &SystemTagListUsingGETParams{
		HTTPClient: client,
	}
}

/* SystemTagListUsingGETParams contains all the parameters to send to the API endpoint
   for the system tag list using g e t operation.

   Typically these are written to a http.Request.
*/
type SystemTagListUsingGETParams struct {

	/* TagValue.

	   tagValue
	*/
	TagValue *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the system tag list using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemTagListUsingGETParams) WithDefaults() *SystemTagListUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the system tag list using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemTagListUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the system tag list using g e t params
func (o *SystemTagListUsingGETParams) WithTimeout(timeout time.Duration) *SystemTagListUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system tag list using g e t params
func (o *SystemTagListUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system tag list using g e t params
func (o *SystemTagListUsingGETParams) WithContext(ctx context.Context) *SystemTagListUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system tag list using g e t params
func (o *SystemTagListUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system tag list using g e t params
func (o *SystemTagListUsingGETParams) WithHTTPClient(client *http.Client) *SystemTagListUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system tag list using g e t params
func (o *SystemTagListUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTagValue adds the tagValue to the system tag list using g e t params
func (o *SystemTagListUsingGETParams) WithTagValue(tagValue *string) *SystemTagListUsingGETParams {
	o.SetTagValue(tagValue)
	return o
}

// SetTagValue adds the tagValue to the system tag list using g e t params
func (o *SystemTagListUsingGETParams) SetTagValue(tagValue *string) {
	o.TagValue = tagValue
}

// WriteToRequest writes these params to a swagger request
func (o *SystemTagListUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.TagValue != nil {

		// query param tagValue
		var qrTagValue string

		if o.TagValue != nil {
			qrTagValue = *o.TagValue
		}
		qTagValue := qrTagValue
		if qTagValue != "" {

			if err := r.SetQueryParam("tagValue", qTagValue); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
