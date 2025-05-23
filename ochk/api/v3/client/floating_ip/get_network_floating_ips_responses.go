// Code generated by go-swagger; DO NOT EDIT.

package floating_ip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
)

// GetNetworkFloatingIpsReader is a Reader for the GetNetworkFloatingIps structure.
type GetNetworkFloatingIpsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkFloatingIpsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkFloatingIpsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNetworkFloatingIpsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetNetworkFloatingIpsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNetworkFloatingIpsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /network/floating-ips] GetNetworkFloatingIps", response, response.Code())
	}
}

// NewGetNetworkFloatingIpsOK creates a GetNetworkFloatingIpsOK with default headers values
func NewGetNetworkFloatingIpsOK() *GetNetworkFloatingIpsOK {
	return &GetNetworkFloatingIpsOK{}
}

/*
GetNetworkFloatingIpsOK describes a response with status code 200, with default header values.

OK
*/
type GetNetworkFloatingIpsOK struct {
	Payload *models.ListFloatingIPResponse
}

// IsSuccess returns true when this get network floating ips o k response has a 2xx status code
func (o *GetNetworkFloatingIpsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network floating ips o k response has a 3xx status code
func (o *GetNetworkFloatingIpsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network floating ips o k response has a 4xx status code
func (o *GetNetworkFloatingIpsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network floating ips o k response has a 5xx status code
func (o *GetNetworkFloatingIpsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network floating ips o k response a status code equal to that given
func (o *GetNetworkFloatingIpsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network floating ips o k response
func (o *GetNetworkFloatingIpsOK) Code() int {
	return 200
}

func (o *GetNetworkFloatingIpsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/floating-ips][%d] getNetworkFloatingIpsOK %s", 200, payload)
}

func (o *GetNetworkFloatingIpsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/floating-ips][%d] getNetworkFloatingIpsOK %s", 200, payload)
}

func (o *GetNetworkFloatingIpsOK) GetPayload() *models.ListFloatingIPResponse {
	return o.Payload
}

func (o *GetNetworkFloatingIpsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListFloatingIPResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkFloatingIpsBadRequest creates a GetNetworkFloatingIpsBadRequest with default headers values
func NewGetNetworkFloatingIpsBadRequest() *GetNetworkFloatingIpsBadRequest {
	return &GetNetworkFloatingIpsBadRequest{}
}

/*
GetNetworkFloatingIpsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetNetworkFloatingIpsBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get network floating ips bad request response has a 2xx status code
func (o *GetNetworkFloatingIpsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network floating ips bad request response has a 3xx status code
func (o *GetNetworkFloatingIpsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network floating ips bad request response has a 4xx status code
func (o *GetNetworkFloatingIpsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get network floating ips bad request response has a 5xx status code
func (o *GetNetworkFloatingIpsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get network floating ips bad request response a status code equal to that given
func (o *GetNetworkFloatingIpsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get network floating ips bad request response
func (o *GetNetworkFloatingIpsBadRequest) Code() int {
	return 400
}

func (o *GetNetworkFloatingIpsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/floating-ips][%d] getNetworkFloatingIpsBadRequest %s", 400, payload)
}

func (o *GetNetworkFloatingIpsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/floating-ips][%d] getNetworkFloatingIpsBadRequest %s", 400, payload)
}

func (o *GetNetworkFloatingIpsBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetNetworkFloatingIpsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkFloatingIpsUnauthorized creates a GetNetworkFloatingIpsUnauthorized with default headers values
func NewGetNetworkFloatingIpsUnauthorized() *GetNetworkFloatingIpsUnauthorized {
	return &GetNetworkFloatingIpsUnauthorized{}
}

/*
GetNetworkFloatingIpsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetNetworkFloatingIpsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get network floating ips unauthorized response has a 2xx status code
func (o *GetNetworkFloatingIpsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network floating ips unauthorized response has a 3xx status code
func (o *GetNetworkFloatingIpsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network floating ips unauthorized response has a 4xx status code
func (o *GetNetworkFloatingIpsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get network floating ips unauthorized response has a 5xx status code
func (o *GetNetworkFloatingIpsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get network floating ips unauthorized response a status code equal to that given
func (o *GetNetworkFloatingIpsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get network floating ips unauthorized response
func (o *GetNetworkFloatingIpsUnauthorized) Code() int {
	return 401
}

func (o *GetNetworkFloatingIpsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/floating-ips][%d] getNetworkFloatingIpsUnauthorized %s", 401, payload)
}

func (o *GetNetworkFloatingIpsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/floating-ips][%d] getNetworkFloatingIpsUnauthorized %s", 401, payload)
}

func (o *GetNetworkFloatingIpsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetNetworkFloatingIpsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkFloatingIpsForbidden creates a GetNetworkFloatingIpsForbidden with default headers values
func NewGetNetworkFloatingIpsForbidden() *GetNetworkFloatingIpsForbidden {
	return &GetNetworkFloatingIpsForbidden{}
}

/*
GetNetworkFloatingIpsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetNetworkFloatingIpsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get network floating ips forbidden response has a 2xx status code
func (o *GetNetworkFloatingIpsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network floating ips forbidden response has a 3xx status code
func (o *GetNetworkFloatingIpsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network floating ips forbidden response has a 4xx status code
func (o *GetNetworkFloatingIpsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get network floating ips forbidden response has a 5xx status code
func (o *GetNetworkFloatingIpsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get network floating ips forbidden response a status code equal to that given
func (o *GetNetworkFloatingIpsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get network floating ips forbidden response
func (o *GetNetworkFloatingIpsForbidden) Code() int {
	return 403
}

func (o *GetNetworkFloatingIpsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/floating-ips][%d] getNetworkFloatingIpsForbidden %s", 403, payload)
}

func (o *GetNetworkFloatingIpsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/floating-ips][%d] getNetworkFloatingIpsForbidden %s", 403, payload)
}

func (o *GetNetworkFloatingIpsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetNetworkFloatingIpsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
