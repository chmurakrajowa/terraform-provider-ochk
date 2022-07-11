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

// DeploymentParam DeploymentParam
//
// swagger:model DeploymentParam
type DeploymentParam struct {

	// param name
	ParamName string `json:"paramName,omitempty"`

	// param type
	// Enum: [HOST_NAME LOGIN_PASSWORD LOGIN_SSH_KEY LOGIN_USERNAME NET_DNS_PRIMARY NET_DNS_SEARCH NET_DNS_SECONDARY NET_DNS_SUFFIX NET_IP_ADDR_01 NET_IP_BROADCAST_01 NET_IP_GATEWAY_01 NET_IP_MASK_01 NET_WINS_ALTENATIVE NET_WINS_PREFERRED]
	ParamType string `json:"paramType,omitempty"`

	// param value
	ParamValue string `json:"paramValue,omitempty"`
}

// Validate validates this deployment param
func (m *DeploymentParam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParamType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var deploymentParamTypeParamTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HOST_NAME","LOGIN_PASSWORD","LOGIN_SSH_KEY","LOGIN_USERNAME","NET_DNS_PRIMARY","NET_DNS_SEARCH","NET_DNS_SECONDARY","NET_DNS_SUFFIX","NET_IP_ADDR_01","NET_IP_BROADCAST_01","NET_IP_GATEWAY_01","NET_IP_MASK_01","NET_WINS_ALTENATIVE","NET_WINS_PREFERRED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deploymentParamTypeParamTypePropEnum = append(deploymentParamTypeParamTypePropEnum, v)
	}
}

const (

	// DeploymentParamParamTypeHOSTNAME captures enum value "HOST_NAME"
	DeploymentParamParamTypeHOSTNAME string = "HOST_NAME"

	// DeploymentParamParamTypeLOGINPASSWORD captures enum value "LOGIN_PASSWORD"
	DeploymentParamParamTypeLOGINPASSWORD string = "LOGIN_PASSWORD"

	// DeploymentParamParamTypeLOGINSSHKEY captures enum value "LOGIN_SSH_KEY"
	DeploymentParamParamTypeLOGINSSHKEY string = "LOGIN_SSH_KEY"

	// DeploymentParamParamTypeLOGINUSERNAME captures enum value "LOGIN_USERNAME"
	DeploymentParamParamTypeLOGINUSERNAME string = "LOGIN_USERNAME"

	// DeploymentParamParamTypeNETDNSPRIMARY captures enum value "NET_DNS_PRIMARY"
	DeploymentParamParamTypeNETDNSPRIMARY string = "NET_DNS_PRIMARY"

	// DeploymentParamParamTypeNETDNSSEARCH captures enum value "NET_DNS_SEARCH"
	DeploymentParamParamTypeNETDNSSEARCH string = "NET_DNS_SEARCH"

	// DeploymentParamParamTypeNETDNSSECONDARY captures enum value "NET_DNS_SECONDARY"
	DeploymentParamParamTypeNETDNSSECONDARY string = "NET_DNS_SECONDARY"

	// DeploymentParamParamTypeNETDNSSUFFIX captures enum value "NET_DNS_SUFFIX"
	DeploymentParamParamTypeNETDNSSUFFIX string = "NET_DNS_SUFFIX"

	// DeploymentParamParamTypeNETIPADDR01 captures enum value "NET_IP_ADDR_01"
	DeploymentParamParamTypeNETIPADDR01 string = "NET_IP_ADDR_01"

	// DeploymentParamParamTypeNETIPBROADCAST01 captures enum value "NET_IP_BROADCAST_01"
	DeploymentParamParamTypeNETIPBROADCAST01 string = "NET_IP_BROADCAST_01"

	// DeploymentParamParamTypeNETIPGATEWAY01 captures enum value "NET_IP_GATEWAY_01"
	DeploymentParamParamTypeNETIPGATEWAY01 string = "NET_IP_GATEWAY_01"

	// DeploymentParamParamTypeNETIPMASK01 captures enum value "NET_IP_MASK_01"
	DeploymentParamParamTypeNETIPMASK01 string = "NET_IP_MASK_01"

	// DeploymentParamParamTypeNETWINSALTENATIVE captures enum value "NET_WINS_ALTENATIVE"
	DeploymentParamParamTypeNETWINSALTENATIVE string = "NET_WINS_ALTENATIVE"

	// DeploymentParamParamTypeNETWINSPREFERRED captures enum value "NET_WINS_PREFERRED"
	DeploymentParamParamTypeNETWINSPREFERRED string = "NET_WINS_PREFERRED"
)

// prop value enum
func (m *DeploymentParam) validateParamTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deploymentParamTypeParamTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeploymentParam) validateParamType(formats strfmt.Registry) error {

	if swag.IsZero(m.ParamType) { // not required
		return nil
	}

	// value enum
	if err := m.validateParamTypeEnum("paramType", "body", m.ParamType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentParam) UnmarshalBinary(b []byte) error {
	var res DeploymentParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}