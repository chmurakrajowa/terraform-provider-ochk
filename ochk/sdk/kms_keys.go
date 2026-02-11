package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"net/http"
	"sync"
)

type KMSKeysProxy struct {
	httpClient *http.Client
	service    *openapi.KeyAPIService
}

func (p *KMSKeysProxy) Create(ctx context.Context, keyInstance openapi.KeyInstance) (*openapi.KeyInstance, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.KmsKeyPut(ctx).KeyInstance(keyInstance)
	post, _, err := action.Execute()
	mutex.Unlock()
	if err != nil {
		return nil, fmt.Errorf("error while creating KMS key: %w", err)
	}

	isSuccess := *post.Success

	if !isSuccess {
		return nil, fmt.Errorf("creating KMS key failed: %s", post.Messages)
	}

	return post.KeyInstance, nil
}

func (p *KMSKeysProxy) Import(ctx context.Context, keyImport openapi.KeyImport) (*openapi.KeyInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.KmsKeyImportPost(ctx).KeyImport(keyImport)
	post, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while importing KMS key: %w", err)
	}
	isSuccess := *post.Success

	if !isSuccess {
		return nil, fmt.Errorf("importing KMS key failed: %s", post.Messages)
	}

	return post.KeyInstance, nil
}

func (p *KMSKeysProxy) Read(ctx context.Context, keyID string) (*openapi.KeyInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.KmsKeyIdGet(ctx, keyID)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading KMS key: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("retrieving KMS key failed: %s", response.Messages)
	}

	return response.KeyInstance, nil
}

func (p *KMSKeysProxy) ListByDisplayName(ctx context.Context, displayName string) ([]openapi.KeyInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.KmsKeyGet(ctx).DisplayName(displayName)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing KMS keys: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing KMS keys failed: %s", response.Messages)
	}

	return response.KeyInstanceCollection, nil
}

func (p *KMSKeysProxy) List(ctx context.Context) ([]openapi.KeyInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.KmsKeyGet(ctx)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing KMS keys: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing KMS keys failed: %s", response.Messages)
	}

	return response.KeyInstanceCollection, nil
}

func (p *KMSKeysProxy) Delete(ctx context.Context, keyID string) error {

	action := p.service.KmsKeyIdDelete(ctx, keyID)
	response, _, err := action.Execute()

	if err != nil {
		return fmt.Errorf("error while deleting KMS key: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return fmt.Errorf("deleting KMS key failed: %s", response.Messages)
	}

	return nil
}
