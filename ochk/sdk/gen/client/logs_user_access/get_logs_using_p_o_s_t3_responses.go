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

// GetLogsUsingPOST3Reader is a Reader for the GetLogsUsingPOST3 structure.
type GetLogsUsingPOST3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogsUsingPOST3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogsUsingPOST3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLogsUsingPOST3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLogsUsingPOST3OK creates a GetLogsUsingPOST3OK with default headers values
func NewGetLogsUsingPOST3OK() *GetLogsUsingPOST3OK {
	return &GetLogsUsingPOST3OK{}
}

/*
GetLogsUsingPOST3OK describes a response with status code 200, with default header values.

OK
*/
type GetLogsUsingPOST3OK struct {
	Payload *models.GetLogsResponse
}

// IsSuccess returns true when this get logs using p o s t3 o k response has a 2xx status code
func (o *GetLogsUsingPOST3OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get logs using p o s t3 o k response has a 3xx status code
func (o *GetLogsUsingPOST3OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs using p o s t3 o k response has a 4xx status code
func (o *GetLogsUsingPOST3OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get logs using p o s t3 o k response has a 5xx status code
func (o *GetLogsUsingPOST3OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get logs using p o s t3 o k response a status code equal to that given
func (o *GetLogsUsingPOST3OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get logs using p o s t3 o k response
func (o *GetLogsUsingPOST3OK) Code() int {
	return 200
}

func (o *GetLogsUsingPOST3OK) Error() string {
	return fmt.Sprintf("[POST /log/user/categories/{logCategoryId}/generate][%d] getLogsUsingPOST3OK  %+v", 200, o.Payload)
}

func (o *GetLogsUsingPOST3OK) String() string {
	return fmt.Sprintf("[POST /log/user/categories/{logCategoryId}/generate][%d] getLogsUsingPOST3OK  %+v", 200, o.Payload)
}

func (o *GetLogsUsingPOST3OK) GetPayload() *models.GetLogsResponse {
	return o.Payload
}

func (o *GetLogsUsingPOST3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetLogsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsUsingPOST3BadRequest creates a GetLogsUsingPOST3BadRequest with default headers values
func NewGetLogsUsingPOST3BadRequest() *GetLogsUsingPOST3BadRequest {
	return &GetLogsUsingPOST3BadRequest{}
}

/*
GetLogsUsingPOST3BadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type GetLogsUsingPOST3BadRequest struct {
}

// IsSuccess returns true when this get logs using p o s t3 bad request response has a 2xx status code
func (o *GetLogsUsingPOST3BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get logs using p o s t3 bad request response has a 3xx status code
func (o *GetLogsUsingPOST3BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs using p o s t3 bad request response has a 4xx status code
func (o *GetLogsUsingPOST3BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get logs using p o s t3 bad request response has a 5xx status code
func (o *GetLogsUsingPOST3BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get logs using p o s t3 bad request response a status code equal to that given
func (o *GetLogsUsingPOST3BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get logs using p o s t3 bad request response
func (o *GetLogsUsingPOST3BadRequest) Code() int {
	return 400
}

func (o *GetLogsUsingPOST3BadRequest) Error() string {
	return fmt.Sprintf("[POST /log/user/categories/{logCategoryId}/generate][%d] getLogsUsingPOST3BadRequest ", 400)
}

func (o *GetLogsUsingPOST3BadRequest) String() string {
	return fmt.Sprintf("[POST /log/user/categories/{logCategoryId}/generate][%d] getLogsUsingPOST3BadRequest ", 400)
}

func (o *GetLogsUsingPOST3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
