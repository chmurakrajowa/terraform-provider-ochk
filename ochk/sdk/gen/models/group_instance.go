// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GroupInstance GroupInstance
//
// swagger:model GroupInstance
type GroupInstance struct {

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

	// principal Id
	PrincipalID *PrincipalID `json:"principalId,omitempty"`

	// principal name
	PrincipalName string `json:"principalName,omitempty"`

	// user instance list
	UserInstanceList []*UserInstance `json:"userInstanceList"`
}

// Validate validates this group instance
func (m *GroupInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupInstanceList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupType(formats); err != nil {
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

func (m *GroupInstance) validateGroupInstanceList(formats strfmt.Registry) error {

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
				}
				return err
			}
		}

	}

	return nil
}

var groupInstanceTypeGroupTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CUSTOM","SSO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupInstanceTypeGroupTypePropEnum = append(groupInstanceTypeGroupTypePropEnum, v)
	}
}

const (

	// GroupInstanceGroupTypeCUSTOM captures enum value "CUSTOM"
	GroupInstanceGroupTypeCUSTOM string = "CUSTOM"

	// GroupInstanceGroupTypeSSO captures enum value "SSO"
	GroupInstanceGroupTypeSSO string = "SSO"
)

// prop value enum
func (m *GroupInstance) validateGroupTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, groupInstanceTypeGroupTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GroupInstance) validateGroupType(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupType) { // not required
		return nil
	}

	// value enum
	if err := m.validateGroupTypeEnum("groupType", "body", m.GroupType); err != nil {
		return err
	}

	return nil
}

func (m *GroupInstance) validatePrincipalID(formats strfmt.Registry) error {

	if swag.IsZero(m.PrincipalID) { // not required
		return nil
	}

	if m.PrincipalID != nil {
		if err := m.PrincipalID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("principalId")
			}
			return err
		}
	}

	return nil
}

func (m *GroupInstance) validateUserInstanceList(formats strfmt.Registry) error {

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
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupInstance) UnmarshalBinary(b []byte) error {
	var res GroupInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
