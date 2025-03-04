// Code generated by go-swagger; DO NOT EDIT.

package router

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
)

// PutNetworkRoutersReader is a Reader for the PutNetworkRouters structure.
type PutNetworkRoutersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworkRoutersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutNetworkRoutersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutNetworkRoutersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutNetworkRoutersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutNetworkRoutersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /network/routers] PutNetworkRouters", response, response.Code())
	}
}

// NewPutNetworkRoutersOK creates a PutNetworkRoutersOK with default headers values
func NewPutNetworkRoutersOK() *PutNetworkRoutersOK {
	return &PutNetworkRoutersOK{}
}

/*
PutNetworkRoutersOK describes a response with status code 200, with default header values.

OK
*/
type PutNetworkRoutersOK struct {
	Payload *models.CreateRouterResponse
}

// IsSuccess returns true when this put network routers o k response has a 2xx status code
func (o *PutNetworkRoutersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put network routers o k response has a 3xx status code
func (o *PutNetworkRoutersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put network routers o k response has a 4xx status code
func (o *PutNetworkRoutersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put network routers o k response has a 5xx status code
func (o *PutNetworkRoutersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put network routers o k response a status code equal to that given
func (o *PutNetworkRoutersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put network routers o k response
func (o *PutNetworkRoutersOK) Code() int {
	return 200
}

func (o *PutNetworkRoutersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/routers][%d] putNetworkRoutersOK %s", 200, payload)
}

func (o *PutNetworkRoutersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/routers][%d] putNetworkRoutersOK %s", 200, payload)
}

func (o *PutNetworkRoutersOK) GetPayload() *models.CreateRouterResponse {
	return o.Payload
}

func (o *PutNetworkRoutersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateRouterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNetworkRoutersBadRequest creates a PutNetworkRoutersBadRequest with default headers values
func NewPutNetworkRoutersBadRequest() *PutNetworkRoutersBadRequest {
	return &PutNetworkRoutersBadRequest{}
}

/*
PutNetworkRoutersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutNetworkRoutersBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put network routers bad request response has a 2xx status code
func (o *PutNetworkRoutersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put network routers bad request response has a 3xx status code
func (o *PutNetworkRoutersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put network routers bad request response has a 4xx status code
func (o *PutNetworkRoutersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put network routers bad request response has a 5xx status code
func (o *PutNetworkRoutersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put network routers bad request response a status code equal to that given
func (o *PutNetworkRoutersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put network routers bad request response
func (o *PutNetworkRoutersBadRequest) Code() int {
	return 400
}

func (o *PutNetworkRoutersBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/routers][%d] putNetworkRoutersBadRequest %s", 400, payload)
}

func (o *PutNetworkRoutersBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/routers][%d] putNetworkRoutersBadRequest %s", 400, payload)
}

func (o *PutNetworkRoutersBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutNetworkRoutersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNetworkRoutersUnauthorized creates a PutNetworkRoutersUnauthorized with default headers values
func NewPutNetworkRoutersUnauthorized() *PutNetworkRoutersUnauthorized {
	return &PutNetworkRoutersUnauthorized{}
}

/*
PutNetworkRoutersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutNetworkRoutersUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put network routers unauthorized response has a 2xx status code
func (o *PutNetworkRoutersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put network routers unauthorized response has a 3xx status code
func (o *PutNetworkRoutersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put network routers unauthorized response has a 4xx status code
func (o *PutNetworkRoutersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put network routers unauthorized response has a 5xx status code
func (o *PutNetworkRoutersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put network routers unauthorized response a status code equal to that given
func (o *PutNetworkRoutersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put network routers unauthorized response
func (o *PutNetworkRoutersUnauthorized) Code() int {
	return 401
}

func (o *PutNetworkRoutersUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/routers][%d] putNetworkRoutersUnauthorized %s", 401, payload)
}

func (o *PutNetworkRoutersUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/routers][%d] putNetworkRoutersUnauthorized %s", 401, payload)
}

func (o *PutNetworkRoutersUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutNetworkRoutersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNetworkRoutersForbidden creates a PutNetworkRoutersForbidden with default headers values
func NewPutNetworkRoutersForbidden() *PutNetworkRoutersForbidden {
	return &PutNetworkRoutersForbidden{}
}

/*
PutNetworkRoutersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutNetworkRoutersForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put network routers forbidden response has a 2xx status code
func (o *PutNetworkRoutersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put network routers forbidden response has a 3xx status code
func (o *PutNetworkRoutersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put network routers forbidden response has a 4xx status code
func (o *PutNetworkRoutersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put network routers forbidden response has a 5xx status code
func (o *PutNetworkRoutersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put network routers forbidden response a status code equal to that given
func (o *PutNetworkRoutersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put network routers forbidden response
func (o *PutNetworkRoutersForbidden) Code() int {
	return 403
}

func (o *PutNetworkRoutersForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/routers][%d] putNetworkRoutersForbidden %s", 403, payload)
}

func (o *PutNetworkRoutersForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/routers][%d] putNetworkRoutersForbidden %s", 403, payload)
}

func (o *PutNetworkRoutersForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutNetworkRoutersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
