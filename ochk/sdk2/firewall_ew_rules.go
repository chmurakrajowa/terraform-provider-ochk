package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type FirewallEWRulesProxy struct {
	httpClient *http.Client
	service    *openapi.DfwRuleAPIService
}

func (p *FirewallEWRulesProxy) Create(ctx context.Context, routerID strfmt.UUID, rule openapi.DfwRule) (*openapi.DfwRule, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkRoutersRouterIdRulesEWPut(ctx, string(routerID)).DfwRule(rule)
	put, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating firewall EW rule: %w", err)
	}

	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("creating firewall EW rule failed: %s", put.Messages)
	}

	return put.DfwRule, nil
}

func (p *FirewallEWRulesProxy) Read(ctx context.Context, routerID strfmt.UUID, ruleID strfmt.UUID) (*openapi.DfwRule, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkRoutersRouterIdRulesEWRuleIdGet(ctx, string(routerID), string(ruleID))
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading firwall EW rule: %w", err)
	}

	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("retrieving firewall EW rule failed: %s", response.Messages)
	}

	return response.RuleInstance, nil
}

func (p *FirewallEWRulesProxy) Update(ctx context.Context, routerID strfmt.UUID, rule *openapi.DfwRule) (*openapi.DfwRule, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkRoutersRouterIdRulesEWRuleIdPut(ctx, string(routerID), rule.GetRuleId())
	put, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while updating firewall EW rule: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("creating updating EW rule failed: %s", put.Messages)
	}

	return put.DfwRule, nil
}

func (p *FirewallEWRulesProxy) ListByDisplayName(ctx context.Context, routerID strfmt.UUID, displayName string) ([]openapi.DfwRule, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkRoutersRouterIdRulesEWGet(ctx, string(routerID)).DisplayName(displayName)
	response, _, err := action.Execute()
	mutex.Unlock()
	if err != nil {
		return nil, fmt.Errorf("error while listing firewall EW rule: %w", err)
	}

	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing firewall EW rule failed: %s", response.Messages)
	}

	return response.RuleInstances, nil
}

func (p *FirewallEWRulesProxy) List(ctx context.Context, routerID strfmt.UUID) ([]openapi.DfwRule, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkRoutersRouterIdRulesEWGet(ctx, string(routerID))
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing firewall EW rule: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing firewall EW rule failed: %s", response.Messages)
	}

	return response.RuleInstances, nil
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
	action := p.service.NetworkRoutersRouterIdRulesEWRuleIdDelete(ctx, string(routerID), string(ruleID))
	response, _, err := action.Execute()

	if err != nil {
		return fmt.Errorf("error while deleting firewall EW rule: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return fmt.Errorf("deleting firewall EW rule failed: %s", response.Messages)
	}

	return nil
}
