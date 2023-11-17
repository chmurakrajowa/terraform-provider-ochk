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

// TagListUsingGETReader is a Reader for the TagListUsingGET structure.
type TagListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TagListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTagListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTagListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /tags] tagListUsingGET", response, response.Code())
	}
}

// NewTagListUsingGETOK creates a TagListUsingGETOK with default headers values
func NewTagListUsingGETOK() *TagListUsingGETOK {
	return &TagListUsingGETOK{}
}

/*
TagListUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type TagListUsingGETOK struct {
	Payload *models.TagListResponse
}

// IsSuccess returns true when this tag list using g e t o k response has a 2xx status code
func (o *TagListUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tag list using g e t o k response has a 3xx status code
func (o *TagListUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tag list using g e t o k response has a 4xx status code
func (o *TagListUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tag list using g e t o k response has a 5xx status code
func (o *TagListUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tag list using g e t o k response a status code equal to that given
func (o *TagListUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tag list using g e t o k response
func (o *TagListUsingGETOK) Code() int {
	return 200
}

func (o *TagListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /tags][%d] tagListUsingGETOK  %+v", 200, o.Payload)
}

func (o *TagListUsingGETOK) String() string {
	return fmt.Sprintf("[GET /tags][%d] tagListUsingGETOK  %+v", 200, o.Payload)
}

func (o *TagListUsingGETOK) GetPayload() *models.TagListResponse {
	return o.Payload
}

func (o *TagListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTagListUsingGETBadRequest creates a TagListUsingGETBadRequest with default headers values
func NewTagListUsingGETBadRequest() *TagListUsingGETBadRequest {
	return &TagListUsingGETBadRequest{}
}

/*
TagListUsingGETBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type TagListUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this tag list using g e t bad request response has a 2xx status code
func (o *TagListUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tag list using g e t bad request response has a 3xx status code
func (o *TagListUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tag list using g e t bad request response has a 4xx status code
func (o *TagListUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this tag list using g e t bad request response has a 5xx status code
func (o *TagListUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this tag list using g e t bad request response a status code equal to that given
func (o *TagListUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the tag list using g e t bad request response
func (o *TagListUsingGETBadRequest) Code() int {
	return 400
}

func (o *TagListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /tags][%d] tagListUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *TagListUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /tags][%d] tagListUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *TagListUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *TagListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
