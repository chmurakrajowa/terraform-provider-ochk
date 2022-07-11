// Code generated by go-swagger; DO NOT EDIT.

package k_m_s_key_rotation_scheduler

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

// NewKeyRotationScheduleDeleteUsingDELETEParams creates a new KeyRotationScheduleDeleteUsingDELETEParams object
// with the default values initialized.
func NewKeyRotationScheduleDeleteUsingDELETEParams() *KeyRotationScheduleDeleteUsingDELETEParams {
	var ()
	return &KeyRotationScheduleDeleteUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewKeyRotationScheduleDeleteUsingDELETEParamsWithTimeout creates a new KeyRotationScheduleDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewKeyRotationScheduleDeleteUsingDELETEParamsWithTimeout(timeout time.Duration) *KeyRotationScheduleDeleteUsingDELETEParams {
	var ()
	return &KeyRotationScheduleDeleteUsingDELETEParams{

		timeout: timeout,
	}
}

// NewKeyRotationScheduleDeleteUsingDELETEParamsWithContext creates a new KeyRotationScheduleDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewKeyRotationScheduleDeleteUsingDELETEParamsWithContext(ctx context.Context) *KeyRotationScheduleDeleteUsingDELETEParams {
	var ()
	return &KeyRotationScheduleDeleteUsingDELETEParams{

		Context: ctx,
	}
}

// NewKeyRotationScheduleDeleteUsingDELETEParamsWithHTTPClient creates a new KeyRotationScheduleDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewKeyRotationScheduleDeleteUsingDELETEParamsWithHTTPClient(client *http.Client) *KeyRotationScheduleDeleteUsingDELETEParams {
	var ()
	return &KeyRotationScheduleDeleteUsingDELETEParams{
		HTTPClient: client,
	}
}

/*KeyRotationScheduleDeleteUsingDELETEParams contains all the parameters to send to the API endpoint
for the key rotation schedule delete using d e l e t e operation typically these are written to a http.Request
*/
type KeyRotationScheduleDeleteUsingDELETEParams struct {

	/*KeyID
	  keyId

	*/
	KeyID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the key rotation schedule delete using d e l e t e params
func (o *KeyRotationScheduleDeleteUsingDELETEParams) WithTimeout(timeout time.Duration) *KeyRotationScheduleDeleteUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the key rotation schedule delete using d e l e t e params
func (o *KeyRotationScheduleDeleteUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the key rotation schedule delete using d e l e t e params
func (o *KeyRotationScheduleDeleteUsingDELETEParams) WithContext(ctx context.Context) *KeyRotationScheduleDeleteUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the key rotation schedule delete using d e l e t e params
func (o *KeyRotationScheduleDeleteUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the key rotation schedule delete using d e l e t e params
func (o *KeyRotationScheduleDeleteUsingDELETEParams) WithHTTPClient(client *http.Client) *KeyRotationScheduleDeleteUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the key rotation schedule delete using d e l e t e params
func (o *KeyRotationScheduleDeleteUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeyID adds the keyID to the key rotation schedule delete using d e l e t e params
func (o *KeyRotationScheduleDeleteUsingDELETEParams) WithKeyID(keyID string) *KeyRotationScheduleDeleteUsingDELETEParams {
	o.SetKeyID(keyID)
	return o
}

// SetKeyID adds the keyId to the key rotation schedule delete using d e l e t e params
func (o *KeyRotationScheduleDeleteUsingDELETEParams) SetKeyID(keyID string) {
	o.KeyID = keyID
}

// WriteToRequest writes these params to a swagger request
func (o *KeyRotationScheduleDeleteUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param keyId
	if err := r.SetPathParam("keyId", o.KeyID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
