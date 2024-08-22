package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBackupPlan() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBackupPlanRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceBackupPlanRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	proxy := meta.(*sdk.Client).BackupPlans

	backupPlanName := d.Get("display_name").(string)

	backupPlans, err := proxy.ListBackupPlanByName(ctx, backupPlanName)

	if err != nil {
		return diag.Errorf("error while listing backup plans: %+v", err)
	}

	if len(backupPlans) < 1 {
		return diag.Errorf("no backup plan for name: %s", backupPlanName)
	}

	if len(backupPlans) > 1 {
		return diag.Errorf("more than one backup plan with name: %s found!", backupPlanName)
	}

	d.SetId(backupPlans[0].BackupPlanID.String())

	if err := d.Set("display_name", backupPlans[0].BackupPlanName); err != nil {
		return diag.Errorf("error setting backup plan name: %+v", err)
	}
	return nil
}
