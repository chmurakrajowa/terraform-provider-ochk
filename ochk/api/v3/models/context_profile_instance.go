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

// ContextProfileInstance context profile instance
//
// swagger:model ContextProfileInstance
type ContextProfileInstance struct {

	// attribute list
	AttributeList []*AttributeInstance `json:"attributeList"`

	// context profile Id
	// Format: uuid
	ContextProfileID strfmt.UUID `json:"contextProfileId,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// external Id
	ExternalID string `json:"externalId,omitempty"`

	// parent path
	ParentPath string `json:"parentPath,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// protection
	Protection string `json:"protection,omitempty"`

	// relative path
	RelativePath string `json:"relativePath,omitempty"`

	// resource type
	ResourceType *ResourceType `json:"resourceType,omitempty"`

	// revision
	Revision int64 `json:"revision,omitempty"`

	// schema value
	SchemaValue string `json:"schemaValue,omitempty"`

	// system owned
	SystemOwned bool `json:"systemOwned,omitempty"`
}

// Validate validates this context profile instance
func (m *ContextProfileInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContextProfileID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContextProfileInstance) validateAttributeList(formats strfmt.Registry) error {
	if swag.IsZero(m.AttributeList) { // not required
		return nil
	}

	for i := 0; i < len(m.AttributeList); i++ {
		if swag.IsZero(m.AttributeList[i]) { // not required
			continue
		}

		if m.AttributeList[i] != nil {
			if err := m.AttributeList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributeList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributeList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContextProfileInstance) validateContextProfileID(formats strfmt.Registry) error {
	if swag.IsZero(m.ContextProfileID) { // not required
		return nil
	}

	if err := validate.FormatOf("contextProfileId", "body", "uuid", m.ContextProfileID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContextProfileInstance) validateResourceType(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceType) { // not required
		return nil
	}

	if m.ResourceType != nil {
		if err := m.ResourceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this context profile instance based on the context it is used
func (m *ContextProfileInstance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributeList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContextProfileInstance) contextValidateAttributeList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AttributeList); i++ {

		if m.AttributeList[i] != nil {

			if swag.IsZero(m.AttributeList[i]) { // not required
				return nil
			}

			if err := m.AttributeList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributeList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributeList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContextProfileInstance) contextValidateResourceType(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceType != nil {

		if swag.IsZero(m.ResourceType) { // not required
			return nil
		}

		if err := m.ResourceType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContextProfileInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContextProfileInstance) UnmarshalBinary(b []byte) error {
	var res ContextProfileInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
