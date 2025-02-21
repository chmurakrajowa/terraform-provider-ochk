package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/nat_rule"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type NatProxy struct {
	httpClient *http.Client
	service    nat_rule.ClientService
}

func (p *NatProxy) Read(ctx context.Context, natRuleID strfmt.UUID) (*models.NATRuleInstance, error) {
	params := &nat_rule.GetNetworkNatRulesRuleIDParams{
		RuleID:     natRuleID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkNatRulesRuleID(params)
	mutex.Unlock()
	if err != nil {
		var notFound *nat_rule.DeleteNetworkNatRulesRuleIDNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading nats: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving nats failed: %s", response.Payload.Messages)
	}

	return response.Payload.NatRuleInstance, nil
}

func (p *NatProxy) ListNatsByName(ctx context.Context, displayName string) ([]*models.NATRuleInstance, error) {
	params := &nat_rule.GetNetworkNatRulesParams{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkNatRules(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing nats: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing nats failed: %s", response.Payload.Messages)
	}

	return response.Payload.NatRuleInstances, nil
}
func (p *NatProxy) List(ctx context.Context) ([]*models.NATRuleInstance, error) {
	params := &nat_rule.GetNetworkNatRulesParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkNatRules(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing nats: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing nats failed: %s", response.Payload.Messages)
	}

	return response.Payload.NatRuleInstances, nil
}

func (p *NatProxy) CreateNat(ctx context.Context, natRuleInstance *models.NATRuleInstance) (*models.RequestInstance, error) {
	if err := natRuleInstance.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating nat struct: %w", err)
	}
	params := &nat_rule.PutNetworkNatRulesParams{
		Body:       natRuleInstance,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}
	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutNetworkNatRules(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating nat: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating nat failed: %s", put.Payload.Messages)
	}

	return put.Payload.RequestInstance, nil
}

func (p *NatProxy) Update(ctx context.Context, natRuleInstance *models.NATRuleInstance) (*models.RequestInstance, error) {
	if err := natRuleInstance.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating nat struct: %w", err)
	}
	params := &nat_rule.PutNetworkNatRulesRuleIDParams{
		RuleID:     natRuleInstance.RuleID,
		Body:       natRuleInstance,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutNetworkNatRulesRuleID(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifying nat: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying nat failed: %s", put.Payload.Messages)
	}

	return put.Payload.RequestInstance, nil
}

func (p *NatProxy) Delete(ctx context.Context, ruleID strfmt.UUID) error {
	params := &nat_rule.DeleteNetworkNatRulesRuleIDParams{
		RuleID:     ruleID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.DeleteNetworkNatRulesRuleID(params)

	if err != nil {
		var badRequest *nat_rule.DeleteNetworkNatRulesRuleIDBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return &NotFoundError{Err: err}
		}

		return fmt.Errorf("error while deleting nat: %w", err)
	}

	if !response.Payload.Success {
		return fmt.Errorf("deleting nat failed: %s", response.Payload.Messages)
	}

	return nil
}
