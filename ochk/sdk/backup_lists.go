package sdk

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
)

type BackupListsProxy struct {
	httpClient *http.Client
	service    *openapi.BackupsAPIService
}

func (p *BackupListsProxy) Read(ctx context.Context, backupPlanID strfmt.UUID, backupListID strfmt.UUID) (*openapi.BackupList, error) {

	action := p.service.BackupsPlansBackupPlanIdListsBackupListIdGet(ctx, string(backupPlanID), string(backupListID))

	response, _, err := action.Execute()
	if err != nil {
		return nil, fmt.Errorf("error while reading backup list: %w", err)
	}

	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("retrieving backup list failed: %s", message)
	}

	return response.BackupList, nil
}

func (p *BackupListsProxy) ListBackupListByName(ctx context.Context, backupPlanID strfmt.UUID, backupListName string) ([]openapi.BackupList, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.BackupsPlansBackupPlanIdListsGet(ctx, string(backupPlanID)).BackupListName(backupListName)
	response, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing backup list: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("listing backup list failed: %s", message)
	}

	return response.BackupListCollection, nil
}

func (p *BackupListsProxy) ListBackupList(ctx context.Context, backupPlanID strfmt.UUID) ([]openapi.BackupList, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.BackupsPlansBackupPlanIdListsGet(ctx, string(backupPlanID))
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing backup list: %w", err)
	}

	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}

		return nil, fmt.Errorf("listing backup list failed: %s", message)
	}

	return response.BackupListCollection, nil
}
