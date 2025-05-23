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

// CustomServiceInstance custom service instance
//
// swagger:model CustomServiceInstance
type CustomServiceInstance struct {

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// creation date
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// external Id
	ExternalID string `json:"externalId,omitempty"`

	// is default
	IsDefault bool `json:"isDefault,omitempty"`

	// l4 port set entries
	L4PortSetEntries []*L4PortSetEntry `json:"l4PortSetEntries"`

	// modification date
	// Format: date-time
	ModificationDate strfmt.DateTime `json:"modificationDate,omitempty"`

	// modified by
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// parent path
	ParentPath string `json:"parentPath,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// project Id
	// Format: uuid
	ProjectID strfmt.UUID `json:"projectId,omitempty"`

	// related fw rules
	RelatedFwRules []*RuleInstance `json:"relatedFwRules"`

	// relative path
	RelativePath string `json:"relativePath,omitempty"`

	// resource type
	ResourceType *ResourceType `json:"resourceType,omitempty"`

	// service Id
	// Format: uuid
	ServiceID strfmt.UUID `json:"serviceId,omitempty"`

	// service type
	ServiceType *ServiceType `json:"serviceType,omitempty"`

	// tags
	Tags []*TagInstance `json:"tags"`

	// tenant instance
	TenantInstance *TenantInstance `json:"tenantInstance,omitempty"`
}

// Validate validates this custom service instance
func (m *CustomServiceInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateL4PortSetEntries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModificationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedFwRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantInstance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomServiceInstance) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomServiceInstance) validateL4PortSetEntries(formats strfmt.Registry) error {
	if swag.IsZero(m.L4PortSetEntries) { // not required
		return nil
	}

	for i := 0; i < len(m.L4PortSetEntries); i++ {
		if swag.IsZero(m.L4PortSetEntries[i]) { // not required
			continue
		}

		if m.L4PortSetEntries[i] != nil {
			if err := m.L4PortSetEntries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("l4PortSetEntries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("l4PortSetEntries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomServiceInstance) validateModificationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ModificationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modificationDate", "body", "date-time", m.ModificationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomServiceInstance) validateProjectID(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectID) { // not required
		return nil
	}

	if err := validate.FormatOf("projectId", "body", "uuid", m.ProjectID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomServiceInstance) validateRelatedFwRules(formats strfmt.Registry) error {
	if swag.IsZero(m.RelatedFwRules) { // not required
		return nil
	}

	for i := 0; i < len(m.RelatedFwRules); i++ {
		if swag.IsZero(m.RelatedFwRules[i]) { // not required
			continue
		}

		if m.RelatedFwRules[i] != nil {
			if err := m.RelatedFwRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relatedFwRules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("relatedFwRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomServiceInstance) validateResourceType(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceType) { // not required
		return nil
	}

	if m.ResourceType != nil {
		if err := m.ResourceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceType")
			}
			return err
		}
	}

	return nil
}

func (m *CustomServiceInstance) validateServiceID(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceID) { // not required
		return nil
	}

	if err := validate.FormatOf("serviceId", "body", "uuid", m.ServiceID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomServiceInstance) validateServiceType(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceType) { // not required
		return nil
	}

	if m.ServiceType != nil {
		if err := m.ServiceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serviceType")
			}
			return err
		}
	}

	return nil
}

func (m *CustomServiceInstance) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomServiceInstance) validateTenantInstance(formats strfmt.Registry) error {
	if swag.IsZero(m.TenantInstance) { // not required
		return nil
	}

	if m.TenantInstance != nil {
		if err := m.TenantInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenantInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tenantInstance")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this custom service instance based on the context it is used
func (m *CustomServiceInstance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateL4PortSetEntries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelatedFwRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTenantInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomServiceInstance) contextValidateL4PortSetEntries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.L4PortSetEntries); i++ {

		if m.L4PortSetEntries[i] != nil {

			if swag.IsZero(m.L4PortSetEntries[i]) { // not required
				return nil
			}

			if err := m.L4PortSetEntries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("l4PortSetEntries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("l4PortSetEntries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomServiceInstance) contextValidateRelatedFwRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RelatedFwRules); i++ {

		if m.RelatedFwRules[i] != nil {

			if swag.IsZero(m.RelatedFwRules[i]) { // not required
				return nil
			}

			if err := m.RelatedFwRules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relatedFwRules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("relatedFwRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomServiceInstance) contextValidateResourceType(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceType != nil {

		if swag.IsZero(m.ResourceType) { // not required
			return nil
		}

		if err := m.ResourceType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceType")
			}
			return err
		}
	}

	return nil
}

func (m *CustomServiceInstance) contextValidateServiceType(ctx context.Context, formats strfmt.Registry) error {

	if m.ServiceType != nil {

		if swag.IsZero(m.ServiceType) { // not required
			return nil
		}

		if err := m.ServiceType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serviceType")
			}
			return err
		}
	}

	return nil
}

func (m *CustomServiceInstance) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {

			if swag.IsZero(m.Tags[i]) { // not required
				return nil
			}

			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomServiceInstance) contextValidateTenantInstance(ctx context.Context, formats strfmt.Registry) error {

	if m.TenantInstance != nil {

		if swag.IsZero(m.TenantInstance) { // not required
			return nil
		}

		if err := m.TenantInstance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenantInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tenantInstance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomServiceInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomServiceInstance) UnmarshalBinary(b []byte) error {
	var res CustomServiceInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
