package sdk

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
)

type FoldersProxy struct {
	httpClient *http.Client
	service    *openapi.FolderAPIService
}

func (p *FoldersProxy) Read(ctx context.Context, projectID strfmt.UUID) (*openapi.FolderInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.FolderProjectIdGet(ctx, string(projectID))
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading folders: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("retrieving folders failed: %s", message)
	}

	return &response.FolderInstanceCollection[0], nil
}

func (p *FoldersProxy) LisFoldersByProjectId(ctx context.Context, projectID strfmt.UUID) ([]openapi.FolderInstance, error) {

	action := p.service.FolderProjectIdGet(ctx, string(projectID))
	response, _, err := action.Execute()

	if err != nil {
		return nil, fmt.Errorf("error while listing folders: %w", err)
	}

	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("listing folders failed: %s", message)
	}

	return response.FolderInstanceCollection, nil
}
