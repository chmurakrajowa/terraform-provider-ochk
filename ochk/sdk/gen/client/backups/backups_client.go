// Code generated by go-swagger; DO NOT EDIT.

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new backups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for backups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	BackupListGetUsingGET(params *BackupListGetUsingGETParams, opts ...ClientOption) (*BackupListGetUsingGETOK, error)

	BackupListListUsingGET(params *BackupListListUsingGETParams, opts ...ClientOption) (*BackupListListUsingGETOK, error)

	BackupPlanGetUsingGET(params *BackupPlanGetUsingGETParams, opts ...ClientOption) (*BackupPlanGetUsingGETOK, error)

	BackupPlanListUsingGET(params *BackupPlanListUsingGETParams, opts ...ClientOption) (*BackupPlanListUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BackupListGetUsingGET gets backup list get

  Get backup list
*/
func (a *Client) BackupListGetUsingGET(params *BackupListGetUsingGETParams, opts ...ClientOption) (*BackupListGetUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupListGetUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "backupListGetUsingGET",
		Method:             "GET",
		PathPattern:        "/backups/plans/{backupPlanId}/lists/{backupListId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupListGetUsingGETReader{formats: a.formats},
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
	success, ok := result.(*BackupListGetUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for backupListGetUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BackupListListUsingGET lists backup lists

  List backup lists
*/
func (a *Client) BackupListListUsingGET(params *BackupListListUsingGETParams, opts ...ClientOption) (*BackupListListUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupListListUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "backupListListUsingGET",
		Method:             "GET",
		PathPattern:        "/backups/plans/{backupPlanId}/lists",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupListListUsingGETReader{formats: a.formats},
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
	success, ok := result.(*BackupListListUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for backupListListUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BackupPlanGetUsingGET gets backup plan

  Get backup plan
*/
func (a *Client) BackupPlanGetUsingGET(params *BackupPlanGetUsingGETParams, opts ...ClientOption) (*BackupPlanGetUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupPlanGetUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "backupPlanGetUsingGET",
		Method:             "GET",
		PathPattern:        "/backups/plans/{backupPlanId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupPlanGetUsingGETReader{formats: a.formats},
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
	success, ok := result.(*BackupPlanGetUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for backupPlanGetUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BackupPlanListUsingGET lists backup plans

  List backup plans
*/
func (a *Client) BackupPlanListUsingGET(params *BackupPlanListUsingGETParams, opts ...ClientOption) (*BackupPlanListUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupPlanListUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "backupPlanListUsingGET",
		Method:             "GET",
		PathPattern:        "/backups/plans",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupPlanListUsingGETReader{formats: a.formats},
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
	success, ok := result.(*BackupPlanListUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for backupPlanListUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
