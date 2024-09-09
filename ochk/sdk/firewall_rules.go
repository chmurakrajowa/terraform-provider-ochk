package sdk

import (
	"context"
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
