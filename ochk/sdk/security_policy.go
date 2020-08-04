package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/security_policies"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
	"net/http"
)

type SecurityPolicyProxy struct {
	httpClient *http.Client
	service    security_policies.ClientService
}

func (p *SecurityPolicyProxy) Read(ctx context.Context, securityPolicyID string) (*models.SecurityPolicy, error) {
	params := &security_policies.SecurityPolicyGetUsingGETParams{
		SecurityPolicyID: securityPolicyID,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	response, err := p.service.SecurityPolicyGetUsingGET(params)
	if err != nil {
		var notFound *security_policies.SecurityPolicyGetUsingGETNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading security policy: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving service failed: %s", response.Payload.Messages)
	}

	return response.Payload.SecurityPolicy, nil
}

func (p *SecurityPolicyProxy) ListByDisplayName(ctx context.Context, displayName string) ([]*models.SecurityPolicy, error) {
	//TODO nie ma jak przekazać display name
	params := &security_policies.SecurityPolicyListUsingGETParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.SecurityPolicyListUsingGET(params)
	if err != nil {
		return nil, fmt.Errorf("error while listing security policies: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing security policies failed: %s", response.Payload.Messages)
	}

	return response.Payload.SecurityPolicyCollection, nil
}
