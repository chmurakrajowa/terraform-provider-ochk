package ochk

import (
	"context"
	"fmt"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	FirewallSNRuleRetryTimeout = 1 * time.Minute
)

func resourceFirewallSNRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFirewallSNRuleCreate,
		ReadContext:   resourceFirewallSNRuleRead,
		UpdateContext: resourceFirewallSNRuleUpdate,
		DeleteContext: resourceFirewallSNRuleDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(FirewallSNRuleRetryTimeout),
			Update: schema.DefaultTimeout(FirewallSNRuleRetryTimeout),
			Delete: schema.DefaultTimeout(FirewallSNRuleRetryTimeout),
		},

		Schema: map[string]*schema.Schema{
			"gateway_policy_id": {
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
		},
	}
}

func resourceFirewallSNRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallSNRules

	gatewayPolicyID := d.Get("gateway_policy_id").(string)

	firewallSNRule := &models.GFWRule{
		DisplayName: d.Get("display_name").(string),
		Action:      d.Get("action").(string),
		Direction:   d.Get("direction").(string),
	}

	if disabled, ok := d.GetOk("disabled"); ok && disabled.(bool) {
		firewallSNRule.Disabled = true
	}

	if ipProtocol, ok := d.GetOk("ip_protocol"); ok {
		firewallSNRule.IPProtocol = ipProtocol.(string)
	}

	if services, ok := d.GetOk("services"); ok {
		firewallSNRule.DefaultServices = expandServicesFromIDs(services.(*schema.Set).List())
	}

	if source, ok := d.GetOk("source"); ok {
		firewallSNRule.Source = expandSecurityGroupFromIDs(source.(*schema.Set).List())
	}

	if destination, ok := d.GetOk("destination"); ok {
		firewallSNRule.Destination = expandSecurityGroupFromIDs(destination.(*schema.Set).List())
	}

	if position, ok := d.GetOk("position"); ok {
		firewallSNRule.Position = expandFirewallRulePosition(position.([]interface{}))
	}

	created, err := proxy.Create(ctx, gatewayPolicyID, firewallSNRule)
	if err != nil {
		return diag.Errorf("error while creating firewall SN rule: %+v", err)
	}

	d.SetId(created.RuleID)

	return resourceFirewallSNRuleRead(ctx, d, meta)
}

func resourceFirewallSNRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallSNRules

	gatewayPolicyID := d.Get("gateway_policy_id").(string)

	firewallSNRule, err := proxy.Read(ctx, gatewayPolicyID, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			d.SetId("")
			return diag.Errorf("firewall SN rule with id %s not found: %+v", d.Id(), err)
		}

		return diag.Errorf("error while reading firewall SN rule: %+v", err)
	}

	if err := d.Set("display_name", firewallSNRule.DisplayName); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("action", firewallSNRule.Action); err != nil {
		return diag.Errorf("error setting action: %+v", err)
	}

	if err := d.Set("direction", firewallSNRule.Direction); err != nil {
		return diag.Errorf("error setting direction: %+v", err)
	}

	if err := d.Set("disabled", firewallSNRule.Disabled); err != nil {
		return diag.Errorf("error setting disabled: %+v", err)
	}

	if err := d.Set("ip_protocol", firewallSNRule.IPProtocol); err != nil {
		return diag.Errorf("error setting ip_protocol: %+v", err)
	}

	if err := d.Set("services", flattenServicesFromIDs(firewallSNRule.DefaultServices)); err != nil {
		return diag.Errorf("error setting services: %+v", err)
	}

	if err := d.Set("source", flattenSecurityGroupFromIDs(firewallSNRule.Source)); err != nil {
		return diag.Errorf("error setting source: %+v", err)
	}

	if err := d.Set("destination", flattenSecurityGroupFromIDs(firewallSNRule.Destination)); err != nil {
		return diag.Errorf("error setting destination: %+v", err)
	}

	if err := d.Set("position", flattenFirewallRulePosition(firewallSNRule.Position)); err != nil {
		return diag.Errorf("error setting position: %+v", err)
	}

	return nil
}

func resourceFirewallSNRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.FromErr(fmt.Errorf("updating firewall SN rule is not implemented"))
}

func resourceFirewallSNRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallSNRules

	gatewayPolicyID := d.Get("gateway_policy_id").(string)

	err := proxy.Delete(ctx, gatewayPolicyID, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			d.SetId("")
			return diag.Errorf("firewall SN rule with id %s not found: %+v", d.Id(), err)
		}

		return diag.Errorf("error while reading deleting firewall SN rule: %+v", err)
	}

	return nil
}
