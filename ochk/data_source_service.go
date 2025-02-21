package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceService() *schema.Resource {
	return &schema.Resource{
		ReadContext: datSourceServiceRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
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

func datSourceServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Services

	displayName := d.Get("display_name").(string)

	services, err := proxy.ListByDisplayName(ctx, displayName)
	if err != nil {
		return diag.Errorf("error while listing services: %+v", err)
	}

	if len(services) < 1 {
		return diag.Errorf("no service found for display_name: %s", displayName)
	}

	if len(services) > 1 {
		return diag.Errorf("more than one service with display_name: %s found!", displayName)
	}

	d.SetId(services[0].ServiceID.String())

	if err := d.Set("created_by", services[0].CreatedBy); err != nil {
		return diag.Errorf("error setting created_by: %+v", err)
	}

	if err := d.Set("created_at", services[0].CreationDate.String()); err != nil {
		return diag.Errorf("error setting created_at: %+v", err)
	}

	if err := d.Set("modified_by", services[0].ModifiedBy); err != nil {
		return diag.Errorf("error setting modified_by: %+v", err)
	}

	if err := d.Set("modified_at", services[0].ModificationDate.String()); err != nil {
		return diag.Errorf("error setting modified_at: %+v", err)
	}

	return nil
}
