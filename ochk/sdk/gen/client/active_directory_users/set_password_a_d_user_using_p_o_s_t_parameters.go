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

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// NewSetPasswordADUserUsingPOSTParams creates a new SetPasswordADUserUsingPOSTParams object
// with the default values initialized.
func NewSetPasswordADUserUsingPOSTParams() *SetPasswordADUserUsingPOSTParams {
	var ()
	return &SetPasswordADUserUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetPasswordADUserUsingPOSTParamsWithTimeout creates a new SetPasswordADUserUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetPasswordADUserUsingPOSTParamsWithTimeout(timeout time.Duration) *SetPasswordADUserUsingPOSTParams {
	var ()
	return &SetPasswordADUserUsingPOSTParams{

		timeout: timeout,
	}
}

// NewSetPasswordADUserUsingPOSTParamsWithContext creates a new SetPasswordADUserUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetPasswordADUserUsingPOSTParamsWithContext(ctx context.Context) *SetPasswordADUserUsingPOSTParams {
	var ()
	return &SetPasswordADUserUsingPOSTParams{

		Context: ctx,
	}
}

// NewSetPasswordADUserUsingPOSTParamsWithHTTPClient creates a new SetPasswordADUserUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetPasswordADUserUsingPOSTParamsWithHTTPClient(client *http.Client) *SetPasswordADUserUsingPOSTParams {
	var ()
	return &SetPasswordADUserUsingPOSTParams{
		HTTPClient: client,
	}
}

/*SetPasswordADUserUsingPOSTParams contains all the parameters to send to the API endpoint
for the set password a d user using p o s t operation typically these are written to a http.Request
*/
type SetPasswordADUserUsingPOSTParams struct {

	/*SamAccountName
	  samAccountName

	*/
	SamAccountName string
	/*SetUserInstancePassRequest
	  setUserInstancePassRequest

	*/
	SetUserInstancePassRequest *models.SetUserInstancePassRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set password a d user using p o s t params
func (o *SetPasswordADUserUsingPOSTParams) WithTimeout(timeout time.Duration) *SetPasswordADUserUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set password a d user using p o s t params
func (o *SetPasswordADUserUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set password a d user using p o s t params
func (o *SetPasswordADUserUsingPOSTParams) WithContext(ctx context.Context) *SetPasswordADUserUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set password a d user using p o s t params
func (o *SetPasswordADUserUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set password a d user using p o s t params
func (o *SetPasswordADUserUsingPOSTParams) WithHTTPClient(client *http.Client) *SetPasswordADUserUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set password a d user using p o s t params
func (o *SetPasswordADUserUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSamAccountName adds the samAccountName to the set password a d user using p o s t params
func (o *SetPasswordADUserUsingPOSTParams) WithSamAccountName(samAccountName string) *SetPasswordADUserUsingPOSTParams {
	o.SetSamAccountName(samAccountName)
	return o
}

// SetSamAccountName adds the samAccountName to the set password a d user using p o s t params
func (o *SetPasswordADUserUsingPOSTParams) SetSamAccountName(samAccountName string) {
	o.SamAccountName = samAccountName
}

// WithSetUserInstancePassRequest adds the setUserInstancePassRequest to the set password a d user using p o s t params
func (o *SetPasswordADUserUsingPOSTParams) WithSetUserInstancePassRequest(setUserInstancePassRequest *models.SetUserInstancePassRequest) *SetPasswordADUserUsingPOSTParams {
	o.SetSetUserInstancePassRequest(setUserInstancePassRequest)
	return o
}

// SetSetUserInstancePassRequest adds the setUserInstancePassRequest to the set password a d user using p o s t params
func (o *SetPasswordADUserUsingPOSTParams) SetSetUserInstancePassRequest(setUserInstancePassRequest *models.SetUserInstancePassRequest) {
	o.SetUserInstancePassRequest = setUserInstancePassRequest
}

// WriteToRequest writes these params to a swagger request
func (o *SetPasswordADUserUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param samAccountName
	if err := r.SetPathParam("samAccountName", o.SamAccountName); err != nil {
		return err
	}

	if o.SetUserInstancePassRequest != nil {
		if err := r.SetBodyParam(o.SetUserInstancePassRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}