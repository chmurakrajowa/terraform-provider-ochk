package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFloatingIPVms() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFloatingIPVmsRead,

		Schema: map[string]*schema.Schema{
			"floating_ip_vms": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"virtual_machine_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"virtual_machine_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"osc_port_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mac_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
func dataSourceFloatingIPVmsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FloatingIPVms
	floatingIPVms, err := proxy.List(ctx)

	if err != nil {
		return diag.Errorf("error while get available floating ip vms: %+v", err)
	}

	if err := d.Set("floating_ip_vms", flattenFloatingIPVmsList(floatingIPVms)); err != nil {
		return diag.Errorf("error setting available floating ip vms: %+v", err)
	}

	d.SetId("floating-ip-vms-list")
	return nil
}
