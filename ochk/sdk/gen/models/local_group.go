// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LocalGroup LocalGroup
//
// swagger:model LocalGroup
type LocalGroup struct {

	// description
	Description string `json:"description,omitempty"`

	// domain
	Domain string `json:"domain,omitempty"`

	// email address
	EmailAddress string `json:"emailAddress,omitempty"`

	// fqdn
	Fqdn string `json:"fqdn,omitempty"`

	// group Id
	GroupID string `json:"groupId,omitempty"`

	// group instance list
	GroupInstanceList []*GroupInstance `json:"groupInstanceList"`

	// group type
	// Enum: [CUSTOM SSO]
	GroupType string `json:"groupType,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// netbios
	Netbios string `json:"netbios,omitempty"`

	// parent groups
	ParentGroups []*GroupInstance `json:"parentGroups"`

	// principal Id
	PrincipalID *PrincipalID `json:"principalId,omitempty"`

	// principal name
	PrincipalName string `json:"principalName,omitempty"`

	// user instance list
	UserInstanceList []*UserInstance `json:"userInstanceList"`
}

// Validate validates this local group
func (m *LocalGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupInstanceList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrincipalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserInstanceList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalGroup) validateGroupInstanceList(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupInstanceList) { // not required
		return nil
	}

	for i := 0; i < len(m.GroupInstanceList); i++ {
		if swag.IsZero(m.GroupInstanceList[i]) { // not required
			continue
		}

		if m.GroupInstanceList[i] != nil {
			if err := m.GroupInstanceList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groupInstanceList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("groupInstanceList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var localGroupTypeGroupTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CUSTOM","SSO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		localGroupTypeGroupTypePropEnum = append(localGroupTypeGroupTypePropEnum, v)
	}
}

const (

	// LocalGroupGroupTypeCUSTOM captures enum value "CUSTOM"
	LocalGroupGroupTypeCUSTOM string = "CUSTOM"

	// LocalGroupGroupTypeSSO captures enum value "SSO"
	LocalGroupGroupTypeSSO string = "SSO"
)

// prop value enum
func (m *LocalGroup) validateGroupTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, localGroupTypeGroupTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LocalGroup) validateGroupType(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupType) { // not required
		return nil
	}

	// value enum
	if err := m.validateGroupTypeEnum("groupType", "body", m.GroupType); err != nil {
		return err
	}

	return nil
}

func (m *LocalGroup) validateParentGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.ParentGroups); i++ {
		if swag.IsZero(m.ParentGroups[i]) { // not required
			continue
		}

		if m.ParentGroups[i] != nil {
			if err := m.ParentGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parentGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parentGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LocalGroup) validatePrincipalID(formats strfmt.Registry) error {
	if swag.IsZero(m.PrincipalID) { // not required
		return nil
	}

	if m.PrincipalID != nil {
		if err := m.PrincipalID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("principalId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("principalId")
			}
			return err
		}
	}

	return nil
}

func (m *LocalGroup) validateUserInstanceList(formats strfmt.Registry) error {
	if swag.IsZero(m.UserInstanceList) { // not required
		return nil
	}

	for i := 0; i < len(m.UserInstanceList); i++ {
		if swag.IsZero(m.UserInstanceList[i]) { // not required
			continue
		}

		if m.UserInstanceList[i] != nil {
			if err := m.UserInstanceList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userInstanceList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userInstanceList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this local group based on the context it is used
func (m *LocalGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroupInstanceList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParentGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrincipalID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserInstanceList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalGroup) contextValidateGroupInstanceList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GroupInstanceList); i++ {

		if m.GroupInstanceList[i] != nil {
			if err := m.GroupInstanceList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groupInstanceList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("groupInstanceList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LocalGroup) contextValidateParentGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ParentGroups); i++ {

		if m.ParentGroups[i] != nil {
			if err := m.ParentGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parentGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parentGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LocalGroup) contextValidatePrincipalID(ctx context.Context, formats strfmt.Registry) error {

	if m.PrincipalID != nil {
		if err := m.PrincipalID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("principalId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("principalId")
			}
			return err
		}
	}

	return nil
}

func (m *LocalGroup) contextValidateUserInstanceList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UserInstanceList); i++ {

		if m.UserInstanceList[i] != nil {
			if err := m.UserInstanceList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userInstanceList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userInstanceList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LocalGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocalGroup) UnmarshalBinary(b []byte) error {
	var res LocalGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
