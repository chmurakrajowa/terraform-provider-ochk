package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFirewallRule() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFirewallRuleRead,

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
	}
}

func dataSourceFirewallRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallRules
	projectId := strfmt.UUID(d.Get("project_id").(string))
	securityGroupId := strfmt.UUID(d.Get("security_group_id").(string))
	ruleId := strfmt.UUID(d.Get("rule_id").(string))

	firewallRule, err := proxy.Read(ctx, ruleId, projectId, securityGroupId)

	if err != nil {
		return diag.Errorf("firewall rule for ruleId %s not found: %+v", ruleId, err)
	}

	//if len(firewallRules) < 1 {
	//	return diag.Errorf("no firewall ew rule found for display_name: %s", displayName)
	//}
	//
	//if len(firewallRules) > 1 {
	//	return diag.Errorf("more than one firewall ew rule with display_name: %s found!", displayName)
	//}

	d.SetId(firewallRule.RuleID.String())

	//if err := d.Set("display_name", firewallEWRules[0].DisplayName); err != nil {
	//	return diag.Errorf("error setting display_name: %+v", err)
	//}
	//
	//if err := d.Set("project_id", firewallEWRules[0].ProjectID); err != nil {
	//	return diag.Errorf("error setting project_id: %+v", err)
	//}
	//
	//if err := d.Set("action", firewallEWRules[0].Action); err != nil {
	//	return diag.Errorf("error setting action: %+v", err)
	//}
	//
	//if err := d.Set("direction", firewallEWRules[0].Direction); err != nil {
	//	return diag.Errorf("error setting direction: %+v", err)
	//}
	//
	//if err := d.Set("disabled", firewallEWRules[0].Disabled); err != nil {
	//	return diag.Errorf("error setting disabled: %+v", err)
	//}
	//
	//if err := d.Set("ip_protocol", firewallEWRules[0].IPProtocol); err != nil {
	//	return diag.Errorf("error setting ip_protocol: %+v", err)
	//}
	//
	//if err := d.Set("services", flattenServicesFromIDs(firewallEWRules[0].DefaultServices)); err != nil {
	//	return diag.Errorf("error setting services: %+v", err)
	//}
	//
	//if err := d.Set("custom_services", flattenCustomServicesFromIDs(firewallEWRules[0].CustomServices)); err != nil {
	//	return diag.Errorf("error setting custom services: %+v", err)
	//}
	//
	//if err := d.Set("source", flattenSecurityGroupFromIDs(firewallEWRules[0].Source)); err != nil {
	//	return diag.Errorf("error setting source: %+v", err)
	//}
	//
	//if err := d.Set("destination", flattenSecurityGroupFromIDs(firewallEWRules[0].Destination)); err != nil {
	//	return diag.Errorf("error setting destination: %+v", err)
	//}
	//
	//if err := d.Set("priority", int(firewallEWRules[0].Priority)); err != nil {
	//	return diag.Errorf("error setting priority: %+v", err)
	//}
	//
	//if err := d.Set("created_by", firewallEWRules[0].CreatedBy); err != nil {
	//	return diag.Errorf("error setting created_by: %+v", err)
	//}
	//
	//if err := d.Set("created_at", firewallEWRules[0].CreationDate.String()); err != nil {
	//	return diag.Errorf("error setting created_at: %+v", err)
	//}
	//
	//if err := d.Set("modified_by", firewallEWRules[0].ModifiedBy); err != nil {
	//	return diag.Errorf("error setting modified_by: %+v", err)
	//}
	//
	//if err := d.Set("modified_at", firewallEWRules[0].ModificationDate.String()); err != nil {
	//	return diag.Errorf("error setting modified_at: %+v", err)
	//}

	return nil
}
