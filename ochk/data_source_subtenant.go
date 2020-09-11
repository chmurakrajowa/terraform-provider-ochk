package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSubtenant() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSubtenantRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory_reserved_size_mb": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"storage_reserved_size_gb": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"users": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"networks": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceSubtenantRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Subtenants

	name := d.Get("name").(string)

	subtenants, err := proxy.ListByName(ctx, name)
	if err != nil {
		return diag.Errorf("error while listing subtenants: %+v", err)
	}

	if len(subtenants) < 1 {
		return diag.Errorf("no subtenant found for name: %s", name)
	}

	if len(subtenants) > 1 {
		return diag.Errorf("more than one subtenant with name: %s found!", name)
	}

	d.SetId(subtenants[0].SubtenantID)

	if err := mapSubtenantToResourceData(d, subtenants[0]); err != nil {
		return diag.Errorf("error while mapping subtenant to resource data: %+v", err)
	}

	return nil
}
