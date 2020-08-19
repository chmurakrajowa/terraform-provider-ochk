// Code generated by go-swagger; DO NOT EDIT.

package subtenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// SubtenantGetUsingGETReader is a Reader for the SubtenantGetUsingGET structure.
type SubtenantGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubtenantGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubtenantGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubtenantGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSubtenantGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSubtenantGetUsingGETOK creates a SubtenantGetUsingGETOK with default headers values
func NewSubtenantGetUsingGETOK() *SubtenantGetUsingGETOK {
	return &SubtenantGetUsingGETOK{}
}

/*SubtenantGetUsingGETOK handles this case with default header values.

OK
*/
type SubtenantGetUsingGETOK struct {
	Payload *models.SubtenantGetResponse
}

func (o *SubtenantGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /subtenants/{subtenantId}][%d] subtenantGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *SubtenantGetUsingGETOK) GetPayload() *models.SubtenantGetResponse {
	return o.Payload
}

func (o *SubtenantGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubtenantGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubtenantGetUsingGETBadRequest creates a SubtenantGetUsingGETBadRequest with default headers values
func NewSubtenantGetUsingGETBadRequest() *SubtenantGetUsingGETBadRequest {
	return &SubtenantGetUsingGETBadRequest{}
}

/*SubtenantGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type SubtenantGetUsingGETBadRequest struct {
}

func (o *SubtenantGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /subtenants/{subtenantId}][%d] subtenantGetUsingGETBadRequest ", 400)
}

func (o *SubtenantGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubtenantGetUsingGETNotFound creates a SubtenantGetUsingGETNotFound with default headers values
func NewSubtenantGetUsingGETNotFound() *SubtenantGetUsingGETNotFound {
	return &SubtenantGetUsingGETNotFound{}
}

/*SubtenantGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type SubtenantGetUsingGETNotFound struct {
}

func (o *SubtenantGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /subtenants/{subtenantId}][%d] subtenantGetUsingGETNotFound ", 404)
}

func (o *SubtenantGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
