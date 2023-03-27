// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IPCollection IpCollection
//
// swagger:model IpCollection
type IPCollection struct {

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// creation date
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// ip collection addresses
	IPCollectionAddresses []string `json:"ipCollectionAddresses"`

	// member type
	// Enum: [IPCOLLECTION SEGMENT VIRTUAL_MACHINE]
	MemberType string `json:"memberType,omitempty"`

	// modification date
	// Format: date-time
	ModificationDate *strfmt.DateTime `json:"modificationDate,omitempty"`

	// modified by
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`
}

// Validate validates this Ip collection
func (m *IPCollection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModificationDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPCollection) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var ipCollectionTypeMemberTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IPCOLLECTION","SEGMENT","VIRTUAL_MACHINE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipCollectionTypeMemberTypePropEnum = append(ipCollectionTypeMemberTypePropEnum, v)
	}
}

const (

	// IPCollectionMemberTypeIPCOLLECTION captures enum value "IPCOLLECTION"
	IPCollectionMemberTypeIPCOLLECTION string = "IPCOLLECTION"

	// IPCollectionMemberTypeSEGMENT captures enum value "SEGMENT"
	IPCollectionMemberTypeSEGMENT string = "SEGMENT"

	// IPCollectionMemberTypeVIRTUALMACHINE captures enum value "VIRTUAL_MACHINE"
	IPCollectionMemberTypeVIRTUALMACHINE string = "VIRTUAL_MACHINE"
)

// prop value enum
func (m *IPCollection) validateMemberTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ipCollectionTypeMemberTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IPCollection) validateMemberType(formats strfmt.Registry) error {
	if swag.IsZero(m.MemberType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMemberTypeEnum("memberType", "body", m.MemberType); err != nil {
		return err
	}

	return nil
}

func (m *IPCollection) validateModificationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ModificationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modificationDate", "body", "date-time", m.ModificationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Ip collection based on context it is used
func (m *IPCollection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IPCollection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPCollection) UnmarshalBinary(b []byte) error {
	var res IPCollection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
