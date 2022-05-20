// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NATRuleListResponse NATRuleListResponse
//
// swagger:model NATRuleListResponse
type NATRuleListResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// nat rule instances
	NatRuleInstances []*NATRuleInstance `json:"natRuleInstances"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this n a t rule list response
func (m *NATRuleListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNatRuleInstances(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NATRuleListResponse) validateNatRuleInstances(formats strfmt.Registry) error {

	if swag.IsZero(m.NatRuleInstances) { // not required
		return nil
	}

	for i := 0; i < len(m.NatRuleInstances); i++ {
		if swag.IsZero(m.NatRuleInstances[i]) { // not required
			continue
		}

		if m.NatRuleInstances[i] != nil {
			if err := m.NatRuleInstances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("natRuleInstances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NATRuleListResponse) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NATRuleListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NATRuleListResponse) UnmarshalBinary(b []byte) error {
	var res NATRuleListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
