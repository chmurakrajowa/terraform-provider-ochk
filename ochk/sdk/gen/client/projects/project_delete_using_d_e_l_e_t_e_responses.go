// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// ProjectDeleteUsingDELETEReader is a Reader for the ProjectDeleteUsingDELETE structure.
type ProjectDeleteUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectDeleteUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectDeleteUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewProjectDeleteUsingDELETECreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectDeleteUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /projects/{projectId}] projectDeleteUsingDELETE", response, response.Code())
	}
}

// NewProjectDeleteUsingDELETEOK creates a ProjectDeleteUsingDELETEOK with default headers values
func NewProjectDeleteUsingDELETEOK() *ProjectDeleteUsingDELETEOK {
	return &ProjectDeleteUsingDELETEOK{}
}

/*
ProjectDeleteUsingDELETEOK describes a response with status code 200, with default header values.

OK
*/
type ProjectDeleteUsingDELETEOK struct {
	Payload *models.ProjectDeleteResponse
}

// IsSuccess returns true when this project delete using d e l e t e o k response has a 2xx status code
func (o *ProjectDeleteUsingDELETEOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project delete using d e l e t e o k response has a 3xx status code
func (o *ProjectDeleteUsingDELETEOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project delete using d e l e t e o k response has a 4xx status code
func (o *ProjectDeleteUsingDELETEOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project delete using d e l e t e o k response has a 5xx status code
func (o *ProjectDeleteUsingDELETEOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project delete using d e l e t e o k response a status code equal to that given
func (o *ProjectDeleteUsingDELETEOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the project delete using d e l e t e o k response
func (o *ProjectDeleteUsingDELETEOK) Code() int {
	return 200
}

func (o *ProjectDeleteUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{projectId}][%d] projectDeleteUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *ProjectDeleteUsingDELETEOK) String() string {
	return fmt.Sprintf("[DELETE /projects/{projectId}][%d] projectDeleteUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *ProjectDeleteUsingDELETEOK) GetPayload() *models.ProjectDeleteResponse {
	return o.Payload
}

func (o *ProjectDeleteUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectDeleteUsingDELETECreated creates a ProjectDeleteUsingDELETECreated with default headers values
func NewProjectDeleteUsingDELETECreated() *ProjectDeleteUsingDELETECreated {
	return &ProjectDeleteUsingDELETECreated{}
}

/*
ProjectDeleteUsingDELETECreated describes a response with status code 201, with default header values.

Entity has been deleted
*/
type ProjectDeleteUsingDELETECreated struct {
	Payload *models.ProjectDeleteResponse
}

// IsSuccess returns true when this project delete using d e l e t e created response has a 2xx status code
func (o *ProjectDeleteUsingDELETECreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project delete using d e l e t e created response has a 3xx status code
func (o *ProjectDeleteUsingDELETECreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project delete using d e l e t e created response has a 4xx status code
func (o *ProjectDeleteUsingDELETECreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this project delete using d e l e t e created response has a 5xx status code
func (o *ProjectDeleteUsingDELETECreated) IsServerError() bool {
	return false
}

// IsCode returns true when this project delete using d e l e t e created response a status code equal to that given
func (o *ProjectDeleteUsingDELETECreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the project delete using d e l e t e created response
func (o *ProjectDeleteUsingDELETECreated) Code() int {
	return 201
}

func (o *ProjectDeleteUsingDELETECreated) Error() string {
	return fmt.Sprintf("[DELETE /projects/{projectId}][%d] projectDeleteUsingDELETECreated  %+v", 201, o.Payload)
}

func (o *ProjectDeleteUsingDELETECreated) String() string {
	return fmt.Sprintf("[DELETE /projects/{projectId}][%d] projectDeleteUsingDELETECreated  %+v", 201, o.Payload)
}

func (o *ProjectDeleteUsingDELETECreated) GetPayload() *models.ProjectDeleteResponse {
	return o.Payload
}

func (o *ProjectDeleteUsingDELETECreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectDeleteUsingDELETEBadRequest creates a ProjectDeleteUsingDELETEBadRequest with default headers values
func NewProjectDeleteUsingDELETEBadRequest() *ProjectDeleteUsingDELETEBadRequest {
	return &ProjectDeleteUsingDELETEBadRequest{}
}

/*
ProjectDeleteUsingDELETEBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type ProjectDeleteUsingDELETEBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this project delete using d e l e t e bad request response has a 2xx status code
func (o *ProjectDeleteUsingDELETEBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project delete using d e l e t e bad request response has a 3xx status code
func (o *ProjectDeleteUsingDELETEBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project delete using d e l e t e bad request response has a 4xx status code
func (o *ProjectDeleteUsingDELETEBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this project delete using d e l e t e bad request response has a 5xx status code
func (o *ProjectDeleteUsingDELETEBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this project delete using d e l e t e bad request response a status code equal to that given
func (o *ProjectDeleteUsingDELETEBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the project delete using d e l e t e bad request response
func (o *ProjectDeleteUsingDELETEBadRequest) Code() int {
	return 400
}

func (o *ProjectDeleteUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /projects/{projectId}][%d] projectDeleteUsingDELETEBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectDeleteUsingDELETEBadRequest) String() string {
	return fmt.Sprintf("[DELETE /projects/{projectId}][%d] projectDeleteUsingDELETEBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectDeleteUsingDELETEBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *ProjectDeleteUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
