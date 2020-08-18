package ochk

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk"
)

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

	securityPolicies, err := proxy.ListByDisplayName(ctx, displayName)
	if err != nil {
		return diag.Errorf("error while listing security policies: %+v", err)
	}

	if len(securityPolicies) < 1 {
		return diag.Errorf("no security policy found for display_name: %s", displayName)
	}

	if len(securityPolicies) > 1 {
		return diag.Errorf("more than one security policy with display_name: %s found!", displayName)
	}

	d.SetId(securityPolicies[0].SecurityPolicyID)

	return nil
}
