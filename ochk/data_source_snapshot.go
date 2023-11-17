package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSnapshot() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSnapshotRead,

		Schema: map[string]*schema.Schema{
			"virtual_machine_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"snapshot_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"snapshot_description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"child_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"power_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSnapshotRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	proxy := meta.(*sdk.Client).Snapshots

	snapshotName := d.Get("display_name").(string)
	virtualMachineID := d.Get("virtual_machine_id").(string)

	snapshots, err := proxy.ListSnapshotsByName(ctx, virtualMachineID, snapshotName)

	if err != nil {
		return diag.Errorf("error while listing snapshots: %+v", err)
	}

	if len(snapshots) < 1 {
		return diag.Errorf("no snapshot for name: %s", snapshotName)
	}

	if len(snapshots) > 1 {
		return diag.Errorf("more than one snapshot with name: %s found!", snapshotName)
	}

	d.SetId(fmt.Sprint(snapshots[0].SnapshotID))

	if err := d.Set("display_name", snapshots[0].SnapshotName); err != nil {
		return diag.Errorf("error setting display name: %+v", err)
	}
	if err := d.Set("virtual_machine_id", snapshots[0].VirtualMachineID); err != nil {
		return diag.Errorf("error setting virtual_machine_id: %+v", err)
	}
	if err := d.Set("snapshot_description", snapshots[0].SnapshotDescription); err != nil {
		return diag.Errorf("error setting snapshot_description: %+v", err)
	}
	if err := d.Set("snapshot_id", snapshots[0].SnapshotID); err != nil {
		return diag.Errorf("error setting snapshot_id: %+v", err)
	}
	if err := d.Set("power_state", snapshots[0].PowerState); err != nil {
		return diag.Errorf("error setting power_state: %+v", err)
	}
	if err := d.Set("parent_id", snapshots[0].ParentSnapshotID); err != nil {
		return diag.Errorf("error setting parent_id: %+v", err)
	}
	if snapshots[0].ChildSnapshots != nil {
		if err := d.Set("child_id", snapshots[0].ChildSnapshots[0].SnapshotID); err != nil {
			return diag.Errorf("error setting child_id: %+v", err)
		}
	}

	return nil
}
