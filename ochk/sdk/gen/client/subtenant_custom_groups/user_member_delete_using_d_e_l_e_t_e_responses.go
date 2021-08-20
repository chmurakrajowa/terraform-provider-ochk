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

// UserMemberDeleteUsingDELETEReader is a Reader for the UserMemberDeleteUsingDELETE structure.
type UserMemberDeleteUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserMemberDeleteUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserMemberDeleteUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserMemberDeleteUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserMemberDeleteUsingDELETEOK creates a UserMemberDeleteUsingDELETEOK with default headers values
func NewUserMemberDeleteUsingDELETEOK() *UserMemberDeleteUsingDELETEOK {
	return &UserMemberDeleteUsingDELETEOK{}
}

/*UserMemberDeleteUsingDELETEOK handles this case with default header values.

OK
*/
type UserMemberDeleteUsingDELETEOK struct {
	Payload *models.DeleteUserMemberResponse
}

func (o *UserMemberDeleteUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /subtenants/{subtenantId}/groups/{parentGroupId}/members/users/{userId}][%d] userMemberDeleteUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *UserMemberDeleteUsingDELETEOK) GetPayload() *models.DeleteUserMemberResponse {
	return o.Payload
}

func (o *UserMemberDeleteUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteUserMemberResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserMemberDeleteUsingDELETEBadRequest creates a UserMemberDeleteUsingDELETEBadRequest with default headers values
func NewUserMemberDeleteUsingDELETEBadRequest() *UserMemberDeleteUsingDELETEBadRequest {
	return &UserMemberDeleteUsingDELETEBadRequest{}
}

/*UserMemberDeleteUsingDELETEBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type UserMemberDeleteUsingDELETEBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *UserMemberDeleteUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /subtenants/{subtenantId}/groups/{parentGroupId}/members/users/{userId}][%d] userMemberDeleteUsingDELETEBadRequest  %+v", 400, o.Payload)
}

func (o *UserMemberDeleteUsingDELETEBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *UserMemberDeleteUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}