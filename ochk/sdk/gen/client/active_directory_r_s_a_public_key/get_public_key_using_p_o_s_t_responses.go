// Code generated by go-swagger; DO NOT EDIT.

package active_directory_r_s_a_public_key

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// GetPublicKeyUsingPOSTReader is a Reader for the GetPublicKeyUsingPOST structure.
type GetPublicKeyUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicKeyUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPublicKeyUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPublicKeyUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPublicKeyUsingPOSTOK creates a GetPublicKeyUsingPOSTOK with default headers values
func NewGetPublicKeyUsingPOSTOK() *GetPublicKeyUsingPOSTOK {
	return &GetPublicKeyUsingPOSTOK{}
}

/* GetPublicKeyUsingPOSTOK describes a response with status code 200, with default header values.

Entity has been imported
*/
type GetPublicKeyUsingPOSTOK struct {
	Payload *models.GetPublicKeyResponse
}

func (o *GetPublicKeyUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /ads/cer][%d] getPublicKeyUsingPOSTOK  %+v", 200, o.Payload)
}
func (o *GetPublicKeyUsingPOSTOK) GetPayload() *models.GetPublicKeyResponse {
	return o.Payload
}

func (o *GetPublicKeyUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetPublicKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPublicKeyUsingPOSTBadRequest creates a GetPublicKeyUsingPOSTBadRequest with default headers values
func NewGetPublicKeyUsingPOSTBadRequest() *GetPublicKeyUsingPOSTBadRequest {
	return &GetPublicKeyUsingPOSTBadRequest{}
}

/* GetPublicKeyUsingPOSTBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type GetPublicKeyUsingPOSTBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *GetPublicKeyUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /ads/cer][%d] getPublicKeyUsingPOSTBadRequest  %+v", 400, o.Payload)
}
func (o *GetPublicKeyUsingPOSTBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *GetPublicKeyUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
