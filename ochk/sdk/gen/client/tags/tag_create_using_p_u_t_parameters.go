// Code generated by go-swagger; DO NOT EDIT.

package tags

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

// NewTagCreateUsingPUTParams creates a new TagCreateUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTagCreateUsingPUTParams() *TagCreateUsingPUTParams {
	return &TagCreateUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTagCreateUsingPUTParamsWithTimeout creates a new TagCreateUsingPUTParams object
// with the ability to set a timeout on a request.
func NewTagCreateUsingPUTParamsWithTimeout(timeout time.Duration) *TagCreateUsingPUTParams {
	return &TagCreateUsingPUTParams{
		timeout: timeout,
	}
}

// NewTagCreateUsingPUTParamsWithContext creates a new TagCreateUsingPUTParams object
// with the ability to set a context for a request.
func NewTagCreateUsingPUTParamsWithContext(ctx context.Context) *TagCreateUsingPUTParams {
	return &TagCreateUsingPUTParams{
		Context: ctx,
	}
}

// NewTagCreateUsingPUTParamsWithHTTPClient creates a new TagCreateUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewTagCreateUsingPUTParamsWithHTTPClient(client *http.Client) *TagCreateUsingPUTParams {
	return &TagCreateUsingPUTParams{
		HTTPClient: client,
	}
}

/*
TagCreateUsingPUTParams contains all the parameters to send to the API endpoint

	for the tag create using p u t operation.

	Typically these are written to a http.Request.
*/
type TagCreateUsingPUTParams struct {

	/* Tag.

	   tag
	*/
	Tag *models.Tag

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tag create using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TagCreateUsingPUTParams) WithDefaults() *TagCreateUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tag create using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TagCreateUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tag create using p u t params
func (o *TagCreateUsingPUTParams) WithTimeout(timeout time.Duration) *TagCreateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tag create using p u t params
func (o *TagCreateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tag create using p u t params
func (o *TagCreateUsingPUTParams) WithContext(ctx context.Context) *TagCreateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tag create using p u t params
func (o *TagCreateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tag create using p u t params
func (o *TagCreateUsingPUTParams) WithHTTPClient(client *http.Client) *TagCreateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tag create using p u t params
func (o *TagCreateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTag adds the tag to the tag create using p u t params
func (o *TagCreateUsingPUTParams) WithTag(tag *models.Tag) *TagCreateUsingPUTParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the tag create using p u t params
func (o *TagCreateUsingPUTParams) SetTag(tag *models.Tag) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *TagCreateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Tag != nil {
		if err := r.SetBodyParam(o.Tag); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
