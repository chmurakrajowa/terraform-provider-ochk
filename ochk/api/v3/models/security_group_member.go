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

// SecurityGroupMember security group member
//
// swagger:model SecurityGroupMember
type SecurityGroupMember struct {

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// external Id
	ExternalID string `json:"externalId,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// member type
	MemberType SecurityGroupMemberType `json:"memberType,omitempty"`
}

// Validate validates this security group member
func (m *SecurityGroupMember) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityGroupMember) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SecurityGroupMember) validateMemberType(formats strfmt.Registry) error {
	if swag.IsZero(m.MemberType) { // not required
		return nil
	}

	if err := m.MemberType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("memberType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("memberType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this security group member based on the context it is used
func (m *SecurityGroupMember) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMemberType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityGroupMember) contextValidateMemberType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.MemberType) { // not required
		return nil
	}

	if err := m.MemberType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("memberType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("memberType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityGroupMember) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityGroupMember) UnmarshalBinary(b []byte) error {
	var res SecurityGroupMember
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
