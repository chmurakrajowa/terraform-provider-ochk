package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/deployments"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"net/http"
)

type DeploymentsProxy struct {
	httpClient *http.Client
	service    deployments.ClientService
}

func (p *DeploymentsProxy) Read(ctx context.Context, deploymentID strfmt.UUID) (*models.DeploymentInstance, error) {
	params := &deployments.GetDeploymentsDeploymentIDParams{
		DeploymentID: deploymentID,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}

	response, err := p.service.GetDeploymentsDeploymentID(params)

	if err != nil {
		var notFound *deployments.GetDeploymentsDeploymentIDNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading deployment: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving deployment failed: %s", response.Payload.Messages)
	}

	return response.Payload.DeploymentInstance, nil
}

func (p *DeploymentsProxy) ListByDisplayName(ctx context.Context, displayName string) ([]*models.DeploymentInstance, error) {
	params := &deployments.GetDeploymentsParams{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	response, err := p.service.GetDeployments(params)
	if err != nil {
		return nil, fmt.Errorf("error while listing deployments by display name %s: %w", displayName, err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing deployments by display name %s failed: %s", displayName, response.Payload.Messages)
	}

	return response.Payload.DeploymentInstanceCollection, nil
}

func (p *DeploymentsProxy) List(ctx context.Context) ([]*models.DeploymentInstance, error) {
	params := &deployments.GetDeploymentsParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.GetDeployments(params)
	if err != nil {
		return nil, fmt.Errorf("error while listing deployments: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing deployments failed: %s", response.Payload.Messages)
	}

	return response.Payload.DeploymentInstanceCollection, nil
}
