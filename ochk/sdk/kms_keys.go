package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/key"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"net/http"
	"sync"
)

type KMSKeysProxy struct {
	httpClient *http.Client
	service    key.ClientService
}

func (p *KMSKeysProxy) Create(ctx context.Context, keyInstance *models.KeyInstance) (*models.KeyInstance, error) {
	params := &key.PutKmsKeyParams{
		Body:       keyInstance,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	post, err := p.service.PutKmsKey(params)
	mutex.Unlock()
	if err != nil {
		return nil, fmt.Errorf("error while creating KMS key: %w", err)
	}

	if !post.Payload.Success {
		return nil, fmt.Errorf("creating KMS key failed: %s", post.Payload.Messages)
	}

	return post.Payload.KeyInstance, nil
}

func (p *KMSKeysProxy) Import(ctx context.Context, keyImport *models.KeyImport) (*models.KeyInstance, error) {
	params := &key.PostKmsKeyImportParams{
		Body:       keyImport,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	post, err := p.service.PostKmsKeyImport(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while importing KMS key: %w", err)
	}

	if !post.Payload.Success {
		return nil, fmt.Errorf("importing KMS key failed: %s", post.Payload.Messages)
	}

	return post.Payload.KeyInstance, nil
}

func (p *KMSKeysProxy) Read(ctx context.Context, keyID string) (*models.KeyInstance, error) {
	params := &key.GetKmsKeyIDParams{
		ID:         keyID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetKmsKeyID(params)
	mutex.Unlock()

	if err != nil {
		var notFound *key.GetKmsKeyIDNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading KMS key: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving KMS key failed: %s", response.Payload.Messages)
	}

	return response.Payload.KeyInstance, nil
}

func (p *KMSKeysProxy) ListByDisplayName(ctx context.Context, displayName string) ([]*models.KeyInstance, error) {
	params := &key.GetKmsKeyParams{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetKmsKey(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing KMS keys: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing KMS keys failed: %s", response.Payload.Messages)
	}

	return response.Payload.KeyInstanceCollection, nil
}

func (p *KMSKeysProxy) List(ctx context.Context) ([]*models.KeyInstance, error) {
	params := &key.GetKmsKeyParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetKmsKey(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing KMS keys: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing KMS keys failed: %s", response.Payload.Messages)
	}

	return response.Payload.KeyInstanceCollection, nil
}

func (p *KMSKeysProxy) Delete(ctx context.Context, keyID string) error {
	params := &key.DeleteKmsKeyIDParams{
		ID:         keyID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.DeleteKmsKeyID(params)

	if err != nil {
		var badRequest *key.DeleteKmsKeyIDBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return &NotFoundError{Err: err}
		}

		return fmt.Errorf("error while deleting KMS key: %w", err)
	}

	if !response.Payload.Success {
		return fmt.Errorf("deleting KMS key failed: %s", response.Payload.Messages)
	}

	return nil
}
