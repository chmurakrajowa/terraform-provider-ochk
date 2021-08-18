package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/system_tags"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/go-openapi/strfmt"
	"net/http"
)

type SystemTagsProxy struct {
	httpClient *http.Client
	service    system_tags.ClientService
}

func (p *SystemTagsProxy) Read(ctx context.Context, systemTagID int32) (*models.SystemTag, error) {
	params := &system_tags.SystemTagGetUsingGETParams{
		SystemTagID: systemTagID,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}
	response, err := p.service.SystemTagGetUsingGET(params)
	if err != nil {
		var notFound *system_tags.SystemTagGetUsingGETNotFound

		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading system tags: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving system tags failed: %s", response.Payload.Messages)
	}

	return response.Payload.SystemTag, nil
}

func (p *SystemTagsProxy) ListSystemTagsByTagName(ctx context.Context, systemTagName string) ([]*models.SystemTag, error) {
	params := &system_tags.SystemTagListUsingGETParams{
		TagValue:   &systemTagName,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.SystemTagListUsingGET(params)

	if err != nil {
		return nil, fmt.Errorf("error while listing system tags: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("Listing system tags failed: %s", response.Payload.Messages)
	}

	return response.Payload.SystemTagCollection, nil
}

func (p *SystemTagsProxy) Create(ctx context.Context, SystemTag *models.SystemTag) (*models.SystemTag, error) {
	if err := SystemTag.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating system tag struct: %w", err)
	}

	params := &system_tags.SystemTagCreateUsingPUTParams{
		SystemTag:  SystemTag,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	_, put, err := p.service.SystemTagCreateUsingPUT(params)

	if err != nil {
		return nil, fmt.Errorf("error while creating system tag: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating system tag failed: %s", put.Payload.Messages)
	}

	return put.Payload.SystemTag, nil
}

func (p *SystemTagsProxy) Update(ctx context.Context, SystemTag *models.SystemTag) (*models.SystemTag, error) {
	if err := SystemTag.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating system tag struct: %w", err)
	}

	params := &system_tags.SystemTagUpdateUsingPUTParams{
		SystemTagID: SystemTag.SystemTagID,
		SystemTag:   SystemTag,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	put, err := p.service.SystemTagUpdateUsingPUT(params)

	if err != nil {
		return nil, fmt.Errorf("error while modifying system tag: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying system tag failed: %s", put.Payload.Messages)
	}

	return put.Payload.SystemTag, nil
}

func (p *SystemTagsProxy) Exists(ctx context.Context, SystemTagID int32) (bool, error) {
	if _, err := p.Read(ctx, SystemTagID); err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}

		return false, fmt.Errorf("error while reading system tag: %w", err)
	}

	return true, nil
}

func (p *SystemTagsProxy) Delete(ctx context.Context, SystemTagID string) error {

	var SystemTagIDInt32 int32
	fmt.Sscan(SystemTagID, &SystemTagIDInt32)

	params := &system_tags.SystemTagDeleteUsingDELETEParams{
		SystemTagID: SystemTagIDInt32,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	response, err := p.service.SystemTagDeleteUsingDELETE(params)
	if err != nil {
		var badRequest *system_tags.SystemTagDeleteUsingDELETEBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return &NotFoundError{Err: err}
		}

		return fmt.Errorf("error while deleting system tag: %w", err)
	}

	if !response.Payload.Success {
		return fmt.Errorf("deleting system tag failed: %s", response.Payload.Messages)
	}

	return nil
}
