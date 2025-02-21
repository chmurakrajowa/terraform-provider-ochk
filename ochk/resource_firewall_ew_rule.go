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

		Importer: &schema.ResourceImporter{
			StateContext: firewallEWRuleStateContextImport,
		},

		Schema: map[string]*schema.Schema{
			"vpc_id": {
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

func firewallEWRuleStateContextImport(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
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

func resourceFirewallEWRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallEWRules

	routerId := strfmt.UUID(d.Get("vpc_id").(string))

	firewallEWRule := mapResourceDataToEWRule(d)

	created, err := proxy.Create(ctx, routerId, firewallEWRule)
	if err != nil {
		return diag.Errorf("error while creating firewall EW rule: %+v", err)
	}

	d.SetId(created.RuleID.String())

	return resourceFirewallEWRuleRead(ctx, d, meta)
}

func resourceFirewallEWRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallEWRules

	routerId := strfmt.UUID(d.Get("vpc_id").(string))

	firewallEWRule, err := proxy.Read(ctx, routerId, strfmt.UUID(d.Id()))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("firewall EW rule with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading firewall EW rule: %+v", err)
	}

	if err := d.Set("display_name", firewallEWRule.DisplayName); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("project_id", firewallEWRule.ProjectID); err != nil {
		return diag.Errorf("error setting project_id: %+v", err)
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

	if len(firewallEWRule.CustomServices) > 0 {
		if err := d.Set("custom_services", flattenCustomServicesFromIDs(firewallEWRule.CustomServices)); err != nil {
			return diag.Errorf("error setting custom services: %+v", err)
		}
	}

	if err := d.Set("source", flattenSecurityGroupFromIDs(firewallEWRule.Source)); err != nil {
		return diag.Errorf("error setting source: %+v", err)
	}

	if err := d.Set("destination", flattenSecurityGroupFromIDs(firewallEWRule.Destination)); err != nil {
		return diag.Errorf("error setting destination: %+v", err)
	}

	if err := d.Set("priority", int(firewallEWRule.Priority)); err != nil {
		return diag.Errorf("error setting priority: %+v", err)
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
	if !d.HasChanges("display_name", "action", "direction", "disabled", "ip_protocol", "services", "custom_services", "source", "destination", "priority") {
		return nil
	}

	proxy := meta.(*sdk.Client).FirewallEWRules

	routerID := strfmt.UUID(d.Get("vpc_id").(string))

	firewallEWRule := mapResourceDataToEWRule(d)
	firewallEWRule.RuleID = strfmt.UUID(d.Id())

	_, err := proxy.Update(ctx, routerID, firewallEWRule)
	if err != nil {
		return diag.Errorf("error while creating firewall EW rule: %+v", err)
	}

	return resourceFirewallEWRuleRead(ctx, d, meta)
}

func castStringToActionEnum(e string) models.Action {
	switch e {
	case "ALLOW":
		return models.ActionALLOW
	case "REJECT":
		return models.ActionREJECT
	case "DROP":
		return models.ActionDROP
	default:
		return ""
	}
}

func castStringToADirectionEnum(e string) models.Direction {
	switch e {
	case "IN_OUT":
		return models.DirectionINOUT
	case "IN":
		return models.DirectionIN
	case "OUT":
		return models.DirectionOUT
	default:
		return ""
	}
}

func castStringToAIPProtocolEnum(e string) models.IPProtocol {
	switch e {
	case "IPV4_IPV6":
		return models.IPProtocolIPV4IPV6
	case "IPV4":
		return models.IPProtocolIPV4
	case "IPV6":
		return models.IPProtocolIPV6
	default:
		return ""
	}
}

func mapResourceDataToEWRule(d *schema.ResourceData) *models.DfwRule {
	rule := &models.DfwRule{
		DisplayName: d.Get("display_name").(string),
		ProjectID:   strfmt.UUID(d.Get("project_id").(string)),
		Action:      castStringToActionEnum(d.Get("action").(string)),
		Direction:   castStringToADirectionEnum(d.Get("direction").(string)),
		Priority:    int64(d.Get("priority").(int)),
	}

	if disabled, ok := d.GetOk("disabled"); ok && disabled.(bool) {
		rule.Disabled = true
	}

	if ipProtocol, ok := d.GetOk("ip_protocol"); ok {
		rule.IPProtocol = castStringToAIPProtocolEnum(ipProtocol.(string))
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

	return rule
}

func resourceFirewallEWRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallEWRules

	routerID := strfmt.UUID(d.Get("vpc_id").(string))

	err := proxy.Delete(ctx, routerID, strfmt.UUID(d.Id()))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("firewall EW rule with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting firewall EW rule: %+v", err)
	}

	return nil
}
