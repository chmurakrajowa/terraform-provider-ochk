package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/ip_a_m_public_ip_allocations"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
	"time"
)

type PublicIPAddressProxy struct {
	httpClient *http.Client
	service    ip_a_m_public_ip_allocations.ClientService
}

func (p *PublicIPAddressProxy) Get(ctx context.Context, allocationId int32) (*models.PublicIPAllocation, error) {
	params := &ip_a_m_public_ip_allocations.AllocationGetUsingGETParams{
		AllocationID: allocationId,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.AllocationGetUsingGET(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while get ipam allocated ip: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing ipam allocated ip failed: %s", response.Payload.Messages)
	}

	return response.Payload.PublicIPAllocation, nil
}

func (p *PublicIPAddressProxy) List(ctx context.Context) ([]*models.PublicIPAllocation, error) {
	params := &ip_a_m_public_ip_allocations.AllocationListUsingGETParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.AllocationListUsingGET(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while list allocated public ip addresses: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing allocated public ip addresses failed: %s", response.Payload.Messages)
	}

	return response.Payload.PublicIPAllocationCollection, nil
}

func (p *PublicIPAddressProxy) ListByName(ctx context.Context, allocationName string) ([]*models.PublicIPAllocation, error) {
	params := &ip_a_m_public_ip_allocations.AllocationListUsingGETParams{
		Name:       &allocationName,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.AllocationListUsingGET(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while list allocated public ip addresses: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing allocated public ip addresses failed: %s", response.Payload.Messages)
	}

	return response.Payload.PublicIPAllocationCollection, nil
}

func (p *PublicIPAddressProxy) ListByIp(ctx context.Context, ipAddress string) ([]*models.PublicIPAllocation, error) {
	params := &ip_a_m_public_ip_allocations.AllocationListUsingGETParams{
		IPAddress:  &ipAddress,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.AllocationListUsingGET(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while get allocated public ip addresse: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing allocated public ip addresses failed: %s", response.Payload.Messages)
	}

	return response.Payload.PublicIPAllocationCollection, nil
}

func (p *PublicIPAddressProxy) Create(ctx context.Context, publicaIPAllocation *models.PublicIPAllocation, timeout time.Duration) (*models.RequestInstance, error) {
	if err := publicaIPAllocation.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating public ip address struct: %w", err)
	}

	params := &ip_a_m_public_ip_allocations.AllocationCreateUsingPUTParams{
		PublicIPAllocation: publicaIPAllocation,
		Context:            ctx,
		HTTPClient:         p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, _, err := p.service.AllocationCreateUsingPUT(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating public ip address allocation: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating public ip allocation failed: %s", put.Payload.Messages)
	}

	return put.Payload.RequestInstance, nil
}

func (p *PublicIPAddressProxy) Update(ctx context.Context, publicaIPAllocation *models.PublicIPAllocation) (*models.RequestInstance, error) {
	if err := publicaIPAllocation.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating public ip allocation struct: %w", err)
	}

	params := &ip_a_m_public_ip_allocations.AllocationUpdateUsingPUTParams{
		AllocationID:       publicaIPAllocation.AllocationID,
		PublicIPAllocation: publicaIPAllocation,
		Context:            ctx,
		HTTPClient:         p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.AllocationUpdateUsingPUT(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifying public ip allocation: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying public ip allocation failed: %s", put.Payload.Messages)
	}

	return put.Payload.RequestInstance, nil
}

func (p *PublicIPAddressProxy) Delete(ctx context.Context, publicaIPAllocationID int32) (*models.RequestInstance, error) {
	params := &ip_a_m_public_ip_allocations.AllocationDeleteUsingDELETEParams{
		AllocationID: publicaIPAllocationID,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}

	response, err := p.service.AllocationDeleteUsingDELETE(params)

	if err != nil {
		var badRequest *ip_a_m_public_ip_allocations.AllocationDeleteUsingDELETEBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while deleting public ip allocation: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("deleting public ip allocation failed: %s", response.Payload.Messages)
	}

	return response.Payload.RequestInstance, nil
}
