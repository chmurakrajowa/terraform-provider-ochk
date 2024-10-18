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
	VpcRetryTimeout = 1 * time.Minute
	E2001           = "TF_ERROR{2001}: Error while creating router(vpc): %+v"
	E2001_UPDATE    = "TF_ERROR{2001}: Error while updating router(vpc): %+v"
	E2002           = "Field %s is not applicable for platform VMWARE. Please remove field from Your input config file."
	E2003           = "Wrong field in resource."
	E2004           = "Error while mapping router (vpc) to resource: %+v"
	E2005           = "Error while reading vpc: %+v"
	E2006           = "Error while deleting vpc: %+v"
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
	proxy_pt := meta.(*sdk.Client).PlatformType

	platformType, _ := proxy_pt.Read(ctx)
	Router, err_vpc := mapResourceDataToVpc(d, platformType)

	if err_vpc != nil {
		return diag.Errorf(E2004, err_vpc)
	}

	created, err := proxy.Create(ctx, Router)
	if err != nil {
		return diag.Errorf(E2001, err)
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

		return diag.Errorf(E2005, err)
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
		return diag.Errorf("error setting autonat_enabled  uuuuuuuuu : %+v", err)
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
	proxy_pt := meta.(*sdk.Client).PlatformType
	platformType, _ := proxy_pt.Read(ctx)

	Router, err_vpc := mapResourceDataToVpc(d, platformType)

	if err_vpc != nil {
		return diag.Errorf(E2004, err_vpc)
	}

	Router.RouterID = strfmt.UUID(d.Id())

	_, err := proxy.Update(ctx, Router)
	if err != nil {
		return diag.Errorf(E2001_UPDATE, err)
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

		return diag.Errorf(E2006, err)
	}

	return nil
}

func mapResourceDataToVpc(d *schema.ResourceData, platformType models.PlatformType) (*models.RouterInstance, diag.Diagnostics) {
	if platformType == models.PlatformTypeOPENSTACK {
		return &models.RouterInstance{
			DisplayName: d.Get("display_name").(string),
			ParentT0ID:  strfmt.UUID(d.Get("vrf_id").(string)),
			ProjectID:   strfmt.UUID(d.Get("project_id").(string)),
			SnatEnabled: d.Get("autonat_enabled").(bool),
			FolderPath:  d.Get("folder_path").(string),
		}, nil
	} else {
		if d.Get("autonat_enabled").(bool) {
			return nil, diag.Errorf(E2003, fmt.Sprintf(E2002, "autonat_enabled"))
		}
		return &models.RouterInstance{
			DisplayName: d.Get("display_name").(string),
			ParentT0ID:  strfmt.UUID(d.Get("vrf_id").(string)),
			ProjectID:   strfmt.UUID(d.Get("project_id").(string)),
			FolderPath:  d.Get("folder_path").(string),
		}, nil
	}
}
