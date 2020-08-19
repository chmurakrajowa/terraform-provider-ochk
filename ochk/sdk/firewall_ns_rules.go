package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/firewall_rules_s_n"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"net/http"
)

type FirewallSNRulesProxy struct {
	httpClient *http.Client
	service    firewall_rules_s_n.ClientService
}

func (p *FirewallSNRulesProxy) Create(ctx context.Context, gatewayPolicyID string, rule *models.GFWRule) (*models.GFWRule, error) {
	if err := rule.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating firewall SN rule struct: %w", err)
	}

	params := &firewall_rules_s_n.GFWRuleCreateUsingPUTParams{
		GatewayPolicyID: gatewayPolicyID,
		GfwRule:         rule,
		Context:         ctx,
		HTTPClient:      p.httpClient,
	}

	_, put, err := p.service.GFWRuleCreateUsingPUT(params)
	if err != nil {
		return nil, fmt.Errorf("error while creating firewall SN rule: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating firewall SN rule failed: %s", put.Payload.Messages)
	}

	return put.Payload.GfwRule, nil
}

func (p *FirewallSNRulesProxy) Update(ctx context.Context, gatewayPolicyID string, rule *models.GFWRule) (*models.GFWRule, error) {
	if err := rule.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating firewall SN rule struct: %w", err)
	}

	params := &firewall_rules_s_n.GFWRuleUpdateUsingPUTParams{
		GatewayPolicyID: gatewayPolicyID,
		RuleID:          rule.RuleID,
		GfwRule:         rule,
		Context:         ctx,
		HTTPClient:      p.httpClient,
	}

	put, err := p.service.GFWRuleUpdateUsingPUT(params)
	if err != nil {
		return nil, fmt.Errorf("error while modifing firewall SN rule: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifing firewall SN rule failed: %s", put.Payload.Messages)
	}

	return put.Payload.GfwRule, nil
}

func (p *FirewallSNRulesProxy) Read(ctx context.Context, gatewayPolicyID string, ruleID string) (*models.GFWRule, error) {
	params := &firewall_rules_s_n.GFWRuleGetUsingGETParams{
		RuleID:          ruleID,
		GatewayPolicyID: gatewayPolicyID,
		Context:         ctx,
		HTTPClient:      p.httpClient,
	}

	response, err := p.service.GFWRuleGetUsingGET(params)
	if err != nil {
		return nil, fmt.Errorf("error while reading firwall SN rule: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving firewall SN rule failed: %s", response.Payload.Messages)
	}

	return response.Payload.RuleInstance, nil
}

func (p *FirewallSNRulesProxy) ListByDisplayName(ctx context.Context, gatewayPolicyID string, displayName string) ([]*models.GFWRule, error) {
	params := &firewall_rules_s_n.GFWRuleListUsingGETParams{
		GatewayPolicyID: gatewayPolicyID,
		DisplayName:     &displayName,
		Context:         ctx,
		HTTPClient:      p.httpClient,
	}

	response, err := p.service.GFWRuleListUsingGET(params)
	if err != nil {
		return nil, fmt.Errorf("error while listing firewall SN rule: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing firewall SN rule failed: %s", response.Payload.Messages)
	}

	return response.Payload.RuleInstances, nil
}

func (p *FirewallSNRulesProxy) Exists(ctx context.Context, gatewayPolicyID string, ruleID string) (bool, error) {
	_, err := p.Read(ctx, gatewayPolicyID, ruleID)
	if err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}

		return false, fmt.Errorf("error while reading firewall SN rule: %w", err)
	}

	return true, nil
}

func (p *FirewallSNRulesProxy) Delete(ctx context.Context, gatewayPolicyID string, ruleID string) error {
	params := &firewall_rules_s_n.GFWRuleDeleteUsingDELETEParams{
		GatewayPolicyID: gatewayPolicyID,
		RuleID:          ruleID,
		Context:         ctx,
		HTTPClient:      p.httpClient,
	}

	response, err := p.service.GFWRuleDeleteUsingDELETE(params)
	if err != nil {
		var badRequest *firewall_rules_s_n.GFWRuleDeleteUsingDELETEBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return &NotFoundError{Err: err}
		}

		return fmt.Errorf("error while deleting firewall SN rule: %w", err)
	}

	if !response.Payload.Success {
		return fmt.Errorf("deleting firewall SN rule failed: %s", response.Payload.Messages)
	}

	return nil
}
