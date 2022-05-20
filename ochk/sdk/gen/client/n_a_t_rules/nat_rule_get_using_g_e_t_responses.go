// Code generated by go-swagger; DO NOT EDIT.

package n_a_t_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// NatRuleGetUsingGETReader is a Reader for the NatRuleGetUsingGET structure.
type NatRuleGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NatRuleGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNatRuleGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNatRuleGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNatRuleGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNatRuleGetUsingGETOK creates a NatRuleGetUsingGETOK with default headers values
func NewNatRuleGetUsingGETOK() *NatRuleGetUsingGETOK {
	return &NatRuleGetUsingGETOK{}
}

/*NatRuleGetUsingGETOK handles this case with default header values.

OK
*/
type NatRuleGetUsingGETOK struct {
	Payload *models.NATRuleGetResponse
}

func (o *NatRuleGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /network/nat/rules/{ruleId}][%d] natRuleGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *NatRuleGetUsingGETOK) GetPayload() *models.NATRuleGetResponse {
	return o.Payload
}

func (o *NatRuleGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NATRuleGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNatRuleGetUsingGETBadRequest creates a NatRuleGetUsingGETBadRequest with default headers values
func NewNatRuleGetUsingGETBadRequest() *NatRuleGetUsingGETBadRequest {
	return &NatRuleGetUsingGETBadRequest{}
}

/*NatRuleGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type NatRuleGetUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *NatRuleGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /network/nat/rules/{ruleId}][%d] natRuleGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *NatRuleGetUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *NatRuleGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNatRuleGetUsingGETNotFound creates a NatRuleGetUsingGETNotFound with default headers values
func NewNatRuleGetUsingGETNotFound() *NatRuleGetUsingGETNotFound {
	return &NatRuleGetUsingGETNotFound{}
}

/*NatRuleGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type NatRuleGetUsingGETNotFound struct {
}

func (o *NatRuleGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /network/nat/rules/{ruleId}][%d] natRuleGetUsingGETNotFound ", 404)
}

func (o *NatRuleGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
