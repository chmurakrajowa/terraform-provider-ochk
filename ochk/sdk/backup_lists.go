package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/backups"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type BackupListsProxy struct {
	httpClient *http.Client
	service    backups.ClientService
}

func (p *BackupListsProxy) Read(ctx context.Context, backupPlanID strfmt.UUID, backupListID strfmt.UUID) (*models.BackupList, error) {
	params := &backups.GetBackupsPlansBackupPlanIDListsBackupListIDParams{
		BackupListID: backupListID,
		BackupPlanID: backupPlanID,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}

	response, err := p.service.GetBackupsPlansBackupPlanIDListsBackupListID(params)

	if err != nil {
		var notFound *backups.GetBackupsPlansBackupPlanIDListsBackupListIDNotFound

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

func (p *BackupListsProxy) ListBackupListByName(ctx context.Context, backupPlanID strfmt.UUID, backupListName string) ([]*models.BackupList, error) {
	params := &backups.GetBackupsPlansBackupPlanIDListsParams{
		BackupPlanID:   backupPlanID,
		BackupListName: &backupListName,
		Context:        ctx,
		HTTPClient:     p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetBackupsPlansBackupPlanIDLists(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing backup list: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("Listing backup list failed: %s", response.Payload.Messages)
	}

	return response.Payload.BackupListCollection, nil
}

func (p *BackupListsProxy) ListBackupList(ctx context.Context, backupPlanID strfmt.UUID) ([]*models.BackupList, error) {
	params := &backups.GetBackupsPlansBackupPlanIDListsParams{
		BackupPlanID: backupPlanID,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetBackupsPlansBackupPlanIDLists(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing backup list: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("Listing backup list failed: %s", response.Payload.Messages)
	}

	return response.Payload.BackupListCollection, nil
}
