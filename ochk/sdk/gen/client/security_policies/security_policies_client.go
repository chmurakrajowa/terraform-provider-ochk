// Code generated by go-swagger; DO NOT EDIT.

package security_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new security policies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for security policies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	SecurityPolicyGetUsingGET(params *SecurityPolicyGetUsingGETParams) (*SecurityPolicyGetUsingGETOK, error)

	SecurityPolicyListUsingGET(params *SecurityPolicyListUsingGETParams) (*SecurityPolicyListUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SecurityPolicyGetUsingGET gets

  Get security policy from NSX-T
*/
func (a *Client) SecurityPolicyGetUsingGET(params *SecurityPolicyGetUsingGETParams) (*SecurityPolicyGetUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecurityPolicyGetUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "securityPolicyGetUsingGET",
		Method:             "GET",
		PathPattern:        "/network/firewall/security-policies/{SecurityPolicyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SecurityPolicyGetUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecurityPolicyGetUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for securityPolicyGetUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecurityPolicyListUsingGET lists

  List security policies from NSX-T
*/
func (a *Client) SecurityPolicyListUsingGET(params *SecurityPolicyListUsingGETParams) (*SecurityPolicyListUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecurityPolicyListUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "securityPolicyListUsingGET",
		Method:             "GET",
		PathPattern:        "/network/firewall/security-policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SecurityPolicyListUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecurityPolicyListUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for securityPolicyListUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
