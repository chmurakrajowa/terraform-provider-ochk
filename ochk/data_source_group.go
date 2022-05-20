package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGroupRead,

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
			"child_groups": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Groups

	name := d.Get("name").(string)

	groups, err := proxy.ListByName(ctx, name)
	if err != nil {
		return diag.Errorf("error while listing groups: %+v", err)
	}

	if len(groups) < 1 {
		return diag.Errorf("no group found for name: %s", name)
	}

	if len(groups) > 1 {
		return diag.Errorf("more than one group with name: %s found!", name)
	}

	d.SetId(groups[0].GroupID)

	if err := d.Set("name", groups[0].Name); err != nil {
		return diag.Errorf("error setting group name: %+v", err)
	}
	if err := d.Set("description", groups[0].Description); err != nil {
		return diag.Errorf("error setting group description: %+v", err)
	}
	if err := d.Set("group_type", groups[0].GroupType); err != nil {
		return diag.Errorf("error setting group_type: %+v", err)
	}
	if err := d.Set("domain", groups[0].Domain); err != nil {
		return diag.Errorf("error setting group domain: %+v", err)
	}
	if err := d.Set("users", flattenUserInstancesFromIDs(groups[0].UserInstanceList)); err != nil {
		return diag.Errorf("error setting users: %+v", err)
	}
	if err := d.Set("child_groups", flattenGroupInstancesFromIDs(groups[0].GroupInstanceList)); err != nil {
		return diag.Errorf("error setting child groups: %+v", err)
	}
	return nil
}
