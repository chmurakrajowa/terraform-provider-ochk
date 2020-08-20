// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkInterfaceInstance NetworkInterfaceInstance
//
// swagger:model NetworkInterfaceInstance
type NetworkInterfaceInstance struct {

	// label
	Label string `json:"label,omitempty"`

	// mac
	Mac string `json:"mac,omitempty"`

	// mac type
	// Enum: [ASSIGNED GENERATED MANUAL]
	MacType string `json:"macType,omitempty"`

	// network external Id
	NetworkExternalID string `json:"networkExternalId,omitempty"`

	// network instance
	NetworkInstance *NetworkInstance `json:"networkInstance,omitempty"`

	// network interface Id
	NetworkInterfaceID string `json:"networkInterfaceId,omitempty"`

	// nic state
	// Enum: [CONNECTED NOT_CONNECTED RECOVERABLE_ERROR UNKNOWN UNRECOVERABLE_ERROR]
	NicState string `json:"nicState,omitempty"`

	// nic type
	// Enum: [E1000 E1000E PCNET32 VMXNET VMXNET2 VMXNET3]
	NicType string `json:"nicType,omitempty"`

	// pci slot number
	PciSlotNumber int64 `json:"pciSlotNumber,omitempty"`

	// upt compatibility enabled
	UptCompatibilityEnabled bool `json:"uptCompatibilityEnabled,omitempty"`

	// wake on lan enabled
	WakeOnLanEnabled bool `json:"wakeOnLanEnabled,omitempty"`
}

// Validate validates this network interface instance
func (m *NetworkInterfaceInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMacType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkInstance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNicState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNicType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var networkInterfaceInstanceTypeMacTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ASSIGNED","GENERATED","MANUAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkInterfaceInstanceTypeMacTypePropEnum = append(networkInterfaceInstanceTypeMacTypePropEnum, v)
	}
}

const (

	// NetworkInterfaceInstanceMacTypeASSIGNED captures enum value "ASSIGNED"
	NetworkInterfaceInstanceMacTypeASSIGNED string = "ASSIGNED"

	// NetworkInterfaceInstanceMacTypeGENERATED captures enum value "GENERATED"
	NetworkInterfaceInstanceMacTypeGENERATED string = "GENERATED"

	// NetworkInterfaceInstanceMacTypeMANUAL captures enum value "MANUAL"
	NetworkInterfaceInstanceMacTypeMANUAL string = "MANUAL"
)

// prop value enum
func (m *NetworkInterfaceInstance) validateMacTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkInterfaceInstanceTypeMacTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkInterfaceInstance) validateMacType(formats strfmt.Registry) error {

	if swag.IsZero(m.MacType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMacTypeEnum("macType", "body", m.MacType); err != nil {
		return err
	}

	return nil
}

func (m *NetworkInterfaceInstance) validateNetworkInstance(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkInstance) { // not required
		return nil
	}

	if m.NetworkInstance != nil {
		if err := m.NetworkInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkInstance")
			}
			return err
		}
	}

	return nil
}

var networkInterfaceInstanceTypeNicStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CONNECTED","NOT_CONNECTED","RECOVERABLE_ERROR","UNKNOWN","UNRECOVERABLE_ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkInterfaceInstanceTypeNicStatePropEnum = append(networkInterfaceInstanceTypeNicStatePropEnum, v)
	}
}

const (

	// NetworkInterfaceInstanceNicStateCONNECTED captures enum value "CONNECTED"
	NetworkInterfaceInstanceNicStateCONNECTED string = "CONNECTED"

	// NetworkInterfaceInstanceNicStateNOTCONNECTED captures enum value "NOT_CONNECTED"
	NetworkInterfaceInstanceNicStateNOTCONNECTED string = "NOT_CONNECTED"

	// NetworkInterfaceInstanceNicStateRECOVERABLEERROR captures enum value "RECOVERABLE_ERROR"
	NetworkInterfaceInstanceNicStateRECOVERABLEERROR string = "RECOVERABLE_ERROR"

	// NetworkInterfaceInstanceNicStateUNKNOWN captures enum value "UNKNOWN"
	NetworkInterfaceInstanceNicStateUNKNOWN string = "UNKNOWN"

	// NetworkInterfaceInstanceNicStateUNRECOVERABLEERROR captures enum value "UNRECOVERABLE_ERROR"
	NetworkInterfaceInstanceNicStateUNRECOVERABLEERROR string = "UNRECOVERABLE_ERROR"
)

// prop value enum
func (m *NetworkInterfaceInstance) validateNicStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkInterfaceInstanceTypeNicStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkInterfaceInstance) validateNicState(formats strfmt.Registry) error {

	if swag.IsZero(m.NicState) { // not required
		return nil
	}

	// value enum
	if err := m.validateNicStateEnum("nicState", "body", m.NicState); err != nil {
		return err
	}

	return nil
}

var networkInterfaceInstanceTypeNicTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["E1000","E1000E","PCNET32","VMXNET","VMXNET2","VMXNET3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkInterfaceInstanceTypeNicTypePropEnum = append(networkInterfaceInstanceTypeNicTypePropEnum, v)
	}
}

const (

	// NetworkInterfaceInstanceNicTypeE1000 captures enum value "E1000"
	NetworkInterfaceInstanceNicTypeE1000 string = "E1000"

	// NetworkInterfaceInstanceNicTypeE1000E captures enum value "E1000E"
	NetworkInterfaceInstanceNicTypeE1000E string = "E1000E"

	// NetworkInterfaceInstanceNicTypePCNET32 captures enum value "PCNET32"
	NetworkInterfaceInstanceNicTypePCNET32 string = "PCNET32"

	// NetworkInterfaceInstanceNicTypeVMXNET captures enum value "VMXNET"
	NetworkInterfaceInstanceNicTypeVMXNET string = "VMXNET"

	// NetworkInterfaceInstanceNicTypeVMXNET2 captures enum value "VMXNET2"
	NetworkInterfaceInstanceNicTypeVMXNET2 string = "VMXNET2"

	// NetworkInterfaceInstanceNicTypeVMXNET3 captures enum value "VMXNET3"
	NetworkInterfaceInstanceNicTypeVMXNET3 string = "VMXNET3"
)

// prop value enum
func (m *NetworkInterfaceInstance) validateNicTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkInterfaceInstanceTypeNicTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkInterfaceInstance) validateNicType(formats strfmt.Registry) error {

	if swag.IsZero(m.NicType) { // not required
		return nil
	}

	// value enum
	if err := m.validateNicTypeEnum("nicType", "body", m.NicType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkInterfaceInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkInterfaceInstance) UnmarshalBinary(b []byte) error {
	var res NetworkInterfaceInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
