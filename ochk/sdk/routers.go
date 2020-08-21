package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/routers"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/go-openapi/strfmt"
	"net/http"
)

type RoutersProxy struct {
	httpClient *http.Client
	service    routers.ClientService
}

func (p *RoutersProxy) Read(ctx context.Context, routerID string) (*models.RouterInstance, error) {
	params := &routers.RouterGetUsingGETParams{
		RouterID:   routerID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.RouterGetUsingGET(params)
	if err != nil {
		var notFound *routers.RouterGetUsingGETNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading routers: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving routers failed: %s", response.Payload.Messages)
	}

	return response.Payload.RouterInstance, nil
}

func (p *RoutersProxy) ListByDisplayName(ctx context.Context, displayName string) ([]*models.RouterInstance, error) {
	params := &routers.RouterListUsingGETParams{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	response, err := p.service.RouterListUsingGET(params)
	if err != nil {
		return nil, fmt.Errorf("error while listing routers: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing routers failed: %s", response.Payload.Messages)
	}

	return response.Payload.RouterCollection, nil
}

func (p *RoutersProxy) Create(ctx context.Context, Router *models.RouterInstance) (*models.RouterInstance, error) {
	if err := Router.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating router struct: %w", err)
	}

	params := &routers.RouterCreateUsingPUTParams{
		RouterInstance: Router,
		Context:        ctx,
		HTTPClient:     p.httpClient,
	}

	_, put, err := p.service.RouterCreateUsingPUT(params)
	if err != nil {
		return nil, fmt.Errorf("error while creating router: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating router failed: %s", put.Payload.Messages)
	}

	return put.Payload.RouterInstance, nil
}

func (p *RoutersProxy) Update(ctx context.Context, Router *models.RouterInstance) (*models.RouterInstance, error) {
	if err := Router.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating router struct: %w", err)
	}

	params := &routers.RouterUpdateUsingPUTParams{
		RouterID:       Router.RouterID,
		RouterInstance: Router,
		Context:        ctx,
		HTTPClient:     p.httpClient,
	}

	put, _, err := p.service.RouterUpdateUsingPUT(params)
	if err != nil {
		return nil, fmt.Errorf("error while modifying router: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying router failed: %s", put.Payload.Messages)
	}

	return put.Payload.RouterInstance, nil
}

func (p *RoutersProxy) Exists(ctx context.Context, RouterID string) (bool, error) {
	if _, err := p.Read(ctx, RouterID); err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}

		return false, fmt.Errorf("error while reading router: %w", err)
	}

	return true, nil
}

func (p *RoutersProxy) Delete(ctx context.Context, RouterID string) error {
	params := &routers.RouterDeleteUsingDELETEParams{
		RouterID:   RouterID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, _, err := p.service.RouterDeleteUsingDELETE(params)
	if err != nil {
		var badRequest *routers.RouterDeleteUsingDELETEBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return &NotFoundError{Err: err}
		}

		return fmt.Errorf("error while deleting router: %w", err)
	}

	if !response.Payload.Success {
		return fmt.Errorf("deleting router failed: %s", response.Payload.Messages)
	}

	return nil
}
