// Code generated by go-swagger; DO NOT EDIT.

package ip_a_m_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ip a m services API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ip a m services API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	IpamServiceGetUsingGET(params *IpamServiceGetUsingGETParams, opts ...ClientOption) (*IpamServiceGetUsingGETOK, error)

	IpamServicesListUsingGET(params *IpamServicesListUsingGETParams, opts ...ClientOption) (*IpamServicesListUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
IpamServiceGetUsingGET gets

Get IPAM service
*/
func (a *Client) IpamServiceGetUsingGET(params *IpamServiceGetUsingGETParams, opts ...ClientOption) (*IpamServiceGetUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamServiceGetUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipamServiceGetUsingGET",
		Method:             "GET",
		PathPattern:        "/ipam/services/{serviceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IpamServiceGetUsingGETReader{formats: a.formats},
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
	success, ok := result.(*IpamServiceGetUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipamServiceGetUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
IpamServicesListUsingGET lists

List IPAM services
*/
func (a *Client) IpamServicesListUsingGET(params *IpamServicesListUsingGETParams, opts ...ClientOption) (*IpamServicesListUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamServicesListUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipamServicesListUsingGET",
		Method:             "GET",
		PathPattern:        "/ipam/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IpamServicesListUsingGETReader{formats: a.formats},
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
	success, ok := result.(*IpamServicesListUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipamServicesListUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
