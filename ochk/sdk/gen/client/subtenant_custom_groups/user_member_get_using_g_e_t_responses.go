// Code generated by go-swagger; DO NOT EDIT.

package subtenant_custom_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// UserMemberGetUsingGETReader is a Reader for the UserMemberGetUsingGET structure.
type UserMemberGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserMemberGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserMemberGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserMemberGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserMemberGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserMemberGetUsingGETOK creates a UserMemberGetUsingGETOK with default headers values
func NewUserMemberGetUsingGETOK() *UserMemberGetUsingGETOK {
	return &UserMemberGetUsingGETOK{}
}

/* UserMemberGetUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type UserMemberGetUsingGETOK struct {
	Payload *models.UserGetResponse
}

func (o *UserMemberGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /subtenants/{subtenantId}/groups/{parentGroupId}/members/users/{userId}][%d] userMemberGetUsingGETOK  %+v", 200, o.Payload)
}
func (o *UserMemberGetUsingGETOK) GetPayload() *models.UserGetResponse {
	return o.Payload
}

func (o *UserMemberGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserMemberGetUsingGETBadRequest creates a UserMemberGetUsingGETBadRequest with default headers values
func NewUserMemberGetUsingGETBadRequest() *UserMemberGetUsingGETBadRequest {
	return &UserMemberGetUsingGETBadRequest{}
}

/* UserMemberGetUsingGETBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type UserMemberGetUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *UserMemberGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /subtenants/{subtenantId}/groups/{parentGroupId}/members/users/{userId}][%d] userMemberGetUsingGETBadRequest  %+v", 400, o.Payload)
}
func (o *UserMemberGetUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *UserMemberGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserMemberGetUsingGETNotFound creates a UserMemberGetUsingGETNotFound with default headers values
func NewUserMemberGetUsingGETNotFound() *UserMemberGetUsingGETNotFound {
	return &UserMemberGetUsingGETNotFound{}
}

/* UserMemberGetUsingGETNotFound describes a response with status code 404, with default header values.

Entity not found.
*/
type UserMemberGetUsingGETNotFound struct {
}

func (o *UserMemberGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /subtenants/{subtenantId}/groups/{parentGroupId}/members/users/{userId}][%d] userMemberGetUsingGETNotFound ", 404)
}

func (o *UserMemberGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
