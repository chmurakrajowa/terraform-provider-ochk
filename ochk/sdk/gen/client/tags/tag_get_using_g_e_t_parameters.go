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
	"github.com/go-openapi/swag"
)

// NewTagGetUsingGETParams creates a new TagGetUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTagGetUsingGETParams() *TagGetUsingGETParams {
	return &TagGetUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTagGetUsingGETParamsWithTimeout creates a new TagGetUsingGETParams object
// with the ability to set a timeout on a request.
func NewTagGetUsingGETParamsWithTimeout(timeout time.Duration) *TagGetUsingGETParams {
	return &TagGetUsingGETParams{
		timeout: timeout,
	}
}

// NewTagGetUsingGETParamsWithContext creates a new TagGetUsingGETParams object
// with the ability to set a context for a request.
func NewTagGetUsingGETParamsWithContext(ctx context.Context) *TagGetUsingGETParams {
	return &TagGetUsingGETParams{
		Context: ctx,
	}
}

// NewTagGetUsingGETParamsWithHTTPClient creates a new TagGetUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewTagGetUsingGETParamsWithHTTPClient(client *http.Client) *TagGetUsingGETParams {
	return &TagGetUsingGETParams{
		HTTPClient: client,
	}
}

/*
TagGetUsingGETParams contains all the parameters to send to the API endpoint

	for the tag get using g e t operation.

	Typically these are written to a http.Request.
*/
type TagGetUsingGETParams struct {

	/* TagID.

	   tagId

	   Format: int32
	*/
	TagID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tag get using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TagGetUsingGETParams) WithDefaults() *TagGetUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tag get using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TagGetUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tag get using g e t params
func (o *TagGetUsingGETParams) WithTimeout(timeout time.Duration) *TagGetUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tag get using g e t params
func (o *TagGetUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tag get using g e t params
func (o *TagGetUsingGETParams) WithContext(ctx context.Context) *TagGetUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tag get using g e t params
func (o *TagGetUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tag get using g e t params
func (o *TagGetUsingGETParams) WithHTTPClient(client *http.Client) *TagGetUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tag get using g e t params
func (o *TagGetUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTagID adds the tagID to the tag get using g e t params
func (o *TagGetUsingGETParams) WithTagID(tagID int32) *TagGetUsingGETParams {
	o.SetTagID(tagID)
	return o
}

// SetTagID adds the tagId to the tag get using g e t params
func (o *TagGetUsingGETParams) SetTagID(tagID int32) {
	o.TagID = tagID
}

// WriteToRequest writes these params to a swagger request
func (o *TagGetUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tagId
	if err := r.SetPathParam("tagId", swag.FormatInt32(o.TagID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
