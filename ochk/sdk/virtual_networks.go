package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type VirtualNetworksProxy struct {
	httpClient *http.Client
	service    *openapi.VirtualNetworkAPIService
}

func (p *VirtualNetworksProxy) Create(ctx context.Context, virtualNetwork openapi.VirtualNetworkInstance) (*openapi.RequestInstance, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworksPut(ctx).VirtualNetworkInstance(virtualNetwork)

	put, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating virtual network: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("creating virtual network failed: %s", put.Messages)
	}

	return put.RequestInstance, nil
}

func (p *VirtualNetworksProxy) Update(ctx context.Context, virtualNetwork openapi.VirtualNetworkInstance) (*openapi.RequestInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworksVirtualNetworkIdPut(ctx, virtualNetwork.GetVirtualNetworkId()).VirtualNetworkInstance(virtualNetwork)
	put, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifying virtual network: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("modifying virtual network failed: %s", put.Messages)
	}

	return put.RequestInstance, nil
}

func (p *VirtualNetworksProxy) Read(ctx context.Context, virtualNetworkID strfmt.UUID) (*openapi.VirtualNetworkInstance, error) {
	if virtualNetworkID == "" {
		return nil, fmt.Errorf("empty virtual network ID")
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworksVirtualNetworkIdGet(ctx, string(virtualNetworkID))
	put, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {

		return nil, fmt.Errorf("error while reading virtual network: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("retrieving virtual network failed: %s", put.Messages)
	}

	return put.VirtualNetworkInstance, nil
}

func (p *VirtualNetworksProxy) ListByDisplayName(ctx context.Context, displayName string) ([]openapi.VirtualNetworkInstance, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworksGet(ctx).DisplayName(displayName)
	response, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing virtual networks: %w", err)
	}

	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing virtual networks failed: %s", response.Messages)
	}

	return response.VirtualNetworkInstanceCollection, nil
}

func (p *VirtualNetworksProxy) List(ctx context.Context) ([]openapi.VirtualNetworkInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworksGet(ctx)
	response, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing virtual networks: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing virtual networks failed: %s", response.Messages)
	}

	return response.VirtualNetworkInstanceCollection, nil
}

func (p *VirtualNetworksProxy) Delete(ctx context.Context, virtualNetworkID strfmt.UUID) (*openapi.RequestInstance, error) {
	action := p.service.NetworksVirtualNetworkIdDelete(ctx, string(virtualNetworkID))
	response, _, err := action.Execute()

	if err != nil {
		return nil, fmt.Errorf("error while deleting virtual network: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("deleting virtual network failed: %s", response.Messages)
	}

	return response.RequestInstance, nil
}
