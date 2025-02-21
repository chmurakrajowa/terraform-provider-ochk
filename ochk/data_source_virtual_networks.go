package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVirtualNetworks() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVirtualNetworksRead,
		Schema: map[string]*schema.Schema{
			"virtual_networks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"virtual_network_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"folder_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipam_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"vpc_id": {
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

func dataSourceVirtualNetworksRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).VirtualNetworks

	virtualNetworks, err := proxy.List(ctx)
	if err != nil {
		return diag.Errorf("error while listing virtual networks: %+v", err)
	}

	if err := d.Set("virtual_networks", flattenVirtualNetworks(virtualNetworks)); err != nil {
		return diag.Errorf("error setting virtual networks list: %v", err)
	}

	d.SetId("virtual-networks-list")

	return nil
}
