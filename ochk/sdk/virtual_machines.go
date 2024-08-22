package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/virtual_machine"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type VirtualMachinesProxy struct {
	httpClient *http.Client
	service    virtual_machine.ClientService
}

func (p *VirtualMachinesProxy) Create(ctx context.Context, virtualMachine *models.VirtualMachineInstance) (*models.RequestInstance, error) {
	if err := virtualMachine.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating virtual machine struct: %w", err)
	}

	params := &virtual_machine.PutVcsVirtualMachinesParams{
		Body:       virtualMachine,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutVcsVirtualMachines(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating virtual machine: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating virtual machine failed: %s", put.Payload.Messages)
	}

	return put.Payload.RequestInstance, nil
}

func (p *VirtualMachinesProxy) Update(ctx context.Context, virtualMachine *models.VirtualMachineInstance) (*models.RequestInstance, error) {
	if err := virtualMachine.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating virtual machine struct: %w", err)
	}

	params := &virtual_machine.PutVcsVirtualMachinesVirtualMachineIDParams{
		VirtualMachineID: virtualMachine.VirtualMachineID,
		Body:             virtualMachine,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}
	mutex := sync.Mutex{}
	mutex.Lock()
	put, err := p.service.PutVcsVirtualMachinesVirtualMachineID(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifying virtual machine: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying virtual machine failed: %s", put.Payload.Messages)
	}

	return put.Payload.RequestInstance, nil
}

func (p *VirtualMachinesProxy) Read(ctx context.Context, VirtualMachineID strfmt.UUID) (*models.VirtualMachineInstance, error) {
	params := &virtual_machine.GetVcsVirtualMachinesVirtualMachineIDParams{
		VirtualMachineID: VirtualMachineID,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetVcsVirtualMachinesVirtualMachineID(params)
	mutex.Unlock()

	if err != nil {
		var notFound *virtual_machine.GetVcsVirtualMachinesVirtualMachineIDNotFound
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

func (p *VirtualMachinesProxy) ListByDisplayName(ctx context.Context, displayName string) ([]*models.VirtualMachineInstance, error) {
	params := &virtual_machine.GetVcsVirtualMachinesParams{
		DisplayName: &displayName,
		Context:     ctx,
		HTTPClient:  p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetVcsVirtualMachines(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing virtual machines: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing virtual machines failed: %s", response.Payload.Messages)
	}

	return response.Payload.VcsVirtualMachineInstanceCollection, nil
}

func (p *VirtualMachinesProxy) List(ctx context.Context) ([]*models.VirtualMachineInstance, error) {
	params := &virtual_machine.GetVcsVirtualMachinesParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetVcsVirtualMachines(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing virtual machines: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing virtual machines failed: %s", response.Payload.Messages)
	}

	return response.Payload.VcsVirtualMachineInstanceCollection, nil
}

func (p *VirtualMachinesProxy) Delete(ctx context.Context, virtualMachineID strfmt.UUID) (*models.RequestInstance, error) {
	params := &virtual_machine.DeleteVcsVirtualMachinesVirtualMachineIDParams{
		VirtualMachineID: virtualMachineID,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	response, err := p.service.DeleteVcsVirtualMachinesVirtualMachineID(params)

	if err != nil {
		var badRequest *virtual_machine.DeleteVcsVirtualMachinesVirtualMachineIDBadRequest
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
