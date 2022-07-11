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

// NatRuleListUsingGETReader is a Reader for the NatRuleListUsingGET structure.
type NatRuleListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NatRuleListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNatRuleListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNatRuleListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNatRuleListUsingGETOK creates a NatRuleListUsingGETOK with default headers values
func NewNatRuleListUsingGETOK() *NatRuleListUsingGETOK {
	return &NatRuleListUsingGETOK{}
}

/*NatRuleListUsingGETOK handles this case with default header values.

OK
*/
type NatRuleListUsingGETOK struct {
	Payload *models.NATRuleListResponse
}

func (o *NatRuleListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /network/nat/rules][%d] natRuleListUsingGETOK  %+v", 200, o.Payload)
}

func (o *NatRuleListUsingGETOK) GetPayload() *models.NATRuleListResponse {
	return o.Payload
}

func (o *NatRuleListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NATRuleListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNatRuleListUsingGETBadRequest creates a NatRuleListUsingGETBadRequest with default headers values
func NewNatRuleListUsingGETBadRequest() *NatRuleListUsingGETBadRequest {
	return &NatRuleListUsingGETBadRequest{}
}

/*NatRuleListUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type NatRuleListUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *NatRuleListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /network/nat/rules][%d] natRuleListUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *NatRuleListUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *NatRuleListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
