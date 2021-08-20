// Code generated by go-swagger; DO NOT EDIT.

package system_tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// SystemTagListUsingGETReader is a Reader for the SystemTagListUsingGET structure.
type SystemTagListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemTagListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemTagListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSystemTagListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSystemTagListUsingGETOK creates a SystemTagListUsingGETOK with default headers values
func NewSystemTagListUsingGETOK() *SystemTagListUsingGETOK {
	return &SystemTagListUsingGETOK{}
}

/*SystemTagListUsingGETOK handles this case with default header values.

OK
*/
type SystemTagListUsingGETOK struct {
	Payload *models.SystemTagListResponse
}

func (o *SystemTagListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /tags/systemTags][%d] systemTagListUsingGETOK  %+v", 200, o.Payload)
}

func (o *SystemTagListUsingGETOK) GetPayload() *models.SystemTagListResponse {
	return o.Payload
}

func (o *SystemTagListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemTagListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSystemTagListUsingGETBadRequest creates a SystemTagListUsingGETBadRequest with default headers values
func NewSystemTagListUsingGETBadRequest() *SystemTagListUsingGETBadRequest {
	return &SystemTagListUsingGETBadRequest{}
}

/*SystemTagListUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type SystemTagListUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *SystemTagListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /tags/systemTags][%d] systemTagListUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *SystemTagListUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *SystemTagListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}