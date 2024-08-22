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

// PostNetworkRoutersReader is a Reader for the PostNetworkRouters structure.
type PostNetworkRoutersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNetworkRoutersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostNetworkRoutersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostNetworkRoutersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostNetworkRoutersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostNetworkRoutersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /network/routers] PostNetworkRouters", response, response.Code())
	}
}

// NewPostNetworkRoutersOK creates a PostNetworkRoutersOK with default headers values
func NewPostNetworkRoutersOK() *PostNetworkRoutersOK {
	return &PostNetworkRoutersOK{}
}

/*
PostNetworkRoutersOK describes a response with status code 200, with default header values.

OK
*/
type PostNetworkRoutersOK struct {
	Payload *models.CreateRouterResponse
}

// IsSuccess returns true when this post network routers o k response has a 2xx status code
func (o *PostNetworkRoutersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post network routers o k response has a 3xx status code
func (o *PostNetworkRoutersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post network routers o k response has a 4xx status code
func (o *PostNetworkRoutersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post network routers o k response has a 5xx status code
func (o *PostNetworkRoutersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post network routers o k response a status code equal to that given
func (o *PostNetworkRoutersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post network routers o k response
func (o *PostNetworkRoutersOK) Code() int {
	return 200
}

func (o *PostNetworkRoutersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/routers][%d] postNetworkRoutersOK %s", 200, payload)
}

func (o *PostNetworkRoutersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/routers][%d] postNetworkRoutersOK %s", 200, payload)
}

func (o *PostNetworkRoutersOK) GetPayload() *models.CreateRouterResponse {
	return o.Payload
}

func (o *PostNetworkRoutersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateRouterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNetworkRoutersBadRequest creates a PostNetworkRoutersBadRequest with default headers values
func NewPostNetworkRoutersBadRequest() *PostNetworkRoutersBadRequest {
	return &PostNetworkRoutersBadRequest{}
}

/*
PostNetworkRoutersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostNetworkRoutersBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this post network routers bad request response has a 2xx status code
func (o *PostNetworkRoutersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post network routers bad request response has a 3xx status code
func (o *PostNetworkRoutersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post network routers bad request response has a 4xx status code
func (o *PostNetworkRoutersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post network routers bad request response has a 5xx status code
func (o *PostNetworkRoutersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post network routers bad request response a status code equal to that given
func (o *PostNetworkRoutersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post network routers bad request response
func (o *PostNetworkRoutersBadRequest) Code() int {
	return 400
}

func (o *PostNetworkRoutersBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/routers][%d] postNetworkRoutersBadRequest %s", 400, payload)
}

func (o *PostNetworkRoutersBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/routers][%d] postNetworkRoutersBadRequest %s", 400, payload)
}

func (o *PostNetworkRoutersBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PostNetworkRoutersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNetworkRoutersUnauthorized creates a PostNetworkRoutersUnauthorized with default headers values
func NewPostNetworkRoutersUnauthorized() *PostNetworkRoutersUnauthorized {
	return &PostNetworkRoutersUnauthorized{}
}

/*
PostNetworkRoutersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostNetworkRoutersUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this post network routers unauthorized response has a 2xx status code
func (o *PostNetworkRoutersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post network routers unauthorized response has a 3xx status code
func (o *PostNetworkRoutersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post network routers unauthorized response has a 4xx status code
func (o *PostNetworkRoutersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post network routers unauthorized response has a 5xx status code
func (o *PostNetworkRoutersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post network routers unauthorized response a status code equal to that given
func (o *PostNetworkRoutersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post network routers unauthorized response
func (o *PostNetworkRoutersUnauthorized) Code() int {
	return 401
}

func (o *PostNetworkRoutersUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/routers][%d] postNetworkRoutersUnauthorized %s", 401, payload)
}

func (o *PostNetworkRoutersUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/routers][%d] postNetworkRoutersUnauthorized %s", 401, payload)
}

func (o *PostNetworkRoutersUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PostNetworkRoutersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNetworkRoutersForbidden creates a PostNetworkRoutersForbidden with default headers values
func NewPostNetworkRoutersForbidden() *PostNetworkRoutersForbidden {
	return &PostNetworkRoutersForbidden{}
}

/*
PostNetworkRoutersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostNetworkRoutersForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this post network routers forbidden response has a 2xx status code
func (o *PostNetworkRoutersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post network routers forbidden response has a 3xx status code
func (o *PostNetworkRoutersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post network routers forbidden response has a 4xx status code
func (o *PostNetworkRoutersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post network routers forbidden response has a 5xx status code
func (o *PostNetworkRoutersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post network routers forbidden response a status code equal to that given
func (o *PostNetworkRoutersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post network routers forbidden response
func (o *PostNetworkRoutersForbidden) Code() int {
	return 403
}

func (o *PostNetworkRoutersForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/routers][%d] postNetworkRoutersForbidden %s", 403, payload)
}

func (o *PostNetworkRoutersForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/routers][%d] postNetworkRoutersForbidden %s", 403, payload)
}

func (o *PostNetworkRoutersForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PostNetworkRoutersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
