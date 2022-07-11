// Code generated by go-swagger; DO NOT EDIT.

package local_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// LocalGroupGetUsingGETReader is a Reader for the LocalGroupGetUsingGET structure.
type LocalGroupGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocalGroupGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocalGroupGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLocalGroupGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLocalGroupGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLocalGroupGetUsingGETOK creates a LocalGroupGetUsingGETOK with default headers values
func NewLocalGroupGetUsingGETOK() *LocalGroupGetUsingGETOK {
	return &LocalGroupGetUsingGETOK{}
}

/*LocalGroupGetUsingGETOK handles this case with default header values.

OK
*/
type LocalGroupGetUsingGETOK struct {
	Payload *models.LocalGroupGetResponse
}

func (o *LocalGroupGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /groups/local/{groupId}][%d] localGroupGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *LocalGroupGetUsingGETOK) GetPayload() *models.LocalGroupGetResponse {
	return o.Payload
}

func (o *LocalGroupGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocalGroupGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocalGroupGetUsingGETBadRequest creates a LocalGroupGetUsingGETBadRequest with default headers values
func NewLocalGroupGetUsingGETBadRequest() *LocalGroupGetUsingGETBadRequest {
	return &LocalGroupGetUsingGETBadRequest{}
}

/*LocalGroupGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type LocalGroupGetUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *LocalGroupGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /groups/local/{groupId}][%d] localGroupGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *LocalGroupGetUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *LocalGroupGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocalGroupGetUsingGETNotFound creates a LocalGroupGetUsingGETNotFound with default headers values
func NewLocalGroupGetUsingGETNotFound() *LocalGroupGetUsingGETNotFound {
	return &LocalGroupGetUsingGETNotFound{}
}

/*LocalGroupGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type LocalGroupGetUsingGETNotFound struct {
}

func (o *LocalGroupGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /groups/local/{groupId}][%d] localGroupGetUsingGETNotFound ", 404)
}

func (o *LocalGroupGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}