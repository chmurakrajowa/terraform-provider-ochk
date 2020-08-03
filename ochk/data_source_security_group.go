package ochk

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk"
)

func dataSourceSecurityGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: datSourceServiceGroupRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"members": {
				Type:     schema.TypeList,
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

func datSourceServiceGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).SecurityGroups

	securityGroup, err := proxy.Read(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			d.SetId("")
			return diag.Errorf("security group with id %s not found: %+v", d.Id(), err)
		}

		return diag.Errorf("error while reading security group: %+v", err)
	}

	if err := d.Set("display_name", securityGroup.DisplayName); err != nil {
		return diag.Errorf("error setting displayName: %+v", err)
	}

	if err := d.Set("members", flattenSecurityGroupMembers(securityGroup.Members)); err != nil {
		return diag.Errorf("error setting members: %+v", err)
	}

	return nil
}
