package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/subtenants"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/go-openapi/strfmt"
	"net/http"
)

type SubtenantsProxy struct {
	httpClient *http.Client
	service    subtenants.ClientService
}

func (p *SubtenantsProxy) Create(ctx context.Context, subtenant *models.SubtenantInstance) (*models.SubtenantInstance, error) {
	if err := subtenant.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating subtenant struct: %w", err)
	}

	params := &subtenants.SubtenantCreateUsingPUTParams{
		SubtenantInstance: subtenant,
		Context:           ctx,
		HTTPClient:        p.httpClient,
	}

	_, put, err := p.service.SubtenantCreateUsingPUT(params)
	if err != nil {
		return nil, fmt.Errorf("error while creating subtenant: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating subtenant failed: %s", put.Payload.Messages)
	}

	return put.Payload.SubtenantInstance, nil
}

func (p *SubtenantsProxy) Update(ctx context.Context, subtenant *models.SubtenantInstance) (*models.SubtenantInstance, error) {
	if err := subtenant.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating subtenant struct: %w", err)
	}

	params := &subtenants.SubtenantUpdateUsingPUTParams{
		SubtenantID:       subtenant.SubtenantID,
		SubtenantInstance: subtenant,
		Context:           ctx,
		HTTPClient:        p.httpClient,
	}

	put, err := p.service.SubtenantUpdateUsingPUT(params)
	if err != nil {
		return nil, fmt.Errorf("error while modifying subtenant: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying subtenant failed: %s", put.Payload.Messages)
	}

	return put.Payload.SubtenantInstance, nil
}

func (p *SubtenantsProxy) Read(ctx context.Context, subtenantID string) (*models.SubtenantInstance, error) {
	params := &subtenants.SubtenantGetUsingGETParams{
		SubtenantID: subtenantID,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	response, err := p.service.SubtenantGetUsingGET(params)
	if err != nil {
		var notFound *subtenants.SubtenantGetUsingGETNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading subtenant: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving subtenant failed: %s", response.Payload.Messages)
	}

	return response.Payload.SubtenantInstance, nil
}

func (p *SubtenantsProxy) ListByDisplayName(ctx context.Context, name string) ([]*models.SubtenantInstance, error) {
	params := &subtenants.SubtenantListUsingGETParams{
		Name:       &name,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.SubtenantListUsingGET(params)
	if err != nil {
		return nil, fmt.Errorf("error while listing subtenants: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing subtenants failed: %s", response.Payload.Messages)
	}

	return response.Payload.SubtenantInstanceCollection, nil
}

func (p *SubtenantsProxy) Exists(ctx context.Context, subtenantID string) (bool, error) {
	if _, err := p.Read(ctx, subtenantID); err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}

		return false, fmt.Errorf("error while reading subtenant: %w", err)
	}

	return true, nil
}

func (p *SubtenantsProxy) Delete(ctx context.Context, subtenantID string) error {
	params := &subtenants.SubtenantDeleteUsingDELETEParams{
		SubtenantID: subtenantID,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	response, err := p.service.SubtenantDeleteUsingDELETE(params)
	if err != nil {
		var badRequest *subtenants.SubtenantDeleteUsingDELETEBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return &NotFoundError{Err: err}
		}

		return fmt.Errorf("error while deleting subtenant: %w", err)
	}

	if !response.Payload.Success {
		return fmt.Errorf("deleting subtenant failed: %s", response.Payload.Messages)
	}

	return nil
}
