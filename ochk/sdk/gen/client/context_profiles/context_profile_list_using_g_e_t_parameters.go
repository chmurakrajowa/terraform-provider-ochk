// Code generated by go-swagger; DO NOT EDIT.

package context_profiles

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

// NewContextProfileListUsingGETParams creates a new ContextProfileListUsingGETParams object
// with the default values initialized.
func NewContextProfileListUsingGETParams() *ContextProfileListUsingGETParams {
	var ()
	return &ContextProfileListUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewContextProfileListUsingGETParamsWithTimeout creates a new ContextProfileListUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewContextProfileListUsingGETParamsWithTimeout(timeout time.Duration) *ContextProfileListUsingGETParams {
	var ()
	return &ContextProfileListUsingGETParams{

		timeout: timeout,
	}
}

// NewContextProfileListUsingGETParamsWithContext creates a new ContextProfileListUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewContextProfileListUsingGETParamsWithContext(ctx context.Context) *ContextProfileListUsingGETParams {
	var ()
	return &ContextProfileListUsingGETParams{

		Context: ctx,
	}
}

// NewContextProfileListUsingGETParamsWithHTTPClient creates a new ContextProfileListUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContextProfileListUsingGETParamsWithHTTPClient(client *http.Client) *ContextProfileListUsingGETParams {
	var ()
	return &ContextProfileListUsingGETParams{
		HTTPClient: client,
	}
}

/*ContextProfileListUsingGETParams contains all the parameters to send to the API endpoint
for the context profile list using g e t operation typically these are written to a http.Request
*/
type ContextProfileListUsingGETParams struct {

	/*DisplayName
	  displayName

	*/
	DisplayName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the context profile list using g e t params
func (o *ContextProfileListUsingGETParams) WithTimeout(timeout time.Duration) *ContextProfileListUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the context profile list using g e t params
func (o *ContextProfileListUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the context profile list using g e t params
func (o *ContextProfileListUsingGETParams) WithContext(ctx context.Context) *ContextProfileListUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the context profile list using g e t params
func (o *ContextProfileListUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the context profile list using g e t params
func (o *ContextProfileListUsingGETParams) WithHTTPClient(client *http.Client) *ContextProfileListUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the context profile list using g e t params
func (o *ContextProfileListUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisplayName adds the displayName to the context profile list using g e t params
func (o *ContextProfileListUsingGETParams) WithDisplayName(displayName *string) *ContextProfileListUsingGETParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the context profile list using g e t params
func (o *ContextProfileListUsingGETParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WriteToRequest writes these params to a swagger request
func (o *ContextProfileListUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
