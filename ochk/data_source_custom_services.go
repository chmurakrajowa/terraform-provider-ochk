package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCustomServices() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCustomServicesRead,

		Schema: map[string]*schema.Schema{
			"custom_services": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"custom_service_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCustomServicesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).CustomServices

	customServices, err := proxy.ListCustomServices(ctx)
	if err != nil {
		return diag.Errorf("error while listing custom services: %+v", err)
	}

	if err := d.Set("custom_services", flattenCustomServices(customServices)); err != nil {
		return diag.Errorf("error setting custom services list: %v", err)
	}

	d.SetId("custom-services-list")

	return nil
}
