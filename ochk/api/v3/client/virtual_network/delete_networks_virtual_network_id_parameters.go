// Code generated by go-swagger; DO NOT EDIT.

package virtual_network

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

// NewDeleteNetworksVirtualNetworkIDParams creates a new DeleteNetworksVirtualNetworkIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNetworksVirtualNetworkIDParams() *DeleteNetworksVirtualNetworkIDParams {
	return &DeleteNetworksVirtualNetworkIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNetworksVirtualNetworkIDParamsWithTimeout creates a new DeleteNetworksVirtualNetworkIDParams object
// with the ability to set a timeout on a request.
func NewDeleteNetworksVirtualNetworkIDParamsWithTimeout(timeout time.Duration) *DeleteNetworksVirtualNetworkIDParams {
	return &DeleteNetworksVirtualNetworkIDParams{
		timeout: timeout,
	}
}

// NewDeleteNetworksVirtualNetworkIDParamsWithContext creates a new DeleteNetworksVirtualNetworkIDParams object
// with the ability to set a context for a request.
func NewDeleteNetworksVirtualNetworkIDParamsWithContext(ctx context.Context) *DeleteNetworksVirtualNetworkIDParams {
	return &DeleteNetworksVirtualNetworkIDParams{
		Context: ctx,
	}
}

// NewDeleteNetworksVirtualNetworkIDParamsWithHTTPClient creates a new DeleteNetworksVirtualNetworkIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNetworksVirtualNetworkIDParamsWithHTTPClient(client *http.Client) *DeleteNetworksVirtualNetworkIDParams {
	return &DeleteNetworksVirtualNetworkIDParams{
		HTTPClient: client,
	}
}

/*
DeleteNetworksVirtualNetworkIDParams contains all the parameters to send to the API endpoint

	for the delete networks virtual network ID operation.

	Typically these are written to a http.Request.
*/
type DeleteNetworksVirtualNetworkIDParams struct {

	// VirtualNetworkID.
	//
	// Format: uuid
	VirtualNetworkID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete networks virtual network ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworksVirtualNetworkIDParams) WithDefaults() *DeleteNetworksVirtualNetworkIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete networks virtual network ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworksVirtualNetworkIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete networks virtual network ID params
func (o *DeleteNetworksVirtualNetworkIDParams) WithTimeout(timeout time.Duration) *DeleteNetworksVirtualNetworkIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete networks virtual network ID params
func (o *DeleteNetworksVirtualNetworkIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete networks virtual network ID params
func (o *DeleteNetworksVirtualNetworkIDParams) WithContext(ctx context.Context) *DeleteNetworksVirtualNetworkIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete networks virtual network ID params
func (o *DeleteNetworksVirtualNetworkIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete networks virtual network ID params
func (o *DeleteNetworksVirtualNetworkIDParams) WithHTTPClient(client *http.Client) *DeleteNetworksVirtualNetworkIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete networks virtual network ID params
func (o *DeleteNetworksVirtualNetworkIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVirtualNetworkID adds the virtualNetworkID to the delete networks virtual network ID params
func (o *DeleteNetworksVirtualNetworkIDParams) WithVirtualNetworkID(virtualNetworkID strfmt.UUID) *DeleteNetworksVirtualNetworkIDParams {
	o.SetVirtualNetworkID(virtualNetworkID)
	return o
}

// SetVirtualNetworkID adds the virtualNetworkId to the delete networks virtual network ID params
func (o *DeleteNetworksVirtualNetworkIDParams) SetVirtualNetworkID(virtualNetworkID strfmt.UUID) {
	o.VirtualNetworkID = virtualNetworkID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNetworksVirtualNetworkIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param virtualNetworkId
	if err := r.SetPathParam("virtualNetworkId", o.VirtualNetworkID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
