// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SystemTag SystemTag
//
// swagger:model SystemTag
type SystemTag struct {

	// system tag Id
	SystemTagID int32 `json:"systemTagId,omitempty"`

	// tag value
	TagValue string `json:"tagValue,omitempty"`
}

// Validate validates this system tag
func (m *SystemTag) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this system tag based on context it is used
func (m *SystemTag) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemTag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemTag) UnmarshalBinary(b []byte) error {
	var res SystemTag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
