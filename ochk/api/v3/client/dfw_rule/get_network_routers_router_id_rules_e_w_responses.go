// Code generated by go-swagger; DO NOT EDIT.

package dfw_rule

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

// GetNetworkRoutersRouterIDRulesEWReader is a Reader for the GetNetworkRoutersRouterIDRulesEW structure.
type GetNetworkRoutersRouterIDRulesEWReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkRoutersRouterIDRulesEWReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkRoutersRouterIDRulesEWOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNetworkRoutersRouterIDRulesEWBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetNetworkRoutersRouterIDRulesEWUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNetworkRoutersRouterIDRulesEWForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /network/routers/{routerId}/rules/e-w] GetNetworkRoutersRouterIDRulesEW", response, response.Code())
	}
}

// NewGetNetworkRoutersRouterIDRulesEWOK creates a GetNetworkRoutersRouterIDRulesEWOK with default headers values
func NewGetNetworkRoutersRouterIDRulesEWOK() *GetNetworkRoutersRouterIDRulesEWOK {
	return &GetNetworkRoutersRouterIDRulesEWOK{}
}

/*
GetNetworkRoutersRouterIDRulesEWOK describes a response with status code 200, with default header values.

OK
*/
type GetNetworkRoutersRouterIDRulesEWOK struct {
	Payload *models.ListDfwRulesResponse
}

// IsSuccess returns true when this get network routers router Id rules e w o k response has a 2xx status code
func (o *GetNetworkRoutersRouterIDRulesEWOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network routers router Id rules e w o k response has a 3xx status code
func (o *GetNetworkRoutersRouterIDRulesEWOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network routers router Id rules e w o k response has a 4xx status code
func (o *GetNetworkRoutersRouterIDRulesEWOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network routers router Id rules e w o k response has a 5xx status code
func (o *GetNetworkRoutersRouterIDRulesEWOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network routers router Id rules e w o k response a status code equal to that given
func (o *GetNetworkRoutersRouterIDRulesEWOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network routers router Id rules e w o k response
func (o *GetNetworkRoutersRouterIDRulesEWOK) Code() int {
	return 200
}

func (o *GetNetworkRoutersRouterIDRulesEWOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/routers/{routerId}/rules/e-w][%d] getNetworkRoutersRouterIdRulesEWOK %s", 200, payload)
}

func (o *GetNetworkRoutersRouterIDRulesEWOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/routers/{routerId}/rules/e-w][%d] getNetworkRoutersRouterIdRulesEWOK %s", 200, payload)
}

func (o *GetNetworkRoutersRouterIDRulesEWOK) GetPayload() *models.ListDfwRulesResponse {
	return o.Payload
}

func (o *GetNetworkRoutersRouterIDRulesEWOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListDfwRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkRoutersRouterIDRulesEWBadRequest creates a GetNetworkRoutersRouterIDRulesEWBadRequest with default headers values
func NewGetNetworkRoutersRouterIDRulesEWBadRequest() *GetNetworkRoutersRouterIDRulesEWBadRequest {
	return &GetNetworkRoutersRouterIDRulesEWBadRequest{}
}

/*
GetNetworkRoutersRouterIDRulesEWBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetNetworkRoutersRouterIDRulesEWBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get network routers router Id rules e w bad request response has a 2xx status code
func (o *GetNetworkRoutersRouterIDRulesEWBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network routers router Id rules e w bad request response has a 3xx status code
func (o *GetNetworkRoutersRouterIDRulesEWBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network routers router Id rules e w bad request response has a 4xx status code
func (o *GetNetworkRoutersRouterIDRulesEWBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get network routers router Id rules e w bad request response has a 5xx status code
func (o *GetNetworkRoutersRouterIDRulesEWBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get network routers router Id rules e w bad request response a status code equal to that given
func (o *GetNetworkRoutersRouterIDRulesEWBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get network routers router Id rules e w bad request response
func (o *GetNetworkRoutersRouterIDRulesEWBadRequest) Code() int {
	return 400
}

func (o *GetNetworkRoutersRouterIDRulesEWBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/routers/{routerId}/rules/e-w][%d] getNetworkRoutersRouterIdRulesEWBadRequest %s", 400, payload)
}

func (o *GetNetworkRoutersRouterIDRulesEWBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/routers/{routerId}/rules/e-w][%d] getNetworkRoutersRouterIdRulesEWBadRequest %s", 400, payload)
}

func (o *GetNetworkRoutersRouterIDRulesEWBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetNetworkRoutersRouterIDRulesEWBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkRoutersRouterIDRulesEWUnauthorized creates a GetNetworkRoutersRouterIDRulesEWUnauthorized with default headers values
func NewGetNetworkRoutersRouterIDRulesEWUnauthorized() *GetNetworkRoutersRouterIDRulesEWUnauthorized {
	return &GetNetworkRoutersRouterIDRulesEWUnauthorized{}
}

/*
GetNetworkRoutersRouterIDRulesEWUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetNetworkRoutersRouterIDRulesEWUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get network routers router Id rules e w unauthorized response has a 2xx status code
func (o *GetNetworkRoutersRouterIDRulesEWUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network routers router Id rules e w unauthorized response has a 3xx status code
func (o *GetNetworkRoutersRouterIDRulesEWUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network routers router Id rules e w unauthorized response has a 4xx status code
func (o *GetNetworkRoutersRouterIDRulesEWUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get network routers router Id rules e w unauthorized response has a 5xx status code
func (o *GetNetworkRoutersRouterIDRulesEWUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get network routers router Id rules e w unauthorized response a status code equal to that given
func (o *GetNetworkRoutersRouterIDRulesEWUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get network routers router Id rules e w unauthorized response
func (o *GetNetworkRoutersRouterIDRulesEWUnauthorized) Code() int {
	return 401
}

func (o *GetNetworkRoutersRouterIDRulesEWUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/routers/{routerId}/rules/e-w][%d] getNetworkRoutersRouterIdRulesEWUnauthorized %s", 401, payload)
}

func (o *GetNetworkRoutersRouterIDRulesEWUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/routers/{routerId}/rules/e-w][%d] getNetworkRoutersRouterIdRulesEWUnauthorized %s", 401, payload)
}

func (o *GetNetworkRoutersRouterIDRulesEWUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetNetworkRoutersRouterIDRulesEWUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkRoutersRouterIDRulesEWForbidden creates a GetNetworkRoutersRouterIDRulesEWForbidden with default headers values
func NewGetNetworkRoutersRouterIDRulesEWForbidden() *GetNetworkRoutersRouterIDRulesEWForbidden {
	return &GetNetworkRoutersRouterIDRulesEWForbidden{}
}

/*
GetNetworkRoutersRouterIDRulesEWForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetNetworkRoutersRouterIDRulesEWForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get network routers router Id rules e w forbidden response has a 2xx status code
func (o *GetNetworkRoutersRouterIDRulesEWForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network routers router Id rules e w forbidden response has a 3xx status code
func (o *GetNetworkRoutersRouterIDRulesEWForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network routers router Id rules e w forbidden response has a 4xx status code
func (o *GetNetworkRoutersRouterIDRulesEWForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get network routers router Id rules e w forbidden response has a 5xx status code
func (o *GetNetworkRoutersRouterIDRulesEWForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get network routers router Id rules e w forbidden response a status code equal to that given
func (o *GetNetworkRoutersRouterIDRulesEWForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get network routers router Id rules e w forbidden response
func (o *GetNetworkRoutersRouterIDRulesEWForbidden) Code() int {
	return 403
}

func (o *GetNetworkRoutersRouterIDRulesEWForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/routers/{routerId}/rules/e-w][%d] getNetworkRoutersRouterIdRulesEWForbidden %s", 403, payload)
}

func (o *GetNetworkRoutersRouterIDRulesEWForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/routers/{routerId}/rules/e-w][%d] getNetworkRoutersRouterIdRulesEWForbidden %s", 403, payload)
}

func (o *GetNetworkRoutersRouterIDRulesEWForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetNetworkRoutersRouterIDRulesEWForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
