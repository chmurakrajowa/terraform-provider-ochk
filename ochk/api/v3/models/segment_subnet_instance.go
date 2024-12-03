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

// SegmentSubnetInstance segment subnet instance
//
// swagger:model SegmentSubnetInstance
type SegmentSubnetInstance struct {

	// dhcp ranges
	DhcpRanges []string `json:"dhcpRanges"`

	// dns servers
	DNSServers []*DNSServerInstance `json:"dnsServers"`

	// gateway address c ID r
	GatewayAddressCIDR string `json:"gatewayAddressCIDR,omitempty"`

	// network c ID r
	NetworkCIDR string `json:"networkCIDR,omitempty"`

	// segment subnet Id
	// Format: uuid
	SegmentSubnetID strfmt.UUID `json:"segmentSubnetId,omitempty"`
}

// Validate validates this segment subnet instance
func (m *SegmentSubnetInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDNSServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSegmentSubnetID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SegmentSubnetInstance) validateDNSServers(formats strfmt.Registry) error {
	if swag.IsZero(m.DNSServers) { // not required
		return nil
	}

	for i := 0; i < len(m.DNSServers); i++ {
		if swag.IsZero(m.DNSServers[i]) { // not required
			continue
		}

		if m.DNSServers[i] != nil {
			if err := m.DNSServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dnsServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dnsServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SegmentSubnetInstance) validateSegmentSubnetID(formats strfmt.Registry) error {
	if swag.IsZero(m.SegmentSubnetID) { // not required
		return nil
	}

	if err := validate.FormatOf("segmentSubnetId", "body", "uuid", m.SegmentSubnetID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this segment subnet instance based on the context it is used
func (m *SegmentSubnetInstance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDNSServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SegmentSubnetInstance) contextValidateDNSServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DNSServers); i++ {

		if m.DNSServers[i] != nil {

			if swag.IsZero(m.DNSServers[i]) { // not required
				return nil
			}

			if err := m.DNSServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dnsServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dnsServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SegmentSubnetInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SegmentSubnetInstance) UnmarshalBinary(b []byte) error {
	var res SegmentSubnetInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
