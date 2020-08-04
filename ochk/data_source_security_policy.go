package ochk

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk"
)

//TODO brak testu akceptacyjnego
//TODO pytanie czy chcemy więcej pól odwzorować czy tyle wystarczy
func dataSourceSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: datSourceSecurityPolicyRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func datSourceSecurityPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).SecurityPolicy

	displayName := d.Get("display_name").(string)

	services, err := proxy.ListByDisplayName(ctx, displayName)
	if err != nil {
		return diag.Errorf("error while listing security policies: %+v", err)
	}

	if len(services) < 1 {
		return diag.Errorf("no security policy found for display_name: %s", displayName)
	}

	if len(services) > 1 {
		return diag.Errorf("more than one security policy with display_name: %s found!", displayName)
	}

	d.SetId(services[0].SecurityPolicyID)

	return nil
}
