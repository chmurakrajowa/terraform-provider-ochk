// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// GetProjectsProjectIDReader is a Reader for the GetProjectsProjectID structure.
type GetProjectsProjectIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsProjectIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectsProjectIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectsProjectIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectsProjectIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectsProjectIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectsProjectIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /projects/{projectId}] GetProjectsProjectID", response, response.Code())
	}
}

// NewGetProjectsProjectIDOK creates a GetProjectsProjectIDOK with default headers values
func NewGetProjectsProjectIDOK() *GetProjectsProjectIDOK {
	return &GetProjectsProjectIDOK{}
}

/*
GetProjectsProjectIDOK describes a response with status code 200, with default header values.

OK
*/
type GetProjectsProjectIDOK struct {
	Payload *models.GetProjectResponse
}

// IsSuccess returns true when this get projects project Id o k response has a 2xx status code
func (o *GetProjectsProjectIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get projects project Id o k response has a 3xx status code
func (o *GetProjectsProjectIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get projects project Id o k response has a 4xx status code
func (o *GetProjectsProjectIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get projects project Id o k response has a 5xx status code
func (o *GetProjectsProjectIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get projects project Id o k response a status code equal to that given
func (o *GetProjectsProjectIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get projects project Id o k response
func (o *GetProjectsProjectIDOK) Code() int {
	return 200
}

func (o *GetProjectsProjectIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectsProjectIdOK %s", 200, payload)
}

func (o *GetProjectsProjectIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectsProjectIdOK %s", 200, payload)
}

func (o *GetProjectsProjectIDOK) GetPayload() *models.GetProjectResponse {
	return o.Payload
}

func (o *GetProjectsProjectIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetProjectResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDBadRequest creates a GetProjectsProjectIDBadRequest with default headers values
func NewGetProjectsProjectIDBadRequest() *GetProjectsProjectIDBadRequest {
	return &GetProjectsProjectIDBadRequest{}
}

/*
GetProjectsProjectIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetProjectsProjectIDBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get projects project Id bad request response has a 2xx status code
func (o *GetProjectsProjectIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get projects project Id bad request response has a 3xx status code
func (o *GetProjectsProjectIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get projects project Id bad request response has a 4xx status code
func (o *GetProjectsProjectIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get projects project Id bad request response has a 5xx status code
func (o *GetProjectsProjectIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get projects project Id bad request response a status code equal to that given
func (o *GetProjectsProjectIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get projects project Id bad request response
func (o *GetProjectsProjectIDBadRequest) Code() int {
	return 400
}

func (o *GetProjectsProjectIDBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectsProjectIdBadRequest %s", 400, payload)
}

func (o *GetProjectsProjectIDBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectsProjectIdBadRequest %s", 400, payload)
}

func (o *GetProjectsProjectIDBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetProjectsProjectIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDUnauthorized creates a GetProjectsProjectIDUnauthorized with default headers values
func NewGetProjectsProjectIDUnauthorized() *GetProjectsProjectIDUnauthorized {
	return &GetProjectsProjectIDUnauthorized{}
}

/*
GetProjectsProjectIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetProjectsProjectIDUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get projects project Id unauthorized response has a 2xx status code
func (o *GetProjectsProjectIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get projects project Id unauthorized response has a 3xx status code
func (o *GetProjectsProjectIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get projects project Id unauthorized response has a 4xx status code
func (o *GetProjectsProjectIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get projects project Id unauthorized response has a 5xx status code
func (o *GetProjectsProjectIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get projects project Id unauthorized response a status code equal to that given
func (o *GetProjectsProjectIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get projects project Id unauthorized response
func (o *GetProjectsProjectIDUnauthorized) Code() int {
	return 401
}

func (o *GetProjectsProjectIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectsProjectIdUnauthorized %s", 401, payload)
}

func (o *GetProjectsProjectIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectsProjectIdUnauthorized %s", 401, payload)
}

func (o *GetProjectsProjectIDUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetProjectsProjectIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDForbidden creates a GetProjectsProjectIDForbidden with default headers values
func NewGetProjectsProjectIDForbidden() *GetProjectsProjectIDForbidden {
	return &GetProjectsProjectIDForbidden{}
}

/*
GetProjectsProjectIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetProjectsProjectIDForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get projects project Id forbidden response has a 2xx status code
func (o *GetProjectsProjectIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get projects project Id forbidden response has a 3xx status code
func (o *GetProjectsProjectIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get projects project Id forbidden response has a 4xx status code
func (o *GetProjectsProjectIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get projects project Id forbidden response has a 5xx status code
func (o *GetProjectsProjectIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get projects project Id forbidden response a status code equal to that given
func (o *GetProjectsProjectIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get projects project Id forbidden response
func (o *GetProjectsProjectIDForbidden) Code() int {
	return 403
}

func (o *GetProjectsProjectIDForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectsProjectIdForbidden %s", 403, payload)
}

func (o *GetProjectsProjectIDForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectsProjectIdForbidden %s", 403, payload)
}

func (o *GetProjectsProjectIDForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetProjectsProjectIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDNotFound creates a GetProjectsProjectIDNotFound with default headers values
func NewGetProjectsProjectIDNotFound() *GetProjectsProjectIDNotFound {
	return &GetProjectsProjectIDNotFound{}
}

/*
GetProjectsProjectIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetProjectsProjectIDNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get projects project Id not found response has a 2xx status code
func (o *GetProjectsProjectIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get projects project Id not found response has a 3xx status code
func (o *GetProjectsProjectIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get projects project Id not found response has a 4xx status code
func (o *GetProjectsProjectIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get projects project Id not found response has a 5xx status code
func (o *GetProjectsProjectIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get projects project Id not found response a status code equal to that given
func (o *GetProjectsProjectIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get projects project Id not found response
func (o *GetProjectsProjectIDNotFound) Code() int {
	return 404
}

func (o *GetProjectsProjectIDNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectsProjectIdNotFound %s", 404, payload)
}

func (o *GetProjectsProjectIDNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectsProjectIdNotFound %s", 404, payload)
}

func (o *GetProjectsProjectIDNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetProjectsProjectIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
