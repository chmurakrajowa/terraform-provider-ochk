package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	FirewallEWRuleRetryTimeout = 1 * time.Minute
)

func resourceFirewallEWRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFirewallEWRuleCreate,
		ReadContext:   resourceFirewallEWRuleRead,
		UpdateContext: resourceFirewallEWRuleUpdate,
		DeleteContext: resourceFirewallEWRuleDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(FirewallEWRuleRetryTimeout),
			Update: schema.DefaultTimeout(FirewallEWRuleRetryTimeout),
			Delete: schema.DefaultTimeout(FirewallEWRuleRetryTimeout),
		},

		Schema: map[string]*schema.Schema{
			"security_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ALLOW",
			},
			"direction": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "IN_OUT",
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ip_protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "IPV4_IPV6",
			},
			"services": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"source": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"destination": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"position": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"revise_operation": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
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

func resourceFirewallEWRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallEWRules

	securityPolicyID := d.Get("security_policy_id").(string)

	firewallEWRule := &models.DFWRule{
		DisplayName: d.Get("display_name").(string),
		Action:      d.Get("action").(string),
		Direction:   d.Get("direction").(string),
	}

	if disabled, ok := d.GetOk("disabled"); ok && disabled.(bool) {
		firewallEWRule.Disabled = true
	}

	if ipProtocol, ok := d.GetOk("ip_protocol"); ok {
		firewallEWRule.IPProtocol = ipProtocol.(string)
	}

	if services, ok := d.GetOk("services"); ok {
		firewallEWRule.DefaultServices = expandServicesFromIDs(services.(*schema.Set).List())
	}

	if source, ok := d.GetOk("source"); ok {
		firewallEWRule.Source = expandSecurityGroupFromIDs(source.(*schema.Set).List())
	}

	if destination, ok := d.GetOk("destination"); ok {
		firewallEWRule.Destination = expandSecurityGroupFromIDs(destination.(*schema.Set).List())
	}

	if position, ok := d.GetOk("position"); ok {
		firewallEWRule.Position = expandFirewallRulePosition(position.([]interface{}))
	}

	created, err := proxy.Create(ctx, securityPolicyID, firewallEWRule)
	if err != nil {
		return diag.Errorf("error while creating firewall EW rule: %+v", err)
	}

	d.SetId(created.RuleID)

	return resourceFirewallEWRuleRead(ctx, d, meta)
}

func resourceFirewallEWRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallEWRules

	securityPolicyID := d.Get("security_policy_id").(string)

	firewallEWRule, err := proxy.Read(ctx, securityPolicyID, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			d.SetId("")
			return diag.Errorf("firewall EW rule with id %s not found: %+v", d.Id(), err)
		}

		return diag.Errorf("error while reading firewall EW rule: %+v", err)
	}

	if err := d.Set("display_name", firewallEWRule.DisplayName); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("action", firewallEWRule.Action); err != nil {
		return diag.Errorf("error setting action: %+v", err)
	}

	if err := d.Set("direction", firewallEWRule.Direction); err != nil {
		return diag.Errorf("error setting direction: %+v", err)
	}

	if err := d.Set("disabled", firewallEWRule.Disabled); err != nil {
		return diag.Errorf("error setting disabled: %+v", err)
	}

	if err := d.Set("ip_protocol", firewallEWRule.IPProtocol); err != nil {
		return diag.Errorf("error setting ip_protocol: %+v", err)
	}

	if err := d.Set("services", flattenServicesFromIDs(firewallEWRule.DefaultServices)); err != nil {
		return diag.Errorf("error setting services: %+v", err)
	}

	if err := d.Set("source", flattenSecurityGroupFromIDs(firewallEWRule.Source)); err != nil {
		return diag.Errorf("error setting source: %+v", err)
	}

	if err := d.Set("destination", flattenSecurityGroupFromIDs(firewallEWRule.Destination)); err != nil {
		return diag.Errorf("error setting destination: %+v", err)
	}

	if err := d.Set("position", flattenFirewallRulePosition(firewallEWRule.Position)); err != nil {
		return diag.Errorf("error setting position: %+v", err)
	}

	if err := d.Set("created_by", firewallEWRule.CreatedBy); err != nil {
		return diag.Errorf("error setting created_by: %+v", err)
	}

	if err := d.Set("created_at", firewallEWRule.CreationDate.String()); err != nil {
		return diag.Errorf("error setting created_at: %+v", err)
	}

	if err := d.Set("modified_by", firewallEWRule.ModifiedBy); err != nil {
		return diag.Errorf("error setting modified_by: %+v", err)
	}

	if err := d.Set("modified_at", firewallEWRule.ModificationDate.String()); err != nil {
		return diag.Errorf("error setting modified_at: %+v", err)
	}

	return nil
}

func resourceFirewallEWRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if !d.HasChanges("display_name", "action", "direction", "disabled", "ip_protocol", "services", "source", "destination", "position") {
		return nil
	}

	proxy := meta.(*sdk.Client).FirewallEWRules

	securityPolicyID := d.Get("security_policy_id").(string)

	firewallEWRule := &models.DFWRule{
		RuleID:          d.Id(),
		Action:          d.Get("action").(string),
		DefaultServices: expandServicesFromIDs(d.Get("services").(*schema.Set).List()),
		Destination:     expandSecurityGroupFromIDs(d.Get("destination").(*schema.Set).List()),
		Direction:       d.Get("direction").(string),
		Disabled:        d.Get("disabled").(bool),
		DisplayName:     d.Get("display_name").(string),
		IPProtocol:      d.Get("ip_protocol").(string),
		Source:          expandSecurityGroupFromIDs(d.Get("source").(*schema.Set).List()),
	}

	_, err := proxy.Update(ctx, securityPolicyID, firewallEWRule)
	if err != nil {
		return diag.Errorf("error while creating firewall EW rule: %+v", err)
	}

	return resourceFirewallEWRuleRead(ctx, d, meta)
}

func resourceFirewallEWRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallEWRules

	securityPolicyID := d.Get("security_policy_id").(string)

	err := proxy.Delete(ctx, securityPolicyID, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			d.SetId("")
			return diag.Errorf("firewall EW rule with id %s not found: %+v", d.Id(), err)
		}

		return diag.Errorf("error while deleting firewall EW rule: %+v", err)
	}

	return nil
}
