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
			"security_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rule_id": {
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
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_at": {
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
	name := d.Get("display_name").(string)
	firewallRule, err := proxy.ListByName(ctx, projectId, securityGroupId, name)

	if err != nil {
		return diag.Errorf("firewall rule for ruleId %s not found: %+v", firewallRule, err)
	}

	if len(firewallRule) < 1 {
		return diag.Errorf("no firewall ew rule found for display_name: %s", name)
	}

	if len(firewallRule) > 1 {
		return diag.Errorf("more than one firewall ew rule with display_name: %s found!", name)
	}

	if err := d.Set("display_name", firewallRule[0].Name); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("ether_type", firewallRule[0].EtherType); err != nil {
		return diag.Errorf("error setting ether_type: %+v", err)
	}

	if err := d.Set("description", firewallRule[0].Description); err != nil {
		return diag.Errorf("error setting description: %+v", err)
	}

	if err := d.Set("direction", firewallRule[0].Direction); err != nil {
		return diag.Errorf("error setting direction: %+v", err)
	}

	if err := d.Set("ether_type", firewallRule[0].EtherType); err != nil {
		return diag.Errorf("error setting ether_type: %+v", err)
	}

	// jezeli jest port domyslny czyli wsztskie porty są dopuszczone to API nie zwraca Port_Range_Min , stąd int32 mapuje null'a na 0
	if err := d.Set("port_range_min", firewallRule[0].PortRangeMin); err != nil {
		return diag.Errorf("error setting port_range_min: %+v", err)
	}
	// jezeli jest port domyslny czyli wsztskie porty są dopuszczone to API nie zwraca Port_Range_Max, stąd int32 mapuje null'a na 0

	if err := d.Set("port_range_max", firewallRule[0].PortRangeMax); err != nil {
		return diag.Errorf("error setting port_range_max: %+v", err)
	}

	if err := d.Set("protocol", firewallRule[0].Protocol); err != nil {
		return diag.Errorf("error setting protocol: %+v", err)
	}

	if err := d.Set("remote_ip_prefix", firewallRule[0].RemoteIPPrefix); err != nil {
		return diag.Errorf("error setting remote_ip_prefix: %+v", err)
	}

	if err := d.Set("remote_ip_prefix", firewallRule[0].RemoteIPPrefix); err != nil {
		return diag.Errorf("error setting remote_ip_prefix: %+v", err)
	}

	if err := d.Set("rule_id", firewallRule[0].RuleID); err != nil {
		return diag.Errorf("error setting rule_id: %+v", err)
	}

	if err := d.Set("created_by", firewallRule[0].CreatedBy.DisplayName); err != nil {
		return diag.Errorf("error setting created_by: %+v", err)
	}

	if err := d.Set("created_at", firewallRule[0].CreationDate.String()); err != nil {
		return diag.Errorf("error setting created_at: %+v", err)
	}

	if err := d.Set("modified_by", firewallRule[0].ModifiedBy.DisplayName); err != nil {
		return diag.Errorf("error setting modified_by: %+v", err)
	}

	if err := d.Set("modified_at", firewallRule[0].ModificationDate.String()); err != nil {
		return diag.Errorf("error setting modified_at: %+v", err)
	}

	d.SetId(firewallRule[0].RuleID.String())
	return nil
}
