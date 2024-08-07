// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// TagGetUsingGETReader is a Reader for the TagGetUsingGET structure.
type TagGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TagGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTagGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTagGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTagGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /tags/{tagId}] tagGetUsingGET", response, response.Code())
	}
}

// NewTagGetUsingGETOK creates a TagGetUsingGETOK with default headers values
func NewTagGetUsingGETOK() *TagGetUsingGETOK {
	return &TagGetUsingGETOK{}
}

/*
TagGetUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type TagGetUsingGETOK struct {
	Payload *models.TagGetResponse
}

// IsSuccess returns true when this tag get using g e t o k response has a 2xx status code
func (o *TagGetUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tag get using g e t o k response has a 3xx status code
func (o *TagGetUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tag get using g e t o k response has a 4xx status code
func (o *TagGetUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tag get using g e t o k response has a 5xx status code
func (o *TagGetUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tag get using g e t o k response a status code equal to that given
func (o *TagGetUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tag get using g e t o k response
func (o *TagGetUsingGETOK) Code() int {
	return 200
}

func (o *TagGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /tags/{tagId}][%d] tagGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *TagGetUsingGETOK) String() string {
	return fmt.Sprintf("[GET /tags/{tagId}][%d] tagGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *TagGetUsingGETOK) GetPayload() *models.TagGetResponse {
	return o.Payload
}

func (o *TagGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTagGetUsingGETBadRequest creates a TagGetUsingGETBadRequest with default headers values
func NewTagGetUsingGETBadRequest() *TagGetUsingGETBadRequest {
	return &TagGetUsingGETBadRequest{}
}

/*
TagGetUsingGETBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type TagGetUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this tag get using g e t bad request response has a 2xx status code
func (o *TagGetUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tag get using g e t bad request response has a 3xx status code
func (o *TagGetUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tag get using g e t bad request response has a 4xx status code
func (o *TagGetUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this tag get using g e t bad request response has a 5xx status code
func (o *TagGetUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this tag get using g e t bad request response a status code equal to that given
func (o *TagGetUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the tag get using g e t bad request response
func (o *TagGetUsingGETBadRequest) Code() int {
	return 400
}

func (o *TagGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /tags/{tagId}][%d] tagGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *TagGetUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /tags/{tagId}][%d] tagGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *TagGetUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *TagGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTagGetUsingGETNotFound creates a TagGetUsingGETNotFound with default headers values
func NewTagGetUsingGETNotFound() *TagGetUsingGETNotFound {
	return &TagGetUsingGETNotFound{}
}

/*
TagGetUsingGETNotFound describes a response with status code 404, with default header values.

Entity not found.
*/
type TagGetUsingGETNotFound struct {
}

// IsSuccess returns true when this tag get using g e t not found response has a 2xx status code
func (o *TagGetUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tag get using g e t not found response has a 3xx status code
func (o *TagGetUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tag get using g e t not found response has a 4xx status code
func (o *TagGetUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this tag get using g e t not found response has a 5xx status code
func (o *TagGetUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this tag get using g e t not found response a status code equal to that given
func (o *TagGetUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the tag get using g e t not found response
func (o *TagGetUsingGETNotFound) Code() int {
	return 404
}

func (o *TagGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /tags/{tagId}][%d] tagGetUsingGETNotFound ", 404)
}

func (o *TagGetUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /tags/{tagId}][%d] tagGetUsingGETNotFound ", 404)
}

func (o *TagGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
