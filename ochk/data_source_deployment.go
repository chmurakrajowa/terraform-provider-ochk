package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeployment() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDeploymentRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceDeploymentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Deployments

	displayName := d.Get("display_name").(string)

	deployments, err := proxy.ListByDisplayName(ctx, displayName)
	if err != nil {
		return diag.Errorf("error while listing deployments: %+v", err)
	}

	if len(deployments) < 1 {
		return diag.Errorf("no deployment found for display_name: %s", displayName)
	}

	if len(deployments) > 1 {
		return diag.Errorf("more than one deployment with display_name: %s found!", displayName)
	}

	d.SetId(deployments[0].DeploymentID)

	return nil
}
