package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ProjectRetryTimeout = 5 * time.Minute
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(ProjectRetryTimeout),
			Update: schema.DefaultTimeout(ProjectRetryTimeout),
			Delete: schema.DefaultTimeout(ProjectRetryTimeout),
		},

		Importer: &schema.ResourceImporter{
			StateContext: resourceProjectImportState,
		},

		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vrf_id": {
				Type:     schema.TypeString,
				Required: true,
				StateFunc: func(val any) string {
					return strings.ToLower(val.(string))
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limits_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"memory_reserved_size_mb": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"storage_reserved_size_gb": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vcpu_reserved_quantity": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceProjectImportState(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	d.SetId(strings.ToLower(d.Id()))
	return []*schema.ResourceData{d}, nil
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Projects
	proxy_pt := meta.(*sdk.Client).PlatformType
	platformType, _ := proxy_pt.Read(ctx)

	project := mapResourceDataToProject(d, platformType)

	created, err := proxy.Create(ctx, project)
	if err != nil {
		return diag.Errorf("error while creating project: %+v", err)
	}

	d.SetId(created.ProjectID.String())
	return resourceProjectRead(ctx, d, meta)
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Projects

	project, err := proxy.Read(ctx, strfmt.UUID(d.Id()))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("project with id %s not found: %+v", id, err)
		}
		return diag.Errorf("error while reading project: %+v", err)
	}

	if err := mapProjectToResourceData(d, project); err != nil {
		return diag.Errorf("error while mapping project to resource data: %+v", err)
	}

	return nil
}

func mapProjectToResourceData(d *schema.ResourceData, project *models.ProjectInstance) error {
	if err := d.Set("project_id", strings.ToLower(strfmt.UUID.String(project.ProjectID))); err != nil {
		return fmt.Errorf("error setting project_id: %w", err)
	}

	if err := d.Set("display_name", project.Name); err != nil {
		return fmt.Errorf("error setting name: %w", err)
	}

	if err := d.Set("vrf_id", project.VrfID); err != nil {
		return fmt.Errorf("error setting vrf_id: %w", err)
	}

	if err := d.Set("description", project.Description); err != nil {
		return fmt.Errorf("error setting description: %w", err)
	}

	if err := d.Set("limits_enabled", project.LimitEnabled); err != nil {
		return fmt.Errorf("error setting limits_enabled: %w", err)
	}

	if err := d.Set("memory_reserved_size_mb", project.MemoryReservedSizeMB); err != nil {
		return fmt.Errorf("error setting memory_reserved_size_mb: %w", err)
	}

	if err := d.Set("storage_reserved_size_gb", project.StorageReservedSizeGB); err != nil {
		return fmt.Errorf("error setting storage_reserved_size_gb: %w", err)
	}

	if err := d.Set("vcpu_reserved_quantity", project.CPUReserved); err != nil {
		return fmt.Errorf("error setting vcpu_reserved_quantity: %w", err)
	}
	return nil
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Projects

	proxy_pt := meta.(*sdk.Client).PlatformType
	platformType, _ := proxy_pt.Read(ctx)

	project := mapResourceDataToProject(d, platformType)
	project.ProjectID = strfmt.UUID(d.Id())

	_, err := proxy.Update(ctx, project)
	if err != nil {
		return diag.Errorf("error while modifying project: %+v", err)
	}

	return resourceProjectRead(ctx, d, meta)
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Projects

	err := proxy.Delete(ctx, strfmt.UUID(d.Id()))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("project with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting project: %+v", err)
	}

	return nil
}

func mapResourceDataToProject(d *schema.ResourceData, platformType models.PlatformType) *models.ProjectInstance {
	var factor int64 = 1
	if platformType == "OPENSTACK" {
		factor = 1024
	}
	return &models.ProjectInstance{
		Description:           d.Get("description").(string),
		MemoryReservedSizeMB:  int64(d.Get("memory_reserved_size_mb").(int)) * factor,
		Name:                  d.Get("display_name").(string),
		StorageReservedSizeGB: int64(d.Get("storage_reserved_size_gb").(int)),
		VrfID:                 strfmt.UUID(d.Get("vrf_id").(string)),
		CPUReserved:           int64(d.Get("vcpu_reserved_quantity").(int)),
		LimitEnabled:          d.Get("limits_enabled").(bool),
		ProjectID:             strfmt.UUID(d.Get("project_id").(string)),
	}
}
