package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	SubtenantRetryTimeout = 5 * time.Minute
)

func resourceSubtenant() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSubtenantCreate,
		ReadContext:   resourceSubtenantRead,
		UpdateContext: resourceSubtenantUpdate,
		DeleteContext: resourceSubtenantDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(SubtenantRetryTimeout),
			Update: schema.DefaultTimeout(SubtenantRetryTimeout),
			Delete: schema.DefaultTimeout(SubtenantRetryTimeout),
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"memory_reserved_size_mb": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"storage_reserved_size_gb": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"users": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				MinItems: 1,
				ForceNew: true,
			},
			"networks": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceSubtenantCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Subtenants

	subtenant := mapResourceDataToSubtenant(d)

	created, err := proxy.Create(ctx, subtenant)
	if err != nil {
		return diag.Errorf("error while creating subtenant: %+v", err)
	}

	d.SetId(created.SubtenantID)

	return resourceSubtenantRead(ctx, d, meta)
}

func resourceSubtenantRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Subtenants

	subtenant, err := proxy.Read(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("subtenant with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading subtenant: %+v", err)
	}

	if err := mapSubtenantToResourceData(d, subtenant); err != nil {
		return diag.Errorf("error while mapping subtenant to resource data: %+v", err)
	}

	return nil
}

func mapSubtenantToResourceData(d *schema.ResourceData, subtenant *models.SubtenantInstance) error {
	if err := d.Set("name", subtenant.Name); err != nil {
		return fmt.Errorf("error setting name: %w", err)
	}

	if err := d.Set("email", subtenant.Email); err != nil {
		return fmt.Errorf("error setting email: %w", err)
	}

	if err := d.Set("description", subtenant.Description); err != nil {
		return fmt.Errorf("error setting description: %w", err)
	}

	if err := d.Set("memory_reserved_size_mb", subtenant.MemoryReservedSizeMB); err != nil {
		return fmt.Errorf("error setting memory_reserved_size_mb: %w", err)
	}

	if err := d.Set("storage_reserved_size_gb", subtenant.StorageReservedSizeGB); err != nil {
		return fmt.Errorf("error setting storage_reserved_size_gb: %w", err)
	}

	if err := d.Set("users", flattenUserInstancesFromIDs(subtenant.Users)); err != nil {
		return fmt.Errorf("error setting users: %w", err)
	}

	if err := d.Set("networks", flattenVCSNetworkInstancesFromIDs(subtenant.Networks)); err != nil {
		return fmt.Errorf("error setting networks: %w", err)
	}

	return nil
}

func resourceSubtenantUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Subtenants

	subtenant := mapResourceDataToSubtenant(d)
	subtenant.SubtenantID = d.Id()

	_, err := proxy.Update(ctx, subtenant)
	if err != nil {
		return diag.Errorf("error while modifying subtenant: %+v", err)
	}

	return resourceSubtenantRead(ctx, d, meta)
}

func resourceSubtenantDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Subtenants

	err := proxy.Delete(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("subtenant with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting subtenant: %+v", err)
	}

	return nil
}

func mapResourceDataToSubtenant(d *schema.ResourceData) *models.SubtenantInstance {
	return &models.SubtenantInstance{
		Description:           d.Get("description").(string),
		Email:                 d.Get("email").(string),
		MemoryReservedSizeMB:  int64(d.Get("memory_reserved_size_mb").(int)),
		Name:                  d.Get("name").(string),
		Networks:              expandVCSNetworkInstancesFromIDs(d.Get("networks").(*schema.Set).List()),
		StorageReservedSizeGB: int64(d.Get("storage_reserved_size_gb").(int)),
		Users:                 expandUserInstancesFromIDs(d.Get("users").(*schema.Set).List()),
	}
}
