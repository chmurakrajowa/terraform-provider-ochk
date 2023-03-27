// Code generated by go-swagger; DO NOT EDIT.

package m_c_s_secret

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

// NewMcsGetSecretUsingGETParams creates a new McsGetSecretUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMcsGetSecretUsingGETParams() *McsGetSecretUsingGETParams {
	return &McsGetSecretUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMcsGetSecretUsingGETParamsWithTimeout creates a new McsGetSecretUsingGETParams object
// with the ability to set a timeout on a request.
func NewMcsGetSecretUsingGETParamsWithTimeout(timeout time.Duration) *McsGetSecretUsingGETParams {
	return &McsGetSecretUsingGETParams{
		timeout: timeout,
	}
}

// NewMcsGetSecretUsingGETParamsWithContext creates a new McsGetSecretUsingGETParams object
// with the ability to set a context for a request.
func NewMcsGetSecretUsingGETParamsWithContext(ctx context.Context) *McsGetSecretUsingGETParams {
	return &McsGetSecretUsingGETParams{
		Context: ctx,
	}
}

// NewMcsGetSecretUsingGETParamsWithHTTPClient creates a new McsGetSecretUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewMcsGetSecretUsingGETParamsWithHTTPClient(client *http.Client) *McsGetSecretUsingGETParams {
	return &McsGetSecretUsingGETParams{
		HTTPClient: client,
	}
}

/*
McsGetSecretUsingGETParams contains all the parameters to send to the API endpoint

	for the mcs get secret using g e t operation.

	Typically these are written to a http.Request.
*/
type McsGetSecretUsingGETParams struct {

	/* Username.

	   username
	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the mcs get secret using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *McsGetSecretUsingGETParams) WithDefaults() *McsGetSecretUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the mcs get secret using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *McsGetSecretUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the mcs get secret using g e t params
func (o *McsGetSecretUsingGETParams) WithTimeout(timeout time.Duration) *McsGetSecretUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mcs get secret using g e t params
func (o *McsGetSecretUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mcs get secret using g e t params
func (o *McsGetSecretUsingGETParams) WithContext(ctx context.Context) *McsGetSecretUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mcs get secret using g e t params
func (o *McsGetSecretUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mcs get secret using g e t params
func (o *McsGetSecretUsingGETParams) WithHTTPClient(client *http.Client) *McsGetSecretUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mcs get secret using g e t params
func (o *McsGetSecretUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the mcs get secret using g e t params
func (o *McsGetSecretUsingGETParams) WithUsername(username string) *McsGetSecretUsingGETParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the mcs get secret using g e t params
func (o *McsGetSecretUsingGETParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *McsGetSecretUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param username
	qrUsername := o.Username
	qUsername := qrUsername
	if qUsername != "" {

		if err := r.SetQueryParam("username", qUsername); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
