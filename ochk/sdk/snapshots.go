package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/virtual_machine_snapshot"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"net/http"
)

type SnapshotsProxy struct {
	httpClient *http.Client
	service    virtual_machine_snapshot.ClientService
}

func (p *SnapshotsProxy) Read(ctx context.Context, snapshotID strfmt.UUID, virtualMachineID strfmt.UUID) (*models.SnapshotInstance, error) {
	params := &virtual_machine_snapshot.GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams{
		SnapshotID:       snapshotID,
		VirtualMachineID: virtualMachineID,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	response, err := p.service.GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotID(params)

	if err != nil {
		var notFound *virtual_machine_snapshot.GetVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDNotFound

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

func (p *SnapshotsProxy) ListSnapshotsByName(ctx context.Context, virtualMachineID strfmt.UUID, snapshotName string) ([]*models.SnapshotInstance, error) {
	params := &virtual_machine_snapshot.GetVcsVirtualMachinesVirtualMachineIDSnapshotsParams{
		VirtualMachineID: virtualMachineID,
		DisplayName:      &snapshotName,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	response, err := p.service.GetVcsVirtualMachinesVirtualMachineIDSnapshots(params)

	if err != nil {
		return nil, fmt.Errorf("error while listing snapshots: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing snapshots failed: %s", response.Payload.Messages)
	}
	return response.Payload.SnapshotInstanceCollection, nil
}

func (p *SnapshotsProxy) ListSnapshots(ctx context.Context, virtualMachineID strfmt.UUID) ([]*models.SnapshotInstance, error) {
	params := &virtual_machine_snapshot.GetVcsVirtualMachinesVirtualMachineIDSnapshotsParams{
		VirtualMachineID: virtualMachineID,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	response, err := p.service.GetVcsVirtualMachinesVirtualMachineIDSnapshots(params)

	if err != nil {
		return nil, fmt.Errorf("error while listing snapshots: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing snapshots failed: %s", response.Payload.Messages)
	}

	return response.Payload.SnapshotInstanceCollection, nil
}

func (p *SnapshotsProxy) Create(ctx context.Context, virtualMachineID strfmt.UUID, ram *bool, snapshot *models.SnapshotInstance) (*models.SnapshotInstance, error) {
	if err := snapshot.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating snapshot struct: %w", err)
	}

	params := &virtual_machine_snapshot.PutVcsVirtualMachinesVirtualMachineIDSnapshotsParams{
		VirtualMachineID: virtualMachineID,
		Body:             snapshot,
		//
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	put, err := p.service.PutVcsVirtualMachinesVirtualMachineIDSnapshots(params)
	if err != nil {
		return nil, fmt.Errorf("error while creating snapshot: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating snapshot failed: %s", put.Payload.Messages)
	}

	return put.Payload.SnapshotInstance, nil
}

func (p *SnapshotsProxy) Delete(ctx context.Context, virtualMachineID strfmt.UUID, snapshotID strfmt.UUID) error {

	params := &virtual_machine_snapshot.DeleteVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDParams{
		VirtualMachineID: virtualMachineID,
		SnapshotID:       snapshotID,
		Context:          ctx,
		HTTPClient:       p.httpClient,
	}

	response, err := p.service.DeleteVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotID(params)
	if err != nil {
		var badRequest *virtual_machine_snapshot.DeleteVcsVirtualMachinesVirtualMachineIDSnapshotsSnapshotIDBadRequest
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
