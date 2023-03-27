package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/default_services"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"net/http"
	"sync"
)

type ServicesProxy struct {
	httpClient *http.Client
	service    default_services.ClientService
}

func (p *ServicesProxy) Read(ctx context.Context, serviceID string) (*models.ServiceInstance, error) {
	params := &default_services.ServiceGetUsingGETParams{
		ServiceID:  serviceID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.ServiceGetUsingGET(params)
	mutex.Unlock()

	if err != nil {
		var notFound *default_services.ServiceGetUsingGETNotFound
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
	params := &default_services.ServiceListUsingGETParams{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.ServiceListUsingGET(params)
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
	params := &default_services.ServiceListUsingGETParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.ServiceListUsingGET(params)

	if err != nil {
		return nil, fmt.Errorf("error while listing services: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing services failed: %s", response.Payload.Messages)
	}

	return response.Payload.ServiceInstanceCollection, nil
}
