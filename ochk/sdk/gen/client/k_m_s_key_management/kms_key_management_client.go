// Code generated by go-swagger; DO NOT EDIT.

package k_m_s_key_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new k m s key management API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for k m s key management API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	KeyCreateUsingPUT(params *KeyCreateUsingPUTParams, opts ...ClientOption) (*KeyCreateUsingPUTOK, *KeyCreateUsingPUTCreated, error)

	KeyDeleteUsingDELETE(params *KeyDeleteUsingDELETEParams, opts ...ClientOption) (*KeyDeleteUsingDELETEOK, error)

	KeyExportUsingPOST(params *KeyExportUsingPOSTParams, opts ...ClientOption) (*KeyExportUsingPOSTOK, error)

	KeyGetUsingGET(params *KeyGetUsingGETParams, opts ...ClientOption) (*KeyGetUsingGETOK, error)

	KeyImportUsingPOST(params *KeyImportUsingPOSTParams, opts ...ClientOption) (*KeyImportUsingPOSTOK, *KeyImportUsingPOSTCreated, error)

	KeyListUsingGET(params *KeyListUsingGETParams, opts ...ClientOption) (*KeyListUsingGETOK, error)

	KeyNewVersionUsingPOST(params *KeyNewVersionUsingPOSTParams, opts ...ClientOption) (*KeyNewVersionUsingPOSTOK, *KeyNewVersionUsingPOSTCreated, error)

	KeyRevokeUsingPOST(params *KeyRevokeUsingPOSTParams, opts ...ClientOption) (*KeyRevokeUsingPOSTOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
KeyCreateUsingPUT creates

Create key (KMS)
*/
func (a *Client) KeyCreateUsingPUT(params *KeyCreateUsingPUTParams, opts ...ClientOption) (*KeyCreateUsingPUTOK, *KeyCreateUsingPUTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKeyCreateUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "keyCreateUsingPUT",
		Method:             "PUT",
		PathPattern:        "/kms/key",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KeyCreateUsingPUTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *KeyCreateUsingPUTOK:
		return value, nil, nil
	case *KeyCreateUsingPUTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for k_m_s_key_management: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
KeyDeleteUsingDELETE deletes

Delete key (KMS)
*/
func (a *Client) KeyDeleteUsingDELETE(params *KeyDeleteUsingDELETEParams, opts ...ClientOption) (*KeyDeleteUsingDELETEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKeyDeleteUsingDELETEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "keyDeleteUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/kms/key/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KeyDeleteUsingDELETEReader{formats: a.formats},
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
	success, ok := result.(*KeyDeleteUsingDELETEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for keyDeleteUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
KeyExportUsingPOST exports

Export key (KMS)
*/
func (a *Client) KeyExportUsingPOST(params *KeyExportUsingPOSTParams, opts ...ClientOption) (*KeyExportUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKeyExportUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "keyExportUsingPOST",
		Method:             "POST",
		PathPattern:        "/kms/key/{id}/export",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KeyExportUsingPOSTReader{formats: a.formats},
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
	success, ok := result.(*KeyExportUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for keyExportUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
KeyGetUsingGET gets

Get KMS key
*/
func (a *Client) KeyGetUsingGET(params *KeyGetUsingGETParams, opts ...ClientOption) (*KeyGetUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKeyGetUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "keyGetUsingGET",
		Method:             "GET",
		PathPattern:        "/kms/key/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KeyGetUsingGETReader{formats: a.formats},
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
	success, ok := result.(*KeyGetUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for keyGetUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
KeyImportUsingPOST imports

Import key (KMS)
*/
func (a *Client) KeyImportUsingPOST(params *KeyImportUsingPOSTParams, opts ...ClientOption) (*KeyImportUsingPOSTOK, *KeyImportUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKeyImportUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "keyImportUsingPOST",
		Method:             "POST",
		PathPattern:        "/kms/key/import",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KeyImportUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *KeyImportUsingPOSTOK:
		return value, nil, nil
	case *KeyImportUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for k_m_s_key_management: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
KeyListUsingGET lists

List KMS keys
*/
func (a *Client) KeyListUsingGET(params *KeyListUsingGETParams, opts ...ClientOption) (*KeyListUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKeyListUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "keyListUsingGET",
		Method:             "GET",
		PathPattern:        "/kms/key",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KeyListUsingGETReader{formats: a.formats},
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
	success, ok := result.(*KeyListUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for keyListUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
KeyNewVersionUsingPOST news version

Create new key version (KMS)
*/
func (a *Client) KeyNewVersionUsingPOST(params *KeyNewVersionUsingPOSTParams, opts ...ClientOption) (*KeyNewVersionUsingPOSTOK, *KeyNewVersionUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKeyNewVersionUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "keyNewVersionUsingPOST",
		Method:             "POST",
		PathPattern:        "/kms/key/{id}/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KeyNewVersionUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *KeyNewVersionUsingPOSTOK:
		return value, nil, nil
	case *KeyNewVersionUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for k_m_s_key_management: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
KeyRevokeUsingPOST revokes

Revoke key (KMS)
*/
func (a *Client) KeyRevokeUsingPOST(params *KeyRevokeUsingPOSTParams, opts ...ClientOption) (*KeyRevokeUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKeyRevokeUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "keyRevokeUsingPOST",
		Method:             "POST",
		PathPattern:        "/kms/key/{id}/revoke",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KeyRevokeUsingPOSTReader{formats: a.formats},
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
	success, ok := result.(*KeyRevokeUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for keyRevokeUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
