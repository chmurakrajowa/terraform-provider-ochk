// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SubtenantInstance subtenant instance
//
// swagger:model SubtenantInstance
type SubtenantInstance struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// subtenant Id
	SubtenantID string `json:"subtenantId,omitempty"`
}

// Validate validates this subtenant instance
func (m *SubtenantInstance) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SubtenantInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubtenantInstance) UnmarshalBinary(b []byte) error {
	var res SubtenantInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}