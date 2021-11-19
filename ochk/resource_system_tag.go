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
	SystemTagRetryTimeout = 1 * time.Minute
)

func resourceSystemTag() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSystemTagCreate,
		ReadContext:   resourceSystemTagRead,
		UpdateContext: resourceSystemTagUpdate,
		DeleteContext: resourceSystemTagDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(SystemTagRetryTimeout),
			Update: schema.DefaultTimeout(SystemTagRetryTimeout),
			Delete: schema.DefaultTimeout(SystemTagRetryTimeout),
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSystemTagCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).SystemTags

	SystemTag := mapResourceDataToSystemTag(d)

	created, err := proxy.Create(ctx, SystemTag)
	if err != nil {
		return diag.Errorf("error while creating system tag: %+v", err)
	}

	d.SetId(fmt.Sprint(created.SystemTagID))
	return resourceSystemTagRead(ctx, d, meta)
}

func resourceSystemTagRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).SystemTags

	var systemTagIDInt32 int32
	fmt.Sscan(d.Id(), &systemTagIDInt32)

	SystemTag, err := proxy.Read(ctx, systemTagIDInt32)
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("system tag with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading system tag: %+v", err)
	}

	if err := d.Set("display_name", SystemTag.TagValue); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	return nil
}

func resourceSystemTagUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).SystemTags

	var systemTagIDInt32 int32
	fmt.Sscan(d.Id(), &systemTagIDInt32)

	SystemTag := mapResourceDataToSystemTag(d)
	SystemTag.SystemTagID = systemTagIDInt32

	_, err := proxy.Update(ctx, SystemTag)
	if err != nil {
		return diag.Errorf("error while modifying system tag: %+v", err)
	}
	return resourceSystemTagRead(ctx, d, meta)
}

func resourceSystemTagDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).SystemTags

	err := proxy.Delete(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("system tag with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting system tag: %+v", err)
	}

	return nil
}

func mapResourceDataToSystemTag(d *schema.ResourceData) *models.SystemTag {
	return &models.SystemTag{
		TagValue: d.Get("display_name").(string),
	}
}
