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

// LogicalPortGetResponse LogicalPortGetResponse
//
// swagger:model LogicalPortGetResponse
type LogicalPortGetResponse struct {

	// logical port
	LogicalPort *LogicalPort `json:"logicalPort,omitempty"`

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this logical port get response
func (m *LogicalPortGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogicalPort(formats); err != nil {
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

func (m *LogicalPortGetResponse) validateLogicalPort(formats strfmt.Registry) error {
	if swag.IsZero(m.LogicalPort) { // not required
		return nil
	}

	if m.LogicalPort != nil {
		if err := m.LogicalPort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logicalPort")
			}
			return err
		}
	}

	return nil
}

func (m *LogicalPortGetResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this logical port get response based on the context it is used
func (m *LogicalPortGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogicalPort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogicalPortGetResponse) contextValidateLogicalPort(ctx context.Context, formats strfmt.Registry) error {

	if m.LogicalPort != nil {
		if err := m.LogicalPort.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logicalPort")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogicalPortGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogicalPortGetResponse) UnmarshalBinary(b []byte) error {
	var res LogicalPortGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
