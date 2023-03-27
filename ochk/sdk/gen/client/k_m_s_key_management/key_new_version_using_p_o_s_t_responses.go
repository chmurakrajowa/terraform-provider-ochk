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

// KeyNewVersionUsingPOSTReader is a Reader for the KeyNewVersionUsingPOST structure.
type KeyNewVersionUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyNewVersionUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeyNewVersionUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewKeyNewVersionUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKeyNewVersionUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKeyNewVersionUsingPOSTOK creates a KeyNewVersionUsingPOSTOK with default headers values
func NewKeyNewVersionUsingPOSTOK() *KeyNewVersionUsingPOSTOK {
	return &KeyNewVersionUsingPOSTOK{}
}

/*
KeyNewVersionUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type KeyNewVersionUsingPOSTOK struct {
	Payload *models.CreateNewKmsKeyVersionResponse
}

// IsSuccess returns true when this key new version using p o s t o k response has a 2xx status code
func (o *KeyNewVersionUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this key new version using p o s t o k response has a 3xx status code
func (o *KeyNewVersionUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this key new version using p o s t o k response has a 4xx status code
func (o *KeyNewVersionUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this key new version using p o s t o k response has a 5xx status code
func (o *KeyNewVersionUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this key new version using p o s t o k response a status code equal to that given
func (o *KeyNewVersionUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the key new version using p o s t o k response
func (o *KeyNewVersionUsingPOSTOK) Code() int {
	return 200
}

func (o *KeyNewVersionUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /kms/key/{id}/versions][%d] keyNewVersionUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *KeyNewVersionUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /kms/key/{id}/versions][%d] keyNewVersionUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *KeyNewVersionUsingPOSTOK) GetPayload() *models.CreateNewKmsKeyVersionResponse {
	return o.Payload
}

func (o *KeyNewVersionUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateNewKmsKeyVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyNewVersionUsingPOSTCreated creates a KeyNewVersionUsingPOSTCreated with default headers values
func NewKeyNewVersionUsingPOSTCreated() *KeyNewVersionUsingPOSTCreated {
	return &KeyNewVersionUsingPOSTCreated{}
}

/*
KeyNewVersionUsingPOSTCreated describes a response with status code 201, with default header values.

New key version has been created.
*/
type KeyNewVersionUsingPOSTCreated struct {
	Payload *models.CreateNewKmsKeyVersionResponse
}

// IsSuccess returns true when this key new version using p o s t created response has a 2xx status code
func (o *KeyNewVersionUsingPOSTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this key new version using p o s t created response has a 3xx status code
func (o *KeyNewVersionUsingPOSTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this key new version using p o s t created response has a 4xx status code
func (o *KeyNewVersionUsingPOSTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this key new version using p o s t created response has a 5xx status code
func (o *KeyNewVersionUsingPOSTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this key new version using p o s t created response a status code equal to that given
func (o *KeyNewVersionUsingPOSTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the key new version using p o s t created response
func (o *KeyNewVersionUsingPOSTCreated) Code() int {
	return 201
}

func (o *KeyNewVersionUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /kms/key/{id}/versions][%d] keyNewVersionUsingPOSTCreated  %+v", 201, o.Payload)
}

func (o *KeyNewVersionUsingPOSTCreated) String() string {
	return fmt.Sprintf("[POST /kms/key/{id}/versions][%d] keyNewVersionUsingPOSTCreated  %+v", 201, o.Payload)
}

func (o *KeyNewVersionUsingPOSTCreated) GetPayload() *models.CreateNewKmsKeyVersionResponse {
	return o.Payload
}

func (o *KeyNewVersionUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateNewKmsKeyVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyNewVersionUsingPOSTBadRequest creates a KeyNewVersionUsingPOSTBadRequest with default headers values
func NewKeyNewVersionUsingPOSTBadRequest() *KeyNewVersionUsingPOSTBadRequest {
	return &KeyNewVersionUsingPOSTBadRequest{}
}

/*
KeyNewVersionUsingPOSTBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type KeyNewVersionUsingPOSTBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this key new version using p o s t bad request response has a 2xx status code
func (o *KeyNewVersionUsingPOSTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this key new version using p o s t bad request response has a 3xx status code
func (o *KeyNewVersionUsingPOSTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this key new version using p o s t bad request response has a 4xx status code
func (o *KeyNewVersionUsingPOSTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this key new version using p o s t bad request response has a 5xx status code
func (o *KeyNewVersionUsingPOSTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this key new version using p o s t bad request response a status code equal to that given
func (o *KeyNewVersionUsingPOSTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the key new version using p o s t bad request response
func (o *KeyNewVersionUsingPOSTBadRequest) Code() int {
	return 400
}

func (o *KeyNewVersionUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /kms/key/{id}/versions][%d] keyNewVersionUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *KeyNewVersionUsingPOSTBadRequest) String() string {
	return fmt.Sprintf("[POST /kms/key/{id}/versions][%d] keyNewVersionUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *KeyNewVersionUsingPOSTBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *KeyNewVersionUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
