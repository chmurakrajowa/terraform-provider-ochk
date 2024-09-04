package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFirewallRules() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFirewallRulesRead,

		Schema: map[string]*schema.Schema{
			"security_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"firewall_rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ether_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"direction": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"protocol": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port_range_min": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"port_range_max": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"remote_ip_prefix": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceFirewallRulesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallRules

	projectId := strfmt.UUID(d.Get("project_id").(string))
	securityGroupId := strfmt.UUID(d.Get("security_group_id").(string))

	firewallRules, err := proxy.List(ctx, projectId, securityGroupId)

	if err := d.Set("firewall_rules", flattenFirewallRulesLists(firewallRules)); err != nil {
		return diag.Errorf("error setting firewall_rules list: %v", err)
	}

	if err != nil {
		return diag.Errorf("no firewall_rules found for projectId  or securityGroupId: %v%v", projectId, securityGroupId)
	}

	return nil
}
