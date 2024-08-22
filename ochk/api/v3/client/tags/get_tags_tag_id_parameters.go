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

// NewGetTagsTagIDParams creates a new GetTagsTagIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTagsTagIDParams() *GetTagsTagIDParams {
	return &GetTagsTagIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTagsTagIDParamsWithTimeout creates a new GetTagsTagIDParams object
// with the ability to set a timeout on a request.
func NewGetTagsTagIDParamsWithTimeout(timeout time.Duration) *GetTagsTagIDParams {
	return &GetTagsTagIDParams{
		timeout: timeout,
	}
}

// NewGetTagsTagIDParamsWithContext creates a new GetTagsTagIDParams object
// with the ability to set a context for a request.
func NewGetTagsTagIDParamsWithContext(ctx context.Context) *GetTagsTagIDParams {
	return &GetTagsTagIDParams{
		Context: ctx,
	}
}

// NewGetTagsTagIDParamsWithHTTPClient creates a new GetTagsTagIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTagsTagIDParamsWithHTTPClient(client *http.Client) *GetTagsTagIDParams {
	return &GetTagsTagIDParams{
		HTTPClient: client,
	}
}

/*
GetTagsTagIDParams contains all the parameters to send to the API endpoint

	for the get tags tag ID operation.

	Typically these are written to a http.Request.
*/
type GetTagsTagIDParams struct {

	// TagID.
	//
	// Format: int32
	TagID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get tags tag ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTagsTagIDParams) WithDefaults() *GetTagsTagIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get tags tag ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTagsTagIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get tags tag ID params
func (o *GetTagsTagIDParams) WithTimeout(timeout time.Duration) *GetTagsTagIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tags tag ID params
func (o *GetTagsTagIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tags tag ID params
func (o *GetTagsTagIDParams) WithContext(ctx context.Context) *GetTagsTagIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tags tag ID params
func (o *GetTagsTagIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tags tag ID params
func (o *GetTagsTagIDParams) WithHTTPClient(client *http.Client) *GetTagsTagIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tags tag ID params
func (o *GetTagsTagIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTagID adds the tagID to the get tags tag ID params
func (o *GetTagsTagIDParams) WithTagID(tagID int32) *GetTagsTagIDParams {
	o.SetTagID(tagID)
	return o
}

// SetTagID adds the tagId to the get tags tag ID params
func (o *GetTagsTagIDParams) SetTagID(tagID int32) {
	o.TagID = tagID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTagsTagIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
