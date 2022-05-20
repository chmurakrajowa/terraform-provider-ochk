package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLocalGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLocalGroupRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"users": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"parent_groups": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"child_groups": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceLocalGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).LocalGroups

	name := d.Get("name").(string)

	localGroups, err := proxy.ListByName(ctx, name)
	if err != nil {
		return diag.Errorf("error while listing local groups: %+v", err)
	}

	if len(localGroups) < 1 {
		return diag.Errorf("no local group found for name: %s", name)
	}

	if len(localGroups) > 1 {
		return diag.Errorf("more than one local group with name: %s found!", name)
	}

	d.SetId(localGroups[0].GroupID)

	if err := d.Set("name", localGroups[0].Name); err != nil {
		return diag.Errorf("error setting local group name: %+v", err)
	}
	if err := d.Set("description", localGroups[0].Description); err != nil {
		return diag.Errorf("error setting local group description: %+v", err)
	}
	if err := d.Set("group_type", localGroups[0].GroupType); err != nil {
		return diag.Errorf("error setting local group_type: %+v", err)
	}
	if err := d.Set("domain", localGroups[0].Domain); err != nil {
		return diag.Errorf("error setting local group domain: %+v", err)
	}
	if err := d.Set("users", flattenUserInstancesFromIDs(localGroups[0].UserInstanceList)); err != nil {
		return diag.Errorf("error setting users: %+v", err)
	}
	if err := d.Set("child_groups", flattenGroupInstancesFromIDs(localGroups[0].GroupInstanceList)); err != nil {
		return diag.Errorf("error setting child groups: %+v", err)
	}
	if err := d.Set("parent_groups", flattenGroupInstancesFromIDs(localGroups[0].ParentGroups)); err != nil {
		return diag.Errorf("error setting parent groups: %+v", err)
	}
	return nil
}
