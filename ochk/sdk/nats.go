package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/n_a_t_rules"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type NatProxy struct {
	httpClient *http.Client
	service    n_a_t_rules.ClientService
}

func (p *NatProxy) Read(ctx context.Context, natRuleID string) (*models.NATRuleInstance, error) {
	params := &n_a_t_rules.NatRuleGetUsingGETParams{
		RuleID:     natRuleID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.NatRuleGetUsingGET(params)
	mutex.Unlock()
	if err != nil {
		var notFound *n_a_t_rules.NatRuleGetUsingGETNotFound
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
	params := &n_a_t_rules.NatRuleListUsingGETParams{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.NatRuleListUsingGET(params)
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
	params := &n_a_t_rules.NatRuleListUsingGETParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.NatRuleListUsingGET(params)
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
	params := &n_a_t_rules.NatRuleCreateUsingPUTParams{
		NatRuleInstance: natRuleInstance,
		Context:         ctx,
		HTTPClient:      p.httpClient,
	}
	mutex := sync.Mutex{}
	mutex.Lock()
	_, put, err := p.service.NatRuleCreateUsingPUT(params)
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
	params := &n_a_t_rules.NatRuleUpdateUsingPUTParams{
		RuleID:          natRuleInstance.RuleID,
		NatRuleInstance: natRuleInstance,
		Context:         ctx,
		HTTPClient:      p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.NatRuleUpdateUsingPUT(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifying nat: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying nat failed: %s", put.Payload.Messages)
	}

	return put.Payload.RequestInstance, nil
}

func (p *NatProxy) Delete(ctx context.Context, ruleID string) error {
	params := &n_a_t_rules.NatRuleDeleteUsingDELETEParams{
		RuleID:     ruleID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.NatRuleDeleteUsingDELETE(params)

	if err != nil {
		var badRequest *n_a_t_rules.NatRuleDeleteUsingDELETEBadRequest
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
