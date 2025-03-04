// Code generated by go-swagger; DO NOT EDIT.

package virtual_machine_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
)

// PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsReader is a Reader for the PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshots structure.
type PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /vcs/virtual-machines/{virtualMachineId}/removeAllSnapshots] PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshots", response, response.Code())
	}
}

// NewPostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsOK creates a PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsOK with default headers values
func NewPostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsOK() *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsOK {
	return &PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsOK{}
}

/*
PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsOK describes a response with status code 200, with default header values.

OK
*/
type PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsOK struct {
	Payload *models.RemoveAllSnapshotsResponse
}

// IsSuccess returns true when this post vcs virtual machines virtual machine Id remove all snapshots o k response has a 2xx status code
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post vcs virtual machines virtual machine Id remove all snapshots o k response has a 3xx status code
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post vcs virtual machines virtual machine Id remove all snapshots o k response has a 4xx status code
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post vcs virtual machines virtual machine Id remove all snapshots o k response has a 5xx status code
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post vcs virtual machines virtual machine Id remove all snapshots o k response a status code equal to that given
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post vcs virtual machines virtual machine Id remove all snapshots o k response
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsOK) Code() int {
	return 200
}

func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /vcs/virtual-machines/{virtualMachineId}/removeAllSnapshots][%d] postVcsVirtualMachinesVirtualMachineIdRemoveAllSnapshotsOK %s", 200, payload)
}

func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /vcs/virtual-machines/{virtualMachineId}/removeAllSnapshots][%d] postVcsVirtualMachinesVirtualMachineIdRemoveAllSnapshotsOK %s", 200, payload)
}

func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsOK) GetPayload() *models.RemoveAllSnapshotsResponse {
	return o.Payload
}

func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemoveAllSnapshotsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsBadRequest creates a PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsBadRequest with default headers values
func NewPostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsBadRequest() *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsBadRequest {
	return &PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsBadRequest{}
}

/*
PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this post vcs virtual machines virtual machine Id remove all snapshots bad request response has a 2xx status code
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post vcs virtual machines virtual machine Id remove all snapshots bad request response has a 3xx status code
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post vcs virtual machines virtual machine Id remove all snapshots bad request response has a 4xx status code
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post vcs virtual machines virtual machine Id remove all snapshots bad request response has a 5xx status code
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post vcs virtual machines virtual machine Id remove all snapshots bad request response a status code equal to that given
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post vcs virtual machines virtual machine Id remove all snapshots bad request response
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsBadRequest) Code() int {
	return 400
}

func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /vcs/virtual-machines/{virtualMachineId}/removeAllSnapshots][%d] postVcsVirtualMachinesVirtualMachineIdRemoveAllSnapshotsBadRequest %s", 400, payload)
}

func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /vcs/virtual-machines/{virtualMachineId}/removeAllSnapshots][%d] postVcsVirtualMachinesVirtualMachineIdRemoveAllSnapshotsBadRequest %s", 400, payload)
}

func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsUnauthorized creates a PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsUnauthorized with default headers values
func NewPostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsUnauthorized() *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsUnauthorized {
	return &PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsUnauthorized{}
}

/*
PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this post vcs virtual machines virtual machine Id remove all snapshots unauthorized response has a 2xx status code
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post vcs virtual machines virtual machine Id remove all snapshots unauthorized response has a 3xx status code
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post vcs virtual machines virtual machine Id remove all snapshots unauthorized response has a 4xx status code
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post vcs virtual machines virtual machine Id remove all snapshots unauthorized response has a 5xx status code
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post vcs virtual machines virtual machine Id remove all snapshots unauthorized response a status code equal to that given
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post vcs virtual machines virtual machine Id remove all snapshots unauthorized response
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsUnauthorized) Code() int {
	return 401
}

func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /vcs/virtual-machines/{virtualMachineId}/removeAllSnapshots][%d] postVcsVirtualMachinesVirtualMachineIdRemoveAllSnapshotsUnauthorized %s", 401, payload)
}

func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /vcs/virtual-machines/{virtualMachineId}/removeAllSnapshots][%d] postVcsVirtualMachinesVirtualMachineIdRemoveAllSnapshotsUnauthorized %s", 401, payload)
}

func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsForbidden creates a PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsForbidden with default headers values
func NewPostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsForbidden() *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsForbidden {
	return &PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsForbidden{}
}

/*
PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this post vcs virtual machines virtual machine Id remove all snapshots forbidden response has a 2xx status code
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post vcs virtual machines virtual machine Id remove all snapshots forbidden response has a 3xx status code
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post vcs virtual machines virtual machine Id remove all snapshots forbidden response has a 4xx status code
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post vcs virtual machines virtual machine Id remove all snapshots forbidden response has a 5xx status code
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post vcs virtual machines virtual machine Id remove all snapshots forbidden response a status code equal to that given
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post vcs virtual machines virtual machine Id remove all snapshots forbidden response
func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsForbidden) Code() int {
	return 403
}

func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /vcs/virtual-machines/{virtualMachineId}/removeAllSnapshots][%d] postVcsVirtualMachinesVirtualMachineIdRemoveAllSnapshotsForbidden %s", 403, payload)
}

func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /vcs/virtual-machines/{virtualMachineId}/removeAllSnapshots][%d] postVcsVirtualMachinesVirtualMachineIdRemoveAllSnapshotsForbidden %s", 403, payload)
}

func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PostVcsVirtualMachinesVirtualMachineIDRemoveAllSnapshotsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
