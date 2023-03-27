// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WSOUITokenRequest WSOUITokenRequest
//
// swagger:model WSOUITokenRequest
type WSOUITokenRequest struct {

	// access token
	AccessToken string `json:"accessToken,omitempty"`

	// platform
	Platform string `json:"platform,omitempty"`
}

// Validate validates this w s o UI token request
func (m *WSOUITokenRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this w s o UI token request based on context it is used
func (m *WSOUITokenRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WSOUITokenRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WSOUITokenRequest) UnmarshalBinary(b []byte) error {
	var res WSOUITokenRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}