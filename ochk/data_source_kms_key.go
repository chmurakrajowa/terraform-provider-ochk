package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKMSKey() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKMSKeyRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"version": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"key_usage": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"activation_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_iv": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"object_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"revocation_reason": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sha1_fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sha256_fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceKMSKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).KMSKeys

	displayName := d.Get("display_name").(string)
	version := d.Get("version").(int)

	kmsKeys, err := proxy.ListByDisplayName(ctx, displayName)
	if err != nil {
		return diag.Errorf("error while listing KMS keys: %+v", err)
	}

	if len(kmsKeys) < 1 {
		return diag.Errorf("no KMS key found for display_name: %s", displayName)
	}

	var keyInstance *models.KeyInstance
	for _, key := range kmsKeys {
		if int(key.Version) == version {
			keyInstance = key
			break
		}
	}

	if keyInstance == nil {
		return diag.Errorf("no KMS key found for display_name: %s and version: %d", displayName, version)
	}

	if err := mapKMSKeyToResourceData(d, keyInstance); err != nil {
		return nil
	}
	d.SetId(keyInstance.ID)

	return nil
}
