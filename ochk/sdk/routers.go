package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type RoutersProxy struct {
	httpClient *http.Client
	service    *openapi.RouterAPIService
}

func (p *RoutersProxy) Read(ctx context.Context, routerID strfmt.UUID) (*openapi.RouterInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkRoutersRouterIdGet(ctx, string(routerID))
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading routers: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("retrieving routers failed: %s", response.Messages)
	}

	return response.RouterInstance, nil
}

func (p *RoutersProxy) ListByDisplayName(ctx context.Context, displayName string) ([]openapi.RouterInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkRoutersGet(ctx).DisplayName(displayName)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing routers: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing routers failed: %s", response.Messages)
	}

	return response.RouterCollection, nil
}

func (p *RoutersProxy) List(ctx context.Context) ([]openapi.RouterInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkRoutersGet(ctx)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing routers: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing routers failed: %s", response.Messages)
	}

	return response.RouterCollection, nil
}

func (p *RoutersProxy) Create(ctx context.Context, Router openapi.RouterInstance) (*openapi.RouterInstance, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkRoutersPut(ctx).RouterInstance(Router)
	put, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating router: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("creating router failed: %s", put.Messages)
	}

	return put.RouterInstance, nil
}

func (p *RoutersProxy) Update(ctx context.Context, Router openapi.RouterInstance) (*openapi.RouterInstance, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkRoutersRouterIdPut(ctx, Router.GetRouterId()).RouterInstance(Router)
	put, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifying router: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("modifying router failed: %s", put.Messages)
	}

	return put.RouterInstance, nil
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

	action := p.service.NetworkRoutersRouterIdDelete(ctx, string(RouterID))
	response, _, err := action.Execute()

	if err != nil {
		return fmt.Errorf("error while deleting router: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return fmt.Errorf("deleting router failed: %s", response.Messages)
	}

	return nil
}
