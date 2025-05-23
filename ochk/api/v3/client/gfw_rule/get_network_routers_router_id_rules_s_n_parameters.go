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
)

// NewGetNetworkRoutersRouterIDRulesSNParams creates a new GetNetworkRoutersRouterIDRulesSNParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkRoutersRouterIDRulesSNParams() *GetNetworkRoutersRouterIDRulesSNParams {
	return &GetNetworkRoutersRouterIDRulesSNParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkRoutersRouterIDRulesSNParamsWithTimeout creates a new GetNetworkRoutersRouterIDRulesSNParams object
// with the ability to set a timeout on a request.
func NewGetNetworkRoutersRouterIDRulesSNParamsWithTimeout(timeout time.Duration) *GetNetworkRoutersRouterIDRulesSNParams {
	return &GetNetworkRoutersRouterIDRulesSNParams{
		timeout: timeout,
	}
}

// NewGetNetworkRoutersRouterIDRulesSNParamsWithContext creates a new GetNetworkRoutersRouterIDRulesSNParams object
// with the ability to set a context for a request.
func NewGetNetworkRoutersRouterIDRulesSNParamsWithContext(ctx context.Context) *GetNetworkRoutersRouterIDRulesSNParams {
	return &GetNetworkRoutersRouterIDRulesSNParams{
		Context: ctx,
	}
}

// NewGetNetworkRoutersRouterIDRulesSNParamsWithHTTPClient creates a new GetNetworkRoutersRouterIDRulesSNParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkRoutersRouterIDRulesSNParamsWithHTTPClient(client *http.Client) *GetNetworkRoutersRouterIDRulesSNParams {
	return &GetNetworkRoutersRouterIDRulesSNParams{
		HTTPClient: client,
	}
}

/*
GetNetworkRoutersRouterIDRulesSNParams contains all the parameters to send to the API endpoint

	for the get network routers router ID rules s n operation.

	Typically these are written to a http.Request.
*/
type GetNetworkRoutersRouterIDRulesSNParams struct {

	// DisplayName.
	DisplayName *string

	// RouterID.
	//
	// Format: uuid
	RouterID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network routers router ID rules s n params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkRoutersRouterIDRulesSNParams) WithDefaults() *GetNetworkRoutersRouterIDRulesSNParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network routers router ID rules s n params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkRoutersRouterIDRulesSNParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network routers router ID rules s n params
func (o *GetNetworkRoutersRouterIDRulesSNParams) WithTimeout(timeout time.Duration) *GetNetworkRoutersRouterIDRulesSNParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network routers router ID rules s n params
func (o *GetNetworkRoutersRouterIDRulesSNParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network routers router ID rules s n params
func (o *GetNetworkRoutersRouterIDRulesSNParams) WithContext(ctx context.Context) *GetNetworkRoutersRouterIDRulesSNParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network routers router ID rules s n params
func (o *GetNetworkRoutersRouterIDRulesSNParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network routers router ID rules s n params
func (o *GetNetworkRoutersRouterIDRulesSNParams) WithHTTPClient(client *http.Client) *GetNetworkRoutersRouterIDRulesSNParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network routers router ID rules s n params
func (o *GetNetworkRoutersRouterIDRulesSNParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisplayName adds the displayName to the get network routers router ID rules s n params
func (o *GetNetworkRoutersRouterIDRulesSNParams) WithDisplayName(displayName *string) *GetNetworkRoutersRouterIDRulesSNParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the get network routers router ID rules s n params
func (o *GetNetworkRoutersRouterIDRulesSNParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WithRouterID adds the routerID to the get network routers router ID rules s n params
func (o *GetNetworkRoutersRouterIDRulesSNParams) WithRouterID(routerID strfmt.UUID) *GetNetworkRoutersRouterIDRulesSNParams {
	o.SetRouterID(routerID)
	return o
}

// SetRouterID adds the routerId to the get network routers router ID rules s n params
func (o *GetNetworkRoutersRouterIDRulesSNParams) SetRouterID(routerID strfmt.UUID) {
	o.RouterID = routerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkRoutersRouterIDRulesSNParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DisplayName != nil {

		// query param displayName
		var qrDisplayName string

		if o.DisplayName != nil {
			qrDisplayName = *o.DisplayName
		}
		qDisplayName := qrDisplayName
		if qDisplayName != "" {

			if err := r.SetQueryParam("displayName", qDisplayName); err != nil {
				return err
			}
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
