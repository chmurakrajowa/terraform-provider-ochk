package sdk

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
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
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("listing ipam allocated ip failed: %s", message)
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
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("listing allocated public ip addresses failed: %s", message)
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
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("listing allocated public ip addresses failed: %s", message)
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
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("listing allocated public ip addresses failed: %s", message)
	}

	return response.PublicIpAllocationCollection, nil
}

func (p *PublicIPAddressProxy) Create(ctx context.Context, publicIPAllocation openapi.PublicIpAllocation) (*openapi.PublicIpAllocation, error) {
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
		message := "no error message"
		if value := put.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("creating public ip allocation failed: %s", message)
	}

	return put.PublicIpAllocation, nil
}

func (p *PublicIPAddressProxy) Update(ctx context.Context, publicIPAllocation openapi.PublicIpAllocation) (*openapi.PublicIpAllocation, error) {

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
		message := "no error message"
		if value := put.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("modifying public ip allocation failed: %s", message)
	}

	return put.PublicIpAllocation, nil
}

func (p *PublicIPAddressProxy) Delete(ctx context.Context, publicIPAllocationID int32) error {

	action := p.service.IpamIpaddressPublicAllocationAllocationIdDelete(ctx, publicIPAllocationID)
	response, _, err := action.Execute()

	if err != nil {
		return fmt.Errorf("error while deleting public ip allocation: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return fmt.Errorf("deleting public ip allocation failed: %s", message)
	}

	return nil
}

func (p *PublicIPAddressProxy) Read(ctx context.Context, publicIPAllocationID int32) (*openapi.PublicIpAllocation, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.IpamIpaddressPublicAllocationAllocationIdGet(ctx, publicIPAllocationID)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading public ip adresses: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("retrieving public ip address: failed: %s", message)
	}
	if response.PublicIpAllocation != nil {
		return response.PublicIpAllocation, nil
	} else {
		return nil, nil
	}
}
