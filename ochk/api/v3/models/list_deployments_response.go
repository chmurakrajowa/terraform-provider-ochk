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

// ListDeploymentsResponse list deployments response
//
// swagger:model ListDeploymentsResponse
type ListDeploymentsResponse struct {

	// deployment instance collection
	DeploymentInstanceCollection []*DeploymentInstance `json:"deploymentInstanceCollection"`

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this list deployments response
func (m *ListDeploymentsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeploymentInstanceCollection(formats); err != nil {
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

func (m *ListDeploymentsResponse) validateDeploymentInstanceCollection(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentInstanceCollection) { // not required
		return nil
	}

	for i := 0; i < len(m.DeploymentInstanceCollection); i++ {
		if swag.IsZero(m.DeploymentInstanceCollection[i]) { // not required
			continue
		}

		if m.DeploymentInstanceCollection[i] != nil {
			if err := m.DeploymentInstanceCollection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deploymentInstanceCollection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deploymentInstanceCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListDeploymentsResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this list deployments response based on the context it is used
func (m *ListDeploymentsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeploymentInstanceCollection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListDeploymentsResponse) contextValidateDeploymentInstanceCollection(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DeploymentInstanceCollection); i++ {

		if m.DeploymentInstanceCollection[i] != nil {

			if swag.IsZero(m.DeploymentInstanceCollection[i]) { // not required
				return nil
			}

			if err := m.DeploymentInstanceCollection[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deploymentInstanceCollection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deploymentInstanceCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListDeploymentsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListDeploymentsResponse) UnmarshalBinary(b []byte) error {
	var res ListDeploymentsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
