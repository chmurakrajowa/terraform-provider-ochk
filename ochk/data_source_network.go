package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetwork() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNetworkRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceNetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Networks

	name := d.Get("name").(string)

	networks, err := proxy.ListByName(ctx, name)
	if err != nil {
		return diag.Errorf("error while listing networks: %+v", err)
	}

	if len(networks) < 1 {
		return diag.Errorf("no network found for name: %s", name)
	}

	if len(networks) > 1 {
		return diag.Errorf("more than one network with name: %s found!", name)
	}

	d.SetId(networks[0].NetworkID)

	if err := d.Set("network", networks[0].Network); err != nil {
		return diag.Errorf("error setting network: %+v", err)
	}

	if err := d.Set("network_type", networks[0].NetworkType); err != nil {
		return diag.Errorf("error setting network_type: %+v", err)
	}

	return nil
}
