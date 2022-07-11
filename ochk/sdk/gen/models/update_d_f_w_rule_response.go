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

// UpdateDFWRuleResponse UpdateDFWRuleResponse
//
// swagger:model UpdateDFWRuleResponse
type UpdateDFWRuleResponse struct {

	// dfw rule
	DfwRule *DFWRule `json:"dfwRule,omitempty"`

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

// Validate validates this update d f w rule response
func (m *UpdateDFWRuleResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDfwRule(formats); err != nil {
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

func (m *UpdateDFWRuleResponse) validateDfwRule(formats strfmt.Registry) error {
	if swag.IsZero(m.DfwRule) { // not required
		return nil
	}

	if m.DfwRule != nil {
		if err := m.DfwRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dfwRule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dfwRule")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateDFWRuleResponse) validateRequestInstance(formats strfmt.Registry) error {
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

func (m *UpdateDFWRuleResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update d f w rule response based on the context it is used
func (m *UpdateDFWRuleResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDfwRule(ctx, formats); err != nil {
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

func (m *UpdateDFWRuleResponse) contextValidateDfwRule(ctx context.Context, formats strfmt.Registry) error {

	if m.DfwRule != nil {
		if err := m.DfwRule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dfwRule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dfwRule")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateDFWRuleResponse) contextValidateRequestInstance(ctx context.Context, formats strfmt.Registry) error {

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
func (m *UpdateDFWRuleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateDFWRuleResponse) UnmarshalBinary(b []byte) error {
	var res UpdateDFWRuleResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
