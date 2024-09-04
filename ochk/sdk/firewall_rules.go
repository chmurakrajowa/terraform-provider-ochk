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

	if err != nil {
		return nil, fmt.Errorf("error while listing firewall openstack rule: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing firewall openstack rule failed: %s", response.Payload.Messages)
	}

	return response.Payload.FirewallRuleCollection, nil
}

// /projects/3cda830c-b37f-46dc-be54-f649d31bec66/osc/security-groups/ff765340-2c9e-4dc8-b565-bd1d84793f5d/firewall/06f8f3bc-93d7-4d84-98bd-270e2b9b847f
// /projects/ff765340-2c9e-4dc8-b565-bd1d84793f5d/osc/security-groups/3cda830c-b37f-46dc-be54-f649d31bec66/firewall/06f8f3bc-93d7-4d84-98bd-270e2b9b847f
