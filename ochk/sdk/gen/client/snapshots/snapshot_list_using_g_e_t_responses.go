// Code generated by go-swagger; DO NOT EDIT.

package snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// SnapshotListUsingGETReader is a Reader for the SnapshotListUsingGET structure.
type SnapshotListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapshotListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSnapshotListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /vcs/snapshots] snapshotListUsingGET", response, response.Code())
	}
}

// NewSnapshotListUsingGETOK creates a SnapshotListUsingGETOK with default headers values
func NewSnapshotListUsingGETOK() *SnapshotListUsingGETOK {
	return &SnapshotListUsingGETOK{}
}

/*
SnapshotListUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type SnapshotListUsingGETOK struct {
	Payload *models.SnapshotListResponse
}

// IsSuccess returns true when this snapshot list using g e t o k response has a 2xx status code
func (o *SnapshotListUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapshot list using g e t o k response has a 3xx status code
func (o *SnapshotListUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapshot list using g e t o k response has a 4xx status code
func (o *SnapshotListUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapshot list using g e t o k response has a 5xx status code
func (o *SnapshotListUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snapshot list using g e t o k response a status code equal to that given
func (o *SnapshotListUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snapshot list using g e t o k response
func (o *SnapshotListUsingGETOK) Code() int {
	return 200
}

func (o *SnapshotListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /vcs/snapshots][%d] snapshotListUsingGETOK  %+v", 200, o.Payload)
}

func (o *SnapshotListUsingGETOK) String() string {
	return fmt.Sprintf("[GET /vcs/snapshots][%d] snapshotListUsingGETOK  %+v", 200, o.Payload)
}

func (o *SnapshotListUsingGETOK) GetPayload() *models.SnapshotListResponse {
	return o.Payload
}

func (o *SnapshotListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotListUsingGETBadRequest creates a SnapshotListUsingGETBadRequest with default headers values
func NewSnapshotListUsingGETBadRequest() *SnapshotListUsingGETBadRequest {
	return &SnapshotListUsingGETBadRequest{}
}

/*
SnapshotListUsingGETBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type SnapshotListUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this snapshot list using g e t bad request response has a 2xx status code
func (o *SnapshotListUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this snapshot list using g e t bad request response has a 3xx status code
func (o *SnapshotListUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapshot list using g e t bad request response has a 4xx status code
func (o *SnapshotListUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this snapshot list using g e t bad request response has a 5xx status code
func (o *SnapshotListUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this snapshot list using g e t bad request response a status code equal to that given
func (o *SnapshotListUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the snapshot list using g e t bad request response
func (o *SnapshotListUsingGETBadRequest) Code() int {
	return 400
}

func (o *SnapshotListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /vcs/snapshots][%d] snapshotListUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *SnapshotListUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /vcs/snapshots][%d] snapshotListUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *SnapshotListUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *SnapshotListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
