// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeploymentInstance DeploymentInstance
//
// swagger:model DeploymentInstance
type DeploymentInstance struct {

	// byol
	Byol bool `json:"byol,omitempty"`

	// custom specification Id
	CustomSpecificationID string `json:"customSpecificationId,omitempty"`

	// deployment category
	// Enum: [LINUX WINDOWS]
	DeploymentCategory string `json:"deploymentCategory,omitempty"`

	// deployment Id
	DeploymentID string `json:"deploymentId,omitempty"`

	// deployment initial size g b
	DeploymentInitialSizeGB float32 `json:"deploymentInitialSizeGB,omitempty"`

	// deployment initial size m b
	DeploymentInitialSizeMB float32 `json:"deploymentInitialSizeMB,omitempty"`

	// deployment template pattern
	DeploymentTemplatePattern string `json:"deploymentTemplatePattern,omitempty"`

	// deployment type
	// Enum: [ISO OVF TEMPLATE]
	DeploymentType string `json:"deploymentType,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// iso name
	IsoName string `json:"isoName,omitempty"`

	// ovf name
	OvfName string `json:"ovfName,omitempty"`
}

// Validate validates this deployment instance
func (m *DeploymentInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeploymentCategory(formats); err != nil {
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

var deploymentInstanceTypeDeploymentCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LINUX","WINDOWS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deploymentInstanceTypeDeploymentCategoryPropEnum = append(deploymentInstanceTypeDeploymentCategoryPropEnum, v)
	}
}

const (

	// DeploymentInstanceDeploymentCategoryLINUX captures enum value "LINUX"
	DeploymentInstanceDeploymentCategoryLINUX string = "LINUX"

	// DeploymentInstanceDeploymentCategoryWINDOWS captures enum value "WINDOWS"
	DeploymentInstanceDeploymentCategoryWINDOWS string = "WINDOWS"
)

// prop value enum
func (m *DeploymentInstance) validateDeploymentCategoryEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deploymentInstanceTypeDeploymentCategoryPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeploymentInstance) validateDeploymentCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentCategory) { // not required
		return nil
	}

	// value enum
	if err := m.validateDeploymentCategoryEnum("deploymentCategory", "body", m.DeploymentCategory); err != nil {
		return err
	}

	return nil
}

var deploymentInstanceTypeDeploymentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ISO","OVF","TEMPLATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deploymentInstanceTypeDeploymentTypePropEnum = append(deploymentInstanceTypeDeploymentTypePropEnum, v)
	}
}

const (

	// DeploymentInstanceDeploymentTypeISO captures enum value "ISO"
	DeploymentInstanceDeploymentTypeISO string = "ISO"

	// DeploymentInstanceDeploymentTypeOVF captures enum value "OVF"
	DeploymentInstanceDeploymentTypeOVF string = "OVF"

	// DeploymentInstanceDeploymentTypeTEMPLATE captures enum value "TEMPLATE"
	DeploymentInstanceDeploymentTypeTEMPLATE string = "TEMPLATE"
)

// prop value enum
func (m *DeploymentInstance) validateDeploymentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deploymentInstanceTypeDeploymentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeploymentInstance) validateDeploymentType(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDeploymentTypeEnum("deploymentType", "body", m.DeploymentType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this deployment instance based on context it is used
func (m *DeploymentInstance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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
