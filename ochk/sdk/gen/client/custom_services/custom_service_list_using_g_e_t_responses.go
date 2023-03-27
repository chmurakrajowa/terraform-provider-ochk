// Code generated by go-swagger; DO NOT EDIT.

package custom_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// CustomServiceListUsingGETReader is a Reader for the CustomServiceListUsingGET structure.
type CustomServiceListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomServiceListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomServiceListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCustomServiceListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCustomServiceListUsingGETOK creates a CustomServiceListUsingGETOK with default headers values
func NewCustomServiceListUsingGETOK() *CustomServiceListUsingGETOK {
	return &CustomServiceListUsingGETOK{}
}

/*
CustomServiceListUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type CustomServiceListUsingGETOK struct {
	Payload *models.CustomServiceListResponse
}

// IsSuccess returns true when this custom service list using g e t o k response has a 2xx status code
func (o *CustomServiceListUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this custom service list using g e t o k response has a 3xx status code
func (o *CustomServiceListUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this custom service list using g e t o k response has a 4xx status code
func (o *CustomServiceListUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this custom service list using g e t o k response has a 5xx status code
func (o *CustomServiceListUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this custom service list using g e t o k response a status code equal to that given
func (o *CustomServiceListUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the custom service list using g e t o k response
func (o *CustomServiceListUsingGETOK) Code() int {
	return 200
}

func (o *CustomServiceListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /network/custom-services][%d] customServiceListUsingGETOK  %+v", 200, o.Payload)
}

func (o *CustomServiceListUsingGETOK) String() string {
	return fmt.Sprintf("[GET /network/custom-services][%d] customServiceListUsingGETOK  %+v", 200, o.Payload)
}

func (o *CustomServiceListUsingGETOK) GetPayload() *models.CustomServiceListResponse {
	return o.Payload
}

func (o *CustomServiceListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomServiceListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomServiceListUsingGETBadRequest creates a CustomServiceListUsingGETBadRequest with default headers values
func NewCustomServiceListUsingGETBadRequest() *CustomServiceListUsingGETBadRequest {
	return &CustomServiceListUsingGETBadRequest{}
}

/*
CustomServiceListUsingGETBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type CustomServiceListUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this custom service list using g e t bad request response has a 2xx status code
func (o *CustomServiceListUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this custom service list using g e t bad request response has a 3xx status code
func (o *CustomServiceListUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this custom service list using g e t bad request response has a 4xx status code
func (o *CustomServiceListUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this custom service list using g e t bad request response has a 5xx status code
func (o *CustomServiceListUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this custom service list using g e t bad request response a status code equal to that given
func (o *CustomServiceListUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the custom service list using g e t bad request response
func (o *CustomServiceListUsingGETBadRequest) Code() int {
	return 400
}

func (o *CustomServiceListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /network/custom-services][%d] customServiceListUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *CustomServiceListUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /network/custom-services][%d] customServiceListUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *CustomServiceListUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *CustomServiceListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
