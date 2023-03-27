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

// GetLogsStatsUsingPOST1Reader is a Reader for the GetLogsStatsUsingPOST1 structure.
type GetLogsStatsUsingPOST1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogsStatsUsingPOST1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogsStatsUsingPOST1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLogsStatsUsingPOST1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLogsStatsUsingPOST1OK creates a GetLogsStatsUsingPOST1OK with default headers values
func NewGetLogsStatsUsingPOST1OK() *GetLogsStatsUsingPOST1OK {
	return &GetLogsStatsUsingPOST1OK{}
}

/*
GetLogsStatsUsingPOST1OK describes a response with status code 200, with default header values.

OK
*/
type GetLogsStatsUsingPOST1OK struct {
	Payload *models.GetLogStats
}

// IsSuccess returns true when this get logs stats using p o s t1 o k response has a 2xx status code
func (o *GetLogsStatsUsingPOST1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get logs stats using p o s t1 o k response has a 3xx status code
func (o *GetLogsStatsUsingPOST1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs stats using p o s t1 o k response has a 4xx status code
func (o *GetLogsStatsUsingPOST1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get logs stats using p o s t1 o k response has a 5xx status code
func (o *GetLogsStatsUsingPOST1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get logs stats using p o s t1 o k response a status code equal to that given
func (o *GetLogsStatsUsingPOST1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get logs stats using p o s t1 o k response
func (o *GetLogsStatsUsingPOST1OK) Code() int {
	return 200
}

func (o *GetLogsStatsUsingPOST1OK) Error() string {
	return fmt.Sprintf("[POST /log/stats/user][%d] getLogsStatsUsingPOST1OK  %+v", 200, o.Payload)
}

func (o *GetLogsStatsUsingPOST1OK) String() string {
	return fmt.Sprintf("[POST /log/stats/user][%d] getLogsStatsUsingPOST1OK  %+v", 200, o.Payload)
}

func (o *GetLogsStatsUsingPOST1OK) GetPayload() *models.GetLogStats {
	return o.Payload
}

func (o *GetLogsStatsUsingPOST1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetLogStats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsStatsUsingPOST1BadRequest creates a GetLogsStatsUsingPOST1BadRequest with default headers values
func NewGetLogsStatsUsingPOST1BadRequest() *GetLogsStatsUsingPOST1BadRequest {
	return &GetLogsStatsUsingPOST1BadRequest{}
}

/*
GetLogsStatsUsingPOST1BadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type GetLogsStatsUsingPOST1BadRequest struct {
}

// IsSuccess returns true when this get logs stats using p o s t1 bad request response has a 2xx status code
func (o *GetLogsStatsUsingPOST1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get logs stats using p o s t1 bad request response has a 3xx status code
func (o *GetLogsStatsUsingPOST1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs stats using p o s t1 bad request response has a 4xx status code
func (o *GetLogsStatsUsingPOST1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get logs stats using p o s t1 bad request response has a 5xx status code
func (o *GetLogsStatsUsingPOST1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get logs stats using p o s t1 bad request response a status code equal to that given
func (o *GetLogsStatsUsingPOST1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get logs stats using p o s t1 bad request response
func (o *GetLogsStatsUsingPOST1BadRequest) Code() int {
	return 400
}

func (o *GetLogsStatsUsingPOST1BadRequest) Error() string {
	return fmt.Sprintf("[POST /log/stats/user][%d] getLogsStatsUsingPOST1BadRequest ", 400)
}

func (o *GetLogsStatsUsingPOST1BadRequest) String() string {
	return fmt.Sprintf("[POST /log/stats/user][%d] getLogsStatsUsingPOST1BadRequest ", 400)
}

func (o *GetLogsStatsUsingPOST1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
