// Code generated by go-swagger; DO NOT EDIT.

package log

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

// NewPostLogCategoriesLogCategoryIDGeneratebydslParams creates a new PostLogCategoriesLogCategoryIDGeneratebydslParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostLogCategoriesLogCategoryIDGeneratebydslParams() *PostLogCategoriesLogCategoryIDGeneratebydslParams {
	return &PostLogCategoriesLogCategoryIDGeneratebydslParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostLogCategoriesLogCategoryIDGeneratebydslParamsWithTimeout creates a new PostLogCategoriesLogCategoryIDGeneratebydslParams object
// with the ability to set a timeout on a request.
func NewPostLogCategoriesLogCategoryIDGeneratebydslParamsWithTimeout(timeout time.Duration) *PostLogCategoriesLogCategoryIDGeneratebydslParams {
	return &PostLogCategoriesLogCategoryIDGeneratebydslParams{
		timeout: timeout,
	}
}

// NewPostLogCategoriesLogCategoryIDGeneratebydslParamsWithContext creates a new PostLogCategoriesLogCategoryIDGeneratebydslParams object
// with the ability to set a context for a request.
func NewPostLogCategoriesLogCategoryIDGeneratebydslParamsWithContext(ctx context.Context) *PostLogCategoriesLogCategoryIDGeneratebydslParams {
	return &PostLogCategoriesLogCategoryIDGeneratebydslParams{
		Context: ctx,
	}
}

// NewPostLogCategoriesLogCategoryIDGeneratebydslParamsWithHTTPClient creates a new PostLogCategoriesLogCategoryIDGeneratebydslParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostLogCategoriesLogCategoryIDGeneratebydslParamsWithHTTPClient(client *http.Client) *PostLogCategoriesLogCategoryIDGeneratebydslParams {
	return &PostLogCategoriesLogCategoryIDGeneratebydslParams{
		HTTPClient: client,
	}
}

/*
PostLogCategoriesLogCategoryIDGeneratebydslParams contains all the parameters to send to the API endpoint

	for the post log categories log category ID generatebydsl operation.

	Typically these are written to a http.Request.
*/
type PostLogCategoriesLogCategoryIDGeneratebydslParams struct {

	// Body.
	Body string

	// DataSize.
	//
	// Format: int32
	DataSize *int32

	// LogCategoryID.
	//
	// Format: int32
	LogCategoryID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post log categories log category ID generatebydsl params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLogCategoriesLogCategoryIDGeneratebydslParams) WithDefaults() *PostLogCategoriesLogCategoryIDGeneratebydslParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post log categories log category ID generatebydsl params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLogCategoriesLogCategoryIDGeneratebydslParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post log categories log category ID generatebydsl params
func (o *PostLogCategoriesLogCategoryIDGeneratebydslParams) WithTimeout(timeout time.Duration) *PostLogCategoriesLogCategoryIDGeneratebydslParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post log categories log category ID generatebydsl params
func (o *PostLogCategoriesLogCategoryIDGeneratebydslParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post log categories log category ID generatebydsl params
func (o *PostLogCategoriesLogCategoryIDGeneratebydslParams) WithContext(ctx context.Context) *PostLogCategoriesLogCategoryIDGeneratebydslParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post log categories log category ID generatebydsl params
func (o *PostLogCategoriesLogCategoryIDGeneratebydslParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post log categories log category ID generatebydsl params
func (o *PostLogCategoriesLogCategoryIDGeneratebydslParams) WithHTTPClient(client *http.Client) *PostLogCategoriesLogCategoryIDGeneratebydslParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post log categories log category ID generatebydsl params
func (o *PostLogCategoriesLogCategoryIDGeneratebydslParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post log categories log category ID generatebydsl params
func (o *PostLogCategoriesLogCategoryIDGeneratebydslParams) WithBody(body string) *PostLogCategoriesLogCategoryIDGeneratebydslParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post log categories log category ID generatebydsl params
func (o *PostLogCategoriesLogCategoryIDGeneratebydslParams) SetBody(body string) {
	o.Body = body
}

// WithDataSize adds the dataSize to the post log categories log category ID generatebydsl params
func (o *PostLogCategoriesLogCategoryIDGeneratebydslParams) WithDataSize(dataSize *int32) *PostLogCategoriesLogCategoryIDGeneratebydslParams {
	o.SetDataSize(dataSize)
	return o
}

// SetDataSize adds the dataSize to the post log categories log category ID generatebydsl params
func (o *PostLogCategoriesLogCategoryIDGeneratebydslParams) SetDataSize(dataSize *int32) {
	o.DataSize = dataSize
}

// WithLogCategoryID adds the logCategoryID to the post log categories log category ID generatebydsl params
func (o *PostLogCategoriesLogCategoryIDGeneratebydslParams) WithLogCategoryID(logCategoryID int32) *PostLogCategoriesLogCategoryIDGeneratebydslParams {
	o.SetLogCategoryID(logCategoryID)
	return o
}

// SetLogCategoryID adds the logCategoryId to the post log categories log category ID generatebydsl params
func (o *PostLogCategoriesLogCategoryIDGeneratebydslParams) SetLogCategoryID(logCategoryID int32) {
	o.LogCategoryID = logCategoryID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLogCategoriesLogCategoryIDGeneratebydslParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if o.DataSize != nil {

		// query param dataSize
		var qrDataSize int32

		if o.DataSize != nil {
			qrDataSize = *o.DataSize
		}
		qDataSize := swag.FormatInt32(qrDataSize)
		if qDataSize != "" {

			if err := r.SetQueryParam("dataSize", qDataSize); err != nil {
				return err
			}
		}
	}

	// path param logCategoryId
	if err := r.SetPathParam("logCategoryId", swag.FormatInt32(o.LogCategoryID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
