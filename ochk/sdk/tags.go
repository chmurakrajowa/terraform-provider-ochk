package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"net/http"
)

type TagsProxy struct {
	httpClient *http.Client
	service    *openapi.TagsAPIService
}

func (p *TagsProxy) Read(ctx context.Context, tagID int32) (*openapi.Tag, error) {

	action := p.service.TagsTagIdGet(ctx, tagID)
	response, _, err := action.Execute()
	if err != nil {
		return nil, fmt.Errorf("error while reading tags: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("retrieving tags failed: %s", response.Messages)
	}

	return response.Tag, nil
}

func (p *TagsProxy) ListTagsByTagName(ctx context.Context, tagName string) ([]openapi.Tag, error) {
	action := p.service.TagsGet(ctx).TagValue(tagName)
	response, _, err := action.Execute()

	if err != nil {
		return nil, fmt.Errorf("error while listing tags: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing tags failed: %s", response.Messages)
	}

	return response.TagCollection, nil
}

func (p *TagsProxy) ListTags(ctx context.Context) ([]openapi.Tag, error) {

	action := p.service.TagsGet(ctx)
	response, _, err := action.Execute()

	if err != nil {
		return nil, fmt.Errorf("error while listing tags: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing tags failed: %s", response.Messages)
	}

	return response.TagCollection, nil
}

func (p *TagsProxy) Create(ctx context.Context, tag openapi.Tag) (*openapi.Tag, error) {

	action := p.service.TagsPut(ctx).Tag(tag)
	put, _, err := action.Execute()

	if err != nil {
		return nil, fmt.Errorf("error while creating tag: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("creating tag failed: %s", put.Messages)
	}

	return put.Tag, nil
}

func (p *TagsProxy) Update(ctx context.Context, tag openapi.Tag) (*openapi.Tag, error) {

	action := p.service.TagsTagIdPut(ctx, tag.GetTagId()).Tag(tag)
	put, _, err := action.Execute()

	if err != nil {
		return nil, fmt.Errorf("error while modifying tag: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("modifying tag failed: %s", put.Messages)
	}

	return put.Tag, nil
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

	action := p.service.TagsTagIdDelete(ctx, tagIDInt32)
	response, _, err := action.Execute()

	if err != nil {
		return fmt.Errorf("error while deleting tag: %w", err)
	}

	isSuccess := *response.Success

	if !isSuccess {
		return fmt.Errorf("deleting tag failed: %s", response.Messages)
	}

	return nil
}
