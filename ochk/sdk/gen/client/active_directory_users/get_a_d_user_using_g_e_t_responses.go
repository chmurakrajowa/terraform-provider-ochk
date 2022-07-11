// Code generated by go-swagger; DO NOT EDIT.

package active_directory_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// GetADUserUsingGETReader is a Reader for the GetADUserUsingGET structure.
type GetADUserUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetADUserUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetADUserUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetADUserUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetADUserUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetADUserUsingGETOK creates a GetADUserUsingGETOK with default headers values
func NewGetADUserUsingGETOK() *GetADUserUsingGETOK {
	return &GetADUserUsingGETOK{}
}

/* GetADUserUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetADUserUsingGETOK struct {
	Payload *models.UserInstanceGetResponse
}

func (o *GetADUserUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /ad/integration/users/{samAccountName}][%d] getADUserUsingGETOK  %+v", 200, o.Payload)
}
func (o *GetADUserUsingGETOK) GetPayload() *models.UserInstanceGetResponse {
	return o.Payload
}

func (o *GetADUserUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserInstanceGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetADUserUsingGETBadRequest creates a GetADUserUsingGETBadRequest with default headers values
func NewGetADUserUsingGETBadRequest() *GetADUserUsingGETBadRequest {
	return &GetADUserUsingGETBadRequest{}
}

/* GetADUserUsingGETBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type GetADUserUsingGETBadRequest struct {
}

func (o *GetADUserUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /ad/integration/users/{samAccountName}][%d] getADUserUsingGETBadRequest ", 400)
}

func (o *GetADUserUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetADUserUsingGETNotFound creates a GetADUserUsingGETNotFound with default headers values
func NewGetADUserUsingGETNotFound() *GetADUserUsingGETNotFound {
	return &GetADUserUsingGETNotFound{}
}

/* GetADUserUsingGETNotFound describes a response with status code 404, with default header values.

Entity not found.
*/
type GetADUserUsingGETNotFound struct {
}

func (o *GetADUserUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /ad/integration/users/{samAccountName}][%d] getADUserUsingGETNotFound ", 404)
}

func (o *GetADUserUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
