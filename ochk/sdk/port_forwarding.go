package sdk

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
)

type PortsForwardingProxy struct {
	httpClient *http.Client
	service    *openapi.PortForwardingAPIService
}

func (p *PortsForwardingProxy) Read(ctx context.Context, floatingIpId strfmt.UUID, portForwardingId strfmt.UUID) (*openapi.PortForwarding, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkFloatingIpsFloatingIpIdPortForwardingsPortForwardingIdGet(ctx, string(floatingIpId), string(portForwardingId))
	response, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading port forwarding: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("retrieving port forwarding: failed: %s", message)
	}
	if response.PortForwarding != nil {
		return response.PortForwarding, nil
	} else {
		return nil, nil
	}
}

func (p *PortsForwardingProxy) List(ctx context.Context, floatingIpId strfmt.UUID) ([]openapi.PortForwarding, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkFloatingIpsFloatingIpIdPortForwardingsGet(ctx, string(floatingIpId))
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing ports forwarding: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("listing ports forwarding failed: %s", message)
	}

	return response.PortForwardingCollection, nil
}

func (p *PortsForwardingProxy) ListByName(ctx context.Context, floatingIpId strfmt.UUID, name string) ([]openapi.PortForwarding, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkFloatingIpsFloatingIpIdPortForwardingsGet(ctx, string(floatingIpId)).Name(name)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing ports forwarding: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("listing ports forwarding failed: %s", message)
	}

	return response.PortForwardingCollection, nil
}

func (p *PortsForwardingProxy) Create(ctx context.Context, floatingIpId strfmt.UUID, portForwarding openapi.PortForwarding) (*openapi.PortForwarding, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkFloatingIpsFloatingIpIdPortForwardingsPut(ctx, string(floatingIpId)).PortForwarding(portForwarding)
	put, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating port forwarding: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		message := "no error message"
		if value := put.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("creating port forwarding failed: %s", message)
	}

	return put.PortForwarding, nil
}

func (p *PortsForwardingProxy) Update(ctx context.Context, floatingIpId strfmt.UUID, portForwarding openapi.PortForwarding) (*openapi.PortForwarding, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkFloatingIpsFloatingIpIdPortForwardingsPortForwardingIdPut(ctx, string(floatingIpId), portForwarding.GetPortForwardingId()).PortForwarding(portForwarding)
	put, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while updating port forwarding: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		message := "no error message"
		if value := put.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("creating updating port forwarding: %s", message)
	}

	return put.PortForwarding, nil
}

func (p *PortsForwardingProxy) Delete(ctx context.Context, floatingIpId strfmt.UUID, portForwardingId strfmt.UUID) error {

	action := p.service.NetworkFloatingIpsFloatingIpIdPortForwardingsPortForwardingIdDelete(ctx, string(floatingIpId), string(portForwardingId))
	response, _, err := action.Execute()

	if err != nil {
		return fmt.Errorf("error while deleting port forwarding: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return fmt.Errorf("deleting port forwarding failed: %s", message)
	}

	return nil
}
