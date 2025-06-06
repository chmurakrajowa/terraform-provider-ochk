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

// UpdateNATRuleResponse update n a t rule response
//
// swagger:model UpdateNATRuleResponse
type UpdateNATRuleResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// nat rule instance
	NatRuleInstance *NATRuleInstance `json:"natRuleInstance,omitempty"`

	// request instance
	RequestInstance *RequestInstance `json:"requestInstance,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this update n a t rule response
func (m *UpdateNATRuleResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNatRuleInstance(formats); err != nil {
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

func (m *UpdateNATRuleResponse) validateNatRuleInstance(formats strfmt.Registry) error {
	if swag.IsZero(m.NatRuleInstance) { // not required
		return nil
	}

	if m.NatRuleInstance != nil {
		if err := m.NatRuleInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("natRuleInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("natRuleInstance")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateNATRuleResponse) validateRequestInstance(formats strfmt.Registry) error {
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

func (m *UpdateNATRuleResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update n a t rule response based on the context it is used
func (m *UpdateNATRuleResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNatRuleInstance(ctx, formats); err != nil {
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

func (m *UpdateNATRuleResponse) contextValidateNatRuleInstance(ctx context.Context, formats strfmt.Registry) error {

	if m.NatRuleInstance != nil {

		if swag.IsZero(m.NatRuleInstance) { // not required
			return nil
		}

		if err := m.NatRuleInstance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("natRuleInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("natRuleInstance")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateNATRuleResponse) contextValidateRequestInstance(ctx context.Context, formats strfmt.Registry) error {

	if m.RequestInstance != nil {

		if swag.IsZero(m.RequestInstance) { // not required
			return nil
		}

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
func (m *UpdateNATRuleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNATRuleResponse) UnmarshalBinary(b []byte) error {
	var res UpdateNATRuleResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
