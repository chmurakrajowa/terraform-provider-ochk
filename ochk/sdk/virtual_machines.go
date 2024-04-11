package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/virtual_machines"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type VirtualMachinesProxy struct {
	httpClient *http.Client
	service    virtual_machines.ClientService
}

func (p *VirtualMachinesProxy) Create(ctx context.Context, virtualMachine *models.VcsVirtualMachineInstance) (*models.RequestInstance, error) {
	if err := virtualMachine.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating virtual machine struct: %w", err)
	}

	params := &virtual_machines.VcsVirtualMachineCreateUsingPUTParams{
		VirtualMachine: virtualMachine,
		Context:        ctx,
		HTTPClient:     p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, _, err := p.service.VcsVirtualMachineCreateUsingPUT(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating virtual machine: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating virtual machine failed: %s", put.Payload.Messages)
	}

	return put.Payload.RequestInstance, nil
}

func (p *VirtualMachinesProxy) Update(ctx context.Context, virtualMachine *models.VcsVirtualMachineInstance) (*models.RequestInstance, error) {
	if err := virtualMachine.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating virtual machine struct: %w", err)
	}

	params := &virtual_machines.VcsVirtualMachineUpdateUsingPUTParams{
		VirtualMachineID: virtualMachine.VirtualMachineID,
		VirtualMachine:   virtualMachine,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}
	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.VcsVirtualMachineUpdateUsingPUT(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifying virtual machine: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying virtual machine failed: %s", put.Payload.Messages)
	}

	return put.Payload.RequestInstance, nil
}

func (p *VirtualMachinesProxy) Read(ctx context.Context, VirtualMachineID string) (*models.VcsVirtualMachineInstance, error) {
	params := &virtual_machines.VcsVirtualMachineGetUsingGETParams{
		VirtualMachineID: VirtualMachineID,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.VcsVirtualMachineGetUsingGET(params)
	mutex.Unlock()

	if err != nil {
		var notFound *virtual_machines.VcsVirtualMachineGetUsingGETNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		//virtual_Machines.VcsVirtualMachineS

		return nil, fmt.Errorf("error while reading virtual machine: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving virtual machine failed: %s", response.Payload.Messages)
	}

	return response.Payload.VcsVirtualMachineInstance, nil
}

func (p *VirtualMachinesProxy) ListByDisplayName(ctx context.Context, displayName string) ([]*models.VcsVirtualMachineInstance, error) {
	params := &virtual_machines.VcsVirtualMachineListUsingGETParams{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.VcsVirtualMachineListUsingGET(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing virtual machines: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing virtual machines failed: %s", response.Payload.Messages)
	}

	return response.Payload.VcsVirtualMachineInstanceCollection, nil
}

func (p *VirtualMachinesProxy) List(ctx context.Context) ([]*models.VcsVirtualMachineInstance, error) {
	params := &virtual_machines.VcsVirtualMachineListUsingGETParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.VcsVirtualMachineListUsingGET(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing virtual machines: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing virtual machines failed: %s", response.Payload.Messages)
	}

	return response.Payload.VcsVirtualMachineInstanceCollection, nil
}

func (p *VirtualMachinesProxy) Delete(ctx context.Context, virtualMachineID string) (*models.RequestInstance, error) {
	params := &virtual_machines.VcsVirtualMachineDeleteUsingDELETEParams{
		VirtualMachineID: virtualMachineID,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	response, err := p.service.VcsVirtualMachineDeleteUsingDELETE(params)

	if err != nil {
		var badRequest *virtual_machines.VcsVirtualMachineDeleteUsingDELETEBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while deleting virtual machine: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("deleting virtual machine failed: %s", response.Payload.Messages)
	}

	return response.Payload.RequestInstance, nil
}
