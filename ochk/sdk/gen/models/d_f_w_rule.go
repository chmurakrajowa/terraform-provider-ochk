// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DFWRule DFWRule
//
// swagger:model DFWRule
type DFWRule struct {

	// action
	// Enum: [ALLOW DROP REJECT]
	Action string `json:"action,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// creation date
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate,omitempty"`

	// custom services
	CustomServices []*CustomServiceInstance `json:"customServices"`

	// default services
	DefaultServices []*ServiceInstance `json:"defaultServices"`

	// destination
	Destination []*SecurityGroup `json:"destination"`

	// direction
	// Enum: [IN IN_OUT OUT]
	Direction string `json:"direction,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// ip protocol
	// Enum: [IPV4 IPV4_IPV6 IPV6]
	IPProtocol string `json:"ipProtocol,omitempty"`

	// modification date
	// Format: date-time
	ModificationDate *strfmt.DateTime `json:"modificationDate,omitempty"`

	// modified by
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// position
	Position *Position `json:"position,omitempty"`

	// priority
	Priority int64 `json:"priority,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// rule Id
	RuleID string `json:"ruleId,omitempty"`

	// scope
	Scope []*SecurityGroup `json:"scope"`

	// source
	Source []*SecurityGroup `json:"source"`
}

// Validate validates this d f w rule
func (m *DFWRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModificationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var dFWRuleTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALLOW","DROP","REJECT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dFWRuleTypeActionPropEnum = append(dFWRuleTypeActionPropEnum, v)
	}
}

const (

	// DFWRuleActionALLOW captures enum value "ALLOW"
	DFWRuleActionALLOW string = "ALLOW"

	// DFWRuleActionDROP captures enum value "DROP"
	DFWRuleActionDROP string = "DROP"

	// DFWRuleActionREJECT captures enum value "REJECT"
	DFWRuleActionREJECT string = "REJECT"
)

// prop value enum
func (m *DFWRule) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dFWRuleTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DFWRule) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *DFWRule) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DFWRule) validateCustomServices(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomServices) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomServices); i++ {
		if swag.IsZero(m.CustomServices[i]) { // not required
			continue
		}

		if m.CustomServices[i] != nil {
			if err := m.CustomServices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customServices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customServices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DFWRule) validateDefaultServices(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultServices) { // not required
		return nil
	}

	for i := 0; i < len(m.DefaultServices); i++ {
		if swag.IsZero(m.DefaultServices[i]) { // not required
			continue
		}

		if m.DefaultServices[i] != nil {
			if err := m.DefaultServices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("defaultServices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("defaultServices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DFWRule) validateDestination(formats strfmt.Registry) error {
	if swag.IsZero(m.Destination) { // not required
		return nil
	}

	for i := 0; i < len(m.Destination); i++ {
		if swag.IsZero(m.Destination[i]) { // not required
			continue
		}

		if m.Destination[i] != nil {
			if err := m.Destination[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("destination" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("destination" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var dFWRuleTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IN","IN_OUT","OUT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dFWRuleTypeDirectionPropEnum = append(dFWRuleTypeDirectionPropEnum, v)
	}
}

const (

	// DFWRuleDirectionIN captures enum value "IN"
	DFWRuleDirectionIN string = "IN"

	// DFWRuleDirectionINOUT captures enum value "IN_OUT"
	DFWRuleDirectionINOUT string = "IN_OUT"

	// DFWRuleDirectionOUT captures enum value "OUT"
	DFWRuleDirectionOUT string = "OUT"
)

// prop value enum
func (m *DFWRule) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dFWRuleTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DFWRule) validateDirection(formats strfmt.Registry) error {
	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

var dFWRuleTypeIPProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IPV4","IPV4_IPV6","IPV6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dFWRuleTypeIPProtocolPropEnum = append(dFWRuleTypeIPProtocolPropEnum, v)
	}
}

const (

	// DFWRuleIPProtocolIPV4 captures enum value "IPV4"
	DFWRuleIPProtocolIPV4 string = "IPV4"

	// DFWRuleIPProtocolIPV4IPV6 captures enum value "IPV4_IPV6"
	DFWRuleIPProtocolIPV4IPV6 string = "IPV4_IPV6"

	// DFWRuleIPProtocolIPV6 captures enum value "IPV6"
	DFWRuleIPProtocolIPV6 string = "IPV6"
)

// prop value enum
func (m *DFWRule) validateIPProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dFWRuleTypeIPProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DFWRule) validateIPProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.IPProtocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateIPProtocolEnum("ipProtocol", "body", m.IPProtocol); err != nil {
		return err
	}

	return nil
}

func (m *DFWRule) validateModificationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ModificationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modificationDate", "body", "date-time", m.ModificationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DFWRule) validatePosition(formats strfmt.Registry) error {
	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if m.Position != nil {
		if err := m.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

func (m *DFWRule) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	for i := 0; i < len(m.Scope); i++ {
		if swag.IsZero(m.Scope[i]) { // not required
			continue
		}

		if m.Scope[i] != nil {
			if err := m.Scope[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scope" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("scope" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DFWRule) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	for i := 0; i < len(m.Source); i++ {
		if swag.IsZero(m.Source[i]) { // not required
			continue
		}

		if m.Source[i] != nil {
			if err := m.Source[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("source" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("source" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this d f w rule based on the context it is used
func (m *DFWRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDestination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DFWRule) contextValidateCustomServices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomServices); i++ {

		if m.CustomServices[i] != nil {

			if swag.IsZero(m.CustomServices[i]) { // not required
				return nil
			}

			if err := m.CustomServices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customServices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customServices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DFWRule) contextValidateDefaultServices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DefaultServices); i++ {

		if m.DefaultServices[i] != nil {

			if swag.IsZero(m.DefaultServices[i]) { // not required
				return nil
			}

			if err := m.DefaultServices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("defaultServices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("defaultServices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DFWRule) contextValidateDestination(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Destination); i++ {

		if m.Destination[i] != nil {

			if swag.IsZero(m.Destination[i]) { // not required
				return nil
			}

			if err := m.Destination[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("destination" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("destination" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DFWRule) contextValidatePosition(ctx context.Context, formats strfmt.Registry) error {

	if m.Position != nil {

		if swag.IsZero(m.Position) { // not required
			return nil
		}

		if err := m.Position.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

func (m *DFWRule) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Scope); i++ {

		if m.Scope[i] != nil {

			if swag.IsZero(m.Scope[i]) { // not required
				return nil
			}

			if err := m.Scope[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scope" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("scope" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DFWRule) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Source); i++ {

		if m.Source[i] != nil {

			if swag.IsZero(m.Source[i]) { // not required
				return nil
			}

			if err := m.Source[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("source" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("source" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DFWRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DFWRule) UnmarshalBinary(b []byte) error {
	var res DFWRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
