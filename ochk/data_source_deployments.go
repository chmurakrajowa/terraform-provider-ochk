package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeployments() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDeploymentsRead,

		Schema: map[string]*schema.Schema{
			"deployment_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"deployments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deployment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"deployment_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"initial_size_mb": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDeploymentsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Deployments

	deploymentType := d.Get("deployment_type").(string)

	deployments, err := proxy.List(ctx)
	if err != nil {
		return diag.Errorf("error while listing deployments: %+v", err)
	}

	if len(deployments) < 1 {
		return diag.Errorf("no deployments found for display_type: %s", deploymentType)
	}

	if err := d.Set("deployments", flattenDeployments(deployments, deploymentType)); err != nil {
		return diag.Errorf("error setting deployments list: %v", err)
	}

	if deploymentType == "" {
		d.SetId("deployments-list")
	} else {
		d.SetId("deployments-list-" + deploymentType)
	}

	return nil
}
