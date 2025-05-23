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

// ErrorType error type
//
// swagger:model ErrorType
type ErrorType int32

// for schema
var errorTypeEnum []interface{}

func init() {
	var res []ErrorType
	if err := json.Unmarshal([]byte(`[0,1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		errorTypeEnum = append(errorTypeEnum, v)
	}
}

func (m ErrorType) validateErrorTypeEnum(path, location string, value ErrorType) error {
	if err := validate.EnumCase(path, location, value, errorTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this error type
func (m ErrorType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateErrorTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this error type based on context it is used
func (m ErrorType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
