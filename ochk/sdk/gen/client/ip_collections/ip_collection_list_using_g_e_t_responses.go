// Code generated by go-swagger; DO NOT EDIT.

package ip_collections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// IPCollectionListUsingGETReader is a Reader for the IPCollectionListUsingGET structure.
type IPCollectionListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPCollectionListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIPCollectionListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIPCollectionListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIPCollectionListUsingGETOK creates a IPCollectionListUsingGETOK with default headers values
func NewIPCollectionListUsingGETOK() *IPCollectionListUsingGETOK {
	return &IPCollectionListUsingGETOK{}
}

/* IPCollectionListUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type IPCollectionListUsingGETOK struct {
	Payload *models.IPCollectionListResponse
}

func (o *IPCollectionListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /ipcs][%d] ipCollectionListUsingGETOK  %+v", 200, o.Payload)
}
func (o *IPCollectionListUsingGETOK) GetPayload() *models.IPCollectionListResponse {
	return o.Payload
}

func (o *IPCollectionListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPCollectionListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIPCollectionListUsingGETBadRequest creates a IPCollectionListUsingGETBadRequest with default headers values
func NewIPCollectionListUsingGETBadRequest() *IPCollectionListUsingGETBadRequest {
	return &IPCollectionListUsingGETBadRequest{}
}

/* IPCollectionListUsingGETBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type IPCollectionListUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *IPCollectionListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /ipcs][%d] ipCollectionListUsingGETBadRequest  %+v", 400, o.Payload)
}
func (o *IPCollectionListUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *IPCollectionListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
