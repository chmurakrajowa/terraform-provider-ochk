package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"net/http"
	"sync"
	"time"
)

type PublicIPAddressProxy struct {
	httpClient *http.Client
	service    *openapi.PublicIpAPIService
}

func (p *PublicIPAddressProxy) Get(ctx context.Context, allocationId int32) (*openapi.PublicIpAllocation, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.IpamIpaddressPublicAllocationAllocationIdGet(ctx, allocationId)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while get ipam allocated ip: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing ipam allocated ip failed: %s", response.Messages)
	}

	return response.PublicIpAllocation, nil
}

func (p *PublicIPAddressProxy) List(ctx context.Context) ([]openapi.PublicIpAllocation, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.IpamIpaddressPublicAllocationGet(ctx)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while list allocated public ip addresses: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing allocated public ip addresses failed: %s", response.Messages)
	}

	return response.PublicIpAllocationCollection, nil
}

func (p *PublicIPAddressProxy) ListByName(ctx context.Context, allocationName string) ([]openapi.PublicIpAllocation, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.IpamIpaddressPublicAllocationGet(ctx).Name(allocationName)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while list allocated public ip addresses: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing allocated public ip addresses failed: %s", response.Messages)
	}

	return response.PublicIpAllocationCollection, nil
}

func (p *PublicIPAddressProxy) ListByIp(ctx context.Context, ipAddress string) ([]openapi.PublicIpAllocation, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.IpamIpaddressPublicAllocationGet(ctx).IpAddress(ipAddress)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while get allocated public ip addresse: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing allocated public ip addresses failed: %s", response.Messages)
	}

	return response.PublicIpAllocationCollection, nil
}

func (p *PublicIPAddressProxy) Create(ctx context.Context, publicIPAllocation openapi.PublicIpAllocation, timeout time.Duration) (*openapi.RequestInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.IpamIpaddressPublicAllocationPut(ctx).PublicIpAllocation(publicIPAllocation)
	put, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating public ip address allocation: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("creating public ip allocation failed: %s", put.Messages)
	}

	return put.RequestInstance, nil
}

func (p *PublicIPAddressProxy) Update(ctx context.Context, publicIPAllocation openapi.PublicIpAllocation) (*openapi.RequestInstance, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.IpamIpaddressPublicAllocationAllocationIdPut(ctx, publicIPAllocation.GetAllocationId()).PublicIpAllocation(publicIPAllocation)
	put, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifying public ip allocation: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("modifying public ip allocation failed: %s", put.Messages)
	}

	return put.RequestInstance, nil
}

func (p *PublicIPAddressProxy) Delete(ctx context.Context, publicIPAllocationID int32) (*openapi.RequestInstance, error) {

	action := p.service.IpamIpaddressPublicAllocationAllocationIdDelete(ctx, publicIPAllocationID)
	response, _, err := action.Execute()

	if err != nil {
		return nil, fmt.Errorf("error while deleting public ip allocation: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("deleting public ip allocation failed: %s", response.Messages)
	}

	return response.RequestInstance, nil
}
