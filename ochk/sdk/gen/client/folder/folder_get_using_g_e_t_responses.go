// Code generated by go-swagger; DO NOT EDIT.

package folder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// FolderGetUsingGETReader is a Reader for the FolderGetUsingGET structure.
type FolderGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FolderGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFolderGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFolderGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFolderGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /folder/{projectId}/{id}] folderGetUsingGET", response, response.Code())
	}
}

// NewFolderGetUsingGETOK creates a FolderGetUsingGETOK with default headers values
func NewFolderGetUsingGETOK() *FolderGetUsingGETOK {
	return &FolderGetUsingGETOK{}
}

/*
FolderGetUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type FolderGetUsingGETOK struct {
	Payload *models.FolderGetResponse
}

// IsSuccess returns true when this folder get using g e t o k response has a 2xx status code
func (o *FolderGetUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this folder get using g e t o k response has a 3xx status code
func (o *FolderGetUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this folder get using g e t o k response has a 4xx status code
func (o *FolderGetUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this folder get using g e t o k response has a 5xx status code
func (o *FolderGetUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this folder get using g e t o k response a status code equal to that given
func (o *FolderGetUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the folder get using g e t o k response
func (o *FolderGetUsingGETOK) Code() int {
	return 200
}

func (o *FolderGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /folder/{projectId}/{id}][%d] folderGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *FolderGetUsingGETOK) String() string {
	return fmt.Sprintf("[GET /folder/{projectId}/{id}][%d] folderGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *FolderGetUsingGETOK) GetPayload() *models.FolderGetResponse {
	return o.Payload
}

func (o *FolderGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FolderGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFolderGetUsingGETBadRequest creates a FolderGetUsingGETBadRequest with default headers values
func NewFolderGetUsingGETBadRequest() *FolderGetUsingGETBadRequest {
	return &FolderGetUsingGETBadRequest{}
}

/*
FolderGetUsingGETBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type FolderGetUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this folder get using g e t bad request response has a 2xx status code
func (o *FolderGetUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this folder get using g e t bad request response has a 3xx status code
func (o *FolderGetUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this folder get using g e t bad request response has a 4xx status code
func (o *FolderGetUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this folder get using g e t bad request response has a 5xx status code
func (o *FolderGetUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this folder get using g e t bad request response a status code equal to that given
func (o *FolderGetUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the folder get using g e t bad request response
func (o *FolderGetUsingGETBadRequest) Code() int {
	return 400
}

func (o *FolderGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /folder/{projectId}/{id}][%d] folderGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *FolderGetUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /folder/{projectId}/{id}][%d] folderGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *FolderGetUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *FolderGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFolderGetUsingGETNotFound creates a FolderGetUsingGETNotFound with default headers values
func NewFolderGetUsingGETNotFound() *FolderGetUsingGETNotFound {
	return &FolderGetUsingGETNotFound{}
}

/*
FolderGetUsingGETNotFound describes a response with status code 404, with default header values.

Entity not found.
*/
type FolderGetUsingGETNotFound struct {
}

// IsSuccess returns true when this folder get using g e t not found response has a 2xx status code
func (o *FolderGetUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this folder get using g e t not found response has a 3xx status code
func (o *FolderGetUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this folder get using g e t not found response has a 4xx status code
func (o *FolderGetUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this folder get using g e t not found response has a 5xx status code
func (o *FolderGetUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this folder get using g e t not found response a status code equal to that given
func (o *FolderGetUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the folder get using g e t not found response
func (o *FolderGetUsingGETNotFound) Code() int {
	return 404
}

func (o *FolderGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /folder/{projectId}/{id}][%d] folderGetUsingGETNotFound ", 404)
}

func (o *FolderGetUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /folder/{projectId}/{id}][%d] folderGetUsingGETNotFound ", 404)
}

func (o *FolderGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
