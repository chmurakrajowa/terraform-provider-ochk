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

// KeyRevokeUsingPOSTReader is a Reader for the KeyRevokeUsingPOST structure.
type KeyRevokeUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyRevokeUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeyRevokeUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKeyRevokeUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /kms/key/{id}/revoke] keyRevokeUsingPOST", response, response.Code())
	}
}

// NewKeyRevokeUsingPOSTOK creates a KeyRevokeUsingPOSTOK with default headers values
func NewKeyRevokeUsingPOSTOK() *KeyRevokeUsingPOSTOK {
	return &KeyRevokeUsingPOSTOK{}
}

/*
KeyRevokeUsingPOSTOK describes a response with status code 200, with default header values.

Key has been revoked.
*/
type KeyRevokeUsingPOSTOK struct {
	Payload *models.RevokeKmsKeyResponse
}

// IsSuccess returns true when this key revoke using p o s t o k response has a 2xx status code
func (o *KeyRevokeUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this key revoke using p o s t o k response has a 3xx status code
func (o *KeyRevokeUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this key revoke using p o s t o k response has a 4xx status code
func (o *KeyRevokeUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this key revoke using p o s t o k response has a 5xx status code
func (o *KeyRevokeUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this key revoke using p o s t o k response a status code equal to that given
func (o *KeyRevokeUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the key revoke using p o s t o k response
func (o *KeyRevokeUsingPOSTOK) Code() int {
	return 200
}

func (o *KeyRevokeUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /kms/key/{id}/revoke][%d] keyRevokeUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *KeyRevokeUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /kms/key/{id}/revoke][%d] keyRevokeUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *KeyRevokeUsingPOSTOK) GetPayload() *models.RevokeKmsKeyResponse {
	return o.Payload
}

func (o *KeyRevokeUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RevokeKmsKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyRevokeUsingPOSTBadRequest creates a KeyRevokeUsingPOSTBadRequest with default headers values
func NewKeyRevokeUsingPOSTBadRequest() *KeyRevokeUsingPOSTBadRequest {
	return &KeyRevokeUsingPOSTBadRequest{}
}

/*
KeyRevokeUsingPOSTBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type KeyRevokeUsingPOSTBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this key revoke using p o s t bad request response has a 2xx status code
func (o *KeyRevokeUsingPOSTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this key revoke using p o s t bad request response has a 3xx status code
func (o *KeyRevokeUsingPOSTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this key revoke using p o s t bad request response has a 4xx status code
func (o *KeyRevokeUsingPOSTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this key revoke using p o s t bad request response has a 5xx status code
func (o *KeyRevokeUsingPOSTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this key revoke using p o s t bad request response a status code equal to that given
func (o *KeyRevokeUsingPOSTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the key revoke using p o s t bad request response
func (o *KeyRevokeUsingPOSTBadRequest) Code() int {
	return 400
}

func (o *KeyRevokeUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /kms/key/{id}/revoke][%d] keyRevokeUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *KeyRevokeUsingPOSTBadRequest) String() string {
	return fmt.Sprintf("[POST /kms/key/{id}/revoke][%d] keyRevokeUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *KeyRevokeUsingPOSTBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *KeyRevokeUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
