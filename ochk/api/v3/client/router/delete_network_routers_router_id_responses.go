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

// DeleteNetworkRoutersRouterIDReader is a Reader for the DeleteNetworkRoutersRouterID structure.
type DeleteNetworkRoutersRouterIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkRoutersRouterIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNetworkRoutersRouterIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteNetworkRoutersRouterIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteNetworkRoutersRouterIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteNetworkRoutersRouterIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteNetworkRoutersRouterIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /network/routers/{routerId}] DeleteNetworkRoutersRouterID", response, response.Code())
	}
}

// NewDeleteNetworkRoutersRouterIDOK creates a DeleteNetworkRoutersRouterIDOK with default headers values
func NewDeleteNetworkRoutersRouterIDOK() *DeleteNetworkRoutersRouterIDOK {
	return &DeleteNetworkRoutersRouterIDOK{}
}

/*
DeleteNetworkRoutersRouterIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteNetworkRoutersRouterIDOK struct {
	Payload *models.DeleteRouterResponse
}

// IsSuccess returns true when this delete network routers router Id o k response has a 2xx status code
func (o *DeleteNetworkRoutersRouterIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete network routers router Id o k response has a 3xx status code
func (o *DeleteNetworkRoutersRouterIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete network routers router Id o k response has a 4xx status code
func (o *DeleteNetworkRoutersRouterIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete network routers router Id o k response has a 5xx status code
func (o *DeleteNetworkRoutersRouterIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete network routers router Id o k response a status code equal to that given
func (o *DeleteNetworkRoutersRouterIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete network routers router Id o k response
func (o *DeleteNetworkRoutersRouterIDOK) Code() int {
	return 200
}

func (o *DeleteNetworkRoutersRouterIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/routers/{routerId}][%d] deleteNetworkRoutersRouterIdOK %s", 200, payload)
}

func (o *DeleteNetworkRoutersRouterIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/routers/{routerId}][%d] deleteNetworkRoutersRouterIdOK %s", 200, payload)
}

func (o *DeleteNetworkRoutersRouterIDOK) GetPayload() *models.DeleteRouterResponse {
	return o.Payload
}

func (o *DeleteNetworkRoutersRouterIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteRouterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkRoutersRouterIDBadRequest creates a DeleteNetworkRoutersRouterIDBadRequest with default headers values
func NewDeleteNetworkRoutersRouterIDBadRequest() *DeleteNetworkRoutersRouterIDBadRequest {
	return &DeleteNetworkRoutersRouterIDBadRequest{}
}

/*
DeleteNetworkRoutersRouterIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteNetworkRoutersRouterIDBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this delete network routers router Id bad request response has a 2xx status code
func (o *DeleteNetworkRoutersRouterIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete network routers router Id bad request response has a 3xx status code
func (o *DeleteNetworkRoutersRouterIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete network routers router Id bad request response has a 4xx status code
func (o *DeleteNetworkRoutersRouterIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete network routers router Id bad request response has a 5xx status code
func (o *DeleteNetworkRoutersRouterIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete network routers router Id bad request response a status code equal to that given
func (o *DeleteNetworkRoutersRouterIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete network routers router Id bad request response
func (o *DeleteNetworkRoutersRouterIDBadRequest) Code() int {
	return 400
}

func (o *DeleteNetworkRoutersRouterIDBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/routers/{routerId}][%d] deleteNetworkRoutersRouterIdBadRequest %s", 400, payload)
}

func (o *DeleteNetworkRoutersRouterIDBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/routers/{routerId}][%d] deleteNetworkRoutersRouterIdBadRequest %s", 400, payload)
}

func (o *DeleteNetworkRoutersRouterIDBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *DeleteNetworkRoutersRouterIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkRoutersRouterIDUnauthorized creates a DeleteNetworkRoutersRouterIDUnauthorized with default headers values
func NewDeleteNetworkRoutersRouterIDUnauthorized() *DeleteNetworkRoutersRouterIDUnauthorized {
	return &DeleteNetworkRoutersRouterIDUnauthorized{}
}

/*
DeleteNetworkRoutersRouterIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteNetworkRoutersRouterIDUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this delete network routers router Id unauthorized response has a 2xx status code
func (o *DeleteNetworkRoutersRouterIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete network routers router Id unauthorized response has a 3xx status code
func (o *DeleteNetworkRoutersRouterIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete network routers router Id unauthorized response has a 4xx status code
func (o *DeleteNetworkRoutersRouterIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete network routers router Id unauthorized response has a 5xx status code
func (o *DeleteNetworkRoutersRouterIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete network routers router Id unauthorized response a status code equal to that given
func (o *DeleteNetworkRoutersRouterIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete network routers router Id unauthorized response
func (o *DeleteNetworkRoutersRouterIDUnauthorized) Code() int {
	return 401
}

func (o *DeleteNetworkRoutersRouterIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/routers/{routerId}][%d] deleteNetworkRoutersRouterIdUnauthorized %s", 401, payload)
}

func (o *DeleteNetworkRoutersRouterIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/routers/{routerId}][%d] deleteNetworkRoutersRouterIdUnauthorized %s", 401, payload)
}

func (o *DeleteNetworkRoutersRouterIDUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *DeleteNetworkRoutersRouterIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkRoutersRouterIDForbidden creates a DeleteNetworkRoutersRouterIDForbidden with default headers values
func NewDeleteNetworkRoutersRouterIDForbidden() *DeleteNetworkRoutersRouterIDForbidden {
	return &DeleteNetworkRoutersRouterIDForbidden{}
}

/*
DeleteNetworkRoutersRouterIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteNetworkRoutersRouterIDForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this delete network routers router Id forbidden response has a 2xx status code
func (o *DeleteNetworkRoutersRouterIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete network routers router Id forbidden response has a 3xx status code
func (o *DeleteNetworkRoutersRouterIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete network routers router Id forbidden response has a 4xx status code
func (o *DeleteNetworkRoutersRouterIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete network routers router Id forbidden response has a 5xx status code
func (o *DeleteNetworkRoutersRouterIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete network routers router Id forbidden response a status code equal to that given
func (o *DeleteNetworkRoutersRouterIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete network routers router Id forbidden response
func (o *DeleteNetworkRoutersRouterIDForbidden) Code() int {
	return 403
}

func (o *DeleteNetworkRoutersRouterIDForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/routers/{routerId}][%d] deleteNetworkRoutersRouterIdForbidden %s", 403, payload)
}

func (o *DeleteNetworkRoutersRouterIDForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/routers/{routerId}][%d] deleteNetworkRoutersRouterIdForbidden %s", 403, payload)
}

func (o *DeleteNetworkRoutersRouterIDForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *DeleteNetworkRoutersRouterIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkRoutersRouterIDNotFound creates a DeleteNetworkRoutersRouterIDNotFound with default headers values
func NewDeleteNetworkRoutersRouterIDNotFound() *DeleteNetworkRoutersRouterIDNotFound {
	return &DeleteNetworkRoutersRouterIDNotFound{}
}

/*
DeleteNetworkRoutersRouterIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteNetworkRoutersRouterIDNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this delete network routers router Id not found response has a 2xx status code
func (o *DeleteNetworkRoutersRouterIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete network routers router Id not found response has a 3xx status code
func (o *DeleteNetworkRoutersRouterIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete network routers router Id not found response has a 4xx status code
func (o *DeleteNetworkRoutersRouterIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete network routers router Id not found response has a 5xx status code
func (o *DeleteNetworkRoutersRouterIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete network routers router Id not found response a status code equal to that given
func (o *DeleteNetworkRoutersRouterIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete network routers router Id not found response
func (o *DeleteNetworkRoutersRouterIDNotFound) Code() int {
	return 404
}

func (o *DeleteNetworkRoutersRouterIDNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/routers/{routerId}][%d] deleteNetworkRoutersRouterIdNotFound %s", 404, payload)
}

func (o *DeleteNetworkRoutersRouterIDNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/routers/{routerId}][%d] deleteNetworkRoutersRouterIdNotFound %s", 404, payload)
}

func (o *DeleteNetworkRoutersRouterIDNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *DeleteNetworkRoutersRouterIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
