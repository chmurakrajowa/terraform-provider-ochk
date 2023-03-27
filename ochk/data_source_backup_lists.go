package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBackupLists() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBackupListsRead,
		Schema: map[string]*schema.Schema{
			"backup_plan_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_lists": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_list_id": {
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

func dataSourceBackupListsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	proxy := meta.(*sdk.Client).BackupLists

	backupPlanId := d.Get("backup_plan_id").(string)

	backupLists, err := proxy.ListBackupList(ctx, backupPlanId)

	if err := d.Set("backup_lists", flattenBackupLists(backupLists)); err != nil {
		return diag.Errorf("error setting backup list: %v", err)
	}

	if err != nil {
		return diag.Errorf("error while listing backup list")
	}

	d.SetId("backup-list-" + backupPlanId)

	return nil
}
