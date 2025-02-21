package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/router"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type RoutersProxy struct {
	httpClient *http.Client
	service    router.ClientService
}

func (p *RoutersProxy) Read(ctx context.Context, routerID strfmt.UUID) (*models.RouterInstance, error) {
	params := &router.GetNetworkRoutersRouterIDParams{
		RouterID:   routerID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkRoutersRouterID(params)
	mutex.Unlock()

	if err != nil {
		var notFound *router.GetNetworkRoutersRouterIDNotFound
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
	params := &router.GetNetworkRoutersParams{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkRouters(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing routers: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing routers failed: %s", response.Payload.Messages)
	}

	return response.Payload.RouterCollection, nil
}

func (p *RoutersProxy) List(ctx context.Context) ([]*models.RouterInstance, error) {
	params := &router.GetNetworkRoutersParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkRouters(params)
	mutex.Unlock()

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

	params := &router.PutNetworkRoutersParams{
		Body:       Router,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutNetworkRouters(params)
	mutex.Unlock()

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

	params := &router.PutNetworkRoutersRouterIDParams{
		RouterID:   Router.RouterID,
		Body:       Router,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutNetworkRoutersRouterID(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifying router: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying router failed: %s", put.Payload.Messages)
	}

	return put.Payload.RouterInstance, nil
}

func (p *RoutersProxy) Exists(ctx context.Context, RouterID strfmt.UUID) (bool, error) {
	if _, err := p.Read(ctx, RouterID); err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}

		return false, fmt.Errorf("error while reading router: %w", err)
	}

	return true, nil
}

func (p *RoutersProxy) Delete(ctx context.Context, RouterID strfmt.UUID) error {
	params := &router.DeleteNetworkRoutersRouterIDParams{
		RouterID:   RouterID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.DeleteNetworkRoutersRouterID(params)

	if err != nil {
		var badRequest *router.DeleteNetworkRoutersRouterIDBadRequest
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
