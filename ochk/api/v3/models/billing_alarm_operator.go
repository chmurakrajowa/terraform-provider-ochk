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

// BillingAlarmOperator billing alarm operator
//
// swagger:model BillingAlarmOperator
type BillingAlarmOperator string

func NewBillingAlarmOperator(value BillingAlarmOperator) *BillingAlarmOperator {
	return &value
}

// Pointer returns a pointer to a freshly-allocated BillingAlarmOperator.
func (m BillingAlarmOperator) Pointer() *BillingAlarmOperator {
	return &m
}

const (

	// BillingAlarmOperatorHIGHER captures enum value "HIGHER"
	BillingAlarmOperatorHIGHER BillingAlarmOperator = "HIGHER"

	// BillingAlarmOperatorLOWER captures enum value "LOWER"
	BillingAlarmOperatorLOWER BillingAlarmOperator = "LOWER"

	// BillingAlarmOperatorEQUALS captures enum value "EQUALS"
	BillingAlarmOperatorEQUALS BillingAlarmOperator = "EQUALS"
)

// for schema
var billingAlarmOperatorEnum []interface{}

func init() {
	var res []BillingAlarmOperator
	if err := json.Unmarshal([]byte(`["HIGHER","LOWER","EQUALS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billingAlarmOperatorEnum = append(billingAlarmOperatorEnum, v)
	}
}

func (m BillingAlarmOperator) validateBillingAlarmOperatorEnum(path, location string, value BillingAlarmOperator) error {
	if err := validate.EnumCase(path, location, value, billingAlarmOperatorEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this billing alarm operator
func (m BillingAlarmOperator) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBillingAlarmOperatorEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this billing alarm operator based on context it is used
func (m BillingAlarmOperator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
