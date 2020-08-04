package sdk

import (
	"context"
	"fmt"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/firewall_rules_e_w"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
	"net/http"
)

type FirewallEWRules struct {
	httpClient *http.Client
	service    firewall_rules_e_w.ClientService
}

func (p *FirewallEWRules) Read(ctx context.Context, securityPolicyID string, ruleID string) (*models.DFWRule, error) {
	params := &firewall_rules_e_w.DFWRuleGetUsingGETParams{
		RuleID:           ruleID,
		SecurityPolicyID: securityPolicyID,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	response, err := p.service.DFWRuleGetUsingGET(params)
	if err != nil {
		return nil, fmt.Errorf("error while reading firwall EW rule: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving firewall EW rule failed: %s", response.Payload.Messages)
	}

	return response.Payload.RuleInstance, nil
}
