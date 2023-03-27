// Code generated by go-swagger; DO NOT EDIT.

package virtual_machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// VcsVirtualMachineSnapshotRevertUsingPOSTReader is a Reader for the VcsVirtualMachineSnapshotRevertUsingPOST structure.
type VcsVirtualMachineSnapshotRevertUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VcsVirtualMachineSnapshotRevertUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVcsVirtualMachineSnapshotRevertUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVcsVirtualMachineSnapshotRevertUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVcsVirtualMachineSnapshotRevertUsingPOSTOK creates a VcsVirtualMachineSnapshotRevertUsingPOSTOK with default headers values
func NewVcsVirtualMachineSnapshotRevertUsingPOSTOK() *VcsVirtualMachineSnapshotRevertUsingPOSTOK {
	return &VcsVirtualMachineSnapshotRevertUsingPOSTOK{}
}

/*
VcsVirtualMachineSnapshotRevertUsingPOSTOK describes a response with status code 200, with default header values.

Request has been completed successfully
*/
type VcsVirtualMachineSnapshotRevertUsingPOSTOK struct {
	Payload *models.RevertSnapshotResponse
}

// IsSuccess returns true when this vcs virtual machine snapshot revert using p o s t o k response has a 2xx status code
func (o *VcsVirtualMachineSnapshotRevertUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vcs virtual machine snapshot revert using p o s t o k response has a 3xx status code
func (o *VcsVirtualMachineSnapshotRevertUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vcs virtual machine snapshot revert using p o s t o k response has a 4xx status code
func (o *VcsVirtualMachineSnapshotRevertUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vcs virtual machine snapshot revert using p o s t o k response has a 5xx status code
func (o *VcsVirtualMachineSnapshotRevertUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vcs virtual machine snapshot revert using p o s t o k response a status code equal to that given
func (o *VcsVirtualMachineSnapshotRevertUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vcs virtual machine snapshot revert using p o s t o k response
func (o *VcsVirtualMachineSnapshotRevertUsingPOSTOK) Code() int {
	return 200
}

func (o *VcsVirtualMachineSnapshotRevertUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /vcs/virtual-machines/{virtualMachineId}/snapshots/{snapshotId}/revert][%d] vcsVirtualMachineSnapshotRevertUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *VcsVirtualMachineSnapshotRevertUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /vcs/virtual-machines/{virtualMachineId}/snapshots/{snapshotId}/revert][%d] vcsVirtualMachineSnapshotRevertUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *VcsVirtualMachineSnapshotRevertUsingPOSTOK) GetPayload() *models.RevertSnapshotResponse {
	return o.Payload
}

func (o *VcsVirtualMachineSnapshotRevertUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RevertSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVcsVirtualMachineSnapshotRevertUsingPOSTBadRequest creates a VcsVirtualMachineSnapshotRevertUsingPOSTBadRequest with default headers values
func NewVcsVirtualMachineSnapshotRevertUsingPOSTBadRequest() *VcsVirtualMachineSnapshotRevertUsingPOSTBadRequest {
	return &VcsVirtualMachineSnapshotRevertUsingPOSTBadRequest{}
}

/*
VcsVirtualMachineSnapshotRevertUsingPOSTBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type VcsVirtualMachineSnapshotRevertUsingPOSTBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this vcs virtual machine snapshot revert using p o s t bad request response has a 2xx status code
func (o *VcsVirtualMachineSnapshotRevertUsingPOSTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this vcs virtual machine snapshot revert using p o s t bad request response has a 3xx status code
func (o *VcsVirtualMachineSnapshotRevertUsingPOSTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vcs virtual machine snapshot revert using p o s t bad request response has a 4xx status code
func (o *VcsVirtualMachineSnapshotRevertUsingPOSTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this vcs virtual machine snapshot revert using p o s t bad request response has a 5xx status code
func (o *VcsVirtualMachineSnapshotRevertUsingPOSTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this vcs virtual machine snapshot revert using p o s t bad request response a status code equal to that given
func (o *VcsVirtualMachineSnapshotRevertUsingPOSTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the vcs virtual machine snapshot revert using p o s t bad request response
func (o *VcsVirtualMachineSnapshotRevertUsingPOSTBadRequest) Code() int {
	return 400
}

func (o *VcsVirtualMachineSnapshotRevertUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /vcs/virtual-machines/{virtualMachineId}/snapshots/{snapshotId}/revert][%d] vcsVirtualMachineSnapshotRevertUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *VcsVirtualMachineSnapshotRevertUsingPOSTBadRequest) String() string {
	return fmt.Sprintf("[POST /vcs/virtual-machines/{virtualMachineId}/snapshots/{snapshotId}/revert][%d] vcsVirtualMachineSnapshotRevertUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *VcsVirtualMachineSnapshotRevertUsingPOSTBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *VcsVirtualMachineSnapshotRevertUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
