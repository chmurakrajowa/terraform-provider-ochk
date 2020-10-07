package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	RouterRetryTimeout = 1 * time.Minute
)

func resourceRouter() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouterCreate,
		ReadContext:   resourceRouterRead,
		UpdateContext: resourceRouterUpdate,
		DeleteContext: resourceRouterDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(RouterRetryTimeout),
			Update: schema.DefaultTimeout(RouterRetryTimeout),
			Delete: schema.DefaultTimeout(RouterRetryTimeout),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"router_type": {
				Type:     schema.TypeString,
				Computed: true,
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

func resourceRouterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Routers

	Router := mapResourceDataToRouter(d)

	created, err := proxy.Create(ctx, Router)
	if err != nil {
		return diag.Errorf("error while creating router: %+v", err)
	}

	d.SetId(created.RouterID)

	return resourceRouterRead(ctx, d, meta)
}

func resourceRouterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Routers

	Router, err := proxy.Read(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("router with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading router: %+v", err)
	}

	if err := d.Set("display_name", Router.DisplayName); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
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

func resourceRouterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Routers

	Router := mapResourceDataToRouter(d)
	Router.RouterID = d.Id()

	_, err := proxy.Update(ctx, Router)
	if err != nil {
		return diag.Errorf("error while modifying router: %+v", err)
	}

	return resourceRouterRead(ctx, d, meta)
}

func resourceRouterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Routers

	err := proxy.Delete(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("router with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting router: %+v", err)
	}

	return nil
}

func mapResourceDataToRouter(d *schema.ResourceData) *models.RouterInstance {
	return &models.RouterInstance{
		DisplayName: d.Get("display_name").(string),
	}
}
