package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/tags"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"net/http"
)

type TagsProxy struct {
	httpClient *http.Client
	service    tags.ClientService
}

func (p *TagsProxy) Read(ctx context.Context, tagID int32) (*models.Tag, error) {
	params := &tags.GetTagsTagIDParams{
		TagID:      tagID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}
	response, err := p.service.GetTagsTagID(params)
	if err != nil {
		var notFound *tags.GetTagsTagIDNotFound

		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading tags: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving tags failed: %s", response.Payload.Messages)
	}

	return response.Payload.Tag, nil
}

func (p *TagsProxy) ListTagsByTagName(ctx context.Context, tagName string) ([]*models.Tag, error) {
	params := &tags.GetTagsParams{
		Name:       &tagName,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.GetTags(params)

	if err != nil {
		return nil, fmt.Errorf("error while listing tags: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing tags failed: %s", response.Payload.Messages)
	}

	return response.Payload.TagCollection, nil
}

func (p *TagsProxy) ListTags(ctx context.Context) ([]*models.Tag, error) {
	params := &tags.GetTagsParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.GetTags(params)

	if err != nil {
		return nil, fmt.Errorf("error while listing tags: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing tags failed: %s", response.Payload.Messages)
	}

	return response.Payload.TagCollection, nil
}

func (p *TagsProxy) Create(ctx context.Context, tag *models.Tag) (*models.Tag, error) {
	if err := tag.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating tag struct: %w", err)
	}

	params := &tags.PutTagsParams{
		Body:       tag,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	put, err := p.service.PutTags(params)

	if err != nil {
		return nil, fmt.Errorf("error while creating tag: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating tag failed: %s", put.Payload.Messages)
	}

	return put.Payload.Tag, nil
}

func (p *TagsProxy) Update(ctx context.Context, tag *models.Tag) (*models.Tag, error) {
	if err := tag.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating tag struct: %w", err)
	}

	params := &tags.PutTagsTagIDParams{
		TagID:      tag.TagID,
		Body:       tag,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	put, err := p.service.PutTagsTagID(params)

	if err != nil {
		return nil, fmt.Errorf("error while modifying tag: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying tag failed: %s", put.Payload.Messages)
	}

	return put.Payload.Tag, nil
}

func (p *TagsProxy) Exists(ctx context.Context, tagID int32) (bool, error) {
	if _, err := p.Read(ctx, tagID); err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}

		return false, fmt.Errorf("error while reading tag: %w", err)
	}

	return true, nil
}

func (p *TagsProxy) Delete(ctx context.Context, tagID string) error {

	var tagIDInt32 int32
	_, err := fmt.Sscan(tagID, &tagIDInt32)
	if err != nil {
		return fmt.Errorf("wrong tag_id format: %w", err)
	}

	params := &tags.DeleteTagsTagIDParams{
		TagID:      tagIDInt32,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.DeleteTagsTagID(params)
	if err != nil {
		var badRequest *tags.DeleteTagsTagIDBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return &NotFoundError{Err: err}
		}

		return fmt.Errorf("error while deleting tag: %w", err)
	}

	if !response.Payload.Success {
		return fmt.Errorf("deleting tag failed: %s", response.Payload.Messages)
	}

	return nil
}
