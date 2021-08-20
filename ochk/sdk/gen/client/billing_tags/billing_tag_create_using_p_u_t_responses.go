// Code generated by go-swagger; DO NOT EDIT.

package billing_tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// BillingTagCreateUsingPUTReader is a Reader for the BillingTagCreateUsingPUT structure.
type BillingTagCreateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingTagCreateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBillingTagCreateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewBillingTagCreateUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBillingTagCreateUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBillingTagCreateUsingPUTOK creates a BillingTagCreateUsingPUTOK with default headers values
func NewBillingTagCreateUsingPUTOK() *BillingTagCreateUsingPUTOK {
	return &BillingTagCreateUsingPUTOK{}
}

/*BillingTagCreateUsingPUTOK handles this case with default header values.

OK
*/
type BillingTagCreateUsingPUTOK struct {
	Payload *models.BillingTagCreateResponse
}

func (o *BillingTagCreateUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /tags/billingTags][%d] billingTagCreateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *BillingTagCreateUsingPUTOK) GetPayload() *models.BillingTagCreateResponse {
	return o.Payload
}

func (o *BillingTagCreateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingTagCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingTagCreateUsingPUTCreated creates a BillingTagCreateUsingPUTCreated with default headers values
func NewBillingTagCreateUsingPUTCreated() *BillingTagCreateUsingPUTCreated {
	return &BillingTagCreateUsingPUTCreated{}
}

/*BillingTagCreateUsingPUTCreated handles this case with default header values.

Entity has been created
*/
type BillingTagCreateUsingPUTCreated struct {
	Payload *models.BillingTagCreateResponse
}

func (o *BillingTagCreateUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /tags/billingTags][%d] billingTagCreateUsingPUTCreated  %+v", 201, o.Payload)
}

func (o *BillingTagCreateUsingPUTCreated) GetPayload() *models.BillingTagCreateResponse {
	return o.Payload
}

func (o *BillingTagCreateUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingTagCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingTagCreateUsingPUTBadRequest creates a BillingTagCreateUsingPUTBadRequest with default headers values
func NewBillingTagCreateUsingPUTBadRequest() *BillingTagCreateUsingPUTBadRequest {
	return &BillingTagCreateUsingPUTBadRequest{}
}

/*BillingTagCreateUsingPUTBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type BillingTagCreateUsingPUTBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *BillingTagCreateUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /tags/billingTags][%d] billingTagCreateUsingPUTBadRequest  %+v", 400, o.Payload)
}

func (o *BillingTagCreateUsingPUTBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *BillingTagCreateUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}