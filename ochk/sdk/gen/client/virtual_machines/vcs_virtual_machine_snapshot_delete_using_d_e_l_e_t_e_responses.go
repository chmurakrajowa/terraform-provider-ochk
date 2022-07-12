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

// VcsVirtualMachineSnapshotDeleteUsingDELETEReader is a Reader for the VcsVirtualMachineSnapshotDeleteUsingDELETE structure.
type VcsVirtualMachineSnapshotDeleteUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VcsVirtualMachineSnapshotDeleteUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVcsVirtualMachineSnapshotDeleteUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVcsVirtualMachineSnapshotDeleteUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVcsVirtualMachineSnapshotDeleteUsingDELETEOK creates a VcsVirtualMachineSnapshotDeleteUsingDELETEOK with default headers values
func NewVcsVirtualMachineSnapshotDeleteUsingDELETEOK() *VcsVirtualMachineSnapshotDeleteUsingDELETEOK {
	return &VcsVirtualMachineSnapshotDeleteUsingDELETEOK{}
}

/* VcsVirtualMachineSnapshotDeleteUsingDELETEOK describes a response with status code 200, with default header values.

Request has been completed successfully
*/
type VcsVirtualMachineSnapshotDeleteUsingDELETEOK struct {
	Payload *models.DeleteSnapshotResponse
}

func (o *VcsVirtualMachineSnapshotDeleteUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /vcs/virtual-machines/{virtualMachineId}/snapshots/{snapshotId}][%d] vcsVirtualMachineSnapshotDeleteUsingDELETEOK  %+v", 200, o.Payload)
}
func (o *VcsVirtualMachineSnapshotDeleteUsingDELETEOK) GetPayload() *models.DeleteSnapshotResponse {
	return o.Payload
}

func (o *VcsVirtualMachineSnapshotDeleteUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVcsVirtualMachineSnapshotDeleteUsingDELETEBadRequest creates a VcsVirtualMachineSnapshotDeleteUsingDELETEBadRequest with default headers values
func NewVcsVirtualMachineSnapshotDeleteUsingDELETEBadRequest() *VcsVirtualMachineSnapshotDeleteUsingDELETEBadRequest {
	return &VcsVirtualMachineSnapshotDeleteUsingDELETEBadRequest{}
}

/* VcsVirtualMachineSnapshotDeleteUsingDELETEBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type VcsVirtualMachineSnapshotDeleteUsingDELETEBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *VcsVirtualMachineSnapshotDeleteUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /vcs/virtual-machines/{virtualMachineId}/snapshots/{snapshotId}][%d] vcsVirtualMachineSnapshotDeleteUsingDELETEBadRequest  %+v", 400, o.Payload)
}
func (o *VcsVirtualMachineSnapshotDeleteUsingDELETEBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *VcsVirtualMachineSnapshotDeleteUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
