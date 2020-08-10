package ochk

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk"
)

func dataSourceRouter() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRouterRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceRouterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Routers

	displayName := d.Get("display_name").(string)

	services, err := proxy.ListByDisplayName(ctx, displayName)
	if err != nil {
		return diag.Errorf("error while listing routers: %+v", err)
	}

	if len(services) < 1 {
		return diag.Errorf("no router found for display_name: %s", displayName)
	}

	if len(services) > 1 {
		return diag.Errorf("more than one router with display_name: %s found!", displayName)
	}

	d.SetId(services[0].RouterID)

	return nil
}
