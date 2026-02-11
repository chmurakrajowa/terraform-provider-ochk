package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type FloatingIPAddressProxy struct {
	httpClient *http.Client
	service    *openapi.FloatingIpAPIService
}

func (p *FloatingIPAddressProxy) Read(ctx context.Context, floating_ip_id strfmt.UUID) (*openapi.FloatingIp, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkFloatingIpsFloatingIpIdGet(ctx, string(floating_ip_id))
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading floating ip adresses: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("retrieving floating ip adress: failed: %s", response.Messages)
	}
	if response.FloatingIp != nil {
		return response.FloatingIp, nil
	} else {
		return nil, nil
	}
}

func (p *FloatingIPAddressProxy) List(ctx context.Context) ([]openapi.FloatingIp, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkFloatingIpsGet(ctx)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing floating ip adresses: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing floating ip adress failed: %s", response.Messages)
	}

	return response.FloatingIpCollection, nil
}

func (p *FloatingIPAddressProxy) ListByName(ctx context.Context, name string) ([]openapi.FloatingIp, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkFloatingIpsGet(ctx).Name(name)
	response, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing floating ip adresses: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing floating ip adress failed: %s", response.Messages)
	}

	return response.FloatingIpCollection, nil
}

func (p *FloatingIPAddressProxy) Create(ctx context.Context, floatingIPAllocation openapi.FloatingIp) (*openapi.FloatingIp, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkFloatingIpsPut(ctx).FloatingIp(floatingIPAllocation)
	put, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating floating ip address allocation: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("creating floating ip allocation failed: %s", put.Messages)
	}

	return put.FloatingIp, nil
}

func (p *FloatingIPAddressProxy) Update(ctx context.Context, floatingIp openapi.FloatingIp) (*openapi.FloatingIp, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkFloatingIpsFloatingIpIdPut(ctx, floatingIp.GetFloatingIpId()).FloatingIp(floatingIp)
	put, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while updating floating ip: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("modifying floating ip failed: %s", put.Messages)
	}

	return put.FloatingIp, nil
}

func (p *FloatingIPAddressProxy) Delete(ctx context.Context, floatingIpID strfmt.UUID) error {

	action := p.service.NetworkFloatingIpsFloatingIpIdDelete(ctx, string(floatingIpID))
	response, _, err := action.Execute()

	if err != nil {
		return fmt.Errorf("error while deleting firewall rule: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return fmt.Errorf("deleting floating ip failed: %s", response.Messages)
	}

	return nil
}
