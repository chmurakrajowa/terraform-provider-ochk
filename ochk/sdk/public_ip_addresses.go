package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/public_ip"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
	"time"
)

type PublicIPAddressProxy struct {
	httpClient *http.Client
	service    public_ip.ClientService
}

func (p *PublicIPAddressProxy) Get(ctx context.Context, allocationId int32) (*models.PublicIPAllocation, error) {
	params := &public_ip.GetIpamIpaddressPublicAllocationAllocationIDParams{
		AllocationID: allocationId,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetIpamIpaddressPublicAllocationAllocationID(params)
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
	params := &public_ip.GetIpamIpaddressPublicAllocationParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetIpamIpaddressPublicAllocation(params)
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
	params := &public_ip.GetIpamIpaddressPublicAllocationParams{
		Name:       &allocationName,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetIpamIpaddressPublicAllocation(params)
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
	params := &public_ip.GetIpamIpaddressPublicAllocationParams{
		IPAddress:  &ipAddress,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetIpamIpaddressPublicAllocation(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while get allocated public ip addresse: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing allocated public ip addresses failed: %s", response.Payload.Messages)
	}

	return response.Payload.PublicIPAllocationCollection, nil
}

func (p *PublicIPAddressProxy) Create(ctx context.Context, publicIPAllocation *models.PublicIPAllocation, timeout time.Duration) (*models.RequestInstance, error) {
	if err := publicIPAllocation.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating public ip address struct: %w", err)
	}

	params := &public_ip.PutIpamIpaddressPublicAllocationParams{
		Body:       publicIPAllocation,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutIpamIpaddressPublicAllocation(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating public ip address allocation: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating public ip allocation failed: %s", put.Payload.Messages)
	}

	return put.Payload.RequestInstance, nil
}

func (p *PublicIPAddressProxy) Update(ctx context.Context, publicIPAllocation *models.PublicIPAllocation) (*models.RequestInstance, error) {
	if err := publicIPAllocation.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating public ip allocation struct: %w", err)
	}

	params := &public_ip.PutIpamIpaddressPublicAllocationAllocationIDParams{
		AllocationID: publicIPAllocation.AllocationID,
		Body:         publicIPAllocation,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutIpamIpaddressPublicAllocationAllocationID(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifying public ip allocation: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying public ip allocation failed: %s", put.Payload.Messages)
	}

	return put.Payload.RequestInstance, nil
}

func (p *PublicIPAddressProxy) Delete(ctx context.Context, publicIPAllocationID int32) (*models.RequestInstance, error) {
	params := &public_ip.DeleteIpamIpaddressPublicAllocationAllocationIDParams{
		AllocationID: publicIPAllocationID,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}

	response, err := p.service.DeleteIpamIpaddressPublicAllocationAllocationID(params)

	if err != nil {
		var badRequest *public_ip.GetIpamIpaddressPublicAllocationBadRequest
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
