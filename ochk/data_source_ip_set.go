package ochk

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
)

func dataSourceIPSet() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIPSetRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"addresses": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"address": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIPSetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).IPSets

	displayName := d.Get("display_name").(string)

	ipSets, err := proxy.ListByDisplayName(ctx, displayName)
	if err != nil {
		return diag.Errorf("error while listing IP sets: %+v", err)
	}

	if len(ipSets) < 1 {
		return diag.Errorf("no IP set found for display_name: %s", displayName)
	}

	if len(ipSets) > 1 {
		return diag.Errorf("more than one IP set with display_name: %s found!", displayName)
	}

	d.SetId(ipSets[0].IPSetID)

	if err := d.Set("addresses", flattenIPSetAddresses(ipSets[0].IPSetAddresses)); err != nil {
		return diag.Errorf("error setting addresses: %+v", err)
	}

	return nil
}
