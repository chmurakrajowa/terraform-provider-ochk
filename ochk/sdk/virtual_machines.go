package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/virtual_machines"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"net/http"
)

type VirtualMachinesProxy struct {
	httpClient *http.Client
	service    virtual_machines.ClientService
}

func (p *VirtualMachinesProxy) Read(ctx context.Context, VirtualMachineID string) (*models.VirtualMachine, error) {
	params := &virtual_machines.VirtualMachineGetUsingGETParams{
		VirtualMachineID: VirtualMachineID,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	response, err := p.service.VirtualMachineGetUsingGET(params)
	if err != nil {
		var notFound *virtual_machines.VirtualMachineGetUsingGETNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading virtual machine: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving virtual machine failed: %s", response.Payload.Messages)
	}

	return response.Payload.VirtualMachine, nil
}

func (p *VirtualMachinesProxy) ListByDisplayName(ctx context.Context, displayName string) ([]*models.VirtualMachine, error) {
	params := &virtual_machines.VirtualMachineListUsingGETParams{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	response, err := p.service.VirtualMachineListUsingGET(params)
	if err != nil {
		return nil, fmt.Errorf("error while listing virtual machines: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing virtual machines failed: %s", response.Payload.Messages)
	}

	return response.Payload.VirtualMachineCollection, nil
}
