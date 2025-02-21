package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFirewallEWRule() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFirewallEWRuleRead,

		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Computed: true,
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
				Optional: true,
				Default:  false,
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

func dataSourceFirewallEWRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallEWRules

	routerID := strfmt.UUID(d.Get("vpc_id").(string))

	displayName := d.Get("display_name").(string)

	firewallEWRules, err := proxy.ListByDisplayName(ctx, routerID, displayName)

	if err != nil {
		return diag.Errorf("firewall ew rule for display_name %s not found: %+v", displayName, err)
	}

	if len(firewallEWRules) < 1 {
		return diag.Errorf("no firewall ew rule found for display_name: %s", displayName)
	}

	if len(firewallEWRules) > 1 {
		return diag.Errorf("more than one firewall ew rule with display_name: %s found!", displayName)
	}

	d.SetId(firewallEWRules[0].RuleID.String())

	if err := d.Set("display_name", firewallEWRules[0].DisplayName); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("project_id", firewallEWRules[0].ProjectID); err != nil {
		return diag.Errorf("error setting project_id: %+v", err)
	}

	if err := d.Set("action", firewallEWRules[0].Action); err != nil {
		return diag.Errorf("error setting action: %+v", err)
	}

	if err := d.Set("direction", firewallEWRules[0].Direction); err != nil {
		return diag.Errorf("error setting direction: %+v", err)
	}

	if err := d.Set("disabled", firewallEWRules[0].Disabled); err != nil {
		return diag.Errorf("error setting disabled: %+v", err)
	}

	if err := d.Set("ip_protocol", firewallEWRules[0].IPProtocol); err != nil {
		return diag.Errorf("error setting ip_protocol: %+v", err)
	}

	if err := d.Set("services", flattenServicesFromIDs(firewallEWRules[0].DefaultServices)); err != nil {
		return diag.Errorf("error setting services: %+v", err)
	}

	if err := d.Set("custom_services", flattenCustomServicesFromIDs(firewallEWRules[0].CustomServices)); err != nil {
		return diag.Errorf("error setting custom services: %+v", err)
	}

	if err := d.Set("source", flattenSecurityGroupFromIDs(firewallEWRules[0].Source)); err != nil {
		return diag.Errorf("error setting source: %+v", err)
	}

	if err := d.Set("destination", flattenSecurityGroupFromIDs(firewallEWRules[0].Destination)); err != nil {
		return diag.Errorf("error setting destination: %+v", err)
	}

	if err := d.Set("priority", int(firewallEWRules[0].Priority)); err != nil {
		return diag.Errorf("error setting priority: %+v", err)
	}

	if err := d.Set("created_by", firewallEWRules[0].CreatedBy); err != nil {
		return diag.Errorf("error setting created_by: %+v", err)
	}

	if err := d.Set("created_at", firewallEWRules[0].CreationDate.String()); err != nil {
		return diag.Errorf("error setting created_at: %+v", err)
	}

	if err := d.Set("modified_by", firewallEWRules[0].ModifiedBy); err != nil {
		return diag.Errorf("error setting modified_by: %+v", err)
	}

	if err := d.Set("modified_at", firewallEWRules[0].ModificationDate.String()); err != nil {
		return diag.Errorf("error setting modified_at: %+v", err)
	}

	return nil
}
