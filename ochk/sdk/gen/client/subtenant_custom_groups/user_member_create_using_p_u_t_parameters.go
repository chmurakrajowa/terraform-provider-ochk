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

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// NewUserMemberCreateUsingPUTParams creates a new UserMemberCreateUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserMemberCreateUsingPUTParams() *UserMemberCreateUsingPUTParams {
	return &UserMemberCreateUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserMemberCreateUsingPUTParamsWithTimeout creates a new UserMemberCreateUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUserMemberCreateUsingPUTParamsWithTimeout(timeout time.Duration) *UserMemberCreateUsingPUTParams {
	return &UserMemberCreateUsingPUTParams{
		timeout: timeout,
	}
}

// NewUserMemberCreateUsingPUTParamsWithContext creates a new UserMemberCreateUsingPUTParams object
// with the ability to set a context for a request.
func NewUserMemberCreateUsingPUTParamsWithContext(ctx context.Context) *UserMemberCreateUsingPUTParams {
	return &UserMemberCreateUsingPUTParams{
		Context: ctx,
	}
}

// NewUserMemberCreateUsingPUTParamsWithHTTPClient creates a new UserMemberCreateUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserMemberCreateUsingPUTParamsWithHTTPClient(client *http.Client) *UserMemberCreateUsingPUTParams {
	return &UserMemberCreateUsingPUTParams{
		HTTPClient: client,
	}
}

/* UserMemberCreateUsingPUTParams contains all the parameters to send to the API endpoint
   for the user member create using p u t operation.

   Typically these are written to a http.Request.
*/
type UserMemberCreateUsingPUTParams struct {

	/* ParentGroupID.

	   parentGroupId
	*/
	ParentGroupID string

	/* SubtenantID.

	   subtenantId
	*/
	SubtenantID string

	/* UserInstance.

	   userInstance
	*/
	UserInstance *models.UserInstance

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user member create using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserMemberCreateUsingPUTParams) WithDefaults() *UserMemberCreateUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user member create using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserMemberCreateUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user member create using p u t params
func (o *UserMemberCreateUsingPUTParams) WithTimeout(timeout time.Duration) *UserMemberCreateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user member create using p u t params
func (o *UserMemberCreateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user member create using p u t params
func (o *UserMemberCreateUsingPUTParams) WithContext(ctx context.Context) *UserMemberCreateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user member create using p u t params
func (o *UserMemberCreateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user member create using p u t params
func (o *UserMemberCreateUsingPUTParams) WithHTTPClient(client *http.Client) *UserMemberCreateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user member create using p u t params
func (o *UserMemberCreateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParentGroupID adds the parentGroupID to the user member create using p u t params
func (o *UserMemberCreateUsingPUTParams) WithParentGroupID(parentGroupID string) *UserMemberCreateUsingPUTParams {
	o.SetParentGroupID(parentGroupID)
	return o
}

// SetParentGroupID adds the parentGroupId to the user member create using p u t params
func (o *UserMemberCreateUsingPUTParams) SetParentGroupID(parentGroupID string) {
	o.ParentGroupID = parentGroupID
}

// WithSubtenantID adds the subtenantID to the user member create using p u t params
func (o *UserMemberCreateUsingPUTParams) WithSubtenantID(subtenantID string) *UserMemberCreateUsingPUTParams {
	o.SetSubtenantID(subtenantID)
	return o
}

// SetSubtenantID adds the subtenantId to the user member create using p u t params
func (o *UserMemberCreateUsingPUTParams) SetSubtenantID(subtenantID string) {
	o.SubtenantID = subtenantID
}

// WithUserInstance adds the userInstance to the user member create using p u t params
func (o *UserMemberCreateUsingPUTParams) WithUserInstance(userInstance *models.UserInstance) *UserMemberCreateUsingPUTParams {
	o.SetUserInstance(userInstance)
	return o
}

// SetUserInstance adds the userInstance to the user member create using p u t params
func (o *UserMemberCreateUsingPUTParams) SetUserInstance(userInstance *models.UserInstance) {
	o.UserInstance = userInstance
}

// WriteToRequest writes these params to a swagger request
func (o *UserMemberCreateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param parentGroupId
	if err := r.SetPathParam("parentGroupId", o.ParentGroupID); err != nil {
		return err
	}

	// path param subtenantId
	if err := r.SetPathParam("subtenantId", o.SubtenantID); err != nil {
		return err
	}
	if o.UserInstance != nil {
		if err := r.SetBodyParam(o.UserInstance); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
