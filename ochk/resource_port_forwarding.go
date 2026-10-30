package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	PortForwardingRetryTimeout = 1 * time.Minute
)

func resourcePortForwarding() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePortForwardingCreate,
		ReadContext:   resourcePortForwardingRead,
		UpdateContext: resourcePortForwardingUpdate,
		DeleteContext: resourcePortForwardingDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(PortForwardingRetryTimeout),
			Update: schema.DefaultTimeout(PortForwardingRetryTimeout),
			Delete: schema.DefaultTimeout(PortForwardingRetryTimeout),
		},

		Importer: &schema.ResourceImporter{
			StateContext: portForwardingStateContextImport,
		},

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
				Optional: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"internal_port_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"internal_ip_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"internal_port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"internal_port_range": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"external_port_range": {
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

func portForwardingStateContextImport(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	d.SetId(strings.ToLower(d.Id()))
	return []*schema.ResourceData{d}, nil
}

func resourcePortForwardingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).PortForwarding

	floatingIpId := strfmt.UUID(d.Get("floating_ip_id").(string))

	portForwarding := mapResourceDataToPortForwarding(d)

	created, err := proxy.Create(ctx, floatingIpId, portForwarding)
	if err != nil {
		return diag.Errorf("error while creating port forwarding: %+v", err)
	}

	d.SetId(created.GetPortForwardingId())

	return resourcePortForwardingRead(ctx, d, meta)
}

func resourcePortForwardingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).PortForwarding
	floatingIpId := strfmt.UUID(d.Get("floating_ip_id").(string))
	portForwardingId := strfmt.UUID(d.Id())

	portForwarding, err := proxy.Read(ctx, floatingIpId, portForwardingId)
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("port forwarding with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading port forwarding: %+v", err)
	}

	if err := d.Set("display_name", portForwarding.GetName()); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("floating_ip_id", floatingIpId); err != nil {
		return diag.Errorf("error setting floating_ip_id: %+v", err)
	}

	if err := d.Set("port_forwarding_id", portForwarding.GetPortForwardingId()); err != nil {
		return diag.Errorf("error setting port_forwarding_id: %+v", err)
	}

	if err := d.Set("description", portForwarding.GetDescription()); err != nil {
		return diag.Errorf("error setting description: %+v", err)
	}

	if err := d.Set("protocol", portForwarding.GetProtocol()); err != nil {
		return diag.Errorf("error setting protocol: %+v", err)
	}

	if err := d.Set("internal_port_id", portForwarding.GetInternalPortId()); err != nil {
		return diag.Errorf("error setting internal_port_id: %+v", err)
	}

	if err := d.Set("internal_ip_address", portForwarding.GetInternalIpAddress()); err != nil {
		return diag.Errorf("error setting internal_ip_address: %+v", err)
	}

	if err := d.Set("internal_port", portForwarding.GetInternalPort()); err != nil {
		return diag.Errorf("error setting internal_port: %+v", err)
	}

	if err := d.Set("internal_port_range", portForwarding.GetInternalPortRange()); err != nil {
		return diag.Errorf("error setting internal_port_range: %+v", err)
	}

	if err := d.Set("public_address", portForwarding.GetPublicAddress()); err != nil {
		return diag.Errorf("error setting public_address: %+v", err)
	}

	if err := d.Set("external_port", portForwarding.GetExternalPort()); err != nil {
		return diag.Errorf("error setting external_port: %+v", err)
	}

	if err := d.Set("external_port_range", portForwarding.GetExternalPortRange()); err != nil {
		return diag.Errorf("error setting external_port_range: %+v", err)
	}

	if err := d.Set("created_by", portForwarding.GetCreatedBy()); err != nil {
		return diag.Errorf("error setting created_by: %+v", err)
	}

	if err := d.Set("created_at", portForwarding.GetCreationDate()); err != nil {
		return diag.Errorf("error setting created_at: %+v", err)
	}

	if err := d.Set("modified_by", portForwarding.GetModifiedBy()); err != nil {
		return diag.Errorf("error setting modified_by: %+v", err)
	}

	if err := d.Set("modified_at", portForwarding.GetModificationDate()); err != nil {
		return diag.Errorf("error setting modified_at: %+v", err)
	}

	return nil
}

func resourcePortForwardingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if !d.HasChanges("display_name", "description", "protocol", "internal_port_id", "internal_ip_address", "internal_port", "internal_port_range", "public_address", "external_port", "external_port_range") {
		return nil
	}
	proxy := meta.(*sdk.Client).PortForwarding

	floatingIpId := strfmt.UUID(d.Get("floating_ip_id").(string))
	portForwarding := mapResourceDataToPortForwarding(d)

	idValue := d.Id()
	portForwarding.PortForwardingId = &idValue

	_, err := proxy.Update(ctx, floatingIpId, portForwarding)
	if err != nil {
		return diag.Errorf("error while creating port forwarding: %+v", err)
	}

	return resourcePortForwardingRead(ctx, d, meta)
}

func mapResourceDataToPortForwarding(d *schema.ResourceData) openapi.PortForwarding {
	internalPortIdValue := d.Get("internal_port_id").(string)
	internalPort := d.Get("internal_port").(int)
	internalPortValue := int32(internalPort)

	externalPort := d.Get("external_port").(int)
	externalPortValue := int32(externalPort)
	portForwarding := openapi.PortForwarding{}
	if internalPortValue == 0 || externalPortValue == 0 {
		portForwarding = openapi.PortForwarding{
			Name:              NewNullableString(d.Get("display_name").(string)),
			Description:       NewNullableString(d.Get("description").(string)),
			Protocol:          NewNullableString(d.Get("protocol").(string)),
			InternalPortId:    &internalPortIdValue,
			InternalIpAddress: NewNullableString(d.Get("internal_ip_address").(string)),
			InternalPortRange: NewNullableString(d.Get("internal_port_range").(string)),
			PublicAddress:     NewNullableString(d.Get("public_address").(string)),
			ExternalPortRange: NewNullableString(d.Get("external_port_range").(string)),
		}
	} else {
		portForwarding = openapi.PortForwarding{
			Name:              NewNullableString(d.Get("display_name").(string)),
			Description:       NewNullableString(d.Get("description").(string)),
			Protocol:          NewNullableString(d.Get("protocol").(string)),
			InternalPortId:    &internalPortIdValue,
			InternalIpAddress: NewNullableString(d.Get("internal_ip_address").(string)),
			InternalPort:      NewNullableInt32(internalPortValue),
			InternalPortRange: NewNullableString(d.Get("internal_port_range").(string)),
			PublicAddress:     NewNullableString(d.Get("public_address").(string)),
			ExternalPort:      NewNullableInt32(externalPortValue),
			ExternalPortRange: NewNullableString(d.Get("external_port_range").(string)),
		}
	}
	return portForwarding
}

func resourcePortForwardingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).PortForwarding

	floatingIpID := strfmt.UUID(d.Get("floating_ip_id").(string))

	err := proxy.Delete(ctx, floatingIpID, strfmt.UUID(d.Id()))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("port forwarding with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting port forwarding: %+v", err)
	}

	return nil
}
