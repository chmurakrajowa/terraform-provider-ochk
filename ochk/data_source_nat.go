package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNat() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNatRead,
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrf_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nat_type": {
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
			"action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"service_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"translated_ports": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_network": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"destination_network": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"translated_network": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_network_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func dataSourceNatRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client)

	displayName := d.Get("display_name").(string)

	nats, err := proxy.Nats.ListNatsByName(ctx, displayName)

	if err != nil {
		return diag.Errorf("error while listing nats: %+v", err)
	}

	if len(nats) < 1 {
		return diag.Errorf("no nats found for display name: %s", displayName)
	}

	if len(nats) > 1 {
		return diag.Errorf("more than one nat with display name: %s found!", displayName)
	}

	if err := d.Set("display_name", displayName); err != nil {
		return diag.Errorf("error setting nat display name: %+v", err)
	}

	if err := d.Set("description", nats[0].Description); err != nil {
		return diag.Errorf("error setting description:  %+v", err)
	}

	if err := d.Set("vrf_id", nats[0].TierZeroRouter.RouterID); err != nil {
		return diag.Errorf("error setting vrf_id:  %+v", err)
	}

	if err := d.Set("nat_type", nats[0].NatType); err != nil {
		return diag.Errorf("error setting nat_type:  %+v", err)
	}

	if err := d.Set("created_by", nats[0].CreatedBy); err != nil {
		return diag.Errorf("error setting created_by:  %+v", err)
	}

	if err := d.Set("created_at", nats[0].CreationDate.String()); err != nil {
		return diag.Errorf("error setting created_at:  %+v", err)
	}

	if err := d.Set("modified_at", nats[0].ModificationDate.String()); err != nil {
		return diag.Errorf("error setting modified_at:  %+v", err)
	}

	if err := d.Set("modified_by", nats[0].ModifiedBy); err != nil {
		return diag.Errorf("error setting modified_by:  %+v", err)
	}

	if err := d.Set("action", nats[0].Action); err != nil {
		return diag.Errorf("error setting action:  %+v", err)
	}

	if err := d.Set("enabled", nats[0].Enabled); err != nil {
		return diag.Errorf("error setting enabled: %+v", err)
	}

	if err := d.Set("priority", nats[0].Priority); err != nil {
		return diag.Errorf("error setting priority: %+v", err)
	}

	if err := d.Set("source_network", nats[0].SourceNetwork); err != nil {
		return diag.Errorf("error setting source_network: %+v", err)
	}

	if err := d.Set("destination_network", nats[0].DestinationNetwork); err != nil {
		return diag.Errorf("error setting destination_network: %+v", err)
	}

	if nats[0].NatType == "MANUAL" {
		if nats[0].Action == "DNAT" {
			if nats[0].ServiceInstance != nil && nats[0].TranslatedPorts != "" {
				if err := d.Set("service_id", nats[0].ServiceInstance.ServiceID); err != nil {
					return diag.Errorf("error setting service: %+v", err)
				}
				if err := d.Set("translated_ports", nats[0].TranslatedPorts); err != nil {
					return diag.Errorf("error setting translated_ports: %+v", err)
				}
			}
		}
		if nats[0].Action == "DNAT" || nats[0].Action == "SNAT" {
			if err := d.Set("translated_network", nats[0].TranslatedNetwork); err != nil {
				return diag.Errorf("error setting translated_network: %+v", err)
			}
		}
	}

	if nats[0].NatType == "AUTO" {
		if err := d.Set("virtual_network_id", nats[0].VirtualNetworkInstance.VirtualNetworkID); err != nil {
			return diag.Errorf("error setting virtual_network_id: %+v", err)
		}
	}

	d.SetId(nats[0].RuleID)

	return nil
}
