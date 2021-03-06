// Code generated by go-swagger; DO NOT EDIT.

package edge_clusters

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

// NewEdgeClusterGetUsingGETParams creates a new EdgeClusterGetUsingGETParams object
// with the default values initialized.
func NewEdgeClusterGetUsingGETParams() *EdgeClusterGetUsingGETParams {
	var ()
	return &EdgeClusterGetUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeClusterGetUsingGETParamsWithTimeout creates a new EdgeClusterGetUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEdgeClusterGetUsingGETParamsWithTimeout(timeout time.Duration) *EdgeClusterGetUsingGETParams {
	var ()
	return &EdgeClusterGetUsingGETParams{

		timeout: timeout,
	}
}

// NewEdgeClusterGetUsingGETParamsWithContext creates a new EdgeClusterGetUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewEdgeClusterGetUsingGETParamsWithContext(ctx context.Context) *EdgeClusterGetUsingGETParams {
	var ()
	return &EdgeClusterGetUsingGETParams{

		Context: ctx,
	}
}

// NewEdgeClusterGetUsingGETParamsWithHTTPClient creates a new EdgeClusterGetUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEdgeClusterGetUsingGETParamsWithHTTPClient(client *http.Client) *EdgeClusterGetUsingGETParams {
	var ()
	return &EdgeClusterGetUsingGETParams{
		HTTPClient: client,
	}
}

/*EdgeClusterGetUsingGETParams contains all the parameters to send to the API endpoint
for the edge cluster get using g e t operation typically these are written to a http.Request
*/
type EdgeClusterGetUsingGETParams struct {

	/*EdgeClusterID
	  edgeClusterId

	*/
	EdgeClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edge cluster get using g e t params
func (o *EdgeClusterGetUsingGETParams) WithTimeout(timeout time.Duration) *EdgeClusterGetUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge cluster get using g e t params
func (o *EdgeClusterGetUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge cluster get using g e t params
func (o *EdgeClusterGetUsingGETParams) WithContext(ctx context.Context) *EdgeClusterGetUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge cluster get using g e t params
func (o *EdgeClusterGetUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge cluster get using g e t params
func (o *EdgeClusterGetUsingGETParams) WithHTTPClient(client *http.Client) *EdgeClusterGetUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge cluster get using g e t params
func (o *EdgeClusterGetUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEdgeClusterID adds the edgeClusterID to the edge cluster get using g e t params
func (o *EdgeClusterGetUsingGETParams) WithEdgeClusterID(edgeClusterID string) *EdgeClusterGetUsingGETParams {
	o.SetEdgeClusterID(edgeClusterID)
	return o
}

// SetEdgeClusterID adds the edgeClusterId to the edge cluster get using g e t params
func (o *EdgeClusterGetUsingGETParams) SetEdgeClusterID(edgeClusterID string) {
	o.EdgeClusterID = edgeClusterID
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeClusterGetUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param edgeClusterId
	if err := r.SetPathParam("edgeClusterId", o.EdgeClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
