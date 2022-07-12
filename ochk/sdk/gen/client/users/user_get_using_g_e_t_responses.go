// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// UserGetUsingGETReader is a Reader for the UserGetUsingGET structure.
type UserGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserGetUsingGETOK creates a UserGetUsingGETOK with default headers values
func NewUserGetUsingGETOK() *UserGetUsingGETOK {
	return &UserGetUsingGETOK{}
}

/* UserGetUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type UserGetUsingGETOK struct {
	Payload *models.UserGetResponse
}

func (o *UserGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /users/{userId}][%d] userGetUsingGETOK  %+v", 200, o.Payload)
}
func (o *UserGetUsingGETOK) GetPayload() *models.UserGetResponse {
	return o.Payload
}

func (o *UserGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGetUsingGETBadRequest creates a UserGetUsingGETBadRequest with default headers values
func NewUserGetUsingGETBadRequest() *UserGetUsingGETBadRequest {
	return &UserGetUsingGETBadRequest{}
}

/* UserGetUsingGETBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type UserGetUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *UserGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{userId}][%d] userGetUsingGETBadRequest  %+v", 400, o.Payload)
}
func (o *UserGetUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *UserGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGetUsingGETNotFound creates a UserGetUsingGETNotFound with default headers values
func NewUserGetUsingGETNotFound() *UserGetUsingGETNotFound {
	return &UserGetUsingGETNotFound{}
}

/* UserGetUsingGETNotFound describes a response with status code 404, with default header values.

Entity not found.
*/
type UserGetUsingGETNotFound struct {
}

func (o *UserGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{userId}][%d] userGetUsingGETNotFound ", 404)
}

func (o *UserGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
