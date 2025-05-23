// Code generated by go-swagger; DO NOT EDIT.

package ip_collection

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

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
)

// NewPutIpcsIPCollectionIDParams creates a new PutIpcsIPCollectionIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutIpcsIPCollectionIDParams() *PutIpcsIPCollectionIDParams {
	return &PutIpcsIPCollectionIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutIpcsIPCollectionIDParamsWithTimeout creates a new PutIpcsIPCollectionIDParams object
// with the ability to set a timeout on a request.
func NewPutIpcsIPCollectionIDParamsWithTimeout(timeout time.Duration) *PutIpcsIPCollectionIDParams {
	return &PutIpcsIPCollectionIDParams{
		timeout: timeout,
	}
}

// NewPutIpcsIPCollectionIDParamsWithContext creates a new PutIpcsIPCollectionIDParams object
// with the ability to set a context for a request.
func NewPutIpcsIPCollectionIDParamsWithContext(ctx context.Context) *PutIpcsIPCollectionIDParams {
	return &PutIpcsIPCollectionIDParams{
		Context: ctx,
	}
}

// NewPutIpcsIPCollectionIDParamsWithHTTPClient creates a new PutIpcsIPCollectionIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutIpcsIPCollectionIDParamsWithHTTPClient(client *http.Client) *PutIpcsIPCollectionIDParams {
	return &PutIpcsIPCollectionIDParams{
		HTTPClient: client,
	}
}

/*
PutIpcsIPCollectionIDParams contains all the parameters to send to the API endpoint

	for the put ipcs IP collection ID operation.

	Typically these are written to a http.Request.
*/
type PutIpcsIPCollectionIDParams struct {

	// Body.
	Body *models.IPCollection

	// IPCollectionID.
	//
	// Format: uuid
	IPCollectionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put ipcs IP collection ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutIpcsIPCollectionIDParams) WithDefaults() *PutIpcsIPCollectionIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put ipcs IP collection ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutIpcsIPCollectionIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put ipcs IP collection ID params
func (o *PutIpcsIPCollectionIDParams) WithTimeout(timeout time.Duration) *PutIpcsIPCollectionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put ipcs IP collection ID params
func (o *PutIpcsIPCollectionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put ipcs IP collection ID params
func (o *PutIpcsIPCollectionIDParams) WithContext(ctx context.Context) *PutIpcsIPCollectionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put ipcs IP collection ID params
func (o *PutIpcsIPCollectionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put ipcs IP collection ID params
func (o *PutIpcsIPCollectionIDParams) WithHTTPClient(client *http.Client) *PutIpcsIPCollectionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put ipcs IP collection ID params
func (o *PutIpcsIPCollectionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put ipcs IP collection ID params
func (o *PutIpcsIPCollectionIDParams) WithBody(body *models.IPCollection) *PutIpcsIPCollectionIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put ipcs IP collection ID params
func (o *PutIpcsIPCollectionIDParams) SetBody(body *models.IPCollection) {
	o.Body = body
}

// WithIPCollectionID adds the iPCollectionID to the put ipcs IP collection ID params
func (o *PutIpcsIPCollectionIDParams) WithIPCollectionID(iPCollectionID strfmt.UUID) *PutIpcsIPCollectionIDParams {
	o.SetIPCollectionID(iPCollectionID)
	return o
}

// SetIPCollectionID adds the ipCollectionId to the put ipcs IP collection ID params
func (o *PutIpcsIPCollectionIDParams) SetIPCollectionID(iPCollectionID strfmt.UUID) {
	o.IPCollectionID = iPCollectionID
}

// WriteToRequest writes these params to a swagger request
func (o *PutIpcsIPCollectionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param ipCollectionId
	if err := r.SetPathParam("ipCollectionId", o.IPCollectionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
