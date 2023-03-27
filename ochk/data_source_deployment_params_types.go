package ochk

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeploymentParamsTypes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDeploymentParamsTypesRead,

		Schema: map[string]*schema.Schema{
			"deployment_params_types": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deployment_param_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDeploymentParamsTypesRead(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {

	var out []map[string]interface{}
	in := []string{
		"NET_IP_ADDR_01",
		"NET_IP_MASK_01",
		"NET_IP_GATEWAY_01",
		"NET_IP_BROADCAST_01",
		"NET_DNS_PRIMARY",
		"NET_DNS_SECONDARY",
		"NET_DNS_SUFFIX",
		"NET_DNS_SEARCH",
		"NET_WINS_PREFERRED",
		"NET_WINS_ALTERNATIVE",
		"LOGIN_USERNAME",
		"LOGIN_PASSWORD",
		"LOGIN_SSH_KEY",
		"HOST_NAME",
	}

	for _, v := range in {
		m := make(map[string]interface{})
		m["deployment_param_type"] = v
		out = append(out, m)
	}

	if err := d.Set("deployment_params_types", out); err != nil {
		return diag.Errorf("error setting deployment params types list: %v", err)
	}

	d.SetId("deployment-params-types-list")

	return nil
}
