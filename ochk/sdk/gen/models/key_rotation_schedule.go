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

// KeyRotationSchedule KeyRotationSchedule
//
// swagger:model KeyRotationSchedule
type KeyRotationSchedule struct {

	// cycle
	Cycle bool `json:"cycle,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// entry Id
	EntryID int32 `json:"entryId,omitempty"`

	// interval hours
	IntervalHours int32 `json:"intervalHours,omitempty"`

	// key Id
	KeyID string `json:"keyId,omitempty"`

	// start time
	// Format: date-time
	StartTime *strfmt.DateTime `json:"startTime,omitempty"`

	// static date
	// Format: date-time
	StaticDate *strfmt.DateTime `json:"staticDate,omitempty"`
}

// Validate validates this key rotation schedule
func (m *KeyRotationSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeyRotationSchedule) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KeyRotationSchedule) validateStaticDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StaticDate) { // not required
		return nil
	}

	if err := validate.FormatOf("staticDate", "body", "date-time", m.StaticDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this key rotation schedule based on context it is used
func (m *KeyRotationSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KeyRotationSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyRotationSchedule) UnmarshalBinary(b []byte) error {
	var res KeyRotationSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
