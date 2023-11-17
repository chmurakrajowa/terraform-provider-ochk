// Code generated by go-swagger; DO NOT EDIT.

package k_m_s_key_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// KeyCreateUsingPUTReader is a Reader for the KeyCreateUsingPUT structure.
type KeyCreateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyCreateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeyCreateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewKeyCreateUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKeyCreateUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /kms/key] keyCreateUsingPUT", response, response.Code())
	}
}

// NewKeyCreateUsingPUTOK creates a KeyCreateUsingPUTOK with default headers values
func NewKeyCreateUsingPUTOK() *KeyCreateUsingPUTOK {
	return &KeyCreateUsingPUTOK{}
}

/*
KeyCreateUsingPUTOK describes a response with status code 200, with default header values.

OK
*/
type KeyCreateUsingPUTOK struct {
	Payload *models.CreateKmsKeyResponse
}

// IsSuccess returns true when this key create using p u t o k response has a 2xx status code
func (o *KeyCreateUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this key create using p u t o k response has a 3xx status code
func (o *KeyCreateUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this key create using p u t o k response has a 4xx status code
func (o *KeyCreateUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this key create using p u t o k response has a 5xx status code
func (o *KeyCreateUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this key create using p u t o k response a status code equal to that given
func (o *KeyCreateUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the key create using p u t o k response
func (o *KeyCreateUsingPUTOK) Code() int {
	return 200
}

func (o *KeyCreateUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /kms/key][%d] keyCreateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *KeyCreateUsingPUTOK) String() string {
	return fmt.Sprintf("[PUT /kms/key][%d] keyCreateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *KeyCreateUsingPUTOK) GetPayload() *models.CreateKmsKeyResponse {
	return o.Payload
}

func (o *KeyCreateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateKmsKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyCreateUsingPUTCreated creates a KeyCreateUsingPUTCreated with default headers values
func NewKeyCreateUsingPUTCreated() *KeyCreateUsingPUTCreated {
	return &KeyCreateUsingPUTCreated{}
}

/*
KeyCreateUsingPUTCreated describes a response with status code 201, with default header values.

Entity has been created
*/
type KeyCreateUsingPUTCreated struct {
	Payload *models.CreateKmsKeyResponse
}

// IsSuccess returns true when this key create using p u t created response has a 2xx status code
func (o *KeyCreateUsingPUTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this key create using p u t created response has a 3xx status code
func (o *KeyCreateUsingPUTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this key create using p u t created response has a 4xx status code
func (o *KeyCreateUsingPUTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this key create using p u t created response has a 5xx status code
func (o *KeyCreateUsingPUTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this key create using p u t created response a status code equal to that given
func (o *KeyCreateUsingPUTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the key create using p u t created response
func (o *KeyCreateUsingPUTCreated) Code() int {
	return 201
}

func (o *KeyCreateUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /kms/key][%d] keyCreateUsingPUTCreated  %+v", 201, o.Payload)
}

func (o *KeyCreateUsingPUTCreated) String() string {
	return fmt.Sprintf("[PUT /kms/key][%d] keyCreateUsingPUTCreated  %+v", 201, o.Payload)
}

func (o *KeyCreateUsingPUTCreated) GetPayload() *models.CreateKmsKeyResponse {
	return o.Payload
}

func (o *KeyCreateUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateKmsKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyCreateUsingPUTBadRequest creates a KeyCreateUsingPUTBadRequest with default headers values
func NewKeyCreateUsingPUTBadRequest() *KeyCreateUsingPUTBadRequest {
	return &KeyCreateUsingPUTBadRequest{}
}

/*
KeyCreateUsingPUTBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type KeyCreateUsingPUTBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this key create using p u t bad request response has a 2xx status code
func (o *KeyCreateUsingPUTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this key create using p u t bad request response has a 3xx status code
func (o *KeyCreateUsingPUTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this key create using p u t bad request response has a 4xx status code
func (o *KeyCreateUsingPUTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this key create using p u t bad request response has a 5xx status code
func (o *KeyCreateUsingPUTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this key create using p u t bad request response a status code equal to that given
func (o *KeyCreateUsingPUTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the key create using p u t bad request response
func (o *KeyCreateUsingPUTBadRequest) Code() int {
	return 400
}

func (o *KeyCreateUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /kms/key][%d] keyCreateUsingPUTBadRequest  %+v", 400, o.Payload)
}

func (o *KeyCreateUsingPUTBadRequest) String() string {
	return fmt.Sprintf("[PUT /kms/key][%d] keyCreateUsingPUTBadRequest  %+v", 400, o.Payload)
}

func (o *KeyCreateUsingPUTBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *KeyCreateUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
