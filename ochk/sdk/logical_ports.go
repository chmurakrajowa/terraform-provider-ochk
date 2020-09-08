package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/logical_ports"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"net/http"
)

type LogicalPortsProxy struct {
	httpClient *http.Client
	service    logical_ports.ClientService
}

func (p *LogicalPortsProxy) Read(ctx context.Context, ipSetID string) (*models.LogicalPort, error) {
	params := &logical_ports.LogicalPortGetUsingGETParams{
		LogicalPortID: ipSetID,
		Context:       ctx,
		HTTPClient:    p.httpClient,
	}

	response, err := p.service.LogicalPortGetUsingGET(params)
	if err != nil {
		var notFound *logical_ports.LogicalPortGetUsingGETNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading logical port: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving logical port failed: %s", response.Payload.Messages)
	}

	return response.Payload.LogicalPort, nil
}

func (p *LogicalPortsProxy) ListByDisplayName(ctx context.Context, displayName string) ([]*models.LogicalPort, error) {
	params := &logical_ports.LogicalPortListUsingGET1Params{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	response, err := p.service.LogicalPortListUsingGET1(params)
	if err != nil {
		return nil, fmt.Errorf("error while listing logical ports by display name %s: %w", displayName, err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing logical ports by display name %s failed: %s", displayName, response.Payload.Messages)
	}

	return response.Payload.LogicalPortCollection, nil
}
