package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBillingTag() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBillingTagRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceBillingTagRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	proxy := meta.(*sdk.Client).BillingTags

	billingTagName := d.Get("display_name").(string)

	billingTags, err := proxy.ListBillingTagsByTagName(ctx, billingTagName)

	if err != nil {
		return diag.Errorf("error while listing billing tags: %+v", err)
	}

	if len(billingTags) < 1 {
		return diag.Errorf("no billing tag for name: %s", billingTagName)
	}

	if len(billingTags) > 1 {
		return diag.Errorf("more than one billing tag with name: %s found!", billingTagName)
	}

	d.SetId(fmt.Sprint(billingTags[0].BillingTagID))

	if err := d.Set("display_name", billingTags[0].TagValue); err != nil {
		return diag.Errorf("error setting billing tag name: %+v", err)
	}
	return nil
}
