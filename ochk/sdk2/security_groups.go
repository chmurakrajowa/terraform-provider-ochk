package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type SecurityGroupsProxy struct {
	httpClient *http.Client
	service    *openapi.SecurityGroupAPIService
}

func (p *SecurityGroupsProxy) Create(ctx context.Context, securityGroup openapi.SecurityGroup) (*openapi.SecurityGroup, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkSecurityGroupsPut(ctx).SecurityGroup(securityGroup)
	put, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating security group: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("creating security group failed: %s", put.Messages)
	}

	return put.SecurityGroup, nil
}

func (p *SecurityGroupsProxy) Update(ctx context.Context, securityGroup openapi.SecurityGroup) (*openapi.SecurityGroup, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkSecurityGroupsGroupIdPut(ctx, securityGroup.GetId()).SecurityGroup(securityGroup)
	put, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifying security group: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("modifying security group failed: %s", put.Messages)
	}

	return put.SecurityGroup, nil
}

func (p *SecurityGroupsProxy) Read(ctx context.Context, securityGroupID strfmt.UUID) (*openapi.SecurityGroup, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkSecurityGroupsGroupIdGet(ctx, string(securityGroupID))
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading security group: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("retrieving security group failed: %s", response.Messages)
	}

	return response.SecurityGroup, nil
}

func (p *SecurityGroupsProxy) ListByDisplayName(ctx context.Context, displayName string) ([]openapi.SecurityGroup, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkSecurityGroupsGet(ctx).DisplayName(displayName)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing security groups: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing security groups failed: %s", response.Messages)
	}

	return response.SecurityGroupCollection, nil
}

func (p *SecurityGroupsProxy) List(ctx context.Context) ([]openapi.SecurityGroup, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkSecurityGroupsGet(ctx)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing security groups: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing security groups failed: %s", response.Messages)
	}

	return response.SecurityGroupCollection, nil
}

func (p *SecurityGroupsProxy) Exists(ctx context.Context, securityGroupID strfmt.UUID) (bool, error) {
	if _, err := p.Read(ctx, securityGroupID); err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}

		return false, fmt.Errorf("error while reading security group: %w", err)
	}

	return true, nil
}

func (p *SecurityGroupsProxy) Delete(ctx context.Context, securityGroupID strfmt.UUID) error {

	action := p.service.NetworkSecurityGroupsGroupIdDelete(ctx, string(securityGroupID))
	response, _, err := action.Execute()
	if err != nil {
		return fmt.Errorf("error while deleting security group: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return fmt.Errorf("deleting security group failed: %s", response.Messages)
	}

	return nil
}
