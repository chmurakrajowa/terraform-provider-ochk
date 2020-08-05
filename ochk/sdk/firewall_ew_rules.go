package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/firewall_rules_e_w"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
	"net/http"
)

type FirewallEWRules struct {
	httpClient *http.Client
	service    firewall_rules_e_w.ClientService
}

//TODO przy create wysyłane są też daty, które powinny być opcjonalne i wyliczone po stronie backendu
func (p *FirewallEWRules) Create(ctx context.Context, securityPolicyID string, rule *models.DFWRule) (*models.DFWRule, error) {
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

func (p *FirewallEWRules) ListByDisplayName(ctx context.Context, securityPolicyID string, displayName string) ([]*models.DFWRule, error) {
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
		return nil, fmt.Errorf("listing firewall ER rule failed: %s", response.Payload.Messages)
	}

	return response.Payload.RuleInstances, nil
}

func (p *FirewallEWRules) Exists(ctx context.Context, securityPolicyID string, ruleID string) (bool, error) {
	_, err := p.Read(ctx, securityPolicyID, ruleID)
	if err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}

		return false, fmt.Errorf("error while reading firewall EW rule: %w", err)
	}

	return true, nil
}

func (p *FirewallEWRules) Delete(ctx context.Context, securityPolicyID string, ruleID string) error {
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
