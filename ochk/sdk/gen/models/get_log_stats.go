// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetLogStats GetLogStats
//
// swagger:model GetLogStats
type GetLogStats struct {

	// log category stats
	LogCategoryStats []*LogCategoryStats `json:"logCategoryStats"`

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this get log stats
func (m *GetLogStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogCategoryStats(formats); err != nil {
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

func (m *GetLogStats) validateLogCategoryStats(formats strfmt.Registry) error {
	if swag.IsZero(m.LogCategoryStats) { // not required
		return nil
	}

	for i := 0; i < len(m.LogCategoryStats); i++ {
		if swag.IsZero(m.LogCategoryStats[i]) { // not required
			continue
		}

		if m.LogCategoryStats[i] != nil {
			if err := m.LogCategoryStats[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logCategoryStats" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("logCategoryStats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetLogStats) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get log stats based on the context it is used
func (m *GetLogStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogCategoryStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetLogStats) contextValidateLogCategoryStats(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LogCategoryStats); i++ {

		if m.LogCategoryStats[i] != nil {

			if swag.IsZero(m.LogCategoryStats[i]) { // not required
				return nil
			}

			if err := m.LogCategoryStats[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logCategoryStats" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("logCategoryStats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetLogStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetLogStats) UnmarshalBinary(b []byte) error {
	var res GetLogStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
