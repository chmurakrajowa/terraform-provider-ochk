// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NATRuleInstance NATRuleInstance
//
// swagger:model NATRuleInstance
type NATRuleInstance struct {

	// action
	// Enum: [DNAT NO_SNAT SNAT]
	Action string `json:"action,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// creation date
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// destination network
	DestinationNetwork string `json:"destinationNetwork,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// firewall match
	// Enum: [BYPASS MATCH_EXTERNAL_ADDRESS MATCH_INTERNAL_ADDRESS]
	FirewallMatch string `json:"firewallMatch,omitempty"`

	// modification date
	// Format: date-time
	ModificationDate *strfmt.DateTime `json:"modificationDate,omitempty"`

	// modified by
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// nat type
	// Enum: [AUTO MANUAL]
	NatType string `json:"natType,omitempty"`

	// priority
	Priority int64 `json:"priority,omitempty"`

	// rule Id
	RuleID string `json:"ruleId,omitempty"`

	// service
	Service string `json:"service,omitempty"`

	// service instance
	ServiceInstance *ServiceInstance `json:"serviceInstance,omitempty"`

	// source network
	SourceNetwork string `json:"sourceNetwork,omitempty"`

	// tier zero router
	TierZeroRouter *RouterInstance `json:"tierZeroRouter,omitempty"`

	// translated network
	TranslatedNetwork string `json:"translatedNetwork,omitempty"`

	// translated ports
	TranslatedPorts string `json:"translatedPorts,omitempty"`

	// virtual network instance
	VirtualNetworkInstance *VirtualNetworkInstance `json:"virtualNetworkInstance,omitempty"`
}

// Validate validates this n a t rule instance
func (m *NATRuleInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirewallMatch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModificationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNatType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceInstance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTierZeroRouter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualNetworkInstance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var nATRuleInstanceTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DNAT","NO_SNAT","SNAT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nATRuleInstanceTypeActionPropEnum = append(nATRuleInstanceTypeActionPropEnum, v)
	}
}

const (

	// NATRuleInstanceActionDNAT captures enum value "DNAT"
	NATRuleInstanceActionDNAT string = "DNAT"

	// NATRuleInstanceActionNOSNAT captures enum value "NO_SNAT"
	NATRuleInstanceActionNOSNAT string = "NO_SNAT"

	// NATRuleInstanceActionSNAT captures enum value "SNAT"
	NATRuleInstanceActionSNAT string = "SNAT"
)

// prop value enum
func (m *NATRuleInstance) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nATRuleInstanceTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NATRuleInstance) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *NATRuleInstance) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var nATRuleInstanceTypeFirewallMatchPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BYPASS","MATCH_EXTERNAL_ADDRESS","MATCH_INTERNAL_ADDRESS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nATRuleInstanceTypeFirewallMatchPropEnum = append(nATRuleInstanceTypeFirewallMatchPropEnum, v)
	}
}

const (

	// NATRuleInstanceFirewallMatchBYPASS captures enum value "BYPASS"
	NATRuleInstanceFirewallMatchBYPASS string = "BYPASS"

	// NATRuleInstanceFirewallMatchMATCHEXTERNALADDRESS captures enum value "MATCH_EXTERNAL_ADDRESS"
	NATRuleInstanceFirewallMatchMATCHEXTERNALADDRESS string = "MATCH_EXTERNAL_ADDRESS"

	// NATRuleInstanceFirewallMatchMATCHINTERNALADDRESS captures enum value "MATCH_INTERNAL_ADDRESS"
	NATRuleInstanceFirewallMatchMATCHINTERNALADDRESS string = "MATCH_INTERNAL_ADDRESS"
)

// prop value enum
func (m *NATRuleInstance) validateFirewallMatchEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nATRuleInstanceTypeFirewallMatchPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NATRuleInstance) validateFirewallMatch(formats strfmt.Registry) error {

	if swag.IsZero(m.FirewallMatch) { // not required
		return nil
	}

	// value enum
	if err := m.validateFirewallMatchEnum("firewallMatch", "body", m.FirewallMatch); err != nil {
		return err
	}

	return nil
}

func (m *NATRuleInstance) validateModificationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ModificationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modificationDate", "body", "date-time", m.ModificationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var nATRuleInstanceTypeNatTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AUTO","MANUAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nATRuleInstanceTypeNatTypePropEnum = append(nATRuleInstanceTypeNatTypePropEnum, v)
	}
}

const (

	// NATRuleInstanceNatTypeAUTO captures enum value "AUTO"
	NATRuleInstanceNatTypeAUTO string = "AUTO"

	// NATRuleInstanceNatTypeMANUAL captures enum value "MANUAL"
	NATRuleInstanceNatTypeMANUAL string = "MANUAL"
)

// prop value enum
func (m *NATRuleInstance) validateNatTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nATRuleInstanceTypeNatTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NATRuleInstance) validateNatType(formats strfmt.Registry) error {

	if swag.IsZero(m.NatType) { // not required
		return nil
	}

	// value enum
	if err := m.validateNatTypeEnum("natType", "body", m.NatType); err != nil {
		return err
	}

	return nil
}

func (m *NATRuleInstance) validateServiceInstance(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceInstance) { // not required
		return nil
	}

	if m.ServiceInstance != nil {
		if err := m.ServiceInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceInstance")
			}
			return err
		}
	}

	return nil
}

func (m *NATRuleInstance) validateTierZeroRouter(formats strfmt.Registry) error {

	if swag.IsZero(m.TierZeroRouter) { // not required
		return nil
	}

	if m.TierZeroRouter != nil {
		if err := m.TierZeroRouter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tierZeroRouter")
			}
			return err
		}
	}

	return nil
}

func (m *NATRuleInstance) validateVirtualNetworkInstance(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualNetworkInstance) { // not required
		return nil
	}

	if m.VirtualNetworkInstance != nil {
		if err := m.VirtualNetworkInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualNetworkInstance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NATRuleInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NATRuleInstance) UnmarshalBinary(b []byte) error {
	var res NATRuleInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}