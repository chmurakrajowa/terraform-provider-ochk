// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SubAttributeInstance SubAttributeInstance
//
// swagger:model SubAttributeInstance
type SubAttributeInstance struct {

	// attribute key
	AttributeKey string `json:"attributeKey,omitempty"`

	// data type
	DataType string `json:"dataType,omitempty"`

	// sub attribute Id
	SubAttributeID int32 `json:"subAttributeId,omitempty"`

	// sub attribute value list
	SubAttributeValueList []*SubAttributeValueInstance `json:"subAttributeValueList"`
}

// Validate validates this sub attribute instance
func (m *SubAttributeInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubAttributeValueList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubAttributeInstance) validateSubAttributeValueList(formats strfmt.Registry) error {

	if swag.IsZero(m.SubAttributeValueList) { // not required
		return nil
	}

	for i := 0; i < len(m.SubAttributeValueList); i++ {
		if swag.IsZero(m.SubAttributeValueList[i]) { // not required
			continue
		}

		if m.SubAttributeValueList[i] != nil {
			if err := m.SubAttributeValueList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subAttributeValueList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubAttributeInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubAttributeInstance) UnmarshalBinary(b []byte) error {
	var res SubAttributeInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
