// Code generated by go-swagger; DO NOT EDIT.

package routers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// RouterCreateUsingPUTReader is a Reader for the RouterCreateUsingPUT structure.
type RouterCreateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouterCreateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRouterCreateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewRouterCreateUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRouterCreateUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /network/routers] routerCreateUsingPUT", response, response.Code())
	}
}

// NewRouterCreateUsingPUTOK creates a RouterCreateUsingPUTOK with default headers values
func NewRouterCreateUsingPUTOK() *RouterCreateUsingPUTOK {
	return &RouterCreateUsingPUTOK{}
}

/*
RouterCreateUsingPUTOK describes a response with status code 200, with default header values.

OK
*/
type RouterCreateUsingPUTOK struct {
	Payload *models.CreateRouterResponse
}

// IsSuccess returns true when this router create using p u t o k response has a 2xx status code
func (o *RouterCreateUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this router create using p u t o k response has a 3xx status code
func (o *RouterCreateUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this router create using p u t o k response has a 4xx status code
func (o *RouterCreateUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this router create using p u t o k response has a 5xx status code
func (o *RouterCreateUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this router create using p u t o k response a status code equal to that given
func (o *RouterCreateUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the router create using p u t o k response
func (o *RouterCreateUsingPUTOK) Code() int {
	return 200
}

func (o *RouterCreateUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /network/routers][%d] routerCreateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *RouterCreateUsingPUTOK) String() string {
	return fmt.Sprintf("[PUT /network/routers][%d] routerCreateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *RouterCreateUsingPUTOK) GetPayload() *models.CreateRouterResponse {
	return o.Payload
}

func (o *RouterCreateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateRouterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRouterCreateUsingPUTCreated creates a RouterCreateUsingPUTCreated with default headers values
func NewRouterCreateUsingPUTCreated() *RouterCreateUsingPUTCreated {
	return &RouterCreateUsingPUTCreated{}
}

/*
RouterCreateUsingPUTCreated describes a response with status code 201, with default header values.

Entity has been created
*/
type RouterCreateUsingPUTCreated struct {
	Payload *models.CreateRouterResponse
}

// IsSuccess returns true when this router create using p u t created response has a 2xx status code
func (o *RouterCreateUsingPUTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this router create using p u t created response has a 3xx status code
func (o *RouterCreateUsingPUTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this router create using p u t created response has a 4xx status code
func (o *RouterCreateUsingPUTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this router create using p u t created response has a 5xx status code
func (o *RouterCreateUsingPUTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this router create using p u t created response a status code equal to that given
func (o *RouterCreateUsingPUTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the router create using p u t created response
func (o *RouterCreateUsingPUTCreated) Code() int {
	return 201
}

func (o *RouterCreateUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /network/routers][%d] routerCreateUsingPUTCreated  %+v", 201, o.Payload)
}

func (o *RouterCreateUsingPUTCreated) String() string {
	return fmt.Sprintf("[PUT /network/routers][%d] routerCreateUsingPUTCreated  %+v", 201, o.Payload)
}

func (o *RouterCreateUsingPUTCreated) GetPayload() *models.CreateRouterResponse {
	return o.Payload
}

func (o *RouterCreateUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateRouterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRouterCreateUsingPUTBadRequest creates a RouterCreateUsingPUTBadRequest with default headers values
func NewRouterCreateUsingPUTBadRequest() *RouterCreateUsingPUTBadRequest {
	return &RouterCreateUsingPUTBadRequest{}
}

/*
RouterCreateUsingPUTBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type RouterCreateUsingPUTBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this router create using p u t bad request response has a 2xx status code
func (o *RouterCreateUsingPUTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this router create using p u t bad request response has a 3xx status code
func (o *RouterCreateUsingPUTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this router create using p u t bad request response has a 4xx status code
func (o *RouterCreateUsingPUTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this router create using p u t bad request response has a 5xx status code
func (o *RouterCreateUsingPUTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this router create using p u t bad request response a status code equal to that given
func (o *RouterCreateUsingPUTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the router create using p u t bad request response
func (o *RouterCreateUsingPUTBadRequest) Code() int {
	return 400
}

func (o *RouterCreateUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /network/routers][%d] routerCreateUsingPUTBadRequest  %+v", 400, o.Payload)
}

func (o *RouterCreateUsingPUTBadRequest) String() string {
	return fmt.Sprintf("[PUT /network/routers][%d] routerCreateUsingPUTBadRequest  %+v", 400, o.Payload)
}

func (o *RouterCreateUsingPUTBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *RouterCreateUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
