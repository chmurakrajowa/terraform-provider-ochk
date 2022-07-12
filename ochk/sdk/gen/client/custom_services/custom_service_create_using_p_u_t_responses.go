// Code generated by go-swagger; DO NOT EDIT.

package custom_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// CustomServiceCreateUsingPUTReader is a Reader for the CustomServiceCreateUsingPUT structure.
type CustomServiceCreateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomServiceCreateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomServiceCreateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCustomServiceCreateUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCustomServiceCreateUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCustomServiceCreateUsingPUTOK creates a CustomServiceCreateUsingPUTOK with default headers values
func NewCustomServiceCreateUsingPUTOK() *CustomServiceCreateUsingPUTOK {
	return &CustomServiceCreateUsingPUTOK{}
}

/* CustomServiceCreateUsingPUTOK describes a response with status code 200, with default header values.

OK
*/
type CustomServiceCreateUsingPUTOK struct {
	Payload *models.CreateCustomServiceResponse
}

func (o *CustomServiceCreateUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /network/custom-services][%d] customServiceCreateUsingPUTOK  %+v", 200, o.Payload)
}
func (o *CustomServiceCreateUsingPUTOK) GetPayload() *models.CreateCustomServiceResponse {
	return o.Payload
}

func (o *CustomServiceCreateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateCustomServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomServiceCreateUsingPUTCreated creates a CustomServiceCreateUsingPUTCreated with default headers values
func NewCustomServiceCreateUsingPUTCreated() *CustomServiceCreateUsingPUTCreated {
	return &CustomServiceCreateUsingPUTCreated{}
}

/* CustomServiceCreateUsingPUTCreated describes a response with status code 201, with default header values.

Entity has been created
*/
type CustomServiceCreateUsingPUTCreated struct {
	Payload *models.CreateCustomServiceResponse
}

func (o *CustomServiceCreateUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /network/custom-services][%d] customServiceCreateUsingPUTCreated  %+v", 201, o.Payload)
}
func (o *CustomServiceCreateUsingPUTCreated) GetPayload() *models.CreateCustomServiceResponse {
	return o.Payload
}

func (o *CustomServiceCreateUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateCustomServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomServiceCreateUsingPUTBadRequest creates a CustomServiceCreateUsingPUTBadRequest with default headers values
func NewCustomServiceCreateUsingPUTBadRequest() *CustomServiceCreateUsingPUTBadRequest {
	return &CustomServiceCreateUsingPUTBadRequest{}
}

/* CustomServiceCreateUsingPUTBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type CustomServiceCreateUsingPUTBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *CustomServiceCreateUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /network/custom-services][%d] customServiceCreateUsingPUTBadRequest  %+v", 400, o.Payload)
}
func (o *CustomServiceCreateUsingPUTBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *CustomServiceCreateUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
