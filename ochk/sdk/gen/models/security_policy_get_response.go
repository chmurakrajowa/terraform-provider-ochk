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

// SecurityPolicyGetResponse SecurityPolicyGetResponse
//
// swagger:model SecurityPolicyGetResponse
type SecurityPolicyGetResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// security policy
	SecurityPolicy *SecurityPolicy `json:"securityPolicy,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this security policy get response
func (m *SecurityPolicyGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecurityPolicy(formats); err != nil {
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

func (m *SecurityPolicyGetResponse) validateSecurityPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityPolicy) { // not required
		return nil
	}

	if m.SecurityPolicy != nil {
		if err := m.SecurityPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securityPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityPolicyGetResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this security policy get response based on the context it is used
func (m *SecurityPolicyGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSecurityPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityPolicyGetResponse) contextValidateSecurityPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityPolicy != nil {
		if err := m.SecurityPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securityPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityPolicyGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityPolicyGetResponse) UnmarshalBinary(b []byte) error {
	var res SecurityPolicyGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
