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

// AllocationGetUsingGETReader is a Reader for the AllocationGetUsingGET structure.
type AllocationGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllocationGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllocationGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAllocationGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAllocationGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /ipam/ipaddress/public/allocation/{allocationId}] allocationGetUsingGET", response, response.Code())
	}
}

// NewAllocationGetUsingGETOK creates a AllocationGetUsingGETOK with default headers values
func NewAllocationGetUsingGETOK() *AllocationGetUsingGETOK {
	return &AllocationGetUsingGETOK{}
}

/*
AllocationGetUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type AllocationGetUsingGETOK struct {
	Payload *models.PublicIPAllocationGetResponse
}

// IsSuccess returns true when this allocation get using g e t o k response has a 2xx status code
func (o *AllocationGetUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this allocation get using g e t o k response has a 3xx status code
func (o *AllocationGetUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this allocation get using g e t o k response has a 4xx status code
func (o *AllocationGetUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this allocation get using g e t o k response has a 5xx status code
func (o *AllocationGetUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this allocation get using g e t o k response a status code equal to that given
func (o *AllocationGetUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the allocation get using g e t o k response
func (o *AllocationGetUsingGETOK) Code() int {
	return 200
}

func (o *AllocationGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /ipam/ipaddress/public/allocation/{allocationId}][%d] allocationGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *AllocationGetUsingGETOK) String() string {
	return fmt.Sprintf("[GET /ipam/ipaddress/public/allocation/{allocationId}][%d] allocationGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *AllocationGetUsingGETOK) GetPayload() *models.PublicIPAllocationGetResponse {
	return o.Payload
}

func (o *AllocationGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicIPAllocationGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllocationGetUsingGETBadRequest creates a AllocationGetUsingGETBadRequest with default headers values
func NewAllocationGetUsingGETBadRequest() *AllocationGetUsingGETBadRequest {
	return &AllocationGetUsingGETBadRequest{}
}

/*
AllocationGetUsingGETBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type AllocationGetUsingGETBadRequest struct {
}

// IsSuccess returns true when this allocation get using g e t bad request response has a 2xx status code
func (o *AllocationGetUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this allocation get using g e t bad request response has a 3xx status code
func (o *AllocationGetUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this allocation get using g e t bad request response has a 4xx status code
func (o *AllocationGetUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this allocation get using g e t bad request response has a 5xx status code
func (o *AllocationGetUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this allocation get using g e t bad request response a status code equal to that given
func (o *AllocationGetUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the allocation get using g e t bad request response
func (o *AllocationGetUsingGETBadRequest) Code() int {
	return 400
}

func (o *AllocationGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /ipam/ipaddress/public/allocation/{allocationId}][%d] allocationGetUsingGETBadRequest ", 400)
}

func (o *AllocationGetUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /ipam/ipaddress/public/allocation/{allocationId}][%d] allocationGetUsingGETBadRequest ", 400)
}

func (o *AllocationGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAllocationGetUsingGETNotFound creates a AllocationGetUsingGETNotFound with default headers values
func NewAllocationGetUsingGETNotFound() *AllocationGetUsingGETNotFound {
	return &AllocationGetUsingGETNotFound{}
}

/*
AllocationGetUsingGETNotFound describes a response with status code 404, with default header values.

Entity not found.
*/
type AllocationGetUsingGETNotFound struct {
}

// IsSuccess returns true when this allocation get using g e t not found response has a 2xx status code
func (o *AllocationGetUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this allocation get using g e t not found response has a 3xx status code
func (o *AllocationGetUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this allocation get using g e t not found response has a 4xx status code
func (o *AllocationGetUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this allocation get using g e t not found response has a 5xx status code
func (o *AllocationGetUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this allocation get using g e t not found response a status code equal to that given
func (o *AllocationGetUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the allocation get using g e t not found response
func (o *AllocationGetUsingGETNotFound) Code() int {
	return 404
}

func (o *AllocationGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /ipam/ipaddress/public/allocation/{allocationId}][%d] allocationGetUsingGETNotFound ", 404)
}

func (o *AllocationGetUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /ipam/ipaddress/public/allocation/{allocationId}][%d] allocationGetUsingGETNotFound ", 404)
}

func (o *AllocationGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
