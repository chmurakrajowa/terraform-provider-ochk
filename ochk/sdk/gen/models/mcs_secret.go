// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// McsSecret McsSecret
//
// swagger:model McsSecret
type McsSecret struct {

	// secret
	Secret string `json:"secret,omitempty"`
}

// Validate validates this mcs secret
func (m *McsSecret) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this mcs secret based on context it is used
func (m *McsSecret) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *McsSecret) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *McsSecret) UnmarshalBinary(b []byte) error {
	var res McsSecret
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
