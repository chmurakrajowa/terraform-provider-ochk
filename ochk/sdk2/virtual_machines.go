package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
)

type VirtualMachinesProxy struct {
	httpClient *http.Client
	service    *openapi.VirtualMachineAPIService
}

func (p *VirtualMachinesProxy) Create(ctx context.Context, virtualMachine openapi.VirtualMachineInstance) (*openapi.RequestInstance, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.VcsVirtualMachinesPut(ctx).VirtualMachineInstance(virtualMachine)
	put, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating virtual machine: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("creating virtual machine failed: %s", put.Messages)
	}

	return put.RequestInstance, nil
}

func (p *VirtualMachinesProxy) Update(ctx context.Context, virtualMachine openapi.VirtualMachineInstance) (*openapi.RequestInstance, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.VcsVirtualMachinesVirtualMachineIdPut(ctx, virtualMachine.GetVirtualMachineId()).VirtualMachineInstance(virtualMachine)
	put, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while modifying virtual machine: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("modifying virtual machine failed: %s", put.Messages)
	}

	return put.RequestInstance, nil
}

func (p *VirtualMachinesProxy) Read(ctx context.Context, VirtualMachineID strfmt.UUID) (*openapi.VirtualMachineInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.VcsVirtualMachinesVirtualMachineIdGet(ctx, string(VirtualMachineID))
	response, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading virtual machine: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("retrieving virtual machine failed: %s", response.Messages)
	}

	return response.VcsVirtualMachineInstance, nil
}

func (p *VirtualMachinesProxy) ListByDisplayName(ctx context.Context, displayName string) ([]openapi.VirtualMachineInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.VcsVirtualMachinesGet(ctx).DisplayName(displayName)
	response, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing virtual machines: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing virtual machines failed: %s", response.Messages)
	}

	return response.VcsVirtualMachineInstanceCollection, nil
}

func (p *VirtualMachinesProxy) List(ctx context.Context) ([]openapi.VirtualMachineInstance, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.VcsVirtualMachinesGet(ctx)
	response, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing virtual machines: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing virtual machines failed: %s", response.Messages)
	}

	return response.VcsVirtualMachineInstanceCollection, nil
}

func (p *VirtualMachinesProxy) Delete(ctx context.Context, virtualMachineID strfmt.UUID) (*openapi.RequestInstance, error) {

	action := p.service.VcsVirtualMachinesVirtualMachineIdDelete(ctx, string(virtualMachineID))
	response, _, err := action.Execute()

	if err != nil {
		return nil, fmt.Errorf("error while deleting virtual machine: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("deleting virtual machine failed: %s", response.Messages)
	}

	return response.RequestInstance, nil
}
