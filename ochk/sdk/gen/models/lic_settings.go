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

// LicSettings LicSettings
//
// swagger:model LicSettings
type LicSettings struct {

	// byol enabled
	ByolEnabled bool `json:"byolEnabled,omitempty"`

	// os type
	// Enum: [LINUX WINDOWS]
	OsType string `json:"osType,omitempty"`

	// use dedicated hosts
	UseDedicatedHosts bool `json:"useDedicatedHosts,omitempty"`
}

// Validate validates this lic settings
func (m *LicSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOsType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var licSettingsTypeOsTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LINUX","WINDOWS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		licSettingsTypeOsTypePropEnum = append(licSettingsTypeOsTypePropEnum, v)
	}
}

const (

	// LicSettingsOsTypeLINUX captures enum value "LINUX"
	LicSettingsOsTypeLINUX string = "LINUX"

	// LicSettingsOsTypeWINDOWS captures enum value "WINDOWS"
	LicSettingsOsTypeWINDOWS string = "WINDOWS"
)

// prop value enum
func (m *LicSettings) validateOsTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, licSettingsTypeOsTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LicSettings) validateOsType(formats strfmt.Registry) error {
	if swag.IsZero(m.OsType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOsTypeEnum("osType", "body", m.OsType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this lic settings based on context it is used
func (m *LicSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LicSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicSettings) UnmarshalBinary(b []byte) error {
	var res LicSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
