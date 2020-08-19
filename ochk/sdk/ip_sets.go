package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/ip_sets"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"net/http"
)

type IPSetsProxy struct {
	httpClient *http.Client
	service    ip_sets.ClientService
}

func (p *IPSetsProxy) Read(ctx context.Context, ipSetID string) (*models.IPSet, error) {
	params := &ip_sets.IPSetGetUsingGETParams{
		IPSetID:    ipSetID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.IPSetGetUsingGET(params)
	if err != nil {
		var notFound *ip_sets.IPSetGetUsingGETNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading IP set: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving IP set failed: %s", response.Payload.Messages)
	}

	return response.Payload.IPSet, nil
}

func (p *IPSetsProxy) ListByDisplayName(ctx context.Context, displayName string) ([]*models.IPSet, error) {
	params := &ip_sets.IPSetListUsingGETParams{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	response, err := p.service.IPSetListUsingGET(params)
	if err != nil {
		return nil, fmt.Errorf("error while listing IP sets by display name %s: %w", displayName, err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing IP sets by display name %s failed: %s", displayName, response.Payload.Messages)
	}

	return response.Payload.IPSetCollection, nil
}
