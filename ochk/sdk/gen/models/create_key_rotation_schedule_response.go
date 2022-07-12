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

// CreateKeyRotationScheduleResponse CreateKeyRotationScheduleResponse
//
// swagger:model CreateKeyRotationScheduleResponse
type CreateKeyRotationScheduleResponse struct {

	// key rotation schedule
	KeyRotationSchedule *KeyRotationSchedule `json:"keyRotationSchedule,omitempty"`

	// messages
	Messages string `json:"messages,omitempty"`

	// request instance
	RequestInstance *RequestInstance `json:"requestInstance,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this create key rotation schedule response
func (m *CreateKeyRotationScheduleResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeyRotationSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestInstance(formats); err != nil {
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

func (m *CreateKeyRotationScheduleResponse) validateKeyRotationSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.KeyRotationSchedule) { // not required
		return nil
	}

	if m.KeyRotationSchedule != nil {
		if err := m.KeyRotationSchedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyRotationSchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyRotationSchedule")
			}
			return err
		}
	}

	return nil
}

func (m *CreateKeyRotationScheduleResponse) validateRequestInstance(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestInstance) { // not required
		return nil
	}

	if m.RequestInstance != nil {
		if err := m.RequestInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requestInstance")
			}
			return err
		}
	}

	return nil
}

func (m *CreateKeyRotationScheduleResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create key rotation schedule response based on the context it is used
func (m *CreateKeyRotationScheduleResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKeyRotationSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequestInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateKeyRotationScheduleResponse) contextValidateKeyRotationSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.KeyRotationSchedule != nil {
		if err := m.KeyRotationSchedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyRotationSchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyRotationSchedule")
			}
			return err
		}
	}

	return nil
}

func (m *CreateKeyRotationScheduleResponse) contextValidateRequestInstance(ctx context.Context, formats strfmt.Registry) error {

	if m.RequestInstance != nil {
		if err := m.RequestInstance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requestInstance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateKeyRotationScheduleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateKeyRotationScheduleResponse) UnmarshalBinary(b []byte) error {
	var res CreateKeyRotationScheduleResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
