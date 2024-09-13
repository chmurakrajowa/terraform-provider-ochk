package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePortsForwarding() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePortsForwardingRead,
		Schema: map[string]*schema.Schema{
			"floating_ip_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ports_forwarding": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_forwarding_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"floating_ip_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"internal_port_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
func dataSourcePortsForwardingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).PortForwarding
	floatingIpId := strfmt.UUID(d.Get("floating_ip_id").(string))
	portsForwarding, err := proxy.List(ctx, floatingIpId)

	if err := d.Set("ports_forwarding", flattenPortsForwardingLists(portsForwarding)); err != nil {
		return diag.Errorf("error setting ports forwarding list: %v", err)
	}

	if err != nil {
		return diag.Errorf("error while listing ports forwarding")
	}

	d.SetId("ports-forwarding-list")
	return nil
}
