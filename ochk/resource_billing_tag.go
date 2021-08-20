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
	BillingTagRetryTimeout = 1 * time.Minute
)

func resourceBillingTag() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBillingTagCreate,
		ReadContext:   resourceBillingTagRead,
		UpdateContext: resourceBillingTagUpdate,
		DeleteContext: resourceBillingTagDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(BillingTagRetryTimeout),
			Update: schema.DefaultTimeout(BillingTagRetryTimeout),
			Delete: schema.DefaultTimeout(BillingTagRetryTimeout),
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

func resourceBillingTagCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).BillingTags

	BillingTag := mapResourceDataToBillingTag(d)

	created, err := proxy.Create(ctx, BillingTag)
	if err != nil {
		return diag.Errorf("error while creating billing tag: %+v", err)
	}

	d.SetId(fmt.Sprint(created.BillingTagID))
	return resourceBillingTagRead(ctx, d, meta)
}

func resourceBillingTagRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).BillingTags

	var billingTagIDInt32 int32
	fmt.Sscan(d.Id(), &billingTagIDInt32)

	BillingTag, err := proxy.Read(ctx, billingTagIDInt32)
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("billing tag with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading billing tag: %+v", err)
	}

	if err := d.Set("display_name", BillingTag.TagValue); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	return nil
}

func resourceBillingTagUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).BillingTags

	var BillingTagIDInt32 int32
	fmt.Sscan(d.Id(), &BillingTagIDInt32)

	BillingTag := mapResourceDataToBillingTag(d)
	BillingTag.BillingTagID = BillingTagIDInt32

	_, err := proxy.Update(ctx, BillingTag)
	if err != nil {
		return diag.Errorf("error while modifying billing tag: %+v", err)
	}
	return resourceBillingTagRead(ctx, d, meta)
}

func resourceBillingTagDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).BillingTags

	err := proxy.Delete(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("billing tag with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting billing tag: %+v", err)
	}

	return nil
}

func mapResourceDataToBillingTag(d *schema.ResourceData) *models.BillingTag {
	return &models.BillingTag{
		TagValue: d.Get("display_name").(string),
	}
}
