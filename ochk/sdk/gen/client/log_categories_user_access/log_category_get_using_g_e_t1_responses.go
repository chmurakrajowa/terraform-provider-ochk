// Code generated by go-swagger; DO NOT EDIT.

package log_categories_user_access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// LogCategoryGetUsingGET1Reader is a Reader for the LogCategoryGetUsingGET1 structure.
type LogCategoryGetUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogCategoryGetUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogCategoryGetUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLogCategoryGetUsingGET1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLogCategoryGetUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /log/user/categories/{logCategoryId}] logCategoryGetUsingGET_1", response, response.Code())
	}
}

// NewLogCategoryGetUsingGET1OK creates a LogCategoryGetUsingGET1OK with default headers values
func NewLogCategoryGetUsingGET1OK() *LogCategoryGetUsingGET1OK {
	return &LogCategoryGetUsingGET1OK{}
}

/*
LogCategoryGetUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type LogCategoryGetUsingGET1OK struct {
	Payload *models.LogCategoryGetResponse
}

// IsSuccess returns true when this log category get using g e t1 o k response has a 2xx status code
func (o *LogCategoryGetUsingGET1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this log category get using g e t1 o k response has a 3xx status code
func (o *LogCategoryGetUsingGET1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log category get using g e t1 o k response has a 4xx status code
func (o *LogCategoryGetUsingGET1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this log category get using g e t1 o k response has a 5xx status code
func (o *LogCategoryGetUsingGET1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this log category get using g e t1 o k response a status code equal to that given
func (o *LogCategoryGetUsingGET1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the log category get using g e t1 o k response
func (o *LogCategoryGetUsingGET1OK) Code() int {
	return 200
}

func (o *LogCategoryGetUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /log/user/categories/{logCategoryId}][%d] logCategoryGetUsingGET1OK  %+v", 200, o.Payload)
}

func (o *LogCategoryGetUsingGET1OK) String() string {
	return fmt.Sprintf("[GET /log/user/categories/{logCategoryId}][%d] logCategoryGetUsingGET1OK  %+v", 200, o.Payload)
}

func (o *LogCategoryGetUsingGET1OK) GetPayload() *models.LogCategoryGetResponse {
	return o.Payload
}

func (o *LogCategoryGetUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogCategoryGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogCategoryGetUsingGET1BadRequest creates a LogCategoryGetUsingGET1BadRequest with default headers values
func NewLogCategoryGetUsingGET1BadRequest() *LogCategoryGetUsingGET1BadRequest {
	return &LogCategoryGetUsingGET1BadRequest{}
}

/*
LogCategoryGetUsingGET1BadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type LogCategoryGetUsingGET1BadRequest struct {
}

// IsSuccess returns true when this log category get using g e t1 bad request response has a 2xx status code
func (o *LogCategoryGetUsingGET1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this log category get using g e t1 bad request response has a 3xx status code
func (o *LogCategoryGetUsingGET1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log category get using g e t1 bad request response has a 4xx status code
func (o *LogCategoryGetUsingGET1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this log category get using g e t1 bad request response has a 5xx status code
func (o *LogCategoryGetUsingGET1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this log category get using g e t1 bad request response a status code equal to that given
func (o *LogCategoryGetUsingGET1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the log category get using g e t1 bad request response
func (o *LogCategoryGetUsingGET1BadRequest) Code() int {
	return 400
}

func (o *LogCategoryGetUsingGET1BadRequest) Error() string {
	return fmt.Sprintf("[GET /log/user/categories/{logCategoryId}][%d] logCategoryGetUsingGET1BadRequest ", 400)
}

func (o *LogCategoryGetUsingGET1BadRequest) String() string {
	return fmt.Sprintf("[GET /log/user/categories/{logCategoryId}][%d] logCategoryGetUsingGET1BadRequest ", 400)
}

func (o *LogCategoryGetUsingGET1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLogCategoryGetUsingGET1NotFound creates a LogCategoryGetUsingGET1NotFound with default headers values
func NewLogCategoryGetUsingGET1NotFound() *LogCategoryGetUsingGET1NotFound {
	return &LogCategoryGetUsingGET1NotFound{}
}

/*
LogCategoryGetUsingGET1NotFound describes a response with status code 404, with default header values.

Entity not found.
*/
type LogCategoryGetUsingGET1NotFound struct {
}

// IsSuccess returns true when this log category get using g e t1 not found response has a 2xx status code
func (o *LogCategoryGetUsingGET1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this log category get using g e t1 not found response has a 3xx status code
func (o *LogCategoryGetUsingGET1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log category get using g e t1 not found response has a 4xx status code
func (o *LogCategoryGetUsingGET1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this log category get using g e t1 not found response has a 5xx status code
func (o *LogCategoryGetUsingGET1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this log category get using g e t1 not found response a status code equal to that given
func (o *LogCategoryGetUsingGET1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the log category get using g e t1 not found response
func (o *LogCategoryGetUsingGET1NotFound) Code() int {
	return 404
}

func (o *LogCategoryGetUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /log/user/categories/{logCategoryId}][%d] logCategoryGetUsingGET1NotFound ", 404)
}

func (o *LogCategoryGetUsingGET1NotFound) String() string {
	return fmt.Sprintf("[GET /log/user/categories/{logCategoryId}][%d] logCategoryGetUsingGET1NotFound ", 404)
}

func (o *LogCategoryGetUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
