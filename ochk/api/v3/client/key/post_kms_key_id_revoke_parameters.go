// Code generated by go-swagger; DO NOT EDIT.

package key

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

// NewPostKmsKeyIDRevokeParams creates a new PostKmsKeyIDRevokeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostKmsKeyIDRevokeParams() *PostKmsKeyIDRevokeParams {
	return &PostKmsKeyIDRevokeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostKmsKeyIDRevokeParamsWithTimeout creates a new PostKmsKeyIDRevokeParams object
// with the ability to set a timeout on a request.
func NewPostKmsKeyIDRevokeParamsWithTimeout(timeout time.Duration) *PostKmsKeyIDRevokeParams {
	return &PostKmsKeyIDRevokeParams{
		timeout: timeout,
	}
}

// NewPostKmsKeyIDRevokeParamsWithContext creates a new PostKmsKeyIDRevokeParams object
// with the ability to set a context for a request.
func NewPostKmsKeyIDRevokeParamsWithContext(ctx context.Context) *PostKmsKeyIDRevokeParams {
	return &PostKmsKeyIDRevokeParams{
		Context: ctx,
	}
}

// NewPostKmsKeyIDRevokeParamsWithHTTPClient creates a new PostKmsKeyIDRevokeParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostKmsKeyIDRevokeParamsWithHTTPClient(client *http.Client) *PostKmsKeyIDRevokeParams {
	return &PostKmsKeyIDRevokeParams{
		HTTPClient: client,
	}
}

/*
PostKmsKeyIDRevokeParams contains all the parameters to send to the API endpoint

	for the post kms key ID revoke operation.

	Typically these are written to a http.Request.
*/
type PostKmsKeyIDRevokeParams struct {

	// Body.
	Body *models.KeyRevocation

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post kms key ID revoke params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostKmsKeyIDRevokeParams) WithDefaults() *PostKmsKeyIDRevokeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post kms key ID revoke params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostKmsKeyIDRevokeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post kms key ID revoke params
func (o *PostKmsKeyIDRevokeParams) WithTimeout(timeout time.Duration) *PostKmsKeyIDRevokeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post kms key ID revoke params
func (o *PostKmsKeyIDRevokeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post kms key ID revoke params
func (o *PostKmsKeyIDRevokeParams) WithContext(ctx context.Context) *PostKmsKeyIDRevokeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post kms key ID revoke params
func (o *PostKmsKeyIDRevokeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post kms key ID revoke params
func (o *PostKmsKeyIDRevokeParams) WithHTTPClient(client *http.Client) *PostKmsKeyIDRevokeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post kms key ID revoke params
func (o *PostKmsKeyIDRevokeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post kms key ID revoke params
func (o *PostKmsKeyIDRevokeParams) WithBody(body *models.KeyRevocation) *PostKmsKeyIDRevokeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post kms key ID revoke params
func (o *PostKmsKeyIDRevokeParams) SetBody(body *models.KeyRevocation) {
	o.Body = body
}

// WithID adds the id to the post kms key ID revoke params
func (o *PostKmsKeyIDRevokeParams) WithID(id string) *PostKmsKeyIDRevokeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post kms key ID revoke params
func (o *PostKmsKeyIDRevokeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostKmsKeyIDRevokeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
