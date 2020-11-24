package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/custom_services"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/go-openapi/strfmt"
	"net/http"
)

type CustomServicesProxy struct {
	httpClient *http.Client
	service    custom_services.ClientService
}

func (p *CustomServicesProxy) Create(ctx context.Context, customService *models.CustomServiceInstance) (*models.CustomServiceInstance, error) {
	if err := customService.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating custom service struct: %w", err)
	}

	params := &custom_services.CustomServiceCreateUsingPUTParams{
		CustomServiceInstance: customService,
		Context:               ctx,
		HTTPClient:            p.httpClient,
	}

	_, put, err := p.service.CustomServiceCreateUsingPUT(params)
	if err != nil {
		return nil, fmt.Errorf("error while creating custom service: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating custom service failed: %s", put.Payload.Messages)
	}

	return put.Payload.CustomServiceInstance, nil
}

func (p *CustomServicesProxy) Update(ctx context.Context, customService *models.CustomServiceInstance) (*models.CustomServiceInstance, error) {
	if err := customService.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating custom service struct: %w", err)
	}

	params := &custom_services.CustomServiceUpdateUsingPUTParams{
		ServiceID:             customService.ServiceID,
		CustomServiceInstance: customService,
		Context:               ctx,
		HTTPClient:            p.httpClient,
	}

	put, _, err := p.service.CustomServiceUpdateUsingPUT(params)
	if err != nil {
		return nil, fmt.Errorf("error while modifying custom service: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying custom service failed: %s", put.Payload.Messages)
	}

	return put.Payload.CustomServiceInstance, nil
}

func (p *CustomServicesProxy) Read(ctx context.Context, customServiceID string) (*models.CustomServiceInstance, error) {
	params := &custom_services.CustomServiceGetUsingGETParams{
		ServiceID:  customServiceID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.CustomServiceGetUsingGET(params)
	if err != nil {
		var notFound *custom_services.CustomServiceGetUsingGETNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading custom service: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving custom service failed: %s", response.Payload.Messages)
	}

	return response.Payload.CustomServiceInstance, nil
}

func (p *CustomServicesProxy) ListByDisplayName(ctx context.Context, displayName string) ([]*models.CustomServiceInstance, error) {
	params := &custom_services.CustomServiceListUsingGETParams{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	response, err := p.service.CustomServiceListUsingGET(params)
	if err != nil {
		return nil, fmt.Errorf("error while listing custom services: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing custom services failed: %s", response.Payload.Messages)
	}

	return response.Payload.CustomServiceInstanceCollection, nil
}

func (p *CustomServicesProxy) Exists(ctx context.Context, customServiceID string) (bool, error) {
	if _, err := p.Read(ctx, customServiceID); err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}

		return false, fmt.Errorf("error while reading custom service: %w", err)
	}

	return true, nil
}

func (p *CustomServicesProxy) Delete(ctx context.Context, customServiceID string) error {
	params := &custom_services.CustomServiceDeleteUsingDELETEParams{
		ServiceID:  customServiceID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, _, err := p.service.CustomServiceDeleteUsingDELETE(params)
	if err != nil {
		var badRequest *custom_services.CustomServiceDeleteUsingDELETEBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return &NotFoundError{Err: err}
		}

		return fmt.Errorf("error while deleting custom service: %w", err)
	}

	if !response.Payload.Success {
		return fmt.Errorf("deleting custom service failed: %s", response.Payload.Messages)
	}

	return nil
}
