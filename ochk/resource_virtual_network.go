package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/go-cty/cty"
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
			"build_in": {
				Type:     schema.TypeBool,
				Computed: true,
				Default:  nil,
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
							ValidateDiagFunc: func(v any, p cty.Path) diag.Diagnostics {
								value := v.(string)
								platformType := sdk.PLATFORM_TYPE
								var diags diag.Diagnostics
								if value != "" && platformType == "VMWARE" {
									diagnostic := diag.Diagnostic{
										Severity: diag.Error,
										Summary:  fmt.Sprintf("Unsupported value for platform type: %s", platformType),
										Detail:   fmt.Sprintf("Value %q is not supported for platform type: %s", p[0], platformType),
									}
									diags = append(diags, diagnostic)
								}
								return diags
							},
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

func mapVirtualNetworkToResourceData(d *schema.ResourceData, virtualNetwork *openapi.VirtualNetworkInstance) diag.Diagnostics {
	if err := d.Set("display_name", virtualNetwork.GetDisplayName()); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("project_id", virtualNetwork.GetProjectId()); err != nil {
		return diag.Errorf("error setting project_id: %+v", err)
	}

	if err := d.Set("build_in", virtualNetwork.GetBuildIn()); err != nil {
		return diag.Errorf("error setting buildIn: %+v", err)
	}

	if err := d.Set("folder_path", virtualNetwork.GetFolderPath()); err != nil {
		return diag.Errorf("error setting folder_path: %+v", err)
	}

	if err := d.Set("ipam_enabled", virtualNetwork.IpamEnabled); err != nil {
		return diag.Errorf("error setting ipam_enabled: %+v", err)
	}

	if err := d.Set("gateway_address", virtualNetwork.GatewayAddress.Get()); err != nil {
		return diag.Errorf("error setting gateway_address: %+v", err)
	}

	if err := d.Set("subnet_mask", virtualNetwork.SubnetMask.Get()); err != nil {
		return diag.Errorf("error setting subnet_mask: %+v", err)
	}

	if err := d.Set("vpc_id", virtualNetwork.GetRouterRefId()); err != nil {
		return diag.Errorf("error setting vpc: %+v", err)
	}

	if virtualNetwork.Subnet != nil {
		if err := d.Set("subnet_gateway_address_cidr", virtualNetwork.Subnet.GatewayAddressCIDR.Get()); err != nil {
			return diag.Errorf("error setting subnet_gateway_address_cidr: %+v", err)
		}
		if err := d.Set("subnet_network_cidr", virtualNetwork.Subnet.NetworkCIDR.Get()); err != nil {
			return diag.Errorf("error setting subnet_network_cidr: %+v", err)
		}
		if virtualNetwork.Subnet.DnsServers != nil && len(virtualNetwork.Subnet.DnsServers) > 0 {
			if err := d.Set("dns_servers", flattenDnsServers(virtualNetwork.Subnet.DnsServers)); err != nil {
				return diag.Errorf("error setting dns_servers: %+v", err)
			}
		}
	}

	return nil
}

func resourceVirtualNetworkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkClient := meta.(*sdk.Client)

	virtualNetwork := mapResourceDataToVirtualNetwork(d)
	virtualNetwork.VirtualNetworkId = NewNullableString(d.Id())

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

func mapResourceDataToVirtualNetwork(d *schema.ResourceData) openapi.VirtualNetworkInstance {
	IpamEnabledValue := d.Get("ipam_enabled").(bool)
	virtualNetworkInstance := openapi.VirtualNetworkInstance{
		DisplayName:      NewNullableString(d.Get("display_name").(string)),
		GatewayAddress:   NewNullableString(d.Get("gateway_address").(string)),
		IpamEnabled:      &IpamEnabledValue,
		RouterRefId:      NewNullableString(d.Get("vpc_id").(string)),
		SubnetMask:       NewNullableString(d.Get("subnet_mask").(string)),
		ProjectId:        NewNullableString(d.Get("project_id").(string)),
		VirtualNetworkId: openapi.NullableString{},
	}

	subnetGatewayAddressCidr, subnetGatewayAddressCidrOk := d.GetOk("subnet_gateway_address_cidr")
	subnetNetworkCidr, subnetNetworkCidrOk := d.GetOk("subnet_network_cidr")
	dnsServers, dnsServersOk := d.GetOk("dns_servers")

	if subnetGatewayAddressCidrOk || subnetNetworkCidrOk {
		virtualNetworkInstance.Subnet = &openapi.SegmentSubnetInstance{}

		if subnetGatewayAddressCidrOk {
			virtualNetworkInstance.Subnet.GatewayAddressCIDR = NewNullableString(subnetGatewayAddressCidr.(string))
		}
		if subnetNetworkCidrOk {
			virtualNetworkInstance.Subnet.NetworkCIDR = NewNullableString(subnetNetworkCidr.(string))
		}

		if dnsServersOk {
			virtualNetworkInstance.Subnet.DnsServers = expandDnsServers(dnsServers.(*schema.Set).List())
		}
	}

	return virtualNetworkInstance
}

func expandDnsServers(in []interface{}) []openapi.DnsServerInstance {
	if len(in) == 0 {
		return nil
	}

	var out = make([]openapi.DnsServerInstance, len(in))
	for i, v := range in {
		m := v.(map[string]interface{})
		member := &openapi.DnsServerInstance{}
		if address, ok := m["address"].(string); ok {
			member.Address = NewNullableString(address)
		}
		out[i] = *member
	}
	return out
}
