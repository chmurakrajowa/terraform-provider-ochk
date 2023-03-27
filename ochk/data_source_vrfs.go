package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVrfs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVrfsRead,

		Schema: map[string]*schema.Schema{
			"vrfs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vrf_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceVrfsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Routers

	vrfs, err := proxy.List(ctx)
	if err != nil {
		return diag.Errorf("error while listing vrfs: %+v", err)
	}

	if err := d.Set("vrfs", flattenVrfs(vrfs)); err != nil {
		return diag.Errorf("error settingv rfs list: %v", err)
	}

	d.SetId("vrfs-list")

	return nil
}
