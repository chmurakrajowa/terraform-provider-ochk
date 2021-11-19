// Code generated by go-swagger; DO NOT EDIT.

package ip_collections

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

// NewIPCollectionUpdateUsingPUTParams creates a new IPCollectionUpdateUsingPUTParams object
// with the default values initialized.
func NewIPCollectionUpdateUsingPUTParams() *IPCollectionUpdateUsingPUTParams {
	var ()
	return &IPCollectionUpdateUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPCollectionUpdateUsingPUTParamsWithTimeout creates a new IPCollectionUpdateUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPCollectionUpdateUsingPUTParamsWithTimeout(timeout time.Duration) *IPCollectionUpdateUsingPUTParams {
	var ()
	return &IPCollectionUpdateUsingPUTParams{

		timeout: timeout,
	}
}

// NewIPCollectionUpdateUsingPUTParamsWithContext creates a new IPCollectionUpdateUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPCollectionUpdateUsingPUTParamsWithContext(ctx context.Context) *IPCollectionUpdateUsingPUTParams {
	var ()
	return &IPCollectionUpdateUsingPUTParams{

		Context: ctx,
	}
}

// NewIPCollectionUpdateUsingPUTParamsWithHTTPClient creates a new IPCollectionUpdateUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPCollectionUpdateUsingPUTParamsWithHTTPClient(client *http.Client) *IPCollectionUpdateUsingPUTParams {
	var ()
	return &IPCollectionUpdateUsingPUTParams{
		HTTPClient: client,
	}
}

/*IPCollectionUpdateUsingPUTParams contains all the parameters to send to the API endpoint
for the ip collection update using p u t operation typically these are written to a http.Request
*/
type IPCollectionUpdateUsingPUTParams struct {

	/*IPCollection
	  ipCollection

	*/
	IPCollection *models.IPCollection
	/*IPCollectionID
	  ipCollectionId

	*/
	IPCollectionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ip collection update using p u t params
func (o *IPCollectionUpdateUsingPUTParams) WithTimeout(timeout time.Duration) *IPCollectionUpdateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ip collection update using p u t params
func (o *IPCollectionUpdateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ip collection update using p u t params
func (o *IPCollectionUpdateUsingPUTParams) WithContext(ctx context.Context) *IPCollectionUpdateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ip collection update using p u t params
func (o *IPCollectionUpdateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ip collection update using p u t params
func (o *IPCollectionUpdateUsingPUTParams) WithHTTPClient(client *http.Client) *IPCollectionUpdateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ip collection update using p u t params
func (o *IPCollectionUpdateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIPCollection adds the iPCollection to the ip collection update using p u t params
func (o *IPCollectionUpdateUsingPUTParams) WithIPCollection(iPCollection *models.IPCollection) *IPCollectionUpdateUsingPUTParams {
	o.SetIPCollection(iPCollection)
	return o
}

// SetIPCollection adds the ipCollection to the ip collection update using p u t params
func (o *IPCollectionUpdateUsingPUTParams) SetIPCollection(iPCollection *models.IPCollection) {
	o.IPCollection = iPCollection
}

// WithIPCollectionID adds the iPCollectionID to the ip collection update using p u t params
func (o *IPCollectionUpdateUsingPUTParams) WithIPCollectionID(iPCollectionID string) *IPCollectionUpdateUsingPUTParams {
	o.SetIPCollectionID(iPCollectionID)
	return o
}

// SetIPCollectionID adds the ipCollectionId to the ip collection update using p u t params
func (o *IPCollectionUpdateUsingPUTParams) SetIPCollectionID(iPCollectionID string) {
	o.IPCollectionID = iPCollectionID
}

// WriteToRequest writes these params to a swagger request
func (o *IPCollectionUpdateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IPCollection != nil {
		if err := r.SetBodyParam(o.IPCollection); err != nil {
			return err
		}
	}

	// path param ipCollectionId
	if err := r.SetPathParam("ipCollectionId", o.IPCollectionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
