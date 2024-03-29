// Code generated by go-swagger; DO NOT EDIT.

package snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new snapshots API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for snapshots API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SnapshotGetUsingGET(params *SnapshotGetUsingGETParams, opts ...ClientOption) (*SnapshotGetUsingGETOK, error)

	SnapshotListUsingGET(params *SnapshotListUsingGETParams, opts ...ClientOption) (*SnapshotListUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
SnapshotGetUsingGET gets

Get vSphere vCenter virtual machine snapshot
*/
func (a *Client) SnapshotGetUsingGET(params *SnapshotGetUsingGETParams, opts ...ClientOption) (*SnapshotGetUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSnapshotGetUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "snapshotGetUsingGET",
		Method:             "GET",
		PathPattern:        "/vcs/snapshots/{snapshotId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SnapshotGetUsingGETReader{formats: a.formats},
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
	success, ok := result.(*SnapshotGetUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for snapshotGetUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SnapshotListUsingGET lists

List vSphere vCenter virtual machines snapshots
*/
func (a *Client) SnapshotListUsingGET(params *SnapshotListUsingGETParams, opts ...ClientOption) (*SnapshotListUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSnapshotListUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "snapshotListUsingGET",
		Method:             "GET",
		PathPattern:        "/vcs/snapshots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SnapshotListUsingGETReader{formats: a.formats},
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
	success, ok := result.(*SnapshotListUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for snapshotListUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
