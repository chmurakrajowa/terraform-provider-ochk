package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/firewall_rules_e_w"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/go-openapi/strfmt"
	"net/http"
)

type FirewallEWRulesProxy struct {
	httpClient *http.Client
	service    firewall_rules_e_w.ClientService
}

func (p *FirewallEWRulesProxy) Create(ctx context.Context, securityPolicyID string, rule *models.DFWRule) (*models.DFWRule, error) {
	if err := rule.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating firewall EW rule struct: %w", err)
	}

	params := &firewall_rules_e_w.DFWRuleCreateUsingPUTParams{
		SecurityPolicyID: securityPolicyID,
		DfwRule:          rule,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	_, put, err := p.service.DFWRuleCreateUsingPUT(params)
	if err != nil {
		return nil, fmt.Errorf("error while creating firewall EW rule: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating firewall EW rule failed: %s", put.Payload.Messages)
	}

	return put.Payload.DfwRule, nil
}

func (p *FirewallEWRulesProxy) Read(ctx context.Context, securityPolicyID string, ruleID string) (*models.DFWRule, error) {
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

func (p *FirewallEWRulesProxy) Update(ctx context.Context, securityPolicyID string, rule *models.DFWRule) (*models.DFWRule, error) {
	if err := rule.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating firewall EW rule struct: %w", err)
	}

	params := &firewall_rules_e_w.DFWRuleUpdateUsingPUTParams{
		SecurityPolicyID: securityPolicyID,
		RuleID:           rule.RuleID,
		DfwRule:          rule,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	put, err := p.service.DFWRuleUpdateUsingPUT(params)
	if err != nil {
		return nil, fmt.Errorf("error while updating firewall EW rule: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating updating EW rule failed: %s", put.Payload.Messages)
	}

	return put.Payload.DfwRule, nil
}

func (p *FirewallEWRulesProxy) ListByDisplayName(ctx context.Context, securityPolicyID string, displayName string) ([]*models.DFWRule, error) {
	params := &firewall_rules_e_w.DFWRuleListUsingGETParams{
		SecurityPolicyID: securityPolicyID,
		DisplayName:      &displayName,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	response, err := p.service.DFWRuleListUsingGET(params)
	if err != nil {
		return nil, fmt.Errorf("error while listing firewall EW rule: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing firewall EW rule failed: %s", response.Payload.Messages)
	}

	return response.Payload.RuleInstances, nil
}

func (p *FirewallEWRulesProxy) List(ctx context.Context, securityPolicyID string) ([]*models.DFWRule, error) {
	params := &firewall_rules_e_w.DFWRuleListUsingGETParams{
		SecurityPolicyID: securityPolicyID,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	response, err := p.service.DFWRuleListUsingGET(params)
	if err != nil {
		return nil, fmt.Errorf("error while listing firewall EW rule: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing firewall EW rule failed: %s", response.Payload.Messages)
	}

	return response.Payload.RuleInstances, nil
}

func (p *FirewallEWRulesProxy) Exists(ctx context.Context, securityPolicyID string, ruleID string) (bool, error) {
	if _, err := p.Read(ctx, securityPolicyID, ruleID); err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}

		return false, fmt.Errorf("error while reading firewall EW rule: %w", err)
	}

	return true, nil
}

func (p *FirewallEWRulesProxy) Delete(ctx context.Context, securityPolicyID string, ruleID string) error {
	params := &firewall_rules_e_w.DFWRuleDeleteUsingDELETEParams{
		SecurityPolicyID: securityPolicyID,
		RuleID:           ruleID,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	response, err := p.service.DFWRuleDeleteUsingDELETE(params)
	if err != nil {
		var badRequest *firewall_rules_e_w.DFWRuleDeleteUsingDELETEBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return &NotFoundError{Err: err}
		}

		return fmt.Errorf("error while deleting firewall EW rule: %w", err)
	}

	if !response.Payload.Success {
		return fmt.Errorf("deleting firewall EW rule failed: %s", response.Payload.Messages)
	}

	return nil
}
