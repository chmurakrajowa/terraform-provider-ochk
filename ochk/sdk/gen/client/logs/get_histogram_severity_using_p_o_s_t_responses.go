// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// GetHistogramSeverityUsingPOSTReader is a Reader for the GetHistogramSeverityUsingPOST structure.
type GetHistogramSeverityUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHistogramSeverityUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHistogramSeverityUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetHistogramSeverityUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /log/histogram/severity] getHistogramSeverityUsingPOST", response, response.Code())
	}
}

// NewGetHistogramSeverityUsingPOSTOK creates a GetHistogramSeverityUsingPOSTOK with default headers values
func NewGetHistogramSeverityUsingPOSTOK() *GetHistogramSeverityUsingPOSTOK {
	return &GetHistogramSeverityUsingPOSTOK{}
}

/*
GetHistogramSeverityUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type GetHistogramSeverityUsingPOSTOK struct {
	Payload *models.GetLogHistogram
}

// IsSuccess returns true when this get histogram severity using p o s t o k response has a 2xx status code
func (o *GetHistogramSeverityUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get histogram severity using p o s t o k response has a 3xx status code
func (o *GetHistogramSeverityUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get histogram severity using p o s t o k response has a 4xx status code
func (o *GetHistogramSeverityUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get histogram severity using p o s t o k response has a 5xx status code
func (o *GetHistogramSeverityUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get histogram severity using p o s t o k response a status code equal to that given
func (o *GetHistogramSeverityUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get histogram severity using p o s t o k response
func (o *GetHistogramSeverityUsingPOSTOK) Code() int {
	return 200
}

func (o *GetHistogramSeverityUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /log/histogram/severity][%d] getHistogramSeverityUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *GetHistogramSeverityUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /log/histogram/severity][%d] getHistogramSeverityUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *GetHistogramSeverityUsingPOSTOK) GetPayload() *models.GetLogHistogram {
	return o.Payload
}

func (o *GetHistogramSeverityUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetLogHistogram)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHistogramSeverityUsingPOSTBadRequest creates a GetHistogramSeverityUsingPOSTBadRequest with default headers values
func NewGetHistogramSeverityUsingPOSTBadRequest() *GetHistogramSeverityUsingPOSTBadRequest {
	return &GetHistogramSeverityUsingPOSTBadRequest{}
}

/*
GetHistogramSeverityUsingPOSTBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type GetHistogramSeverityUsingPOSTBadRequest struct {
}

// IsSuccess returns true when this get histogram severity using p o s t bad request response has a 2xx status code
func (o *GetHistogramSeverityUsingPOSTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get histogram severity using p o s t bad request response has a 3xx status code
func (o *GetHistogramSeverityUsingPOSTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get histogram severity using p o s t bad request response has a 4xx status code
func (o *GetHistogramSeverityUsingPOSTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get histogram severity using p o s t bad request response has a 5xx status code
func (o *GetHistogramSeverityUsingPOSTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get histogram severity using p o s t bad request response a status code equal to that given
func (o *GetHistogramSeverityUsingPOSTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get histogram severity using p o s t bad request response
func (o *GetHistogramSeverityUsingPOSTBadRequest) Code() int {
	return 400
}

func (o *GetHistogramSeverityUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /log/histogram/severity][%d] getHistogramSeverityUsingPOSTBadRequest ", 400)
}

func (o *GetHistogramSeverityUsingPOSTBadRequest) String() string {
	return fmt.Sprintf("[POST /log/histogram/severity][%d] getHistogramSeverityUsingPOSTBadRequest ", 400)
}

func (o *GetHistogramSeverityUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
