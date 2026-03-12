package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"net/http"
)

type SnapshotsProxy struct {
	httpClient *http.Client
	service    *openapi.VirtualMachineSnapshotAPIService
}

func (p *SnapshotsProxy) Read(ctx context.Context, snapshotID strfmt.UUID, virtualMachineID strfmt.UUID) (*openapi.SnapshotInstance, error) {

	action := p.service.VcsVirtualMachinesVirtualMachineIdSnapshotsSnapshotIdGet(ctx, string(virtualMachineID), string(snapshotID))
	response, _, err := action.Execute()

	if err != nil {
		return nil, fmt.Errorf("error while reading snapshot: %w", err)
	}

	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("retrieving snapshot failed: %s", response.Messages)
	}

	return response.SnapshotInstance, nil
}

func (p *SnapshotsProxy) ListSnapshotsByName(ctx context.Context, virtualMachineID strfmt.UUID, snapshotName string) ([]openapi.SnapshotInstance, error) {
	action := p.service.VcsVirtualMachinesVirtualMachineIdSnapshotsGet(ctx, string(virtualMachineID)).DisplayName(snapshotName)
	response, _, err := action.Execute()

	if err != nil {
		return nil, fmt.Errorf("error while listing snapshots: %w", err)
	}

	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing snapshots failed: %s", response.Messages)
	}
	return response.SnapshotInstanceCollection, nil
}

func (p *SnapshotsProxy) ListSnapshots(ctx context.Context, virtualMachineID strfmt.UUID) ([]openapi.SnapshotInstance, error) {
	action := p.service.VcsVirtualMachinesVirtualMachineIdSnapshotsGet(ctx, string(virtualMachineID))
	response, _, err := action.Execute()

	if err != nil {
		return nil, fmt.Errorf("error while listing snapshots: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing snapshots failed: %s", response.Messages)
	}

	return response.SnapshotInstanceCollection, nil
}

func (p *SnapshotsProxy) Create(ctx context.Context, virtualMachineID strfmt.UUID, ram bool, snapshot openapi.SnapshotInstance) (*openapi.SnapshotInstance, *http.Response, error) {

	action := p.service.VcsVirtualMachinesVirtualMachineIdSnapshotsPut(ctx, string(virtualMachineID)).SnapshotInstance(snapshot).RamSnapshot(ram)
	put, httpResponse, err := action.Execute()

	if httpResponse.StatusCode == 504 {
		return nil, httpResponse, fmt.Errorf("error while creating snapshot: Timeout error. %+v", err)
	}
	if err != nil {
		return nil, httpResponse, fmt.Errorf("error while creating snapshot: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, httpResponse, fmt.Errorf("creating snapshot failed: %s", put.Messages)
	}

	return put.SnapshotInstance, httpResponse, nil
}

/*func (p *SnapshotsProxy) Update(ctx context.Context, virtualMachineID strfmt.UUID, snapshotId strfmt.UUID, snapshot openapi.SnapshotInstance) (*openapi.SnapshotInstance, *http.Response, error) {

	action := p.service.VcsVirtualMachinesVirtualMachineIdSnapshotsSnapshotIdPut(ctx, string(virtualMachineID), string(snapshotId)).SnapshotInstance(snapshot)
	put, httpResponse, err := action.Execute()

	if err != nil {
		return nil, httpResponse, fmt.Errorf("error while updating snapshot: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, httpResponse, fmt.Errorf("updating snapshot failed: %s", put.Messages)
	}

	return put.SnapshotInstance, httpResponse, nil
}*/

func (p *SnapshotsProxy) Delete(ctx context.Context, virtualMachineID strfmt.UUID, snapshotID strfmt.UUID) error {

	action := p.service.VcsVirtualMachinesVirtualMachineIdSnapshotsSnapshotIdDelete(ctx, string(virtualMachineID), string(snapshotID))
	response, _, err := action.Execute()
	if err != nil {
		return fmt.Errorf("error while deleting snapshot: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return fmt.Errorf("deleting snapshot failed: %s", response.Messages)
	}

	return nil
}
