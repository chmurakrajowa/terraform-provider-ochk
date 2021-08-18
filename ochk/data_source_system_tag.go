package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSystemTag() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSystemTagRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceSystemTagRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	proxy := meta.(*sdk.Client).SystemTags

	systemTagName := d.Get("display_name").(string)

	systemTags, err := proxy.ListSystemTagsByTagName(ctx, systemTagName)

	if err != nil {
		return diag.Errorf("error while listing billing tags: %+v", err)
	}

	if len(systemTags) < 1 {
		return diag.Errorf("no billing tag for name: %s", systemTagName)
	}

	if len(systemTags) > 1 {
		return diag.Errorf("more than one billing tag with name: %s found!", systemTagName)
	}

	d.SetId(fmt.Sprint(systemTags[0].SystemTagID))

	if err := d.Set("display_name", systemTags[0].TagValue); err != nil {
		return diag.Errorf("error setting billing tag name: %+v", err)
	}
	return nil
}
