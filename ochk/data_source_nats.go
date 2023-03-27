package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAutoNats() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAutoNatsRead,
		Schema: map[string]*schema.Schema{
			"auto_nats": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_nat_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"virtual_network_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"vrf_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceManualNats() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceManualNatsRead,
		Schema: map[string]*schema.Schema{
			"manual_nats": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manual_nat_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"vrf_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_network": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"destination_network": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAutoNatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client)

	nats, err := proxy.Nats.List(ctx)

	if err != nil {
		return diag.Errorf("error while listing nats: %+v", err)
	}

	if err := d.Set("auto_nats", flattenAutoNats(nats)); err != nil {
		return diag.Errorf("error setting auto_nats: %v", err)
	}

	d.SetId("auto-nats-list")

	return nil
}

func dataSourceManualNatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client)

	nats, err := proxy.Nats.List(ctx)

	if err != nil {
		return diag.Errorf("error while listing nats: %+v", err)
	}

	if err := d.Set("manual_nats", flattenManualNats(nats)); err != nil {
		return diag.Errorf("error setting manual_nats: %v", err)
	}

	d.SetId("maual-nats-list")

	return nil
}
