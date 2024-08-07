// Code generated by go-swagger; DO NOT EDIT.

package log_stats_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// GetLogsCountUsingPOST1Reader is a Reader for the GetLogsCountUsingPOST1 structure.
type GetLogsCountUsingPOST1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogsCountUsingPOST1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogsCountUsingPOST1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLogsCountUsingPOST1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /log/stats/user/{logCategoryId}] getLogsCountUsingPOST_1", response, response.Code())
	}
}

// NewGetLogsCountUsingPOST1OK creates a GetLogsCountUsingPOST1OK with default headers values
func NewGetLogsCountUsingPOST1OK() *GetLogsCountUsingPOST1OK {
	return &GetLogsCountUsingPOST1OK{}
}

/*
GetLogsCountUsingPOST1OK describes a response with status code 200, with default header values.

OK
*/
type GetLogsCountUsingPOST1OK struct {
	Payload *models.GetLogCategoryStats
}

// IsSuccess returns true when this get logs count using p o s t1 o k response has a 2xx status code
func (o *GetLogsCountUsingPOST1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get logs count using p o s t1 o k response has a 3xx status code
func (o *GetLogsCountUsingPOST1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs count using p o s t1 o k response has a 4xx status code
func (o *GetLogsCountUsingPOST1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get logs count using p o s t1 o k response has a 5xx status code
func (o *GetLogsCountUsingPOST1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get logs count using p o s t1 o k response a status code equal to that given
func (o *GetLogsCountUsingPOST1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get logs count using p o s t1 o k response
func (o *GetLogsCountUsingPOST1OK) Code() int {
	return 200
}

func (o *GetLogsCountUsingPOST1OK) Error() string {
	return fmt.Sprintf("[POST /log/stats/user/{logCategoryId}][%d] getLogsCountUsingPOST1OK  %+v", 200, o.Payload)
}

func (o *GetLogsCountUsingPOST1OK) String() string {
	return fmt.Sprintf("[POST /log/stats/user/{logCategoryId}][%d] getLogsCountUsingPOST1OK  %+v", 200, o.Payload)
}

func (o *GetLogsCountUsingPOST1OK) GetPayload() *models.GetLogCategoryStats {
	return o.Payload
}

func (o *GetLogsCountUsingPOST1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetLogCategoryStats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsCountUsingPOST1BadRequest creates a GetLogsCountUsingPOST1BadRequest with default headers values
func NewGetLogsCountUsingPOST1BadRequest() *GetLogsCountUsingPOST1BadRequest {
	return &GetLogsCountUsingPOST1BadRequest{}
}

/*
GetLogsCountUsingPOST1BadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type GetLogsCountUsingPOST1BadRequest struct {
}

// IsSuccess returns true when this get logs count using p o s t1 bad request response has a 2xx status code
func (o *GetLogsCountUsingPOST1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get logs count using p o s t1 bad request response has a 3xx status code
func (o *GetLogsCountUsingPOST1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs count using p o s t1 bad request response has a 4xx status code
func (o *GetLogsCountUsingPOST1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get logs count using p o s t1 bad request response has a 5xx status code
func (o *GetLogsCountUsingPOST1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get logs count using p o s t1 bad request response a status code equal to that given
func (o *GetLogsCountUsingPOST1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get logs count using p o s t1 bad request response
func (o *GetLogsCountUsingPOST1BadRequest) Code() int {
	return 400
}

func (o *GetLogsCountUsingPOST1BadRequest) Error() string {
	return fmt.Sprintf("[POST /log/stats/user/{logCategoryId}][%d] getLogsCountUsingPOST1BadRequest ", 400)
}

func (o *GetLogsCountUsingPOST1BadRequest) String() string {
	return fmt.Sprintf("[POST /log/stats/user/{logCategoryId}][%d] getLogsCountUsingPOST1BadRequest ", 400)
}

func (o *GetLogsCountUsingPOST1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
