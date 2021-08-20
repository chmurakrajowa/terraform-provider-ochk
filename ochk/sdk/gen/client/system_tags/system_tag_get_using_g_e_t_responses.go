// Code generated by go-swagger; DO NOT EDIT.

package system_tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// SystemTagGetUsingGETReader is a Reader for the SystemTagGetUsingGET structure.
type SystemTagGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemTagGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemTagGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSystemTagGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSystemTagGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSystemTagGetUsingGETOK creates a SystemTagGetUsingGETOK with default headers values
func NewSystemTagGetUsingGETOK() *SystemTagGetUsingGETOK {
	return &SystemTagGetUsingGETOK{}
}

/*SystemTagGetUsingGETOK handles this case with default header values.

OK
*/
type SystemTagGetUsingGETOK struct {
	Payload *models.SystemTagGetResponse
}

func (o *SystemTagGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /tags/systemTags/{systemTagId}][%d] systemTagGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *SystemTagGetUsingGETOK) GetPayload() *models.SystemTagGetResponse {
	return o.Payload
}

func (o *SystemTagGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemTagGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSystemTagGetUsingGETBadRequest creates a SystemTagGetUsingGETBadRequest with default headers values
func NewSystemTagGetUsingGETBadRequest() *SystemTagGetUsingGETBadRequest {
	return &SystemTagGetUsingGETBadRequest{}
}

/*SystemTagGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type SystemTagGetUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *SystemTagGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /tags/systemTags/{systemTagId}][%d] systemTagGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *SystemTagGetUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *SystemTagGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSystemTagGetUsingGETNotFound creates a SystemTagGetUsingGETNotFound with default headers values
func NewSystemTagGetUsingGETNotFound() *SystemTagGetUsingGETNotFound {
	return &SystemTagGetUsingGETNotFound{}
}

/*SystemTagGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type SystemTagGetUsingGETNotFound struct {
}

func (o *SystemTagGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /tags/systemTags/{systemTagId}][%d] systemTagGetUsingGETNotFound ", 404)
}

func (o *SystemTagGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}