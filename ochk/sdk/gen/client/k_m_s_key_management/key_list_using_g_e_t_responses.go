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

// KeyListUsingGETReader is a Reader for the KeyListUsingGET structure.
type KeyListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeyListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKeyListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewKeyListUsingGETOK creates a KeyListUsingGETOK with default headers values
func NewKeyListUsingGETOK() *KeyListUsingGETOK {
	return &KeyListUsingGETOK{}
}

/*KeyListUsingGETOK handles this case with default header values.

OK
*/
type KeyListUsingGETOK struct {
	Payload *models.KeyListResponse
}

func (o *KeyListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /kms/key][%d] keyListUsingGETOK  %+v", 200, o.Payload)
}

func (o *KeyListUsingGETOK) GetPayload() *models.KeyListResponse {
	return o.Payload
}

func (o *KeyListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyListUsingGETBadRequest creates a KeyListUsingGETBadRequest with default headers values
func NewKeyListUsingGETBadRequest() *KeyListUsingGETBadRequest {
	return &KeyListUsingGETBadRequest{}
}

/*KeyListUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type KeyListUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *KeyListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /kms/key][%d] keyListUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *KeyListUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *KeyListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
