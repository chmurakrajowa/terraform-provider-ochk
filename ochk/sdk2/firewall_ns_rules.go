package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type FirewallSNRulesProxy struct {
	httpClient *http.Client
	service    *openapi.GfwRuleAPIService
}

func (p *FirewallSNRulesProxy) Create(ctx context.Context, routerID strfmt.UUID, rule openapi.GfwRule) (*openapi.GfwRule, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkRoutersRouterIdRulesSNPut(ctx, string(routerID)).GfwRule(rule)
	put, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating firewall SN rule: %w", err)
	}

	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("creating firewall SN rule failed: %s", put.Messages)
	}

	return put.GfwRule, nil
}

func (p *FirewallSNRulesProxy) Update(ctx context.Context, routerID strfmt.UUID, rule *openapi.GfwRule) (*openapi.GfwRule, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkRoutersRouterIdRulesSNRuleIdPut(ctx, string(routerID), rule.GetRuleId())
	put, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifing firewall SN rule: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("modifing firewall SN rule failed: %s", put.Messages)
	}

	return put.GfwRule, nil
}

func (p *FirewallSNRulesProxy) Read(ctx context.Context, routerID strfmt.UUID, ruleID strfmt.UUID) (*openapi.GfwRule, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkRoutersRouterIdRulesSNRuleIdGet(ctx, string(routerID), string(ruleID))
	response, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading firwall SN rule: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("retrieving firewall SN rule failed: %s", response.Messages)
	}

	return response.RuleInstance, nil
}

func (p *FirewallSNRulesProxy) ListByDisplayName(ctx context.Context, routerID strfmt.UUID, displayName string) ([]openapi.GfwRule, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkRoutersRouterIdRulesSNGet(ctx, string(routerID)).DisplayName(displayName)
	response, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing firewall SN rule: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing firewall SN rule failed: %s", response.Messages)
	}

	return response.RuleInstances, nil
}

func (p *FirewallSNRulesProxy) List(ctx context.Context, routerID strfmt.UUID) ([]openapi.GfwRule, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkRoutersRouterIdRulesSNGet(ctx, string(routerID))
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing firewall SN rule: %w", err)
	}

	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing firewall SN rule failed: %s", response.Messages)
	}

	return response.RuleInstances, nil
}

func (p *FirewallSNRulesProxy) Exists(ctx context.Context, routerID strfmt.UUID, ruleID strfmt.UUID) (bool, error) {
	if _, err := p.Read(ctx, routerID, ruleID); err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}

		return false, fmt.Errorf("error while reading firewall SN rule: %w", err)
	}

	return true, nil
}

func (p *FirewallSNRulesProxy) Delete(ctx context.Context, routerID strfmt.UUID, ruleID strfmt.UUID) error {

	action := p.service.NetworkRoutersRouterIdRulesSNRuleIdDelete(ctx, string(routerID), string(ruleID))
	response, _, err := action.Execute()

	if err != nil {
		return fmt.Errorf("error while deleting firewall SN rule: %w", err)
	}

	isSuccess := *response.Success

	if !isSuccess {
		return fmt.Errorf("deleting firewall SN rule failed: %s", response.Messages)
	}

	return nil
}
