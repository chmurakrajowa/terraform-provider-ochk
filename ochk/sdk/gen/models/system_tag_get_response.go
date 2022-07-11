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

// SystemTagGetResponse SystemTagGetResponse
//
// swagger:model SystemTagGetResponse
type SystemTagGetResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// system tag
	SystemTag *SystemTag `json:"systemTag,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this system tag get response
func (m *SystemTagGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSystemTag(formats); err != nil {
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

func (m *SystemTagGetResponse) validateSystemTag(formats strfmt.Registry) error {
	if swag.IsZero(m.SystemTag) { // not required
		return nil
	}

	if m.SystemTag != nil {
		if err := m.SystemTag.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("systemTag")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("systemTag")
			}
			return err
		}
	}

	return nil
}

func (m *SystemTagGetResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this system tag get response based on the context it is used
func (m *SystemTagGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSystemTag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemTagGetResponse) contextValidateSystemTag(ctx context.Context, formats strfmt.Registry) error {

	if m.SystemTag != nil {
		if err := m.SystemTag.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("systemTag")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("systemTag")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SystemTagGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemTagGetResponse) UnmarshalBinary(b []byte) error {
	var res SystemTagGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
