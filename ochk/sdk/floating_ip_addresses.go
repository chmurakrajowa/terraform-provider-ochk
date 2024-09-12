package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/floating_ip"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type FloatingIPAddressProxy struct {
	httpClient *http.Client
	service    floating_ip.ClientService
}

func (p *FloatingIPAddressProxy) Read(ctx context.Context, floating_ip_id strfmt.UUID) (*models.FloatingIP, error) {
	params := &floating_ip.GetNetworkFloatingIpsFloatingIPIDParams{
		FloatingIPID: floating_ip_id,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkFloatingIpsFloatingIPID(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading floating ip adresses: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving floating ip adress: failed: %s", response.Payload.Messages)
	}
	if response.Payload.FloatingIP != nil {
		return response.Payload.FloatingIP, nil
	} else {
		return nil, nil
	}
}

func (p *FloatingIPAddressProxy) List(ctx context.Context) ([]*models.FloatingIP, error) {

	params := &floating_ip.GetNetworkFloatingIpsParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkFloatingIps(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing floating ip adresses: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing floating ip adress failed: %s", response.Payload.Messages)
	}

	return response.Payload.FloatingIPCollection, nil
}

func (p *FloatingIPAddressProxy) ListByName(ctx context.Context, name string) ([]*models.FloatingIP, error) {
	params := &floating_ip.GetNetworkFloatingIpsParams{
		Name:       &name,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkFloatingIps(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing floating ip adresses: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing floating ip adress failed: %s", response.Payload.Messages)
	}

	return response.Payload.FloatingIPCollection, nil
}

func (p *FloatingIPAddressProxy) Create(ctx context.Context, floatingIPAllocation *models.FloatingIP) (*models.FloatingIP, error) {
	if err := floatingIPAllocation.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating floating ip address struct: %w", err)
	}

	params := &floating_ip.PutNetworkFloatingIpsParams{
		Body:       floatingIPAllocation,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutNetworkFloatingIps(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating floating ip address allocation: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating floating ip allocation failed: %s", put.Payload.Messages)
	}

	return put.Payload.FloatingIP, nil
}

func (p *FloatingIPAddressProxy) Update(ctx context.Context, floatingIp *models.FloatingIP) (*models.FloatingIP, error) {
	if err := floatingIp.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating floating ip struct: %w", err)
	}

	params := &floating_ip.PutNetworkFloatingIpsFloatingIPIDParams{
		FloatingIPID: floatingIp.FloatingIPID,
		Body:         floatingIp,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutNetworkFloatingIpsFloatingIPID(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while updating floating ip: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying floating ip failed: %s", put.Payload.Messages)
	}

	return put.Payload.FloatingIP, nil
}

func (p *FloatingIPAddressProxy) Delete(ctx context.Context, floatingIpID strfmt.UUID) error {
	params := &floating_ip.DeleteNetworkFloatingIpsFloatingIPIDParams{
		FloatingIPID: floatingIpID,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}

	response, err := p.service.DeleteNetworkFloatingIpsFloatingIPID(params)

	if err != nil {
		var badRequest *floating_ip.DeleteNetworkFloatingIpsFloatingIPIDBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return &NotFoundError{Err: err}
		}

		return fmt.Errorf("error while deleting firewall rule: %w", err)
	}

	if !response.Payload.Success {
		return fmt.Errorf("deleting floating ip failed: %s", response.Payload.Messages)
	}

	return nil
}
