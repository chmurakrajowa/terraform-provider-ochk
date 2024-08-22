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

// ListGfwRulesResponse list gfw rules response
//
// swagger:model ListGfwRulesResponse
type ListGfwRulesResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// rule instances
	RuleInstances []*GfwRule `json:"ruleInstances"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this list gfw rules response
func (m *ListGfwRulesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRuleInstances(formats); err != nil {
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

func (m *ListGfwRulesResponse) validateRuleInstances(formats strfmt.Registry) error {
	if swag.IsZero(m.RuleInstances) { // not required
		return nil
	}

	for i := 0; i < len(m.RuleInstances); i++ {
		if swag.IsZero(m.RuleInstances[i]) { // not required
			continue
		}

		if m.RuleInstances[i] != nil {
			if err := m.RuleInstances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ruleInstances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ruleInstances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListGfwRulesResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this list gfw rules response based on the context it is used
func (m *ListGfwRulesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRuleInstances(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListGfwRulesResponse) contextValidateRuleInstances(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RuleInstances); i++ {

		if m.RuleInstances[i] != nil {

			if swag.IsZero(m.RuleInstances[i]) { // not required
				return nil
			}

			if err := m.RuleInstances[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ruleInstances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ruleInstances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListGfwRulesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListGfwRulesResponse) UnmarshalBinary(b []byte) error {
	var res ListGfwRulesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
