package ochk

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk"
)

func dataSourceService() *schema.Resource {
	return &schema.Resource{
		ReadContext: datSourceServiceRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
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

	d.SetId(services[0].ServiceID)

	return nil
}
