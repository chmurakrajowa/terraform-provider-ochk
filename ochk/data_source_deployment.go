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
			"deployment_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"deployment_category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"initial_size_gb": {
				Type:     schema.TypeInt,
				Computed: true,
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

	if err := d.Set("deployment_type", deployments[0].DeploymentType); err != nil {
		return diag.Errorf("error setting deployment_type: %+v", err)
	}

	if err := d.Set("deployment_category", deployments[0].DeploymentCategory); err != nil {
		return diag.Errorf("error setting deployment_category: %+v", err)
	}

	if err := d.Set("initial_size_gb", int(deployments[0].DeploymentInitialSizeGB)); err != nil {
		return diag.Errorf("error setting initial_size_gb: %+v", err)
	}

	d.SetId(deployments[0].DeploymentID.String())

	return nil
}
