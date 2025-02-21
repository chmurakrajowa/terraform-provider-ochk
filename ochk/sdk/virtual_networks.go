package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/virtual_network"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type VirtualNetworksProxy struct {
	httpClient *http.Client
	service    virtual_network.ClientService
}

func (p *VirtualNetworksProxy) Create(ctx context.Context, virtualNetwork *models.VirtualNetworkInstance) (*models.RequestInstance, error) {
	if err := virtualNetwork.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating virtual network struct: %w", err)
	}

	params := &virtual_network.PutNetworksParams{
		Body:       virtualNetwork,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutNetworks(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating virtual network: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating virtual network failed: %s", put.Payload.Messages)
	}

	return put.Payload.RequestInstance, nil
}

func (p *VirtualNetworksProxy) Update(ctx context.Context, virtualNetwork *models.VirtualNetworkInstance) (*models.RequestInstance, error) {
	if err := virtualNetwork.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating virtual network struct: %w", err)
	}

	params := &virtual_network.PutNetworksVirtualNetworkIDParams{
		VirtualNetworkID: virtualNetwork.VirtualNetworkID,
		Body:             virtualNetwork,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutNetworksVirtualNetworkID(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifying virtual network: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying virtual network failed: %s", put.Payload.Messages)
	}

	return put.Payload.RequestInstance, nil
}

func (p *VirtualNetworksProxy) Read(ctx context.Context, virtualNetworkID strfmt.UUID) (*models.VirtualNetworkInstance, error) {
	if virtualNetworkID == "" {
		return nil, fmt.Errorf("empty virtual network ID")
	}

	params := &virtual_network.GetNetworksVirtualNetworkIDParams{
		VirtualNetworkID: virtualNetworkID,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworksVirtualNetworkID(params)
	mutex.Unlock()

	if err != nil {
		var notFound *virtual_network.GetNetworksVirtualNetworkIDNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading virtual network: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving virtual network failed: %s", response.Payload.Messages)
	}

	return response.Payload.VirtualNetworkInstance, nil
}

func (p *VirtualNetworksProxy) ListByDisplayName(ctx context.Context, displayName string) ([]*models.VirtualNetworkInstance, error) {
	params := &virtual_network.GetNetworksParams{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworks(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing virtual networks: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing virtual networks failed: %s", response.Payload.Messages)
	}

	return response.Payload.VirtualNetworkInstanceCollection, nil
}

func (p *VirtualNetworksProxy) List(ctx context.Context) ([]*models.VirtualNetworkInstance, error) {
	params := &virtual_network.GetNetworksParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworks(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing virtual networks: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing virtual networks failed: %s", response.Payload.Messages)
	}

	return response.Payload.VirtualNetworkInstanceCollection, nil
}

func (p *VirtualNetworksProxy) Delete(ctx context.Context, virtualNetworkID strfmt.UUID) (*models.RequestInstance, error) {
	params := &virtual_network.DeleteNetworksVirtualNetworkIDParams{
		VirtualNetworkID: virtualNetworkID,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	response, err := p.service.DeleteNetworksVirtualNetworkID(params)

	if err != nil {
		var badRequest *virtual_network.DeleteNetworksVirtualNetworkIDBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while deleting virtual network: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("deleting virtual network failed: %s", response.Payload.Messages)
	}

	return response.Payload.RequestInstance, nil
}
