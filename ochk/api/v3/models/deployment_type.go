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

// DeploymentType deployment type
//
// swagger:model DeploymentType
type DeploymentType string

func NewDeploymentType(value DeploymentType) *DeploymentType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated DeploymentType.
func (m DeploymentType) Pointer() *DeploymentType {
	return &m
}

const (

	// DeploymentTypeTEMPLATE captures enum value "TEMPLATE"
	DeploymentTypeTEMPLATE DeploymentType = "TEMPLATE"

	// DeploymentTypeOVF captures enum value "OVF"
	DeploymentTypeOVF DeploymentType = "OVF"

	// DeploymentTypeISO captures enum value "ISO"
	DeploymentTypeISO DeploymentType = "ISO"
)

// for schema
var deploymentTypeEnum []interface{}

func init() {
	var res []DeploymentType
	if err := json.Unmarshal([]byte(`["TEMPLATE","OVF","ISO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deploymentTypeEnum = append(deploymentTypeEnum, v)
	}
}

func (m DeploymentType) validateDeploymentTypeEnum(path, location string, value DeploymentType) error {
	if err := validate.EnumCase(path, location, value, deploymentTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this deployment type
func (m DeploymentType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDeploymentTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this deployment type based on context it is used
func (m DeploymentType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
