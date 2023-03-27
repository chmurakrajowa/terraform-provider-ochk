// Code generated by go-swagger; DO NOT EDIT.

package ip_collections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// IPCollectionDeleteUsingDELETEReader is a Reader for the IPCollectionDeleteUsingDELETE structure.
type IPCollectionDeleteUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPCollectionDeleteUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIPCollectionDeleteUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIPCollectionDeleteUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIPCollectionDeleteUsingDELETEOK creates a IPCollectionDeleteUsingDELETEOK with default headers values
func NewIPCollectionDeleteUsingDELETEOK() *IPCollectionDeleteUsingDELETEOK {
	return &IPCollectionDeleteUsingDELETEOK{}
}

/*
IPCollectionDeleteUsingDELETEOK describes a response with status code 200, with default header values.

OK
*/
type IPCollectionDeleteUsingDELETEOK struct {
	Payload *models.IPCollectionDeleteResponse
}

// IsSuccess returns true when this ip collection delete using d e l e t e o k response has a 2xx status code
func (o *IPCollectionDeleteUsingDELETEOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ip collection delete using d e l e t e o k response has a 3xx status code
func (o *IPCollectionDeleteUsingDELETEOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ip collection delete using d e l e t e o k response has a 4xx status code
func (o *IPCollectionDeleteUsingDELETEOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ip collection delete using d e l e t e o k response has a 5xx status code
func (o *IPCollectionDeleteUsingDELETEOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ip collection delete using d e l e t e o k response a status code equal to that given
func (o *IPCollectionDeleteUsingDELETEOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ip collection delete using d e l e t e o k response
func (o *IPCollectionDeleteUsingDELETEOK) Code() int {
	return 200
}

func (o *IPCollectionDeleteUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /ipcs/{ipCollectionId}][%d] ipCollectionDeleteUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *IPCollectionDeleteUsingDELETEOK) String() string {
	return fmt.Sprintf("[DELETE /ipcs/{ipCollectionId}][%d] ipCollectionDeleteUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *IPCollectionDeleteUsingDELETEOK) GetPayload() *models.IPCollectionDeleteResponse {
	return o.Payload
}

func (o *IPCollectionDeleteUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPCollectionDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIPCollectionDeleteUsingDELETEBadRequest creates a IPCollectionDeleteUsingDELETEBadRequest with default headers values
func NewIPCollectionDeleteUsingDELETEBadRequest() *IPCollectionDeleteUsingDELETEBadRequest {
	return &IPCollectionDeleteUsingDELETEBadRequest{}
}

/*
IPCollectionDeleteUsingDELETEBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type IPCollectionDeleteUsingDELETEBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this ip collection delete using d e l e t e bad request response has a 2xx status code
func (o *IPCollectionDeleteUsingDELETEBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ip collection delete using d e l e t e bad request response has a 3xx status code
func (o *IPCollectionDeleteUsingDELETEBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ip collection delete using d e l e t e bad request response has a 4xx status code
func (o *IPCollectionDeleteUsingDELETEBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ip collection delete using d e l e t e bad request response has a 5xx status code
func (o *IPCollectionDeleteUsingDELETEBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ip collection delete using d e l e t e bad request response a status code equal to that given
func (o *IPCollectionDeleteUsingDELETEBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the ip collection delete using d e l e t e bad request response
func (o *IPCollectionDeleteUsingDELETEBadRequest) Code() int {
	return 400
}

func (o *IPCollectionDeleteUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /ipcs/{ipCollectionId}][%d] ipCollectionDeleteUsingDELETEBadRequest  %+v", 400, o.Payload)
}

func (o *IPCollectionDeleteUsingDELETEBadRequest) String() string {
	return fmt.Sprintf("[DELETE /ipcs/{ipCollectionId}][%d] ipCollectionDeleteUsingDELETEBadRequest  %+v", 400, o.Payload)
}

func (o *IPCollectionDeleteUsingDELETEBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *IPCollectionDeleteUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
