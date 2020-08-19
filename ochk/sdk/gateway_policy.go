package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/gateway_policies"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"net/http"
)

type GatewayPolicyProxy struct {
	httpClient *http.Client
	service    gateway_policies.ClientService
}

func (p *GatewayPolicyProxy) Read(ctx context.Context, gatewayPolicyID string) (*models.GatewayPolicy, error) {
	params := &gateway_policies.GatewayPolicyGetUsingGETParams{
		GatewayPolicyID: gatewayPolicyID,
		Context:         ctx,
		HTTPClient:      p.httpClient,
	}

	response, err := p.service.GatewayPolicyGetUsingGET(params)
	if err != nil {
		var notFound *gateway_policies.GatewayPolicyGetUsingGETNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading gateway policy: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving gateway policy failed: %s", response.Payload.Messages)
	}

	return response.Payload.GatewayPolicy, nil
}

func (p *GatewayPolicyProxy) ListByDisplayName(ctx context.Context, displayName string) ([]*models.GatewayPolicy, error) {
	params := &gateway_policies.GatewayPolicyListUsingGETParams{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	response, err := p.service.GatewayPolicyListUsingGET(params)
	if err != nil {
		return nil, fmt.Errorf("error while listing gateway policies: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing gateway policies failed: %s", response.Payload.Messages)
	}

	return response.Payload.GatewayPolicyCollection, nil
}
