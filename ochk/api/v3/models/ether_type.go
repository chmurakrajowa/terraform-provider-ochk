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

// EtherType ether type
//
// swagger:model EtherType
type EtherType string

func NewEtherType(value EtherType) *EtherType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated EtherType.
func (m EtherType) Pointer() *EtherType {
	return &m
}

const (

	// EtherTypeIPV4 captures enum value "IPv4"
	EtherTypeIPV4 EtherType = "IPv4"

	// EtherTypeIPV6 captures enum value "IPv6"
	EtherTypeIPV6 EtherType = "IPv6"
)

// for schema
var etherTypeEnum []interface{}

func init() {
	var res []EtherType
	if err := json.Unmarshal([]byte(`["IPv4","IPv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		etherTypeEnum = append(etherTypeEnum, v)
	}
}

func (m EtherType) validateEtherTypeEnum(path, location string, value EtherType) error {
	if err := validate.EnumCase(path, location, value, etherTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this ether type
func (m EtherType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEtherTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this ether type based on context it is used
func (m EtherType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
