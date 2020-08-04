// Code generated by go-swagger; DO NOT EDIT.

package security_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
)

// SecurityGroupListUsingGETReader is a Reader for the SecurityGroupListUsingGET structure.
type SecurityGroupListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityGroupListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityGroupListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSecurityGroupListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSecurityGroupListUsingGETOK creates a SecurityGroupListUsingGETOK with default headers values
func NewSecurityGroupListUsingGETOK() *SecurityGroupListUsingGETOK {
	return &SecurityGroupListUsingGETOK{}
}

/*SecurityGroupListUsingGETOK handles this case with default header values.

OK
*/
type SecurityGroupListUsingGETOK struct {
	Payload *models.SecurityGroupListResponse
}

func (o *SecurityGroupListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /network/security-groups][%d] securityGroupListUsingGETOK  %+v", 200, o.Payload)
}

func (o *SecurityGroupListUsingGETOK) GetPayload() *models.SecurityGroupListResponse {
	return o.Payload
}

func (o *SecurityGroupListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityGroupListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityGroupListUsingGETBadRequest creates a SecurityGroupListUsingGETBadRequest with default headers values
func NewSecurityGroupListUsingGETBadRequest() *SecurityGroupListUsingGETBadRequest {
	return &SecurityGroupListUsingGETBadRequest{}
}

/*SecurityGroupListUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type SecurityGroupListUsingGETBadRequest struct {
}

func (o *SecurityGroupListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /network/security-groups][%d] securityGroupListUsingGETBadRequest ", 400)
}

func (o *SecurityGroupListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
