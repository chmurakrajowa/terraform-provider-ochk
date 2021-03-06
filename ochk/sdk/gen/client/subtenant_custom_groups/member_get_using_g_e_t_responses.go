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

// MemberGetUsingGETReader is a Reader for the MemberGetUsingGET structure.
type MemberGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MemberGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMemberGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMemberGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMemberGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMemberGetUsingGETOK creates a MemberGetUsingGETOK with default headers values
func NewMemberGetUsingGETOK() *MemberGetUsingGETOK {
	return &MemberGetUsingGETOK{}
}

/*MemberGetUsingGETOK handles this case with default header values.

OK
*/
type MemberGetUsingGETOK struct {
	Payload *models.UserGetResponse
}

func (o *MemberGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /subtenants/{subtenantId}/groups/{groupId}/members/{userId}][%d] memberGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *MemberGetUsingGETOK) GetPayload() *models.UserGetResponse {
	return o.Payload
}

func (o *MemberGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMemberGetUsingGETBadRequest creates a MemberGetUsingGETBadRequest with default headers values
func NewMemberGetUsingGETBadRequest() *MemberGetUsingGETBadRequest {
	return &MemberGetUsingGETBadRequest{}
}

/*MemberGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type MemberGetUsingGETBadRequest struct {
}

func (o *MemberGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /subtenants/{subtenantId}/groups/{groupId}/members/{userId}][%d] memberGetUsingGETBadRequest ", 400)
}

func (o *MemberGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMemberGetUsingGETNotFound creates a MemberGetUsingGETNotFound with default headers values
func NewMemberGetUsingGETNotFound() *MemberGetUsingGETNotFound {
	return &MemberGetUsingGETNotFound{}
}

/*MemberGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type MemberGetUsingGETNotFound struct {
}

func (o *MemberGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /subtenants/{subtenantId}/groups/{groupId}/members/{userId}][%d] memberGetUsingGETNotFound ", 404)
}

func (o *MemberGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
