// Code generated by go-swagger; DO NOT EDIT.

package subtenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// SubtenantUpdateUsingPUTReader is a Reader for the SubtenantUpdateUsingPUT structure.
type SubtenantUpdateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubtenantUpdateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubtenantUpdateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewSubtenantUpdateUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubtenantUpdateUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubtenantUpdateUsingPUTOK creates a SubtenantUpdateUsingPUTOK with default headers values
func NewSubtenantUpdateUsingPUTOK() *SubtenantUpdateUsingPUTOK {
	return &SubtenantUpdateUsingPUTOK{}
}

/* SubtenantUpdateUsingPUTOK describes a response with status code 200, with default header values.

OK
*/
type SubtenantUpdateUsingPUTOK struct {
	Payload *models.SubtenantUpdateResponse
}

func (o *SubtenantUpdateUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /subtenants/{subtenantId}][%d] subtenantUpdateUsingPUTOK  %+v", 200, o.Payload)
}
func (o *SubtenantUpdateUsingPUTOK) GetPayload() *models.SubtenantUpdateResponse {
	return o.Payload
}

func (o *SubtenantUpdateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubtenantUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubtenantUpdateUsingPUTCreated creates a SubtenantUpdateUsingPUTCreated with default headers values
func NewSubtenantUpdateUsingPUTCreated() *SubtenantUpdateUsingPUTCreated {
	return &SubtenantUpdateUsingPUTCreated{}
}

/* SubtenantUpdateUsingPUTCreated describes a response with status code 201, with default header values.

Entity has been updated
*/
type SubtenantUpdateUsingPUTCreated struct {
	Payload *models.SubtenantUpdateResponse
}

func (o *SubtenantUpdateUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /subtenants/{subtenantId}][%d] subtenantUpdateUsingPUTCreated  %+v", 201, o.Payload)
}
func (o *SubtenantUpdateUsingPUTCreated) GetPayload() *models.SubtenantUpdateResponse {
	return o.Payload
}

func (o *SubtenantUpdateUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubtenantUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubtenantUpdateUsingPUTBadRequest creates a SubtenantUpdateUsingPUTBadRequest with default headers values
func NewSubtenantUpdateUsingPUTBadRequest() *SubtenantUpdateUsingPUTBadRequest {
	return &SubtenantUpdateUsingPUTBadRequest{}
}

/* SubtenantUpdateUsingPUTBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type SubtenantUpdateUsingPUTBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *SubtenantUpdateUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /subtenants/{subtenantId}][%d] subtenantUpdateUsingPUTBadRequest  %+v", 400, o.Payload)
}
func (o *SubtenantUpdateUsingPUTBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *SubtenantUpdateUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
