package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFirewallSNRules() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFirewallSNRulesRead,

		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"firewall_sn_rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"firewall_sn_rule_id": {
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

func dataSourceFirewallSNRulesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallSNRules

	routerID := d.Get("vpc_id").(string)

	firewallSNRules, err := proxy.List(ctx, routerID)

	if err := d.Set("firewall_sn_rules", flattenFirewallSNRulesLists(firewallSNRules)); err != nil {
		return diag.Errorf("error setting firewall_sn_rules list: %v", err)
	}

	if err != nil {
		return diag.Errorf("no firewall_sn_rules found for router_id: %v", routerID)
	}

	d.SetId("firewall-sn-rules-" + routerID)

	return nil
}
