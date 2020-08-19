// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// VcsVirtualMachineGroupGetUsingGETReader is a Reader for the VcsVirtualMachineGroupGetUsingGET structure.
type VcsVirtualMachineGroupGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VcsVirtualMachineGroupGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVcsVirtualMachineGroupGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVcsVirtualMachineGroupGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewVcsVirtualMachineGroupGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVcsVirtualMachineGroupGetUsingGETOK creates a VcsVirtualMachineGroupGetUsingGETOK with default headers values
func NewVcsVirtualMachineGroupGetUsingGETOK() *VcsVirtualMachineGroupGetUsingGETOK {
	return &VcsVirtualMachineGroupGetUsingGETOK{}
}

/*VcsVirtualMachineGroupGetUsingGETOK handles this case with default header values.

OK
*/
type VcsVirtualMachineGroupGetUsingGETOK struct {
	Payload *models.NetworkGetResponse
}

func (o *VcsVirtualMachineGroupGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /vcs/networks/{networkId}][%d] vcsVirtualMachineGroupGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *VcsVirtualMachineGroupGetUsingGETOK) GetPayload() *models.NetworkGetResponse {
	return o.Payload
}

func (o *VcsVirtualMachineGroupGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVcsVirtualMachineGroupGetUsingGETBadRequest creates a VcsVirtualMachineGroupGetUsingGETBadRequest with default headers values
func NewVcsVirtualMachineGroupGetUsingGETBadRequest() *VcsVirtualMachineGroupGetUsingGETBadRequest {
	return &VcsVirtualMachineGroupGetUsingGETBadRequest{}
}

/*VcsVirtualMachineGroupGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type VcsVirtualMachineGroupGetUsingGETBadRequest struct {
}

func (o *VcsVirtualMachineGroupGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /vcs/networks/{networkId}][%d] vcsVirtualMachineGroupGetUsingGETBadRequest ", 400)
}

func (o *VcsVirtualMachineGroupGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVcsVirtualMachineGroupGetUsingGETNotFound creates a VcsVirtualMachineGroupGetUsingGETNotFound with default headers values
func NewVcsVirtualMachineGroupGetUsingGETNotFound() *VcsVirtualMachineGroupGetUsingGETNotFound {
	return &VcsVirtualMachineGroupGetUsingGETNotFound{}
}

/*VcsVirtualMachineGroupGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type VcsVirtualMachineGroupGetUsingGETNotFound struct {
}

func (o *VcsVirtualMachineGroupGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /vcs/networks/{networkId}][%d] vcsVirtualMachineGroupGetUsingGETNotFound ", 404)
}

func (o *VcsVirtualMachineGroupGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
