// Code generated by go-swagger; DO NOT EDIT.

package gateway_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new gateway policies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for gateway policies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GatewayPolicyGetUsingGET(params *GatewayPolicyGetUsingGETParams, opts ...ClientOption) (*GatewayPolicyGetUsingGETOK, error)

	GatewayPolicyListUsingGET(params *GatewayPolicyListUsingGETParams, opts ...ClientOption) (*GatewayPolicyListUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GatewayPolicyGetUsingGET gets

  Get gateway policies from NSX-T
*/
func (a *Client) GatewayPolicyGetUsingGET(params *GatewayPolicyGetUsingGETParams, opts ...ClientOption) (*GatewayPolicyGetUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGatewayPolicyGetUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "gatewayPolicyGetUsingGET",
		Method:             "GET",
		PathPattern:        "/network/firewall/gateway-policies/{gatewayPolicyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GatewayPolicyGetUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GatewayPolicyGetUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for gatewayPolicyGetUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GatewayPolicyListUsingGET lists

  List gateway policies from NSX-T
*/
func (a *Client) GatewayPolicyListUsingGET(params *GatewayPolicyListUsingGETParams, opts ...ClientOption) (*GatewayPolicyListUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGatewayPolicyListUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "gatewayPolicyListUsingGET",
		Method:             "GET",
		PathPattern:        "/network/firewall/gateway-policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GatewayPolicyListUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GatewayPolicyListUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for gatewayPolicyListUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
