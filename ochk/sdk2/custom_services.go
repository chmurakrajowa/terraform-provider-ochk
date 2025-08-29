package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type CustomServicesProxy struct {
	httpClient *http.Client
	service    openapi.CustomServicesAPIService
}

func (p *CustomServicesProxy) Create(ctx context.Context, customService openapi.CustomServiceInstance) (*openapi.CustomServiceInstance, error) {
	//if err := customService.Validate(strfmt.Default); err != nil {
	//	return nil, fmt.Errorf("error while validating custom service struct: %w", err)
	//}

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkCustomServicesPut(ctx).CustomServiceInstance(customService)
	put, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating custom service: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("creating custom service failed: %s", put.Messages)
	}

	return put.CustomServiceInstance, nil
}

func (p *CustomServicesProxy) Update(ctx context.Context, customService *openapi.CustomServiceInstance) (*openapi.CustomServiceInstance, error) {
	//if err := customService.Validate(strfmt.Default); err != nil {
	//	return nil, fmt.Errorf("error while validating custom service struct: %w", err)
	//}

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkCustomServicesServiceIdPut(ctx, customService.GetServiceId())
	put, _, err := action.Execute()

	mutex.Unlock()
	if err != nil {
		return nil, fmt.Errorf("error while modifying custom service: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("modifying custom service failed: %s", put.Messages)
	}

	return put.CustomServiceInstance, nil
}

func (p *CustomServicesProxy) Read(ctx context.Context, customServiceID strfmt.UUID) (*openapi.CustomServiceInstance, error) {

	action := p.service.NetworkCustomServicesServiceIdGet(ctx, string(customServiceID))
	response, _, err := action.Execute()

	if err != nil {
		return nil, fmt.Errorf("error while reading custom service: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("retrieving custom service failed: %s", response.Messages)
	}

	return response.CustomServiceInstance, nil
}

func (p *CustomServicesProxy) ListByDisplayName(ctx context.Context, displayName string) ([]openapi.CustomServiceInstance, error) {

	action := p.service.NetworkCustomServicesGet(ctx).DisplayName(displayName)
	response, _, err := action.Execute()

	if err != nil {
		return nil, fmt.Errorf("error while listing custom services: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing custom services failed: %s", response.Messages)
	}

	return response.CustomServiceInstanceCollection, nil
}

func (p *CustomServicesProxy) ListCustomServices(ctx context.Context) ([]openapi.CustomServiceInstance, error) {

	action := p.service.NetworkCustomServicesGet(ctx)

	response, _, err := action.Execute()

	if err != nil {
		return nil, fmt.Errorf("error while listing custom services: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing custom services failed: %s", response.Messages)
	}

	return response.CustomServiceInstanceCollection, nil
}

func (p *CustomServicesProxy) Exists(ctx context.Context, customServiceID strfmt.UUID) (bool, error) {
	if _, err := p.Read(ctx, customServiceID); err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}

		return false, fmt.Errorf("error while reading custom service: %w", err)
	}

	return true, nil
}

func (p *CustomServicesProxy) Delete(ctx context.Context, customServiceID strfmt.UUID) error {

	action := p.service.NetworkCustomServicesServiceIdDelete(ctx, string(customServiceID))
	response, _, err := action.Execute()

	if err != nil {
		return fmt.Errorf("error while deleting custom service: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return fmt.Errorf("deleting custom service failed: %s", response.Messages)
	}

	return nil
}
