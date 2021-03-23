package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/virtual_networks"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"time"
)

type VirtualNetworksProxy struct {
	httpClient *http.Client
	service    virtual_networks.ClientService
}

func (p *VirtualNetworksProxy) Create(ctx context.Context, virtualNetwork *models.VirtualNetworkInstance, timeout time.Duration) (*models.RequestInstance, error) {
	if err := virtualNetwork.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating virtual network struct: %w", err)
	}

	params := &virtual_networks.VirtualNetworkCreateUsingPUTParams{
		VirtualNetworkInstance: virtualNetwork,
		Context:                ctx,
		HTTPClient:             p.httpClient,
	}

	_, put, err := p.service.VirtualNetworkCreateUsingPUT(params)
	if err != nil {
		return nil, fmt.Errorf("error while creating virtua network: %w", err)
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

	params := &virtual_networks.VirtualNetworkUpdateUsingPUTParams{
		VirtualNetworkID:       virtualNetwork.VirtualNetworkID,
		VirtualNetworkInstance: virtualNetwork,
		Context:                ctx,
		HTTPClient:             p.httpClient,
	}

	put, _, err := p.service.VirtualNetworkUpdateUsingPUT(params)
	if err != nil {
		return nil, fmt.Errorf("error while modifying virtual network: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying virtual network failed: %s", put.Payload.Messages)
	}

	return put.Payload.RequestInstance, nil
}

func (p *VirtualNetworksProxy) Read(ctx context.Context, virtualNetworkID string) (*models.VirtualNetworkInstance, error) {
	if virtualNetworkID == "" {
		return nil, fmt.Errorf("empty virtual network ID")
	}

	params := &virtual_networks.VirtualNetworkGetUsingGETParams{
		VirtualNetworkID: virtualNetworkID,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	response, err := p.service.VirtualNetworkGetUsingGET(params)
	if err != nil {
		var notFound *virtual_networks.VirtualNetworkGetUsingGETNotFound
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
	params := &virtual_networks.VirtualNetworkListUsingGETParams{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	response, err := p.service.VirtualNetworkListUsingGET(params)
	if err != nil {
		return nil, fmt.Errorf("error while listing virtual networks: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing virtual networks failed: %s", response.Payload.Messages)
	}

	return response.Payload.VirtualNetworkInstanceCollection, nil
}

func (p *VirtualNetworksProxy) Delete(ctx context.Context, virtualNetworkID string) (*models.RequestInstance, error) {
	params := &virtual_networks.VirtualNetworkDeleteUsingDELETEParams{
		VirtualNetworkID: virtualNetworkID,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	response, _, err := p.service.VirtualNetworkDeleteUsingDELETE(params)
	if err != nil {
		var badRequest *virtual_networks.VirtualNetworkDeleteUsingDELETEBadRequest
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
