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

// SystemTagDeleteUsingDELETEReader is a Reader for the SystemTagDeleteUsingDELETE structure.
type SystemTagDeleteUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemTagDeleteUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemTagDeleteUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSystemTagDeleteUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSystemTagDeleteUsingDELETEOK creates a SystemTagDeleteUsingDELETEOK with default headers values
func NewSystemTagDeleteUsingDELETEOK() *SystemTagDeleteUsingDELETEOK {
	return &SystemTagDeleteUsingDELETEOK{}
}

/* SystemTagDeleteUsingDELETEOK describes a response with status code 200, with default header values.

OK
*/
type SystemTagDeleteUsingDELETEOK struct {
	Payload *models.SystemTagDeleteResponse
}

func (o *SystemTagDeleteUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /tags/systemTags/{systemTagId}][%d] systemTagDeleteUsingDELETEOK  %+v", 200, o.Payload)
}
func (o *SystemTagDeleteUsingDELETEOK) GetPayload() *models.SystemTagDeleteResponse {
	return o.Payload
}

func (o *SystemTagDeleteUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemTagDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSystemTagDeleteUsingDELETEBadRequest creates a SystemTagDeleteUsingDELETEBadRequest with default headers values
func NewSystemTagDeleteUsingDELETEBadRequest() *SystemTagDeleteUsingDELETEBadRequest {
	return &SystemTagDeleteUsingDELETEBadRequest{}
}

/* SystemTagDeleteUsingDELETEBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type SystemTagDeleteUsingDELETEBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *SystemTagDeleteUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /tags/systemTags/{systemTagId}][%d] systemTagDeleteUsingDELETEBadRequest  %+v", 400, o.Payload)
}
func (o *SystemTagDeleteUsingDELETEBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *SystemTagDeleteUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
