// Code generated by go-swagger; DO NOT EDIT.

package virtual_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// VirtualNetworkListUsingGETReader is a Reader for the VirtualNetworkListUsingGET structure.
type VirtualNetworkListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualNetworkListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualNetworkListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVirtualNetworkListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualNetworkListUsingGETOK creates a VirtualNetworkListUsingGETOK with default headers values
func NewVirtualNetworkListUsingGETOK() *VirtualNetworkListUsingGETOK {
	return &VirtualNetworkListUsingGETOK{}
}

/*
VirtualNetworkListUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type VirtualNetworkListUsingGETOK struct {
	Payload *models.VirtualNetworkListResponse
}

// IsSuccess returns true when this virtual network list using g e t o k response has a 2xx status code
func (o *VirtualNetworkListUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtual network list using g e t o k response has a 3xx status code
func (o *VirtualNetworkListUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtual network list using g e t o k response has a 4xx status code
func (o *VirtualNetworkListUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtual network list using g e t o k response has a 5xx status code
func (o *VirtualNetworkListUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtual network list using g e t o k response a status code equal to that given
func (o *VirtualNetworkListUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtual network list using g e t o k response
func (o *VirtualNetworkListUsingGETOK) Code() int {
	return 200
}

func (o *VirtualNetworkListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /networks][%d] virtualNetworkListUsingGETOK  %+v", 200, o.Payload)
}

func (o *VirtualNetworkListUsingGETOK) String() string {
	return fmt.Sprintf("[GET /networks][%d] virtualNetworkListUsingGETOK  %+v", 200, o.Payload)
}

func (o *VirtualNetworkListUsingGETOK) GetPayload() *models.VirtualNetworkListResponse {
	return o.Payload
}

func (o *VirtualNetworkListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualNetworkListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualNetworkListUsingGETBadRequest creates a VirtualNetworkListUsingGETBadRequest with default headers values
func NewVirtualNetworkListUsingGETBadRequest() *VirtualNetworkListUsingGETBadRequest {
	return &VirtualNetworkListUsingGETBadRequest{}
}

/*
VirtualNetworkListUsingGETBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type VirtualNetworkListUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this virtual network list using g e t bad request response has a 2xx status code
func (o *VirtualNetworkListUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this virtual network list using g e t bad request response has a 3xx status code
func (o *VirtualNetworkListUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtual network list using g e t bad request response has a 4xx status code
func (o *VirtualNetworkListUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this virtual network list using g e t bad request response has a 5xx status code
func (o *VirtualNetworkListUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this virtual network list using g e t bad request response a status code equal to that given
func (o *VirtualNetworkListUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the virtual network list using g e t bad request response
func (o *VirtualNetworkListUsingGETBadRequest) Code() int {
	return 400
}

func (o *VirtualNetworkListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /networks][%d] virtualNetworkListUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *VirtualNetworkListUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /networks][%d] virtualNetworkListUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *VirtualNetworkListUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *VirtualNetworkListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
