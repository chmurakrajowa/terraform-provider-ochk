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

// SubtenantGetResponse SubtenantGetResponse
//
// swagger:model SubtenantGetResponse
type SubtenantGetResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// subtenant instance
	SubtenantInstance *SubtenantInstance `json:"subtenantInstance,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this subtenant get response
func (m *SubtenantGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubtenantInstance(formats); err != nil {
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

func (m *SubtenantGetResponse) validateSubtenantInstance(formats strfmt.Registry) error {

	if swag.IsZero(m.SubtenantInstance) { // not required
		return nil
	}

	if m.SubtenantInstance != nil {
		if err := m.SubtenantInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subtenantInstance")
			}
			return err
		}
	}

	return nil
}

func (m *SubtenantGetResponse) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubtenantGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubtenantGetResponse) UnmarshalBinary(b []byte) error {
	var res SubtenantGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
