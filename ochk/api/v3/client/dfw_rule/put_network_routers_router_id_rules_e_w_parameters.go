// Code generated by go-swagger; DO NOT EDIT.

package dfw_rule

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

// NewPutNetworkRoutersRouterIDRulesEWParams creates a new PutNetworkRoutersRouterIDRulesEWParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutNetworkRoutersRouterIDRulesEWParams() *PutNetworkRoutersRouterIDRulesEWParams {
	return &PutNetworkRoutersRouterIDRulesEWParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutNetworkRoutersRouterIDRulesEWParamsWithTimeout creates a new PutNetworkRoutersRouterIDRulesEWParams object
// with the ability to set a timeout on a request.
func NewPutNetworkRoutersRouterIDRulesEWParamsWithTimeout(timeout time.Duration) *PutNetworkRoutersRouterIDRulesEWParams {
	return &PutNetworkRoutersRouterIDRulesEWParams{
		timeout: timeout,
	}
}

// NewPutNetworkRoutersRouterIDRulesEWParamsWithContext creates a new PutNetworkRoutersRouterIDRulesEWParams object
// with the ability to set a context for a request.
func NewPutNetworkRoutersRouterIDRulesEWParamsWithContext(ctx context.Context) *PutNetworkRoutersRouterIDRulesEWParams {
	return &PutNetworkRoutersRouterIDRulesEWParams{
		Context: ctx,
	}
}

// NewPutNetworkRoutersRouterIDRulesEWParamsWithHTTPClient creates a new PutNetworkRoutersRouterIDRulesEWParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutNetworkRoutersRouterIDRulesEWParamsWithHTTPClient(client *http.Client) *PutNetworkRoutersRouterIDRulesEWParams {
	return &PutNetworkRoutersRouterIDRulesEWParams{
		HTTPClient: client,
	}
}

/*
PutNetworkRoutersRouterIDRulesEWParams contains all the parameters to send to the API endpoint

	for the put network routers router ID rules e w operation.

	Typically these are written to a http.Request.
*/
type PutNetworkRoutersRouterIDRulesEWParams struct {

	// Body.
	Body *models.DfwRule

	// RouterID.
	//
	// Format: uuid
	RouterID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put network routers router ID rules e w params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNetworkRoutersRouterIDRulesEWParams) WithDefaults() *PutNetworkRoutersRouterIDRulesEWParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put network routers router ID rules e w params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNetworkRoutersRouterIDRulesEWParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put network routers router ID rules e w params
func (o *PutNetworkRoutersRouterIDRulesEWParams) WithTimeout(timeout time.Duration) *PutNetworkRoutersRouterIDRulesEWParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put network routers router ID rules e w params
func (o *PutNetworkRoutersRouterIDRulesEWParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put network routers router ID rules e w params
func (o *PutNetworkRoutersRouterIDRulesEWParams) WithContext(ctx context.Context) *PutNetworkRoutersRouterIDRulesEWParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put network routers router ID rules e w params
func (o *PutNetworkRoutersRouterIDRulesEWParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put network routers router ID rules e w params
func (o *PutNetworkRoutersRouterIDRulesEWParams) WithHTTPClient(client *http.Client) *PutNetworkRoutersRouterIDRulesEWParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put network routers router ID rules e w params
func (o *PutNetworkRoutersRouterIDRulesEWParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put network routers router ID rules e w params
func (o *PutNetworkRoutersRouterIDRulesEWParams) WithBody(body *models.DfwRule) *PutNetworkRoutersRouterIDRulesEWParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put network routers router ID rules e w params
func (o *PutNetworkRoutersRouterIDRulesEWParams) SetBody(body *models.DfwRule) {
	o.Body = body
}

// WithRouterID adds the routerID to the put network routers router ID rules e w params
func (o *PutNetworkRoutersRouterIDRulesEWParams) WithRouterID(routerID strfmt.UUID) *PutNetworkRoutersRouterIDRulesEWParams {
	o.SetRouterID(routerID)
	return o
}

// SetRouterID adds the routerId to the put network routers router ID rules e w params
func (o *PutNetworkRoutersRouterIDRulesEWParams) SetRouterID(routerID strfmt.UUID) {
	o.RouterID = routerID
}

// WriteToRequest writes these params to a swagger request
func (o *PutNetworkRoutersRouterIDRulesEWParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param routerId
	if err := r.SetPathParam("routerId", o.RouterID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
