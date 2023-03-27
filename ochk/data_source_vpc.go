package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVpc() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVpcRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vrf_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"folder_path": {
				Type:     schema.TypeString,
				Computed: true,
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

func dataSourceVpcRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Routers

	displayName := d.Get("display_name").(string)

	routers, err := proxy.ListByDisplayName(ctx, displayName)
	if err != nil {
		return diag.Errorf("error while listing vpcs: %+v", err)
	}

	if len(routers) < 1 {
		return diag.Errorf("no vpc found for display_name: %s", displayName)
	}

	if len(routers) > 1 {
		return diag.Errorf("more than one vpc with display_name: %s found!", displayName)
	}

	d.SetId(routers[0].RouterID)

	if err := d.Set("vrf_id", routers[0].ParentT0ID); err != nil {
		return diag.Errorf("error setting vrf_id: %+v", err)
	}

	if err := d.Set("project_id", routers[0].ProjectID); err != nil {
		return diag.Errorf("error setting project_id: %+v", err)
	}

	if err := d.Set("folder_path", routers[0].FolderPath); err != nil {
		return diag.Errorf("error setting folder_path: %+v", err)
	}

	if err := d.Set("created_by", routers[0].CreatedBy); err != nil {
		return diag.Errorf("error setting created_by: %+v", err)
	}

	if err := d.Set("created_at", routers[0].CreationDate.String()); err != nil {
		return diag.Errorf("error setting created_at: %+v", err)
	}

	if err := d.Set("modified_by", routers[0].ModifiedBy); err != nil {
		return diag.Errorf("error setting modified_by: %+v", err)
	}

	if err := d.Set("modified_at", routers[0].ModificationDate.String()); err != nil {
		return diag.Errorf("error setting modified_at: %+v", err)
	}

	return nil
}
