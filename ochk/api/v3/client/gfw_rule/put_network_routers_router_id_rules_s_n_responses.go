// Code generated by go-swagger; DO NOT EDIT.

package gfw_rule

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

// PutNetworkRoutersRouterIDRulesSNReader is a Reader for the PutNetworkRoutersRouterIDRulesSN structure.
type PutNetworkRoutersRouterIDRulesSNReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworkRoutersRouterIDRulesSNReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutNetworkRoutersRouterIDRulesSNOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutNetworkRoutersRouterIDRulesSNBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutNetworkRoutersRouterIDRulesSNUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutNetworkRoutersRouterIDRulesSNForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /network/routers/{routerId}/rules/s-n] PutNetworkRoutersRouterIDRulesSN", response, response.Code())
	}
}

// NewPutNetworkRoutersRouterIDRulesSNOK creates a PutNetworkRoutersRouterIDRulesSNOK with default headers values
func NewPutNetworkRoutersRouterIDRulesSNOK() *PutNetworkRoutersRouterIDRulesSNOK {
	return &PutNetworkRoutersRouterIDRulesSNOK{}
}

/*
PutNetworkRoutersRouterIDRulesSNOK describes a response with status code 200, with default header values.

OK
*/
type PutNetworkRoutersRouterIDRulesSNOK struct {
	Payload *models.CreateGfwRuleResponse
}

// IsSuccess returns true when this put network routers router Id rules s n o k response has a 2xx status code
func (o *PutNetworkRoutersRouterIDRulesSNOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put network routers router Id rules s n o k response has a 3xx status code
func (o *PutNetworkRoutersRouterIDRulesSNOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put network routers router Id rules s n o k response has a 4xx status code
func (o *PutNetworkRoutersRouterIDRulesSNOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put network routers router Id rules s n o k response has a 5xx status code
func (o *PutNetworkRoutersRouterIDRulesSNOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put network routers router Id rules s n o k response a status code equal to that given
func (o *PutNetworkRoutersRouterIDRulesSNOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put network routers router Id rules s n o k response
func (o *PutNetworkRoutersRouterIDRulesSNOK) Code() int {
	return 200
}

func (o *PutNetworkRoutersRouterIDRulesSNOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/routers/{routerId}/rules/s-n][%d] putNetworkRoutersRouterIdRulesSNOK %s", 200, payload)
}

func (o *PutNetworkRoutersRouterIDRulesSNOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/routers/{routerId}/rules/s-n][%d] putNetworkRoutersRouterIdRulesSNOK %s", 200, payload)
}

func (o *PutNetworkRoutersRouterIDRulesSNOK) GetPayload() *models.CreateGfwRuleResponse {
	return o.Payload
}

func (o *PutNetworkRoutersRouterIDRulesSNOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateGfwRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNetworkRoutersRouterIDRulesSNBadRequest creates a PutNetworkRoutersRouterIDRulesSNBadRequest with default headers values
func NewPutNetworkRoutersRouterIDRulesSNBadRequest() *PutNetworkRoutersRouterIDRulesSNBadRequest {
	return &PutNetworkRoutersRouterIDRulesSNBadRequest{}
}

/*
PutNetworkRoutersRouterIDRulesSNBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutNetworkRoutersRouterIDRulesSNBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put network routers router Id rules s n bad request response has a 2xx status code
func (o *PutNetworkRoutersRouterIDRulesSNBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put network routers router Id rules s n bad request response has a 3xx status code
func (o *PutNetworkRoutersRouterIDRulesSNBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put network routers router Id rules s n bad request response has a 4xx status code
func (o *PutNetworkRoutersRouterIDRulesSNBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put network routers router Id rules s n bad request response has a 5xx status code
func (o *PutNetworkRoutersRouterIDRulesSNBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put network routers router Id rules s n bad request response a status code equal to that given
func (o *PutNetworkRoutersRouterIDRulesSNBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put network routers router Id rules s n bad request response
func (o *PutNetworkRoutersRouterIDRulesSNBadRequest) Code() int {
	return 400
}

func (o *PutNetworkRoutersRouterIDRulesSNBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/routers/{routerId}/rules/s-n][%d] putNetworkRoutersRouterIdRulesSNBadRequest %s", 400, payload)
}

func (o *PutNetworkRoutersRouterIDRulesSNBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/routers/{routerId}/rules/s-n][%d] putNetworkRoutersRouterIdRulesSNBadRequest %s", 400, payload)
}

func (o *PutNetworkRoutersRouterIDRulesSNBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutNetworkRoutersRouterIDRulesSNBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNetworkRoutersRouterIDRulesSNUnauthorized creates a PutNetworkRoutersRouterIDRulesSNUnauthorized with default headers values
func NewPutNetworkRoutersRouterIDRulesSNUnauthorized() *PutNetworkRoutersRouterIDRulesSNUnauthorized {
	return &PutNetworkRoutersRouterIDRulesSNUnauthorized{}
}

/*
PutNetworkRoutersRouterIDRulesSNUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutNetworkRoutersRouterIDRulesSNUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put network routers router Id rules s n unauthorized response has a 2xx status code
func (o *PutNetworkRoutersRouterIDRulesSNUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put network routers router Id rules s n unauthorized response has a 3xx status code
func (o *PutNetworkRoutersRouterIDRulesSNUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put network routers router Id rules s n unauthorized response has a 4xx status code
func (o *PutNetworkRoutersRouterIDRulesSNUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put network routers router Id rules s n unauthorized response has a 5xx status code
func (o *PutNetworkRoutersRouterIDRulesSNUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put network routers router Id rules s n unauthorized response a status code equal to that given
func (o *PutNetworkRoutersRouterIDRulesSNUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put network routers router Id rules s n unauthorized response
func (o *PutNetworkRoutersRouterIDRulesSNUnauthorized) Code() int {
	return 401
}

func (o *PutNetworkRoutersRouterIDRulesSNUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/routers/{routerId}/rules/s-n][%d] putNetworkRoutersRouterIdRulesSNUnauthorized %s", 401, payload)
}

func (o *PutNetworkRoutersRouterIDRulesSNUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/routers/{routerId}/rules/s-n][%d] putNetworkRoutersRouterIdRulesSNUnauthorized %s", 401, payload)
}

func (o *PutNetworkRoutersRouterIDRulesSNUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutNetworkRoutersRouterIDRulesSNUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNetworkRoutersRouterIDRulesSNForbidden creates a PutNetworkRoutersRouterIDRulesSNForbidden with default headers values
func NewPutNetworkRoutersRouterIDRulesSNForbidden() *PutNetworkRoutersRouterIDRulesSNForbidden {
	return &PutNetworkRoutersRouterIDRulesSNForbidden{}
}

/*
PutNetworkRoutersRouterIDRulesSNForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutNetworkRoutersRouterIDRulesSNForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put network routers router Id rules s n forbidden response has a 2xx status code
func (o *PutNetworkRoutersRouterIDRulesSNForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put network routers router Id rules s n forbidden response has a 3xx status code
func (o *PutNetworkRoutersRouterIDRulesSNForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put network routers router Id rules s n forbidden response has a 4xx status code
func (o *PutNetworkRoutersRouterIDRulesSNForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put network routers router Id rules s n forbidden response has a 5xx status code
func (o *PutNetworkRoutersRouterIDRulesSNForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put network routers router Id rules s n forbidden response a status code equal to that given
func (o *PutNetworkRoutersRouterIDRulesSNForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put network routers router Id rules s n forbidden response
func (o *PutNetworkRoutersRouterIDRulesSNForbidden) Code() int {
	return 403
}

func (o *PutNetworkRoutersRouterIDRulesSNForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/routers/{routerId}/rules/s-n][%d] putNetworkRoutersRouterIdRulesSNForbidden %s", 403, payload)
}

func (o *PutNetworkRoutersRouterIDRulesSNForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/routers/{routerId}/rules/s-n][%d] putNetworkRoutersRouterIdRulesSNForbidden %s", 403, payload)
}

func (o *PutNetworkRoutersRouterIDRulesSNForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutNetworkRoutersRouterIDRulesSNForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
