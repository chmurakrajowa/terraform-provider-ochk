package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type NatProxy struct {
	httpClient *http.Client
	service    *openapi.NatRuleAPIService
}

func (p *NatProxy) Read(ctx context.Context, natRuleID strfmt.UUID) (*openapi.NATRuleInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkNatRulesRuleIdGet(ctx, string(natRuleID))
	response, _, err := action.Execute()
	mutex.Unlock()
	if err != nil {
		return nil, fmt.Errorf("error while reading nats: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("retrieving nats failed: %s", response.Messages)
	}

	return response.NatRuleInstance, nil
}

func (p *NatProxy) ListNatsByName(ctx context.Context, displayName string) ([]openapi.NATRuleInstance, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkNatRulesGet(ctx).DisplayName(displayName)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing nats: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing nats failed: %s", response.Messages)
	}

	return response.NatRuleInstances, nil
}
func (p *NatProxy) List(ctx context.Context) ([]openapi.NATRuleInstance, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkNatRulesGet(ctx)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing nats: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing nats failed: %s", response.Messages)
	}

	return response.NatRuleInstances, nil
}

func (p *NatProxy) CreateNat(ctx context.Context, natRuleInstance openapi.NATRuleInstance) (*openapi.RequestInstance, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkNatRulesPut(ctx).NATRuleInstance(natRuleInstance)
	put, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating nat: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("creating nat failed: %s", put.Messages)
	}

	return put.RequestInstance, nil
}

func (p *NatProxy) Update(ctx context.Context, natRuleInstance *openapi.NATRuleInstance) (*openapi.RequestInstance, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkNatRulesRuleIdPut(ctx, natRuleInstance.GetRuleId())
	put, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifying nat: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("modifying nat failed: %s", put.Messages)
	}

	return put.RequestInstance, nil
}

func (p *NatProxy) Delete(ctx context.Context, ruleID strfmt.UUID) error {

	action := p.service.NetworkNatRulesRuleIdDelete(ctx, string(ruleID))
	response, _, err := action.Execute()

	if err != nil {
		return fmt.Errorf("error while deleting nat: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return fmt.Errorf("deleting nat failed: %s", response.Messages)
	}

	return nil
}
