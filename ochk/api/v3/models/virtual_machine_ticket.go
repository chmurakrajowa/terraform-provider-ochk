// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VirtualMachineTicket virtual machine ticket
//
// swagger:model VirtualMachineTicket
type VirtualMachineTicket struct {

	// host
	Host string `json:"host,omitempty"`

	// ticket
	Ticket string `json:"ticket,omitempty"`
}

// Validate validates this virtual machine ticket
func (m *VirtualMachineTicket) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this virtual machine ticket based on context it is used
func (m *VirtualMachineTicket) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VirtualMachineTicket) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualMachineTicket) UnmarshalBinary(b []byte) error {
	var res VirtualMachineTicket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
