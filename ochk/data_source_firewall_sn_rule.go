package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFirewallSNRule() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFirewallSNRuleRead,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(FirewallSNRuleRetryTimeout),
			Update: schema.DefaultTimeout(FirewallSNRuleRetryTimeout),
			Delete: schema.DefaultTimeout(FirewallSNRuleRetryTimeout),
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"router_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"direction": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ip_protocol": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"services": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"custom_services": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"source": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"destination": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"priority": {
				Type:     schema.TypeInt,
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

func dataSourceFirewallSNRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallSNRules

	routerID := d.Get("router_id").(string)

	displayName := d.Get("display_name").(string)

	firewallSNRules, err := proxy.ListByDisplayName(ctx, routerID, displayName)
	if err != nil {
		return diag.Errorf("firewall sn rule for display_name %s not found: %+v", displayName, err)
	}

	if len(firewallSNRules) < 1 {
		return diag.Errorf("no firewall sn rule found for display_name: %s", displayName)
	}

	if len(firewallSNRules) > 1 {
		return diag.Errorf("more than one firewall sn rule with display_name: %s found!", displayName)
	}

	d.SetId(firewallSNRules[0].RuleID)

	if err := d.Set("router_id", flattenRouterInstancesFromScope(firewallSNRules[0].Scope)); err != nil {
		return diag.Errorf("error setting router_id: %+v", err)
	}

	if err := d.Set("display_name", firewallSNRules[0].DisplayName); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("action", firewallSNRules[0].Action); err != nil {
		return diag.Errorf("error setting action: %+v", err)
	}

	if err := d.Set("direction", firewallSNRules[0].Direction); err != nil {
		return diag.Errorf("error setting direction: %+v", err)
	}

	if err := d.Set("disabled", firewallSNRules[0].Disabled); err != nil {
		return diag.Errorf("error setting disabled: %+v", err)
	}

	if err := d.Set("ip_protocol", firewallSNRules[0].IPProtocol); err != nil {
		return diag.Errorf("error setting ip_protocol: %+v", err)
	}

	if err := d.Set("services", flattenServicesFromIDs(firewallSNRules[0].DefaultServices)); err != nil {
		return diag.Errorf("error setting services: %+v", err)
	}

	if err := d.Set("custom_services", flattenCustomServicesFromIDs(firewallSNRules[0].CustomServices)); err != nil {
		return diag.Errorf("error setting custom services: %+v", err)
	}

	if err := d.Set("source", flattenSecurityGroupFromIDs(firewallSNRules[0].Source)); err != nil {
		return diag.Errorf("error setting source: %+v", err)
	}

	if err := d.Set("destination", flattenSecurityGroupFromIDs(firewallSNRules[0].Destination)); err != nil {
		return diag.Errorf("error setting destination: %+v", err)
	}

	if err := d.Set("priority", firewallSNRules[0].Priority); err != nil {
		return diag.Errorf("error setting priority: %+v", err)
	}

	if err := d.Set("created_by", firewallSNRules[0].CreatedBy); err != nil {
		return diag.Errorf("error setting created_by: %+v", err)
	}

	if err := d.Set("created_at", firewallSNRules[0].CreationDate.String()); err != nil {
		return diag.Errorf("error setting created_at: %+v", err)
	}

	if err := d.Set("modified_by", firewallSNRules[0].ModifiedBy); err != nil {
		return diag.Errorf("error setting modified_by: %+v", err)
	}

	if err := d.Set("modified_at", firewallSNRules[0].ModificationDate.String()); err != nil {
		return diag.Errorf("error setting modified_at: %+v", err)
	}

	return nil
}
