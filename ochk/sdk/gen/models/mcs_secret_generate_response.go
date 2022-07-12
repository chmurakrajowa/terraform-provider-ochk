// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// McsSecretGenerateResponse McsSecretGenerateResponse
//
// swagger:model McsSecretGenerateResponse
type McsSecretGenerateResponse struct {

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

// Validate validates this mcs secret generate response
func (m *McsSecretGenerateResponse) Validate(formats strfmt.Registry) error {
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

func (m *McsSecretGenerateResponse) validateMcsSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.McsSecret) { // not required
		return nil
	}

	if m.McsSecret != nil {
		if err := m.McsSecret.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mcsSecret")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mcsSecret")
			}
			return err
		}
	}

	return nil
}

func (m *McsSecretGenerateResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this mcs secret generate response based on the context it is used
func (m *McsSecretGenerateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMcsSecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *McsSecretGenerateResponse) contextValidateMcsSecret(ctx context.Context, formats strfmt.Registry) error {

	if m.McsSecret != nil {
		if err := m.McsSecret.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mcsSecret")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mcsSecret")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *McsSecretGenerateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *McsSecretGenerateResponse) UnmarshalBinary(b []byte) error {
	var res McsSecretGenerateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
