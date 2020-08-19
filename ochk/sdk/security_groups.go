package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/security_groups"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"net/http"
)

type SecurityGroupsProxy struct {
	httpClient *http.Client
	service    security_groups.ClientService
}

func (p *SecurityGroupsProxy) Create(ctx context.Context, securityGroup *models.SecurityGroup) (*models.SecurityGroup, error) {
	if err := securityGroup.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating security group struct: %w", err)
	}

	params := &security_groups.SecurityGroupCreateUsingPUTParams{
		SecurityGroup: securityGroup,
		Context:       ctx,
		HTTPClient:    p.httpClient,
	}

	_, put, err := p.service.SecurityGroupCreateUsingPUT(params)
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

	params := &security_groups.SecurityGroupUpdateUsingPUTParams{
		GroupID:       securityGroup.ID,
		SecurityGroup: securityGroup,
		Context:       ctx,
		HTTPClient:    p.httpClient,
	}

	put, err := p.service.SecurityGroupUpdateUsingPUT(params)
	if err != nil {
		return nil, fmt.Errorf("error while modifying security group: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying security group failed: %s", put.Payload.Messages)
	}

	return put.Payload.SecurityGroup, nil
}

func (p *SecurityGroupsProxy) Read(ctx context.Context, securityGroupID string) (*models.SecurityGroup, error) {
	params := &security_groups.SecurityGroupGetUsingGETParams{
		GroupID:    securityGroupID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.SecurityGroupGetUsingGET(params)
	if err != nil {
		var notFound *security_groups.SecurityGroupGetUsingGETNotFound
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
	params := &security_groups.SecurityGroupListUsingGETParams{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	response, err := p.service.SecurityGroupListUsingGET(params)
	if err != nil {
		return nil, fmt.Errorf("error while listing security groups: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing security groups failed: %s", response.Payload.Messages)
	}

	return response.Payload.SecurityGroupCollection, nil
}

func (p *SecurityGroupsProxy) Exists(ctx context.Context, securityGroupID string) (bool, error) {
	_, err := p.Read(ctx, securityGroupID)
	if err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}

		return false, fmt.Errorf("error while reading security group: %w", err)
	}

	return true, nil
}

func (p *SecurityGroupsProxy) Delete(ctx context.Context, securityGroupID string) error {
	params := &security_groups.SecurityGroupDeleteUsingDELETEParams{
		GroupID:    securityGroupID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.SecurityGroupDeleteUsingDELETE(params)
	if err != nil {
		var badRequest *security_groups.SecurityGroupDeleteUsingDELETEBadRequest
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
