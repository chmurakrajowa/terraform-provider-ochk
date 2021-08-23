package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/backups"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"net/http"
)

type BackupListsProxy struct {
	httpClient *http.Client
	service    backups.ClientService
}

func (p *BackupListsProxy) Read(ctx context.Context, backupPlanID string, backupListID string) (*models.BackupList, error) {
	params := &backups.BackupListGetUsingGETParams{
		BackupListID: backupListID,
		BackupPlanID: backupPlanID,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}
	response, err := p.service.BackupListGetUsingGET(params)

	if err != nil {
		var notFound *backups.BackupListGetUsingGETNotFound

		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading backup list: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving backup list failed: %s", response.Payload.Messages)
	}

	return response.Payload.BackupList, nil
}

func (p *BackupListsProxy) ListBackupListByName(ctx context.Context, backupPlanID string, backupListName string) ([]*models.BackupList, error) {
	params := &backups.BackupListListUsingGETParams{
		BackupPlanID:   backupPlanID,
		BackupListName: &backupListName,
		Context:        ctx,
		HTTPClient:     p.httpClient,
	}

	response, err := p.service.BackupListListUsingGET(params)

	if err != nil {
		return nil, fmt.Errorf("error while listing backup list: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("Listing backup list failed: %s", response.Payload.Messages)
	}

	return response.Payload.BackupListCollection, nil
}
