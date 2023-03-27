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

// AccountDeleteUsingDELETEReader is a Reader for the AccountDeleteUsingDELETE structure.
type AccountDeleteUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountDeleteUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountDeleteUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewAccountDeleteUsingDELETECreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccountDeleteUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAccountDeleteUsingDELETEOK creates a AccountDeleteUsingDELETEOK with default headers values
func NewAccountDeleteUsingDELETEOK() *AccountDeleteUsingDELETEOK {
	return &AccountDeleteUsingDELETEOK{}
}

/*
AccountDeleteUsingDELETEOK describes a response with status code 200, with default header values.

OK
*/
type AccountDeleteUsingDELETEOK struct {
	Payload *models.AccountDeleteResponse
}

// IsSuccess returns true when this account delete using d e l e t e o k response has a 2xx status code
func (o *AccountDeleteUsingDELETEOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account delete using d e l e t e o k response has a 3xx status code
func (o *AccountDeleteUsingDELETEOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account delete using d e l e t e o k response has a 4xx status code
func (o *AccountDeleteUsingDELETEOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account delete using d e l e t e o k response has a 5xx status code
func (o *AccountDeleteUsingDELETEOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account delete using d e l e t e o k response a status code equal to that given
func (o *AccountDeleteUsingDELETEOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account delete using d e l e t e o k response
func (o *AccountDeleteUsingDELETEOK) Code() int {
	return 200
}

func (o *AccountDeleteUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /billing/accounts/{accountId}][%d] accountDeleteUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *AccountDeleteUsingDELETEOK) String() string {
	return fmt.Sprintf("[DELETE /billing/accounts/{accountId}][%d] accountDeleteUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *AccountDeleteUsingDELETEOK) GetPayload() *models.AccountDeleteResponse {
	return o.Payload
}

func (o *AccountDeleteUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountDeleteUsingDELETECreated creates a AccountDeleteUsingDELETECreated with default headers values
func NewAccountDeleteUsingDELETECreated() *AccountDeleteUsingDELETECreated {
	return &AccountDeleteUsingDELETECreated{}
}

/*
AccountDeleteUsingDELETECreated describes a response with status code 201, with default header values.

Entity has been deleted
*/
type AccountDeleteUsingDELETECreated struct {
	Payload *models.AccountDeleteResponse
}

// IsSuccess returns true when this account delete using d e l e t e created response has a 2xx status code
func (o *AccountDeleteUsingDELETECreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account delete using d e l e t e created response has a 3xx status code
func (o *AccountDeleteUsingDELETECreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account delete using d e l e t e created response has a 4xx status code
func (o *AccountDeleteUsingDELETECreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this account delete using d e l e t e created response has a 5xx status code
func (o *AccountDeleteUsingDELETECreated) IsServerError() bool {
	return false
}

// IsCode returns true when this account delete using d e l e t e created response a status code equal to that given
func (o *AccountDeleteUsingDELETECreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the account delete using d e l e t e created response
func (o *AccountDeleteUsingDELETECreated) Code() int {
	return 201
}

func (o *AccountDeleteUsingDELETECreated) Error() string {
	return fmt.Sprintf("[DELETE /billing/accounts/{accountId}][%d] accountDeleteUsingDELETECreated  %+v", 201, o.Payload)
}

func (o *AccountDeleteUsingDELETECreated) String() string {
	return fmt.Sprintf("[DELETE /billing/accounts/{accountId}][%d] accountDeleteUsingDELETECreated  %+v", 201, o.Payload)
}

func (o *AccountDeleteUsingDELETECreated) GetPayload() *models.AccountDeleteResponse {
	return o.Payload
}

func (o *AccountDeleteUsingDELETECreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountDeleteUsingDELETEBadRequest creates a AccountDeleteUsingDELETEBadRequest with default headers values
func NewAccountDeleteUsingDELETEBadRequest() *AccountDeleteUsingDELETEBadRequest {
	return &AccountDeleteUsingDELETEBadRequest{}
}

/*
AccountDeleteUsingDELETEBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type AccountDeleteUsingDELETEBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this account delete using d e l e t e bad request response has a 2xx status code
func (o *AccountDeleteUsingDELETEBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account delete using d e l e t e bad request response has a 3xx status code
func (o *AccountDeleteUsingDELETEBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account delete using d e l e t e bad request response has a 4xx status code
func (o *AccountDeleteUsingDELETEBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this account delete using d e l e t e bad request response has a 5xx status code
func (o *AccountDeleteUsingDELETEBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this account delete using d e l e t e bad request response a status code equal to that given
func (o *AccountDeleteUsingDELETEBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the account delete using d e l e t e bad request response
func (o *AccountDeleteUsingDELETEBadRequest) Code() int {
	return 400
}

func (o *AccountDeleteUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /billing/accounts/{accountId}][%d] accountDeleteUsingDELETEBadRequest  %+v", 400, o.Payload)
}

func (o *AccountDeleteUsingDELETEBadRequest) String() string {
	return fmt.Sprintf("[DELETE /billing/accounts/{accountId}][%d] accountDeleteUsingDELETEBadRequest  %+v", 400, o.Payload)
}

func (o *AccountDeleteUsingDELETEBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *AccountDeleteUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
