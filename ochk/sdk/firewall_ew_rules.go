package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/firewall_rules_e_w"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type FirewallEWRulesProxy struct {
	httpClient *http.Client
	service    firewall_rules_e_w.ClientService
}

func (p *FirewallEWRulesProxy) Create(ctx context.Context, routerID string, rule *models.DFWRule) (*models.DFWRule, error) {
	if err := rule.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating firewall EW rule struct: %w", err)
	}

	params := &firewall_rules_e_w.DfwRuleCreateUsingPUTParams{
		RouterID:   routerID,
		DfwRule:    rule,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	_, put, err := p.service.DfwRuleCreateUsingPUT(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating firewall EW rule: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating firewall EW rule failed: %s", put.Payload.Messages)
	}

	return put.Payload.DfwRule, nil
}

func (p *FirewallEWRulesProxy) Read(ctx context.Context, routerID string, ruleID string) (*models.DFWRule, error) {
	params := &firewall_rules_e_w.DfwRuleGetUsingGETParams{
		RuleID:     ruleID,
		RouterID:   routerID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.DfwRuleGetUsingGET(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading firwall EW rule: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving firewall EW rule failed: %s", response.Payload.Messages)
	}

	return response.Payload.RuleInstance, nil
}

func (p *FirewallEWRulesProxy) Update(ctx context.Context, routerID string, rule *models.DFWRule) (*models.DFWRule, error) {
	if err := rule.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating firewall EW rule struct: %w", err)
	}

	params := &firewall_rules_e_w.DfwRuleUpdateUsingPUTParams{
		RouterID:   routerID,
		RuleID:     rule.RuleID,
		DfwRule:    rule,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, _, err := p.service.DfwRuleUpdateUsingPUT(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while updating firewall EW rule: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating updating EW rule failed: %s", put.Payload.Messages)
	}

	return put.Payload.DfwRule, nil
}

func (p *FirewallEWRulesProxy) ListByDisplayName(ctx context.Context, routerID string, displayName string) ([]*models.DFWRule, error) {
	params := &firewall_rules_e_w.DfwRuleListUsingGETParams{
		RouterID:    routerID,
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.DfwRuleListUsingGET(params)
	mutex.Unlock()
	if err != nil {
		return nil, fmt.Errorf("error while listing firewall EW rule: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing firewall EW rule failed: %s", response.Payload.Messages)
	}

	return response.Payload.RuleInstances, nil
}

func (p *FirewallEWRulesProxy) List(ctx context.Context, routerID string) ([]*models.DFWRule, error) {
	params := &firewall_rules_e_w.DfwRuleListUsingGETParams{
		RouterID:   routerID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.DfwRuleListUsingGET(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing firewall EW rule: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing firewall EW rule failed: %s", response.Payload.Messages)
	}

	return response.Payload.RuleInstances, nil
}

func (p *FirewallEWRulesProxy) Exists(ctx context.Context, routerID string, ruleID string) (bool, error) {
	if _, err := p.Read(ctx, routerID, ruleID); err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}

		return false, fmt.Errorf("error while reading firewall EW rule: %w", err)
	}

	return true, nil
}

func (p *FirewallEWRulesProxy) Delete(ctx context.Context, routerID string, ruleID string) error {
	params := &firewall_rules_e_w.DfwRuleDeleteUsingDELETEParams{
		RouterID:   routerID,
		RuleID:     ruleID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, _, err := p.service.DfwRuleDeleteUsingDELETE(params)

	if err != nil {
		var badRequest *firewall_rules_e_w.DfwRuleDeleteUsingDELETEBadRequest
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
