// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// DeploymentGetUsingGETReader is a Reader for the DeploymentGetUsingGET structure.
type DeploymentGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeploymentGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeploymentGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeploymentGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeploymentGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeploymentGetUsingGETOK creates a DeploymentGetUsingGETOK with default headers values
func NewDeploymentGetUsingGETOK() *DeploymentGetUsingGETOK {
	return &DeploymentGetUsingGETOK{}
}

/*DeploymentGetUsingGETOK handles this case with default header values.

OK
*/
type DeploymentGetUsingGETOK struct {
	Payload *models.DeploymentGetResponse
}

func (o *DeploymentGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /deployments/{deploymentId}][%d] deploymentGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *DeploymentGetUsingGETOK) GetPayload() *models.DeploymentGetResponse {
	return o.Payload
}

func (o *DeploymentGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeploymentGetUsingGETBadRequest creates a DeploymentGetUsingGETBadRequest with default headers values
func NewDeploymentGetUsingGETBadRequest() *DeploymentGetUsingGETBadRequest {
	return &DeploymentGetUsingGETBadRequest{}
}

/*DeploymentGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type DeploymentGetUsingGETBadRequest struct {
}

func (o *DeploymentGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /deployments/{deploymentId}][%d] deploymentGetUsingGETBadRequest ", 400)
}

func (o *DeploymentGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeploymentGetUsingGETNotFound creates a DeploymentGetUsingGETNotFound with default headers values
func NewDeploymentGetUsingGETNotFound() *DeploymentGetUsingGETNotFound {
	return &DeploymentGetUsingGETNotFound{}
}

/*DeploymentGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type DeploymentGetUsingGETNotFound struct {
}

func (o *DeploymentGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /deployments/{deploymentId}][%d] deploymentGetUsingGETNotFound ", 404)
}

func (o *DeploymentGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}