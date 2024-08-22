// Code generated by go-swagger; DO NOT EDIT.

package log_stats_user

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

// PostLogStatsUserLogCategoryIDReader is a Reader for the PostLogStatsUserLogCategoryID structure.
type PostLogStatsUserLogCategoryIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLogStatsUserLogCategoryIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLogStatsUserLogCategoryIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLogStatsUserLogCategoryIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostLogStatsUserLogCategoryIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLogStatsUserLogCategoryIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /log/stats/user/{logCategoryId}] PostLogStatsUserLogCategoryID", response, response.Code())
	}
}

// NewPostLogStatsUserLogCategoryIDOK creates a PostLogStatsUserLogCategoryIDOK with default headers values
func NewPostLogStatsUserLogCategoryIDOK() *PostLogStatsUserLogCategoryIDOK {
	return &PostLogStatsUserLogCategoryIDOK{}
}

/*
PostLogStatsUserLogCategoryIDOK describes a response with status code 200, with default header values.

OK
*/
type PostLogStatsUserLogCategoryIDOK struct {
	Payload *models.GetLogCategoryStats
}

// IsSuccess returns true when this post log stats user log category Id o k response has a 2xx status code
func (o *PostLogStatsUserLogCategoryIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post log stats user log category Id o k response has a 3xx status code
func (o *PostLogStatsUserLogCategoryIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post log stats user log category Id o k response has a 4xx status code
func (o *PostLogStatsUserLogCategoryIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post log stats user log category Id o k response has a 5xx status code
func (o *PostLogStatsUserLogCategoryIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post log stats user log category Id o k response a status code equal to that given
func (o *PostLogStatsUserLogCategoryIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post log stats user log category Id o k response
func (o *PostLogStatsUserLogCategoryIDOK) Code() int {
	return 200
}

func (o *PostLogStatsUserLogCategoryIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /log/stats/user/{logCategoryId}][%d] postLogStatsUserLogCategoryIdOK %s", 200, payload)
}

func (o *PostLogStatsUserLogCategoryIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /log/stats/user/{logCategoryId}][%d] postLogStatsUserLogCategoryIdOK %s", 200, payload)
}

func (o *PostLogStatsUserLogCategoryIDOK) GetPayload() *models.GetLogCategoryStats {
	return o.Payload
}

func (o *PostLogStatsUserLogCategoryIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetLogCategoryStats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLogStatsUserLogCategoryIDBadRequest creates a PostLogStatsUserLogCategoryIDBadRequest with default headers values
func NewPostLogStatsUserLogCategoryIDBadRequest() *PostLogStatsUserLogCategoryIDBadRequest {
	return &PostLogStatsUserLogCategoryIDBadRequest{}
}

/*
PostLogStatsUserLogCategoryIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostLogStatsUserLogCategoryIDBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this post log stats user log category Id bad request response has a 2xx status code
func (o *PostLogStatsUserLogCategoryIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post log stats user log category Id bad request response has a 3xx status code
func (o *PostLogStatsUserLogCategoryIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post log stats user log category Id bad request response has a 4xx status code
func (o *PostLogStatsUserLogCategoryIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post log stats user log category Id bad request response has a 5xx status code
func (o *PostLogStatsUserLogCategoryIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post log stats user log category Id bad request response a status code equal to that given
func (o *PostLogStatsUserLogCategoryIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post log stats user log category Id bad request response
func (o *PostLogStatsUserLogCategoryIDBadRequest) Code() int {
	return 400
}

func (o *PostLogStatsUserLogCategoryIDBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /log/stats/user/{logCategoryId}][%d] postLogStatsUserLogCategoryIdBadRequest %s", 400, payload)
}

func (o *PostLogStatsUserLogCategoryIDBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /log/stats/user/{logCategoryId}][%d] postLogStatsUserLogCategoryIdBadRequest %s", 400, payload)
}

func (o *PostLogStatsUserLogCategoryIDBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PostLogStatsUserLogCategoryIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLogStatsUserLogCategoryIDUnauthorized creates a PostLogStatsUserLogCategoryIDUnauthorized with default headers values
func NewPostLogStatsUserLogCategoryIDUnauthorized() *PostLogStatsUserLogCategoryIDUnauthorized {
	return &PostLogStatsUserLogCategoryIDUnauthorized{}
}

/*
PostLogStatsUserLogCategoryIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostLogStatsUserLogCategoryIDUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this post log stats user log category Id unauthorized response has a 2xx status code
func (o *PostLogStatsUserLogCategoryIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post log stats user log category Id unauthorized response has a 3xx status code
func (o *PostLogStatsUserLogCategoryIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post log stats user log category Id unauthorized response has a 4xx status code
func (o *PostLogStatsUserLogCategoryIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post log stats user log category Id unauthorized response has a 5xx status code
func (o *PostLogStatsUserLogCategoryIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post log stats user log category Id unauthorized response a status code equal to that given
func (o *PostLogStatsUserLogCategoryIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post log stats user log category Id unauthorized response
func (o *PostLogStatsUserLogCategoryIDUnauthorized) Code() int {
	return 401
}

func (o *PostLogStatsUserLogCategoryIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /log/stats/user/{logCategoryId}][%d] postLogStatsUserLogCategoryIdUnauthorized %s", 401, payload)
}

func (o *PostLogStatsUserLogCategoryIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /log/stats/user/{logCategoryId}][%d] postLogStatsUserLogCategoryIdUnauthorized %s", 401, payload)
}

func (o *PostLogStatsUserLogCategoryIDUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PostLogStatsUserLogCategoryIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLogStatsUserLogCategoryIDForbidden creates a PostLogStatsUserLogCategoryIDForbidden with default headers values
func NewPostLogStatsUserLogCategoryIDForbidden() *PostLogStatsUserLogCategoryIDForbidden {
	return &PostLogStatsUserLogCategoryIDForbidden{}
}

/*
PostLogStatsUserLogCategoryIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostLogStatsUserLogCategoryIDForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this post log stats user log category Id forbidden response has a 2xx status code
func (o *PostLogStatsUserLogCategoryIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post log stats user log category Id forbidden response has a 3xx status code
func (o *PostLogStatsUserLogCategoryIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post log stats user log category Id forbidden response has a 4xx status code
func (o *PostLogStatsUserLogCategoryIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post log stats user log category Id forbidden response has a 5xx status code
func (o *PostLogStatsUserLogCategoryIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post log stats user log category Id forbidden response a status code equal to that given
func (o *PostLogStatsUserLogCategoryIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post log stats user log category Id forbidden response
func (o *PostLogStatsUserLogCategoryIDForbidden) Code() int {
	return 403
}

func (o *PostLogStatsUserLogCategoryIDForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /log/stats/user/{logCategoryId}][%d] postLogStatsUserLogCategoryIdForbidden %s", 403, payload)
}

func (o *PostLogStatsUserLogCategoryIDForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /log/stats/user/{logCategoryId}][%d] postLogStatsUserLogCategoryIdForbidden %s", 403, payload)
}

func (o *PostLogStatsUserLogCategoryIDForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PostLogStatsUserLogCategoryIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
