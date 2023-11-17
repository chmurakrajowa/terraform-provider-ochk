// Code generated by go-swagger; DO NOT EDIT.

package billing_vm_report

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// GenerateVMBillingUsingGETReader is a Reader for the GenerateVMBillingUsingGET structure.
type GenerateVMBillingUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateVMBillingUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateVMBillingUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /billing/accounts/{accountId}/reports/{billingPeriodId}/vm] generateVMBillingUsingGET", response, response.Code())
	}
}

// NewGenerateVMBillingUsingGETOK creates a GenerateVMBillingUsingGETOK with default headers values
func NewGenerateVMBillingUsingGETOK() *GenerateVMBillingUsingGETOK {
	return &GenerateVMBillingUsingGETOK{}
}

/*
GenerateVMBillingUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GenerateVMBillingUsingGETOK struct {
	Payload *models.GenerateVirtualMachineBillingReportResponse
}

// IsSuccess returns true when this generate Vm billing using g e t o k response has a 2xx status code
func (o *GenerateVMBillingUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this generate Vm billing using g e t o k response has a 3xx status code
func (o *GenerateVMBillingUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate Vm billing using g e t o k response has a 4xx status code
func (o *GenerateVMBillingUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate Vm billing using g e t o k response has a 5xx status code
func (o *GenerateVMBillingUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this generate Vm billing using g e t o k response a status code equal to that given
func (o *GenerateVMBillingUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the generate Vm billing using g e t o k response
func (o *GenerateVMBillingUsingGETOK) Code() int {
	return 200
}

func (o *GenerateVMBillingUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /billing/accounts/{accountId}/reports/{billingPeriodId}/vm][%d] generateVmBillingUsingGETOK  %+v", 200, o.Payload)
}

func (o *GenerateVMBillingUsingGETOK) String() string {
	return fmt.Sprintf("[GET /billing/accounts/{accountId}/reports/{billingPeriodId}/vm][%d] generateVmBillingUsingGETOK  %+v", 200, o.Payload)
}

func (o *GenerateVMBillingUsingGETOK) GetPayload() *models.GenerateVirtualMachineBillingReportResponse {
	return o.Payload
}

func (o *GenerateVMBillingUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenerateVirtualMachineBillingReportResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
