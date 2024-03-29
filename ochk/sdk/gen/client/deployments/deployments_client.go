// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new deployments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for deployments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeploymentGetUsingGET(params *DeploymentGetUsingGETParams, opts ...ClientOption) (*DeploymentGetUsingGETOK, error)

	DeploymentListUsingGET(params *DeploymentListUsingGETParams, opts ...ClientOption) (*DeploymentListUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeploymentGetUsingGET gets

Get deployment object
*/
func (a *Client) DeploymentGetUsingGET(params *DeploymentGetUsingGETParams, opts ...ClientOption) (*DeploymentGetUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeploymentGetUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deploymentGetUsingGET",
		Method:             "GET",
		PathPattern:        "/deployments/{deploymentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeploymentGetUsingGETReader{formats: a.formats},
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
	success, ok := result.(*DeploymentGetUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deploymentGetUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeploymentListUsingGET lists

List deployments from server
*/
func (a *Client) DeploymentListUsingGET(params *DeploymentListUsingGETParams, opts ...ClientOption) (*DeploymentListUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeploymentListUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deploymentListUsingGET",
		Method:             "GET",
		PathPattern:        "/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeploymentListUsingGETReader{formats: a.formats},
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
	success, ok := result.(*DeploymentListUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deploymentListUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
