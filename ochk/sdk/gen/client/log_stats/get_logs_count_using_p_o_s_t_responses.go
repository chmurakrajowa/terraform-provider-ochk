// Code generated by go-swagger; DO NOT EDIT.

package log_stats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// GetLogsCountUsingPOSTReader is a Reader for the GetLogsCountUsingPOST structure.
type GetLogsCountUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogsCountUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogsCountUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLogsCountUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLogsCountUsingPOSTOK creates a GetLogsCountUsingPOSTOK with default headers values
func NewGetLogsCountUsingPOSTOK() *GetLogsCountUsingPOSTOK {
	return &GetLogsCountUsingPOSTOK{}
}

/*
GetLogsCountUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type GetLogsCountUsingPOSTOK struct {
	Payload *models.GetLogCategoryStats
}

// IsSuccess returns true when this get logs count using p o s t o k response has a 2xx status code
func (o *GetLogsCountUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get logs count using p o s t o k response has a 3xx status code
func (o *GetLogsCountUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs count using p o s t o k response has a 4xx status code
func (o *GetLogsCountUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get logs count using p o s t o k response has a 5xx status code
func (o *GetLogsCountUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get logs count using p o s t o k response a status code equal to that given
func (o *GetLogsCountUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get logs count using p o s t o k response
func (o *GetLogsCountUsingPOSTOK) Code() int {
	return 200
}

func (o *GetLogsCountUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /log/stats/{logCategoryId}][%d] getLogsCountUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *GetLogsCountUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /log/stats/{logCategoryId}][%d] getLogsCountUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *GetLogsCountUsingPOSTOK) GetPayload() *models.GetLogCategoryStats {
	return o.Payload
}

func (o *GetLogsCountUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetLogCategoryStats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsCountUsingPOSTBadRequest creates a GetLogsCountUsingPOSTBadRequest with default headers values
func NewGetLogsCountUsingPOSTBadRequest() *GetLogsCountUsingPOSTBadRequest {
	return &GetLogsCountUsingPOSTBadRequest{}
}

/*
GetLogsCountUsingPOSTBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type GetLogsCountUsingPOSTBadRequest struct {
}

// IsSuccess returns true when this get logs count using p o s t bad request response has a 2xx status code
func (o *GetLogsCountUsingPOSTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get logs count using p o s t bad request response has a 3xx status code
func (o *GetLogsCountUsingPOSTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs count using p o s t bad request response has a 4xx status code
func (o *GetLogsCountUsingPOSTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get logs count using p o s t bad request response has a 5xx status code
func (o *GetLogsCountUsingPOSTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get logs count using p o s t bad request response a status code equal to that given
func (o *GetLogsCountUsingPOSTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get logs count using p o s t bad request response
func (o *GetLogsCountUsingPOSTBadRequest) Code() int {
	return 400
}

func (o *GetLogsCountUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /log/stats/{logCategoryId}][%d] getLogsCountUsingPOSTBadRequest ", 400)
}

func (o *GetLogsCountUsingPOSTBadRequest) String() string {
	return fmt.Sprintf("[POST /log/stats/{logCategoryId}][%d] getLogsCountUsingPOSTBadRequest ", 400)
}

func (o *GetLogsCountUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}