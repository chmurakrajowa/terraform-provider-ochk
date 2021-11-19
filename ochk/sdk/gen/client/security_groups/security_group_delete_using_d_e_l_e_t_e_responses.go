// Code generated by go-swagger; DO NOT EDIT.

package security_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// SecurityGroupDeleteUsingDELETEReader is a Reader for the SecurityGroupDeleteUsingDELETE structure.
type SecurityGroupDeleteUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityGroupDeleteUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityGroupDeleteUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewSecurityGroupDeleteUsingDELETECreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSecurityGroupDeleteUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSecurityGroupDeleteUsingDELETEOK creates a SecurityGroupDeleteUsingDELETEOK with default headers values
func NewSecurityGroupDeleteUsingDELETEOK() *SecurityGroupDeleteUsingDELETEOK {
	return &SecurityGroupDeleteUsingDELETEOK{}
}

/*SecurityGroupDeleteUsingDELETEOK handles this case with default header values.

OK
*/
type SecurityGroupDeleteUsingDELETEOK struct {
	Payload *models.DeleteSecurityGroupResponse
}

func (o *SecurityGroupDeleteUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /network/security-groups/{groupId}][%d] securityGroupDeleteUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *SecurityGroupDeleteUsingDELETEOK) GetPayload() *models.DeleteSecurityGroupResponse {
	return o.Payload
}

func (o *SecurityGroupDeleteUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteSecurityGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityGroupDeleteUsingDELETECreated creates a SecurityGroupDeleteUsingDELETECreated with default headers values
func NewSecurityGroupDeleteUsingDELETECreated() *SecurityGroupDeleteUsingDELETECreated {
	return &SecurityGroupDeleteUsingDELETECreated{}
}

/*SecurityGroupDeleteUsingDELETECreated handles this case with default header values.

Entity has been deleted
*/
type SecurityGroupDeleteUsingDELETECreated struct {
	Payload *models.DeleteSecurityGroupResponse
}

func (o *SecurityGroupDeleteUsingDELETECreated) Error() string {
	return fmt.Sprintf("[DELETE /network/security-groups/{groupId}][%d] securityGroupDeleteUsingDELETECreated  %+v", 201, o.Payload)
}

func (o *SecurityGroupDeleteUsingDELETECreated) GetPayload() *models.DeleteSecurityGroupResponse {
	return o.Payload
}

func (o *SecurityGroupDeleteUsingDELETECreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteSecurityGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityGroupDeleteUsingDELETEBadRequest creates a SecurityGroupDeleteUsingDELETEBadRequest with default headers values
func NewSecurityGroupDeleteUsingDELETEBadRequest() *SecurityGroupDeleteUsingDELETEBadRequest {
	return &SecurityGroupDeleteUsingDELETEBadRequest{}
}

/*SecurityGroupDeleteUsingDELETEBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type SecurityGroupDeleteUsingDELETEBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *SecurityGroupDeleteUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /network/security-groups/{groupId}][%d] securityGroupDeleteUsingDELETEBadRequest  %+v", 400, o.Payload)
}

func (o *SecurityGroupDeleteUsingDELETEBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *SecurityGroupDeleteUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
