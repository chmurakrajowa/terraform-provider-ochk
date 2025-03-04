// Code generated by go-swagger; DO NOT EDIT.

package gfw_rule

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

// NewPutNetworkRoutersRouterIDRulesSNParams creates a new PutNetworkRoutersRouterIDRulesSNParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutNetworkRoutersRouterIDRulesSNParams() *PutNetworkRoutersRouterIDRulesSNParams {
	return &PutNetworkRoutersRouterIDRulesSNParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutNetworkRoutersRouterIDRulesSNParamsWithTimeout creates a new PutNetworkRoutersRouterIDRulesSNParams object
// with the ability to set a timeout on a request.
func NewPutNetworkRoutersRouterIDRulesSNParamsWithTimeout(timeout time.Duration) *PutNetworkRoutersRouterIDRulesSNParams {
	return &PutNetworkRoutersRouterIDRulesSNParams{
		timeout: timeout,
	}
}

// NewPutNetworkRoutersRouterIDRulesSNParamsWithContext creates a new PutNetworkRoutersRouterIDRulesSNParams object
// with the ability to set a context for a request.
func NewPutNetworkRoutersRouterIDRulesSNParamsWithContext(ctx context.Context) *PutNetworkRoutersRouterIDRulesSNParams {
	return &PutNetworkRoutersRouterIDRulesSNParams{
		Context: ctx,
	}
}

// NewPutNetworkRoutersRouterIDRulesSNParamsWithHTTPClient creates a new PutNetworkRoutersRouterIDRulesSNParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutNetworkRoutersRouterIDRulesSNParamsWithHTTPClient(client *http.Client) *PutNetworkRoutersRouterIDRulesSNParams {
	return &PutNetworkRoutersRouterIDRulesSNParams{
		HTTPClient: client,
	}
}

/*
PutNetworkRoutersRouterIDRulesSNParams contains all the parameters to send to the API endpoint

	for the put network routers router ID rules s n operation.

	Typically these are written to a http.Request.
*/
type PutNetworkRoutersRouterIDRulesSNParams struct {

	// Body.
	Body *models.GfwRule

	// RouterID.
	//
	// Format: uuid
	RouterID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put network routers router ID rules s n params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNetworkRoutersRouterIDRulesSNParams) WithDefaults() *PutNetworkRoutersRouterIDRulesSNParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put network routers router ID rules s n params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNetworkRoutersRouterIDRulesSNParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put network routers router ID rules s n params
func (o *PutNetworkRoutersRouterIDRulesSNParams) WithTimeout(timeout time.Duration) *PutNetworkRoutersRouterIDRulesSNParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put network routers router ID rules s n params
func (o *PutNetworkRoutersRouterIDRulesSNParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put network routers router ID rules s n params
func (o *PutNetworkRoutersRouterIDRulesSNParams) WithContext(ctx context.Context) *PutNetworkRoutersRouterIDRulesSNParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put network routers router ID rules s n params
func (o *PutNetworkRoutersRouterIDRulesSNParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put network routers router ID rules s n params
func (o *PutNetworkRoutersRouterIDRulesSNParams) WithHTTPClient(client *http.Client) *PutNetworkRoutersRouterIDRulesSNParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put network routers router ID rules s n params
func (o *PutNetworkRoutersRouterIDRulesSNParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put network routers router ID rules s n params
func (o *PutNetworkRoutersRouterIDRulesSNParams) WithBody(body *models.GfwRule) *PutNetworkRoutersRouterIDRulesSNParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put network routers router ID rules s n params
func (o *PutNetworkRoutersRouterIDRulesSNParams) SetBody(body *models.GfwRule) {
	o.Body = body
}

// WithRouterID adds the routerID to the put network routers router ID rules s n params
func (o *PutNetworkRoutersRouterIDRulesSNParams) WithRouterID(routerID strfmt.UUID) *PutNetworkRoutersRouterIDRulesSNParams {
	o.SetRouterID(routerID)
	return o
}

// SetRouterID adds the routerId to the put network routers router ID rules s n params
func (o *PutNetworkRoutersRouterIDRulesSNParams) SetRouterID(routerID strfmt.UUID) {
	o.RouterID = routerID
}

// WriteToRequest writes these params to a swagger request
func (o *PutNetworkRoutersRouterIDRulesSNParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
