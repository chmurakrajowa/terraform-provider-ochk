package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/port_forwarding"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type PortsForwardingProxy struct {
	httpClient *http.Client
	service    port_forwarding.ClientService
}

func (p *PortsForwardingProxy) Read(ctx context.Context, floatingIpId strfmt.UUID, portForwardingId strfmt.UUID) (*models.PortForwarding, error) {
	params := &port_forwarding.GetNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDParams{
		FloatingIPID:     floatingIpId,
		PortForwardingID: portForwardingId,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingID(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading port forwarding: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving port forwarding: failed: %s", response.Payload.Messages)
	}
	if response.Payload.PortForwarding != nil {
		return response.Payload.PortForwarding, nil
	} else {
		return nil, nil
	}
}

func (p *PortsForwardingProxy) List(ctx context.Context, floatingIpId strfmt.UUID) ([]*models.PortForwarding, error) {

	params := &port_forwarding.GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams{
		FloatingIPID: floatingIpId,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkFloatingIpsFloatingIPIDPortForwardings(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing ports forwarding: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing ports forwarding failed: %s", response.Payload.Messages)
	}

	return response.Payload.PortForwardingCollection, nil
}

func (p *PortsForwardingProxy) ListByName(ctx context.Context, floatingIpId strfmt.UUID, name string) ([]*models.PortForwarding, error) {
	params := &port_forwarding.GetNetworkFloatingIpsFloatingIPIDPortForwardingsParams{
		FloatingIPID: floatingIpId,
		Name:         &name,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkFloatingIpsFloatingIPIDPortForwardings(params)
	mutex.Unlock()

	//if true {
	//	return nil, fmt.Errorf("<<<<<<<<>>>>>>>>: %w")
	//}

	if err != nil {
		return nil, fmt.Errorf("error while listing ports forwarding: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing ports forwarding failed: %s", response.Payload.Messages)
	}

	return response.Payload.PortForwardingCollection, nil
}

func (p *PortsForwardingProxy) Create(ctx context.Context, floatingIpId strfmt.UUID, portForwarding *models.PortForwarding) (*models.PortForwarding, error) {
	if err := portForwarding.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating port forwarding struct: %w", err)
	}

	params := &port_forwarding.PutNetworkFloatingIpsFloatingIPIDPortForwardingsParams{
		FloatingIPID: floatingIpId,
		Body:         portForwarding,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutNetworkFloatingIpsFloatingIPIDPortForwardings(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating port forwarding: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating port forwarding failed: %s", put.Payload.Messages)
	}

	return put.Payload.PortForwarding, nil
}

func (p *PortsForwardingProxy) Update(ctx context.Context, floatingIpId strfmt.UUID, portForwarding *models.PortForwarding) (*models.PortForwarding, error) {
	if err := portForwarding.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating port forwarding struct: %w", err)
	}

	params := &port_forwarding.PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDParams{
		FloatingIPID:     floatingIpId,
		PortForwardingID: portForwarding.PortForwardingID,
		Body:             portForwarding,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingID(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while updating port forwarding: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating updating port forwarding: %s", put.Payload.Messages)
	}

	return put.Payload.PortForwarding, nil
}

func (p *PortsForwardingProxy) Delete(ctx context.Context, floatingIpId strfmt.UUID, portForwardingId strfmt.UUID) error {
	params := &port_forwarding.DeleteNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDParams{
		FloatingIPID:     floatingIpId,
		PortForwardingID: portForwardingId,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	response, err := p.service.DeleteNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingID(params)

	if err != nil {
		var badRequest *port_forwarding.DeleteNetworkFloatingIpsFloatingIPIDPortForwardingsPortForwardingIDBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return &NotFoundError{Err: err}
		}

		return fmt.Errorf("error while deleting port forwarding: %w", err)
	}

	if !response.Payload.Success {
		return fmt.Errorf("deleting port forwarding failed: %s", response.Payload.Messages)
	}

	return nil
}
