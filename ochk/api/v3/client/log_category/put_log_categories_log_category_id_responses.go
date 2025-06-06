// Code generated by go-swagger; DO NOT EDIT.

package log_category

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

// PutLogCategoriesLogCategoryIDReader is a Reader for the PutLogCategoriesLogCategoryID structure.
type PutLogCategoriesLogCategoryIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLogCategoriesLogCategoryIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutLogCategoriesLogCategoryIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutLogCategoriesLogCategoryIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutLogCategoriesLogCategoryIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutLogCategoriesLogCategoryIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /log/categories/{logCategoryId}] PutLogCategoriesLogCategoryID", response, response.Code())
	}
}

// NewPutLogCategoriesLogCategoryIDOK creates a PutLogCategoriesLogCategoryIDOK with default headers values
func NewPutLogCategoriesLogCategoryIDOK() *PutLogCategoriesLogCategoryIDOK {
	return &PutLogCategoriesLogCategoryIDOK{}
}

/*
PutLogCategoriesLogCategoryIDOK describes a response with status code 200, with default header values.

OK
*/
type PutLogCategoriesLogCategoryIDOK struct {
	Payload *models.UpdateLogCategoriesResponse
}

// IsSuccess returns true when this put log categories log category Id o k response has a 2xx status code
func (o *PutLogCategoriesLogCategoryIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put log categories log category Id o k response has a 3xx status code
func (o *PutLogCategoriesLogCategoryIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put log categories log category Id o k response has a 4xx status code
func (o *PutLogCategoriesLogCategoryIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put log categories log category Id o k response has a 5xx status code
func (o *PutLogCategoriesLogCategoryIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put log categories log category Id o k response a status code equal to that given
func (o *PutLogCategoriesLogCategoryIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put log categories log category Id o k response
func (o *PutLogCategoriesLogCategoryIDOK) Code() int {
	return 200
}

func (o *PutLogCategoriesLogCategoryIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /log/categories/{logCategoryId}][%d] putLogCategoriesLogCategoryIdOK %s", 200, payload)
}

func (o *PutLogCategoriesLogCategoryIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /log/categories/{logCategoryId}][%d] putLogCategoriesLogCategoryIdOK %s", 200, payload)
}

func (o *PutLogCategoriesLogCategoryIDOK) GetPayload() *models.UpdateLogCategoriesResponse {
	return o.Payload
}

func (o *PutLogCategoriesLogCategoryIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateLogCategoriesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLogCategoriesLogCategoryIDBadRequest creates a PutLogCategoriesLogCategoryIDBadRequest with default headers values
func NewPutLogCategoriesLogCategoryIDBadRequest() *PutLogCategoriesLogCategoryIDBadRequest {
	return &PutLogCategoriesLogCategoryIDBadRequest{}
}

/*
PutLogCategoriesLogCategoryIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutLogCategoriesLogCategoryIDBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put log categories log category Id bad request response has a 2xx status code
func (o *PutLogCategoriesLogCategoryIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put log categories log category Id bad request response has a 3xx status code
func (o *PutLogCategoriesLogCategoryIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put log categories log category Id bad request response has a 4xx status code
func (o *PutLogCategoriesLogCategoryIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put log categories log category Id bad request response has a 5xx status code
func (o *PutLogCategoriesLogCategoryIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put log categories log category Id bad request response a status code equal to that given
func (o *PutLogCategoriesLogCategoryIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put log categories log category Id bad request response
func (o *PutLogCategoriesLogCategoryIDBadRequest) Code() int {
	return 400
}

func (o *PutLogCategoriesLogCategoryIDBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /log/categories/{logCategoryId}][%d] putLogCategoriesLogCategoryIdBadRequest %s", 400, payload)
}

func (o *PutLogCategoriesLogCategoryIDBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /log/categories/{logCategoryId}][%d] putLogCategoriesLogCategoryIdBadRequest %s", 400, payload)
}

func (o *PutLogCategoriesLogCategoryIDBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutLogCategoriesLogCategoryIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLogCategoriesLogCategoryIDUnauthorized creates a PutLogCategoriesLogCategoryIDUnauthorized with default headers values
func NewPutLogCategoriesLogCategoryIDUnauthorized() *PutLogCategoriesLogCategoryIDUnauthorized {
	return &PutLogCategoriesLogCategoryIDUnauthorized{}
}

/*
PutLogCategoriesLogCategoryIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutLogCategoriesLogCategoryIDUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put log categories log category Id unauthorized response has a 2xx status code
func (o *PutLogCategoriesLogCategoryIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put log categories log category Id unauthorized response has a 3xx status code
func (o *PutLogCategoriesLogCategoryIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put log categories log category Id unauthorized response has a 4xx status code
func (o *PutLogCategoriesLogCategoryIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put log categories log category Id unauthorized response has a 5xx status code
func (o *PutLogCategoriesLogCategoryIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put log categories log category Id unauthorized response a status code equal to that given
func (o *PutLogCategoriesLogCategoryIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put log categories log category Id unauthorized response
func (o *PutLogCategoriesLogCategoryIDUnauthorized) Code() int {
	return 401
}

func (o *PutLogCategoriesLogCategoryIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /log/categories/{logCategoryId}][%d] putLogCategoriesLogCategoryIdUnauthorized %s", 401, payload)
}

func (o *PutLogCategoriesLogCategoryIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /log/categories/{logCategoryId}][%d] putLogCategoriesLogCategoryIdUnauthorized %s", 401, payload)
}

func (o *PutLogCategoriesLogCategoryIDUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutLogCategoriesLogCategoryIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLogCategoriesLogCategoryIDForbidden creates a PutLogCategoriesLogCategoryIDForbidden with default headers values
func NewPutLogCategoriesLogCategoryIDForbidden() *PutLogCategoriesLogCategoryIDForbidden {
	return &PutLogCategoriesLogCategoryIDForbidden{}
}

/*
PutLogCategoriesLogCategoryIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutLogCategoriesLogCategoryIDForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put log categories log category Id forbidden response has a 2xx status code
func (o *PutLogCategoriesLogCategoryIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put log categories log category Id forbidden response has a 3xx status code
func (o *PutLogCategoriesLogCategoryIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put log categories log category Id forbidden response has a 4xx status code
func (o *PutLogCategoriesLogCategoryIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put log categories log category Id forbidden response has a 5xx status code
func (o *PutLogCategoriesLogCategoryIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put log categories log category Id forbidden response a status code equal to that given
func (o *PutLogCategoriesLogCategoryIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put log categories log category Id forbidden response
func (o *PutLogCategoriesLogCategoryIDForbidden) Code() int {
	return 403
}

func (o *PutLogCategoriesLogCategoryIDForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /log/categories/{logCategoryId}][%d] putLogCategoriesLogCategoryIdForbidden %s", 403, payload)
}

func (o *PutLogCategoriesLogCategoryIDForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /log/categories/{logCategoryId}][%d] putLogCategoriesLogCategoryIdForbidden %s", 403, payload)
}

func (o *PutLogCategoriesLogCategoryIDForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutLogCategoriesLogCategoryIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
