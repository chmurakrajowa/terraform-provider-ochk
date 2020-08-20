// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConnectivityStrategy ConnectivityStrategy
//
// swagger:model ConnectivityStrategy
type ConnectivityStrategy struct {

	// connectivity strategy
	ConnectivityStrategy string `json:"connectivityStrategy,omitempty"`

	// connectivity strategy Id
	ConnectivityStrategyID int32 `json:"connectivityStrategyId,omitempty"`
}

// Validate validates this connectivity strategy
func (m *ConnectivityStrategy) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConnectivityStrategy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectivityStrategy) UnmarshalBinary(b []byte) error {
	var res ConnectivityStrategy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
