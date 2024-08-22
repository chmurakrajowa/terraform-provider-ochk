package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/folder"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type FoldersProxy struct {
	httpClient *http.Client
	service    folder.ClientService
}

func (p *FoldersProxy) Read(ctx context.Context, projectID strfmt.UUID) (*models.FolderInstance, error) {
	params := &folder.GetFolderProjectIDIDParams{
		ProjectID:  projectID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetFolderProjectIDID(params)
	mutex.Unlock()

	if err != nil {
		var notFound *folder.GetFolderProjectIDIDNotFound

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

func (p *FoldersProxy) LisFoldersByProjectId(ctx context.Context, projectID strfmt.UUID) ([]*models.FolderInstance, error) {
	params := &folder.GetFolderProjectIDParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
		ProjectID:  projectID,
	}

	response, err := p.service.GetFolderProjectID(params)

	if err != nil {
		return nil, fmt.Errorf("error while listing folders: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing folders failed: %s", response.Payload.Messages)
	}

	return response.Payload.FolderInstanceCollection, nil
}
