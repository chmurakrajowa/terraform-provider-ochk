package ochk

import (
	"context"
	"fmt"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	SecurityGroupRetryTimeout = 1 * time.Minute
)

func resourceSecurityGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceGroupCreate,
		Read:   resourceServiceGroupRead,
		Update: resourceServiceGroupUpdate,
		Delete: resourceServiceGroupDelete,

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
				Type:     schema.TypeList,
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

func resourceServiceGroupCreate(d *schema.ResourceData, meta interface{}) error {
	proxy := meta.(*sdk.Client).SecurityGroups

	ctx := context.Background()

	securityGroup := &models.SecurityGroup{
		DisplayName: d.Get("display_name").(string),
		Members:     expandSecurityGroupMembers(d.Get("members").([]interface{})),
	}

	created, err := proxy.Create(ctx, securityGroup)
	if err != nil {
		return fmt.Errorf("error while creating security group: %w", err)
	}

	d.SetId(created.ID)

	return nil
}

func resourceServiceGroupRead(d *schema.ResourceData, meta interface{}) error {
	proxy := meta.(*sdk.Client).SecurityGroups

	ctx := context.Background()

	securityGroup, err := proxy.Read(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			d.SetId("")
			return fmt.Errorf("security group with id %s not found: %w", d.Id(), err)
		}

		return fmt.Errorf("error while reading security group: %w", err)
	}

	if err := d.Set("display_name", securityGroup.DisplayName); err != nil {
		return fmt.Errorf("error setting displayName: %w", err)
	}

	if err := d.Set("members", flattenSecurityGroupMembers(securityGroup.Members)); err != nil {
		return fmt.Errorf("error setting members: %w", err)
	}

	return nil
}

func resourceServiceGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("updating SecurityGroup is not implemented")
}

func resourceServiceGroupDelete(d *schema.ResourceData, meta interface{}) error {
	proxy := meta.(*sdk.Client).SecurityGroups

	ctx := context.Background()

	err := proxy.Delete(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			d.SetId("")
			return fmt.Errorf("security group with id %s not found: %w", d.Id(), err)
		}

		return fmt.Errorf("error while reading deleting group: %w", err)
	}

	return nil
}
