package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/default_services"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type ServicesProxy struct {
	httpClient *http.Client
	service    default_services.ClientService
}

func (p *ServicesProxy) Read(ctx context.Context, serviceID strfmt.UUID) (*models.ServiceInstance, error) {
	params := &default_services.GetNetworkDefaultServicesServiceIDParams{
		ServiceID:  serviceID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkDefaultServicesServiceID(params)
	mutex.Unlock()

	if err != nil {
		var notFound *default_services.GetNetworkDefaultServicesServiceIDNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading service: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving service failed: %s", response.Payload.Messages)
	}

	return response.Payload.ServiceInstance, nil
}

func (p *ServicesProxy) ListByDisplayName(ctx context.Context, displayName string) ([]*models.ServiceInstance, error) {
	params := &default_services.GetNetworkDefaultServicesParams{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkDefaultServices(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing services: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing services failed: %s", response.Payload.Messages)
	}

	return response.Payload.ServiceInstanceCollection, nil
}

func (p *ServicesProxy) ListServices(ctx context.Context) ([]*models.ServiceInstance, error) {
	params := &default_services.GetNetworkDefaultServicesParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.GetNetworkDefaultServices(params)

	if err != nil {
		return nil, fmt.Errorf("error while listing services: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing services failed: %s", response.Payload.Messages)
	}

	return response.Payload.ServiceInstanceCollection, nil
}
