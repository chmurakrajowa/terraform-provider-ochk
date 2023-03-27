// Code generated by go-swagger; DO NOT EDIT.

package log_categories_user_access

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

// NewLogCategoryGetUsingGET1Params creates a new LogCategoryGetUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLogCategoryGetUsingGET1Params() *LogCategoryGetUsingGET1Params {
	return &LogCategoryGetUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewLogCategoryGetUsingGET1ParamsWithTimeout creates a new LogCategoryGetUsingGET1Params object
// with the ability to set a timeout on a request.
func NewLogCategoryGetUsingGET1ParamsWithTimeout(timeout time.Duration) *LogCategoryGetUsingGET1Params {
	return &LogCategoryGetUsingGET1Params{
		timeout: timeout,
	}
}

// NewLogCategoryGetUsingGET1ParamsWithContext creates a new LogCategoryGetUsingGET1Params object
// with the ability to set a context for a request.
func NewLogCategoryGetUsingGET1ParamsWithContext(ctx context.Context) *LogCategoryGetUsingGET1Params {
	return &LogCategoryGetUsingGET1Params{
		Context: ctx,
	}
}

// NewLogCategoryGetUsingGET1ParamsWithHTTPClient creates a new LogCategoryGetUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewLogCategoryGetUsingGET1ParamsWithHTTPClient(client *http.Client) *LogCategoryGetUsingGET1Params {
	return &LogCategoryGetUsingGET1Params{
		HTTPClient: client,
	}
}

/*
LogCategoryGetUsingGET1Params contains all the parameters to send to the API endpoint

	for the log category get using g e t 1 operation.

	Typically these are written to a http.Request.
*/
type LogCategoryGetUsingGET1Params struct {

	/* LogCategoryID.

	   logCategoryId

	   Format: int32
	*/
	LogCategoryID int32

	/* UserID.

	   userId
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the log category get using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogCategoryGetUsingGET1Params) WithDefaults() *LogCategoryGetUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the log category get using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogCategoryGetUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the log category get using g e t 1 params
func (o *LogCategoryGetUsingGET1Params) WithTimeout(timeout time.Duration) *LogCategoryGetUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the log category get using g e t 1 params
func (o *LogCategoryGetUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the log category get using g e t 1 params
func (o *LogCategoryGetUsingGET1Params) WithContext(ctx context.Context) *LogCategoryGetUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the log category get using g e t 1 params
func (o *LogCategoryGetUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the log category get using g e t 1 params
func (o *LogCategoryGetUsingGET1Params) WithHTTPClient(client *http.Client) *LogCategoryGetUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the log category get using g e t 1 params
func (o *LogCategoryGetUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLogCategoryID adds the logCategoryID to the log category get using g e t 1 params
func (o *LogCategoryGetUsingGET1Params) WithLogCategoryID(logCategoryID int32) *LogCategoryGetUsingGET1Params {
	o.SetLogCategoryID(logCategoryID)
	return o
}

// SetLogCategoryID adds the logCategoryId to the log category get using g e t 1 params
func (o *LogCategoryGetUsingGET1Params) SetLogCategoryID(logCategoryID int32) {
	o.LogCategoryID = logCategoryID
}

// WithUserID adds the userID to the log category get using g e t 1 params
func (o *LogCategoryGetUsingGET1Params) WithUserID(userID string) *LogCategoryGetUsingGET1Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the log category get using g e t 1 params
func (o *LogCategoryGetUsingGET1Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *LogCategoryGetUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param logCategoryId
	if err := r.SetPathParam("logCategoryId", swag.FormatInt32(o.LogCategoryID)); err != nil {
		return err
	}

	// query param userId
	qrUserID := o.UserID
	qUserID := qrUserID
	if qUserID != "" {

		if err := r.SetQueryParam("userId", qUserID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}