package sdk

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
)

type ServicesProxy struct {
	httpClient *http.Client
	service    *openapi.DefaultServicesAPIService
}

func (p *ServicesProxy) Read(ctx context.Context, serviceID strfmt.UUID) (*openapi.ServiceInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkDefaultServicesServiceIdGet(ctx, string(serviceID))
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading service: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("retrieving service failed: %s", message)
	}

	return response.ServiceInstance, nil
}

func (p *ServicesProxy) ListByDisplayName(ctx context.Context, displayName string) ([]openapi.ServiceInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkDefaultServicesGet(ctx).DisplayName(displayName)
	response, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing services: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("listing services failed: %s", message)
	}

	return response.ServiceInstanceCollection, nil
}

func (p *ServicesProxy) ListServices(ctx context.Context) ([]openapi.ServiceInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkDefaultServicesGet(ctx)
	response, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing services: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("listing services failed: %s", message)
	}

	return response.ServiceInstanceCollection, nil
}
