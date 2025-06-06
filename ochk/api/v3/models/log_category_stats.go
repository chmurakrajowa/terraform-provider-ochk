// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LogCategoryStats log category stats
//
// swagger:model LogCategoryStats
type LogCategoryStats struct {

	// log category
	LogCategory *LogCategory `json:"logCategory,omitempty"`

	// logs count
	LogsCount int64 `json:"logsCount,omitempty"`
}

// Validate validates this log category stats
func (m *LogCategoryStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogCategory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogCategoryStats) validateLogCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.LogCategory) { // not required
		return nil
	}

	if m.LogCategory != nil {
		if err := m.LogCategory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logCategory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logCategory")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this log category stats based on the context it is used
func (m *LogCategoryStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogCategoryStats) contextValidateLogCategory(ctx context.Context, formats strfmt.Registry) error {

	if m.LogCategory != nil {

		if swag.IsZero(m.LogCategory) { // not required
			return nil
		}

		if err := m.LogCategory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logCategory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logCategory")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogCategoryStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogCategoryStats) UnmarshalBinary(b []byte) error {
	var res LogCategoryStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
