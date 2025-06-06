// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AllocationPool allocation pool
//
// swagger:model AllocationPool
type AllocationPool struct {

	// end
	End string `json:"end,omitempty"`

	// start
	Start string `json:"start,omitempty"`
}

// Validate validates this allocation pool
func (m *AllocationPool) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this allocation pool based on context it is used
func (m *AllocationPool) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AllocationPool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllocationPool) UnmarshalBinary(b []byte) error {
	var res AllocationPool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
