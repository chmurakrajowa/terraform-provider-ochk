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
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			//TODO czy jest wymagane?
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			//TODO czy jest wymagane?
			"direction": {
				Type:     schema.TypeString,
				Required: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			//TODO czy jest wymagane?
			"ip_protocol": {
				Type:     schema.TypeString,
				Optional: true,
			},
			//TODO czy kolejność ma znaczenie?
			"services": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			//TODO czy kolejność ma znaczenie?
			"source": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			//TODO czy kolejność ma znaczenie?
			"destination": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			//TODO czy jest wymagane?
			"position": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_id": {
							Type:     schema.TypeString,
							Required: true,
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

func resourceFirewallEWRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallEWFRules

	d.Get("ipProtocol").(*schema.Set).List()

	FirewallEWRule := &models.DFWRule{
		DisplayName: d.Get("display_name").(string),
		Action:      d.Get("action").(string),
		Direction:   d.Get("directions").(string),
	}

	if disabled, ok := d.GetOk("disabled"); ok && disabled.(bool) {
		FirewallEWRule.Disabled = true
	}

	if ipProtocol, ok := d.GetOk("ip_protocol"); ok {
		FirewallEWRule.IPProtocol = ipProtocol.(string)
	}

	if services, ok := d.GetOk("ipProtocol"); ok {
		FirewallEWRule.DefaultServices = expandServicesFromIDs(services.(*schema.Set).List())
	}

	if source, ok := d.GetOk("source"); ok {
		FirewallEWRule.Source = expandSecurityGroupFromIDs(source.(*schema.Set).List())
	}

	if destination, ok := d.GetOk("source"); ok {
		FirewallEWRule.Destination = expandSecurityGroupFromIDs(destination.(*schema.Set).List())
	}

	if destination, ok := d.GetOk("position"); ok {
		FirewallEWRule.Position = expandSecurityGroupFromIDs(destination.(*schema.Set).List())
	}

	created, err := proxy.Create(ctx, FirewallEWRule)
	if err != nil {
		return diag.Errorf("error while creating security PolicyRule: %+v", err)
	}

	d.SetId(created.ID)

	return resourceFirewallEWRuleRead(ctx, d, meta)
}

func resourceFirewallEWRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallEWRule

	FirewallEWRule, err := proxy.Read(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			d.SetId("")
			return diag.Errorf("security PolicyRule with id %s not found: %+v", d.Id(), err)
		}

		return diag.Errorf("error while reading security PolicyRule: %+v", err)
	}

	if err := d.Set("display_name", FirewallEWRule.DisplayName); err != nil {
		return diag.Errorf("error setting displayName: %+v", err)
	}

	if err := d.Set("members", flattenFirewallEWRuleMembers(FirewallEWRule.Members)); err != nil {
		return diag.Errorf("error setting members: %+v", err)
	}

	return nil
}

func resourceFirewallEWRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.FromErr(fmt.Errorf("updating FirewallEWRule is not implemented"))
}

func resourceFirewallEWRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallEWRules

	err := proxy.Delete(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			d.SetId("")
			return diag.Errorf("security PolicyRule with id %s not found: %+v", d.Id(), err)
		}

		return diag.Errorf("error while reading deleting PolicyRule: %+v", err)
	}

	return nil
}
