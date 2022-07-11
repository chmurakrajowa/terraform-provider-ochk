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

// VirtualNetworkGetUsingGETReader is a Reader for the VirtualNetworkGetUsingGET structure.
type VirtualNetworkGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualNetworkGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualNetworkGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVirtualNetworkGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewVirtualNetworkGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualNetworkGetUsingGETOK creates a VirtualNetworkGetUsingGETOK with default headers values
func NewVirtualNetworkGetUsingGETOK() *VirtualNetworkGetUsingGETOK {
	return &VirtualNetworkGetUsingGETOK{}
}

/*VirtualNetworkGetUsingGETOK handles this case with default header values.

OK
*/
type VirtualNetworkGetUsingGETOK struct {
	Payload *models.VirtualNetworkGetResponse
}

func (o *VirtualNetworkGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /networks/{virtualNetworkId}][%d] virtualNetworkGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *VirtualNetworkGetUsingGETOK) GetPayload() *models.VirtualNetworkGetResponse {
	return o.Payload
}

func (o *VirtualNetworkGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualNetworkGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualNetworkGetUsingGETBadRequest creates a VirtualNetworkGetUsingGETBadRequest with default headers values
func NewVirtualNetworkGetUsingGETBadRequest() *VirtualNetworkGetUsingGETBadRequest {
	return &VirtualNetworkGetUsingGETBadRequest{}
}

/*VirtualNetworkGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type VirtualNetworkGetUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *VirtualNetworkGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /networks/{virtualNetworkId}][%d] virtualNetworkGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *VirtualNetworkGetUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *VirtualNetworkGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualNetworkGetUsingGETNotFound creates a VirtualNetworkGetUsingGETNotFound with default headers values
func NewVirtualNetworkGetUsingGETNotFound() *VirtualNetworkGetUsingGETNotFound {
	return &VirtualNetworkGetUsingGETNotFound{}
}

/*VirtualNetworkGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type VirtualNetworkGetUsingGETNotFound struct {
}

func (o *VirtualNetworkGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /networks/{virtualNetworkId}][%d] virtualNetworkGetUsingGETNotFound ", 404)
}

func (o *VirtualNetworkGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
