package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePortForwarding() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePortForwardingRead,

		Schema: map[string]*schema.Schema{
			"floating_ip_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port_forwarding_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"external_port_range": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"internal_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"internal_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"internal_port_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"internal_port_range": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"protocol": {
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

func dataSourcePortForwardingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).PortForwarding
	floatingIpId := strfmt.UUID(d.Get("floating_ip_id").(string))
	name := d.Get("display_name").(string)
	portForwarding, err := proxy.ListByName(ctx, floatingIpId, name)

	if err != nil {
		return diag.Errorf("error while listing ports forwarding: %+v", err)
	}

	if len(portForwarding) < 1 {
		return diag.Errorf("no port forwarding found for display_name: %s", name)
	}

	if len(portForwarding) > 1 {
		return diag.Errorf("more than one firewall ew rule with display_name: %s found!", name)
	}

	if err := d.Set("display_name", portForwarding[0].GetName()); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("description", portForwarding[0].GetDescription()); err != nil {
		return diag.Errorf("error setting description: %+v", err)
	}

	if err := d.Set("port_forwarding_id", portForwarding[0].GetPortForwardingId()); err != nil {
		return diag.Errorf("error setting port_forwarding_id: %+v", err)
	}

	if err := d.Set("public_address", portForwarding[0].GetPublicAddress()); err != nil {
		return diag.Errorf("error setting public_address: %+v", err)
	}

	if err := d.Set("external_port", portForwarding[0].GetExternalPort()); err != nil {
		return diag.Errorf("error setting external_port: %+v", err)
	}

	if err := d.Set("external_port_range", portForwarding[0].GetExternalPortRange()); err != nil {
		return diag.Errorf("error setting external_port_range: %+v", err)
	}

	if err := d.Set("internal_port", portForwarding[0].GetInternalPort()); err != nil {
		return diag.Errorf("error setting internal_port: %+v", err)
	}

	if err := d.Set("internal_ip_address", portForwarding[0].GetInternalIpAddress()); err != nil {
		return diag.Errorf("error setting internal_ip_address: %+v", err)
	}

	if err := d.Set("internal_port_id", portForwarding[0].GetInternalPortId()); err != nil {
		return diag.Errorf("error setting internal_port_id: %+v", err)
	}

	if err := d.Set("internal_port_range", portForwarding[0].GetInternalPortRange()); err != nil {
		return diag.Errorf("error setting internal_port_range: %+v", err)
	}

	if err := d.Set("protocol", portForwarding[0].GetProtocol()); err != nil {
		return diag.Errorf("error setting protocol: %+v", err)
	}

	if err := d.Set("created_by", portForwarding[0].GetCreatedBy()); err != nil {
		return diag.Errorf("error setting created_by: %+v", err)
	}

	if err := d.Set("created_at", portForwarding[0].GetCreationDate()); err != nil {
		return diag.Errorf("error setting created_at: %+v", err)
	}

	if err := d.Set("modified_by", portForwarding[0].GetModifiedBy()); err != nil {
		return diag.Errorf("error setting modified_by: %+v", err)
	}

	if err := d.Set("modified_at", portForwarding[0].GetModificationDate()); err != nil {
		return diag.Errorf("error setting modified_at: %+v", err)
	}

	d.SetId(portForwarding[0].GetPortForwardingId())
	return nil
}
