package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/firewall_rule"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type FirewallRulesProxy struct {
	httpClient *http.Client
	service    firewall_rule.ClientService
}

func (p *FirewallRulesProxy) Read(ctx context.Context, projectId strfmt.UUID, securityGroupId strfmt.UUID, ruleId strfmt.UUID) (*models.FirewallRule, error) {
	params := &firewall_rule.GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDParams{
		RuleID:          ruleId,
		ProjectID:       projectId,
		SecurityGroupID: securityGroupId,
		Context:         ctx,
		HTTPClient:      p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleID(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading firwall openstack rule: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving firwall openstack rule: failed: %s", response.Payload.Messages)
	}
	if response.Payload.FirewallRule != nil {
		return response.Payload.FirewallRule, nil
	} else {
		return nil, nil
	}
}

func (p *FirewallRulesProxy) List(ctx context.Context, projectId strfmt.UUID, securityGroupId strfmt.UUID) ([]*models.FirewallRule, error) {

	params := &firewall_rule.GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallParams{
		ProjectID:       projectId,
		SecurityGroupID: securityGroupId,
		Context:         ctx,
		HTTPClient:      p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewall(params)
	mutex.Unlock()

	if false {
		return nil, fmt.Errorf("openstack rule response.Payload.Success: %w", response)
	}

	if err != nil {
		return nil, fmt.Errorf("error while listing firewall openstack rule: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing firewall openstack rule failed: %s", response.Payload.Messages)
	}

	return response.Payload.FirewallRuleCollection, nil
}

func (p *FirewallRulesProxy) ListByName(ctx context.Context, projectId strfmt.UUID, securityGroupId strfmt.UUID, name string) ([]*models.FirewallRule, error) {
	params := &firewall_rule.GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallParams{
		ProjectID:       projectId,
		SecurityGroupID: securityGroupId,
		Name:            &name,
		Context:         ctx,
		HTTPClient:      p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewall(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing firewall rules: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing firewall rules failed: %s", response.Payload.Messages)
	}

	return response.Payload.FirewallRuleCollection, nil
}

func (p *FirewallRulesProxy) Create(ctx context.Context, projectId strfmt.UUID, securityGroupId strfmt.UUID, rule *models.FirewallRule) (*models.FirewallRule, error) {
	if err := rule.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating firewall rule struct: %w", err)
	}

	params := &firewall_rule.PutProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallParams{
		ProjectID:       projectId,
		SecurityGroupID: securityGroupId,
		Body:            rule,
		Context:         ctx,
		HTTPClient:      p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewall(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating firewall rule: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating firewall rule failed: %s", put.Payload.Messages)
	}

	return put.Payload.FirewallRule, nil
}

func (p *FirewallRulesProxy) Update(ctx context.Context, projectId strfmt.UUID, securityGroupId strfmt.UUID, rule *models.FirewallRule) (*models.FirewallRule, error) {
	if err := rule.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating firewall EW rule struct: %w", err)
	}

	params := &firewall_rule.PutProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDParams{
		ProjectID:       projectId,
		SecurityGroupID: securityGroupId,
		RuleID:          rule.RuleID,
		Body:            rule,
		Context:         ctx,
		HTTPClient:      p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleID(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while updating firewall rule: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating updating rule failed: %s", put.Payload.Messages)
	}

	return put.Payload.FirewallRule, nil
}

func (p *FirewallRulesProxy) Delete(ctx context.Context, projectId strfmt.UUID, securityGroupId strfmt.UUID, ruleID strfmt.UUID) error {
	params := &firewall_rule.DeleteProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDParams{
		ProjectID:       projectId,
		SecurityGroupID: securityGroupId,
		RuleID:          ruleID,
		Context:         ctx,
		HTTPClient:      p.httpClient,
	}

	response, err := p.service.DeleteProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleID(params)

	if err != nil {
		var badRequest *firewall_rule.DeleteProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return &NotFoundError{Err: err}
		}

		return fmt.Errorf("error while deleting firewall rule: %w", err)
	}

	if !response.Payload.Success {
		return fmt.Errorf("deleting firewall rule failed: %s", response.Payload.Messages)
	}

	return nil
}
