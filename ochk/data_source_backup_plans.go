package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBackupPlans() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBackupPlansRead,
		Schema: map[string]*schema.Schema{
			"backup_plans": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_plan_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceBackupPlansRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	proxy := meta.(*sdk.Client).BackupPlans

	backupPlans, err := proxy.ListBackupPlans(ctx)

	if err := d.Set("backup_plans", flattenBackupPlans(backupPlans)); err != nil {
		return diag.Errorf("error setting backup plans: %v", err)
	}

	if err != nil {
		return diag.Errorf("error while listing backup plans")
	}

	d.SetId("backup-plans-list")

	return nil
}
