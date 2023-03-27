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
	TagRetryTimeout = 1 * time.Minute
)

func resourceTag() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTagCreate,
		ReadContext:   resourceTagRead,
		UpdateContext: resourceTagUpdate,
		DeleteContext: resourceTagDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(TagRetryTimeout),
			Update: schema.DefaultTimeout(TagRetryTimeout),
			Delete: schema.DefaultTimeout(TagRetryTimeout),
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceTagCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Tags

	tag := mapResourceDataToTag(d)

	created, err := proxy.Create(ctx, tag)
	if err != nil {
		return diag.Errorf("error while creating tag: %+v", err)
	}

	d.SetId(fmt.Sprint(created.TagID))
	return resourceTagRead(ctx, d, meta)
}

func resourceTagRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Tags

	var tagIDInt32 int32
	_, err := fmt.Sscan(d.Id(), &tagIDInt32)
	if err != nil {
		return diag.Errorf("wrong tag id format: %+v", err)
	}

	tag, err := proxy.Read(ctx, tagIDInt32)
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("tag with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading tag: %+v", err)
	}

	if err := d.Set("display_name", tag.TagValue); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("project_id", tag.ProjectID); err != nil {
		return diag.Errorf("error setting project_id: %+v", err)
	}

	return nil
}

func resourceTagUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Tags

	var tagIDInt32 int32
	_, err := fmt.Sscan(d.Id(), &tagIDInt32)
	if err != nil {
		return diag.Errorf("wrong tag id format: %+v", err)
	}

	tag := mapResourceDataToTag(d)
	tag.TagID = tagIDInt32

	_, err = proxy.Update(ctx, tag)
	if err != nil {
		return diag.Errorf("error while modifying tag: %+v", err)
	}
	return resourceTagRead(ctx, d, meta)
}

func resourceTagDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Tags

	err := proxy.Delete(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("tag with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting tag: %+v", err)
	}

	return nil
}

func mapResourceDataToTag(d *schema.ResourceData) *models.Tag {
	return &models.Tag{
		TagValue:  d.Get("display_name").(string),
		ProjectID: d.Get("project_id").(string),
	}
}
