package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFloatingIPAddresses() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFloatingIPAddressesRead,
		Schema: map[string]*schema.Schema{
			"floating_ip_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"floating_ip_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_address": {
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
func dataSourceFloatingIPAddressesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FloatingIPAddresses
	floatingIPAddresses, err := proxy.List(ctx)

	if err != nil {
		return diag.Errorf("error while get available floating ip addresses: %+v", err)
	}

	if err := d.Set("floating_ip_addresses", flattenFloatingIPAddressList(floatingIPAddresses)); err != nil {
		return diag.Errorf("error setting available floating ip address: %+v", err)
	}

	d.SetId("floating-ip-addresses-list")
	return nil
}
