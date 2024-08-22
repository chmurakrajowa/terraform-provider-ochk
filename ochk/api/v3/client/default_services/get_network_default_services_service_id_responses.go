// Code generated by go-swagger; DO NOT EDIT.

package default_services

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

// GetNetworkDefaultServicesServiceIDReader is a Reader for the GetNetworkDefaultServicesServiceID structure.
type GetNetworkDefaultServicesServiceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkDefaultServicesServiceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkDefaultServicesServiceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNetworkDefaultServicesServiceIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetNetworkDefaultServicesServiceIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNetworkDefaultServicesServiceIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNetworkDefaultServicesServiceIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /network/default-services/{serviceId}] GetNetworkDefaultServicesServiceID", response, response.Code())
	}
}

// NewGetNetworkDefaultServicesServiceIDOK creates a GetNetworkDefaultServicesServiceIDOK with default headers values
func NewGetNetworkDefaultServicesServiceIDOK() *GetNetworkDefaultServicesServiceIDOK {
	return &GetNetworkDefaultServicesServiceIDOK{}
}

/*
GetNetworkDefaultServicesServiceIDOK describes a response with status code 200, with default header values.

OK
*/
type GetNetworkDefaultServicesServiceIDOK struct {
	Payload *models.GetServiceResponse
}

// IsSuccess returns true when this get network default services service Id o k response has a 2xx status code
func (o *GetNetworkDefaultServicesServiceIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network default services service Id o k response has a 3xx status code
func (o *GetNetworkDefaultServicesServiceIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network default services service Id o k response has a 4xx status code
func (o *GetNetworkDefaultServicesServiceIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network default services service Id o k response has a 5xx status code
func (o *GetNetworkDefaultServicesServiceIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network default services service Id o k response a status code equal to that given
func (o *GetNetworkDefaultServicesServiceIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network default services service Id o k response
func (o *GetNetworkDefaultServicesServiceIDOK) Code() int {
	return 200
}

func (o *GetNetworkDefaultServicesServiceIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/default-services/{serviceId}][%d] getNetworkDefaultServicesServiceIdOK %s", 200, payload)
}

func (o *GetNetworkDefaultServicesServiceIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/default-services/{serviceId}][%d] getNetworkDefaultServicesServiceIdOK %s", 200, payload)
}

func (o *GetNetworkDefaultServicesServiceIDOK) GetPayload() *models.GetServiceResponse {
	return o.Payload
}

func (o *GetNetworkDefaultServicesServiceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkDefaultServicesServiceIDBadRequest creates a GetNetworkDefaultServicesServiceIDBadRequest with default headers values
func NewGetNetworkDefaultServicesServiceIDBadRequest() *GetNetworkDefaultServicesServiceIDBadRequest {
	return &GetNetworkDefaultServicesServiceIDBadRequest{}
}

/*
GetNetworkDefaultServicesServiceIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetNetworkDefaultServicesServiceIDBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get network default services service Id bad request response has a 2xx status code
func (o *GetNetworkDefaultServicesServiceIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network default services service Id bad request response has a 3xx status code
func (o *GetNetworkDefaultServicesServiceIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network default services service Id bad request response has a 4xx status code
func (o *GetNetworkDefaultServicesServiceIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get network default services service Id bad request response has a 5xx status code
func (o *GetNetworkDefaultServicesServiceIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get network default services service Id bad request response a status code equal to that given
func (o *GetNetworkDefaultServicesServiceIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get network default services service Id bad request response
func (o *GetNetworkDefaultServicesServiceIDBadRequest) Code() int {
	return 400
}

func (o *GetNetworkDefaultServicesServiceIDBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/default-services/{serviceId}][%d] getNetworkDefaultServicesServiceIdBadRequest %s", 400, payload)
}

func (o *GetNetworkDefaultServicesServiceIDBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/default-services/{serviceId}][%d] getNetworkDefaultServicesServiceIdBadRequest %s", 400, payload)
}

func (o *GetNetworkDefaultServicesServiceIDBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetNetworkDefaultServicesServiceIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkDefaultServicesServiceIDUnauthorized creates a GetNetworkDefaultServicesServiceIDUnauthorized with default headers values
func NewGetNetworkDefaultServicesServiceIDUnauthorized() *GetNetworkDefaultServicesServiceIDUnauthorized {
	return &GetNetworkDefaultServicesServiceIDUnauthorized{}
}

/*
GetNetworkDefaultServicesServiceIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetNetworkDefaultServicesServiceIDUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get network default services service Id unauthorized response has a 2xx status code
func (o *GetNetworkDefaultServicesServiceIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network default services service Id unauthorized response has a 3xx status code
func (o *GetNetworkDefaultServicesServiceIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network default services service Id unauthorized response has a 4xx status code
func (o *GetNetworkDefaultServicesServiceIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get network default services service Id unauthorized response has a 5xx status code
func (o *GetNetworkDefaultServicesServiceIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get network default services service Id unauthorized response a status code equal to that given
func (o *GetNetworkDefaultServicesServiceIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get network default services service Id unauthorized response
func (o *GetNetworkDefaultServicesServiceIDUnauthorized) Code() int {
	return 401
}

func (o *GetNetworkDefaultServicesServiceIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/default-services/{serviceId}][%d] getNetworkDefaultServicesServiceIdUnauthorized %s", 401, payload)
}

func (o *GetNetworkDefaultServicesServiceIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/default-services/{serviceId}][%d] getNetworkDefaultServicesServiceIdUnauthorized %s", 401, payload)
}

func (o *GetNetworkDefaultServicesServiceIDUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetNetworkDefaultServicesServiceIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkDefaultServicesServiceIDForbidden creates a GetNetworkDefaultServicesServiceIDForbidden with default headers values
func NewGetNetworkDefaultServicesServiceIDForbidden() *GetNetworkDefaultServicesServiceIDForbidden {
	return &GetNetworkDefaultServicesServiceIDForbidden{}
}

/*
GetNetworkDefaultServicesServiceIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetNetworkDefaultServicesServiceIDForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get network default services service Id forbidden response has a 2xx status code
func (o *GetNetworkDefaultServicesServiceIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network default services service Id forbidden response has a 3xx status code
func (o *GetNetworkDefaultServicesServiceIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network default services service Id forbidden response has a 4xx status code
func (o *GetNetworkDefaultServicesServiceIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get network default services service Id forbidden response has a 5xx status code
func (o *GetNetworkDefaultServicesServiceIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get network default services service Id forbidden response a status code equal to that given
func (o *GetNetworkDefaultServicesServiceIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get network default services service Id forbidden response
func (o *GetNetworkDefaultServicesServiceIDForbidden) Code() int {
	return 403
}

func (o *GetNetworkDefaultServicesServiceIDForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/default-services/{serviceId}][%d] getNetworkDefaultServicesServiceIdForbidden %s", 403, payload)
}

func (o *GetNetworkDefaultServicesServiceIDForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/default-services/{serviceId}][%d] getNetworkDefaultServicesServiceIdForbidden %s", 403, payload)
}

func (o *GetNetworkDefaultServicesServiceIDForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetNetworkDefaultServicesServiceIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkDefaultServicesServiceIDNotFound creates a GetNetworkDefaultServicesServiceIDNotFound with default headers values
func NewGetNetworkDefaultServicesServiceIDNotFound() *GetNetworkDefaultServicesServiceIDNotFound {
	return &GetNetworkDefaultServicesServiceIDNotFound{}
}

/*
GetNetworkDefaultServicesServiceIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetNetworkDefaultServicesServiceIDNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get network default services service Id not found response has a 2xx status code
func (o *GetNetworkDefaultServicesServiceIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network default services service Id not found response has a 3xx status code
func (o *GetNetworkDefaultServicesServiceIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network default services service Id not found response has a 4xx status code
func (o *GetNetworkDefaultServicesServiceIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get network default services service Id not found response has a 5xx status code
func (o *GetNetworkDefaultServicesServiceIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get network default services service Id not found response a status code equal to that given
func (o *GetNetworkDefaultServicesServiceIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get network default services service Id not found response
func (o *GetNetworkDefaultServicesServiceIDNotFound) Code() int {
	return 404
}

func (o *GetNetworkDefaultServicesServiceIDNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/default-services/{serviceId}][%d] getNetworkDefaultServicesServiceIdNotFound %s", 404, payload)
}

func (o *GetNetworkDefaultServicesServiceIDNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/default-services/{serviceId}][%d] getNetworkDefaultServicesServiceIdNotFound %s", 404, payload)
}

func (o *GetNetworkDefaultServicesServiceIDNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetNetworkDefaultServicesServiceIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
