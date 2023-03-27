package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFolders() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFoldersRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"folders": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"folder_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"folder_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"folder_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceFoldersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	proxy := meta.(*sdk.Client).Folders

	projectID := d.Get("project_id").(string)

	folders, err := proxy.LisFoldersByProjectId(ctx, projectID)

	if err := d.Set("folders", flattenFolders(folders)); err != nil {
		return diag.Errorf("error setting folders list: %v", err)
	}

	if err != nil {
		return diag.Errorf("error while listing folders")
	}

	d.SetId("folders-list-" + projectID)

	return nil
}
