package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/snapshots"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/virtual_machines"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/go-openapi/strfmt"
	"net/http"
)

type SnapshotsProxy struct {
	httpClient *http.Client
	service    virtual_machines.ClientService
}

func (p *SnapshotsProxy) Read(ctx context.Context, snapshotID string, virtualMachineID string) (*models.SnapshotInstance, error) {
	params := &virtual_machines.VcsVirtualMachineSnapshotGetUsingGETParams{
		SnapshotID:       snapshotID,
		VirtualMachineID: virtualMachineID,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	response, err := p.service.VcsVirtualMachineSnapshotGetUsingGET(params)

	if err != nil {
		var notFound *snapshots.SnapshotGetUsingGETNotFound

		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading snapshot: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving snapshot failed: %s", response.Payload.Messages)
	}

	return response.Payload.SnapshotInstance, nil
}

func (p *SnapshotsProxy) ListSnapshotsByName(ctx context.Context, virtualMachineID string, snapshotName string) ([]*models.SnapshotInstance, error) {
	params := &virtual_machines.VcsVirtualMachineSnapshotListUsingGETParams{
		VirtualMachineID: virtualMachineID,
		DisplayName:      &snapshotName,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	response, err := p.service.VcsVirtualMachineSnapshotListUsingGET(params)

	if err != nil {
		return nil, fmt.Errorf("error while listing snapshots: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing snapshots failed: %s", response.Payload.Messages)
	}
	return response.Payload.SnapshotInstanceCollection, nil
}

func (p *SnapshotsProxy) ListSnapshots(ctx context.Context, virtualMachineID string) ([]*models.SnapshotInstance, error) {
	params := &virtual_machines.VcsVirtualMachineSnapshotListUsingGETParams{
		VirtualMachineID: virtualMachineID,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	response, err := p.service.VcsVirtualMachineSnapshotListUsingGET(params)

	if err != nil {
		return nil, fmt.Errorf("error while listing snapshots: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("Listing snapshots failed: %s", response.Payload.Messages)
	}

	return response.Payload.SnapshotInstanceCollection, nil
}

func (p *SnapshotsProxy) Create(ctx context.Context, virtualMachineID string, ram *bool, snapshot *models.SnapshotInstance) (*models.SnapshotInstance, error) {
	if err := snapshot.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating snapshot struct: %w", err)
	}

	params := &virtual_machines.VcsVirtualMachineSnapshotCreateUsingPUTParams{
		VirtualMachineID: virtualMachineID,
		SnapshotInstance: snapshot,
		RAMSnapshot:      ram,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	put, _, err := p.service.VcsVirtualMachineSnapshotCreateUsingPUT(params)
	if err != nil {
		return nil, fmt.Errorf("error while creating snapshot: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating snapshot failed: %s", put.Payload.Messages)
	}

	return put.Payload.SnapshotInstance, nil
}

func (p *SnapshotsProxy) Delete(ctx context.Context, virtualMachineID string, snapshotID string) error {

	params := &virtual_machines.VcsVirtualMachineSnapshotDeleteUsingDELETEParams{
		VirtualMachineID: virtualMachineID,
		SnapshotID:       snapshotID,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	response, err := p.service.VcsVirtualMachineSnapshotDeleteUsingDELETE(params)
	if err != nil {
		var badRequest *snapshots.SnapshotGetUsingGETBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return &NotFoundError{Err: err}
		}

		return fmt.Errorf("error while deleting snapshot: %w", err)
	}

	if !response.Payload.Success {
		return fmt.Errorf("deleting snapshot failed: %s", response.Payload.Messages)
	}

	return nil
}
