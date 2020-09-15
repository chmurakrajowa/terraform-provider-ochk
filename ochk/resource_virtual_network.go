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
	VirtualNetworkRetryTimeout = 20 * time.Minute
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
				ForceNew: true,
			},
			"primary_dns_address": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"secondary_dns_address": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"dns_suffix": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"dns_search_suffix": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"primary_wins_address": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"secondary_wins_address": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"gateway_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_mask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_dhcp_ranges": {
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 0,
				Elem:     &schema.Schema{Type: schema.TypeString},
				ForceNew: true,
			},
			"subnet_gateway_address_cidr": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_network_cidr": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"router": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subtenants": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				MinItems: 1,
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
			id := d.Id()
			d.SetId("")
			return diag.Errorf("virtual network with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading virtual network: %+v", err)
	}

	if err := d.Set("display_name", virtualNetwork.DisplayName); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("ipam_enabled", virtualNetwork.IpamEnabled); err != nil {
		return diag.Errorf("error setting ipam_enabled: %+v", err)
	}
	if err := d.Set("primary_dns_address", virtualNetwork.PrimaryDNSAddress); err != nil {
		return diag.Errorf("error setting primary_dns_address: %+v", err)
	}
	if err := d.Set("secondary_dns_address", virtualNetwork.SecondaryDNSAddress); err != nil {
		return diag.Errorf("error setting secondary_dns_address: %+v", err)
	}
	if err := d.Set("dns_suffix", virtualNetwork.DNSSuffix); err != nil {
		return diag.Errorf("error setting dns_suffix: %+v", err)
	}
	if err := d.Set("dns_search_suffix", virtualNetwork.DNSSearchSuffix); err != nil {
		return diag.Errorf("error setting dns_search_suffix: %+v", err)
	}
	if err := d.Set("primary_wins_address", virtualNetwork.PrimaryWinsAddress); err != nil {
		return diag.Errorf("error setting primary_wins_address: %+v", err)
	}
	if err := d.Set("secondary_wins_address", virtualNetwork.SecondaryWinsAddress); err != nil {
		return diag.Errorf("error setting secondary_wins_address: %+v", err)
	}
	if err := d.Set("gateway_address", virtualNetwork.GatewayAddress); err != nil {
		return diag.Errorf("error setting gateway_address: %+v", err)
	}
	if err := d.Set("subnet_mask", virtualNetwork.SubnetMask); err != nil {
		return diag.Errorf("error setting subnet_mask: %+v", err)
	}
	if err := d.Set("router", virtualNetwork.RouterRefID); err != nil {
		return diag.Errorf("error setting router: %+v", err)
	}
	if virtualNetwork.Subnet != nil {
		if err := d.Set("subnet_dhcp_ranges", flattenStringSlice(virtualNetwork.Subnet.DhcpRanges)); err != nil {
			return diag.Errorf("error setting subnet_dhcp_ranges: %+v", err)
		}
		if err := d.Set("subnet_gateway_address_cidr", virtualNetwork.Subnet.GatewayAddressCIDR); err != nil {
			return diag.Errorf("error setting subnet_gateway_address_cidr: %+v", err)
		}
		if err := d.Set("subnet_network_cidr", virtualNetwork.Subnet.NetworkCIDR); err != nil {
			return diag.Errorf("error setting subnet_network_cidr: %+v", err)
		}
	}

	if err := d.Set("subtenants", flattenUUIDs(virtualNetwork.SubtenantRefIds)); err != nil {
		return diag.Errorf("error setting subtenants: %+v", err)
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
			id := d.Id()
			d.SetId("")
			return diag.Errorf("virtual network with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting virtual network: %+v", err)
	}

	return nil
}

func mapResourceDataToVirtualNetwork(d *schema.ResourceData) *models.VirtualNetworkInstance {
	virtualNetworkInstance := models.VirtualNetworkInstance{
		DisplayName:          d.Get("display_name").(string),
		DNSSearchSuffix:      d.Get("dns_search_suffix").(string),
		DNSSuffix:            d.Get("dns_suffix").(string),
		GatewayAddress:       d.Get("gateway_address").(string),
		IpamEnabled:          d.Get("ipam_enabled").(bool),
		PrimaryDNSAddress:    d.Get("primary_dns_address").(string),
		PrimaryWinsAddress:   d.Get("primary_wins_address").(string),
		RouterRefID:          d.Get("router").(string),
		SecondaryDNSAddress:  d.Get("secondary_dns_address").(string),
		SecondaryWinsAddress: d.Get("secondary_wins_address").(string),
		SubnetMask:           d.Get("subnet_mask").(string),
		//FIXME should not be UUID
		SubtenantRefIds:  mapInterfaceSliceToUUIDSlice(d.Get("subtenants").(*schema.Set).List()),
		VirtualNetworkID: d.Id(),
	}

	subnetDhcpRanges, subnetDhcpRangesOk := d.GetOk("subnet_dhcp_ranges")
	subnetGatewayAddressCidr, subnetGatewayAddressCidrOk := d.GetOk("subnet_gateway_address_cidr")
	subnetNetworkCidr, subnetNetworkCidrOk := d.GetOk("subnet_network_cidr")
	if subnetDhcpRangesOk || subnetGatewayAddressCidrOk || subnetNetworkCidrOk {
		virtualNetworkInstance.Subnet = &models.SegmentSubnetInstance{}

		if subnetDhcpRangesOk {
			virtualNetworkInstance.Subnet.DhcpRanges = transformSetToStringSlice(subnetDhcpRanges.(*schema.Set))
		}
		if subnetGatewayAddressCidrOk {
			virtualNetworkInstance.Subnet.GatewayAddressCIDR = subnetGatewayAddressCidr.(string)
		}
		if subnetNetworkCidrOk {
			virtualNetworkInstance.Subnet.NetworkCIDR = subnetNetworkCidr.(string)
		}
	}

	return &virtualNetworkInstance
}
