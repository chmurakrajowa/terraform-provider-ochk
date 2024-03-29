// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetPublicKeyRequest GetPublicKeyRequest
//
// swagger:model GetPublicKeyRequest
type GetPublicKeyRequest struct {

	// sha256 check sum
	Sha256CheckSum string `json:"sha256CheckSum,omitempty"`
}

// Validate validates this get public key request
func (m *GetPublicKeyRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get public key request based on context it is used
func (m *GetPublicKeyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetPublicKeyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetPublicKeyRequest) UnmarshalBinary(b []byte) error {
	var res GetPublicKeyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
