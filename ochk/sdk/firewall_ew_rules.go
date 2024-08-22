package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/dfw_rule"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type FirewallEWRulesProxy struct {
	httpClient *http.Client
	service    dfw_rule.ClientService
}

func (p *FirewallEWRulesProxy) Create(ctx context.Context, routerID strfmt.UUID, rule *models.DfwRule) (*models.DfwRule, error) {
	if err := rule.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating firewall EW rule struct: %w", err)
	}

	params := &dfw_rule.PutNetworkRoutersRouterIDRulesEWParams{
		RouterID:   routerID,
		Body:       rule,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutNetworkRoutersRouterIDRulesEW(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating firewall EW rule: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating firewall EW rule failed: %s", put.Payload.Messages)
	}

	return put.Payload.DfwRule, nil
}

func (p *FirewallEWRulesProxy) Read(ctx context.Context, routerID strfmt.UUID, ruleID strfmt.UUID) (*models.DfwRule, error) {
	params := &dfw_rule.GetNetworkRoutersRouterIDRulesEWRuleIDParams{
		RuleID:     ruleID,
		RouterID:   routerID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkRoutersRouterIDRulesEWRuleID(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading firwall EW rule: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving firewall EW rule failed: %s", response.Payload.Messages)
	}

	return response.Payload.RuleInstance, nil
}

func (p *FirewallEWRulesProxy) Update(ctx context.Context, routerID strfmt.UUID, rule *models.DfwRule) (*models.DfwRule, error) {
	if err := rule.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating firewall EW rule struct: %w", err)
	}

	params := &dfw_rule.PutNetworkRoutersRouterIDRulesEWRuleIDParams{
		RouterID:   routerID,
		RuleID:     rule.RuleID,
		Body:       rule,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutNetworkRoutersRouterIDRulesEWRuleID(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while updating firewall EW rule: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating updating EW rule failed: %s", put.Payload.Messages)
	}

	return put.Payload.DfwRule, nil
}

func (p *FirewallEWRulesProxy) ListByDisplayName(ctx context.Context, routerID strfmt.UUID, displayName string) ([]*models.DfwRule, error) {
	params := &dfw_rule.GetNetworkRoutersRouterIDRulesEWParams{
		RouterID:    routerID,
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkRoutersRouterIDRulesEW(params)
	mutex.Unlock()
	if err != nil {
		return nil, fmt.Errorf("error while listing firewall EW rule: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing firewall EW rule failed: %s", response.Payload.Messages)
	}

	return response.Payload.RuleInstances, nil
}

func (p *FirewallEWRulesProxy) List(ctx context.Context, routerID strfmt.UUID) ([]*models.DfwRule, error) {
	params := &dfw_rule.GetNetworkRoutersRouterIDRulesEWParams{
		RouterID:   routerID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkRoutersRouterIDRulesEW(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing firewall EW rule: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing firewall EW rule failed: %s", response.Payload.Messages)
	}

	return response.Payload.RuleInstances, nil
}

func (p *FirewallEWRulesProxy) Exists(ctx context.Context, routerID strfmt.UUID, ruleID strfmt.UUID) (bool, error) {
	if _, err := p.Read(ctx, routerID, ruleID); err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}

		return false, fmt.Errorf("error while reading firewall EW rule: %w", err)
	}

	return true, nil
}

func (p *FirewallEWRulesProxy) Delete(ctx context.Context, routerID strfmt.UUID, ruleID strfmt.UUID) error {
	params := &dfw_rule.DeleteNetworkRoutersRouterIDRulesEWRuleIDParams{
		RouterID:   routerID,
		RuleID:     ruleID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.DeleteNetworkRoutersRouterIDRulesEWRuleID(params)

	if err != nil {
		var badRequest *dfw_rule.DeleteNetworkRoutersRouterIDRulesEWRuleIDBadRequest
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
