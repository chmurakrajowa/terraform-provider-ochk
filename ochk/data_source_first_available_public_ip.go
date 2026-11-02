package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFirstAvailablePublicIp() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFirstAvailablePublicIPRead,
		Schema: map[string]*schema.Schema{
			"ip_address_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func dataSourceFirstAvailablePublicIPRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).AvailablePublicIp
	floatingIPVms, err := proxy.Get(ctx)

	if err != nil {
		return diag.Errorf("error while get first available public ip address: %+v", err)
	}

	if err := d.Set("ip_address_id", floatingIPVms.GetIpAddressId()); err != nil {
		return diag.Errorf("error setting deployment_type: %+v", err)
	}

	if err := d.Set("ip_address", floatingIPVms.GetIpAddress()); err != nil {
		return diag.Errorf("error setting deployment_category: %+v", err)
	}

	d.SetId(floatingIPVms.GetIpAddressId())
	return nil
}
