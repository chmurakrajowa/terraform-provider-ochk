// Code generated by go-swagger; DO NOT EDIT.

package key

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

// PostKmsKeyIDExportReader is a Reader for the PostKmsKeyIDExport structure.
type PostKmsKeyIDExportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostKmsKeyIDExportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostKmsKeyIDExportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostKmsKeyIDExportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostKmsKeyIDExportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostKmsKeyIDExportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostKmsKeyIDExportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /kms/key/{id}/export] PostKmsKeyIDExport", response, response.Code())
	}
}

// NewPostKmsKeyIDExportOK creates a PostKmsKeyIDExportOK with default headers values
func NewPostKmsKeyIDExportOK() *PostKmsKeyIDExportOK {
	return &PostKmsKeyIDExportOK{}
}

/*
PostKmsKeyIDExportOK describes a response with status code 200, with default header values.

OK
*/
type PostKmsKeyIDExportOK struct {
	Payload *models.ExportKmsKeyResponse
}

// IsSuccess returns true when this post kms key Id export o k response has a 2xx status code
func (o *PostKmsKeyIDExportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post kms key Id export o k response has a 3xx status code
func (o *PostKmsKeyIDExportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post kms key Id export o k response has a 4xx status code
func (o *PostKmsKeyIDExportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post kms key Id export o k response has a 5xx status code
func (o *PostKmsKeyIDExportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post kms key Id export o k response a status code equal to that given
func (o *PostKmsKeyIDExportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post kms key Id export o k response
func (o *PostKmsKeyIDExportOK) Code() int {
	return 200
}

func (o *PostKmsKeyIDExportOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kms/key/{id}/export][%d] postKmsKeyIdExportOK %s", 200, payload)
}

func (o *PostKmsKeyIDExportOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kms/key/{id}/export][%d] postKmsKeyIdExportOK %s", 200, payload)
}

func (o *PostKmsKeyIDExportOK) GetPayload() *models.ExportKmsKeyResponse {
	return o.Payload
}

func (o *PostKmsKeyIDExportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportKmsKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKmsKeyIDExportBadRequest creates a PostKmsKeyIDExportBadRequest with default headers values
func NewPostKmsKeyIDExportBadRequest() *PostKmsKeyIDExportBadRequest {
	return &PostKmsKeyIDExportBadRequest{}
}

/*
PostKmsKeyIDExportBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostKmsKeyIDExportBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this post kms key Id export bad request response has a 2xx status code
func (o *PostKmsKeyIDExportBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post kms key Id export bad request response has a 3xx status code
func (o *PostKmsKeyIDExportBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post kms key Id export bad request response has a 4xx status code
func (o *PostKmsKeyIDExportBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post kms key Id export bad request response has a 5xx status code
func (o *PostKmsKeyIDExportBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post kms key Id export bad request response a status code equal to that given
func (o *PostKmsKeyIDExportBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post kms key Id export bad request response
func (o *PostKmsKeyIDExportBadRequest) Code() int {
	return 400
}

func (o *PostKmsKeyIDExportBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kms/key/{id}/export][%d] postKmsKeyIdExportBadRequest %s", 400, payload)
}

func (o *PostKmsKeyIDExportBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kms/key/{id}/export][%d] postKmsKeyIdExportBadRequest %s", 400, payload)
}

func (o *PostKmsKeyIDExportBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PostKmsKeyIDExportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKmsKeyIDExportUnauthorized creates a PostKmsKeyIDExportUnauthorized with default headers values
func NewPostKmsKeyIDExportUnauthorized() *PostKmsKeyIDExportUnauthorized {
	return &PostKmsKeyIDExportUnauthorized{}
}

/*
PostKmsKeyIDExportUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostKmsKeyIDExportUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this post kms key Id export unauthorized response has a 2xx status code
func (o *PostKmsKeyIDExportUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post kms key Id export unauthorized response has a 3xx status code
func (o *PostKmsKeyIDExportUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post kms key Id export unauthorized response has a 4xx status code
func (o *PostKmsKeyIDExportUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post kms key Id export unauthorized response has a 5xx status code
func (o *PostKmsKeyIDExportUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post kms key Id export unauthorized response a status code equal to that given
func (o *PostKmsKeyIDExportUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post kms key Id export unauthorized response
func (o *PostKmsKeyIDExportUnauthorized) Code() int {
	return 401
}

func (o *PostKmsKeyIDExportUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kms/key/{id}/export][%d] postKmsKeyIdExportUnauthorized %s", 401, payload)
}

func (o *PostKmsKeyIDExportUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kms/key/{id}/export][%d] postKmsKeyIdExportUnauthorized %s", 401, payload)
}

func (o *PostKmsKeyIDExportUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PostKmsKeyIDExportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKmsKeyIDExportForbidden creates a PostKmsKeyIDExportForbidden with default headers values
func NewPostKmsKeyIDExportForbidden() *PostKmsKeyIDExportForbidden {
	return &PostKmsKeyIDExportForbidden{}
}

/*
PostKmsKeyIDExportForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostKmsKeyIDExportForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this post kms key Id export forbidden response has a 2xx status code
func (o *PostKmsKeyIDExportForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post kms key Id export forbidden response has a 3xx status code
func (o *PostKmsKeyIDExportForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post kms key Id export forbidden response has a 4xx status code
func (o *PostKmsKeyIDExportForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post kms key Id export forbidden response has a 5xx status code
func (o *PostKmsKeyIDExportForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post kms key Id export forbidden response a status code equal to that given
func (o *PostKmsKeyIDExportForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post kms key Id export forbidden response
func (o *PostKmsKeyIDExportForbidden) Code() int {
	return 403
}

func (o *PostKmsKeyIDExportForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kms/key/{id}/export][%d] postKmsKeyIdExportForbidden %s", 403, payload)
}

func (o *PostKmsKeyIDExportForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kms/key/{id}/export][%d] postKmsKeyIdExportForbidden %s", 403, payload)
}

func (o *PostKmsKeyIDExportForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PostKmsKeyIDExportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKmsKeyIDExportNotFound creates a PostKmsKeyIDExportNotFound with default headers values
func NewPostKmsKeyIDExportNotFound() *PostKmsKeyIDExportNotFound {
	return &PostKmsKeyIDExportNotFound{}
}

/*
PostKmsKeyIDExportNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostKmsKeyIDExportNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this post kms key Id export not found response has a 2xx status code
func (o *PostKmsKeyIDExportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post kms key Id export not found response has a 3xx status code
func (o *PostKmsKeyIDExportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post kms key Id export not found response has a 4xx status code
func (o *PostKmsKeyIDExportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post kms key Id export not found response has a 5xx status code
func (o *PostKmsKeyIDExportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post kms key Id export not found response a status code equal to that given
func (o *PostKmsKeyIDExportNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the post kms key Id export not found response
func (o *PostKmsKeyIDExportNotFound) Code() int {
	return 404
}

func (o *PostKmsKeyIDExportNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kms/key/{id}/export][%d] postKmsKeyIdExportNotFound %s", 404, payload)
}

func (o *PostKmsKeyIDExportNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kms/key/{id}/export][%d] postKmsKeyIdExportNotFound %s", 404, payload)
}

func (o *PostKmsKeyIDExportNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PostKmsKeyIDExportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
