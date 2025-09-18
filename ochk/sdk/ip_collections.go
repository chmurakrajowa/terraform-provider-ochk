package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type IPCollectionsProxy struct {
	httpClient *http.Client
	service    *openapi.IpCollectionAPIService
}

func (p *IPCollectionsProxy) Read(ctx context.Context, ipCollectionID strfmt.UUID) (*openapi.IpCollection, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.IpcsIpCollectionIdGet(ctx, string(ipCollectionID))
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading IP collection: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("retrieving IP collection failed: %s", response.Messages)
	}

	return response.IpCollection, nil
}

func (p *IPCollectionsProxy) ListByDisplayName(ctx context.Context, displayName string) ([]openapi.IpCollection, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.IpcsGet(ctx).Name(displayName)
	response, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing IP collections by display name %s: %w", displayName, err)
	}

	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing IP collections by display name %s failed: %s", displayName, response.Messages)
	}

	return response.IpCollectionSet, nil
}

func (p *IPCollectionsProxy) List(ctx context.Context) ([]openapi.IpCollection, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.IpcsGet(ctx)
	response, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing IP collections: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing IP collections failed: %s", response.Messages)
	}

	return response.IpCollectionSet, nil
}

func (p *IPCollectionsProxy) Create(ctx context.Context, IPCollection openapi.IpCollection) (*openapi.IpCollection, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.IpcsPut(ctx).IpCollection(IPCollection)
	put, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating ip collection: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("creating ip collection failed: %s", put.Messages)
	}

	return put.IpCollection, nil
}

func (p *IPCollectionsProxy) Update(ctx context.Context, IPCollection openapi.IpCollection) (*openapi.IpCollection, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.IpcsIpCollectionIdPut(ctx, IPCollection.GetId()).IpCollection(IPCollection)
	put, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifying ip collection: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("modifying ip collection failed: %s", put.Messages)
	}

	return put.IpCollection, nil
}

func (p *IPCollectionsProxy) Exists(ctx context.Context, IPCollectionID strfmt.UUID) (bool, error) {
	if _, err := p.Read(ctx, IPCollectionID); err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}

		return false, fmt.Errorf("error while reading ip collection: %w", err)
	}

	return true, nil
}

func (p *IPCollectionsProxy) Delete(ctx context.Context, IPCollectionID strfmt.UUID) error {

	action := p.service.IpcsIpCollectionIdDelete(ctx, string(IPCollectionID))
	response, _, err := action.Execute()
	if err != nil {
		return fmt.Errorf("error while deleting ip collection: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return fmt.Errorf("deleting ip collection failed: %s", response.Messages)
	}

	return nil
}
