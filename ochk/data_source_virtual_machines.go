package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVirtualMachines() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVirtualMachinesRead,

		Schema: map[string]*schema.Schema{
			"virtual_machines": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"virtual_machine_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"folder_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceVirtualMachinesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).VirtualMachines

	virtualMachines, err := proxy.List(ctx)
	if err != nil {
		return diag.Errorf("error while listing virtual machines: %+v", err)
	}

	if err := d.Set("virtual_machines", flattenVirtualMachines(virtualMachines)); err != nil {
		return diag.Errorf("error setting virtual machines list: %v", err)
	}

	d.SetId("virtual-machines-list")

	return nil
}
