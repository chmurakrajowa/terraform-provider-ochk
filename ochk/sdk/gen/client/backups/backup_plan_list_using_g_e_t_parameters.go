// Code generated by go-swagger; DO NOT EDIT.

package backups

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

// NewBackupPlanListUsingGETParams creates a new BackupPlanListUsingGETParams object
// with the default values initialized.
func NewBackupPlanListUsingGETParams() *BackupPlanListUsingGETParams {
	var ()
	return &BackupPlanListUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBackupPlanListUsingGETParamsWithTimeout creates a new BackupPlanListUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBackupPlanListUsingGETParamsWithTimeout(timeout time.Duration) *BackupPlanListUsingGETParams {
	var ()
	return &BackupPlanListUsingGETParams{

		timeout: timeout,
	}
}

// NewBackupPlanListUsingGETParamsWithContext creates a new BackupPlanListUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewBackupPlanListUsingGETParamsWithContext(ctx context.Context) *BackupPlanListUsingGETParams {
	var ()
	return &BackupPlanListUsingGETParams{

		Context: ctx,
	}
}

// NewBackupPlanListUsingGETParamsWithHTTPClient creates a new BackupPlanListUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBackupPlanListUsingGETParamsWithHTTPClient(client *http.Client) *BackupPlanListUsingGETParams {
	var ()
	return &BackupPlanListUsingGETParams{
		HTTPClient: client,
	}
}

/*BackupPlanListUsingGETParams contains all the parameters to send to the API endpoint
for the backup plan list using g e t operation typically these are written to a http.Request
*/
type BackupPlanListUsingGETParams struct {

	/*BackupPlanName
	  backupPlanName

	*/
	BackupPlanName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the backup plan list using g e t params
func (o *BackupPlanListUsingGETParams) WithTimeout(timeout time.Duration) *BackupPlanListUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backup plan list using g e t params
func (o *BackupPlanListUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backup plan list using g e t params
func (o *BackupPlanListUsingGETParams) WithContext(ctx context.Context) *BackupPlanListUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backup plan list using g e t params
func (o *BackupPlanListUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backup plan list using g e t params
func (o *BackupPlanListUsingGETParams) WithHTTPClient(client *http.Client) *BackupPlanListUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backup plan list using g e t params
func (o *BackupPlanListUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackupPlanName adds the backupPlanName to the backup plan list using g e t params
func (o *BackupPlanListUsingGETParams) WithBackupPlanName(backupPlanName *string) *BackupPlanListUsingGETParams {
	o.SetBackupPlanName(backupPlanName)
	return o
}

// SetBackupPlanName adds the backupPlanName to the backup plan list using g e t params
func (o *BackupPlanListUsingGETParams) SetBackupPlanName(backupPlanName *string) {
	o.BackupPlanName = backupPlanName
}

// WriteToRequest writes these params to a swagger request
func (o *BackupPlanListUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BackupPlanName != nil {

		// query param backupPlanName
		var qrBackupPlanName string
		if o.BackupPlanName != nil {
			qrBackupPlanName = *o.BackupPlanName
		}
		qBackupPlanName := qrBackupPlanName
		if qBackupPlanName != "" {
			if err := r.SetQueryParam("backupPlanName", qBackupPlanName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
