// Code generated by go-swagger; DO NOT EDIT.

package ip_a_m_public_ip_allocations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// AllocationUpdateUsingPUTReader is a Reader for the AllocationUpdateUsingPUT structure.
type AllocationUpdateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllocationUpdateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllocationUpdateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAllocationUpdateUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAllocationUpdateUsingPUTOK creates a AllocationUpdateUsingPUTOK with default headers values
func NewAllocationUpdateUsingPUTOK() *AllocationUpdateUsingPUTOK {
	return &AllocationUpdateUsingPUTOK{}
}

/*AllocationUpdateUsingPUTOK handles this case with default header values.

Entity has been created
*/
type AllocationUpdateUsingPUTOK struct {
	Payload *models.UpdatePublicIPAllocationResponse
}

func (o *AllocationUpdateUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/ipaddress/public/allocation/{allocationId}][%d] allocationUpdateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *AllocationUpdateUsingPUTOK) GetPayload() *models.UpdatePublicIPAllocationResponse {
	return o.Payload
}

func (o *AllocationUpdateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdatePublicIPAllocationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllocationUpdateUsingPUTBadRequest creates a AllocationUpdateUsingPUTBadRequest with default headers values
func NewAllocationUpdateUsingPUTBadRequest() *AllocationUpdateUsingPUTBadRequest {
	return &AllocationUpdateUsingPUTBadRequest{}
}

/*AllocationUpdateUsingPUTBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type AllocationUpdateUsingPUTBadRequest struct {
}

func (o *AllocationUpdateUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /ipam/ipaddress/public/allocation/{allocationId}][%d] allocationUpdateUsingPUTBadRequest ", 400)
}

func (o *AllocationUpdateUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}