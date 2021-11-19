// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// McsSecretGetResponse McsSecretGetResponse
//
// swagger:model McsSecretGetResponse
type McsSecretGetResponse struct {

	// mcs secret
	McsSecret *McsSecret `json:"mcsSecret,omitempty"`

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this mcs secret get response
func (m *McsSecretGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMcsSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *McsSecretGetResponse) validateMcsSecret(formats strfmt.Registry) error {

	if swag.IsZero(m.McsSecret) { // not required
		return nil
	}

	if m.McsSecret != nil {
		if err := m.McsSecret.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mcsSecret")
			}
			return err
		}
	}

	return nil
}

func (m *McsSecretGetResponse) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *McsSecretGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *McsSecretGetResponse) UnmarshalBinary(b []byte) error {
	var res McsSecretGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
