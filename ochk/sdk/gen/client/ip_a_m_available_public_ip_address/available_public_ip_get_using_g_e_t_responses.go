// Code generated by go-swagger; DO NOT EDIT.

package ip_a_m_available_public_ip_address

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// AvailablePublicIPGetUsingGETReader is a Reader for the AvailablePublicIPGetUsingGET structure.
type AvailablePublicIPGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AvailablePublicIPGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAvailablePublicIPGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAvailablePublicIPGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAvailablePublicIPGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAvailablePublicIPGetUsingGETOK creates a AvailablePublicIPGetUsingGETOK with default headers values
func NewAvailablePublicIPGetUsingGETOK() *AvailablePublicIPGetUsingGETOK {
	return &AvailablePublicIPGetUsingGETOK{}
}

/*AvailablePublicIPGetUsingGETOK handles this case with default header values.

OK
*/
type AvailablePublicIPGetUsingGETOK struct {
	Payload *models.AvailablePublicIPGetResponse
}

func (o *AvailablePublicIPGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /ipam/ipaddress/public/available][%d] availablePublicIpGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *AvailablePublicIPGetUsingGETOK) GetPayload() *models.AvailablePublicIPGetResponse {
	return o.Payload
}

func (o *AvailablePublicIPGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AvailablePublicIPGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAvailablePublicIPGetUsingGETBadRequest creates a AvailablePublicIPGetUsingGETBadRequest with default headers values
func NewAvailablePublicIPGetUsingGETBadRequest() *AvailablePublicIPGetUsingGETBadRequest {
	return &AvailablePublicIPGetUsingGETBadRequest{}
}

/*AvailablePublicIPGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type AvailablePublicIPGetUsingGETBadRequest struct {
}

func (o *AvailablePublicIPGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /ipam/ipaddress/public/available][%d] availablePublicIpGetUsingGETBadRequest ", 400)
}

func (o *AvailablePublicIPGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAvailablePublicIPGetUsingGETNotFound creates a AvailablePublicIPGetUsingGETNotFound with default headers values
func NewAvailablePublicIPGetUsingGETNotFound() *AvailablePublicIPGetUsingGETNotFound {
	return &AvailablePublicIPGetUsingGETNotFound{}
}

/*AvailablePublicIPGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type AvailablePublicIPGetUsingGETNotFound struct {
}

func (o *AvailablePublicIPGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /ipam/ipaddress/public/available][%d] availablePublicIpGetUsingGETNotFound ", 404)
}

func (o *AvailablePublicIPGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}