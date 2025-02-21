// Code generated by go-swagger; DO NOT EDIT.

package ip_collection

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

// PutIpcsIPCollectionIDReader is a Reader for the PutIpcsIPCollectionID structure.
type PutIpcsIPCollectionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutIpcsIPCollectionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutIpcsIPCollectionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutIpcsIPCollectionIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutIpcsIPCollectionIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutIpcsIPCollectionIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutIpcsIPCollectionIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /ipcs/{ipCollectionId}] PutIpcsIPCollectionID", response, response.Code())
	}
}

// NewPutIpcsIPCollectionIDOK creates a PutIpcsIPCollectionIDOK with default headers values
func NewPutIpcsIPCollectionIDOK() *PutIpcsIPCollectionIDOK {
	return &PutIpcsIPCollectionIDOK{}
}

/*
PutIpcsIPCollectionIDOK describes a response with status code 200, with default header values.

OK
*/
type PutIpcsIPCollectionIDOK struct {
	Payload *models.UpdateIPCollectionResponse
}

// IsSuccess returns true when this put ipcs Ip collection Id o k response has a 2xx status code
func (o *PutIpcsIPCollectionIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put ipcs Ip collection Id o k response has a 3xx status code
func (o *PutIpcsIPCollectionIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put ipcs Ip collection Id o k response has a 4xx status code
func (o *PutIpcsIPCollectionIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put ipcs Ip collection Id o k response has a 5xx status code
func (o *PutIpcsIPCollectionIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put ipcs Ip collection Id o k response a status code equal to that given
func (o *PutIpcsIPCollectionIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put ipcs Ip collection Id o k response
func (o *PutIpcsIPCollectionIDOK) Code() int {
	return 200
}

func (o *PutIpcsIPCollectionIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipcs/{ipCollectionId}][%d] putIpcsIpCollectionIdOK %s", 200, payload)
}

func (o *PutIpcsIPCollectionIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipcs/{ipCollectionId}][%d] putIpcsIpCollectionIdOK %s", 200, payload)
}

func (o *PutIpcsIPCollectionIDOK) GetPayload() *models.UpdateIPCollectionResponse {
	return o.Payload
}

func (o *PutIpcsIPCollectionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateIPCollectionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIpcsIPCollectionIDBadRequest creates a PutIpcsIPCollectionIDBadRequest with default headers values
func NewPutIpcsIPCollectionIDBadRequest() *PutIpcsIPCollectionIDBadRequest {
	return &PutIpcsIPCollectionIDBadRequest{}
}

/*
PutIpcsIPCollectionIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutIpcsIPCollectionIDBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put ipcs Ip collection Id bad request response has a 2xx status code
func (o *PutIpcsIPCollectionIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put ipcs Ip collection Id bad request response has a 3xx status code
func (o *PutIpcsIPCollectionIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put ipcs Ip collection Id bad request response has a 4xx status code
func (o *PutIpcsIPCollectionIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put ipcs Ip collection Id bad request response has a 5xx status code
func (o *PutIpcsIPCollectionIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put ipcs Ip collection Id bad request response a status code equal to that given
func (o *PutIpcsIPCollectionIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put ipcs Ip collection Id bad request response
func (o *PutIpcsIPCollectionIDBadRequest) Code() int {
	return 400
}

func (o *PutIpcsIPCollectionIDBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipcs/{ipCollectionId}][%d] putIpcsIpCollectionIdBadRequest %s", 400, payload)
}

func (o *PutIpcsIPCollectionIDBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipcs/{ipCollectionId}][%d] putIpcsIpCollectionIdBadRequest %s", 400, payload)
}

func (o *PutIpcsIPCollectionIDBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutIpcsIPCollectionIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIpcsIPCollectionIDUnauthorized creates a PutIpcsIPCollectionIDUnauthorized with default headers values
func NewPutIpcsIPCollectionIDUnauthorized() *PutIpcsIPCollectionIDUnauthorized {
	return &PutIpcsIPCollectionIDUnauthorized{}
}

/*
PutIpcsIPCollectionIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutIpcsIPCollectionIDUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put ipcs Ip collection Id unauthorized response has a 2xx status code
func (o *PutIpcsIPCollectionIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put ipcs Ip collection Id unauthorized response has a 3xx status code
func (o *PutIpcsIPCollectionIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put ipcs Ip collection Id unauthorized response has a 4xx status code
func (o *PutIpcsIPCollectionIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put ipcs Ip collection Id unauthorized response has a 5xx status code
func (o *PutIpcsIPCollectionIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put ipcs Ip collection Id unauthorized response a status code equal to that given
func (o *PutIpcsIPCollectionIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put ipcs Ip collection Id unauthorized response
func (o *PutIpcsIPCollectionIDUnauthorized) Code() int {
	return 401
}

func (o *PutIpcsIPCollectionIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipcs/{ipCollectionId}][%d] putIpcsIpCollectionIdUnauthorized %s", 401, payload)
}

func (o *PutIpcsIPCollectionIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipcs/{ipCollectionId}][%d] putIpcsIpCollectionIdUnauthorized %s", 401, payload)
}

func (o *PutIpcsIPCollectionIDUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutIpcsIPCollectionIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIpcsIPCollectionIDForbidden creates a PutIpcsIPCollectionIDForbidden with default headers values
func NewPutIpcsIPCollectionIDForbidden() *PutIpcsIPCollectionIDForbidden {
	return &PutIpcsIPCollectionIDForbidden{}
}

/*
PutIpcsIPCollectionIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutIpcsIPCollectionIDForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put ipcs Ip collection Id forbidden response has a 2xx status code
func (o *PutIpcsIPCollectionIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put ipcs Ip collection Id forbidden response has a 3xx status code
func (o *PutIpcsIPCollectionIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put ipcs Ip collection Id forbidden response has a 4xx status code
func (o *PutIpcsIPCollectionIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put ipcs Ip collection Id forbidden response has a 5xx status code
func (o *PutIpcsIPCollectionIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put ipcs Ip collection Id forbidden response a status code equal to that given
func (o *PutIpcsIPCollectionIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put ipcs Ip collection Id forbidden response
func (o *PutIpcsIPCollectionIDForbidden) Code() int {
	return 403
}

func (o *PutIpcsIPCollectionIDForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipcs/{ipCollectionId}][%d] putIpcsIpCollectionIdForbidden %s", 403, payload)
}

func (o *PutIpcsIPCollectionIDForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipcs/{ipCollectionId}][%d] putIpcsIpCollectionIdForbidden %s", 403, payload)
}

func (o *PutIpcsIPCollectionIDForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutIpcsIPCollectionIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIpcsIPCollectionIDNotFound creates a PutIpcsIPCollectionIDNotFound with default headers values
func NewPutIpcsIPCollectionIDNotFound() *PutIpcsIPCollectionIDNotFound {
	return &PutIpcsIPCollectionIDNotFound{}
}

/*
PutIpcsIPCollectionIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutIpcsIPCollectionIDNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put ipcs Ip collection Id not found response has a 2xx status code
func (o *PutIpcsIPCollectionIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put ipcs Ip collection Id not found response has a 3xx status code
func (o *PutIpcsIPCollectionIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put ipcs Ip collection Id not found response has a 4xx status code
func (o *PutIpcsIPCollectionIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put ipcs Ip collection Id not found response has a 5xx status code
func (o *PutIpcsIPCollectionIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put ipcs Ip collection Id not found response a status code equal to that given
func (o *PutIpcsIPCollectionIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the put ipcs Ip collection Id not found response
func (o *PutIpcsIPCollectionIDNotFound) Code() int {
	return 404
}

func (o *PutIpcsIPCollectionIDNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipcs/{ipCollectionId}][%d] putIpcsIpCollectionIdNotFound %s", 404, payload)
}

func (o *PutIpcsIPCollectionIDNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipcs/{ipCollectionId}][%d] putIpcsIpCollectionIdNotFound %s", 404, payload)
}

func (o *PutIpcsIPCollectionIDNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutIpcsIPCollectionIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
