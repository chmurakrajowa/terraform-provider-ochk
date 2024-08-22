package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVirtualMachine() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVirtualMachineRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
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

	if err := d.Set("display_name", virtualMachines[0].VirtualMachineName); err != nil {
		return diag.Errorf("error setting virtual machine display_name: %v", err)
	}

	if err := d.Set("folder_path", virtualMachines[0].FolderPath); err != nil {
		return diag.Errorf("error setting virtual machine folder_path: %v", err)
	}

	if err := d.Set("project_id", virtualMachines[0].ProjectID); err != nil {
		return diag.Errorf("error setting virtual machine project_id: %v", err)
	}
	d.SetId(virtualMachines[0].VirtualMachineID.String())

	return nil
}
