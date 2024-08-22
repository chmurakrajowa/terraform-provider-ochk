// Code generated by go-swagger; DO NOT EDIT.

package custom_services

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

// GetNetworkCustomServicesReader is a Reader for the GetNetworkCustomServices structure.
type GetNetworkCustomServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkCustomServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkCustomServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNetworkCustomServicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetNetworkCustomServicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNetworkCustomServicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /network/custom-services] GetNetworkCustomServices", response, response.Code())
	}
}

// NewGetNetworkCustomServicesOK creates a GetNetworkCustomServicesOK with default headers values
func NewGetNetworkCustomServicesOK() *GetNetworkCustomServicesOK {
	return &GetNetworkCustomServicesOK{}
}

/*
GetNetworkCustomServicesOK describes a response with status code 200, with default header values.

OK
*/
type GetNetworkCustomServicesOK struct {
	Payload *models.ListCustomServicesResponse
}

// IsSuccess returns true when this get network custom services o k response has a 2xx status code
func (o *GetNetworkCustomServicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network custom services o k response has a 3xx status code
func (o *GetNetworkCustomServicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network custom services o k response has a 4xx status code
func (o *GetNetworkCustomServicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network custom services o k response has a 5xx status code
func (o *GetNetworkCustomServicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network custom services o k response a status code equal to that given
func (o *GetNetworkCustomServicesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network custom services o k response
func (o *GetNetworkCustomServicesOK) Code() int {
	return 200
}

func (o *GetNetworkCustomServicesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/custom-services][%d] getNetworkCustomServicesOK %s", 200, payload)
}

func (o *GetNetworkCustomServicesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/custom-services][%d] getNetworkCustomServicesOK %s", 200, payload)
}

func (o *GetNetworkCustomServicesOK) GetPayload() *models.ListCustomServicesResponse {
	return o.Payload
}

func (o *GetNetworkCustomServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListCustomServicesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkCustomServicesBadRequest creates a GetNetworkCustomServicesBadRequest with default headers values
func NewGetNetworkCustomServicesBadRequest() *GetNetworkCustomServicesBadRequest {
	return &GetNetworkCustomServicesBadRequest{}
}

/*
GetNetworkCustomServicesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetNetworkCustomServicesBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get network custom services bad request response has a 2xx status code
func (o *GetNetworkCustomServicesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network custom services bad request response has a 3xx status code
func (o *GetNetworkCustomServicesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network custom services bad request response has a 4xx status code
func (o *GetNetworkCustomServicesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get network custom services bad request response has a 5xx status code
func (o *GetNetworkCustomServicesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get network custom services bad request response a status code equal to that given
func (o *GetNetworkCustomServicesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get network custom services bad request response
func (o *GetNetworkCustomServicesBadRequest) Code() int {
	return 400
}

func (o *GetNetworkCustomServicesBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/custom-services][%d] getNetworkCustomServicesBadRequest %s", 400, payload)
}

func (o *GetNetworkCustomServicesBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/custom-services][%d] getNetworkCustomServicesBadRequest %s", 400, payload)
}

func (o *GetNetworkCustomServicesBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetNetworkCustomServicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkCustomServicesUnauthorized creates a GetNetworkCustomServicesUnauthorized with default headers values
func NewGetNetworkCustomServicesUnauthorized() *GetNetworkCustomServicesUnauthorized {
	return &GetNetworkCustomServicesUnauthorized{}
}

/*
GetNetworkCustomServicesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetNetworkCustomServicesUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get network custom services unauthorized response has a 2xx status code
func (o *GetNetworkCustomServicesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network custom services unauthorized response has a 3xx status code
func (o *GetNetworkCustomServicesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network custom services unauthorized response has a 4xx status code
func (o *GetNetworkCustomServicesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get network custom services unauthorized response has a 5xx status code
func (o *GetNetworkCustomServicesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get network custom services unauthorized response a status code equal to that given
func (o *GetNetworkCustomServicesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get network custom services unauthorized response
func (o *GetNetworkCustomServicesUnauthorized) Code() int {
	return 401
}

func (o *GetNetworkCustomServicesUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/custom-services][%d] getNetworkCustomServicesUnauthorized %s", 401, payload)
}

func (o *GetNetworkCustomServicesUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/custom-services][%d] getNetworkCustomServicesUnauthorized %s", 401, payload)
}

func (o *GetNetworkCustomServicesUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetNetworkCustomServicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkCustomServicesForbidden creates a GetNetworkCustomServicesForbidden with default headers values
func NewGetNetworkCustomServicesForbidden() *GetNetworkCustomServicesForbidden {
	return &GetNetworkCustomServicesForbidden{}
}

/*
GetNetworkCustomServicesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetNetworkCustomServicesForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get network custom services forbidden response has a 2xx status code
func (o *GetNetworkCustomServicesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network custom services forbidden response has a 3xx status code
func (o *GetNetworkCustomServicesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network custom services forbidden response has a 4xx status code
func (o *GetNetworkCustomServicesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get network custom services forbidden response has a 5xx status code
func (o *GetNetworkCustomServicesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get network custom services forbidden response a status code equal to that given
func (o *GetNetworkCustomServicesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get network custom services forbidden response
func (o *GetNetworkCustomServicesForbidden) Code() int {
	return 403
}

func (o *GetNetworkCustomServicesForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/custom-services][%d] getNetworkCustomServicesForbidden %s", 403, payload)
}

func (o *GetNetworkCustomServicesForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/custom-services][%d] getNetworkCustomServicesForbidden %s", 403, payload)
}

func (o *GetNetworkCustomServicesForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetNetworkCustomServicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
