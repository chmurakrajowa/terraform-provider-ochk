// Code generated by go-swagger; DO NOT EDIT.

package active_directory_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new active directory groups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for active directory groups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateADGroupUsingPUT(params *CreateADGroupUsingPUTParams) (*CreateADGroupUsingPUTOK, *CreateADGroupUsingPUTCreated, error)

	DeleteADGroupUsingDELETE(params *DeleteADGroupUsingDELETEParams) (*DeleteADGroupUsingDELETEOK, error)

	GetADGroupUsingGET(params *GetADGroupUsingGETParams) (*GetADGroupUsingGETOK, error)

	ListADGroupUsingGET(params *ListADGroupUsingGETParams) (*ListADGroupUsingGETOK, error)

	UpdateADGroupUsingPUT(params *UpdateADGroupUsingPUTParams) (*UpdateADGroupUsingPUTOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateADGroupUsingPUT creates group

  Create group
*/
func (a *Client) CreateADGroupUsingPUT(params *CreateADGroupUsingPUTParams) (*CreateADGroupUsingPUTOK, *CreateADGroupUsingPUTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateADGroupUsingPUTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createADGroupUsingPUT",
		Method:             "PUT",
		PathPattern:        "/ad/integration/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateADGroupUsingPUTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateADGroupUsingPUTOK:
		return value, nil, nil
	case *CreateADGroupUsingPUTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for active_directory_groups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteADGroupUsingDELETE deletes group

  Delete group
*/
func (a *Client) DeleteADGroupUsingDELETE(params *DeleteADGroupUsingDELETEParams) (*DeleteADGroupUsingDELETEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteADGroupUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteADGroupUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/ad/integration/groups/{samAccountName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteADGroupUsingDELETEReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteADGroupUsingDELETEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteADGroupUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetADGroupUsingGET gets group info

  Get group info
*/
func (a *Client) GetADGroupUsingGET(params *GetADGroupUsingGETParams) (*GetADGroupUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetADGroupUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getADGroupUsingGET",
		Method:             "GET",
		PathPattern:        "/ad/integration/groups/{samAccountName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetADGroupUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetADGroupUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getADGroupUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListADGroupUsingGET lists groups

  List groups
*/
func (a *Client) ListADGroupUsingGET(params *ListADGroupUsingGETParams) (*ListADGroupUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListADGroupUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listADGroupUsingGET",
		Method:             "GET",
		PathPattern:        "/ad/integration/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListADGroupUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListADGroupUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listADGroupUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateADGroupUsingPUT updates group

  Update group
*/
func (a *Client) UpdateADGroupUsingPUT(params *UpdateADGroupUsingPUTParams) (*UpdateADGroupUsingPUTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateADGroupUsingPUTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateADGroupUsingPUT",
		Method:             "PUT",
		PathPattern:        "/ad/integration/groups/{samAccountName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateADGroupUsingPUTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateADGroupUsingPUTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateADGroupUsingPUT: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}