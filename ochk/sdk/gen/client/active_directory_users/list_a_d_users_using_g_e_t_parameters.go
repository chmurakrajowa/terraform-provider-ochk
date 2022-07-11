// Code generated by go-swagger; DO NOT EDIT.

package active_directory_users

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

// NewListADUsersUsingGETParams creates a new ListADUsersUsingGETParams object
// with the default values initialized.
func NewListADUsersUsingGETParams() *ListADUsersUsingGETParams {

	return &ListADUsersUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListADUsersUsingGETParamsWithTimeout creates a new ListADUsersUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListADUsersUsingGETParamsWithTimeout(timeout time.Duration) *ListADUsersUsingGETParams {

	return &ListADUsersUsingGETParams{

		timeout: timeout,
	}
}

// NewListADUsersUsingGETParamsWithContext creates a new ListADUsersUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewListADUsersUsingGETParamsWithContext(ctx context.Context) *ListADUsersUsingGETParams {

	return &ListADUsersUsingGETParams{

		Context: ctx,
	}
}

// NewListADUsersUsingGETParamsWithHTTPClient creates a new ListADUsersUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListADUsersUsingGETParamsWithHTTPClient(client *http.Client) *ListADUsersUsingGETParams {

	return &ListADUsersUsingGETParams{
		HTTPClient: client,
	}
}

/*ListADUsersUsingGETParams contains all the parameters to send to the API endpoint
for the list a d users using g e t operation typically these are written to a http.Request
*/
type ListADUsersUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list a d users using g e t params
func (o *ListADUsersUsingGETParams) WithTimeout(timeout time.Duration) *ListADUsersUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list a d users using g e t params
func (o *ListADUsersUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list a d users using g e t params
func (o *ListADUsersUsingGETParams) WithContext(ctx context.Context) *ListADUsersUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list a d users using g e t params
func (o *ListADUsersUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list a d users using g e t params
func (o *ListADUsersUsingGETParams) WithHTTPClient(client *http.Client) *ListADUsersUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list a d users using g e t params
func (o *ListADUsersUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListADUsersUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}