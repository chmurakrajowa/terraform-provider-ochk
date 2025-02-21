package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFirewallEWRules() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFirewallEWRulesRead,

		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"firewall_ew_rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"firewall_ew_rule_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_id": {
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

func dataSourceFirewallEWRulesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallEWRules

	routerID := strfmt.UUID(d.Get("vpc_id").(string))

	firewallEWRules, err := proxy.List(ctx, routerID)

	if err := d.Set("firewall_ew_rules", flattenFirewallEWRulesLists(firewallEWRules)); err != nil {
		return diag.Errorf("error setting firewall_ew_rules list: %v", err)
	}

	if err != nil {
		return diag.Errorf("no firewall_ew_rules found for router_id: %v", routerID)
	}

	d.SetId("firewall-ew-rules-" + routerID.String())

	return nil
}
