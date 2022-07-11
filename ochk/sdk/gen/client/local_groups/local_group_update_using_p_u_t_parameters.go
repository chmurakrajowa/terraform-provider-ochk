// Code generated by go-swagger; DO NOT EDIT.

package local_groups

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

// NewLocalGroupUpdateUsingPUTParams creates a new LocalGroupUpdateUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocalGroupUpdateUsingPUTParams() *LocalGroupUpdateUsingPUTParams {
	return &LocalGroupUpdateUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocalGroupUpdateUsingPUTParamsWithTimeout creates a new LocalGroupUpdateUsingPUTParams object
// with the ability to set a timeout on a request.
func NewLocalGroupUpdateUsingPUTParamsWithTimeout(timeout time.Duration) *LocalGroupUpdateUsingPUTParams {
	return &LocalGroupUpdateUsingPUTParams{
		timeout: timeout,
	}
}

// NewLocalGroupUpdateUsingPUTParamsWithContext creates a new LocalGroupUpdateUsingPUTParams object
// with the ability to set a context for a request.
func NewLocalGroupUpdateUsingPUTParamsWithContext(ctx context.Context) *LocalGroupUpdateUsingPUTParams {
	return &LocalGroupUpdateUsingPUTParams{
		Context: ctx,
	}
}

// NewLocalGroupUpdateUsingPUTParamsWithHTTPClient creates a new LocalGroupUpdateUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocalGroupUpdateUsingPUTParamsWithHTTPClient(client *http.Client) *LocalGroupUpdateUsingPUTParams {
	return &LocalGroupUpdateUsingPUTParams{
		HTTPClient: client,
	}
}

/* LocalGroupUpdateUsingPUTParams contains all the parameters to send to the API endpoint
   for the local group update using p u t operation.

   Typically these are written to a http.Request.
*/
type LocalGroupUpdateUsingPUTParams struct {

	/* GroupID.

	   groupId
	*/
	GroupID string

	/* LocalGroup.

	   localGroup
	*/
	LocalGroup *models.LocalGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the local group update using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocalGroupUpdateUsingPUTParams) WithDefaults() *LocalGroupUpdateUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the local group update using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocalGroupUpdateUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the local group update using p u t params
func (o *LocalGroupUpdateUsingPUTParams) WithTimeout(timeout time.Duration) *LocalGroupUpdateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the local group update using p u t params
func (o *LocalGroupUpdateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the local group update using p u t params
func (o *LocalGroupUpdateUsingPUTParams) WithContext(ctx context.Context) *LocalGroupUpdateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the local group update using p u t params
func (o *LocalGroupUpdateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the local group update using p u t params
func (o *LocalGroupUpdateUsingPUTParams) WithHTTPClient(client *http.Client) *LocalGroupUpdateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the local group update using p u t params
func (o *LocalGroupUpdateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the local group update using p u t params
func (o *LocalGroupUpdateUsingPUTParams) WithGroupID(groupID string) *LocalGroupUpdateUsingPUTParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the local group update using p u t params
func (o *LocalGroupUpdateUsingPUTParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithLocalGroup adds the localGroup to the local group update using p u t params
func (o *LocalGroupUpdateUsingPUTParams) WithLocalGroup(localGroup *models.LocalGroup) *LocalGroupUpdateUsingPUTParams {
	o.SetLocalGroup(localGroup)
	return o
}

// SetLocalGroup adds the localGroup to the local group update using p u t params
func (o *LocalGroupUpdateUsingPUTParams) SetLocalGroup(localGroup *models.LocalGroup) {
	o.LocalGroup = localGroup
}

// WriteToRequest writes these params to a swagger request
func (o *LocalGroupUpdateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
	}
	if o.LocalGroup != nil {
		if err := r.SetBodyParam(o.LocalGroup); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
