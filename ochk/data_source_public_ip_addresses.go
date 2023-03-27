package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePublicIPAddresses() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePublicIPAddressesRead,
		Schema: map[string]*schema.Schema{
			"public_ip_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"public_ip_address_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address": {
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
func dataSourcePublicIPAddressesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).PublicIPAddresses

	publicIPAddresses, err := proxy.List(ctx)

	if err != nil {
		return diag.Errorf("error while get available public ip addresses: %+v", err)
	}

	if err := d.Set("public_ip_addresses", flattenPublicIPAddress(publicIPAddresses)); err != nil {
		return diag.Errorf("error setting available public ip address: %+v", err)
	}

	d.SetId("public-ip-addresses-list")

	return nil
}
