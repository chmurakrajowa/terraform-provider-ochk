// Code generated by go-swagger; DO NOT EDIT.

package billing_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// AccountGetUsingGETReader is a Reader for the AccountGetUsingGET structure.
type AccountGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccountGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccountGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAccountGetUsingGETOK creates a AccountGetUsingGETOK with default headers values
func NewAccountGetUsingGETOK() *AccountGetUsingGETOK {
	return &AccountGetUsingGETOK{}
}

/*
AccountGetUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type AccountGetUsingGETOK struct {
	Payload *models.AccountGetResponse
}

// IsSuccess returns true when this account get using g e t o k response has a 2xx status code
func (o *AccountGetUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account get using g e t o k response has a 3xx status code
func (o *AccountGetUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account get using g e t o k response has a 4xx status code
func (o *AccountGetUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account get using g e t o k response has a 5xx status code
func (o *AccountGetUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account get using g e t o k response a status code equal to that given
func (o *AccountGetUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account get using g e t o k response
func (o *AccountGetUsingGETOK) Code() int {
	return 200
}

func (o *AccountGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /billing/accounts/{accountId}][%d] accountGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *AccountGetUsingGETOK) String() string {
	return fmt.Sprintf("[GET /billing/accounts/{accountId}][%d] accountGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *AccountGetUsingGETOK) GetPayload() *models.AccountGetResponse {
	return o.Payload
}

func (o *AccountGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountGetUsingGETBadRequest creates a AccountGetUsingGETBadRequest with default headers values
func NewAccountGetUsingGETBadRequest() *AccountGetUsingGETBadRequest {
	return &AccountGetUsingGETBadRequest{}
}

/*
AccountGetUsingGETBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type AccountGetUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this account get using g e t bad request response has a 2xx status code
func (o *AccountGetUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account get using g e t bad request response has a 3xx status code
func (o *AccountGetUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account get using g e t bad request response has a 4xx status code
func (o *AccountGetUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this account get using g e t bad request response has a 5xx status code
func (o *AccountGetUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this account get using g e t bad request response a status code equal to that given
func (o *AccountGetUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the account get using g e t bad request response
func (o *AccountGetUsingGETBadRequest) Code() int {
	return 400
}

func (o *AccountGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /billing/accounts/{accountId}][%d] accountGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *AccountGetUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /billing/accounts/{accountId}][%d] accountGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *AccountGetUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *AccountGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountGetUsingGETNotFound creates a AccountGetUsingGETNotFound with default headers values
func NewAccountGetUsingGETNotFound() *AccountGetUsingGETNotFound {
	return &AccountGetUsingGETNotFound{}
}

/*
AccountGetUsingGETNotFound describes a response with status code 404, with default header values.

Entity not found.
*/
type AccountGetUsingGETNotFound struct {
}

// IsSuccess returns true when this account get using g e t not found response has a 2xx status code
func (o *AccountGetUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account get using g e t not found response has a 3xx status code
func (o *AccountGetUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account get using g e t not found response has a 4xx status code
func (o *AccountGetUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this account get using g e t not found response has a 5xx status code
func (o *AccountGetUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this account get using g e t not found response a status code equal to that given
func (o *AccountGetUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the account get using g e t not found response
func (o *AccountGetUsingGETNotFound) Code() int {
	return 404
}

func (o *AccountGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /billing/accounts/{accountId}][%d] accountGetUsingGETNotFound ", 404)
}

func (o *AccountGetUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /billing/accounts/{accountId}][%d] accountGetUsingGETNotFound ", 404)
}

func (o *AccountGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
