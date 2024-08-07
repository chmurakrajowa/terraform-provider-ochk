// Code generated by go-swagger; DO NOT EDIT.

package log_categories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// LogCategoryUpdateUsingPUTReader is a Reader for the LogCategoryUpdateUsingPUT structure.
type LogCategoryUpdateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogCategoryUpdateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogCategoryUpdateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLogCategoryUpdateUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /log/categories/{logCategoryId}] logCategoryUpdateUsingPUT", response, response.Code())
	}
}

// NewLogCategoryUpdateUsingPUTOK creates a LogCategoryUpdateUsingPUTOK with default headers values
func NewLogCategoryUpdateUsingPUTOK() *LogCategoryUpdateUsingPUTOK {
	return &LogCategoryUpdateUsingPUTOK{}
}

/*
LogCategoryUpdateUsingPUTOK describes a response with status code 200, with default header values.

Request has been completed successfully
*/
type LogCategoryUpdateUsingPUTOK struct {
	Payload *models.UpdateLogCategoryResponse
}

// IsSuccess returns true when this log category update using p u t o k response has a 2xx status code
func (o *LogCategoryUpdateUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this log category update using p u t o k response has a 3xx status code
func (o *LogCategoryUpdateUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log category update using p u t o k response has a 4xx status code
func (o *LogCategoryUpdateUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this log category update using p u t o k response has a 5xx status code
func (o *LogCategoryUpdateUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this log category update using p u t o k response a status code equal to that given
func (o *LogCategoryUpdateUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the log category update using p u t o k response
func (o *LogCategoryUpdateUsingPUTOK) Code() int {
	return 200
}

func (o *LogCategoryUpdateUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /log/categories/{logCategoryId}][%d] logCategoryUpdateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *LogCategoryUpdateUsingPUTOK) String() string {
	return fmt.Sprintf("[PUT /log/categories/{logCategoryId}][%d] logCategoryUpdateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *LogCategoryUpdateUsingPUTOK) GetPayload() *models.UpdateLogCategoryResponse {
	return o.Payload
}

func (o *LogCategoryUpdateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateLogCategoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogCategoryUpdateUsingPUTBadRequest creates a LogCategoryUpdateUsingPUTBadRequest with default headers values
func NewLogCategoryUpdateUsingPUTBadRequest() *LogCategoryUpdateUsingPUTBadRequest {
	return &LogCategoryUpdateUsingPUTBadRequest{}
}

/*
LogCategoryUpdateUsingPUTBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type LogCategoryUpdateUsingPUTBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this log category update using p u t bad request response has a 2xx status code
func (o *LogCategoryUpdateUsingPUTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this log category update using p u t bad request response has a 3xx status code
func (o *LogCategoryUpdateUsingPUTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log category update using p u t bad request response has a 4xx status code
func (o *LogCategoryUpdateUsingPUTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this log category update using p u t bad request response has a 5xx status code
func (o *LogCategoryUpdateUsingPUTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this log category update using p u t bad request response a status code equal to that given
func (o *LogCategoryUpdateUsingPUTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the log category update using p u t bad request response
func (o *LogCategoryUpdateUsingPUTBadRequest) Code() int {
	return 400
}

func (o *LogCategoryUpdateUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /log/categories/{logCategoryId}][%d] logCategoryUpdateUsingPUTBadRequest  %+v", 400, o.Payload)
}

func (o *LogCategoryUpdateUsingPUTBadRequest) String() string {
	return fmt.Sprintf("[PUT /log/categories/{logCategoryId}][%d] logCategoryUpdateUsingPUTBadRequest  %+v", 400, o.Payload)
}

func (o *LogCategoryUpdateUsingPUTBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *LogCategoryUpdateUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
