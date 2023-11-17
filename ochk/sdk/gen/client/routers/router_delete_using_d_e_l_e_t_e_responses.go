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

// RouterDeleteUsingDELETEReader is a Reader for the RouterDeleteUsingDELETE structure.
type RouterDeleteUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouterDeleteUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRouterDeleteUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewRouterDeleteUsingDELETECreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRouterDeleteUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /network/routers/{routerId}] routerDeleteUsingDELETE", response, response.Code())
	}
}

// NewRouterDeleteUsingDELETEOK creates a RouterDeleteUsingDELETEOK with default headers values
func NewRouterDeleteUsingDELETEOK() *RouterDeleteUsingDELETEOK {
	return &RouterDeleteUsingDELETEOK{}
}

/*
RouterDeleteUsingDELETEOK describes a response with status code 200, with default header values.

OK
*/
type RouterDeleteUsingDELETEOK struct {
	Payload *models.DeleteRouterResponse
}

// IsSuccess returns true when this router delete using d e l e t e o k response has a 2xx status code
func (o *RouterDeleteUsingDELETEOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this router delete using d e l e t e o k response has a 3xx status code
func (o *RouterDeleteUsingDELETEOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this router delete using d e l e t e o k response has a 4xx status code
func (o *RouterDeleteUsingDELETEOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this router delete using d e l e t e o k response has a 5xx status code
func (o *RouterDeleteUsingDELETEOK) IsServerError() bool {
	return false
}

// IsCode returns true when this router delete using d e l e t e o k response a status code equal to that given
func (o *RouterDeleteUsingDELETEOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the router delete using d e l e t e o k response
func (o *RouterDeleteUsingDELETEOK) Code() int {
	return 200
}

func (o *RouterDeleteUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /network/routers/{routerId}][%d] routerDeleteUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *RouterDeleteUsingDELETEOK) String() string {
	return fmt.Sprintf("[DELETE /network/routers/{routerId}][%d] routerDeleteUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *RouterDeleteUsingDELETEOK) GetPayload() *models.DeleteRouterResponse {
	return o.Payload
}

func (o *RouterDeleteUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteRouterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRouterDeleteUsingDELETECreated creates a RouterDeleteUsingDELETECreated with default headers values
func NewRouterDeleteUsingDELETECreated() *RouterDeleteUsingDELETECreated {
	return &RouterDeleteUsingDELETECreated{}
}

/*
RouterDeleteUsingDELETECreated describes a response with status code 201, with default header values.

Entity has been deleted
*/
type RouterDeleteUsingDELETECreated struct {
	Payload *models.DeleteRouterResponse
}

// IsSuccess returns true when this router delete using d e l e t e created response has a 2xx status code
func (o *RouterDeleteUsingDELETECreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this router delete using d e l e t e created response has a 3xx status code
func (o *RouterDeleteUsingDELETECreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this router delete using d e l e t e created response has a 4xx status code
func (o *RouterDeleteUsingDELETECreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this router delete using d e l e t e created response has a 5xx status code
func (o *RouterDeleteUsingDELETECreated) IsServerError() bool {
	return false
}

// IsCode returns true when this router delete using d e l e t e created response a status code equal to that given
func (o *RouterDeleteUsingDELETECreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the router delete using d e l e t e created response
func (o *RouterDeleteUsingDELETECreated) Code() int {
	return 201
}

func (o *RouterDeleteUsingDELETECreated) Error() string {
	return fmt.Sprintf("[DELETE /network/routers/{routerId}][%d] routerDeleteUsingDELETECreated  %+v", 201, o.Payload)
}

func (o *RouterDeleteUsingDELETECreated) String() string {
	return fmt.Sprintf("[DELETE /network/routers/{routerId}][%d] routerDeleteUsingDELETECreated  %+v", 201, o.Payload)
}

func (o *RouterDeleteUsingDELETECreated) GetPayload() *models.DeleteRouterResponse {
	return o.Payload
}

func (o *RouterDeleteUsingDELETECreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteRouterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRouterDeleteUsingDELETEBadRequest creates a RouterDeleteUsingDELETEBadRequest with default headers values
func NewRouterDeleteUsingDELETEBadRequest() *RouterDeleteUsingDELETEBadRequest {
	return &RouterDeleteUsingDELETEBadRequest{}
}

/*
RouterDeleteUsingDELETEBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type RouterDeleteUsingDELETEBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this router delete using d e l e t e bad request response has a 2xx status code
func (o *RouterDeleteUsingDELETEBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this router delete using d e l e t e bad request response has a 3xx status code
func (o *RouterDeleteUsingDELETEBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this router delete using d e l e t e bad request response has a 4xx status code
func (o *RouterDeleteUsingDELETEBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this router delete using d e l e t e bad request response has a 5xx status code
func (o *RouterDeleteUsingDELETEBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this router delete using d e l e t e bad request response a status code equal to that given
func (o *RouterDeleteUsingDELETEBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the router delete using d e l e t e bad request response
func (o *RouterDeleteUsingDELETEBadRequest) Code() int {
	return 400
}

func (o *RouterDeleteUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /network/routers/{routerId}][%d] routerDeleteUsingDELETEBadRequest  %+v", 400, o.Payload)
}

func (o *RouterDeleteUsingDELETEBadRequest) String() string {
	return fmt.Sprintf("[DELETE /network/routers/{routerId}][%d] routerDeleteUsingDELETEBadRequest  %+v", 400, o.Payload)
}

func (o *RouterDeleteUsingDELETEBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *RouterDeleteUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
