package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/k_m_s_key_management"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"net/http"
)

type KMSKeysProxy struct {
	httpClient *http.Client
	service    k_m_s_key_management.ClientService
}

func (p *KMSKeysProxy) Create(ctx context.Context, keyInstance *models.KeyInstance) (*models.KeyInstance, error) {
	params := &k_m_s_key_management.KeyCreateUsingPUTParams{
		KeyInstance: keyInstance,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	_, post, err := p.service.KeyCreateUsingPUT(params)
	if err != nil {
		return nil, fmt.Errorf("error while creating KMS key: %w", err)
	}

	if !post.Payload.Success {
		return nil, fmt.Errorf("creating KMS key failed: %s", post.Payload.Messages)
	}

	return post.Payload.KeyInstance, nil
}

func (p *KMSKeysProxy) Import(ctx context.Context, keyImport *models.KeyImport) (*models.KeyInstance, error) {
	params := &k_m_s_key_management.KeyImportUsingPOSTParams{
		KeyImport:  keyImport,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	_, post, err := p.service.KeyImportUsingPOST(params)
	if err != nil {
		return nil, fmt.Errorf("error while importing KMS key: %w", err)
	}

	if !post.Payload.Success {
		return nil, fmt.Errorf("importing KMS key failed: %s", post.Payload.Messages)
	}

	return post.Payload.KeyInstance, nil
}

func (p *KMSKeysProxy) Read(ctx context.Context, keyID string) (*models.KeyInstance, error) {
	params := &k_m_s_key_management.KeyGetUsingGETParams{
		ID:         keyID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.KeyGetUsingGET(params)
	if err != nil {
		var notFound *k_m_s_key_management.KeyGetUsingGETNotFound
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
	params := &k_m_s_key_management.KeyListUsingGETParams{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	response, err := p.service.KeyListUsingGET(params)
	if err != nil {
		return nil, fmt.Errorf("error while listing KMS keys: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing KMS keys failed: %s", response.Payload.Messages)
	}

	return response.Payload.KeyInstanceCollection, nil
}

func (p *KMSKeysProxy) Delete(ctx context.Context, keyID string) error {
	params := &k_m_s_key_management.KeyDeleteUsingDELETEParams{
		ID:         keyID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.KeyDeleteUsingDELETE(params)
	if err != nil {
		var badRequest *k_m_s_key_management.KeyDeleteUsingDELETEBadRequest
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
