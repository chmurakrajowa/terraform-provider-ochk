package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	VpcRetryTimeout = 1 * time.Minute
)

func resourceVpc() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVpcCreate,
		ReadContext:   resourceVpcRead,
		UpdateContext: resourceVpcUpdate,
		DeleteContext: resourceVpcDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(VpcRetryTimeout),
			Update: schema.DefaultTimeout(VpcRetryTimeout),
			Delete: schema.DefaultTimeout(VpcRetryTimeout),
		},

		Importer: &schema.ResourceImporter{
			StateContext: resourceVpcImportState,
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vrf_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"autonat_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"folder_path": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/",
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceVpcImportState(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	d.SetId(strings.ToLower(d.Id()))
	return []*schema.ResourceData{d}, nil
}

func resourceVpcCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Routers

	Router := mapResourceDataToVpc(d)

	created, err := proxy.Create(ctx, Router)
	if err != nil {
		return diag.Errorf("error while creating vpc: %+v", err)
	}

	d.SetId(created.RouterID.String())

	return resourceVpcRead(ctx, d, meta)
}

func resourceVpcRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Routers

	Router, err := proxy.Read(ctx, strfmt.UUID(d.Id()))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("vpc with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading vpc: %+v", err)
	}

	if err := d.Set("vrf_id", Router.ParentT0ID); err != nil {
		return diag.Errorf("error setting vrf_id: %+v", err)
	}

	if err := d.Set("project_id", Router.ProjectID); err != nil {
		return diag.Errorf("error setting project_id: %+v", err)
	}

	if err := d.Set("display_name", Router.DisplayName); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("autonat_enabled", Router.SnatEnabled); err != nil {
		return diag.Errorf("error setting autonat_enabled: %+v", err)
	}

	if err := d.Set("folder_path", Router.FolderPath); err != nil {
		return diag.Errorf("error setting folder_path: %+v", err)
	}

	if err := d.Set("created_by", Router.CreatedBy); err != nil {
		return diag.Errorf("error setting created_by: %+v", err)
	}

	if err := d.Set("created_at", Router.CreationDate.String()); err != nil {
		return diag.Errorf("error setting created_at: %+v", err)
	}

	if err := d.Set("modified_by", Router.ModifiedBy); err != nil {
		return diag.Errorf("error setting modified_by: %+v", err)
	}

	if err := d.Set("modified_at", Router.ModificationDate.String()); err != nil {
		return diag.Errorf("error setting modified_at: %+v", err)
	}

	return nil
}

func resourceVpcUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Routers

	Router := mapResourceDataToVpc(d)
	Router.RouterID = strfmt.UUID(d.Id())

	_, err := proxy.Update(ctx, Router)
	if err != nil {
		return diag.Errorf("error while modifying vpc: %+v", err)
	}

	return resourceVpcRead(ctx, d, meta)
}

func resourceVpcDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Routers

	err := proxy.Delete(ctx, strfmt.UUID(d.Id()))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("vpc with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting vpc: %+v", err)
	}

	return nil
}

func mapResourceDataToVpc(d *schema.ResourceData) *models.RouterInstance {
	return &models.RouterInstance{
		DisplayName: d.Get("display_name").(string),
		ParentT0ID:  strfmt.UUID(d.Get("vrf_id").(string)),
		ProjectID:   strfmt.UUID(d.Get("project_id").(string)),
		SnatEnabled: d.Get("autonat_enabled").(bool),
		FolderPath:  d.Get("folder_path").(string),
	}
}
