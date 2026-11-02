package sdk

import (
	"context"
	"fmt"
	"io"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"

	"net/http"
	"sync"
)

type ProjectsProxy struct {
	httpClient *http.Client
	service    *openapi.ProjectsAPIService
}

func (p *ProjectsProxy) Create(ctx context.Context, project openapi.ProjectInstance) (*openapi.ProjectInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.ProjectsPut(ctx).ProjectInstance(project)
	put, resp, err := action.Execute()

	if resp != nil {
		bodyBytes, _ := io.ReadAll(resp.Body)
		if resp.StatusCode == 400 {
			return nil, fmt.Errorf("error while creating project: %s", string(bodyBytes))
		}
	}
	mutex.Unlock()
	if err != nil {
		return nil, fmt.Errorf("error while creating project: %w", err)
	}

	isSuccess := put.Success
	if !*isSuccess {
		message := "no error message"
		if value := put.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("creating project failed: %s", message)
	}

	return put.ProjectInstance, nil
}

func (p *ProjectsProxy) Update(ctx context.Context, project openapi.ProjectInstance) (*openapi.ProjectInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.ProjectsProjectIdPut(ctx, project.GetProjectId()).ProjectInstance(project)
	put, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifying project: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		message := "no error message"
		if value := put.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("modifying project failed: %s", message)
	}

	return put.ProjectInstance, nil
}

func (p *ProjectsProxy) Read(ctx context.Context, projectID strfmt.UUID) (*openapi.ProjectInstance, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.ProjectsProjectIdGet(ctx, string(projectID))
	response, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading project: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("retrieving project failed: %s", message)
	}

	return response.ProjectInstance, nil
}

func (p *ProjectsProxy) ListByName(ctx context.Context, name string) ([]openapi.ProjectInstance, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.ProjectsGet(ctx).Name(name)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing projects: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("listing projects failed: %s", message)
	}

	return response.ProjectInstanceCollection, nil
}
func (p *ProjectsProxy) List(ctx context.Context) ([]openapi.ProjectInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.ProjectsGet(ctx)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing projects: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("listing projects failed: %s", message)
	}

	return response.ProjectInstanceCollection, nil
}

func (p *ProjectsProxy) Exists(ctx context.Context, projectID strfmt.UUID) (bool, error) {
	if _, err := p.Read(ctx, projectID); err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}

		return false, fmt.Errorf("error while reading project: %w", err)
	}

	return true, nil
}

func (p *ProjectsProxy) Delete(ctx context.Context, projectID strfmt.UUID) error {
	action := p.service.ProjectsProjectIdDelete(ctx, string(projectID))
	response, _, err := action.Execute()

	if err != nil {
		return fmt.Errorf("error while deleting project: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return fmt.Errorf("deleting project failed: %s", message)
	}

	return nil
}
