// Code generated by go-swagger; DO NOT EDIT.

package logs_user_access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// GetLogsUsingPOST2Reader is a Reader for the GetLogsUsingPOST2 structure.
type GetLogsUsingPOST2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogsUsingPOST2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogsUsingPOST2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLogsUsingPOST2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /log/user/categories/{logCategoryId}/generatebydsl] getLogsUsingPOST_2", response, response.Code())
	}
}

// NewGetLogsUsingPOST2OK creates a GetLogsUsingPOST2OK with default headers values
func NewGetLogsUsingPOST2OK() *GetLogsUsingPOST2OK {
	return &GetLogsUsingPOST2OK{}
}

/*
GetLogsUsingPOST2OK describes a response with status code 200, with default header values.

OK
*/
type GetLogsUsingPOST2OK struct {
	Payload *models.GetLogsResponse
}

// IsSuccess returns true when this get logs using p o s t2 o k response has a 2xx status code
func (o *GetLogsUsingPOST2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get logs using p o s t2 o k response has a 3xx status code
func (o *GetLogsUsingPOST2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs using p o s t2 o k response has a 4xx status code
func (o *GetLogsUsingPOST2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get logs using p o s t2 o k response has a 5xx status code
func (o *GetLogsUsingPOST2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get logs using p o s t2 o k response a status code equal to that given
func (o *GetLogsUsingPOST2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get logs using p o s t2 o k response
func (o *GetLogsUsingPOST2OK) Code() int {
	return 200
}

func (o *GetLogsUsingPOST2OK) Error() string {
	return fmt.Sprintf("[POST /log/user/categories/{logCategoryId}/generatebydsl][%d] getLogsUsingPOST2OK  %+v", 200, o.Payload)
}

func (o *GetLogsUsingPOST2OK) String() string {
	return fmt.Sprintf("[POST /log/user/categories/{logCategoryId}/generatebydsl][%d] getLogsUsingPOST2OK  %+v", 200, o.Payload)
}

func (o *GetLogsUsingPOST2OK) GetPayload() *models.GetLogsResponse {
	return o.Payload
}

func (o *GetLogsUsingPOST2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetLogsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsUsingPOST2BadRequest creates a GetLogsUsingPOST2BadRequest with default headers values
func NewGetLogsUsingPOST2BadRequest() *GetLogsUsingPOST2BadRequest {
	return &GetLogsUsingPOST2BadRequest{}
}

/*
GetLogsUsingPOST2BadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type GetLogsUsingPOST2BadRequest struct {
}

// IsSuccess returns true when this get logs using p o s t2 bad request response has a 2xx status code
func (o *GetLogsUsingPOST2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get logs using p o s t2 bad request response has a 3xx status code
func (o *GetLogsUsingPOST2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs using p o s t2 bad request response has a 4xx status code
func (o *GetLogsUsingPOST2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get logs using p o s t2 bad request response has a 5xx status code
func (o *GetLogsUsingPOST2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get logs using p o s t2 bad request response a status code equal to that given
func (o *GetLogsUsingPOST2BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get logs using p o s t2 bad request response
func (o *GetLogsUsingPOST2BadRequest) Code() int {
	return 400
}

func (o *GetLogsUsingPOST2BadRequest) Error() string {
	return fmt.Sprintf("[POST /log/user/categories/{logCategoryId}/generatebydsl][%d] getLogsUsingPOST2BadRequest ", 400)
}

func (o *GetLogsUsingPOST2BadRequest) String() string {
	return fmt.Sprintf("[POST /log/user/categories/{logCategoryId}/generatebydsl][%d] getLogsUsingPOST2BadRequest ", 400)
}

func (o *GetLogsUsingPOST2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
