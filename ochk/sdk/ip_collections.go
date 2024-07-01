package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/ip_collections"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type IPCollectionsProxy struct {
	httpClient *http.Client
	service    ip_collections.ClientService
}

func (p *IPCollectionsProxy) Read(ctx context.Context, ipCollectionID string) (*models.IPCollection, error) {
	params := &ip_collections.IPCollectionGetUsingGETParams{
		IPCollectionID: ipCollectionID,
		Context:        ctx,
		HTTPClient:     p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.IPCollectionGetUsingGET(params)
	mutex.Unlock()

	if err != nil {
		var notFound *ip_collections.IPCollectionGetUsingGETNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading IP collection: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving IP collection failed: %s", response.Payload.Messages)
	}

	return response.Payload.IPCollection, nil
}

func (p *IPCollectionsProxy) ListByDisplayName(ctx context.Context, displayName string) ([]*models.IPCollection, error) {
	params := &ip_collections.IPCollectionListUsingGETParams{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.IPCollectionListUsingGET(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing IP collections by display name %s: %w", displayName, err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing IP collections by display name %s failed: %s", displayName, response.Payload.Messages)
	}

	return response.Payload.IPCollectionSet, nil
}

func (p *IPCollectionsProxy) List(ctx context.Context) ([]*models.IPCollection, error) {
	params := &ip_collections.IPCollectionListUsingGETParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.IPCollectionListUsingGET(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing IP collections: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing IP collections failed: %s", response.Payload.Messages)
	}

	return response.Payload.IPCollectionSet, nil
}

func (p *IPCollectionsProxy) Create(ctx context.Context, IPCollection *models.IPCollection) (*models.IPCollection, error) {
	if err := IPCollection.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating ip collection struct: %w", err)
	}

	params := &ip_collections.IPCollectionCreateUsingPUTParams{
		IPCollection: IPCollection,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, _, err := p.service.IPCollectionCreateUsingPUT(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating ip collection: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating ip collection failed: %s", put.Payload.Messages)
	}

	return put.Payload.IPCollection, nil
}

func (p *IPCollectionsProxy) Update(ctx context.Context, IPCollection *models.IPCollection) (*models.IPCollection, error) {
	if err := IPCollection.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating ip collection struct: %w", err)
	}

	params := &ip_collections.IPCollectionUpdateUsingPUTParams{
		IPCollectionID: IPCollection.ID,
		IPCollection:   IPCollection,
		Context:        ctx,
		HTTPClient:     p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.IPCollectionUpdateUsingPUT(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifying ip collection: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying ip collection failed: %s", put.Payload.Messages)
	}

	return put.Payload.IPCollection, nil
}

func (p *IPCollectionsProxy) Exists(ctx context.Context, IPCollectionID string) (bool, error) {
	if _, err := p.Read(ctx, IPCollectionID); err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}

		return false, fmt.Errorf("error while reading ip collection: %w", err)
	}

	return true, nil
}

func (p *IPCollectionsProxy) Delete(ctx context.Context, IPCollectionID string) error {
	params := &ip_collections.IPCollectionDeleteUsingDELETEParams{
		IPCollectionID: IPCollectionID,
		Context:        ctx,
		HTTPClient:     p.httpClient,
	}

	response, err := p.service.IPCollectionDeleteUsingDELETE(params)

	if err != nil {
		var badRequest *ip_collections.IPCollectionDeleteUsingDELETEBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return &NotFoundError{Err: err}
		}

		return fmt.Errorf("error while deleting ip collection: %w", err)
	}

	if !response.Payload.Success {
		return fmt.Errorf("deleting ip collection failed: %s", response.Payload.Messages)
	}

	return nil
}
