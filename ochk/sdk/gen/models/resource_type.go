// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResourceType resource type
//
// swagger:model ResourceType
type ResourceType struct {

	// resource type Id
	ResourceTypeID int32 `json:"resourceTypeId,omitempty"`

	// resource type name
	ResourceTypeName string `json:"resourceTypeName,omitempty"`
}

// Validate validates this resource type
func (m *ResourceType) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceType) UnmarshalBinary(b []byte) error {
	var res ResourceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
