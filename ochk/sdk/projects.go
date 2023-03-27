package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/projects"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type ProjectsProxy struct {
	httpClient *http.Client
	service    projects.ClientService
}

func (p *ProjectsProxy) Create(ctx context.Context, project *models.ProjectInstance) (*models.ProjectInstance, error) {
	if err := project.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating project struct: %w", err)
	}

	params := &projects.ProjectCreateUsingPUTParams{
		ProjectInstance: project,
		Context:         ctx,
		HTTPClient:      p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	_, put, err := p.service.ProjectCreateUsingPUT(params)
	mutex.Unlock()
	if err != nil {
		return nil, fmt.Errorf("error while creating project: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating project failed: %s", put.Payload.Messages)
	}

	return put.Payload.ProjectInstance, nil
}

func (p *ProjectsProxy) Update(ctx context.Context, project *models.ProjectInstance) (*models.ProjectInstance, error) {
	if err := project.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating project struct: %w", err)
	}

	params := &projects.ProjectUpdateUsingPUTParams{
		ProjectID:       project.ProjectID,
		ProjectInstance: project,
		Context:         ctx,
		HTTPClient:      p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, _, err := p.service.ProjectUpdateUsingPUT(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifying project: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying project failed: %s", put.Payload.Messages)
	}

	return put.Payload.ProjectInstance, nil
}

func (p *ProjectsProxy) Read(ctx context.Context, projectID string) (*models.ProjectInstance, error) {
	params := &projects.ProjectGetUsingGETParams{
		ProjectID:  projectID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.ProjectGetUsingGET(params)
	mutex.Unlock()

	if err != nil {
		var notFound *projects.ProjectGetUsingGETNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading project: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving project failed: %s", response.Payload.Messages)
	}

	return response.Payload.ProjectInstance, nil
}

func (p *ProjectsProxy) ListByName(ctx context.Context, name string) ([]*models.ProjectInstance, error) {
	params := &projects.ProjectListUsingGETParams{
		Name:       &name,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.ProjectListUsingGET(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing projects: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing projects failed: %s", response.Payload.Messages)
	}

	return response.Payload.ProjectInstanceCollection, nil
}
func (p *ProjectsProxy) List(ctx context.Context) ([]*models.ProjectInstance, error) {
	params := &projects.ProjectListUsingGETParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.ProjectListUsingGET(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing projects: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing projects failed: %s", response.Payload.Messages)
	}

	return response.Payload.ProjectInstanceCollection, nil
}

func (p *ProjectsProxy) Exists(ctx context.Context, projectID string) (bool, error) {
	if _, err := p.Read(ctx, projectID); err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}

		return false, fmt.Errorf("error while reading project: %w", err)
	}

	return true, nil
}

func (p *ProjectsProxy) Delete(ctx context.Context, projectID string) error {
	params := &projects.ProjectDeleteUsingDELETEParams{
		ProjectID:  projectID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, _, err := p.service.ProjectDeleteUsingDELETE(params)

	if err != nil {
		var badRequest *projects.ProjectDeleteUsingDELETEBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return &NotFoundError{Err: err}
		}

		return fmt.Errorf("error while deleting project: %w", err)
	}

	if !response.Payload.Success {
		return fmt.Errorf("deleting project failed: %s", response.Payload.Messages)
	}

	return nil
}
