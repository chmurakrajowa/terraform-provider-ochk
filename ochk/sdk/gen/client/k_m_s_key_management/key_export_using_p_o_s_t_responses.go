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

// KeyExportUsingPOSTReader is a Reader for the KeyExportUsingPOST structure.
type KeyExportUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyExportUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeyExportUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKeyExportUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /kms/key/{id}/export] keyExportUsingPOST", response, response.Code())
	}
}

// NewKeyExportUsingPOSTOK creates a KeyExportUsingPOSTOK with default headers values
func NewKeyExportUsingPOSTOK() *KeyExportUsingPOSTOK {
	return &KeyExportUsingPOSTOK{}
}

/*
KeyExportUsingPOSTOK describes a response with status code 200, with default header values.

Key has been exported successfully.
*/
type KeyExportUsingPOSTOK struct {
	Payload *models.ExportKmsKeyResponse
}

// IsSuccess returns true when this key export using p o s t o k response has a 2xx status code
func (o *KeyExportUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this key export using p o s t o k response has a 3xx status code
func (o *KeyExportUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this key export using p o s t o k response has a 4xx status code
func (o *KeyExportUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this key export using p o s t o k response has a 5xx status code
func (o *KeyExportUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this key export using p o s t o k response a status code equal to that given
func (o *KeyExportUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the key export using p o s t o k response
func (o *KeyExportUsingPOSTOK) Code() int {
	return 200
}

func (o *KeyExportUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /kms/key/{id}/export][%d] keyExportUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *KeyExportUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /kms/key/{id}/export][%d] keyExportUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *KeyExportUsingPOSTOK) GetPayload() *models.ExportKmsKeyResponse {
	return o.Payload
}

func (o *KeyExportUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportKmsKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyExportUsingPOSTBadRequest creates a KeyExportUsingPOSTBadRequest with default headers values
func NewKeyExportUsingPOSTBadRequest() *KeyExportUsingPOSTBadRequest {
	return &KeyExportUsingPOSTBadRequest{}
}

/*
KeyExportUsingPOSTBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type KeyExportUsingPOSTBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this key export using p o s t bad request response has a 2xx status code
func (o *KeyExportUsingPOSTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this key export using p o s t bad request response has a 3xx status code
func (o *KeyExportUsingPOSTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this key export using p o s t bad request response has a 4xx status code
func (o *KeyExportUsingPOSTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this key export using p o s t bad request response has a 5xx status code
func (o *KeyExportUsingPOSTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this key export using p o s t bad request response a status code equal to that given
func (o *KeyExportUsingPOSTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the key export using p o s t bad request response
func (o *KeyExportUsingPOSTBadRequest) Code() int {
	return 400
}

func (o *KeyExportUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /kms/key/{id}/export][%d] keyExportUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *KeyExportUsingPOSTBadRequest) String() string {
	return fmt.Sprintf("[POST /kms/key/{id}/export][%d] keyExportUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *KeyExportUsingPOSTBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *KeyExportUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
