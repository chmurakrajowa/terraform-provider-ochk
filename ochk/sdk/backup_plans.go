package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type BackupPlansProxy struct {
	httpClient *http.Client
	service    *openapi.BackupsAPIService
}

func (p *BackupPlansProxy) Read(ctx context.Context, backupPlanID strfmt.UUID) (*openapi.BackupPlan, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.BackupsPlansBackupPlanIdGet(ctx, string(backupPlanID))
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading backup plan: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("retrieving backup plan failed: %s", response.Messages)
	}

	return response.BackupPlan, nil
}

func (p *BackupPlansProxy) ListBackupPlanByName(ctx context.Context, backupPlanName string) ([]openapi.BackupPlan, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.BackupsPlansGet(ctx).BackupPlanName(backupPlanName)
	response, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing backup plans: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("Listing backup plans failed: %s", response.Messages)
	}

	return response.BackupPlanCollection, nil
}

func (p *BackupPlansProxy) ListBackupPlans(ctx context.Context) ([]openapi.BackupPlan, error) {

	action := p.service.BackupsPlansGet(ctx)
	response, _, err := action.Execute()

	if err != nil {
		return nil, fmt.Errorf("error while listing backup plans: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("Listing backup plans failed: %s", response.Messages)
	}

	return response.BackupPlanCollection, nil
}
