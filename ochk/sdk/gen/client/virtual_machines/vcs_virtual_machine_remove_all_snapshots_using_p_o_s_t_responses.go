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

// VcsVirtualMachineRemoveAllSnapshotsUsingPOSTReader is a Reader for the VcsVirtualMachineRemoveAllSnapshotsUsingPOST structure.
type VcsVirtualMachineRemoveAllSnapshotsUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVcsVirtualMachineRemoveAllSnapshotsUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVcsVirtualMachineRemoveAllSnapshotsUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVcsVirtualMachineRemoveAllSnapshotsUsingPOSTOK creates a VcsVirtualMachineRemoveAllSnapshotsUsingPOSTOK with default headers values
func NewVcsVirtualMachineRemoveAllSnapshotsUsingPOSTOK() *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTOK {
	return &VcsVirtualMachineRemoveAllSnapshotsUsingPOSTOK{}
}

/*
VcsVirtualMachineRemoveAllSnapshotsUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type VcsVirtualMachineRemoveAllSnapshotsUsingPOSTOK struct {
	Payload *models.RemoveAllSnapshotsResponse
}

// IsSuccess returns true when this vcs virtual machine remove all snapshots using p o s t o k response has a 2xx status code
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vcs virtual machine remove all snapshots using p o s t o k response has a 3xx status code
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vcs virtual machine remove all snapshots using p o s t o k response has a 4xx status code
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vcs virtual machine remove all snapshots using p o s t o k response has a 5xx status code
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vcs virtual machine remove all snapshots using p o s t o k response a status code equal to that given
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vcs virtual machine remove all snapshots using p o s t o k response
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTOK) Code() int {
	return 200
}

func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /vcs/virtual-machines/{virtualMachineId}/removeAllSnapshots][%d] vcsVirtualMachineRemoveAllSnapshotsUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /vcs/virtual-machines/{virtualMachineId}/removeAllSnapshots][%d] vcsVirtualMachineRemoveAllSnapshotsUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTOK) GetPayload() *models.RemoveAllSnapshotsResponse {
	return o.Payload
}

func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemoveAllSnapshotsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVcsVirtualMachineRemoveAllSnapshotsUsingPOSTBadRequest creates a VcsVirtualMachineRemoveAllSnapshotsUsingPOSTBadRequest with default headers values
func NewVcsVirtualMachineRemoveAllSnapshotsUsingPOSTBadRequest() *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTBadRequest {
	return &VcsVirtualMachineRemoveAllSnapshotsUsingPOSTBadRequest{}
}

/*
VcsVirtualMachineRemoveAllSnapshotsUsingPOSTBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type VcsVirtualMachineRemoveAllSnapshotsUsingPOSTBadRequest struct {
}

// IsSuccess returns true when this vcs virtual machine remove all snapshots using p o s t bad request response has a 2xx status code
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this vcs virtual machine remove all snapshots using p o s t bad request response has a 3xx status code
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vcs virtual machine remove all snapshots using p o s t bad request response has a 4xx status code
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this vcs virtual machine remove all snapshots using p o s t bad request response has a 5xx status code
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this vcs virtual machine remove all snapshots using p o s t bad request response a status code equal to that given
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the vcs virtual machine remove all snapshots using p o s t bad request response
func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTBadRequest) Code() int {
	return 400
}

func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /vcs/virtual-machines/{virtualMachineId}/removeAllSnapshots][%d] vcsVirtualMachineRemoveAllSnapshotsUsingPOSTBadRequest ", 400)
}

func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTBadRequest) String() string {
	return fmt.Sprintf("[POST /vcs/virtual-machines/{virtualMachineId}/removeAllSnapshots][%d] vcsVirtualMachineRemoveAllSnapshotsUsingPOSTBadRequest ", 400)
}

func (o *VcsVirtualMachineRemoveAllSnapshotsUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
