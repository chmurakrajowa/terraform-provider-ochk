package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBackupList() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBackupListRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backup_plan_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceBackupListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	proxy := meta.(*sdk.Client).BackupLists

	backupListName := d.Get("display_name").(string)
	backupPlanId := d.Get("backup_plan_id").(string)

	backupLists, err := proxy.ListBackupListByName(ctx, backupPlanId, backupListName)

	if err != nil {
		return diag.Errorf("error while listing backup list: %+v", err)
	}

	if len(backupLists) < 1 {
		return diag.Errorf("no backup list for name: %s", backupListName)
	}

	if len(backupLists) > 1 {
		return diag.Errorf("more than one backup list with name: %s found!", backupListName)
	}

	d.SetId(backupLists[0].BackupListID)

	if err := d.Set("display_name", backupLists[0].BackupListName); err != nil {
		return diag.Errorf("error setting backup list name: %+v", err)
	}
	return nil
}
