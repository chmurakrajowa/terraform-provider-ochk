package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"strings"
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

		Importer: &schema.ResourceImporter{
			StateContext: resourceVirtualNetworkImportState,
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"folder_path": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/",
			},
			"ipam_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
			"subnet_gateway_address_cidr": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_network_cidr": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dns_servers": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:     schema.TypeString,
							Required: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceVirtualNetworkImportState(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	d.SetId(strings.ToLower(d.Id()))
	return []*schema.ResourceData{d}, nil
}

func resourceVirtualNetworkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*sdk.Client)

	virtualNetwork := mapResourceDataToVirtualNetwork(d)

	request, err := client.VirtualNetworks.Create(ctx, virtualNetwork)
	if err != nil {
		return diag.Errorf("error while creating virtual network: %+v", err)
	}

	err, resourceID := client.Requests.FetchResourceID(ctx, d.Timeout(schema.TimeoutCreate), request)
	if err != nil {
		return diag.Errorf("error while fetching virtual network request state: %+v", err)
	}

	d.SetId(resourceID.String())

	return resourceVirtualNetworkRead(ctx, d, meta)
}

func resourceVirtualNetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).VirtualNetworks

	virtualNetwork, err := proxy.Read(ctx, strfmt.UUID(d.Id()))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("virtual network with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading virtual network: %+v", err)
	}

	if err := mapVirtualNetworkToResourceData(d, virtualNetwork); err != nil {
		return err
	}

	return nil
}

func mapVirtualNetworkToResourceData(d *schema.ResourceData, virtualNetwork *models.VirtualNetworkInstance) diag.Diagnostics {
	if err := d.Set("display_name", virtualNetwork.DisplayName); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("project_id", virtualNetwork.ProjectID); err != nil {
		return diag.Errorf("error setting project_id: %+v", err)
	}

	if err := d.Set("folder_path", virtualNetwork.FolderPath); err != nil {
		return diag.Errorf("error setting folder_path: %+v", err)
	}

	if err := d.Set("ipam_enabled", virtualNetwork.IpamEnabled); err != nil {
		return diag.Errorf("error setting ipam_enabled: %+v", err)
	}

	if err := d.Set("gateway_address", virtualNetwork.GatewayAddress); err != nil {
		return diag.Errorf("error setting gateway_address: %+v", err)
	}

	if err := d.Set("subnet_mask", virtualNetwork.SubnetMask); err != nil {
		return diag.Errorf("error setting subnet_mask: %+v", err)
	}

	if err := d.Set("vpc_id", virtualNetwork.RouterRefID); err != nil {
		return diag.Errorf("error setting vpc: %+v", err)
	}

	if virtualNetwork.Subnet != nil {
		if err := d.Set("subnet_gateway_address_cidr", virtualNetwork.Subnet.GatewayAddressCIDR); err != nil {
			return diag.Errorf("error setting subnet_gateway_address_cidr: %+v", err)
		}
		if err := d.Set("subnet_network_cidr", virtualNetwork.Subnet.NetworkCIDR); err != nil {
			return diag.Errorf("error setting subnet_network_cidr: %+v", err)
		}
		if virtualNetwork.Subnet.DNSServers != nil && len(virtualNetwork.Subnet.DNSServers) > 0 {
			if err := d.Set("dns_servers", flattenDnsServers(virtualNetwork.Subnet.DNSServers)); err != nil {
				return diag.Errorf("error setting dns_servers: %+v", err)
			}
		}
	}

	return nil
}

func resourceVirtualNetworkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkClient := meta.(*sdk.Client)

	virtualNetwork := mapResourceDataToVirtualNetwork(d)
	virtualNetwork.VirtualNetworkID = strfmt.UUID(d.Id())

	request, err := sdkClient.VirtualNetworks.Update(ctx, virtualNetwork)
	if err != nil {
		return diag.Errorf("error while modifying virtual network: %+v", err)
	}

	err, resourceID := sdkClient.Requests.FetchResourceID(ctx, d.Timeout(schema.TimeoutCreate), request)
	if err != nil {
		return diag.Errorf("error while fetching virtual network request state: %+v", err)
	}

	d.SetId(resourceID.String())

	return resourceVirtualNetworkRead(ctx, d, meta)
}

func resourceVirtualNetworkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkClient := meta.(*sdk.Client)

	request, err := sdkClient.VirtualNetworks.Delete(ctx, strfmt.UUID(d.Id()))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("virtual network with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting virtual network: %+v", err)
	}

	err, _ = sdkClient.Requests.FetchResourceID(ctx, d.Timeout(schema.TimeoutCreate), request)
	if err != nil {
		return diag.Errorf("error while fetching virtual network request state: %+v", err)
	}

	return nil
}

func mapResourceDataToVirtualNetwork(d *schema.ResourceData) *models.VirtualNetworkInstance {
	virtualNetworkInstance := models.VirtualNetworkInstance{
		DisplayName:      d.Get("display_name").(string),
		GatewayAddress:   d.Get("gateway_address").(string),
		IpamEnabled:      d.Get("ipam_enabled").(bool),
		RouterRefID:      strfmt.UUID(d.Get("vpc_id").(string)),
		SubnetMask:       d.Get("subnet_mask").(string),
		ProjectID:        strfmt.UUID(d.Get("project_id").(string)),
		VirtualNetworkID: strfmt.UUID(d.Id()),
	}

	subnetGatewayAddressCidr, subnetGatewayAddressCidrOk := d.GetOk("subnet_gateway_address_cidr")
	subnetNetworkCidr, subnetNetworkCidrOk := d.GetOk("subnet_network_cidr")
	dnsServers, dnsServersOk := d.GetOk("dns_servers")

	if subnetGatewayAddressCidrOk || subnetNetworkCidrOk {
		virtualNetworkInstance.Subnet = &models.SegmentSubnetInstance{}

		if subnetGatewayAddressCidrOk {
			virtualNetworkInstance.Subnet.GatewayAddressCIDR = subnetGatewayAddressCidr.(string)
		}
		if subnetNetworkCidrOk {
			virtualNetworkInstance.Subnet.NetworkCIDR = subnetNetworkCidr.(string)
		}

		if dnsServersOk {
			virtualNetworkInstance.Subnet.DNSServers = expandDnsServers(dnsServers.(*schema.Set).List())
		}
	}

	return &virtualNetworkInstance
}

func expandDnsServers(in []interface{}) []*models.DNSServerInstance {
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.DNSServerInstance, len(in))
	for i, v := range in {
		m := v.(map[string]interface{})
		member := &models.DNSServerInstance{}
		if address, ok := m["address"].(string); ok {
			member.Address = address
		}
		out[i] = member
	}
	return out
}
