package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLogicalPort() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLogicalPortRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceLogicalPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).LogicalPorts

	displayName := d.Get("display_name").(string)

	logicalPorts, err := proxy.ListByDisplayName(ctx, displayName)
	if err != nil {
		return diag.Errorf("error while listing logical ports: %+v", err)
	}

	if len(logicalPorts) < 1 {
		return diag.Errorf("no logical port found for display_name: %s", displayName)
	}

	if len(logicalPorts) > 1 {
		return diag.Errorf("more than one logical port with display_name: %s found!", displayName)
	}

	d.SetId(logicalPorts[0].ID)

	return nil
}
