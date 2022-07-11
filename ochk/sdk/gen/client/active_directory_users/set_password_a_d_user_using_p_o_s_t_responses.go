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

// SetPasswordADUserUsingPOSTReader is a Reader for the SetPasswordADUserUsingPOST structure.
type SetPasswordADUserUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetPasswordADUserUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetPasswordADUserUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetPasswordADUserUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetPasswordADUserUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetPasswordADUserUsingPOSTOK creates a SetPasswordADUserUsingPOSTOK with default headers values
func NewSetPasswordADUserUsingPOSTOK() *SetPasswordADUserUsingPOSTOK {
	return &SetPasswordADUserUsingPOSTOK{}
}

/* SetPasswordADUserUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type SetPasswordADUserUsingPOSTOK struct {
	Payload *models.SetUserInstancePasswordResponse
}

func (o *SetPasswordADUserUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /ad/integration/users/{samAccountName}/setPassword][%d] setPasswordADUserUsingPOSTOK  %+v", 200, o.Payload)
}
func (o *SetPasswordADUserUsingPOSTOK) GetPayload() *models.SetUserInstancePasswordResponse {
	return o.Payload
}

func (o *SetPasswordADUserUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SetUserInstancePasswordResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetPasswordADUserUsingPOSTBadRequest creates a SetPasswordADUserUsingPOSTBadRequest with default headers values
func NewSetPasswordADUserUsingPOSTBadRequest() *SetPasswordADUserUsingPOSTBadRequest {
	return &SetPasswordADUserUsingPOSTBadRequest{}
}

/* SetPasswordADUserUsingPOSTBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type SetPasswordADUserUsingPOSTBadRequest struct {
}

func (o *SetPasswordADUserUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /ad/integration/users/{samAccountName}/setPassword][%d] setPasswordADUserUsingPOSTBadRequest ", 400)
}

func (o *SetPasswordADUserUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetPasswordADUserUsingPOSTNotFound creates a SetPasswordADUserUsingPOSTNotFound with default headers values
func NewSetPasswordADUserUsingPOSTNotFound() *SetPasswordADUserUsingPOSTNotFound {
	return &SetPasswordADUserUsingPOSTNotFound{}
}

/* SetPasswordADUserUsingPOSTNotFound describes a response with status code 404, with default header values.

Entity not found.
*/
type SetPasswordADUserUsingPOSTNotFound struct {
}

func (o *SetPasswordADUserUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /ad/integration/users/{samAccountName}/setPassword][%d] setPasswordADUserUsingPOSTNotFound ", 404)
}

func (o *SetPasswordADUserUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
