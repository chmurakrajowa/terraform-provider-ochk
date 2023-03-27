// Code generated by go-swagger; DO NOT EDIT.

package requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// RequestGetUsingGETReader is a Reader for the RequestGetUsingGET structure.
type RequestGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RequestGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRequestGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRequestGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRequestGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRequestGetUsingGETOK creates a RequestGetUsingGETOK with default headers values
func NewRequestGetUsingGETOK() *RequestGetUsingGETOK {
	return &RequestGetUsingGETOK{}
}

/*
RequestGetUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type RequestGetUsingGETOK struct {
	Payload *models.RequestInstanceGetResponse
}

// IsSuccess returns true when this request get using g e t o k response has a 2xx status code
func (o *RequestGetUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this request get using g e t o k response has a 3xx status code
func (o *RequestGetUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this request get using g e t o k response has a 4xx status code
func (o *RequestGetUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this request get using g e t o k response has a 5xx status code
func (o *RequestGetUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this request get using g e t o k response a status code equal to that given
func (o *RequestGetUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the request get using g e t o k response
func (o *RequestGetUsingGETOK) Code() int {
	return 200
}

func (o *RequestGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /request/{requestId}][%d] requestGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *RequestGetUsingGETOK) String() string {
	return fmt.Sprintf("[GET /request/{requestId}][%d] requestGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *RequestGetUsingGETOK) GetPayload() *models.RequestInstanceGetResponse {
	return o.Payload
}

func (o *RequestGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestInstanceGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestGetUsingGETBadRequest creates a RequestGetUsingGETBadRequest with default headers values
func NewRequestGetUsingGETBadRequest() *RequestGetUsingGETBadRequest {
	return &RequestGetUsingGETBadRequest{}
}

/*
RequestGetUsingGETBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type RequestGetUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this request get using g e t bad request response has a 2xx status code
func (o *RequestGetUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this request get using g e t bad request response has a 3xx status code
func (o *RequestGetUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this request get using g e t bad request response has a 4xx status code
func (o *RequestGetUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this request get using g e t bad request response has a 5xx status code
func (o *RequestGetUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this request get using g e t bad request response a status code equal to that given
func (o *RequestGetUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the request get using g e t bad request response
func (o *RequestGetUsingGETBadRequest) Code() int {
	return 400
}

func (o *RequestGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /request/{requestId}][%d] requestGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *RequestGetUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /request/{requestId}][%d] requestGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *RequestGetUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *RequestGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestGetUsingGETNotFound creates a RequestGetUsingGETNotFound with default headers values
func NewRequestGetUsingGETNotFound() *RequestGetUsingGETNotFound {
	return &RequestGetUsingGETNotFound{}
}

/*
RequestGetUsingGETNotFound describes a response with status code 404, with default header values.

Entity not found.
*/
type RequestGetUsingGETNotFound struct {
}

// IsSuccess returns true when this request get using g e t not found response has a 2xx status code
func (o *RequestGetUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this request get using g e t not found response has a 3xx status code
func (o *RequestGetUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this request get using g e t not found response has a 4xx status code
func (o *RequestGetUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this request get using g e t not found response has a 5xx status code
func (o *RequestGetUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this request get using g e t not found response a status code equal to that given
func (o *RequestGetUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the request get using g e t not found response
func (o *RequestGetUsingGETNotFound) Code() int {
	return 404
}

func (o *RequestGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /request/{requestId}][%d] requestGetUsingGETNotFound ", 404)
}

func (o *RequestGetUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /request/{requestId}][%d] requestGetUsingGETNotFound ", 404)
}

func (o *RequestGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
