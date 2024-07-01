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

// DeploymentListUsingGETReader is a Reader for the DeploymentListUsingGET structure.
type DeploymentListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeploymentListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeploymentListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeploymentListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /deployments] deploymentListUsingGET", response, response.Code())
	}
}

// NewDeploymentListUsingGETOK creates a DeploymentListUsingGETOK with default headers values
func NewDeploymentListUsingGETOK() *DeploymentListUsingGETOK {
	return &DeploymentListUsingGETOK{}
}

/*
DeploymentListUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type DeploymentListUsingGETOK struct {
	Payload *models.DeploymentListResponse
}

// IsSuccess returns true when this deployment list using g e t o k response has a 2xx status code
func (o *DeploymentListUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this deployment list using g e t o k response has a 3xx status code
func (o *DeploymentListUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this deployment list using g e t o k response has a 4xx status code
func (o *DeploymentListUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this deployment list using g e t o k response has a 5xx status code
func (o *DeploymentListUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this deployment list using g e t o k response a status code equal to that given
func (o *DeploymentListUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the deployment list using g e t o k response
func (o *DeploymentListUsingGETOK) Code() int {
	return 200
}

func (o *DeploymentListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /deployments][%d] deploymentListUsingGETOK  %+v", 200, o.Payload)
}

func (o *DeploymentListUsingGETOK) String() string {
	return fmt.Sprintf("[GET /deployments][%d] deploymentListUsingGETOK  %+v", 200, o.Payload)
}

func (o *DeploymentListUsingGETOK) GetPayload() *models.DeploymentListResponse {
	return o.Payload
}

func (o *DeploymentListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeploymentListUsingGETBadRequest creates a DeploymentListUsingGETBadRequest with default headers values
func NewDeploymentListUsingGETBadRequest() *DeploymentListUsingGETBadRequest {
	return &DeploymentListUsingGETBadRequest{}
}

/*
DeploymentListUsingGETBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type DeploymentListUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

// IsSuccess returns true when this deployment list using g e t bad request response has a 2xx status code
func (o *DeploymentListUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this deployment list using g e t bad request response has a 3xx status code
func (o *DeploymentListUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this deployment list using g e t bad request response has a 4xx status code
func (o *DeploymentListUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this deployment list using g e t bad request response has a 5xx status code
func (o *DeploymentListUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this deployment list using g e t bad request response a status code equal to that given
func (o *DeploymentListUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the deployment list using g e t bad request response
func (o *DeploymentListUsingGETBadRequest) Code() int {
	return 400
}

func (o *DeploymentListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /deployments][%d] deploymentListUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *DeploymentListUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /deployments][%d] deploymentListUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *DeploymentListUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *DeploymentListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
