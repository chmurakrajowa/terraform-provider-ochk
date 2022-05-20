package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/local_groups"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"time"
)

type LocalGroupsProxy struct {
	httpClient *http.Client
	service    local_groups.ClientService
}

func (p *LocalGroupsProxy) Read(ctx context.Context, groupID string) (*models.LocalGroup, error) {
	params := &local_groups.LocalGroupGetUsingGETParams{
		GroupID:    groupID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.LocalGroupGetUsingGET(params)
	if err != nil {
		var notFound *local_groups.LocalGroupGetUsingGETNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading groups: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving groups failed: %s", response.Payload.Messages)
	}

	return response.Payload.LocalGroup, nil
}

func (p *LocalGroupsProxy) ListByName(ctx context.Context, groupName string) ([]*models.LocalGroup, error) {
	params := &local_groups.LocalGroupListUsingGETParams{
		Name:       &groupName,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.LocalGroupListUsingGET(params)

	if err != nil {
		return nil, fmt.Errorf("error while listing users: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing users failed: %s", response.Payload.Messages)
	}

	return response.Payload.LocalGroupInstanceCollection, nil
}

func (p *LocalGroupsProxy) Create(ctx context.Context, localGroup *models.LocalGroup, timeout time.Duration) (*models.RequestInstance, error) {
	if err := localGroup.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating virtual network struct: %w", err)
	}

	params := &local_groups.LocalGroupCreateUsingPUTParams{
		LocalGroup: localGroup,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	put, _, err := p.service.LocalGroupCreateUsingPUT(params)
	if err != nil {
		return nil, fmt.Errorf("error while creating local group: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating local group failed: %s", put.Payload.Messages)
	}

	return put.Payload.RequestInstance, nil
}

func (p *LocalGroupsProxy) Update(ctx context.Context, localGroup *models.LocalGroup) (*models.RequestInstance, error) {
	if err := localGroup.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating local group struct: %w", err)
	}

	params := &local_groups.LocalGroupUpdateUsingPUTParams{
		GroupID:    localGroup.GroupID,
		LocalGroup: localGroup,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	put, err := p.service.LocalGroupUpdateUsingPUT(params)
	if err != nil {
		return nil, fmt.Errorf("error while modifying local group: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying local group failed: %s", put.Payload.Messages)
	}

	return put.Payload.RequestInstance, nil
}

func (p *LocalGroupsProxy) Delete(ctx context.Context, localGroupID string) (*models.RequestInstance, error) {
	params := &local_groups.LocalGroupDeleteUsingDELETEParams{
		GroupID:    localGroupID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.LocalGroupDeleteUsingDELETE(params)
	if err != nil {
		var badRequest *local_groups.LocalGroupDeleteUsingDELETEBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while deleting local group: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("deleting local group failed: %s", response.Payload.Messages)
	}

	return response.Payload.RequestInstance, nil
}
