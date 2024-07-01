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

// GetHistogramCategoryUsingPOST1Reader is a Reader for the GetHistogramCategoryUsingPOST1 structure.
type GetHistogramCategoryUsingPOST1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHistogramCategoryUsingPOST1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHistogramCategoryUsingPOST1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetHistogramCategoryUsingPOST1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /log/histogram/user/category] getHistogramCategoryUsingPOST_1", response, response.Code())
	}
}

// NewGetHistogramCategoryUsingPOST1OK creates a GetHistogramCategoryUsingPOST1OK with default headers values
func NewGetHistogramCategoryUsingPOST1OK() *GetHistogramCategoryUsingPOST1OK {
	return &GetHistogramCategoryUsingPOST1OK{}
}

/*
GetHistogramCategoryUsingPOST1OK describes a response with status code 200, with default header values.

OK
*/
type GetHistogramCategoryUsingPOST1OK struct {
	Payload *models.GetLogHistogram
}

// IsSuccess returns true when this get histogram category using p o s t1 o k response has a 2xx status code
func (o *GetHistogramCategoryUsingPOST1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get histogram category using p o s t1 o k response has a 3xx status code
func (o *GetHistogramCategoryUsingPOST1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get histogram category using p o s t1 o k response has a 4xx status code
func (o *GetHistogramCategoryUsingPOST1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get histogram category using p o s t1 o k response has a 5xx status code
func (o *GetHistogramCategoryUsingPOST1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get histogram category using p o s t1 o k response a status code equal to that given
func (o *GetHistogramCategoryUsingPOST1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get histogram category using p o s t1 o k response
func (o *GetHistogramCategoryUsingPOST1OK) Code() int {
	return 200
}

func (o *GetHistogramCategoryUsingPOST1OK) Error() string {
	return fmt.Sprintf("[POST /log/histogram/user/category][%d] getHistogramCategoryUsingPOST1OK  %+v", 200, o.Payload)
}

func (o *GetHistogramCategoryUsingPOST1OK) String() string {
	return fmt.Sprintf("[POST /log/histogram/user/category][%d] getHistogramCategoryUsingPOST1OK  %+v", 200, o.Payload)
}

func (o *GetHistogramCategoryUsingPOST1OK) GetPayload() *models.GetLogHistogram {
	return o.Payload
}

func (o *GetHistogramCategoryUsingPOST1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetLogHistogram)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHistogramCategoryUsingPOST1BadRequest creates a GetHistogramCategoryUsingPOST1BadRequest with default headers values
func NewGetHistogramCategoryUsingPOST1BadRequest() *GetHistogramCategoryUsingPOST1BadRequest {
	return &GetHistogramCategoryUsingPOST1BadRequest{}
}

/*
GetHistogramCategoryUsingPOST1BadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type GetHistogramCategoryUsingPOST1BadRequest struct {
}

// IsSuccess returns true when this get histogram category using p o s t1 bad request response has a 2xx status code
func (o *GetHistogramCategoryUsingPOST1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get histogram category using p o s t1 bad request response has a 3xx status code
func (o *GetHistogramCategoryUsingPOST1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get histogram category using p o s t1 bad request response has a 4xx status code
func (o *GetHistogramCategoryUsingPOST1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get histogram category using p o s t1 bad request response has a 5xx status code
func (o *GetHistogramCategoryUsingPOST1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get histogram category using p o s t1 bad request response a status code equal to that given
func (o *GetHistogramCategoryUsingPOST1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get histogram category using p o s t1 bad request response
func (o *GetHistogramCategoryUsingPOST1BadRequest) Code() int {
	return 400
}

func (o *GetHistogramCategoryUsingPOST1BadRequest) Error() string {
	return fmt.Sprintf("[POST /log/histogram/user/category][%d] getHistogramCategoryUsingPOST1BadRequest ", 400)
}

func (o *GetHistogramCategoryUsingPOST1BadRequest) String() string {
	return fmt.Sprintf("[POST /log/histogram/user/category][%d] getHistogramCategoryUsingPOST1BadRequest ", 400)
}

func (o *GetHistogramCategoryUsingPOST1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
