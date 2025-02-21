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

// DeploymentInstance deployment instance
//
// swagger:model DeploymentInstance
type DeploymentInstance struct {

	// byol
	Byol bool `json:"byol,omitempty"`

	// compute site
	ComputeSite *ComputeSiteInstance `json:"computeSite,omitempty"`

	// custom command instance list
	CustomCommandInstanceList []*CustomCommandInstance `json:"customCommandInstanceList"`

	// custom specification Id
	// Format: uuid
	CustomSpecificationID strfmt.UUID `json:"customSpecificationId,omitempty"`

	// deployment category
	DeploymentCategory DeploymentCategory `json:"deploymentCategory,omitempty"`

	// deployment Id
	// Format: uuid
	DeploymentID strfmt.UUID `json:"deploymentId,omitempty"`

	// deployment initial size g b
	DeploymentInitialSizeGB float32 `json:"deploymentInitialSizeGB,omitempty"`

	// deployment template pattern
	DeploymentTemplatePattern string `json:"deploymentTemplatePattern,omitempty"`

	// deployment type
	DeploymentType DeploymentType `json:"deploymentType,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// iso name
	IsoName string `json:"isoName,omitempty"`

	// osc image external Id
	OscImageExternalID string `json:"oscImageExternalId,omitempty"`

	// ovf name
	OvfName string `json:"ovfName,omitempty"`
}

// Validate validates this deployment instance
func (m *DeploymentInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComputeSite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomCommandInstanceList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomSpecificationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentInstance) validateComputeSite(formats strfmt.Registry) error {
	if swag.IsZero(m.ComputeSite) { // not required
		return nil
	}

	if m.ComputeSite != nil {
		if err := m.ComputeSite.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("computeSite")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("computeSite")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentInstance) validateCustomCommandInstanceList(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomCommandInstanceList) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomCommandInstanceList); i++ {
		if swag.IsZero(m.CustomCommandInstanceList[i]) { // not required
			continue
		}

		if m.CustomCommandInstanceList[i] != nil {
			if err := m.CustomCommandInstanceList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customCommandInstanceList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customCommandInstanceList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeploymentInstance) validateCustomSpecificationID(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomSpecificationID) { // not required
		return nil
	}

	if err := validate.FormatOf("customSpecificationId", "body", "uuid", m.CustomSpecificationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentInstance) validateDeploymentCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentCategory) { // not required
		return nil
	}

	if err := m.DeploymentCategory.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deploymentCategory")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("deploymentCategory")
		}
		return err
	}

	return nil
}

func (m *DeploymentInstance) validateDeploymentID(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentID) { // not required
		return nil
	}

	if err := validate.FormatOf("deploymentId", "body", "uuid", m.DeploymentID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentInstance) validateDeploymentType(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentType) { // not required
		return nil
	}

	if err := m.DeploymentType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deploymentType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("deploymentType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this deployment instance based on the context it is used
func (m *DeploymentInstance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComputeSite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomCommandInstanceList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeploymentCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeploymentType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentInstance) contextValidateComputeSite(ctx context.Context, formats strfmt.Registry) error {

	if m.ComputeSite != nil {

		if swag.IsZero(m.ComputeSite) { // not required
			return nil
		}

		if err := m.ComputeSite.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("computeSite")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("computeSite")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentInstance) contextValidateCustomCommandInstanceList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomCommandInstanceList); i++ {

		if m.CustomCommandInstanceList[i] != nil {

			if swag.IsZero(m.CustomCommandInstanceList[i]) { // not required
				return nil
			}

			if err := m.CustomCommandInstanceList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customCommandInstanceList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customCommandInstanceList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeploymentInstance) contextValidateDeploymentCategory(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.DeploymentCategory) { // not required
		return nil
	}

	if err := m.DeploymentCategory.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deploymentCategory")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("deploymentCategory")
		}
		return err
	}

	return nil
}

func (m *DeploymentInstance) contextValidateDeploymentType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.DeploymentType) { // not required
		return nil
	}

	if err := m.DeploymentType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deploymentType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("deploymentType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentInstance) UnmarshalBinary(b []byte) error {
	var res DeploymentInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
