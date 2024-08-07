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

// DeploymentGetResponse DeploymentGetResponse
//
// swagger:model DeploymentGetResponse
type DeploymentGetResponse struct {

	// deployment instance
	DeploymentInstance *DeploymentInstance `json:"deploymentInstance,omitempty"`

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this deployment get response
func (m *DeploymentGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeploymentInstance(formats); err != nil {
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

func (m *DeploymentGetResponse) validateDeploymentInstance(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentInstance) { // not required
		return nil
	}

	if m.DeploymentInstance != nil {
		if err := m.DeploymentInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deploymentInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deploymentInstance")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentGetResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this deployment get response based on the context it is used
func (m *DeploymentGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeploymentInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentGetResponse) contextValidateDeploymentInstance(ctx context.Context, formats strfmt.Registry) error {

	if m.DeploymentInstance != nil {

		if swag.IsZero(m.DeploymentInstance) { // not required
			return nil
		}

		if err := m.DeploymentInstance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deploymentInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deploymentInstance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentGetResponse) UnmarshalBinary(b []byte) error {
	var res DeploymentGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
