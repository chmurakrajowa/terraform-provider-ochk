// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BackupPlanGetResponse BackupPlanGetResponse
//
// swagger:model BackupPlanGetResponse
type BackupPlanGetResponse struct {

	// backup plan
	BackupPlan *BackupPlan `json:"backupPlan,omitempty"`

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this backup plan get response
func (m *BackupPlanGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupPlan(formats); err != nil {
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

func (m *BackupPlanGetResponse) validateBackupPlan(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupPlan) { // not required
		return nil
	}

	if m.BackupPlan != nil {
		if err := m.BackupPlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backupPlan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backupPlan")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanGetResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this backup plan get response based on the context it is used
func (m *BackupPlanGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupPlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupPlanGetResponse) contextValidateBackupPlan(ctx context.Context, formats strfmt.Registry) error {

	if m.BackupPlan != nil {
		if err := m.BackupPlan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backupPlan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backupPlan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupPlanGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupPlanGetResponse) UnmarshalBinary(b []byte) error {
	var res BackupPlanGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
