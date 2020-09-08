package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/users"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"net/http"
)

type UsersProxy struct {
	httpClient *http.Client
	service    users.ClientService
}

func (p *UsersProxy) Read(ctx context.Context, userID string) (*models.UserInstance, error) {
	params := &users.UserGetUsingGETParams{
		UserID:     userID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.UserGetUsingGET(params)
	if err != nil {
		var notFound *users.UserGetUsingGETNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading users: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving users failed: %s", response.Payload.Messages)
	}

	return response.Payload.UserInstance, nil
}

func (p *UsersProxy) ListByDisplayName(ctx context.Context, userName string) ([]*models.UserInstance, error) {
	params := &users.UserListUsingGETParams{
		UserName:   &userName,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.UserListUsingGET(params)
	if err != nil {
		return nil, fmt.Errorf("error while listing users: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing users failed: %s", response.Payload.Messages)
	}

	return response.Payload.UserInstanceCollection, nil
}
