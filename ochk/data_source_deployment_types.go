package ochk

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeploymentTypes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDeploymentTypesRead,

		Schema: map[string]*schema.Schema{
			"deployment_types": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deployment_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDeploymentTypesRead(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {

	var out []map[string]interface{}
	in := [3]string{"TEMPLATE", "OVF", "ISO"}

	for _, v := range in {
		m := make(map[string]interface{})
		m["deployment_type"] = v
		out = append(out, m)
	}

	if err := d.Set("deployment_types", out); err != nil {
		return diag.Errorf("error setting deployment types list: %v", err)
	}

	d.SetId("deployment-types-list")

	return nil
}
