// Code generated by go-swagger; DO NOT EDIT.

package firewall_rules_e_w

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
)

// DFWRuleGetUsingGETReader is a Reader for the DFWRuleGetUsingGET structure.
type DFWRuleGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DFWRuleGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDFWRuleGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDFWRuleGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDFWRuleGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDFWRuleGetUsingGETOK creates a DFWRuleGetUsingGETOK with default headers values
func NewDFWRuleGetUsingGETOK() *DFWRuleGetUsingGETOK {
	return &DFWRuleGetUsingGETOK{}
}

/*DFWRuleGetUsingGETOK handles this case with default header values.

OK
*/
type DFWRuleGetUsingGETOK struct {
	Payload *models.DFWRuleGetResponse
}

func (o *DFWRuleGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /network/firewall/security-policies/{SecurityPolicyId}/rules/{RuleId}][%d] dFWRuleGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *DFWRuleGetUsingGETOK) GetPayload() *models.DFWRuleGetResponse {
	return o.Payload
}

func (o *DFWRuleGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DFWRuleGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDFWRuleGetUsingGETBadRequest creates a DFWRuleGetUsingGETBadRequest with default headers values
func NewDFWRuleGetUsingGETBadRequest() *DFWRuleGetUsingGETBadRequest {
	return &DFWRuleGetUsingGETBadRequest{}
}

/*DFWRuleGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type DFWRuleGetUsingGETBadRequest struct {
}

func (o *DFWRuleGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /network/firewall/security-policies/{SecurityPolicyId}/rules/{RuleId}][%d] dFWRuleGetUsingGETBadRequest ", 400)
}

func (o *DFWRuleGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDFWRuleGetUsingGETNotFound creates a DFWRuleGetUsingGETNotFound with default headers values
func NewDFWRuleGetUsingGETNotFound() *DFWRuleGetUsingGETNotFound {
	return &DFWRuleGetUsingGETNotFound{}
}

/*DFWRuleGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type DFWRuleGetUsingGETNotFound struct {
}

func (o *DFWRuleGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /network/firewall/security-policies/{SecurityPolicyId}/rules/{RuleId}][%d] dFWRuleGetUsingGETNotFound ", 404)
}

func (o *DFWRuleGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
