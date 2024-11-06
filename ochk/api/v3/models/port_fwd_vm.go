// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PortFwdVM port fwd Vm
//
// swagger:model PortFwdVm
type PortFwdVM struct {

	// ip address
	IPAddress string `json:"ipAddress,omitempty"`

	// mac address
	MacAddress string `json:"macAddress,omitempty"`

	// network interface label
	NetworkInterfaceLabel string `json:"networkInterfaceLabel,omitempty"`

	// network name
	NetworkName string `json:"networkName,omitempty"`

	// osc port Id
	// Format: uuid
	OscPortID strfmt.UUID `json:"oscPortId,omitempty"`

	// virtual machine Id
	// Format: uuid
	VirtualMachineID strfmt.UUID `json:"virtualMachineId,omitempty"`

	// virtual machine name
	VirtualMachineName string `json:"virtualMachineName,omitempty"`

	// vm external Id
	VMExternalID string `json:"vmExternalId,omitempty"`
}

// Validate validates this port fwd Vm
func (m *PortFwdVM) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOscPortID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualMachineID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortFwdVM) validateOscPortID(formats strfmt.Registry) error {
	if swag.IsZero(m.OscPortID) { // not required
		return nil
	}

	if err := validate.FormatOf("oscPortId", "body", "uuid", m.OscPortID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PortFwdVM) validateVirtualMachineID(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualMachineID) { // not required
		return nil
	}

	if err := validate.FormatOf("virtualMachineId", "body", "uuid", m.VirtualMachineID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this port fwd Vm based on context it is used
func (m *PortFwdVM) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortFwdVM) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortFwdVM) UnmarshalBinary(b []byte) error {
	var res PortFwdVM
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
