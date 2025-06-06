// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NATType n a t type
//
// swagger:model NATType
type NATType string

func NewNATType(value NATType) *NATType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated NATType.
func (m NATType) Pointer() *NATType {
	return &m
}

const (

	// NATTypeMANUAL captures enum value "MANUAL"
	NATTypeMANUAL NATType = "MANUAL"

	// NATTypeAUTO captures enum value "AUTO"
	NATTypeAUTO NATType = "AUTO"
)

// for schema
var nATTypeEnum []interface{}

func init() {
	var res []NATType
	if err := json.Unmarshal([]byte(`["MANUAL","AUTO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nATTypeEnum = append(nATTypeEnum, v)
	}
}

func (m NATType) validateNATTypeEnum(path, location string, value NATType) error {
	if err := validate.EnumCase(path, location, value, nATTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this n a t type
func (m NATType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateNATTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this n a t type based on context it is used
func (m NATType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
