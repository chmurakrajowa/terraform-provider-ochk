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

// MemberDeleteUsingDELETEReader is a Reader for the MemberDeleteUsingDELETE structure.
type MemberDeleteUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MemberDeleteUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMemberDeleteUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMemberDeleteUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMemberDeleteUsingDELETEOK creates a MemberDeleteUsingDELETEOK with default headers values
func NewMemberDeleteUsingDELETEOK() *MemberDeleteUsingDELETEOK {
	return &MemberDeleteUsingDELETEOK{}
}

/*MemberDeleteUsingDELETEOK handles this case with default header values.

OK
*/
type MemberDeleteUsingDELETEOK struct {
	Payload *models.DeleteMemberResponse
}

func (o *MemberDeleteUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /subtenants/{subtenantId}/groups/{groupId}/members/{userId}][%d] memberDeleteUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *MemberDeleteUsingDELETEOK) GetPayload() *models.DeleteMemberResponse {
	return o.Payload
}

func (o *MemberDeleteUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteMemberResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMemberDeleteUsingDELETEBadRequest creates a MemberDeleteUsingDELETEBadRequest with default headers values
func NewMemberDeleteUsingDELETEBadRequest() *MemberDeleteUsingDELETEBadRequest {
	return &MemberDeleteUsingDELETEBadRequest{}
}

/*MemberDeleteUsingDELETEBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type MemberDeleteUsingDELETEBadRequest struct {
}

func (o *MemberDeleteUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /subtenants/{subtenantId}/groups/{groupId}/members/{userId}][%d] memberDeleteUsingDELETEBadRequest ", 400)
}

func (o *MemberDeleteUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}