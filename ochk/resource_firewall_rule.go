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
	FirewallRuleRetryTimeout = 1 * time.Minute
	E3001                    = "TF_ERROR{3001}: Error while creating firewall rule: %+v"
	E3001_UPDATE             = "TF_ERROR{3001}: Error while updating firewall rule: %+v"
	E3002                    = "Field %s or %s is required. Please put %s or %s field to input config file."
	E3003                    = "Error while mapping firewall rule to resource: %+v"
	E3004                    = "Error while reading firewall rule: %+v"
	E3005                    = "Error while deleting firewall rule: %+v"
)

func resourceFirewallRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFirewallRuleCreate,
		ReadContext:   resourceFirewallRuleRead,
		UpdateContext: resourceFirewallRuleUpdate,
		DeleteContext: resourceFirewallRuleDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(FirewallRuleRetryTimeout),
			Update: schema.DefaultTimeout(FirewallRuleRetryTimeout),
			Delete: schema.DefaultTimeout(FirewallRuleRetryTimeout),
		},

		Importer: &schema.ResourceImporter{
			StateContext: firewallRuleStateContextImport,
		},

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
				Optional: true,
			},
			"ether_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"direction": {
				Type:     schema.TypeString,
				Required: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port_range_min": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"port_range_max": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"remote_ip_prefix": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dest_security_group": {
				Type:     schema.TypeString,
				Optional: true,
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

func firewallRuleStateContextImport(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
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

func resourceFirewallRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallRules

	projectId := strfmt.UUID(d.Get("project_id").(string))
	securityGroupId := strfmt.UUID(d.Get("security_group_id").(string))

	firewallRule, rule_sg := mapResourceDataToRule(d)

	if rule_sg != nil {
		return diag.Errorf(E3003, rule_sg)
	}

	created, err := proxy.Create(ctx, projectId, securityGroupId, firewallRule)
	if err != nil {
		return diag.Errorf("error while creating firewall rule: %+v", err)
	}

	d.SetId(created.RuleID.String())

	return resourceFirewallRuleRead(ctx, d, meta)
}

func resourceFirewallRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallRules
	projectId := strfmt.UUID(d.Get("project_id").(string))
	securityGroupId := strfmt.UUID(d.Get("security_group_id").(string))
	ruleId := strfmt.UUID(d.Id())

	firewallRule, err := proxy.Read(ctx, projectId, securityGroupId, ruleId)
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("firewall rule with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading firewall rule: %+v", err)
	}

	if err := d.Set("display_name", firewallRule.Name); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("project_id", projectId); err != nil {
		return diag.Errorf("error setting project_id: %+v", err)
	}

	if err := d.Set("security_group_id", securityGroupId); err != nil {
		return diag.Errorf("error setting security_group_id: %+v", err)
	}

	if err := d.Set("rule_id", firewallRule.RuleID); err != nil {
		return diag.Errorf("error setting rule_id: %+v", err)
	}

	if err := d.Set("direction", firewallRule.Direction); err != nil {
		return diag.Errorf("error setting direction: %+v", err)
	}

	if err := d.Set("description", firewallRule.Description); err != nil {
		return diag.Errorf("error setting description: %+v", err)
	}

	if err := d.Set("ether_type", firewallRule.EtherType); err != nil {
		return diag.Errorf("error setting ether_type: %+v", err)
	}

	if err := d.Set("protocol", firewallRule.Protocol); err != nil {
		return diag.Errorf("error setting protocol: %+v", err)
	}

	if err := d.Set("port_range_min", firewallRule.PortRangeMin); err != nil {
		return diag.Errorf("error setting port_range_min: %+v", err)
	}

	if err := d.Set("port_range_max", firewallRule.PortRangeMax); err != nil {
		return diag.Errorf("error setting port_range_max: %+v", err)
	}

	if err := d.Set("remote_ip_prefix", firewallRule.RemoteIPPrefix); err != nil {
		return diag.Errorf("error setting remote_ip_prefix: %+v", err)
	}

	if err := d.Set("created_by", firewallRule.CreatedBy.DisplayName); err != nil {
		return diag.Errorf("error setting created_by: %+v", err)
	}

	if err := d.Set("created_at", firewallRule.CreationDate.String()); err != nil {
		return diag.Errorf("error setting created_at: %+v", err)
	}

	if err := d.Set("modified_by", firewallRule.ModifiedBy.DisplayName); err != nil {
		return diag.Errorf("error setting modified_by: %+v", err)
	}

	if err := d.Set("modified_at", firewallRule.ModificationDate.String()); err != nil {
		return diag.Errorf("error setting modified_at: %+v", err)
	}

	return nil
}

func resourceFirewallRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if !d.HasChanges("display_name", "description", "ether_type", "direction", "protocol", "port_range_min", "port_range_max", "remote_ip_prefix") {
		return nil
	}
	proxy := meta.(*sdk.Client).FirewallRules

	securityGroupID := strfmt.UUID(d.Get("security_group_id").(string))
	projectID := strfmt.UUID(d.Get("project_id").(string))
	firewallRule, err_rule := mapResourceDataToRule(d)

	if err_rule != nil {
		return diag.Errorf(E3001_UPDATE, err_rule)
	}
	firewallRule.RuleID = strfmt.UUID(d.Id())

	_, err := proxy.Update(ctx, projectID, securityGroupID, firewallRule)
	if err != nil {
		return diag.Errorf("error while creating firewall rule: %+v", err)
	}

	return resourceFirewallRuleRead(ctx, d, meta)
}

func mapResourceDataToRule(d *schema.ResourceData) (*models.FirewallRule, diag.Diagnostics) {

	if d.Get("dest_security_group").(string) == "" && d.Get("remote_ip_prefix").(string) == "" {
		fmt.Printf("dest_security_group or remote_ip_prefix has to be set in input resource.")
		return nil, diag.Errorf(E3001, fmt.Sprintf(E3002, "dest_security_group", "remote_ip_prefix", "dest_security_group", "remote_ip_prefix"))
	}
	if d.Get("dest_security_group").(string) != "" {
		rule := &models.FirewallRule{
			Name:              d.Get("display_name").(string),
			Description:       d.Get("description").(string),
			ProjectExternalID: strfmt.UUID(d.Get("project_id").(string)),
			EtherType:         models.EtherType(d.Get("ether_type").(string)),
			Direction:         models.Direction1(d.Get("direction").(string)),
			Protocol:          models.Protocol(d.Get("protocol").(string)),
			PortRangeMax:      int64(d.Get("port_range_max").(int)),
			PortRangeMin:      int64(d.Get("port_range_min").(int)),
			RemoteIPPrefix:    d.Get("remote_ip_prefix").(string),
			SecurityGroup:     expandSecurityGroup(d.Get("dest_security_group").(string)),
		}
		return rule, nil
	} else {
		rule := &models.FirewallRule{
			Name:              d.Get("display_name").(string),
			Description:       d.Get("description").(string),
			ProjectExternalID: strfmt.UUID(d.Get("project_id").(string)),
			EtherType:         models.EtherType(d.Get("ether_type").(string)),
			Direction:         models.Direction1(d.Get("direction").(string)),
			Protocol:          models.Protocol(d.Get("protocol").(string)),
			PortRangeMax:      int64(d.Get("port_range_max").(int)),
			PortRangeMin:      int64(d.Get("port_range_min").(int)),
			RemoteIPPrefix:    d.Get("remote_ip_prefix").(string),
		}
		return rule, nil
	}
}

func resourceFirewallRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallRules

	securityGroupID := strfmt.UUID(d.Get("security_group_id").(string))
	projectID := strfmt.UUID(d.Get("project_id").(string))

	err := proxy.Delete(ctx, projectID, securityGroupID, strfmt.UUID(d.Id()))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("firewall rule with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting firewall rule: %+v", err)
	}

	return nil
}

func expandSecurityGroup(dest_security_group string) *models.SecurityGroup {
	sg_dest := &models.SecurityGroup{
		ID: strfmt.UUID(dest_security_group),
	}
	return sg_dest
}
