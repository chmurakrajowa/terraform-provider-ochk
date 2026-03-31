package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	PublicIpRetryTimeout = 5 * time.Minute
)

func resourcePublicIp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePublicIpCreate,
		ReadContext:   resourcePublicIpRead,
		UpdateContext: resourcePublicIpUpdate,
		DeleteContext: resourcePublicIpDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(PublicIpRetryTimeout),
			Update: schema.DefaultTimeout(PublicIpRetryTimeout),
			Delete: schema.DefaultTimeout(PublicIpRetryTimeout),
		},

		Importer: &schema.ResourceImporter{
			StateContext: resourcePublicIpImportState,
		},

		Schema: map[string]*schema.Schema{
			"public_ip_address_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourcePublicIpImportState(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	d.SetId(strings.ToLower(d.Id()))
	return []*schema.ResourceData{d}, nil
}

func resourcePublicIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).PublicIPAddresses
	proxy_pt := meta.(*sdk.Client).PlatformType
	platformType, _ := proxy_pt.Read(ctx)

	public_ip_address := mapResourceDataToPublicIp(d, platformType)

	created, err := proxy.Create(ctx, public_ip_address, 1000)
	if err != nil {
		return diag.Errorf("error while creating public ip address: %+v", err)
	}

	d.SetId(created.GetRequestId())
	return resourcePublicIpRead(ctx, d, meta)
}

func resourcePublicIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

func mapPublicIpToResourceData(d *schema.ResourceData, project openapi.ProjectInstance) error {

	return nil
}

func resourcePublicIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
}

func resourcePublicIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).PublicIPAddresses

	err := proxy.Delete(ctx, d.Id()) //Delete(ctx, strfmt.UUID(d.Id()))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("public ip address with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting public ip address: %+v", err)
	}

	return nil
}

func mapResourceDataToPublicIp(d *schema.ResourceData, platformType openapi.PlatformType) openapi.PublicIpAllocation {
	PublicIpAddressValue := d.Get("ip_address").(string)
	IpAddressValue := d.Get("public_ip_address_id").(string)
	PublicIpAddress := openapi.PublicIpAddress{}
	PublicIpAddress.IpAddressId = NewNullableString(IpAddressValue)
	PublicIpAddress.IpAddress = NewNullableString(PublicIpAddressValue)

	return openapi.PublicIpAllocation{
		PublicIpAddress: &PublicIpAddress,
		Name:            NewNullableString(d.Get("display_name").(string)),
	}
}
