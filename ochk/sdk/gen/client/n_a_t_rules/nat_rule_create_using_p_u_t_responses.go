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

// NatRuleCreateUsingPUTReader is a Reader for the NatRuleCreateUsingPUT structure.
type NatRuleCreateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NatRuleCreateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNatRuleCreateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewNatRuleCreateUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNatRuleCreateUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNatRuleCreateUsingPUTOK creates a NatRuleCreateUsingPUTOK with default headers values
func NewNatRuleCreateUsingPUTOK() *NatRuleCreateUsingPUTOK {
	return &NatRuleCreateUsingPUTOK{}
}

/* NatRuleCreateUsingPUTOK describes a response with status code 200, with default header values.

OK
*/
type NatRuleCreateUsingPUTOK struct {
	Payload *models.CreateNATRuleResponse
}

func (o *NatRuleCreateUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /network/nat/rules][%d] natRuleCreateUsingPUTOK  %+v", 200, o.Payload)
}
func (o *NatRuleCreateUsingPUTOK) GetPayload() *models.CreateNATRuleResponse {
	return o.Payload
}

func (o *NatRuleCreateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateNATRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNatRuleCreateUsingPUTCreated creates a NatRuleCreateUsingPUTCreated with default headers values
func NewNatRuleCreateUsingPUTCreated() *NatRuleCreateUsingPUTCreated {
	return &NatRuleCreateUsingPUTCreated{}
}

/* NatRuleCreateUsingPUTCreated describes a response with status code 201, with default header values.

Entity has been created
*/
type NatRuleCreateUsingPUTCreated struct {
	Payload *models.CreateNATRuleResponse
}

func (o *NatRuleCreateUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /network/nat/rules][%d] natRuleCreateUsingPUTCreated  %+v", 201, o.Payload)
}
func (o *NatRuleCreateUsingPUTCreated) GetPayload() *models.CreateNATRuleResponse {
	return o.Payload
}

func (o *NatRuleCreateUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateNATRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNatRuleCreateUsingPUTBadRequest creates a NatRuleCreateUsingPUTBadRequest with default headers values
func NewNatRuleCreateUsingPUTBadRequest() *NatRuleCreateUsingPUTBadRequest {
	return &NatRuleCreateUsingPUTBadRequest{}
}

/* NatRuleCreateUsingPUTBadRequest describes a response with status code 400, with default header values.

Bad request, error occurred. For more details see log messages.
*/
type NatRuleCreateUsingPUTBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *NatRuleCreateUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /network/nat/rules][%d] natRuleCreateUsingPUTBadRequest  %+v", 400, o.Payload)
}
func (o *NatRuleCreateUsingPUTBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *NatRuleCreateUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
