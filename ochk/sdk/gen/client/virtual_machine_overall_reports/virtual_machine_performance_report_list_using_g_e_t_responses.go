// Code generated by go-swagger; DO NOT EDIT.

package virtual_machine_overall_reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// VirtualMachinePerformanceReportListUsingGETReader is a Reader for the VirtualMachinePerformanceReportListUsingGET structure.
type VirtualMachinePerformanceReportListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualMachinePerformanceReportListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualMachinePerformanceReportListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVirtualMachinePerformanceReportListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewVirtualMachinePerformanceReportListUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualMachinePerformanceReportListUsingGETOK creates a VirtualMachinePerformanceReportListUsingGETOK with default headers values
func NewVirtualMachinePerformanceReportListUsingGETOK() *VirtualMachinePerformanceReportListUsingGETOK {
	return &VirtualMachinePerformanceReportListUsingGETOK{}
}

/* VirtualMachinePerformanceReportListUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type VirtualMachinePerformanceReportListUsingGETOK struct {
	Payload *models.VirtualMachinePerformanceReportGetResponse
}

func (o *VirtualMachinePerformanceReportListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /billing/virtual-machine/performance][%d] virtualMachinePerformanceReportListUsingGETOK  %+v", 200, o.Payload)
}
func (o *VirtualMachinePerformanceReportListUsingGETOK) GetPayload() *models.VirtualMachinePerformanceReportGetResponse {
	return o.Payload
}

func (o *VirtualMachinePerformanceReportListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachinePerformanceReportGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualMachinePerformanceReportListUsingGETBadRequest creates a VirtualMachinePerformanceReportListUsingGETBadRequest with default headers values
func NewVirtualMachinePerformanceReportListUsingGETBadRequest() *VirtualMachinePerformanceReportListUsingGETBadRequest {
	return &VirtualMachinePerformanceReportListUsingGETBadRequest{}
}

/* VirtualMachinePerformanceReportListUsingGETBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type VirtualMachinePerformanceReportListUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *VirtualMachinePerformanceReportListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /billing/virtual-machine/performance][%d] virtualMachinePerformanceReportListUsingGETBadRequest  %+v", 400, o.Payload)
}
func (o *VirtualMachinePerformanceReportListUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *VirtualMachinePerformanceReportListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualMachinePerformanceReportListUsingGETNotFound creates a VirtualMachinePerformanceReportListUsingGETNotFound with default headers values
func NewVirtualMachinePerformanceReportListUsingGETNotFound() *VirtualMachinePerformanceReportListUsingGETNotFound {
	return &VirtualMachinePerformanceReportListUsingGETNotFound{}
}

/* VirtualMachinePerformanceReportListUsingGETNotFound describes a response with status code 404, with default header values.

Entity not found.
*/
type VirtualMachinePerformanceReportListUsingGETNotFound struct {
}

func (o *VirtualMachinePerformanceReportListUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /billing/virtual-machine/performance][%d] virtualMachinePerformanceReportListUsingGETNotFound ", 404)
}

func (o *VirtualMachinePerformanceReportListUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
