// Code generated by go-swagger; DO NOT EDIT.

package tags

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

// DeleteTagsTagIDReader is a Reader for the DeleteTagsTagID structure.
type DeleteTagsTagIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTagsTagIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTagsTagIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTagsTagIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteTagsTagIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTagsTagIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTagsTagIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /tags/{tagId}] DeleteTagsTagID", response, response.Code())
	}
}

// NewDeleteTagsTagIDOK creates a DeleteTagsTagIDOK with default headers values
func NewDeleteTagsTagIDOK() *DeleteTagsTagIDOK {
	return &DeleteTagsTagIDOK{}
}

/*
DeleteTagsTagIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteTagsTagIDOK struct {
	Payload *models.DeleteTagResponse
}

// IsSuccess returns true when this delete tags tag Id o k response has a 2xx status code
func (o *DeleteTagsTagIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete tags tag Id o k response has a 3xx status code
func (o *DeleteTagsTagIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tags tag Id o k response has a 4xx status code
func (o *DeleteTagsTagIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete tags tag Id o k response has a 5xx status code
func (o *DeleteTagsTagIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete tags tag Id o k response a status code equal to that given
func (o *DeleteTagsTagIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete tags tag Id o k response
func (o *DeleteTagsTagIDOK) Code() int {
	return 200
}

func (o *DeleteTagsTagIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /tags/{tagId}][%d] deleteTagsTagIdOK %s", 200, payload)
}

func (o *DeleteTagsTagIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /tags/{tagId}][%d] deleteTagsTagIdOK %s", 200, payload)
}

func (o *DeleteTagsTagIDOK) GetPayload() *models.DeleteTagResponse {
	return o.Payload
}

func (o *DeleteTagsTagIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteTagResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTagsTagIDBadRequest creates a DeleteTagsTagIDBadRequest with default headers values
func NewDeleteTagsTagIDBadRequest() *DeleteTagsTagIDBadRequest {
	return &DeleteTagsTagIDBadRequest{}
}

/*
DeleteTagsTagIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteTagsTagIDBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this delete tags tag Id bad request response has a 2xx status code
func (o *DeleteTagsTagIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete tags tag Id bad request response has a 3xx status code
func (o *DeleteTagsTagIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tags tag Id bad request response has a 4xx status code
func (o *DeleteTagsTagIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete tags tag Id bad request response has a 5xx status code
func (o *DeleteTagsTagIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete tags tag Id bad request response a status code equal to that given
func (o *DeleteTagsTagIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete tags tag Id bad request response
func (o *DeleteTagsTagIDBadRequest) Code() int {
	return 400
}

func (o *DeleteTagsTagIDBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /tags/{tagId}][%d] deleteTagsTagIdBadRequest %s", 400, payload)
}

func (o *DeleteTagsTagIDBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /tags/{tagId}][%d] deleteTagsTagIdBadRequest %s", 400, payload)
}

func (o *DeleteTagsTagIDBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *DeleteTagsTagIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTagsTagIDUnauthorized creates a DeleteTagsTagIDUnauthorized with default headers values
func NewDeleteTagsTagIDUnauthorized() *DeleteTagsTagIDUnauthorized {
	return &DeleteTagsTagIDUnauthorized{}
}

/*
DeleteTagsTagIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteTagsTagIDUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this delete tags tag Id unauthorized response has a 2xx status code
func (o *DeleteTagsTagIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete tags tag Id unauthorized response has a 3xx status code
func (o *DeleteTagsTagIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tags tag Id unauthorized response has a 4xx status code
func (o *DeleteTagsTagIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete tags tag Id unauthorized response has a 5xx status code
func (o *DeleteTagsTagIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete tags tag Id unauthorized response a status code equal to that given
func (o *DeleteTagsTagIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete tags tag Id unauthorized response
func (o *DeleteTagsTagIDUnauthorized) Code() int {
	return 401
}

func (o *DeleteTagsTagIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /tags/{tagId}][%d] deleteTagsTagIdUnauthorized %s", 401, payload)
}

func (o *DeleteTagsTagIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /tags/{tagId}][%d] deleteTagsTagIdUnauthorized %s", 401, payload)
}

func (o *DeleteTagsTagIDUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *DeleteTagsTagIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTagsTagIDForbidden creates a DeleteTagsTagIDForbidden with default headers values
func NewDeleteTagsTagIDForbidden() *DeleteTagsTagIDForbidden {
	return &DeleteTagsTagIDForbidden{}
}

/*
DeleteTagsTagIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteTagsTagIDForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this delete tags tag Id forbidden response has a 2xx status code
func (o *DeleteTagsTagIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete tags tag Id forbidden response has a 3xx status code
func (o *DeleteTagsTagIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tags tag Id forbidden response has a 4xx status code
func (o *DeleteTagsTagIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete tags tag Id forbidden response has a 5xx status code
func (o *DeleteTagsTagIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete tags tag Id forbidden response a status code equal to that given
func (o *DeleteTagsTagIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete tags tag Id forbidden response
func (o *DeleteTagsTagIDForbidden) Code() int {
	return 403
}

func (o *DeleteTagsTagIDForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /tags/{tagId}][%d] deleteTagsTagIdForbidden %s", 403, payload)
}

func (o *DeleteTagsTagIDForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /tags/{tagId}][%d] deleteTagsTagIdForbidden %s", 403, payload)
}

func (o *DeleteTagsTagIDForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *DeleteTagsTagIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTagsTagIDNotFound creates a DeleteTagsTagIDNotFound with default headers values
func NewDeleteTagsTagIDNotFound() *DeleteTagsTagIDNotFound {
	return &DeleteTagsTagIDNotFound{}
}

/*
DeleteTagsTagIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteTagsTagIDNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this delete tags tag Id not found response has a 2xx status code
func (o *DeleteTagsTagIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete tags tag Id not found response has a 3xx status code
func (o *DeleteTagsTagIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tags tag Id not found response has a 4xx status code
func (o *DeleteTagsTagIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete tags tag Id not found response has a 5xx status code
func (o *DeleteTagsTagIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete tags tag Id not found response a status code equal to that given
func (o *DeleteTagsTagIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete tags tag Id not found response
func (o *DeleteTagsTagIDNotFound) Code() int {
	return 404
}

func (o *DeleteTagsTagIDNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /tags/{tagId}][%d] deleteTagsTagIdNotFound %s", 404, payload)
}

func (o *DeleteTagsTagIDNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /tags/{tagId}][%d] deleteTagsTagIdNotFound %s", 404, payload)
}

func (o *DeleteTagsTagIDNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *DeleteTagsTagIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
