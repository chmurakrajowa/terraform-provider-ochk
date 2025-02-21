// Code generated by go-swagger; DO NOT EDIT.

package port_forwarding

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

// PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDReader is a Reader for the PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingID structure.
type PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /network/floating-ips/{floatingIpId}/port-forwardings/{portForwardingId}] PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingID", response, response.Code())
	}
}

// NewPutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDOK creates a PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDOK with default headers values
func NewPutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDOK() *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDOK {
	return &PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDOK{}
}

/*
PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDOK describes a response with status code 200, with default header values.

OK
*/
type PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDOK struct {
	Payload *models.UpdatePortForwardingResponse
}

// IsSuccess returns true when this put network floating ips floating Ip Id port forwardings port forwarding Id o k response has a 2xx status code
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put network floating ips floating Ip Id port forwardings port forwarding Id o k response has a 3xx status code
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put network floating ips floating Ip Id port forwardings port forwarding Id o k response has a 4xx status code
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put network floating ips floating Ip Id port forwardings port forwarding Id o k response has a 5xx status code
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put network floating ips floating Ip Id port forwardings port forwarding Id o k response a status code equal to that given
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put network floating ips floating Ip Id port forwardings port forwarding Id o k response
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDOK) Code() int {
	return 200
}

func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/floating-ips/{floatingIpId}/port-forwardings/{portForwardingId}][%d] putNetworkFloatingIpsFloatingIpIdPortForwardingsPortForwardingIdOK %s", 200, payload)
}

func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/floating-ips/{floatingIpId}/port-forwardings/{portForwardingId}][%d] putNetworkFloatingIpsFloatingIpIdPortForwardingsPortForwardingIdOK %s", 200, payload)
}

func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDOK) GetPayload() *models.UpdatePortForwardingResponse {
	return o.Payload
}

func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdatePortForwardingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDBadRequest creates a PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDBadRequest with default headers values
func NewPutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDBadRequest() *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDBadRequest {
	return &PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDBadRequest{}
}

/*
PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put network floating ips floating Ip Id port forwardings port forwarding Id bad request response has a 2xx status code
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put network floating ips floating Ip Id port forwardings port forwarding Id bad request response has a 3xx status code
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put network floating ips floating Ip Id port forwardings port forwarding Id bad request response has a 4xx status code
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put network floating ips floating Ip Id port forwardings port forwarding Id bad request response has a 5xx status code
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put network floating ips floating Ip Id port forwardings port forwarding Id bad request response a status code equal to that given
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put network floating ips floating Ip Id port forwardings port forwarding Id bad request response
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDBadRequest) Code() int {
	return 400
}

func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/floating-ips/{floatingIpId}/port-forwardings/{portForwardingId}][%d] putNetworkFloatingIpsFloatingIpIdPortForwardingsPortForwardingIdBadRequest %s", 400, payload)
}

func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/floating-ips/{floatingIpId}/port-forwardings/{portForwardingId}][%d] putNetworkFloatingIpsFloatingIpIdPortForwardingsPortForwardingIdBadRequest %s", 400, payload)
}

func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDUnauthorized creates a PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDUnauthorized with default headers values
func NewPutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDUnauthorized() *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDUnauthorized {
	return &PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDUnauthorized{}
}

/*
PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put network floating ips floating Ip Id port forwardings port forwarding Id unauthorized response has a 2xx status code
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put network floating ips floating Ip Id port forwardings port forwarding Id unauthorized response has a 3xx status code
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put network floating ips floating Ip Id port forwardings port forwarding Id unauthorized response has a 4xx status code
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put network floating ips floating Ip Id port forwardings port forwarding Id unauthorized response has a 5xx status code
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put network floating ips floating Ip Id port forwardings port forwarding Id unauthorized response a status code equal to that given
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put network floating ips floating Ip Id port forwardings port forwarding Id unauthorized response
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDUnauthorized) Code() int {
	return 401
}

func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/floating-ips/{floatingIpId}/port-forwardings/{portForwardingId}][%d] putNetworkFloatingIpsFloatingIpIdPortForwardingsPortForwardingIdUnauthorized %s", 401, payload)
}

func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/floating-ips/{floatingIpId}/port-forwardings/{portForwardingId}][%d] putNetworkFloatingIpsFloatingIpIdPortForwardingsPortForwardingIdUnauthorized %s", 401, payload)
}

func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDForbidden creates a PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDForbidden with default headers values
func NewPutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDForbidden() *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDForbidden {
	return &PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDForbidden{}
}

/*
PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this put network floating ips floating Ip Id port forwardings port forwarding Id forbidden response has a 2xx status code
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put network floating ips floating Ip Id port forwardings port forwarding Id forbidden response has a 3xx status code
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put network floating ips floating Ip Id port forwardings port forwarding Id forbidden response has a 4xx status code
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put network floating ips floating Ip Id port forwardings port forwarding Id forbidden response has a 5xx status code
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put network floating ips floating Ip Id port forwardings port forwarding Id forbidden response a status code equal to that given
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put network floating ips floating Ip Id port forwardings port forwarding Id forbidden response
func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDForbidden) Code() int {
	return 403
}

func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/floating-ips/{floatingIpId}/port-forwardings/{portForwardingId}][%d] putNetworkFloatingIpsFloatingIpIdPortForwardingsPortForwardingIdForbidden %s", 403, payload)
}

func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/floating-ips/{floatingIpId}/port-forwardings/{portForwardingId}][%d] putNetworkFloatingIpsFloatingIpIdPortForwardingsPortForwardingIdForbidden %s", 403, payload)
}

func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
