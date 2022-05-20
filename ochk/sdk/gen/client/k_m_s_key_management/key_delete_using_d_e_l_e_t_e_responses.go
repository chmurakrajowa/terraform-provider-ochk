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

// KeyDeleteUsingDELETEReader is a Reader for the KeyDeleteUsingDELETE structure.
type KeyDeleteUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyDeleteUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeyDeleteUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKeyDeleteUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKeyDeleteUsingDELETEOK creates a KeyDeleteUsingDELETEOK with default headers values
func NewKeyDeleteUsingDELETEOK() *KeyDeleteUsingDELETEOK {
	return &KeyDeleteUsingDELETEOK{}
}

/* KeyDeleteUsingDELETEOK describes a response with status code 200, with default header values.

Entity has been deleted
*/
type KeyDeleteUsingDELETEOK struct {
	Payload *models.DeleteKmsKeyResponse
}

func (o *KeyDeleteUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /kms/key/{id}][%d] keyDeleteUsingDELETEOK  %+v", 200, o.Payload)
}
func (o *KeyDeleteUsingDELETEOK) GetPayload() *models.DeleteKmsKeyResponse {
	return o.Payload
}

func (o *KeyDeleteUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteKmsKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyDeleteUsingDELETEBadRequest creates a KeyDeleteUsingDELETEBadRequest with default headers values
func NewKeyDeleteUsingDELETEBadRequest() *KeyDeleteUsingDELETEBadRequest {
	return &KeyDeleteUsingDELETEBadRequest{}
}

/* KeyDeleteUsingDELETEBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type KeyDeleteUsingDELETEBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *KeyDeleteUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /kms/key/{id}][%d] keyDeleteUsingDELETEBadRequest  %+v", 400, o.Payload)
}
func (o *KeyDeleteUsingDELETEBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *KeyDeleteUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
