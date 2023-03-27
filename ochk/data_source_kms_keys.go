package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKMSKeys() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKMSKeysRead,

		Schema: map[string]*schema.Schema{
			"kms_keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_usage": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceKMSKeysRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).KMSKeys

	kmsKeys, err := proxy.List(ctx)
	if err != nil {
		return diag.Errorf("error while listing KMS keys: %+v", err)
	}

	if err := d.Set("kms_keys", flattenKMSKeys(kmsKeys)); err != nil {
		return diag.Errorf("error setting ip_collections: %v", err)
	}

	d.SetId("kms-keys-list")

	return nil
}
