// Code generated by go-swagger; DO NOT EDIT.

package active_directory_public_key

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// GetPublicKeyUsingGETReader is a Reader for the GetPublicKeyUsingGET structure.
type GetPublicKeyUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicKeyUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPublicKeyUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPublicKeyUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPublicKeyUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicKeyUsingGETOK creates a GetPublicKeyUsingGETOK with default headers values
func NewGetPublicKeyUsingGETOK() *GetPublicKeyUsingGETOK {
	return &GetPublicKeyUsingGETOK{}
}

/*GetPublicKeyUsingGETOK handles this case with default header values.

OK
*/
type GetPublicKeyUsingGETOK struct {
	Payload *models.PublicKeyInstanceGetResponse
}

func (o *GetPublicKeyUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /ad/integration/pk][%d] getPublicKeyUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetPublicKeyUsingGETOK) GetPayload() *models.PublicKeyInstanceGetResponse {
	return o.Payload
}

func (o *GetPublicKeyUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicKeyInstanceGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPublicKeyUsingGETBadRequest creates a GetPublicKeyUsingGETBadRequest with default headers values
func NewGetPublicKeyUsingGETBadRequest() *GetPublicKeyUsingGETBadRequest {
	return &GetPublicKeyUsingGETBadRequest{}
}

/*GetPublicKeyUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type GetPublicKeyUsingGETBadRequest struct {
}

func (o *GetPublicKeyUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /ad/integration/pk][%d] getPublicKeyUsingGETBadRequest ", 400)
}

func (o *GetPublicKeyUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPublicKeyUsingGETNotFound creates a GetPublicKeyUsingGETNotFound with default headers values
func NewGetPublicKeyUsingGETNotFound() *GetPublicKeyUsingGETNotFound {
	return &GetPublicKeyUsingGETNotFound{}
}

/*GetPublicKeyUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type GetPublicKeyUsingGETNotFound struct {
}

func (o *GetPublicKeyUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /ad/integration/pk][%d] getPublicKeyUsingGETNotFound ", 404)
}

func (o *GetPublicKeyUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}