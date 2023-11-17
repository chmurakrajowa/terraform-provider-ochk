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

// GetLogsUsingPOST1Reader is a Reader for the GetLogsUsingPOST1 structure.
type GetLogsUsingPOST1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogsUsingPOST1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogsUsingPOST1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLogsUsingPOST1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /log/categories/{logCategoryId}/generate] getLogsUsingPOST_1", response, response.Code())
	}
}

// NewGetLogsUsingPOST1OK creates a GetLogsUsingPOST1OK with default headers values
func NewGetLogsUsingPOST1OK() *GetLogsUsingPOST1OK {
	return &GetLogsUsingPOST1OK{}
}

/*
GetLogsUsingPOST1OK describes a response with status code 200, with default header values.

Request has been completed successfully
*/
type GetLogsUsingPOST1OK struct {
	Payload *models.GetLogsResponse
}

// IsSuccess returns true when this get logs using p o s t1 o k response has a 2xx status code
func (o *GetLogsUsingPOST1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get logs using p o s t1 o k response has a 3xx status code
func (o *GetLogsUsingPOST1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs using p o s t1 o k response has a 4xx status code
func (o *GetLogsUsingPOST1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get logs using p o s t1 o k response has a 5xx status code
func (o *GetLogsUsingPOST1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get logs using p o s t1 o k response a status code equal to that given
func (o *GetLogsUsingPOST1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get logs using p o s t1 o k response
func (o *GetLogsUsingPOST1OK) Code() int {
	return 200
}

func (o *GetLogsUsingPOST1OK) Error() string {
	return fmt.Sprintf("[POST /log/categories/{logCategoryId}/generate][%d] getLogsUsingPOST1OK  %+v", 200, o.Payload)
}

func (o *GetLogsUsingPOST1OK) String() string {
	return fmt.Sprintf("[POST /log/categories/{logCategoryId}/generate][%d] getLogsUsingPOST1OK  %+v", 200, o.Payload)
}

func (o *GetLogsUsingPOST1OK) GetPayload() *models.GetLogsResponse {
	return o.Payload
}

func (o *GetLogsUsingPOST1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetLogsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsUsingPOST1BadRequest creates a GetLogsUsingPOST1BadRequest with default headers values
func NewGetLogsUsingPOST1BadRequest() *GetLogsUsingPOST1BadRequest {
	return &GetLogsUsingPOST1BadRequest{}
}

/*
GetLogsUsingPOST1BadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type GetLogsUsingPOST1BadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this get logs using p o s t1 bad request response has a 2xx status code
func (o *GetLogsUsingPOST1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get logs using p o s t1 bad request response has a 3xx status code
func (o *GetLogsUsingPOST1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs using p o s t1 bad request response has a 4xx status code
func (o *GetLogsUsingPOST1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get logs using p o s t1 bad request response has a 5xx status code
func (o *GetLogsUsingPOST1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get logs using p o s t1 bad request response a status code equal to that given
func (o *GetLogsUsingPOST1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get logs using p o s t1 bad request response
func (o *GetLogsUsingPOST1BadRequest) Code() int {
	return 400
}

func (o *GetLogsUsingPOST1BadRequest) Error() string {
	return fmt.Sprintf("[POST /log/categories/{logCategoryId}/generate][%d] getLogsUsingPOST1BadRequest  %+v", 400, o.Payload)
}

func (o *GetLogsUsingPOST1BadRequest) String() string {
	return fmt.Sprintf("[POST /log/categories/{logCategoryId}/generate][%d] getLogsUsingPOST1BadRequest  %+v", 400, o.Payload)
}

func (o *GetLogsUsingPOST1BadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *GetLogsUsingPOST1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
