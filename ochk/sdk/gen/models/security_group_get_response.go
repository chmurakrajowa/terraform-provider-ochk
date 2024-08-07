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

// SecurityGroupGetResponse SecurityGroupGetResponse
//
// swagger:model SecurityGroupGetResponse
type SecurityGroupGetResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// security group
	SecurityGroup *SecurityGroup `json:"securityGroup,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this security group get response
func (m *SecurityGroupGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecurityGroup(formats); err != nil {
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

func (m *SecurityGroupGetResponse) validateSecurityGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityGroup) { // not required
		return nil
	}

	if m.SecurityGroup != nil {
		if err := m.SecurityGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securityGroup")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityGroupGetResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this security group get response based on the context it is used
func (m *SecurityGroupGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSecurityGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityGroupGetResponse) contextValidateSecurityGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityGroup != nil {

		if swag.IsZero(m.SecurityGroup) { // not required
			return nil
		}

		if err := m.SecurityGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securityGroup")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityGroupGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityGroupGetResponse) UnmarshalBinary(b []byte) error {
	var res SecurityGroupGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
