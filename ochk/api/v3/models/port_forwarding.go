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

// PortForwarding port forwarding
//
// swagger:model PortForwarding
type PortForwarding struct {

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// creation date
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// external Id
	ExternalID string `json:"externalId,omitempty"`

	// external port
	ExternalPort int32 `json:"externalPort,omitempty"`

	// external port range
	ExternalPortRange string `json:"externalPortRange,omitempty"`

	// floating Ip Id
	// Format: uuid
	FloatingIPID strfmt.UUID `json:"floatingIpId,omitempty"`

	// internal Ip address
	InternalIPAddress string `json:"internalIpAddress,omitempty"`

	// internal port
	InternalPort int32 `json:"internalPort,omitempty"`

	// internal port Id
	// Format: uuid
	InternalPortID strfmt.UUID `json:"internalPortId,omitempty"`

	// internal port range
	InternalPortRange string `json:"internalPortRange,omitempty"`

	// modification date
	// Format: date-time
	ModificationDate strfmt.DateTime `json:"modificationDate,omitempty"`

	// modified by
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// port forwarding Id
	// Format: uuid
	PortForwardingID strfmt.UUID `json:"portForwardingId,omitempty"`

	// protocol
	Protocol string `json:"protocol,omitempty"`

	// public address
	PublicAddress string `json:"publicAddress,omitempty"`
}

// Validate validates this port forwarding
func (m *PortForwarding) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFloatingIPID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternalPortID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModificationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortForwardingID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortForwarding) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PortForwarding) validateFloatingIPID(formats strfmt.Registry) error {
	if swag.IsZero(m.FloatingIPID) { // not required
		return nil
	}

	if err := validate.FormatOf("floatingIpId", "body", "uuid", m.FloatingIPID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PortForwarding) validateInternalPortID(formats strfmt.Registry) error {
	if swag.IsZero(m.InternalPortID) { // not required
		return nil
	}

	if err := validate.FormatOf("internalPortId", "body", "uuid", m.InternalPortID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PortForwarding) validateModificationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ModificationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modificationDate", "body", "date-time", m.ModificationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PortForwarding) validatePortForwardingID(formats strfmt.Registry) error {
	if swag.IsZero(m.PortForwardingID) { // not required
		return nil
	}

	if err := validate.FormatOf("portForwardingId", "body", "uuid", m.PortForwardingID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this port forwarding based on context it is used
func (m *PortForwarding) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortForwarding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortForwarding) UnmarshalBinary(b []byte) error {
	var res PortForwarding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
