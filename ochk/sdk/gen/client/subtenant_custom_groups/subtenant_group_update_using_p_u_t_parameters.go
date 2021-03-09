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

// NewSubtenantGroupUpdateUsingPUTParams creates a new SubtenantGroupUpdateUsingPUTParams object
// with the default values initialized.
func NewSubtenantGroupUpdateUsingPUTParams() *SubtenantGroupUpdateUsingPUTParams {
	var ()
	return &SubtenantGroupUpdateUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubtenantGroupUpdateUsingPUTParamsWithTimeout creates a new SubtenantGroupUpdateUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubtenantGroupUpdateUsingPUTParamsWithTimeout(timeout time.Duration) *SubtenantGroupUpdateUsingPUTParams {
	var ()
	return &SubtenantGroupUpdateUsingPUTParams{

		timeout: timeout,
	}
}

// NewSubtenantGroupUpdateUsingPUTParamsWithContext creates a new SubtenantGroupUpdateUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubtenantGroupUpdateUsingPUTParamsWithContext(ctx context.Context) *SubtenantGroupUpdateUsingPUTParams {
	var ()
	return &SubtenantGroupUpdateUsingPUTParams{

		Context: ctx,
	}
}

// NewSubtenantGroupUpdateUsingPUTParamsWithHTTPClient creates a new SubtenantGroupUpdateUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubtenantGroupUpdateUsingPUTParamsWithHTTPClient(client *http.Client) *SubtenantGroupUpdateUsingPUTParams {
	var ()
	return &SubtenantGroupUpdateUsingPUTParams{
		HTTPClient: client,
	}
}

/*SubtenantGroupUpdateUsingPUTParams contains all the parameters to send to the API endpoint
for the subtenant group update using p u t operation typically these are written to a http.Request
*/
type SubtenantGroupUpdateUsingPUTParams struct {

	/*GroupID
	  groupId

	*/
	GroupID string
	/*GroupInstance
	  groupInstance

	*/
	GroupInstance *models.GroupInstance
	/*SubtenantID
	  subtenantId

	*/
	SubtenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the subtenant group update using p u t params
func (o *SubtenantGroupUpdateUsingPUTParams) WithTimeout(timeout time.Duration) *SubtenantGroupUpdateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subtenant group update using p u t params
func (o *SubtenantGroupUpdateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subtenant group update using p u t params
func (o *SubtenantGroupUpdateUsingPUTParams) WithContext(ctx context.Context) *SubtenantGroupUpdateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subtenant group update using p u t params
func (o *SubtenantGroupUpdateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subtenant group update using p u t params
func (o *SubtenantGroupUpdateUsingPUTParams) WithHTTPClient(client *http.Client) *SubtenantGroupUpdateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subtenant group update using p u t params
func (o *SubtenantGroupUpdateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the subtenant group update using p u t params
func (o *SubtenantGroupUpdateUsingPUTParams) WithGroupID(groupID string) *SubtenantGroupUpdateUsingPUTParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the subtenant group update using p u t params
func (o *SubtenantGroupUpdateUsingPUTParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithGroupInstance adds the groupInstance to the subtenant group update using p u t params
func (o *SubtenantGroupUpdateUsingPUTParams) WithGroupInstance(groupInstance *models.GroupInstance) *SubtenantGroupUpdateUsingPUTParams {
	o.SetGroupInstance(groupInstance)
	return o
}

// SetGroupInstance adds the groupInstance to the subtenant group update using p u t params
func (o *SubtenantGroupUpdateUsingPUTParams) SetGroupInstance(groupInstance *models.GroupInstance) {
	o.GroupInstance = groupInstance
}

// WithSubtenantID adds the subtenantID to the subtenant group update using p u t params
func (o *SubtenantGroupUpdateUsingPUTParams) WithSubtenantID(subtenantID string) *SubtenantGroupUpdateUsingPUTParams {
	o.SetSubtenantID(subtenantID)
	return o
}

// SetSubtenantID adds the subtenantId to the subtenant group update using p u t params
func (o *SubtenantGroupUpdateUsingPUTParams) SetSubtenantID(subtenantID string) {
	o.SubtenantID = subtenantID
}

// WriteToRequest writes these params to a swagger request
func (o *SubtenantGroupUpdateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
	}

	if o.GroupInstance != nil {
		if err := r.SetBodyParam(o.GroupInstance); err != nil {
			return err
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