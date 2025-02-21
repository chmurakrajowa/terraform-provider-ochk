// Code generated by go-swagger; DO NOT EDIT.

package log_user

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

// NewPostLogUserCategoriesLogCategoryIDGeneratebydslParams creates a new PostLogUserCategoriesLogCategoryIDGeneratebydslParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostLogUserCategoriesLogCategoryIDGeneratebydslParams() *PostLogUserCategoriesLogCategoryIDGeneratebydslParams {
	return &PostLogUserCategoriesLogCategoryIDGeneratebydslParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostLogUserCategoriesLogCategoryIDGeneratebydslParamsWithTimeout creates a new PostLogUserCategoriesLogCategoryIDGeneratebydslParams object
// with the ability to set a timeout on a request.
func NewPostLogUserCategoriesLogCategoryIDGeneratebydslParamsWithTimeout(timeout time.Duration) *PostLogUserCategoriesLogCategoryIDGeneratebydslParams {
	return &PostLogUserCategoriesLogCategoryIDGeneratebydslParams{
		timeout: timeout,
	}
}

// NewPostLogUserCategoriesLogCategoryIDGeneratebydslParamsWithContext creates a new PostLogUserCategoriesLogCategoryIDGeneratebydslParams object
// with the ability to set a context for a request.
func NewPostLogUserCategoriesLogCategoryIDGeneratebydslParamsWithContext(ctx context.Context) *PostLogUserCategoriesLogCategoryIDGeneratebydslParams {
	return &PostLogUserCategoriesLogCategoryIDGeneratebydslParams{
		Context: ctx,
	}
}

// NewPostLogUserCategoriesLogCategoryIDGeneratebydslParamsWithHTTPClient creates a new PostLogUserCategoriesLogCategoryIDGeneratebydslParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostLogUserCategoriesLogCategoryIDGeneratebydslParamsWithHTTPClient(client *http.Client) *PostLogUserCategoriesLogCategoryIDGeneratebydslParams {
	return &PostLogUserCategoriesLogCategoryIDGeneratebydslParams{
		HTTPClient: client,
	}
}

/*
PostLogUserCategoriesLogCategoryIDGeneratebydslParams contains all the parameters to send to the API endpoint

	for the post log user categories log category ID generatebydsl operation.

	Typically these are written to a http.Request.
*/
type PostLogUserCategoriesLogCategoryIDGeneratebydslParams struct {

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

	// UserID.
	//
	// Format: uuid
	UserID *strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post log user categories log category ID generatebydsl params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLogUserCategoriesLogCategoryIDGeneratebydslParams) WithDefaults() *PostLogUserCategoriesLogCategoryIDGeneratebydslParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post log user categories log category ID generatebydsl params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLogUserCategoriesLogCategoryIDGeneratebydslParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post log user categories log category ID generatebydsl params
func (o *PostLogUserCategoriesLogCategoryIDGeneratebydslParams) WithTimeout(timeout time.Duration) *PostLogUserCategoriesLogCategoryIDGeneratebydslParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post log user categories log category ID generatebydsl params
func (o *PostLogUserCategoriesLogCategoryIDGeneratebydslParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post log user categories log category ID generatebydsl params
func (o *PostLogUserCategoriesLogCategoryIDGeneratebydslParams) WithContext(ctx context.Context) *PostLogUserCategoriesLogCategoryIDGeneratebydslParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post log user categories log category ID generatebydsl params
func (o *PostLogUserCategoriesLogCategoryIDGeneratebydslParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post log user categories log category ID generatebydsl params
func (o *PostLogUserCategoriesLogCategoryIDGeneratebydslParams) WithHTTPClient(client *http.Client) *PostLogUserCategoriesLogCategoryIDGeneratebydslParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post log user categories log category ID generatebydsl params
func (o *PostLogUserCategoriesLogCategoryIDGeneratebydslParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post log user categories log category ID generatebydsl params
func (o *PostLogUserCategoriesLogCategoryIDGeneratebydslParams) WithBody(body string) *PostLogUserCategoriesLogCategoryIDGeneratebydslParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post log user categories log category ID generatebydsl params
func (o *PostLogUserCategoriesLogCategoryIDGeneratebydslParams) SetBody(body string) {
	o.Body = body
}

// WithDataSize adds the dataSize to the post log user categories log category ID generatebydsl params
func (o *PostLogUserCategoriesLogCategoryIDGeneratebydslParams) WithDataSize(dataSize *int32) *PostLogUserCategoriesLogCategoryIDGeneratebydslParams {
	o.SetDataSize(dataSize)
	return o
}

// SetDataSize adds the dataSize to the post log user categories log category ID generatebydsl params
func (o *PostLogUserCategoriesLogCategoryIDGeneratebydslParams) SetDataSize(dataSize *int32) {
	o.DataSize = dataSize
}

// WithLogCategoryID adds the logCategoryID to the post log user categories log category ID generatebydsl params
func (o *PostLogUserCategoriesLogCategoryIDGeneratebydslParams) WithLogCategoryID(logCategoryID int32) *PostLogUserCategoriesLogCategoryIDGeneratebydslParams {
	o.SetLogCategoryID(logCategoryID)
	return o
}

// SetLogCategoryID adds the logCategoryId to the post log user categories log category ID generatebydsl params
func (o *PostLogUserCategoriesLogCategoryIDGeneratebydslParams) SetLogCategoryID(logCategoryID int32) {
	o.LogCategoryID = logCategoryID
}

// WithUserID adds the userID to the post log user categories log category ID generatebydsl params
func (o *PostLogUserCategoriesLogCategoryIDGeneratebydslParams) WithUserID(userID *strfmt.UUID) *PostLogUserCategoriesLogCategoryIDGeneratebydslParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the post log user categories log category ID generatebydsl params
func (o *PostLogUserCategoriesLogCategoryIDGeneratebydslParams) SetUserID(userID *strfmt.UUID) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLogUserCategoriesLogCategoryIDGeneratebydslParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.UserID != nil {

		// query param userId
		var qrUserID strfmt.UUID

		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := qrUserID.String()
		if qUserID != "" {

			if err := r.SetQueryParam("userId", qUserID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
