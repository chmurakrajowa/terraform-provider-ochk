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
	VirtualNetworkRetryTimeout = 5 * time.Minute
)

func resourceVirtualNetwork() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVirtualNetworkCreate,
		ReadContext:   resourceVirtualNetworkRead,
		UpdateContext: resourceVirtualNetworkUpdate,
		DeleteContext: resourceVirtualNetworkDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(VirtualNetworkRetryTimeout),
			Update: schema.DefaultTimeout(VirtualNetworkRetryTimeout),
			Delete: schema.DefaultTimeout(VirtualNetworkRetryTimeout),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ipam_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceVirtualNetworkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*sdk.Client)

	virtualNetwork := mapResourceDataToVirtualNetwork(d)

	request, err := client.VirtualNetworks.Create(ctx, virtualNetwork, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.Errorf("error while creating virtual network: %+v", err)
	}

	err, resourceID := client.Requests.FetchResourceID(ctx, d.Timeout(schema.TimeoutCreate), request)
	if err != nil {
		return diag.Errorf("error while fetching virtual network request state: %+v", err)
	}

	d.SetId(resourceID)

	return resourceVirtualNetworkRead(ctx, d, meta)
}

func resourceVirtualNetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).VirtualNetworks

	virtualNetwork, err := proxy.Read(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			d.SetId("")
			return diag.Errorf("virtual network with id %s not found: %+v", d.Id(), err)
		}

		return diag.Errorf("error while reading virtual network: %+v", err)
	}

	if err := d.Set("display_name", virtualNetwork.DisplayName); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("ipam_enabled", virtualNetwork.IpamEnabled); err != nil {
		return diag.Errorf("error setting ipam_enabled: %+v", err)
	}

	return nil
}

func resourceVirtualNetworkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).VirtualNetworks

	virtualNetwork := mapResourceDataToVirtualNetwork(d)
	virtualNetwork.VirtualNetworkID = d.Id()

	_, err := proxy.Update(ctx, virtualNetwork)
	if err != nil {
		return diag.Errorf("error while modifying virtual network: %+v", err)
	}

	return resourceVirtualNetworkRead(ctx, d, meta)
}

func resourceVirtualNetworkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).VirtualNetworks

	err := proxy.Delete(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			d.SetId("")
			return diag.Errorf("virtual network with id %s not found: %+v", d.Id(), err)
		}

		return diag.Errorf("error while deleting virtual network: %+v", err)
	}

	return nil
}

func mapResourceDataToVirtualNetwork(d *schema.ResourceData) *models.VirtualNetworkInstance {
	virtualNetworkInstance := models.VirtualNetworkInstance{
		DisplayName: d.Get("display_name").(string),
	}

	if ipamEnabled, ok := d.GetOk("ipam_enabled"); ok && ipamEnabled.(bool) {
		virtualNetworkInstance.IpamEnabled = true
	}

	return &virtualNetworkInstance
}
