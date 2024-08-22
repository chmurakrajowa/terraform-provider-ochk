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

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
)

// NewPostLogUserCategoriesLogCategoryIDGenerateParams creates a new PostLogUserCategoriesLogCategoryIDGenerateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostLogUserCategoriesLogCategoryIDGenerateParams() *PostLogUserCategoriesLogCategoryIDGenerateParams {
	return &PostLogUserCategoriesLogCategoryIDGenerateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostLogUserCategoriesLogCategoryIDGenerateParamsWithTimeout creates a new PostLogUserCategoriesLogCategoryIDGenerateParams object
// with the ability to set a timeout on a request.
func NewPostLogUserCategoriesLogCategoryIDGenerateParamsWithTimeout(timeout time.Duration) *PostLogUserCategoriesLogCategoryIDGenerateParams {
	return &PostLogUserCategoriesLogCategoryIDGenerateParams{
		timeout: timeout,
	}
}

// NewPostLogUserCategoriesLogCategoryIDGenerateParamsWithContext creates a new PostLogUserCategoriesLogCategoryIDGenerateParams object
// with the ability to set a context for a request.
func NewPostLogUserCategoriesLogCategoryIDGenerateParamsWithContext(ctx context.Context) *PostLogUserCategoriesLogCategoryIDGenerateParams {
	return &PostLogUserCategoriesLogCategoryIDGenerateParams{
		Context: ctx,
	}
}

// NewPostLogUserCategoriesLogCategoryIDGenerateParamsWithHTTPClient creates a new PostLogUserCategoriesLogCategoryIDGenerateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostLogUserCategoriesLogCategoryIDGenerateParamsWithHTTPClient(client *http.Client) *PostLogUserCategoriesLogCategoryIDGenerateParams {
	return &PostLogUserCategoriesLogCategoryIDGenerateParams{
		HTTPClient: client,
	}
}

/*
PostLogUserCategoriesLogCategoryIDGenerateParams contains all the parameters to send to the API endpoint

	for the post log user categories log category ID generate operation.

	Typically these are written to a http.Request.
*/
type PostLogUserCategoriesLogCategoryIDGenerateParams struct {

	// Body.
	Body *models.QueryFilter

	// DataSize.
	//
	// Format: int32
	DataSize *int32

	// LastIndex.
	//
	// Format: int64
	LastIndex *int64

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

// WithDefaults hydrates default values in the post log user categories log category ID generate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLogUserCategoriesLogCategoryIDGenerateParams) WithDefaults() *PostLogUserCategoriesLogCategoryIDGenerateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post log user categories log category ID generate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLogUserCategoriesLogCategoryIDGenerateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post log user categories log category ID generate params
func (o *PostLogUserCategoriesLogCategoryIDGenerateParams) WithTimeout(timeout time.Duration) *PostLogUserCategoriesLogCategoryIDGenerateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post log user categories log category ID generate params
func (o *PostLogUserCategoriesLogCategoryIDGenerateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post log user categories log category ID generate params
func (o *PostLogUserCategoriesLogCategoryIDGenerateParams) WithContext(ctx context.Context) *PostLogUserCategoriesLogCategoryIDGenerateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post log user categories log category ID generate params
func (o *PostLogUserCategoriesLogCategoryIDGenerateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post log user categories log category ID generate params
func (o *PostLogUserCategoriesLogCategoryIDGenerateParams) WithHTTPClient(client *http.Client) *PostLogUserCategoriesLogCategoryIDGenerateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post log user categories log category ID generate params
func (o *PostLogUserCategoriesLogCategoryIDGenerateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post log user categories log category ID generate params
func (o *PostLogUserCategoriesLogCategoryIDGenerateParams) WithBody(body *models.QueryFilter) *PostLogUserCategoriesLogCategoryIDGenerateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post log user categories log category ID generate params
func (o *PostLogUserCategoriesLogCategoryIDGenerateParams) SetBody(body *models.QueryFilter) {
	o.Body = body
}

// WithDataSize adds the dataSize to the post log user categories log category ID generate params
func (o *PostLogUserCategoriesLogCategoryIDGenerateParams) WithDataSize(dataSize *int32) *PostLogUserCategoriesLogCategoryIDGenerateParams {
	o.SetDataSize(dataSize)
	return o
}

// SetDataSize adds the dataSize to the post log user categories log category ID generate params
func (o *PostLogUserCategoriesLogCategoryIDGenerateParams) SetDataSize(dataSize *int32) {
	o.DataSize = dataSize
}

// WithLastIndex adds the lastIndex to the post log user categories log category ID generate params
func (o *PostLogUserCategoriesLogCategoryIDGenerateParams) WithLastIndex(lastIndex *int64) *PostLogUserCategoriesLogCategoryIDGenerateParams {
	o.SetLastIndex(lastIndex)
	return o
}

// SetLastIndex adds the lastIndex to the post log user categories log category ID generate params
func (o *PostLogUserCategoriesLogCategoryIDGenerateParams) SetLastIndex(lastIndex *int64) {
	o.LastIndex = lastIndex
}

// WithLogCategoryID adds the logCategoryID to the post log user categories log category ID generate params
func (o *PostLogUserCategoriesLogCategoryIDGenerateParams) WithLogCategoryID(logCategoryID int32) *PostLogUserCategoriesLogCategoryIDGenerateParams {
	o.SetLogCategoryID(logCategoryID)
	return o
}

// SetLogCategoryID adds the logCategoryId to the post log user categories log category ID generate params
func (o *PostLogUserCategoriesLogCategoryIDGenerateParams) SetLogCategoryID(logCategoryID int32) {
	o.LogCategoryID = logCategoryID
}

// WithUserID adds the userID to the post log user categories log category ID generate params
func (o *PostLogUserCategoriesLogCategoryIDGenerateParams) WithUserID(userID *strfmt.UUID) *PostLogUserCategoriesLogCategoryIDGenerateParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the post log user categories log category ID generate params
func (o *PostLogUserCategoriesLogCategoryIDGenerateParams) SetUserID(userID *strfmt.UUID) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLogUserCategoriesLogCategoryIDGenerateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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

	if o.LastIndex != nil {

		// query param lastIndex
		var qrLastIndex int64

		if o.LastIndex != nil {
			qrLastIndex = *o.LastIndex
		}
		qLastIndex := swag.FormatInt64(qrLastIndex)
		if qLastIndex != "" {

			if err := r.SetQueryParam("lastIndex", qLastIndex); err != nil {
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
