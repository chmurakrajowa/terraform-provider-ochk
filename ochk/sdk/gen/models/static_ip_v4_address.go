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

// StaticIPV4Address StaticIPv4Address
//
// swagger:model StaticIPv4Address
type StaticIPV4Address struct {

	// IP sort value
	IPSortValue int64 `json:"IPSortValue,omitempty"`

	// IPv4 address
	IPV4Address string `json:"IPv4Address,omitempty"`

	// created date
	// Format: date-time
	CreatedDate *strfmt.DateTime `json:"createdDate,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// external reference Id
	ExternalReferenceID string `json:"externalReferenceId,omitempty"`

	// external reference name
	ExternalReferenceName string `json:"externalReferenceName,omitempty"`

	// host name
	HostName string `json:"hostName,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last modified date
	// Format: date-time
	LastModifiedDate *strfmt.DateTime `json:"lastModifiedDate,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// network interface card offset
	NetworkInterfaceCardOffset string `json:"networkInterfaceCardOffset,omitempty"`

	// state
	// Enum: [ALLOCATED DESTROYED EXPIRED UNALLOCATED]
	State string `json:"state,omitempty"`

	// state value
	StateValue int32 `json:"stateValue,omitempty"`

	// static IPv4 address Id
	StaticIPV4AddressID string `json:"staticIPv4AddressId,omitempty"`

	// virtual machine Id
	VirtualMachineID string `json:"virtualMachineId,omitempty"`

	// virtual machine name
	VirtualMachineName string `json:"virtualMachineName,omitempty"`
}

// Validate validates this static IPv4 address
func (m *StaticIPV4Address) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StaticIPV4Address) validateCreatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StaticIPV4Address) validateLastModifiedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("lastModifiedDate", "body", "date-time", m.LastModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var staticIpv4AddressTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALLOCATED","DESTROYED","EXPIRED","UNALLOCATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		staticIpv4AddressTypeStatePropEnum = append(staticIpv4AddressTypeStatePropEnum, v)
	}
}

const (

	// StaticIPV4AddressStateALLOCATED captures enum value "ALLOCATED"
	StaticIPV4AddressStateALLOCATED string = "ALLOCATED"

	// StaticIPV4AddressStateDESTROYED captures enum value "DESTROYED"
	StaticIPV4AddressStateDESTROYED string = "DESTROYED"

	// StaticIPV4AddressStateEXPIRED captures enum value "EXPIRED"
	StaticIPV4AddressStateEXPIRED string = "EXPIRED"

	// StaticIPV4AddressStateUNALLOCATED captures enum value "UNALLOCATED"
	StaticIPV4AddressStateUNALLOCATED string = "UNALLOCATED"
)

// prop value enum
func (m *StaticIPV4Address) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, staticIpv4AddressTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StaticIPV4Address) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StaticIPV4Address) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StaticIPV4Address) UnmarshalBinary(b []byte) error {
	var res StaticIPV4Address
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
