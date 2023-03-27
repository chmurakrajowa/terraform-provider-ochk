package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/folder"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"net/http"
	"sync"
)

type FoldersProxy struct {
	httpClient *http.Client
	service    folder.ClientService
}

func (p *FoldersProxy) Read(ctx context.Context, projectID string) (*models.FolderInstance, error) {
	params := &folder.FolderGetUsingGETParams{
		ProjectID:  projectID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.FolderGetUsingGET(params)
	mutex.Unlock()

	if err != nil {
		var notFound *folder.FolderGetUsingGETNotFound

		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading folders: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving folders failed: %s", response.Payload.Messages)
	}

	return response.Payload.FolderInstance, nil
}

func (p *FoldersProxy) LisFoldersByProjectId(ctx context.Context, projectID string) ([]*models.FolderInstance, error) {
	params := &folder.FolderListUsingGETParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
		ProjectID:  projectID,
	}

	response, err := p.service.FolderListUsingGET(params)

	if err != nil {
		return nil, fmt.Errorf("error while listing folders: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("Listing folders failed: %s", response.Payload.Messages)
	}

	return response.Payload.FolderInstanceCollection, nil
}
