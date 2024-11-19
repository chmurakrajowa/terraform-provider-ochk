package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSecurityGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: datSourceSecurityGroupsRead,

		Schema: map[string]*schema.Schema{
			"security_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"security_group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
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

func datSourceSecurityGroupsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).SecurityGroups

	securityGroups, err := proxy.List(ctx)
	if err != nil {
		return diag.Errorf("error while listing security groups: %+v", err)
	}

	if err := d.Set("security_groups", flattenSecurityGroups(securityGroups)); err != nil {
		return diag.Errorf("error setting security_groups: %v", err)
	}

	d.SetId("security-groups-list")

	return nil
}
