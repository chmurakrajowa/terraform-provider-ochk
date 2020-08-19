package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSecurityGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: datSourceSecurityGroupRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"members": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datSourceSecurityGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).SecurityGroups

	displayName := d.Get("display_name").(string)

	securityGroups, err := proxy.ListByDisplayName(ctx, displayName)
	if err != nil {
		return diag.Errorf("error while listing security group: %+v", err)
	}

	if len(securityGroups) < 1 {
		return diag.Errorf("no security group found for display_name: %s", displayName)
	}

	if len(securityGroups) > 1 {
		return diag.Errorf("more than one security group with display_name: %s found!", displayName)
	}

	d.SetId(securityGroups[0].ID)

	if err := d.Set("members", flattenSecurityGroupMembers(securityGroups[0].Members)); err != nil {
		return diag.Errorf("error setting members: %+v", err)
	}

	return nil
}
