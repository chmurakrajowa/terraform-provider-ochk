package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/networks"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"net/http"
)

type NetworksProxy struct {
	httpClient *http.Client
	service    networks.ClientService
}

func (p *NetworksProxy) Read(ctx context.Context, NetworkID string) (*models.VCSNetworkInstance, error) {
	params := &networks.VcsVirtualMachineGroupGetUsingGETParams{
		NetworkID:  NetworkID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.VcsVirtualMachineGroupGetUsingGET(params)
	if err != nil {
		var notFound *networks.VcsVirtualMachineGroupGetUsingGETNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading vcs network: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving vcs network failed: %s", response.Payload.Messages)
	}

	return response.Payload.VcsNetworkInstance, nil
}

func (p *NetworksProxy) ListByName(ctx context.Context, name string) ([]*models.VCSNetworkInstance, error) {
	params := &networks.VcsVirtualMachineListUsingGETParams{
		Name:       &name,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.VcsVirtualMachineListUsingGET(params)
	if err != nil {
		return nil, fmt.Errorf("error while listing vcs networks: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing vcs networks failed: %s", response.Payload.Messages)
	}

	return response.Payload.VcsNetworkInstanceCollection, nil
}
