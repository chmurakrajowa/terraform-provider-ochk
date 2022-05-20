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

// LocalGroupGetResponse LocalGroupGetResponse
//
// swagger:model LocalGroupGetResponse
type LocalGroupGetResponse struct {

	// local group
	LocalGroup *LocalGroup `json:"localGroup,omitempty"`

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this local group get response
func (m *LocalGroupGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocalGroup(formats); err != nil {
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

func (m *LocalGroupGetResponse) validateLocalGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.LocalGroup) { // not required
		return nil
	}

	if m.LocalGroup != nil {
		if err := m.LocalGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localGroup")
			}
			return err
		}
	}

	return nil
}

func (m *LocalGroupGetResponse) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LocalGroupGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocalGroupGetResponse) UnmarshalBinary(b []byte) error {
	var res LocalGroupGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
