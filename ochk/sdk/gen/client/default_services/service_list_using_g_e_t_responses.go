// Code generated by go-swagger; DO NOT EDIT.

package default_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// ServiceListUsingGETReader is a Reader for the ServiceListUsingGET structure.
type ServiceListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /network/default-services] serviceListUsingGET", response, response.Code())
	}
}

// NewServiceListUsingGETOK creates a ServiceListUsingGETOK with default headers values
func NewServiceListUsingGETOK() *ServiceListUsingGETOK {
	return &ServiceListUsingGETOK{}
}

/*
ServiceListUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type ServiceListUsingGETOK struct {
	Payload *models.ServiceListResponse
}

// IsSuccess returns true when this service list using g e t o k response has a 2xx status code
func (o *ServiceListUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service list using g e t o k response has a 3xx status code
func (o *ServiceListUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service list using g e t o k response has a 4xx status code
func (o *ServiceListUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service list using g e t o k response has a 5xx status code
func (o *ServiceListUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service list using g e t o k response a status code equal to that given
func (o *ServiceListUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service list using g e t o k response
func (o *ServiceListUsingGETOK) Code() int {
	return 200
}

func (o *ServiceListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /network/default-services][%d] serviceListUsingGETOK  %+v", 200, o.Payload)
}

func (o *ServiceListUsingGETOK) String() string {
	return fmt.Sprintf("[GET /network/default-services][%d] serviceListUsingGETOK  %+v", 200, o.Payload)
}

func (o *ServiceListUsingGETOK) GetPayload() *models.ServiceListResponse {
	return o.Payload
}

func (o *ServiceListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceListUsingGETBadRequest creates a ServiceListUsingGETBadRequest with default headers values
func NewServiceListUsingGETBadRequest() *ServiceListUsingGETBadRequest {
	return &ServiceListUsingGETBadRequest{}
}

/*
ServiceListUsingGETBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type ServiceListUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this service list using g e t bad request response has a 2xx status code
func (o *ServiceListUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service list using g e t bad request response has a 3xx status code
func (o *ServiceListUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service list using g e t bad request response has a 4xx status code
func (o *ServiceListUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service list using g e t bad request response has a 5xx status code
func (o *ServiceListUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service list using g e t bad request response a status code equal to that given
func (o *ServiceListUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the service list using g e t bad request response
func (o *ServiceListUsingGETBadRequest) Code() int {
	return 400
}

func (o *ServiceListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /network/default-services][%d] serviceListUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceListUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /network/default-services][%d] serviceListUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceListUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *ServiceListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
