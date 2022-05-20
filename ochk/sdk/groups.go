package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/groups"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"net/http"
)

type GroupsProxy struct {
	httpClient *http.Client
	service    groups.ClientService
}

func (p *GroupsProxy) Read(ctx context.Context, groupID string) (*models.GroupInstance, error) {
	params := &groups.GroupGetUsingGETParams{
		GroupID:    groupID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.GroupGetUsingGET(params)
	if err != nil {
		var notFound *groups.GroupGetUsingGETNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading groups: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving groups failed: %s", response.Payload.Messages)
	}

	return response.Payload.GroupInstance, nil
}

func (p *GroupsProxy) ListByName(ctx context.Context, groupName string) ([]*models.GroupInstance, error) {
	params := &groups.GroupListUsingGETParams{
		Name:       &groupName,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.GroupListUsingGET(params)

	if err != nil {
		return nil, fmt.Errorf("error while listing users: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing users failed: %s", response.Payload.Messages)
	}

	return response.Payload.GroupInstanceCollection, nil
}
