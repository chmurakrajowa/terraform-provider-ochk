// Code generated by go-swagger; DO NOT EDIT.

package subtenant_custom_groups

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

// NewSubtenantGroupListUsingGETParams creates a new SubtenantGroupListUsingGETParams object
// with the default values initialized.
func NewSubtenantGroupListUsingGETParams() *SubtenantGroupListUsingGETParams {
	var ()
	return &SubtenantGroupListUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubtenantGroupListUsingGETParamsWithTimeout creates a new SubtenantGroupListUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubtenantGroupListUsingGETParamsWithTimeout(timeout time.Duration) *SubtenantGroupListUsingGETParams {
	var ()
	return &SubtenantGroupListUsingGETParams{

		timeout: timeout,
	}
}

// NewSubtenantGroupListUsingGETParamsWithContext creates a new SubtenantGroupListUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubtenantGroupListUsingGETParamsWithContext(ctx context.Context) *SubtenantGroupListUsingGETParams {
	var ()
	return &SubtenantGroupListUsingGETParams{

		Context: ctx,
	}
}

// NewSubtenantGroupListUsingGETParamsWithHTTPClient creates a new SubtenantGroupListUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubtenantGroupListUsingGETParamsWithHTTPClient(client *http.Client) *SubtenantGroupListUsingGETParams {
	var ()
	return &SubtenantGroupListUsingGETParams{
		HTTPClient: client,
	}
}

/*SubtenantGroupListUsingGETParams contains all the parameters to send to the API endpoint
for the subtenant group list using g e t operation typically these are written to a http.Request
*/
type SubtenantGroupListUsingGETParams struct {

	/*DisplayName
	  displayName

	*/
	DisplayName *string
	/*SubtenantID
	  subtenantId

	*/
	SubtenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the subtenant group list using g e t params
func (o *SubtenantGroupListUsingGETParams) WithTimeout(timeout time.Duration) *SubtenantGroupListUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subtenant group list using g e t params
func (o *SubtenantGroupListUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subtenant group list using g e t params
func (o *SubtenantGroupListUsingGETParams) WithContext(ctx context.Context) *SubtenantGroupListUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subtenant group list using g e t params
func (o *SubtenantGroupListUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subtenant group list using g e t params
func (o *SubtenantGroupListUsingGETParams) WithHTTPClient(client *http.Client) *SubtenantGroupListUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subtenant group list using g e t params
func (o *SubtenantGroupListUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisplayName adds the displayName to the subtenant group list using g e t params
func (o *SubtenantGroupListUsingGETParams) WithDisplayName(displayName *string) *SubtenantGroupListUsingGETParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the subtenant group list using g e t params
func (o *SubtenantGroupListUsingGETParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WithSubtenantID adds the subtenantID to the subtenant group list using g e t params
func (o *SubtenantGroupListUsingGETParams) WithSubtenantID(subtenantID string) *SubtenantGroupListUsingGETParams {
	o.SetSubtenantID(subtenantID)
	return o
}

// SetSubtenantID adds the subtenantId to the subtenant group list using g e t params
func (o *SubtenantGroupListUsingGETParams) SetSubtenantID(subtenantID string) {
	o.SubtenantID = subtenantID
}

// WriteToRequest writes these params to a swagger request
func (o *SubtenantGroupListUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param subtenantId
	if err := r.SetPathParam("subtenantId", o.SubtenantID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
