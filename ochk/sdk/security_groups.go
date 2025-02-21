package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/security_group"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type SecurityGroupsProxy struct {
	httpClient *http.Client
	service    security_group.ClientService
}

func (p *SecurityGroupsProxy) Create(ctx context.Context, securityGroup *models.SecurityGroup) (*models.SecurityGroup, error) {
	if err := securityGroup.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating security group struct: %w", err)
	}

	params := &security_group.PutNetworkSecurityGroupsParams{
		Body:       securityGroup,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutNetworkSecurityGroups(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating security group: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating security group failed: %s", put.Payload.Messages)
	}

	return put.Payload.SecurityGroup, nil
}

func (p *SecurityGroupsProxy) Update(ctx context.Context, securityGroup *models.SecurityGroup) (*models.SecurityGroup, error) {
	if err := securityGroup.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating security group struct: %w", err)
	}

	params := &security_group.PutNetworkSecurityGroupsGroupIDParams{
		GroupID:    securityGroup.ID,
		Body:       securityGroup,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutNetworkSecurityGroupsGroupID(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifying security group: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying security group failed: %s", put.Payload.Messages)
	}

	return put.Payload.SecurityGroup, nil
}

func (p *SecurityGroupsProxy) Read(ctx context.Context, securityGroupID strfmt.UUID) (*models.SecurityGroup, error) {
	params := &security_group.GetNetworkSecurityGroupsGroupIDParams{
		GroupID:    securityGroupID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkSecurityGroupsGroupID(params)
	mutex.Unlock()

	if err != nil {
		var notFound *security_group.GetNetworkSecurityGroupsGroupIDNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading security group: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving security group failed: %s", response.Payload.Messages)
	}

	return response.Payload.SecurityGroup, nil
}

func (p *SecurityGroupsProxy) ListByDisplayName(ctx context.Context, displayName string) ([]*models.SecurityGroup, error) {
	params := &security_group.GetNetworkSecurityGroupsParams{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkSecurityGroups(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing security groups: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing security groups failed: %s", response.Payload.Messages)
	}

	return response.Payload.SecurityGroupCollection, nil
}

func (p *SecurityGroupsProxy) List(ctx context.Context) ([]*models.SecurityGroup, error) {
	params := &security_group.GetNetworkSecurityGroupsParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkSecurityGroups(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing security groups: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing security groups failed: %s", response.Payload.Messages)
	}

	return response.Payload.SecurityGroupCollection, nil
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
	params := &security_group.DeleteNetworkSecurityGroupsGroupIDParams{
		GroupID:    securityGroupID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.DeleteNetworkSecurityGroupsGroupID(params)

	if err != nil {
		var badRequest *security_group.DeleteNetworkSecurityGroupsGroupIDBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return &NotFoundError{Err: err}
		}

		return fmt.Errorf("error while deleting security group: %w", err)
	}

	if !response.Payload.Success {
		return fmt.Errorf("deleting security group failed: %s", response.Payload.Messages)
	}

	return nil
}
