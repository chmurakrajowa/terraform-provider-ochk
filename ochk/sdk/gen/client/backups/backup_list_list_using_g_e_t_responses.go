// Code generated by go-swagger; DO NOT EDIT.

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// BackupListListUsingGETReader is a Reader for the BackupListListUsingGET structure.
type BackupListListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupListListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupListListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBackupListListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /backups/plans/{backupPlanId}/lists] backupListListUsingGET", response, response.Code())
	}
}

// NewBackupListListUsingGETOK creates a BackupListListUsingGETOK with default headers values
func NewBackupListListUsingGETOK() *BackupListListUsingGETOK {
	return &BackupListListUsingGETOK{}
}

/*
BackupListListUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type BackupListListUsingGETOK struct {
	Payload *models.BackupListListResponse
}

// IsSuccess returns true when this backup list list using g e t o k response has a 2xx status code
func (o *BackupListListUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backup list list using g e t o k response has a 3xx status code
func (o *BackupListListUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup list list using g e t o k response has a 4xx status code
func (o *BackupListListUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup list list using g e t o k response has a 5xx status code
func (o *BackupListListUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this backup list list using g e t o k response a status code equal to that given
func (o *BackupListListUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the backup list list using g e t o k response
func (o *BackupListListUsingGETOK) Code() int {
	return 200
}

func (o *BackupListListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /backups/plans/{backupPlanId}/lists][%d] backupListListUsingGETOK  %+v", 200, o.Payload)
}

func (o *BackupListListUsingGETOK) String() string {
	return fmt.Sprintf("[GET /backups/plans/{backupPlanId}/lists][%d] backupListListUsingGETOK  %+v", 200, o.Payload)
}

func (o *BackupListListUsingGETOK) GetPayload() *models.BackupListListResponse {
	return o.Payload
}

func (o *BackupListListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupListListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupListListUsingGETBadRequest creates a BackupListListUsingGETBadRequest with default headers values
func NewBackupListListUsingGETBadRequest() *BackupListListUsingGETBadRequest {
	return &BackupListListUsingGETBadRequest{}
}

/*
BackupListListUsingGETBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type BackupListListUsingGETBadRequest struct {
}

// IsSuccess returns true when this backup list list using g e t bad request response has a 2xx status code
func (o *BackupListListUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup list list using g e t bad request response has a 3xx status code
func (o *BackupListListUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup list list using g e t bad request response has a 4xx status code
func (o *BackupListListUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup list list using g e t bad request response has a 5xx status code
func (o *BackupListListUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this backup list list using g e t bad request response a status code equal to that given
func (o *BackupListListUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the backup list list using g e t bad request response
func (o *BackupListListUsingGETBadRequest) Code() int {
	return 400
}

func (o *BackupListListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /backups/plans/{backupPlanId}/lists][%d] backupListListUsingGETBadRequest ", 400)
}

func (o *BackupListListUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /backups/plans/{backupPlanId}/lists][%d] backupListListUsingGETBadRequest ", 400)
}

func (o *BackupListListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
