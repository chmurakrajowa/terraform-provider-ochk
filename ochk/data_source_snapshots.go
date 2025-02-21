package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSnapshots() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSnapshotsRead,
		Schema: map[string]*schema.Schema{
			"virtual_machine_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"snapshots": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"snapshot_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"virtual_machine_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"child_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSnapshotsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	proxy := meta.(*sdk.Client).Snapshots
	virtualMachineID := strfmt.UUID(d.Get("virtual_machine_id").(string))

	snapshots, err := proxy.ListSnapshots(ctx, virtualMachineID)

	if err := d.Set("snapshots", flattenSnapshot(snapshots)); err != nil {
		return diag.Errorf("error setting snapshots list: %v", err)
	}

	if err != nil {
		return diag.Errorf("error while listing snapshots: %+v", err)
	}

	d.SetId("snapshots-list " + virtualMachineID.String())

	return nil
}
