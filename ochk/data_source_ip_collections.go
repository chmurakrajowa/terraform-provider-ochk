package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIPCollections() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIPCollectionsRead,

		Schema: map[string]*schema.Schema{
			"ip_collections": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_collection_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_addresses": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"project_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIPCollectionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).IPCollections

	ipCollections, err := proxy.List(ctx)
	if err != nil {
		return diag.Errorf("error while listing ip_collections: %+v", err)
	}

	if err := d.Set("ip_collections", flattenIPCollections(ipCollections)); err != nil {
		return diag.Errorf("error setting ip_collections: %v", err)
	}

	d.SetId("ip-collections-list")

	return nil
}
