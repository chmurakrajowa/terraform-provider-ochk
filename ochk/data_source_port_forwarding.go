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
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
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

	if err := d.Set("display_name", portForwarding[0].Name); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("description", portForwarding[0].Description); err != nil {
		return diag.Errorf("error setting description: %+v", err)
	}

	if err := d.Set("port_forwarding_id", portForwarding[0].PortForwardingID); err != nil {
		return diag.Errorf("error setting port_forwarding_id: %+v", err)
	}

	if err := d.Set("public_address", portForwarding[0].PublicAddress); err != nil {
		return diag.Errorf("error setting public_address: %+v", err)
	}

	if err := d.Set("external_port", portForwarding[0].ExternalPort); err != nil {
		return diag.Errorf("error setting external_port: %+v", err)
	}

	if err := d.Set("external_port_range", portForwarding[0].ExternalPortRange); err != nil {
		return diag.Errorf("error setting external_port_range: %+v", err)
	}

	if err := d.Set("internal_port", portForwarding[0].InternalPort); err != nil {
		return diag.Errorf("error setting internal_port: %+v", err)
	}

	if err := d.Set("internal_ip_address", portForwarding[0].InternalIPAddress); err != nil {
		return diag.Errorf("error setting internal_ip_address: %+v", err)
	}

	if err := d.Set("internal_port_id", portForwarding[0].InternalPortID); err != nil {
		return diag.Errorf("error setting internal_port_id: %+v", err)
	}

	if err := d.Set("internal_port_range", portForwarding[0].InternalPortRange); err != nil {
		return diag.Errorf("error setting internal_port_range: %+v", err)
	}

	if err := d.Set("protocol", portForwarding[0].Protocol); err != nil {
		return diag.Errorf("error setting protocol: %+v", err)
	}

	if err := d.Set("created_by", portForwarding[0].CreatedBy); err != nil {
		return diag.Errorf("error setting created_by: %+v", err)
	}

	if err := d.Set("created_at", portForwarding[0].CreationDate.String()); err != nil {
		return diag.Errorf("error setting created_at: %+v", err)
	}

	if err := d.Set("modified_by", portForwarding[0].ModifiedBy); err != nil {
		return diag.Errorf("error setting modified_by: %+v", err)
	}

	if err := d.Set("modified_at", portForwarding[0].ModificationDate.String()); err != nil {
		return diag.Errorf("error setting modified_at: %+v", err)
	}

	d.SetId(portForwarding[0].PortForwardingID.String())
	return nil
}
