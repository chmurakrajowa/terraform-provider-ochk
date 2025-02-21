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

type BackupPlansProxy struct {
	httpClient *http.Client
	service    backups.ClientService
}

func (p *BackupPlansProxy) Read(ctx context.Context, backupPlanID strfmt.UUID) (*models.BackupPlan, error) {
	params := &backups.GetBackupsPlansBackupPlanIDParams{
		BackupPlanID: backupPlanID,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetBackupsPlansBackupPlanID(params)
	mutex.Unlock()

	if err != nil {
		var notFound *backups.GetBackupsPlansBackupPlanIDNotFound

		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading backup plan: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving backup plan failed: %s", response.Payload.Messages)
	}

	return response.Payload.BackupPlan, nil
}

func (p *BackupPlansProxy) ListBackupPlanByName(ctx context.Context, backupPlanName string) ([]*models.BackupPlan, error) {
	params := &backups.GetBackupsPlansParams{
		BackupPlanName: &backupPlanName,
		Context:        ctx,
		HTTPClient:     p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetBackupsPlans(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing backup plans: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("Listing backup plans failed: %s", response.Payload.Messages)
	}

	return response.Payload.BackupPlanCollection, nil
}

func (p *BackupPlansProxy) ListBackupPlans(ctx context.Context) ([]*models.BackupPlan, error) {
	params := &backups.GetBackupsPlansParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.GetBackupsPlans(params)

	if err != nil {
		return nil, fmt.Errorf("error while listing backup plans: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("Listing backup plans failed: %s", response.Payload.Messages)
	}

	return response.Payload.BackupPlanCollection, nil
}
