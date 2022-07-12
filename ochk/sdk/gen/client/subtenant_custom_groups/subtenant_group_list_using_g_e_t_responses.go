// Code generated by go-swagger; DO NOT EDIT.

package subtenant_custom_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// SubtenantGroupListUsingGETReader is a Reader for the SubtenantGroupListUsingGET structure.
type SubtenantGroupListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubtenantGroupListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubtenantGroupListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubtenantGroupListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubtenantGroupListUsingGETOK creates a SubtenantGroupListUsingGETOK with default headers values
func NewSubtenantGroupListUsingGETOK() *SubtenantGroupListUsingGETOK {
	return &SubtenantGroupListUsingGETOK{}
}

/* SubtenantGroupListUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type SubtenantGroupListUsingGETOK struct {
	Payload *models.GroupListResponse
}

func (o *SubtenantGroupListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /subtenants/{subtenantId}/groups][%d] subtenantGroupListUsingGETOK  %+v", 200, o.Payload)
}
func (o *SubtenantGroupListUsingGETOK) GetPayload() *models.GroupListResponse {
	return o.Payload
}

func (o *SubtenantGroupListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubtenantGroupListUsingGETBadRequest creates a SubtenantGroupListUsingGETBadRequest with default headers values
func NewSubtenantGroupListUsingGETBadRequest() *SubtenantGroupListUsingGETBadRequest {
	return &SubtenantGroupListUsingGETBadRequest{}
}

/* SubtenantGroupListUsingGETBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type SubtenantGroupListUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *SubtenantGroupListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /subtenants/{subtenantId}/groups][%d] subtenantGroupListUsingGETBadRequest  %+v", 400, o.Payload)
}
func (o *SubtenantGroupListUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *SubtenantGroupListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
