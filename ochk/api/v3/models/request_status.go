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

// RequestStatus request status
//
// swagger:model RequestStatus
type RequestStatus string

func NewRequestStatus(value RequestStatus) *RequestStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated RequestStatus.
func (m RequestStatus) Pointer() *RequestStatus {
	return &m
}

const (

	// RequestStatusSUCCESS captures enum value "SUCCESS"
	RequestStatusSUCCESS RequestStatus = "SUCCESS"

	// RequestStatusFAILED captures enum value "FAILED"
	RequestStatusFAILED RequestStatus = "FAILED"
)

// for schema
var requestStatusEnum []interface{}

func init() {
	var res []RequestStatus
	if err := json.Unmarshal([]byte(`["SUCCESS","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		requestStatusEnum = append(requestStatusEnum, v)
	}
}

func (m RequestStatus) validateRequestStatusEnum(path, location string, value RequestStatus) error {
	if err := validate.EnumCase(path, location, value, requestStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this request status
func (m RequestStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRequestStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this request status based on context it is used
func (m RequestStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
