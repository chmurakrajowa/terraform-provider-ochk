package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVirtualNetwork() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVirtualNetworkRead,
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ipam_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"primary_dns_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secondary_dns_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_suffix": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_search_suffix": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"primary_wins_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secondary_wins_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_mask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_gateway_address_cidr": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_network_cidr": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"router": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subtenants": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceVirtualNetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).VirtualNetworks

	name := d.Get("display_name").(string)

	virtualNetworks, err := proxy.ListByDisplayName(ctx, name)
	if err != nil {
		return diag.Errorf("error while listing virtual networks: %+v", err)
	}

	if len(virtualNetworks) < 1 {
		return diag.Errorf("no virtual network found for name: %s", name)
	}

	if len(virtualNetworks) > 1 {
		return diag.Errorf("more than one virtual network with name: %s found!", name)
	}

	d.SetId(virtualNetworks[0].VirtualNetworkID)

	if err := mapVirtualNetworkToResourceData(d, virtualNetworks[0]); err != nil {
		return err
	}

	return nil
}
