// Code generated by go-swagger; DO NOT EDIT.

package logical_ports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// LogicalPortGetUsingGETReader is a Reader for the LogicalPortGetUsingGET structure.
type LogicalPortGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogicalPortGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogicalPortGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLogicalPortGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLogicalPortGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLogicalPortGetUsingGETOK creates a LogicalPortGetUsingGETOK with default headers values
func NewLogicalPortGetUsingGETOK() *LogicalPortGetUsingGETOK {
	return &LogicalPortGetUsingGETOK{}
}

/*LogicalPortGetUsingGETOK handles this case with default header values.

OK
*/
type LogicalPortGetUsingGETOK struct {
	Payload *models.LogicalPortGetResponse
}

func (o *LogicalPortGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /network/logical-ports/{logicalPortId}][%d] logicalPortGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *LogicalPortGetUsingGETOK) GetPayload() *models.LogicalPortGetResponse {
	return o.Payload
}

func (o *LogicalPortGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogicalPortGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogicalPortGetUsingGETBadRequest creates a LogicalPortGetUsingGETBadRequest with default headers values
func NewLogicalPortGetUsingGETBadRequest() *LogicalPortGetUsingGETBadRequest {
	return &LogicalPortGetUsingGETBadRequest{}
}

/*LogicalPortGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type LogicalPortGetUsingGETBadRequest struct {
}

func (o *LogicalPortGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /network/logical-ports/{logicalPortId}][%d] logicalPortGetUsingGETBadRequest ", 400)
}

func (o *LogicalPortGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLogicalPortGetUsingGETNotFound creates a LogicalPortGetUsingGETNotFound with default headers values
func NewLogicalPortGetUsingGETNotFound() *LogicalPortGetUsingGETNotFound {
	return &LogicalPortGetUsingGETNotFound{}
}

/*LogicalPortGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type LogicalPortGetUsingGETNotFound struct {
}

func (o *LogicalPortGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /network/logical-ports/{logicalPortId}][%d] logicalPortGetUsingGETNotFound ", 404)
}

func (o *LogicalPortGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
