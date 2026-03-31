package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"strconv"
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

	public_ip_address := mapResourceDataToPublicIp(d)

	created, err := proxy.Create(ctx, public_ip_address)
	if err != nil {
		return diag.Errorf("error while creating public ip address: %+v", err)
	}
	allocation_id_value := (string)(created.GetAllocationId())
	d.SetId(allocation_id_value)
	return resourcePublicIpRead(ctx, d, meta)
}

func resourcePublicIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	proxy := meta.(*sdk.Client).PublicIPAddresses

	public_ip_id, _ := strconv.Atoi(d.Id())
	public_ip, err := proxy.Read(ctx, (int32)(public_ip_id))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("public ip with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading public ip: %+v", err)
	}

	if err := d.Set("display_name", public_ip.GetName()); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("ip_address", public_ip.GetPublicIpAddress().IpAddress); err != nil {
		return diag.Errorf("error setting description: %+v", err)
	}

	return nil
}

func resourcePublicIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	if !d.HasChanges("display_name") {
		return nil
	}
	proxy := meta.(*sdk.Client).PublicIPAddresses

	publicIp := mapResourceDataToPublicIp(d)
	id_value, _ := strconv.Atoi(d.Id())
	var id = (int32)(id_value)
	publicIp.AllocationId = &id

	_, err := proxy.Update(ctx, publicIp)
	if err != nil {
		return diag.Errorf("error while modifying floating ip: %+v", err)
	}

	return resourceFloatingIpRead(ctx, d, meta)
}

func resourcePublicIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).PublicIPAddresses

	id_value, _ := strconv.Atoi(d.Id())
	err := proxy.Delete(ctx, (int32)(id_value))
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

func mapResourceDataToPublicIp(d *schema.ResourceData) openapi.PublicIpAllocation {
	PublicIpAddressValue := d.Get("ip_address").(string)
	IpAddressValue := d.Get("public_ip_address_id").(string)
	PublicIpAddress := openapi.PublicIpAddress{}
	PublicIpAddress.IpAddressId = NewNullableString(IpAddressValue)
	PublicIpAddress.IpAddress = NewNullableString(PublicIpAddressValue)

	id_v, _ := strconv.Atoi(d.Get("public_ip_address_id").(string))
	id_value := (int32)(id_v)

	return openapi.PublicIpAllocation{
		AllocationId:    &id_value,
		PublicIpAddress: &PublicIpAddress,
		Name:            NewNullableString(d.Get("display_name").(string)),
	}
}
