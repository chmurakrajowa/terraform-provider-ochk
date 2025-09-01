package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/go-openapi/strfmt"

	"net/http"
)

type DeploymentsProxy struct {
	httpClient *http.Client
	service    openapi.DeploymentsAPIService
}

func (p *DeploymentsProxy) Read(ctx context.Context, deploymentID strfmt.UUID) (*openapi.DeploymentInstance, error) {
	action := p.service.DeploymentsDeploymentIdGet(ctx, string(deploymentID))
	response, _, err := action.Execute()

	if err != nil {
		return nil, fmt.Errorf("error while reading deployment: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("retrieving deployment failed: %s", response.Messages)
	}

	return response.DeploymentInstance, nil
}

func (p *DeploymentsProxy) ListByDisplayName(ctx context.Context, displayName string) ([]openapi.DeploymentInstance, error) {

	action := p.service.DeploymentsGet(ctx).DisplayName(displayName)
	response, _, err := action.Execute()

	if err != nil {
		return nil, fmt.Errorf("error while listing deployments by display name %s: %w", displayName, err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing deployments by display name %s failed: %s", displayName, response.Messages)
	}

	return response.DeploymentInstanceCollection, nil
}

func (p *DeploymentsProxy) List(ctx context.Context) ([]openapi.DeploymentInstance, error) {
	action := p.service.DeploymentsGet(ctx)
	response, _, err := action.Execute()

	if err != nil {
		return nil, fmt.Errorf("error while listing deployments: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing deployments failed: %s", response.Messages)
	}

	return response.DeploymentInstanceCollection, nil
}
