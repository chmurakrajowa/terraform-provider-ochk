package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/routers"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
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
