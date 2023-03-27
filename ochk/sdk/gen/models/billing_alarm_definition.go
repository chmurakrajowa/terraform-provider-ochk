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

// BillingAlarmDefinition BillingAlarmDefinition
//
// swagger:model BillingAlarmDefinition
type BillingAlarmDefinition struct {

	// account Id
	AccountID string `json:"accountId,omitempty"`

	// alarm definition description
	AlarmDefinitionDescription string `json:"alarmDefinitionDescription,omitempty"`

	// alarm definition Id
	AlarmDefinitionID string `json:"alarmDefinitionId,omitempty"`

	// alarm definition name
	AlarmDefinitionName string `json:"alarmDefinitionName,omitempty"`

	// alarm enabled
	AlarmEnabled bool `json:"alarmEnabled,omitempty"`

	// all criteria must be meet
	AllCriteriaMustBeMeet bool `json:"allCriteriaMustBeMeet,omitempty"`

	// block resource deployment
	BlockResourceDeployment bool `json:"blockResourceDeployment,omitempty"`

	// cpu allocation alarm enabled
	CPUAllocationAlarmEnabled bool `json:"cpuAllocationAlarmEnabled,omitempty"`

	// cpu allocation operator
	// Enum: [EQUALS HIGHER LOWER]
	CPUAllocationOperator string `json:"cpuAllocationOperator,omitempty"`

	// cpu allocation threshold value
	CPUAllocationThresholdValue float32 `json:"cpuAllocationThresholdValue,omitempty"`

	// cpu usage alarm enabled
	CPUUsageAlarmEnabled bool `json:"cpuUsageAlarmEnabled,omitempty"`

	// cpu usage operator
	// Enum: [EQUALS HIGHER LOWER]
	CPUUsageOperator string `json:"cpuUsageOperator,omitempty"`

	// cpu usage threshold value
	CPUUsageThresholdValue float32 `json:"cpuUsageThresholdValue,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// creation date
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate,omitempty"`

	// group instance list
	GroupInstanceList []*GroupInstance `json:"groupInstanceList"`

	// modification date
	// Format: date-time
	ModificationDate *strfmt.DateTime `json:"modificationDate,omitempty"`

	// modified by
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// ram allocation alarm enabled
	RAMAllocationAlarmEnabled bool `json:"ramAllocationAlarmEnabled,omitempty"`

	// ram allocation operator
	// Enum: [EQUALS HIGHER LOWER]
	RAMAllocationOperator string `json:"ramAllocationOperator,omitempty"`

	// ram allocation threshold value
	RAMAllocationThresholdValue float32 `json:"ramAllocationThresholdValue,omitempty"`

	// ram usage alarm enabled
	RAMUsageAlarmEnabled bool `json:"ramUsageAlarmEnabled,omitempty"`

	// ram usage operator
	// Enum: [EQUALS HIGHER LOWER]
	RAMUsageOperator string `json:"ramUsageOperator,omitempty"`

	// ram usage threshold value
	RAMUsageThresholdValue float32 `json:"ramUsageThresholdValue,omitempty"`

	// send e mail notification
	SendEMailNotification bool `json:"sendEMailNotification,omitempty"`

	// send notification
	SendNotification bool `json:"sendNotification,omitempty"`

	// storage allocation alarm enabled
	StorageAllocationAlarmEnabled bool `json:"storageAllocationAlarmEnabled,omitempty"`

	// storage allocation operator
	// Enum: [EQUALS HIGHER LOWER]
	StorageAllocationOperator string `json:"storageAllocationOperator,omitempty"`

	// storage allocation threshold value
	StorageAllocationThresholdValue float32 `json:"storageAllocationThresholdValue,omitempty"`

	// storage usage alarm enabled
	StorageUsageAlarmEnabled bool `json:"storageUsageAlarmEnabled,omitempty"`

	// storage usage operator
	// Enum: [EQUALS HIGHER LOWER]
	StorageUsageOperator string `json:"storageUsageOperator,omitempty"`

	// storage usage threshold value
	StorageUsageThresholdValue float32 `json:"storageUsageThresholdValue,omitempty"`

	// tag Id
	TagID int32 `json:"tagId,omitempty"`

	// total cost threshold value
	TotalCostThresholdValue float32 `json:"totalCostThresholdValue,omitempty"`

	// total cost usage alarm enabled
	TotalCostUsageAlarmEnabled bool `json:"totalCostUsageAlarmEnabled,omitempty"`

	// total cost usage operator
	// Enum: [EQUALS HIGHER LOWER]
	TotalCostUsageOperator string `json:"totalCostUsageOperator,omitempty"`

	// user instance list
	UserInstanceList []*UserInstance `json:"userInstanceList"`
}

// Validate validates this billing alarm definition
func (m *BillingAlarmDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUAllocationOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUUsageOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupInstanceList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModificationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRAMAllocationOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRAMUsageOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageAllocationOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageUsageOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCostUsageOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserInstanceList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var billingAlarmDefinitionTypeCPUAllocationOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EQUALS","HIGHER","LOWER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billingAlarmDefinitionTypeCPUAllocationOperatorPropEnum = append(billingAlarmDefinitionTypeCPUAllocationOperatorPropEnum, v)
	}
}

const (

	// BillingAlarmDefinitionCPUAllocationOperatorEQUALS captures enum value "EQUALS"
	BillingAlarmDefinitionCPUAllocationOperatorEQUALS string = "EQUALS"

	// BillingAlarmDefinitionCPUAllocationOperatorHIGHER captures enum value "HIGHER"
	BillingAlarmDefinitionCPUAllocationOperatorHIGHER string = "HIGHER"

	// BillingAlarmDefinitionCPUAllocationOperatorLOWER captures enum value "LOWER"
	BillingAlarmDefinitionCPUAllocationOperatorLOWER string = "LOWER"
)

// prop value enum
func (m *BillingAlarmDefinition) validateCPUAllocationOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, billingAlarmDefinitionTypeCPUAllocationOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BillingAlarmDefinition) validateCPUAllocationOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.CPUAllocationOperator) { // not required
		return nil
	}

	// value enum
	if err := m.validateCPUAllocationOperatorEnum("cpuAllocationOperator", "body", m.CPUAllocationOperator); err != nil {
		return err
	}

	return nil
}

var billingAlarmDefinitionTypeCPUUsageOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EQUALS","HIGHER","LOWER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billingAlarmDefinitionTypeCPUUsageOperatorPropEnum = append(billingAlarmDefinitionTypeCPUUsageOperatorPropEnum, v)
	}
}

const (

	// BillingAlarmDefinitionCPUUsageOperatorEQUALS captures enum value "EQUALS"
	BillingAlarmDefinitionCPUUsageOperatorEQUALS string = "EQUALS"

	// BillingAlarmDefinitionCPUUsageOperatorHIGHER captures enum value "HIGHER"
	BillingAlarmDefinitionCPUUsageOperatorHIGHER string = "HIGHER"

	// BillingAlarmDefinitionCPUUsageOperatorLOWER captures enum value "LOWER"
	BillingAlarmDefinitionCPUUsageOperatorLOWER string = "LOWER"
)

// prop value enum
func (m *BillingAlarmDefinition) validateCPUUsageOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, billingAlarmDefinitionTypeCPUUsageOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BillingAlarmDefinition) validateCPUUsageOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.CPUUsageOperator) { // not required
		return nil
	}

	// value enum
	if err := m.validateCPUUsageOperatorEnum("cpuUsageOperator", "body", m.CPUUsageOperator); err != nil {
		return err
	}

	return nil
}

func (m *BillingAlarmDefinition) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BillingAlarmDefinition) validateGroupInstanceList(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupInstanceList) { // not required
		return nil
	}

	for i := 0; i < len(m.GroupInstanceList); i++ {
		if swag.IsZero(m.GroupInstanceList[i]) { // not required
			continue
		}

		if m.GroupInstanceList[i] != nil {
			if err := m.GroupInstanceList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groupInstanceList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("groupInstanceList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BillingAlarmDefinition) validateModificationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ModificationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modificationDate", "body", "date-time", m.ModificationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var billingAlarmDefinitionTypeRAMAllocationOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EQUALS","HIGHER","LOWER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billingAlarmDefinitionTypeRAMAllocationOperatorPropEnum = append(billingAlarmDefinitionTypeRAMAllocationOperatorPropEnum, v)
	}
}

const (

	// BillingAlarmDefinitionRAMAllocationOperatorEQUALS captures enum value "EQUALS"
	BillingAlarmDefinitionRAMAllocationOperatorEQUALS string = "EQUALS"

	// BillingAlarmDefinitionRAMAllocationOperatorHIGHER captures enum value "HIGHER"
	BillingAlarmDefinitionRAMAllocationOperatorHIGHER string = "HIGHER"

	// BillingAlarmDefinitionRAMAllocationOperatorLOWER captures enum value "LOWER"
	BillingAlarmDefinitionRAMAllocationOperatorLOWER string = "LOWER"
)

// prop value enum
func (m *BillingAlarmDefinition) validateRAMAllocationOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, billingAlarmDefinitionTypeRAMAllocationOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BillingAlarmDefinition) validateRAMAllocationOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.RAMAllocationOperator) { // not required
		return nil
	}

	// value enum
	if err := m.validateRAMAllocationOperatorEnum("ramAllocationOperator", "body", m.RAMAllocationOperator); err != nil {
		return err
	}

	return nil
}

var billingAlarmDefinitionTypeRAMUsageOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EQUALS","HIGHER","LOWER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billingAlarmDefinitionTypeRAMUsageOperatorPropEnum = append(billingAlarmDefinitionTypeRAMUsageOperatorPropEnum, v)
	}
}

const (

	// BillingAlarmDefinitionRAMUsageOperatorEQUALS captures enum value "EQUALS"
	BillingAlarmDefinitionRAMUsageOperatorEQUALS string = "EQUALS"

	// BillingAlarmDefinitionRAMUsageOperatorHIGHER captures enum value "HIGHER"
	BillingAlarmDefinitionRAMUsageOperatorHIGHER string = "HIGHER"

	// BillingAlarmDefinitionRAMUsageOperatorLOWER captures enum value "LOWER"
	BillingAlarmDefinitionRAMUsageOperatorLOWER string = "LOWER"
)

// prop value enum
func (m *BillingAlarmDefinition) validateRAMUsageOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, billingAlarmDefinitionTypeRAMUsageOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BillingAlarmDefinition) validateRAMUsageOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.RAMUsageOperator) { // not required
		return nil
	}

	// value enum
	if err := m.validateRAMUsageOperatorEnum("ramUsageOperator", "body", m.RAMUsageOperator); err != nil {
		return err
	}

	return nil
}

var billingAlarmDefinitionTypeStorageAllocationOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EQUALS","HIGHER","LOWER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billingAlarmDefinitionTypeStorageAllocationOperatorPropEnum = append(billingAlarmDefinitionTypeStorageAllocationOperatorPropEnum, v)
	}
}

const (

	// BillingAlarmDefinitionStorageAllocationOperatorEQUALS captures enum value "EQUALS"
	BillingAlarmDefinitionStorageAllocationOperatorEQUALS string = "EQUALS"

	// BillingAlarmDefinitionStorageAllocationOperatorHIGHER captures enum value "HIGHER"
	BillingAlarmDefinitionStorageAllocationOperatorHIGHER string = "HIGHER"

	// BillingAlarmDefinitionStorageAllocationOperatorLOWER captures enum value "LOWER"
	BillingAlarmDefinitionStorageAllocationOperatorLOWER string = "LOWER"
)

// prop value enum
func (m *BillingAlarmDefinition) validateStorageAllocationOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, billingAlarmDefinitionTypeStorageAllocationOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BillingAlarmDefinition) validateStorageAllocationOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageAllocationOperator) { // not required
		return nil
	}

	// value enum
	if err := m.validateStorageAllocationOperatorEnum("storageAllocationOperator", "body", m.StorageAllocationOperator); err != nil {
		return err
	}

	return nil
}

var billingAlarmDefinitionTypeStorageUsageOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EQUALS","HIGHER","LOWER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billingAlarmDefinitionTypeStorageUsageOperatorPropEnum = append(billingAlarmDefinitionTypeStorageUsageOperatorPropEnum, v)
	}
}

const (

	// BillingAlarmDefinitionStorageUsageOperatorEQUALS captures enum value "EQUALS"
	BillingAlarmDefinitionStorageUsageOperatorEQUALS string = "EQUALS"

	// BillingAlarmDefinitionStorageUsageOperatorHIGHER captures enum value "HIGHER"
	BillingAlarmDefinitionStorageUsageOperatorHIGHER string = "HIGHER"

	// BillingAlarmDefinitionStorageUsageOperatorLOWER captures enum value "LOWER"
	BillingAlarmDefinitionStorageUsageOperatorLOWER string = "LOWER"
)

// prop value enum
func (m *BillingAlarmDefinition) validateStorageUsageOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, billingAlarmDefinitionTypeStorageUsageOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BillingAlarmDefinition) validateStorageUsageOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageUsageOperator) { // not required
		return nil
	}

	// value enum
	if err := m.validateStorageUsageOperatorEnum("storageUsageOperator", "body", m.StorageUsageOperator); err != nil {
		return err
	}

	return nil
}

var billingAlarmDefinitionTypeTotalCostUsageOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EQUALS","HIGHER","LOWER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billingAlarmDefinitionTypeTotalCostUsageOperatorPropEnum = append(billingAlarmDefinitionTypeTotalCostUsageOperatorPropEnum, v)
	}
}

const (

	// BillingAlarmDefinitionTotalCostUsageOperatorEQUALS captures enum value "EQUALS"
	BillingAlarmDefinitionTotalCostUsageOperatorEQUALS string = "EQUALS"

	// BillingAlarmDefinitionTotalCostUsageOperatorHIGHER captures enum value "HIGHER"
	BillingAlarmDefinitionTotalCostUsageOperatorHIGHER string = "HIGHER"

	// BillingAlarmDefinitionTotalCostUsageOperatorLOWER captures enum value "LOWER"
	BillingAlarmDefinitionTotalCostUsageOperatorLOWER string = "LOWER"
)

// prop value enum
func (m *BillingAlarmDefinition) validateTotalCostUsageOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, billingAlarmDefinitionTypeTotalCostUsageOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BillingAlarmDefinition) validateTotalCostUsageOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.TotalCostUsageOperator) { // not required
		return nil
	}

	// value enum
	if err := m.validateTotalCostUsageOperatorEnum("totalCostUsageOperator", "body", m.TotalCostUsageOperator); err != nil {
		return err
	}

	return nil
}

func (m *BillingAlarmDefinition) validateUserInstanceList(formats strfmt.Registry) error {
	if swag.IsZero(m.UserInstanceList) { // not required
		return nil
	}

	for i := 0; i < len(m.UserInstanceList); i++ {
		if swag.IsZero(m.UserInstanceList[i]) { // not required
			continue
		}

		if m.UserInstanceList[i] != nil {
			if err := m.UserInstanceList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userInstanceList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userInstanceList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this billing alarm definition based on the context it is used
func (m *BillingAlarmDefinition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroupInstanceList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserInstanceList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingAlarmDefinition) contextValidateGroupInstanceList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GroupInstanceList); i++ {

		if m.GroupInstanceList[i] != nil {
			if err := m.GroupInstanceList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groupInstanceList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("groupInstanceList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BillingAlarmDefinition) contextValidateUserInstanceList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UserInstanceList); i++ {

		if m.UserInstanceList[i] != nil {
			if err := m.UserInstanceList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userInstanceList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userInstanceList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillingAlarmDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingAlarmDefinition) UnmarshalBinary(b []byte) error {
	var res BillingAlarmDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
