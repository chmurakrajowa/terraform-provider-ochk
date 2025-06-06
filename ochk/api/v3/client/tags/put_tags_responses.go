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

// PutTagsReader is a Reader for the PutTags structure.
type PutTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /tags] PutTags", response, response.Code())
	}
}

// NewPutTagsOK creates a PutTagsOK with default headers values
func NewPutTagsOK() *PutTagsOK {
	return &PutTagsOK{}
}

/*
PutTagsOK describes a response with status code 200, with default header values.

OK
*/
type PutTagsOK struct {
	Payload *models.CreateTagResponse
}

// IsSuccess returns true when this put tags o k response has a 2xx status code
func (o *PutTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put tags o k response has a 3xx status code
func (o *PutTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put tags o k response has a 4xx status code
func (o *PutTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put tags o k response has a 5xx status code
func (o *PutTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put tags o k response a status code equal to that given
func (o *PutTagsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put tags o k response
func (o *PutTagsOK) Code() int {
	return 200
}

func (o *PutTagsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tags][%d] putTagsOK %s", 200, payload)
}

func (o *PutTagsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tags][%d] putTagsOK %s", 200, payload)
}

func (o *PutTagsOK) GetPayload() *models.CreateTagResponse {
	return o.Payload
}

func (o *PutTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateTagResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTagsBadRequest creates a PutTagsBadRequest with default headers values
func NewPutTagsBadRequest() *PutTagsBadRequest {
	return &PutTagsBadRequest{}
}

/*
PutTagsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutTagsBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put tags bad request response has a 2xx status code
func (o *PutTagsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put tags bad request response has a 3xx status code
func (o *PutTagsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put tags bad request response has a 4xx status code
func (o *PutTagsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put tags bad request response has a 5xx status code
func (o *PutTagsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put tags bad request response a status code equal to that given
func (o *PutTagsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put tags bad request response
func (o *PutTagsBadRequest) Code() int {
	return 400
}

func (o *PutTagsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tags][%d] putTagsBadRequest %s", 400, payload)
}

func (o *PutTagsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tags][%d] putTagsBadRequest %s", 400, payload)
}

func (o *PutTagsBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTagsUnauthorized creates a PutTagsUnauthorized with default headers values
func NewPutTagsUnauthorized() *PutTagsUnauthorized {
	return &PutTagsUnauthorized{}
}

/*
PutTagsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutTagsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put tags unauthorized response has a 2xx status code
func (o *PutTagsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put tags unauthorized response has a 3xx status code
func (o *PutTagsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put tags unauthorized response has a 4xx status code
func (o *PutTagsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put tags unauthorized response has a 5xx status code
func (o *PutTagsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put tags unauthorized response a status code equal to that given
func (o *PutTagsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put tags unauthorized response
func (o *PutTagsUnauthorized) Code() int {
	return 401
}

func (o *PutTagsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tags][%d] putTagsUnauthorized %s", 401, payload)
}

func (o *PutTagsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tags][%d] putTagsUnauthorized %s", 401, payload)
}

func (o *PutTagsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTagsForbidden creates a PutTagsForbidden with default headers values
func NewPutTagsForbidden() *PutTagsForbidden {
	return &PutTagsForbidden{}
}

/*
PutTagsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutTagsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put tags forbidden response has a 2xx status code
func (o *PutTagsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put tags forbidden response has a 3xx status code
func (o *PutTagsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put tags forbidden response has a 4xx status code
func (o *PutTagsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put tags forbidden response has a 5xx status code
func (o *PutTagsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put tags forbidden response a status code equal to that given
func (o *PutTagsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put tags forbidden response
func (o *PutTagsForbidden) Code() int {
	return 403
}

func (o *PutTagsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tags][%d] putTagsForbidden %s", 403, payload)
}

func (o *PutTagsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tags][%d] putTagsForbidden %s", 403, payload)
}

func (o *PutTagsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
