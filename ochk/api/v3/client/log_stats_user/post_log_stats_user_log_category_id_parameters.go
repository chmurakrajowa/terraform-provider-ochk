// Code generated by go-swagger; DO NOT EDIT.

package log_stats_user

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

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
)

// NewPostLogStatsUserLogCategoryIDParams creates a new PostLogStatsUserLogCategoryIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostLogStatsUserLogCategoryIDParams() *PostLogStatsUserLogCategoryIDParams {
	return &PostLogStatsUserLogCategoryIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostLogStatsUserLogCategoryIDParamsWithTimeout creates a new PostLogStatsUserLogCategoryIDParams object
// with the ability to set a timeout on a request.
func NewPostLogStatsUserLogCategoryIDParamsWithTimeout(timeout time.Duration) *PostLogStatsUserLogCategoryIDParams {
	return &PostLogStatsUserLogCategoryIDParams{
		timeout: timeout,
	}
}

// NewPostLogStatsUserLogCategoryIDParamsWithContext creates a new PostLogStatsUserLogCategoryIDParams object
// with the ability to set a context for a request.
func NewPostLogStatsUserLogCategoryIDParamsWithContext(ctx context.Context) *PostLogStatsUserLogCategoryIDParams {
	return &PostLogStatsUserLogCategoryIDParams{
		Context: ctx,
	}
}

// NewPostLogStatsUserLogCategoryIDParamsWithHTTPClient creates a new PostLogStatsUserLogCategoryIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostLogStatsUserLogCategoryIDParamsWithHTTPClient(client *http.Client) *PostLogStatsUserLogCategoryIDParams {
	return &PostLogStatsUserLogCategoryIDParams{
		HTTPClient: client,
	}
}

/*
PostLogStatsUserLogCategoryIDParams contains all the parameters to send to the API endpoint

	for the post log stats user log category ID operation.

	Typically these are written to a http.Request.
*/
type PostLogStatsUserLogCategoryIDParams struct {

	// Body.
	Body *models.QueryFilter

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

// WithDefaults hydrates default values in the post log stats user log category ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLogStatsUserLogCategoryIDParams) WithDefaults() *PostLogStatsUserLogCategoryIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post log stats user log category ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLogStatsUserLogCategoryIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post log stats user log category ID params
func (o *PostLogStatsUserLogCategoryIDParams) WithTimeout(timeout time.Duration) *PostLogStatsUserLogCategoryIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post log stats user log category ID params
func (o *PostLogStatsUserLogCategoryIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post log stats user log category ID params
func (o *PostLogStatsUserLogCategoryIDParams) WithContext(ctx context.Context) *PostLogStatsUserLogCategoryIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post log stats user log category ID params
func (o *PostLogStatsUserLogCategoryIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post log stats user log category ID params
func (o *PostLogStatsUserLogCategoryIDParams) WithHTTPClient(client *http.Client) *PostLogStatsUserLogCategoryIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post log stats user log category ID params
func (o *PostLogStatsUserLogCategoryIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post log stats user log category ID params
func (o *PostLogStatsUserLogCategoryIDParams) WithBody(body *models.QueryFilter) *PostLogStatsUserLogCategoryIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post log stats user log category ID params
func (o *PostLogStatsUserLogCategoryIDParams) SetBody(body *models.QueryFilter) {
	o.Body = body
}

// WithLogCategoryID adds the logCategoryID to the post log stats user log category ID params
func (o *PostLogStatsUserLogCategoryIDParams) WithLogCategoryID(logCategoryID int32) *PostLogStatsUserLogCategoryIDParams {
	o.SetLogCategoryID(logCategoryID)
	return o
}

// SetLogCategoryID adds the logCategoryId to the post log stats user log category ID params
func (o *PostLogStatsUserLogCategoryIDParams) SetLogCategoryID(logCategoryID int32) {
	o.LogCategoryID = logCategoryID
}

// WithUserID adds the userID to the post log stats user log category ID params
func (o *PostLogStatsUserLogCategoryIDParams) WithUserID(userID *strfmt.UUID) *PostLogStatsUserLogCategoryIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the post log stats user log category ID params
func (o *PostLogStatsUserLogCategoryIDParams) SetUserID(userID *strfmt.UUID) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLogStatsUserLogCategoryIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
