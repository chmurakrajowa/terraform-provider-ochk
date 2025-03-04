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

// PowerState power state
//
// swagger:model PowerState
type PowerState string

func NewPowerState(value PowerState) *PowerState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated PowerState.
func (m PowerState) Pointer() *PowerState {
	return &m
}

const (

	// PowerStatePoweredOff captures enum value "poweredOff"
	PowerStatePoweredOff PowerState = "poweredOff"

	// PowerStatePoweredOn captures enum value "poweredOn"
	PowerStatePoweredOn PowerState = "poweredOn"

	// PowerStateSuspended captures enum value "suspended"
	PowerStateSuspended PowerState = "suspended"
)

// for schema
var powerStateEnum []interface{}

func init() {
	var res []PowerState
	if err := json.Unmarshal([]byte(`["poweredOff","poweredOn","suspended"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerStateEnum = append(powerStateEnum, v)
	}
}

func (m PowerState) validatePowerStateEnum(path, location string, value PowerState) error {
	if err := validate.EnumCase(path, location, value, powerStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this power state
func (m PowerState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePowerStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this power state based on context it is used
func (m PowerState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
