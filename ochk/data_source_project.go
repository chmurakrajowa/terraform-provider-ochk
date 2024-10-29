package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProject() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProjectRead,
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vrf_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"limits_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"memory_reserved_size_mb": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"storage_reserved_size_gb": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vcpu_reserved_quantity": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Projects

	name := d.Get("display_name").(string)

	projects, err := proxy.ListByName(ctx, name)
	if err != nil {
		return diag.Errorf("error while listing projects: %+v", err)
	}

	if len(projects) < 1 {
		return diag.Errorf("no project found for name: %s", name)
	}

	if len(projects) > 1 {
		return diag.Errorf("more than one project with name: %s found!", name)
	}

	if err := d.Set("vrf_id", projects[0].VrfID); err != nil {
		return diag.Errorf("error setting vrf_id:  %+v", err)
	}

	d.SetId(projects[0].ProjectID.String())

	if err := d.Set("display_name", projects[0].Name); err != nil {
		return diag.Errorf("error setting name: %s", err)
	}

	if err := d.Set("vrf_id", projects[0].VrfID); err != nil {
		return diag.Errorf("error setting vrf_id: %s", err)
	}

	if err := d.Set("description", projects[0].Description); err != nil {
		return diag.Errorf("error setting description: %s", err)
	}

	if err := d.Set("limits_enabled", projects[0].LimitEnabled); err != nil {
		return diag.Errorf("error setting limits_enabled: %s", err)
	}

	if err := d.Set("memory_reserved_size_mb", projects[0].MemoryReservedSizeMB); err != nil {
		return diag.Errorf("error setting memory_reserved_size_mb: %s", err)
	}

	if err := d.Set("storage_reserved_size_gb", projects[0].StorageReservedSizeGB); err != nil {
		return diag.Errorf("error setting storage_reserved_size_gb: %s", err)
	}

	if err := d.Set("vcpu_reserved_quantity", projects[0].CPUReserved); err != nil {
		return diag.Errorf("error setting vcpu_reserved_quantity: %s", err)
	}
	return nil
}
