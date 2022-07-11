// Code generated by go-swagger; DO NOT EDIT.

package active_directory_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// GetADGroupUsingGETReader is a Reader for the GetADGroupUsingGET structure.
type GetADGroupUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetADGroupUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetADGroupUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetADGroupUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetADGroupUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetADGroupUsingGETOK creates a GetADGroupUsingGETOK with default headers values
func NewGetADGroupUsingGETOK() *GetADGroupUsingGETOK {
	return &GetADGroupUsingGETOK{}
}

/*GetADGroupUsingGETOK handles this case with default header values.

OK
*/
type GetADGroupUsingGETOK struct {
	Payload *models.GroupInstanceGetResponse
}

func (o *GetADGroupUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /ad/integration/groups/{samAccountName}][%d] getADGroupUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetADGroupUsingGETOK) GetPayload() *models.GroupInstanceGetResponse {
	return o.Payload
}

func (o *GetADGroupUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupInstanceGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetADGroupUsingGETBadRequest creates a GetADGroupUsingGETBadRequest with default headers values
func NewGetADGroupUsingGETBadRequest() *GetADGroupUsingGETBadRequest {
	return &GetADGroupUsingGETBadRequest{}
}

/*GetADGroupUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type GetADGroupUsingGETBadRequest struct {
}

func (o *GetADGroupUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /ad/integration/groups/{samAccountName}][%d] getADGroupUsingGETBadRequest ", 400)
}

func (o *GetADGroupUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetADGroupUsingGETNotFound creates a GetADGroupUsingGETNotFound with default headers values
func NewGetADGroupUsingGETNotFound() *GetADGroupUsingGETNotFound {
	return &GetADGroupUsingGETNotFound{}
}

/*GetADGroupUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type GetADGroupUsingGETNotFound struct {
}

func (o *GetADGroupUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /ad/integration/groups/{samAccountName}][%d] getADGroupUsingGETNotFound ", 404)
}

func (o *GetADGroupUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
