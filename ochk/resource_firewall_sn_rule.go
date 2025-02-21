package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"strings"
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

		Importer: &schema.ResourceImporter{
			StateContext: firewallSNRuleStateContextImport,
		},

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
			"custom_services": {
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
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
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

func firewallSNRuleStateContextImport(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	parts := strings.SplitN(d.Id(), "/", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return nil, fmt.Errorf("unexpected format of ID (%s), expected format: vpc_id/rule_id", d.Id())
	}
	d.SetId(strings.ToLower(parts[1]))
	if err := d.Set("vpc_id", strings.ToLower(parts[0])); err != nil {
		return nil, fmt.Errorf("cannot set vpc_id: (%s)", parts[0])
	}
	return []*schema.ResourceData{d}, nil
}

func resourceFirewallSNRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallSNRules

	routerID := strfmt.UUID(d.Get("vpc_id").(string))

	rule := mapResourceDataToGFWRule(d)

	created, err := proxy.Create(ctx, routerID, rule)
	if err != nil {
		return diag.Errorf("error while creating firewall SN rule: %+v", err)
	}

	d.SetId(created.RuleID.String())

	return resourceFirewallSNRuleRead(ctx, d, meta)
}

func resourceFirewallSNRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallSNRules

	routerID := strfmt.UUID(d.Get("vpc_id").(string))

	firewallSNRule, err := proxy.Read(ctx, routerID, strfmt.UUID(d.Id()))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("firewall SN rule with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading firewall SN rule: %+v", err)
	}

	if err := d.Set("vpc_id", flattenRouterInstancesFromScope(firewallSNRule.Scope)); err != nil {
		return diag.Errorf("error setting vpc_id: %+v", err)
	}

	if err := d.Set("display_name", firewallSNRule.DisplayName); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("project_id", firewallSNRule.ProjectID); err != nil {
		return diag.Errorf("error setting project_id: %+v", err)
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

	if len(firewallSNRule.CustomServices) > 0 {
		if err := d.Set("custom_services", flattenCustomServicesFromIDs(firewallSNRule.CustomServices)); err != nil {
			return diag.Errorf("error setting custom services: %+v", err)
		}
	}

	if err := d.Set("source", flattenSecurityGroupFromIDs(firewallSNRule.Source)); err != nil {
		return diag.Errorf("error setting source: %+v", err)
	}

	if err := d.Set("destination", flattenSecurityGroupFromIDs(firewallSNRule.Destination)); err != nil {
		return diag.Errorf("error setting destination: %+v", err)
	}

	if err := d.Set("priority", firewallSNRule.Priority); err != nil {
		return diag.Errorf("error setting priority: %+v", err)
	}

	if err := d.Set("created_by", firewallSNRule.CreatedBy); err != nil {
		return diag.Errorf("error setting created_by: %+v", err)
	}

	if err := d.Set("created_at", firewallSNRule.CreationDate.String()); err != nil {
		return diag.Errorf("error setting created_at: %+v", err)
	}

	if err := d.Set("modified_by", firewallSNRule.ModifiedBy); err != nil {
		return diag.Errorf("error setting modified_by: %+v", err)
	}

	if err := d.Set("modified_at", firewallSNRule.ModificationDate.String()); err != nil {
		return diag.Errorf("error setting modified_at: %+v", err)
	}

	return nil
}

func resourceFirewallSNRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	proxy := meta.(*sdk.Client).FirewallSNRules

	if !d.HasChanges("display_name", "action", "direction", "disabled", "ip_protocol", "services", "custom_services", "source", "destination", "priority") {
		return nil
	}

	routerID := strfmt.UUID(d.Get("vpc_id").(string))

	rule := mapResourceDataToGFWRule(d)
	rule.RuleID = strfmt.UUID(d.Id())

	_, err := proxy.Update(ctx, routerID, rule)
	if err != nil {
		return diag.Errorf("error while modifying firewall SN rule: %+v", err)
	}

	return resourceFirewallSNRuleRead(ctx, d, meta)
}

func mapResourceDataToGFWRule(d *schema.ResourceData) *models.GfwRule {
	rule := &models.GfwRule{
		DisplayName: d.Get("display_name").(string),
		ProjectID:   strfmt.UUID(d.Get("project_id").(string)),
		Action:      castStringToActionEnum(d.Get("action").(string)),
		Direction:   castStringToADirectionEnum(d.Get("direction").(string)),
		Disabled:    d.Get("disabled").(bool),
		IPProtocol:  castStringToAIPProtocolEnum(d.Get("ip_protocol").(string)),
		Priority:    int64(d.Get("priority").(int)),
	}

	if services, ok := d.GetOk("services"); ok {
		rule.DefaultServices = expandServicesFromIDs(services.(*schema.Set).List())
	}

	if customServices, ok := d.GetOk("custom_services"); ok {
		rule.CustomServices = expandCustomServicesFromIDs(customServices.(*schema.Set).List())
	}

	if source, ok := d.GetOk("source"); ok {
		rule.Source = expandSecurityGroupFromIDs(source.(*schema.Set).List())
	}

	if destination, ok := d.GetOk("destination"); ok {
		rule.Destination = expandSecurityGroupFromIDs(destination.(*schema.Set).List())
	}

	if routerId, ok := d.GetOk("vpc_id"); ok {

		var RouterInstanceList = make([]*models.RouterInstance, 1)
		routerInstanceId := &models.RouterInstance{
			RouterID: strfmt.UUID(routerId.(string)),
		}
		RouterInstanceList[0] = routerInstanceId
		rule.Scope = RouterInstanceList
	}

	//if position, ok := d.GetOk("position"); ok {
	//	rule.Position = expandFirewallRulePosition(position.([]interface{}))
	//}

	return rule
}

func resourceFirewallSNRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallSNRules

	routerID := strfmt.UUID(d.Get("vpc_id").(string))

	err := proxy.Delete(ctx, routerID, strfmt.UUID(d.Id()))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("firewall SN rule with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting firewall SN rule: %+v", err)
	}

	return nil
}
