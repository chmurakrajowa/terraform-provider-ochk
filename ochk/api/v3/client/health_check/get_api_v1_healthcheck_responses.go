// Code generated by go-swagger; DO NOT EDIT.

package health_check

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

// GetAPIV1HealthcheckReader is a Reader for the GetAPIV1Healthcheck structure.
type GetAPIV1HealthcheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV1HealthcheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV1HealthcheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV1HealthcheckBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/healthcheck] GetAPIV1Healthcheck", response, response.Code())
	}
}

// NewGetAPIV1HealthcheckOK creates a GetAPIV1HealthcheckOK with default headers values
func NewGetAPIV1HealthcheckOK() *GetAPIV1HealthcheckOK {
	return &GetAPIV1HealthcheckOK{}
}

/*
GetAPIV1HealthcheckOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIV1HealthcheckOK struct {
	Payload *models.GetHealthCheckResponse
}

// IsSuccess returns true when this get Api v1 healthcheck o k response has a 2xx status code
func (o *GetAPIV1HealthcheckOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api v1 healthcheck o k response has a 3xx status code
func (o *GetAPIV1HealthcheckOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v1 healthcheck o k response has a 4xx status code
func (o *GetAPIV1HealthcheckOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api v1 healthcheck o k response has a 5xx status code
func (o *GetAPIV1HealthcheckOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v1 healthcheck o k response a status code equal to that given
func (o *GetAPIV1HealthcheckOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api v1 healthcheck o k response
func (o *GetAPIV1HealthcheckOK) Code() int {
	return 200
}

func (o *GetAPIV1HealthcheckOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/healthcheck][%d] getApiV1HealthcheckOK %s", 200, payload)
}

func (o *GetAPIV1HealthcheckOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/healthcheck][%d] getApiV1HealthcheckOK %s", 200, payload)
}

func (o *GetAPIV1HealthcheckOK) GetPayload() *models.GetHealthCheckResponse {
	return o.Payload
}

func (o *GetAPIV1HealthcheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetHealthCheckResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV1HealthcheckBadRequest creates a GetAPIV1HealthcheckBadRequest with default headers values
func NewGetAPIV1HealthcheckBadRequest() *GetAPIV1HealthcheckBadRequest {
	return &GetAPIV1HealthcheckBadRequest{}
}

/*
GetAPIV1HealthcheckBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIV1HealthcheckBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get Api v1 healthcheck bad request response has a 2xx status code
func (o *GetAPIV1HealthcheckBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api v1 healthcheck bad request response has a 3xx status code
func (o *GetAPIV1HealthcheckBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v1 healthcheck bad request response has a 4xx status code
func (o *GetAPIV1HealthcheckBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Api v1 healthcheck bad request response has a 5xx status code
func (o *GetAPIV1HealthcheckBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v1 healthcheck bad request response a status code equal to that given
func (o *GetAPIV1HealthcheckBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get Api v1 healthcheck bad request response
func (o *GetAPIV1HealthcheckBadRequest) Code() int {
	return 400
}

func (o *GetAPIV1HealthcheckBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/healthcheck][%d] getApiV1HealthcheckBadRequest %s", 400, payload)
}

func (o *GetAPIV1HealthcheckBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/healthcheck][%d] getApiV1HealthcheckBadRequest %s", 400, payload)
}

func (o *GetAPIV1HealthcheckBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetAPIV1HealthcheckBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
