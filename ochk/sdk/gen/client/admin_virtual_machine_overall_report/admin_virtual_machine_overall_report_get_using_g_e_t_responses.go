// Code generated by go-swagger; DO NOT EDIT.

package admin_virtual_machine_overall_report

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// AdminVirtualMachineOverallReportGetUsingGETReader is a Reader for the AdminVirtualMachineOverallReportGetUsingGET structure.
type AdminVirtualMachineOverallReportGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminVirtualMachineOverallReportGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminVirtualMachineOverallReportGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminVirtualMachineOverallReportGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAdminVirtualMachineOverallReportGetUsingGETOK creates a AdminVirtualMachineOverallReportGetUsingGETOK with default headers values
func NewAdminVirtualMachineOverallReportGetUsingGETOK() *AdminVirtualMachineOverallReportGetUsingGETOK {
	return &AdminVirtualMachineOverallReportGetUsingGETOK{}
}

/*AdminVirtualMachineOverallReportGetUsingGETOK handles this case with default header values.

OK
*/
type AdminVirtualMachineOverallReportGetUsingGETOK struct {
	Payload *models.VirtualMachineSummaryReportGetResponse
}

func (o *AdminVirtualMachineOverallReportGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /admin/virtual-machine/summary/{tenantId}][%d] adminVirtualMachineOverallReportGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *AdminVirtualMachineOverallReportGetUsingGETOK) GetPayload() *models.VirtualMachineSummaryReportGetResponse {
	return o.Payload
}

func (o *AdminVirtualMachineOverallReportGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineSummaryReportGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminVirtualMachineOverallReportGetUsingGETBadRequest creates a AdminVirtualMachineOverallReportGetUsingGETBadRequest with default headers values
func NewAdminVirtualMachineOverallReportGetUsingGETBadRequest() *AdminVirtualMachineOverallReportGetUsingGETBadRequest {
	return &AdminVirtualMachineOverallReportGetUsingGETBadRequest{}
}

/*AdminVirtualMachineOverallReportGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type AdminVirtualMachineOverallReportGetUsingGETBadRequest struct {
}

func (o *AdminVirtualMachineOverallReportGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /admin/virtual-machine/summary/{tenantId}][%d] adminVirtualMachineOverallReportGetUsingGETBadRequest ", 400)
}

func (o *AdminVirtualMachineOverallReportGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
