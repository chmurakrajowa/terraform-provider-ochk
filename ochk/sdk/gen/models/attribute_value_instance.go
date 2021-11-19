// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AttributeValueInstance AttributeValueInstance
//
// swagger:model AttributeValueInstance
type AttributeValueInstance struct {

	// value
	Value string `json:"value,omitempty"`

	// value Id
	ValueID int32 `json:"valueId,omitempty"`
}

// Validate validates this attribute value instance
func (m *AttributeValueInstance) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this attribute value instance based on context it is used
func (m *AttributeValueInstance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AttributeValueInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttributeValueInstance) UnmarshalBinary(b []byte) error {
	var res AttributeValueInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
