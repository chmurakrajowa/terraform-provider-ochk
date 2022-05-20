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

// CreateADUserUsingPUTReader is a Reader for the CreateADUserUsingPUT structure.
type CreateADUserUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateADUserUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateADUserUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateADUserUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateADUserUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateADUserUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateADUserUsingPUTOK creates a CreateADUserUsingPUTOK with default headers values
func NewCreateADUserUsingPUTOK() *CreateADUserUsingPUTOK {
	return &CreateADUserUsingPUTOK{}
}

/*CreateADUserUsingPUTOK handles this case with default header values.

OK
*/
type CreateADUserUsingPUTOK struct {
	Payload *models.CreateUserInstanceResponse
}

func (o *CreateADUserUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /ad/integration/users][%d] createADUserUsingPUTOK  %+v", 200, o.Payload)
}

func (o *CreateADUserUsingPUTOK) GetPayload() *models.CreateUserInstanceResponse {
	return o.Payload
}

func (o *CreateADUserUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateUserInstanceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateADUserUsingPUTCreated creates a CreateADUserUsingPUTCreated with default headers values
func NewCreateADUserUsingPUTCreated() *CreateADUserUsingPUTCreated {
	return &CreateADUserUsingPUTCreated{}
}

/*CreateADUserUsingPUTCreated handles this case with default header values.

Entity has been created
*/
type CreateADUserUsingPUTCreated struct {
	Payload *models.CreateUserInstanceResponse
}

func (o *CreateADUserUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /ad/integration/users][%d] createADUserUsingPUTCreated  %+v", 201, o.Payload)
}

func (o *CreateADUserUsingPUTCreated) GetPayload() *models.CreateUserInstanceResponse {
	return o.Payload
}

func (o *CreateADUserUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateUserInstanceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateADUserUsingPUTBadRequest creates a CreateADUserUsingPUTBadRequest with default headers values
func NewCreateADUserUsingPUTBadRequest() *CreateADUserUsingPUTBadRequest {
	return &CreateADUserUsingPUTBadRequest{}
}

/*CreateADUserUsingPUTBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type CreateADUserUsingPUTBadRequest struct {
}

func (o *CreateADUserUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /ad/integration/users][%d] createADUserUsingPUTBadRequest ", 400)
}

func (o *CreateADUserUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateADUserUsingPUTNotFound creates a CreateADUserUsingPUTNotFound with default headers values
func NewCreateADUserUsingPUTNotFound() *CreateADUserUsingPUTNotFound {
	return &CreateADUserUsingPUTNotFound{}
}

/*CreateADUserUsingPUTNotFound handles this case with default header values.

Entity not found.
*/
type CreateADUserUsingPUTNotFound struct {
}

func (o *CreateADUserUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /ad/integration/users][%d] createADUserUsingPUTNotFound ", 404)
}

func (o *CreateADUserUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
