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

// VcsVirtualMachineGetUsingGETReader is a Reader for the VcsVirtualMachineGetUsingGET structure.
type VcsVirtualMachineGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VcsVirtualMachineGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVcsVirtualMachineGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVcsVirtualMachineGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewVcsVirtualMachineGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /vcs/virtual-machines/{virtualMachineId}] vcsVirtualMachineGetUsingGET", response, response.Code())
	}
}

// NewVcsVirtualMachineGetUsingGETOK creates a VcsVirtualMachineGetUsingGETOK with default headers values
func NewVcsVirtualMachineGetUsingGETOK() *VcsVirtualMachineGetUsingGETOK {
	return &VcsVirtualMachineGetUsingGETOK{}
}

/*
VcsVirtualMachineGetUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type VcsVirtualMachineGetUsingGETOK struct {
	Payload *models.VcsVirtualMachineGetResponse
}

// IsSuccess returns true when this vcs virtual machine get using g e t o k response has a 2xx status code
func (o *VcsVirtualMachineGetUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vcs virtual machine get using g e t o k response has a 3xx status code
func (o *VcsVirtualMachineGetUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vcs virtual machine get using g e t o k response has a 4xx status code
func (o *VcsVirtualMachineGetUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vcs virtual machine get using g e t o k response has a 5xx status code
func (o *VcsVirtualMachineGetUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vcs virtual machine get using g e t o k response a status code equal to that given
func (o *VcsVirtualMachineGetUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vcs virtual machine get using g e t o k response
func (o *VcsVirtualMachineGetUsingGETOK) Code() int {
	return 200
}

func (o *VcsVirtualMachineGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /vcs/virtual-machines/{virtualMachineId}][%d] vcsVirtualMachineGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *VcsVirtualMachineGetUsingGETOK) String() string {
	return fmt.Sprintf("[GET /vcs/virtual-machines/{virtualMachineId}][%d] vcsVirtualMachineGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *VcsVirtualMachineGetUsingGETOK) GetPayload() *models.VcsVirtualMachineGetResponse {
	return o.Payload
}

func (o *VcsVirtualMachineGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VcsVirtualMachineGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVcsVirtualMachineGetUsingGETBadRequest creates a VcsVirtualMachineGetUsingGETBadRequest with default headers values
func NewVcsVirtualMachineGetUsingGETBadRequest() *VcsVirtualMachineGetUsingGETBadRequest {
	return &VcsVirtualMachineGetUsingGETBadRequest{}
}

/*
VcsVirtualMachineGetUsingGETBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type VcsVirtualMachineGetUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this vcs virtual machine get using g e t bad request response has a 2xx status code
func (o *VcsVirtualMachineGetUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this vcs virtual machine get using g e t bad request response has a 3xx status code
func (o *VcsVirtualMachineGetUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vcs virtual machine get using g e t bad request response has a 4xx status code
func (o *VcsVirtualMachineGetUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this vcs virtual machine get using g e t bad request response has a 5xx status code
func (o *VcsVirtualMachineGetUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this vcs virtual machine get using g e t bad request response a status code equal to that given
func (o *VcsVirtualMachineGetUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the vcs virtual machine get using g e t bad request response
func (o *VcsVirtualMachineGetUsingGETBadRequest) Code() int {
	return 400
}

func (o *VcsVirtualMachineGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /vcs/virtual-machines/{virtualMachineId}][%d] vcsVirtualMachineGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *VcsVirtualMachineGetUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /vcs/virtual-machines/{virtualMachineId}][%d] vcsVirtualMachineGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *VcsVirtualMachineGetUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *VcsVirtualMachineGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVcsVirtualMachineGetUsingGETNotFound creates a VcsVirtualMachineGetUsingGETNotFound with default headers values
func NewVcsVirtualMachineGetUsingGETNotFound() *VcsVirtualMachineGetUsingGETNotFound {
	return &VcsVirtualMachineGetUsingGETNotFound{}
}

/*
VcsVirtualMachineGetUsingGETNotFound describes a response with status code 404, with default header values.

Entity not found.
*/
type VcsVirtualMachineGetUsingGETNotFound struct {
}

// IsSuccess returns true when this vcs virtual machine get using g e t not found response has a 2xx status code
func (o *VcsVirtualMachineGetUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this vcs virtual machine get using g e t not found response has a 3xx status code
func (o *VcsVirtualMachineGetUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vcs virtual machine get using g e t not found response has a 4xx status code
func (o *VcsVirtualMachineGetUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this vcs virtual machine get using g e t not found response has a 5xx status code
func (o *VcsVirtualMachineGetUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this vcs virtual machine get using g e t not found response a status code equal to that given
func (o *VcsVirtualMachineGetUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the vcs virtual machine get using g e t not found response
func (o *VcsVirtualMachineGetUsingGETNotFound) Code() int {
	return 404
}

func (o *VcsVirtualMachineGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /vcs/virtual-machines/{virtualMachineId}][%d] vcsVirtualMachineGetUsingGETNotFound ", 404)
}

func (o *VcsVirtualMachineGetUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /vcs/virtual-machines/{virtualMachineId}][%d] vcsVirtualMachineGetUsingGETNotFound ", 404)
}

func (o *VcsVirtualMachineGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
