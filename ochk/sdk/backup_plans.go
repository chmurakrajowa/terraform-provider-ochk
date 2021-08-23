package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/backups"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"net/http"
)

type BackupPlansProxy struct {
	httpClient *http.Client
	service    backups.ClientService
}

func (p *BackupPlansProxy) Read(ctx context.Context, packupPlanID string) (*models.BackupPlan, error) {
	params := &backups.BackupPlanGetUsingGETParams{
		BackupPlanID: packupPlanID,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}

	response, err := p.service.BackupPlanGetUsingGET(params)

	if err != nil {
		var notFound *backups.BackupPlanGetUsingGETNotFound

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
	params := &backups.BackupPlanListUsingGETParams{
		BackupPlanName: &backupPlanName,
		Context:        ctx,
		HTTPClient:     p.httpClient,
	}

	response, err := p.service.BackupPlanListUsingGET(params)

	if err != nil {
		return nil, fmt.Errorf("error while listing backup plans: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("Listing backup plans failed: %s", response.Payload.Messages)
	}

	return response.Payload.BackupPlanCollection, nil
}
