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
)

// AttributeInstance attribute instance
//
// swagger:model AttributeInstance
type AttributeInstance struct {

	// attribute Id
	AttributeID int32 `json:"attributeId,omitempty"`

	// attribute key
	AttributeKey string `json:"attributeKey,omitempty"`

	// attribute value list
	AttributeValueList []*AttributeValueInstance `json:"attributeValueList"`

	// data type
	DataType string `json:"dataType,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// is a l g type
	IsALGType bool `json:"isALGType,omitempty"`

	// sub attribute list
	SubAttributeList []*SubAttributeInstance `json:"subAttributeList"`
}

// Validate validates this attribute instance
func (m *AttributeInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeValueList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubAttributeList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AttributeInstance) validateAttributeValueList(formats strfmt.Registry) error {
	if swag.IsZero(m.AttributeValueList) { // not required
		return nil
	}

	for i := 0; i < len(m.AttributeValueList); i++ {
		if swag.IsZero(m.AttributeValueList[i]) { // not required
			continue
		}

		if m.AttributeValueList[i] != nil {
			if err := m.AttributeValueList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributeValueList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributeValueList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AttributeInstance) validateSubAttributeList(formats strfmt.Registry) error {
	if swag.IsZero(m.SubAttributeList) { // not required
		return nil
	}

	for i := 0; i < len(m.SubAttributeList); i++ {
		if swag.IsZero(m.SubAttributeList[i]) { // not required
			continue
		}

		if m.SubAttributeList[i] != nil {
			if err := m.SubAttributeList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subAttributeList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subAttributeList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this attribute instance based on the context it is used
func (m *AttributeInstance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributeValueList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubAttributeList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AttributeInstance) contextValidateAttributeValueList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AttributeValueList); i++ {

		if m.AttributeValueList[i] != nil {

			if swag.IsZero(m.AttributeValueList[i]) { // not required
				return nil
			}

			if err := m.AttributeValueList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributeValueList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributeValueList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AttributeInstance) contextValidateSubAttributeList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SubAttributeList); i++ {

		if m.SubAttributeList[i] != nil {

			if swag.IsZero(m.SubAttributeList[i]) { // not required
				return nil
			}

			if err := m.SubAttributeList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subAttributeList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subAttributeList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AttributeInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttributeInstance) UnmarshalBinary(b []byte) error {
	var res AttributeInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
