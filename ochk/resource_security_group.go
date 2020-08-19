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
	SecurityGroupRetryTimeout = 1 * time.Minute
)

func resourceSecurityGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSecurityGroupCreate,
		ReadContext:   resourceSecurityGroupRead,
		UpdateContext: resourceSecurityGroupUpdate,
		DeleteContext: resourceSecurityGroupDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(SecurityGroupRetryTimeout),
			Update: schema.DefaultTimeout(SecurityGroupRetryTimeout),
			Delete: schema.DefaultTimeout(SecurityGroupRetryTimeout),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"members": {
				Type:     schema.TypeSet,
				MinItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceSecurityGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).SecurityGroups

	securityGroup := mapResourceDataToSecurityGroup(d)

	created, err := proxy.Create(ctx, securityGroup)
	if err != nil {
		return diag.Errorf("error while creating security group: %+v", err)
	}

	d.SetId(created.ID)

	return resourceSecurityGroupRead(ctx, d, meta)
}

func resourceSecurityGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).SecurityGroups

	securityGroup, err := proxy.Read(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			d.SetId("")
			return diag.Errorf("security group with id %s not found: %+v", d.Id(), err)
		}

		return diag.Errorf("error while reading security group: %+v", err)
	}

	if err := d.Set("display_name", securityGroup.DisplayName); err != nil {
		return diag.Errorf("error setting displayName: %+v", err)
	}

	if err := d.Set("members", flattenSecurityGroupMembers(securityGroup.Members)); err != nil {
		return diag.Errorf("error setting members: %+v", err)
	}

	return nil
}

func resourceSecurityGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).SecurityGroups

	securityGroup := mapResourceDataToSecurityGroup(d)
	securityGroup.ID = d.Id()

	_, err := proxy.Update(ctx, securityGroup)
	if err != nil {
		return diag.Errorf("error while modifying security group: %+v", err)
	}

	return resourceSecurityGroupRead(ctx, d, meta)
}

func resourceSecurityGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).SecurityGroups

	err := proxy.Delete(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			d.SetId("")
			return diag.Errorf("security group with id %s not found: %+v", d.Id(), err)
		}

		return diag.Errorf("error while reading deleting group: %+v", err)
	}

	return nil
}

func mapResourceDataToSecurityGroup(d *schema.ResourceData) *models.SecurityGroup {
	return &models.SecurityGroup{
		DisplayName: d.Get("display_name").(string),
		Members:     expandSecurityGroupMembers(d.Get("members").(*schema.Set).List()),
	}
}
