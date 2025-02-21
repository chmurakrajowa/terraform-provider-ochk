// Code generated by go-swagger; DO NOT EDIT.

package nat_rule

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

// PutNetworkNatRulesRuleIDReader is a Reader for the PutNetworkNatRulesRuleID structure.
type PutNetworkNatRulesRuleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworkNatRulesRuleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutNetworkNatRulesRuleIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutNetworkNatRulesRuleIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutNetworkNatRulesRuleIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutNetworkNatRulesRuleIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutNetworkNatRulesRuleIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /network/nat/rules/{ruleId}] PutNetworkNatRulesRuleID", response, response.Code())
	}
}

// NewPutNetworkNatRulesRuleIDOK creates a PutNetworkNatRulesRuleIDOK with default headers values
func NewPutNetworkNatRulesRuleIDOK() *PutNetworkNatRulesRuleIDOK {
	return &PutNetworkNatRulesRuleIDOK{}
}

/*
PutNetworkNatRulesRuleIDOK describes a response with status code 200, with default header values.

OK
*/
type PutNetworkNatRulesRuleIDOK struct {
	Payload *models.UpdateNATRuleResponse
}

// IsSuccess returns true when this put network nat rules rule Id o k response has a 2xx status code
func (o *PutNetworkNatRulesRuleIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put network nat rules rule Id o k response has a 3xx status code
func (o *PutNetworkNatRulesRuleIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put network nat rules rule Id o k response has a 4xx status code
func (o *PutNetworkNatRulesRuleIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put network nat rules rule Id o k response has a 5xx status code
func (o *PutNetworkNatRulesRuleIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put network nat rules rule Id o k response a status code equal to that given
func (o *PutNetworkNatRulesRuleIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put network nat rules rule Id o k response
func (o *PutNetworkNatRulesRuleIDOK) Code() int {
	return 200
}

func (o *PutNetworkNatRulesRuleIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/nat/rules/{ruleId}][%d] putNetworkNatRulesRuleIdOK %s", 200, payload)
}

func (o *PutNetworkNatRulesRuleIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/nat/rules/{ruleId}][%d] putNetworkNatRulesRuleIdOK %s", 200, payload)
}

func (o *PutNetworkNatRulesRuleIDOK) GetPayload() *models.UpdateNATRuleResponse {
	return o.Payload
}

func (o *PutNetworkNatRulesRuleIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateNATRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNetworkNatRulesRuleIDBadRequest creates a PutNetworkNatRulesRuleIDBadRequest with default headers values
func NewPutNetworkNatRulesRuleIDBadRequest() *PutNetworkNatRulesRuleIDBadRequest {
	return &PutNetworkNatRulesRuleIDBadRequest{}
}

/*
PutNetworkNatRulesRuleIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutNetworkNatRulesRuleIDBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put network nat rules rule Id bad request response has a 2xx status code
func (o *PutNetworkNatRulesRuleIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put network nat rules rule Id bad request response has a 3xx status code
func (o *PutNetworkNatRulesRuleIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put network nat rules rule Id bad request response has a 4xx status code
func (o *PutNetworkNatRulesRuleIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put network nat rules rule Id bad request response has a 5xx status code
func (o *PutNetworkNatRulesRuleIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put network nat rules rule Id bad request response a status code equal to that given
func (o *PutNetworkNatRulesRuleIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put network nat rules rule Id bad request response
func (o *PutNetworkNatRulesRuleIDBadRequest) Code() int {
	return 400
}

func (o *PutNetworkNatRulesRuleIDBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/nat/rules/{ruleId}][%d] putNetworkNatRulesRuleIdBadRequest %s", 400, payload)
}

func (o *PutNetworkNatRulesRuleIDBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/nat/rules/{ruleId}][%d] putNetworkNatRulesRuleIdBadRequest %s", 400, payload)
}

func (o *PutNetworkNatRulesRuleIDBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutNetworkNatRulesRuleIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNetworkNatRulesRuleIDUnauthorized creates a PutNetworkNatRulesRuleIDUnauthorized with default headers values
func NewPutNetworkNatRulesRuleIDUnauthorized() *PutNetworkNatRulesRuleIDUnauthorized {
	return &PutNetworkNatRulesRuleIDUnauthorized{}
}

/*
PutNetworkNatRulesRuleIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutNetworkNatRulesRuleIDUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put network nat rules rule Id unauthorized response has a 2xx status code
func (o *PutNetworkNatRulesRuleIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put network nat rules rule Id unauthorized response has a 3xx status code
func (o *PutNetworkNatRulesRuleIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put network nat rules rule Id unauthorized response has a 4xx status code
func (o *PutNetworkNatRulesRuleIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put network nat rules rule Id unauthorized response has a 5xx status code
func (o *PutNetworkNatRulesRuleIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put network nat rules rule Id unauthorized response a status code equal to that given
func (o *PutNetworkNatRulesRuleIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put network nat rules rule Id unauthorized response
func (o *PutNetworkNatRulesRuleIDUnauthorized) Code() int {
	return 401
}

func (o *PutNetworkNatRulesRuleIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/nat/rules/{ruleId}][%d] putNetworkNatRulesRuleIdUnauthorized %s", 401, payload)
}

func (o *PutNetworkNatRulesRuleIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/nat/rules/{ruleId}][%d] putNetworkNatRulesRuleIdUnauthorized %s", 401, payload)
}

func (o *PutNetworkNatRulesRuleIDUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutNetworkNatRulesRuleIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNetworkNatRulesRuleIDForbidden creates a PutNetworkNatRulesRuleIDForbidden with default headers values
func NewPutNetworkNatRulesRuleIDForbidden() *PutNetworkNatRulesRuleIDForbidden {
	return &PutNetworkNatRulesRuleIDForbidden{}
}

/*
PutNetworkNatRulesRuleIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutNetworkNatRulesRuleIDForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put network nat rules rule Id forbidden response has a 2xx status code
func (o *PutNetworkNatRulesRuleIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put network nat rules rule Id forbidden response has a 3xx status code
func (o *PutNetworkNatRulesRuleIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put network nat rules rule Id forbidden response has a 4xx status code
func (o *PutNetworkNatRulesRuleIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put network nat rules rule Id forbidden response has a 5xx status code
func (o *PutNetworkNatRulesRuleIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put network nat rules rule Id forbidden response a status code equal to that given
func (o *PutNetworkNatRulesRuleIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put network nat rules rule Id forbidden response
func (o *PutNetworkNatRulesRuleIDForbidden) Code() int {
	return 403
}

func (o *PutNetworkNatRulesRuleIDForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/nat/rules/{ruleId}][%d] putNetworkNatRulesRuleIdForbidden %s", 403, payload)
}

func (o *PutNetworkNatRulesRuleIDForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/nat/rules/{ruleId}][%d] putNetworkNatRulesRuleIdForbidden %s", 403, payload)
}

func (o *PutNetworkNatRulesRuleIDForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutNetworkNatRulesRuleIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNetworkNatRulesRuleIDNotFound creates a PutNetworkNatRulesRuleIDNotFound with default headers values
func NewPutNetworkNatRulesRuleIDNotFound() *PutNetworkNatRulesRuleIDNotFound {
	return &PutNetworkNatRulesRuleIDNotFound{}
}

/*
PutNetworkNatRulesRuleIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutNetworkNatRulesRuleIDNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put network nat rules rule Id not found response has a 2xx status code
func (o *PutNetworkNatRulesRuleIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put network nat rules rule Id not found response has a 3xx status code
func (o *PutNetworkNatRulesRuleIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put network nat rules rule Id not found response has a 4xx status code
func (o *PutNetworkNatRulesRuleIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put network nat rules rule Id not found response has a 5xx status code
func (o *PutNetworkNatRulesRuleIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put network nat rules rule Id not found response a status code equal to that given
func (o *PutNetworkNatRulesRuleIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the put network nat rules rule Id not found response
func (o *PutNetworkNatRulesRuleIDNotFound) Code() int {
	return 404
}

func (o *PutNetworkNatRulesRuleIDNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/nat/rules/{ruleId}][%d] putNetworkNatRulesRuleIdNotFound %s", 404, payload)
}

func (o *PutNetworkNatRulesRuleIDNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/nat/rules/{ruleId}][%d] putNetworkNatRulesRuleIdNotFound %s", 404, payload)
}

func (o *PutNetworkNatRulesRuleIDNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutNetworkNatRulesRuleIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
