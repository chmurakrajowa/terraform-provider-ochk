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

// NetworkInstance NetworkInstance
//
// swagger:model NetworkInstance
type NetworkInstance struct {

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// creation date
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate,omitempty"`

	// modification date
	// Format: date-time
	ModificationDate *strfmt.DateTime `json:"modificationDate,omitempty"`

	// modified by
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// network Id
	NetworkID string `json:"networkId,omitempty"`

	// network type
	// Enum: [DISTRIBUTED_PORTGROUP HOST_DEVICE OPAQUE_NETWORK STANDARD_PORTGROUP]
	NetworkType string `json:"networkType,omitempty"`
}

// Validate validates this network instance
func (m *NetworkInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModificationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkInstance) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NetworkInstance) validateModificationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ModificationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modificationDate", "body", "date-time", m.ModificationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var networkInstanceTypeNetworkTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DISTRIBUTED_PORTGROUP","HOST_DEVICE","OPAQUE_NETWORK","STANDARD_PORTGROUP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkInstanceTypeNetworkTypePropEnum = append(networkInstanceTypeNetworkTypePropEnum, v)
	}
}

const (

	// NetworkInstanceNetworkTypeDISTRIBUTEDPORTGROUP captures enum value "DISTRIBUTED_PORTGROUP"
	NetworkInstanceNetworkTypeDISTRIBUTEDPORTGROUP string = "DISTRIBUTED_PORTGROUP"

	// NetworkInstanceNetworkTypeHOSTDEVICE captures enum value "HOST_DEVICE"
	NetworkInstanceNetworkTypeHOSTDEVICE string = "HOST_DEVICE"

	// NetworkInstanceNetworkTypeOPAQUENETWORK captures enum value "OPAQUE_NETWORK"
	NetworkInstanceNetworkTypeOPAQUENETWORK string = "OPAQUE_NETWORK"

	// NetworkInstanceNetworkTypeSTANDARDPORTGROUP captures enum value "STANDARD_PORTGROUP"
	NetworkInstanceNetworkTypeSTANDARDPORTGROUP string = "STANDARD_PORTGROUP"
)

// prop value enum
func (m *NetworkInstance) validateNetworkTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkInstanceTypeNetworkTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkInstance) validateNetworkType(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkType) { // not required
		return nil
	}

	// value enum
	if err := m.validateNetworkTypeEnum("networkType", "body", m.NetworkType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkInstance) UnmarshalBinary(b []byte) error {
	var res NetworkInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
