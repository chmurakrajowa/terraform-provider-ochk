package ochk

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
)

func dataSourceVirtualMachine() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVirtualMachineRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceVirtualMachineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).VirtualMachines

	displayName := d.Get("display_name").(string)

	virtualMachines, err := proxy.ListByDisplayName(ctx, displayName)
	if err != nil {
		return diag.Errorf("error while listing virtual machines: %+v", err)
	}

	if len(virtualMachines) < 1 {
		return diag.Errorf("no virtual machine found for display_name: %s", displayName)
	}

	if len(virtualMachines) > 1 {
		return diag.Errorf("more than one virtual machine with display_name: %s found!", displayName)
	}

	d.SetId(virtualMachines[0].VirtualMachineID)

	return nil
}
