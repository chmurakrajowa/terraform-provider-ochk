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
	"github.com/go-openapi/validate"
)

// OldListBillingPeriodResponse old list billing period response
//
// swagger:model OldListBillingPeriodResponse
type OldListBillingPeriodResponse struct {

	// billing period collection
	BillingPeriodCollection []*BillingPeriodInstance `json:"billingPeriodCollection"`

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this old list billing period response
func (m *OldListBillingPeriodResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingPeriodCollection(formats); err != nil {
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

func (m *OldListBillingPeriodResponse) validateBillingPeriodCollection(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingPeriodCollection) { // not required
		return nil
	}

	for i := 0; i < len(m.BillingPeriodCollection); i++ {
		if swag.IsZero(m.BillingPeriodCollection[i]) { // not required
			continue
		}

		if m.BillingPeriodCollection[i] != nil {
			if err := m.BillingPeriodCollection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("billingPeriodCollection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("billingPeriodCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OldListBillingPeriodResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this old list billing period response based on the context it is used
func (m *OldListBillingPeriodResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBillingPeriodCollection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OldListBillingPeriodResponse) contextValidateBillingPeriodCollection(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BillingPeriodCollection); i++ {

		if m.BillingPeriodCollection[i] != nil {

			if swag.IsZero(m.BillingPeriodCollection[i]) { // not required
				return nil
			}

			if err := m.BillingPeriodCollection[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("billingPeriodCollection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("billingPeriodCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OldListBillingPeriodResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OldListBillingPeriodResponse) UnmarshalBinary(b []byte) error {
	var res OldListBillingPeriodResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
