package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
	"time"
)

const (
	autoNatRetryTimeout = 1 * time.Minute
)

func resourceAutoNat() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAutoNatCreate,
		ReadContext:   resourceNatRead,
		UpdateContext: resourceAutoNatUpdate,
		DeleteContext: resourceNatDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(autoNatRetryTimeout),
			Update: schema.DefaultTimeout(autoNatRetryTimeout),
			Delete: schema.DefaultTimeout(autoNatRetryTimeout),
		},
		Importer: &schema.ResourceImporter{
			StateContext: resourceNatImportState,
		},
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"virtual_network_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"priority": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"service_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_network": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"translated_network": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"destination_network": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceManualNat() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceManualNatCreate,
		ReadContext:   resourceNatRead,
		UpdateContext: resourceManualNatUpdate,
		DeleteContext: resourceNatDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(autoNatRetryTimeout),
			Update: schema.DefaultTimeout(autoNatRetryTimeout),
			Delete: schema.DefaultTimeout(autoNatRetryTimeout),
		},
		Importer: &schema.ResourceImporter{
			StateContext: resourceNatImportState,
		},
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"vrf_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"nat_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"service_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"destination_network": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_network": {
				Type:     schema.TypeString,
				Required: true,
			},
			"translated_network": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"translated_ports": {
				Type:     schema.TypeString,
				Optional: true,
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

func resourceNatImportState(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	d.SetId(strings.ToLower(d.Id()))
	return []*schema.ResourceData{d}, nil
}

func resourceNatRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	proxy := meta.(*sdk.Client).Nats

	Nat, err := proxy.Read(ctx, strfmt.UUID(d.Id()))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("nat with id %s not found: %+v", id, err)
		}
		return diag.Errorf("error while reading nat: %+v", err)
	}

	if err := d.Set("display_name", Nat.DisplayName); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("vrf_id", Nat.TierZeroRouter.RouterId); err != nil {
		return diag.Errorf("error setting vrf_id: %+v", err)
	}

	if err := d.Set("nat_type", Nat.NatType); err != nil {
		return diag.Errorf("error setting nat_type: %+v", err)
	}

	if err := d.Set("created_by", Nat.CreatedBy); err != nil {
		return diag.Errorf("error setting created_by: %+v", err)
	}

	if err := d.Set("created_at", Nat.GetCreationDate()); err != nil {
		return diag.Errorf("error setting created_at: %+v", err)
	}

	if err := d.Set("modified_by", Nat.ModifiedBy); err != nil {
		return diag.Errorf("error setting modified_by: %+v", err)
	}

	if err := d.Set("modified_at", Nat.GetModificationDate()); err != nil {
		return diag.Errorf("error setting modified_at: %+v", err)
	}

	if err := d.Set("action", Nat.Action); err != nil {
		return diag.Errorf("error setting action: %+v", err)
	}

	if err := d.Set("priority", Nat.Priority); err != nil {
		return diag.Errorf("error setting priority: %+v", err)
	}

	if *Nat.NatType == "AUTO" {
		if err := d.Set("virtual_network_id", Nat.VirtualNetworkInstance.VirtualNetworkId); err != nil {
			return diag.Errorf("error setting virtual_network_id: %+v", err)
		}
	}

	if *Nat.NatType == "MANUAL" {
		if Nat.Description == "UNKNOWN" || Nat.Description == "" {
			if err := d.Set("description", ""); err != nil {
				return diag.Errorf("error setting description: %+v", err)
			}
		} else {
			if err := d.Set("description", Nat.Description); err != nil {
				return diag.Errorf("error setting description: %+v", err)
			}
		}

		if err := d.Set("enabled", Nat.Enabled); err != nil {
			return diag.Errorf("error setting enabled: %+v", err)
		}

		if err := d.Set("source_network", Nat.SourceNetwork); err != nil {
			return diag.Errorf("error setting source_network: %+v", err)
		}

		if err := d.Set("destination_network", Nat.DestinationNetwork); err != nil {
			return diag.Errorf("error setting destination_network: %+v", err)
		}

		if Nat.Action == "DNAT" || Nat.Action == "SNAT" {
			if err := d.Set("translated_network", Nat.TranslatedNetwork); err != nil {
				return diag.Errorf("error setting translated_network: %+v", err)
			}
		}

		if Nat.Action == "DNAT" {
			if Nat.TranslatedPorts != "" && Nat.ServiceInstance.ServiceId != "" {
				if err := d.Set("service_id", Nat.ServiceInstance.ServiceId); err != nil {
					return diag.Errorf("error setting service_id: %+v", err)
				}

				if err := d.Set("translated_ports", Nat.TranslatedPorts); err != nil {
					return diag.Errorf("error setting translated_ports: %+v", err)
				}
			}
		}
	}
	return nil
}

func resourceAutoNatCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*sdk.Client)

	nat := mapResourceDataToAutoNat(d)

	request, err := client.Nats.CreateNat(ctx, nat)
	if err != nil {
		return diag.Errorf("error while creating nat: %+v", err)
	}

	err, resourceID := client.Requests.FetchResourceID(ctx, d.Timeout(schema.TimeoutCreate), request)
	if err != nil {
		return diag.Errorf("error while fetching nat request state: %+v", err)
	}

	d.SetId(resourceID.String())

	return resourceNatRead(ctx, d, meta)

}

func resourceManualNatCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*sdk.Client)

	nat := mapResourceDataToManualNat(d)

	request, err := client.Nats.CreateNat(ctx, nat)
	if err != nil {
		return diag.Errorf("error while creating nat: %+v", err)
	}

	err, resourceID := client.Requests.FetchResourceID(ctx, d.Timeout(schema.TimeoutCreate), request)
	if err != nil {
		return diag.Errorf("error while fetching nat request state: %+v", err)
	}

	d.SetId(resourceID.String())

	return resourceNatRead(ctx, d, meta)
}

func resourceAutoNatUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Nats

	Nat := mapResourceDataToAutoNat(d)
	Nat.RuleId = NewNullableString(d.Id())

	_, err := proxy.Update(ctx, &Nat)
	if err != nil {
		return diag.Errorf("error while modifying nat: %+v", err)
	}

	return resourceNatRead(ctx, d, meta)
}

func resourceManualNatUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Nats

	Nat := mapResourceDataToManualNat(d)
	Nat.RuleId = NewNullableString(d.Id())

	_, err := proxy.Update(ctx, &Nat)
	if err != nil {
		return diag.Errorf("error while modifying nat: %+v", err)
	}

	return resourceNatRead(ctx, d, meta)
}

func resourceNatDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Nats

	err := proxy.Delete(ctx, strfmt.UUID(d.Id()))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("nat with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting nat: %+v", err)
	}

	return nil
}

func mapResourceDataToAutoNat(d *schema.ResourceData) openapi.NATRuleInstance {
	return openapi.NATRuleInstance{
		DisplayName:            NewNullableString(d.Get("display_name").(string)),
		VirtualNetworkInstance: mapResourceDataToVirtualNetworkInstance(d),
		NatType:                "AUTO",
	}
}

func mapResourceDataToManualNat(d *schema.ResourceData) openapi.NATRuleInstance {

	publicPriorityInt64 := int64(d.Get("priority").(int))

	natManualRule := openapi.NATRuleInstance{
		DisplayName:        NewNullableString(d.Get("display_name").(string)),
		Description:        NewNullableString(d.Get("description").(string)),
		Enabled:            NewNullableBool(d.Get("enabled").(bool)),
		TierZeroRouter:     mapResourceDataToVrfRouter(d),
		Action:             openapi.NATRuleAction(d.Get("action").(string)),
		Priority:           NewNullableInt64(publicPriorityInt64),
		SourceNetwork:      NewNullableString(d.Get("source_network").(string)),
		DestinationNetwork: NewNullableString(d.Get("destination_network").(string)),
		NatType:            "MANUAL",
	}
	natAction := d.Get("action")

	if natAction == "DNAT" {
		if d.Get("translated_ports") != "" && d.Get("service_id") != "" {
			natManualRule.ServiceInstance = mapResourceDataToServiceInstance(d)
			natManualRule.TranslatedPorts = NewNullableString(d.Get("translated_ports").(string))
		}
	}

	if natAction == "DNAT" || natAction == "SNAT" {
		natManualRule.TranslatedNetwork = NewNullableString(d.Get("translated_network").(string))
	}

	return natManualRule
}

func mapResourceDataToVrfRouter(d *schema.ResourceData) openapi.RouterInstance {
	var vrf_id = d.Get("vrf_id").(string)
	routerInstance := openapi.RouterInstance{
		RouterId: NewNullableString(vrf_id),
	}
	return routerInstance
}

func mapResourceDataToVirtualNetworkInstance(d *schema.ResourceData) openapi.VirtualNetworkInstance {
	virtualNetworkInstance := openapi.VirtualNetworkInstance{
		VirtualNetworkId: NewNullableString(d.Get("virtual_network_id").(string)),
	}
	return virtualNetworkInstance
}

func mapResourceDataToServiceInstance(d *schema.ResourceData) openapi.ServiceInstance {
	serviceInstance := openapi.ServiceInstance{
		ServiceId: NewNullableString(d.Get("service_id").(string)),
	}
	return serviceInstance
}
