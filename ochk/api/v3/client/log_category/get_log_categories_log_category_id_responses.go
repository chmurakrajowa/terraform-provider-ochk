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

// GetLogCategoriesLogCategoryIDReader is a Reader for the GetLogCategoriesLogCategoryID structure.
type GetLogCategoriesLogCategoryIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogCategoriesLogCategoryIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogCategoriesLogCategoryIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLogCategoriesLogCategoryIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLogCategoriesLogCategoryIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLogCategoriesLogCategoryIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /log/categories/{logCategoryId}] GetLogCategoriesLogCategoryID", response, response.Code())
	}
}

// NewGetLogCategoriesLogCategoryIDOK creates a GetLogCategoriesLogCategoryIDOK with default headers values
func NewGetLogCategoriesLogCategoryIDOK() *GetLogCategoriesLogCategoryIDOK {
	return &GetLogCategoriesLogCategoryIDOK{}
}

/*
GetLogCategoriesLogCategoryIDOK describes a response with status code 200, with default header values.

OK
*/
type GetLogCategoriesLogCategoryIDOK struct {
	Payload *models.GetLogCategoryResponse
}

// IsSuccess returns true when this get log categories log category Id o k response has a 2xx status code
func (o *GetLogCategoriesLogCategoryIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get log categories log category Id o k response has a 3xx status code
func (o *GetLogCategoriesLogCategoryIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get log categories log category Id o k response has a 4xx status code
func (o *GetLogCategoriesLogCategoryIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get log categories log category Id o k response has a 5xx status code
func (o *GetLogCategoriesLogCategoryIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get log categories log category Id o k response a status code equal to that given
func (o *GetLogCategoriesLogCategoryIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get log categories log category Id o k response
func (o *GetLogCategoriesLogCategoryIDOK) Code() int {
	return 200
}

func (o *GetLogCategoriesLogCategoryIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /log/categories/{logCategoryId}][%d] getLogCategoriesLogCategoryIdOK %s", 200, payload)
}

func (o *GetLogCategoriesLogCategoryIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /log/categories/{logCategoryId}][%d] getLogCategoriesLogCategoryIdOK %s", 200, payload)
}

func (o *GetLogCategoriesLogCategoryIDOK) GetPayload() *models.GetLogCategoryResponse {
	return o.Payload
}

func (o *GetLogCategoriesLogCategoryIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetLogCategoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogCategoriesLogCategoryIDBadRequest creates a GetLogCategoriesLogCategoryIDBadRequest with default headers values
func NewGetLogCategoriesLogCategoryIDBadRequest() *GetLogCategoriesLogCategoryIDBadRequest {
	return &GetLogCategoriesLogCategoryIDBadRequest{}
}

/*
GetLogCategoriesLogCategoryIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetLogCategoriesLogCategoryIDBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get log categories log category Id bad request response has a 2xx status code
func (o *GetLogCategoriesLogCategoryIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get log categories log category Id bad request response has a 3xx status code
func (o *GetLogCategoriesLogCategoryIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get log categories log category Id bad request response has a 4xx status code
func (o *GetLogCategoriesLogCategoryIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get log categories log category Id bad request response has a 5xx status code
func (o *GetLogCategoriesLogCategoryIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get log categories log category Id bad request response a status code equal to that given
func (o *GetLogCategoriesLogCategoryIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get log categories log category Id bad request response
func (o *GetLogCategoriesLogCategoryIDBadRequest) Code() int {
	return 400
}

func (o *GetLogCategoriesLogCategoryIDBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /log/categories/{logCategoryId}][%d] getLogCategoriesLogCategoryIdBadRequest %s", 400, payload)
}

func (o *GetLogCategoriesLogCategoryIDBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /log/categories/{logCategoryId}][%d] getLogCategoriesLogCategoryIdBadRequest %s", 400, payload)
}

func (o *GetLogCategoriesLogCategoryIDBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetLogCategoriesLogCategoryIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogCategoriesLogCategoryIDUnauthorized creates a GetLogCategoriesLogCategoryIDUnauthorized with default headers values
func NewGetLogCategoriesLogCategoryIDUnauthorized() *GetLogCategoriesLogCategoryIDUnauthorized {
	return &GetLogCategoriesLogCategoryIDUnauthorized{}
}

/*
GetLogCategoriesLogCategoryIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetLogCategoriesLogCategoryIDUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get log categories log category Id unauthorized response has a 2xx status code
func (o *GetLogCategoriesLogCategoryIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get log categories log category Id unauthorized response has a 3xx status code
func (o *GetLogCategoriesLogCategoryIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get log categories log category Id unauthorized response has a 4xx status code
func (o *GetLogCategoriesLogCategoryIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get log categories log category Id unauthorized response has a 5xx status code
func (o *GetLogCategoriesLogCategoryIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get log categories log category Id unauthorized response a status code equal to that given
func (o *GetLogCategoriesLogCategoryIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get log categories log category Id unauthorized response
func (o *GetLogCategoriesLogCategoryIDUnauthorized) Code() int {
	return 401
}

func (o *GetLogCategoriesLogCategoryIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /log/categories/{logCategoryId}][%d] getLogCategoriesLogCategoryIdUnauthorized %s", 401, payload)
}

func (o *GetLogCategoriesLogCategoryIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /log/categories/{logCategoryId}][%d] getLogCategoriesLogCategoryIdUnauthorized %s", 401, payload)
}

func (o *GetLogCategoriesLogCategoryIDUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetLogCategoriesLogCategoryIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogCategoriesLogCategoryIDForbidden creates a GetLogCategoriesLogCategoryIDForbidden with default headers values
func NewGetLogCategoriesLogCategoryIDForbidden() *GetLogCategoriesLogCategoryIDForbidden {
	return &GetLogCategoriesLogCategoryIDForbidden{}
}

/*
GetLogCategoriesLogCategoryIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetLogCategoriesLogCategoryIDForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get log categories log category Id forbidden response has a 2xx status code
func (o *GetLogCategoriesLogCategoryIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get log categories log category Id forbidden response has a 3xx status code
func (o *GetLogCategoriesLogCategoryIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get log categories log category Id forbidden response has a 4xx status code
func (o *GetLogCategoriesLogCategoryIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get log categories log category Id forbidden response has a 5xx status code
func (o *GetLogCategoriesLogCategoryIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get log categories log category Id forbidden response a status code equal to that given
func (o *GetLogCategoriesLogCategoryIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get log categories log category Id forbidden response
func (o *GetLogCategoriesLogCategoryIDForbidden) Code() int {
	return 403
}

func (o *GetLogCategoriesLogCategoryIDForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /log/categories/{logCategoryId}][%d] getLogCategoriesLogCategoryIdForbidden %s", 403, payload)
}

func (o *GetLogCategoriesLogCategoryIDForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /log/categories/{logCategoryId}][%d] getLogCategoriesLogCategoryIdForbidden %s", 403, payload)
}

func (o *GetLogCategoriesLogCategoryIDForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetLogCategoriesLogCategoryIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
