// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IP6ProfilePath IP 6 profile path
//
// swagger:model IP6ProfilePath
type IP6ProfilePath struct {

	// ip6 profile path
	Ip6ProfilePath string `json:"ip6ProfilePath,omitempty"`

	// ip6 profile path Id
	Ip6ProfilePathID int32 `json:"ip6ProfilePathId,omitempty"`
}

// Validate validates this IP 6 profile path
func (m *IP6ProfilePath) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this IP 6 profile path based on context it is used
func (m *IP6ProfilePath) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IP6ProfilePath) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IP6ProfilePath) UnmarshalBinary(b []byte) error {
	var res IP6ProfilePath
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
