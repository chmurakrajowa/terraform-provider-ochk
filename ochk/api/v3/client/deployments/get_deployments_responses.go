// Code generated by go-swagger; DO NOT EDIT.

package deployments

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

// GetDeploymentsReader is a Reader for the GetDeployments structure.
type GetDeploymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDeploymentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDeploymentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDeploymentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /deployments] GetDeployments", response, response.Code())
	}
}

// NewGetDeploymentsOK creates a GetDeploymentsOK with default headers values
func NewGetDeploymentsOK() *GetDeploymentsOK {
	return &GetDeploymentsOK{}
}

/*
GetDeploymentsOK describes a response with status code 200, with default header values.

OK
*/
type GetDeploymentsOK struct {
	Payload *models.ListDeploymentsResponse
}

// IsSuccess returns true when this get deployments o k response has a 2xx status code
func (o *GetDeploymentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get deployments o k response has a 3xx status code
func (o *GetDeploymentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployments o k response has a 4xx status code
func (o *GetDeploymentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get deployments o k response has a 5xx status code
func (o *GetDeploymentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployments o k response a status code equal to that given
func (o *GetDeploymentsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get deployments o k response
func (o *GetDeploymentsOK) Code() int {
	return 200
}

func (o *GetDeploymentsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /deployments][%d] getDeploymentsOK %s", 200, payload)
}

func (o *GetDeploymentsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /deployments][%d] getDeploymentsOK %s", 200, payload)
}

func (o *GetDeploymentsOK) GetPayload() *models.ListDeploymentsResponse {
	return o.Payload
}

func (o *GetDeploymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListDeploymentsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentsBadRequest creates a GetDeploymentsBadRequest with default headers values
func NewGetDeploymentsBadRequest() *GetDeploymentsBadRequest {
	return &GetDeploymentsBadRequest{}
}

/*
GetDeploymentsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDeploymentsBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get deployments bad request response has a 2xx status code
func (o *GetDeploymentsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployments bad request response has a 3xx status code
func (o *GetDeploymentsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployments bad request response has a 4xx status code
func (o *GetDeploymentsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployments bad request response has a 5xx status code
func (o *GetDeploymentsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployments bad request response a status code equal to that given
func (o *GetDeploymentsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get deployments bad request response
func (o *GetDeploymentsBadRequest) Code() int {
	return 400
}

func (o *GetDeploymentsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /deployments][%d] getDeploymentsBadRequest %s", 400, payload)
}

func (o *GetDeploymentsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /deployments][%d] getDeploymentsBadRequest %s", 400, payload)
}

func (o *GetDeploymentsBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetDeploymentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentsUnauthorized creates a GetDeploymentsUnauthorized with default headers values
func NewGetDeploymentsUnauthorized() *GetDeploymentsUnauthorized {
	return &GetDeploymentsUnauthorized{}
}

/*
GetDeploymentsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDeploymentsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get deployments unauthorized response has a 2xx status code
func (o *GetDeploymentsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployments unauthorized response has a 3xx status code
func (o *GetDeploymentsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployments unauthorized response has a 4xx status code
func (o *GetDeploymentsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployments unauthorized response has a 5xx status code
func (o *GetDeploymentsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployments unauthorized response a status code equal to that given
func (o *GetDeploymentsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get deployments unauthorized response
func (o *GetDeploymentsUnauthorized) Code() int {
	return 401
}

func (o *GetDeploymentsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /deployments][%d] getDeploymentsUnauthorized %s", 401, payload)
}

func (o *GetDeploymentsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /deployments][%d] getDeploymentsUnauthorized %s", 401, payload)
}

func (o *GetDeploymentsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetDeploymentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentsForbidden creates a GetDeploymentsForbidden with default headers values
func NewGetDeploymentsForbidden() *GetDeploymentsForbidden {
	return &GetDeploymentsForbidden{}
}

/*
GetDeploymentsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDeploymentsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get deployments forbidden response has a 2xx status code
func (o *GetDeploymentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployments forbidden response has a 3xx status code
func (o *GetDeploymentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployments forbidden response has a 4xx status code
func (o *GetDeploymentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployments forbidden response has a 5xx status code
func (o *GetDeploymentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployments forbidden response a status code equal to that given
func (o *GetDeploymentsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get deployments forbidden response
func (o *GetDeploymentsForbidden) Code() int {
	return 403
}

func (o *GetDeploymentsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /deployments][%d] getDeploymentsForbidden %s", 403, payload)
}

func (o *GetDeploymentsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /deployments][%d] getDeploymentsForbidden %s", 403, payload)
}

func (o *GetDeploymentsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetDeploymentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
