// Code generated by go-swagger; DO NOT EDIT.

package logs

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

// NewGetHistogramSeverityUsingPOST1Params creates a new GetHistogramSeverityUsingPOST1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHistogramSeverityUsingPOST1Params() *GetHistogramSeverityUsingPOST1Params {
	return &GetHistogramSeverityUsingPOST1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHistogramSeverityUsingPOST1ParamsWithTimeout creates a new GetHistogramSeverityUsingPOST1Params object
// with the ability to set a timeout on a request.
func NewGetHistogramSeverityUsingPOST1ParamsWithTimeout(timeout time.Duration) *GetHistogramSeverityUsingPOST1Params {
	return &GetHistogramSeverityUsingPOST1Params{
		timeout: timeout,
	}
}

// NewGetHistogramSeverityUsingPOST1ParamsWithContext creates a new GetHistogramSeverityUsingPOST1Params object
// with the ability to set a context for a request.
func NewGetHistogramSeverityUsingPOST1ParamsWithContext(ctx context.Context) *GetHistogramSeverityUsingPOST1Params {
	return &GetHistogramSeverityUsingPOST1Params{
		Context: ctx,
	}
}

// NewGetHistogramSeverityUsingPOST1ParamsWithHTTPClient creates a new GetHistogramSeverityUsingPOST1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetHistogramSeverityUsingPOST1ParamsWithHTTPClient(client *http.Client) *GetHistogramSeverityUsingPOST1Params {
	return &GetHistogramSeverityUsingPOST1Params{
		HTTPClient: client,
	}
}

/*
GetHistogramSeverityUsingPOST1Params contains all the parameters to send to the API endpoint

	for the get histogram severity using p o s t 1 operation.

	Typically these are written to a http.Request.
*/
type GetHistogramSeverityUsingPOST1Params struct {

	/* CalendarInterval.

	   calendarInterval
	*/
	CalendarInterval *string

	/* QueryFilter.

	   queryFilter
	*/
	QueryFilter *models.QueryFilter

	/* UserID.

	   userId
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get histogram severity using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHistogramSeverityUsingPOST1Params) WithDefaults() *GetHistogramSeverityUsingPOST1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get histogram severity using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHistogramSeverityUsingPOST1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get histogram severity using p o s t 1 params
func (o *GetHistogramSeverityUsingPOST1Params) WithTimeout(timeout time.Duration) *GetHistogramSeverityUsingPOST1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get histogram severity using p o s t 1 params
func (o *GetHistogramSeverityUsingPOST1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get histogram severity using p o s t 1 params
func (o *GetHistogramSeverityUsingPOST1Params) WithContext(ctx context.Context) *GetHistogramSeverityUsingPOST1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get histogram severity using p o s t 1 params
func (o *GetHistogramSeverityUsingPOST1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get histogram severity using p o s t 1 params
func (o *GetHistogramSeverityUsingPOST1Params) WithHTTPClient(client *http.Client) *GetHistogramSeverityUsingPOST1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get histogram severity using p o s t 1 params
func (o *GetHistogramSeverityUsingPOST1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCalendarInterval adds the calendarInterval to the get histogram severity using p o s t 1 params
func (o *GetHistogramSeverityUsingPOST1Params) WithCalendarInterval(calendarInterval *string) *GetHistogramSeverityUsingPOST1Params {
	o.SetCalendarInterval(calendarInterval)
	return o
}

// SetCalendarInterval adds the calendarInterval to the get histogram severity using p o s t 1 params
func (o *GetHistogramSeverityUsingPOST1Params) SetCalendarInterval(calendarInterval *string) {
	o.CalendarInterval = calendarInterval
}

// WithQueryFilter adds the queryFilter to the get histogram severity using p o s t 1 params
func (o *GetHistogramSeverityUsingPOST1Params) WithQueryFilter(queryFilter *models.QueryFilter) *GetHistogramSeverityUsingPOST1Params {
	o.SetQueryFilter(queryFilter)
	return o
}

// SetQueryFilter adds the queryFilter to the get histogram severity using p o s t 1 params
func (o *GetHistogramSeverityUsingPOST1Params) SetQueryFilter(queryFilter *models.QueryFilter) {
	o.QueryFilter = queryFilter
}

// WithUserID adds the userID to the get histogram severity using p o s t 1 params
func (o *GetHistogramSeverityUsingPOST1Params) WithUserID(userID string) *GetHistogramSeverityUsingPOST1Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get histogram severity using p o s t 1 params
func (o *GetHistogramSeverityUsingPOST1Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetHistogramSeverityUsingPOST1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CalendarInterval != nil {

		// query param calendarInterval
		var qrCalendarInterval string

		if o.CalendarInterval != nil {
			qrCalendarInterval = *o.CalendarInterval
		}
		qCalendarInterval := qrCalendarInterval
		if qCalendarInterval != "" {

			if err := r.SetQueryParam("calendarInterval", qCalendarInterval); err != nil {
				return err
			}
		}
	}
	if o.QueryFilter != nil {
		if err := r.SetBodyParam(o.QueryFilter); err != nil {
			return err
		}
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