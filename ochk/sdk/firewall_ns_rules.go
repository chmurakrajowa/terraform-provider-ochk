package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/gfw_rule"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type FirewallSNRulesProxy struct {
	httpClient *http.Client
	service    gfw_rule.ClientService
}

func (p *FirewallSNRulesProxy) Create(ctx context.Context, routerID strfmt.UUID, rule *models.GfwRule) (*models.GfwRule, error) {
	if err := rule.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating firewall SN rule struct: %w", err)
	}

	params := &gfw_rule.PutNetworkRoutersRouterIDRulesSNParams{
		RouterID:   routerID,
		Body:       rule,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutNetworkRoutersRouterIDRulesSN(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating firewall SN rule: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating firewall SN rule failed: %s", put.Payload.Messages)
	}

	return put.Payload.GfwRule, nil
}

func (p *FirewallSNRulesProxy) Update(ctx context.Context, routerID strfmt.UUID, rule *models.GfwRule) (*models.GfwRule, error) {
	if err := rule.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating firewall SN rule struct: %w", err)
	}

	params := &gfw_rule.PutNetworkRoutersRouterIDRulesSNRuleIDParams{
		RouterID:   routerID,
		RuleID:     rule.RuleID,
		Body:       rule,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutNetworkRoutersRouterIDRulesSNRuleID(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifing firewall SN rule: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifing firewall SN rule failed: %s", put.Payload.Messages)
	}

	return put.Payload.GfwRule, nil
}

func (p *FirewallSNRulesProxy) Read(ctx context.Context, routerID strfmt.UUID, ruleID strfmt.UUID) (*models.GfwRule, error) {
	params := &gfw_rule.GetNetworkRoutersRouterIDRulesSNRuleIDParams{
		RuleID:     ruleID,
		RouterID:   routerID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkRoutersRouterIDRulesSNRuleID(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading firwall SN rule: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving firewall SN rule failed: %s", response.Payload.Messages)
	}

	return response.Payload.RuleInstance, nil
}

func (p *FirewallSNRulesProxy) ListByDisplayName(ctx context.Context, routerID strfmt.UUID, displayName string) ([]*models.GfwRule, error) {
	params := &gfw_rule.GetNetworkRoutersRouterIDRulesSNParams{
		RouterID:    routerID,
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkRoutersRouterIDRulesSN(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing firewall SN rule: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing firewall SN rule failed: %s", response.Payload.Messages)
	}

	return response.Payload.RuleInstances, nil
}

func (p *FirewallSNRulesProxy) List(ctx context.Context, routerID strfmt.UUID) ([]*models.GfwRule, error) {
	params := &gfw_rule.GetNetworkRoutersRouterIDRulesSNParams{
		RouterID:   routerID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkRoutersRouterIDRulesSN(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing firewall SN rule: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing firewall SN rule failed: %s", response.Payload.Messages)
	}

	return response.Payload.RuleInstances, nil
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
	params := &gfw_rule.DeleteNetworkRoutersRouterIDRulesSNRuleIDParams{
		RouterID:   routerID,
		RuleID:     ruleID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.DeleteNetworkRoutersRouterIDRulesSNRuleID(params)

	if err != nil {
		var badRequest *gfw_rule.DeleteNetworkRoutersRouterIDRulesSNRuleIDBadRequest
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
