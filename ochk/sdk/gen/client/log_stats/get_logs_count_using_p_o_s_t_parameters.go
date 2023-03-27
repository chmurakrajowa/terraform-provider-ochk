// Code generated by go-swagger; DO NOT EDIT.

package log_stats

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

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// NewGetLogsCountUsingPOSTParams creates a new GetLogsCountUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLogsCountUsingPOSTParams() *GetLogsCountUsingPOSTParams {
	return &GetLogsCountUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogsCountUsingPOSTParamsWithTimeout creates a new GetLogsCountUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewGetLogsCountUsingPOSTParamsWithTimeout(timeout time.Duration) *GetLogsCountUsingPOSTParams {
	return &GetLogsCountUsingPOSTParams{
		timeout: timeout,
	}
}

// NewGetLogsCountUsingPOSTParamsWithContext creates a new GetLogsCountUsingPOSTParams object
// with the ability to set a context for a request.
func NewGetLogsCountUsingPOSTParamsWithContext(ctx context.Context) *GetLogsCountUsingPOSTParams {
	return &GetLogsCountUsingPOSTParams{
		Context: ctx,
	}
}

// NewGetLogsCountUsingPOSTParamsWithHTTPClient creates a new GetLogsCountUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLogsCountUsingPOSTParamsWithHTTPClient(client *http.Client) *GetLogsCountUsingPOSTParams {
	return &GetLogsCountUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
GetLogsCountUsingPOSTParams contains all the parameters to send to the API endpoint

	for the get logs count using p o s t operation.

	Typically these are written to a http.Request.
*/
type GetLogsCountUsingPOSTParams struct {

	/* LogCategoryID.

	   logCategoryId

	   Format: int32
	*/
	LogCategoryID int32

	/* QueryFilter.

	   queryFilter
	*/
	QueryFilter *models.QueryFilter

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get logs count using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogsCountUsingPOSTParams) WithDefaults() *GetLogsCountUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get logs count using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogsCountUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get logs count using p o s t params
func (o *GetLogsCountUsingPOSTParams) WithTimeout(timeout time.Duration) *GetLogsCountUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get logs count using p o s t params
func (o *GetLogsCountUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get logs count using p o s t params
func (o *GetLogsCountUsingPOSTParams) WithContext(ctx context.Context) *GetLogsCountUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get logs count using p o s t params
func (o *GetLogsCountUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get logs count using p o s t params
func (o *GetLogsCountUsingPOSTParams) WithHTTPClient(client *http.Client) *GetLogsCountUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get logs count using p o s t params
func (o *GetLogsCountUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLogCategoryID adds the logCategoryID to the get logs count using p o s t params
func (o *GetLogsCountUsingPOSTParams) WithLogCategoryID(logCategoryID int32) *GetLogsCountUsingPOSTParams {
	o.SetLogCategoryID(logCategoryID)
	return o
}

// SetLogCategoryID adds the logCategoryId to the get logs count using p o s t params
func (o *GetLogsCountUsingPOSTParams) SetLogCategoryID(logCategoryID int32) {
	o.LogCategoryID = logCategoryID
}

// WithQueryFilter adds the queryFilter to the get logs count using p o s t params
func (o *GetLogsCountUsingPOSTParams) WithQueryFilter(queryFilter *models.QueryFilter) *GetLogsCountUsingPOSTParams {
	o.SetQueryFilter(queryFilter)
	return o
}

// SetQueryFilter adds the queryFilter to the get logs count using p o s t params
func (o *GetLogsCountUsingPOSTParams) SetQueryFilter(queryFilter *models.QueryFilter) {
	o.QueryFilter = queryFilter
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogsCountUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param logCategoryId
	if err := r.SetPathParam("logCategoryId", swag.FormatInt32(o.LogCategoryID)); err != nil {
		return err
	}
	if o.QueryFilter != nil {
		if err := r.SetBodyParam(o.QueryFilter); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
