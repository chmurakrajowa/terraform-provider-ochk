// Code generated by go-swagger; DO NOT EDIT.

package security_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// SecurityPolicyGetUsingGETReader is a Reader for the SecurityPolicyGetUsingGET structure.
type SecurityPolicyGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityPolicyGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityPolicyGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSecurityPolicyGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSecurityPolicyGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSecurityPolicyGetUsingGETOK creates a SecurityPolicyGetUsingGETOK with default headers values
func NewSecurityPolicyGetUsingGETOK() *SecurityPolicyGetUsingGETOK {
	return &SecurityPolicyGetUsingGETOK{}
}

/*SecurityPolicyGetUsingGETOK handles this case with default header values.

OK
*/
type SecurityPolicyGetUsingGETOK struct {
	Payload *models.SecurityPolicyGetResponse
}

func (o *SecurityPolicyGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /network/firewall/security-policies/{SecurityPolicyId}][%d] securityPolicyGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *SecurityPolicyGetUsingGETOK) GetPayload() *models.SecurityPolicyGetResponse {
	return o.Payload
}

func (o *SecurityPolicyGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityPolicyGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityPolicyGetUsingGETBadRequest creates a SecurityPolicyGetUsingGETBadRequest with default headers values
func NewSecurityPolicyGetUsingGETBadRequest() *SecurityPolicyGetUsingGETBadRequest {
	return &SecurityPolicyGetUsingGETBadRequest{}
}

/*SecurityPolicyGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type SecurityPolicyGetUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *SecurityPolicyGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /network/firewall/security-policies/{SecurityPolicyId}][%d] securityPolicyGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *SecurityPolicyGetUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *SecurityPolicyGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityPolicyGetUsingGETNotFound creates a SecurityPolicyGetUsingGETNotFound with default headers values
func NewSecurityPolicyGetUsingGETNotFound() *SecurityPolicyGetUsingGETNotFound {
	return &SecurityPolicyGetUsingGETNotFound{}
}

/*SecurityPolicyGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type SecurityPolicyGetUsingGETNotFound struct {
}

func (o *SecurityPolicyGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /network/firewall/security-policies/{SecurityPolicyId}][%d] securityPolicyGetUsingGETNotFound ", 404)
}

func (o *SecurityPolicyGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
