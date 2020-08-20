// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TransitSubnet TransitSubnet
//
// swagger:model TransitSubnet
type TransitSubnet struct {

	// transit subnet
	TransitSubnet string `json:"transitSubnet,omitempty"`

	// transit subnet Id
	TransitSubnetID int32 `json:"transitSubnetId,omitempty"`
}

// Validate validates this transit subnet
func (m *TransitSubnet) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TransitSubnet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransitSubnet) UnmarshalBinary(b []byte) error {
	var res TransitSubnet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
