package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTag() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTagRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceTagRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	proxy := meta.(*sdk.Client).Tags

	tagName := d.Get("display_name").(string)

	tags, err := proxy.ListTagsByTagName(ctx, tagName)

	if err != nil {
		return diag.Errorf("error while listing tags: %+v", err)
	}

	if len(tags) < 1 {
		return diag.Errorf("no tag for name: %s", tagName)
	}

	if len(tags) > 1 {
		return diag.Errorf("more than one billing tag with name: %s found!", tagName)
	}

	d.SetId(fmt.Sprint(tags[0].TagID))

	if err := d.Set("display_name", tags[0].TagValue); err != nil {
		return diag.Errorf("error setting tag name: %+v", err)
	}
	if err := d.Set("project_id", tags[0].ProjectID); err != nil {
		return diag.Errorf("error setting project_id: %+v", err)
	}

	return nil
}
