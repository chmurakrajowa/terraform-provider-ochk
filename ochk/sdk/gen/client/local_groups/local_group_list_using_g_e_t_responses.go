// Code generated by go-swagger; DO NOT EDIT.

package local_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// LocalGroupListUsingGETReader is a Reader for the LocalGroupListUsingGET structure.
type LocalGroupListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocalGroupListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocalGroupListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLocalGroupListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLocalGroupListUsingGETOK creates a LocalGroupListUsingGETOK with default headers values
func NewLocalGroupListUsingGETOK() *LocalGroupListUsingGETOK {
	return &LocalGroupListUsingGETOK{}
}

/*LocalGroupListUsingGETOK handles this case with default header values.

OK
*/
type LocalGroupListUsingGETOK struct {
	Payload *models.LocalGroupListResponse
}

func (o *LocalGroupListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /groups/local][%d] localGroupListUsingGETOK  %+v", 200, o.Payload)
}

func (o *LocalGroupListUsingGETOK) GetPayload() *models.LocalGroupListResponse {
	return o.Payload
}

func (o *LocalGroupListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocalGroupListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocalGroupListUsingGETBadRequest creates a LocalGroupListUsingGETBadRequest with default headers values
func NewLocalGroupListUsingGETBadRequest() *LocalGroupListUsingGETBadRequest {
	return &LocalGroupListUsingGETBadRequest{}
}

/*LocalGroupListUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type LocalGroupListUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *LocalGroupListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /groups/local][%d] localGroupListUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *LocalGroupListUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *LocalGroupListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
