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
	FloatingIpRetryTimeout = 5 * time.Minute
)

func resourceFloatingIp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFloatingIpCreate,
		ReadContext:   resourceFloatingIpRead,
		UpdateContext: resourceFloatingIpUpdate,
		DeleteContext: resourceFloatingIpDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(FloatingIpRetryTimeout),
			Update: schema.DefaultTimeout(FloatingIpRetryTimeout),
			Delete: schema.DefaultTimeout(FloatingIpRetryTimeout),
		},

		Importer: &schema.ResourceImporter{
			StateContext: resourceFloatingIpImportState,
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_port_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_fixed_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_adress": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceFloatingIpImportState(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	d.SetId(strings.ToLower(d.Id()))
	return []*schema.ResourceData{d}, nil
}

func resourceFloatingIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FloatingIPAddresses

	floating_ip := mapResourceDataToFloatingIp(d)

	created, err := proxy.Create(ctx, floating_ip)
	if err != nil {
		return diag.Errorf("error while allocation floating ip: %+v", err)
	}

	d.SetId(created.FloatingIPID.String())

	return resourceFloatingIpRead(ctx, d, meta)
}

func resourceFloatingIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FloatingIPAddresses
	floating_ip_id := strfmt.UUID(d.Id())

	floating_ip, err := proxy.Read(ctx, floating_ip_id)
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("floating ip with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading floating ip: %+v", err)
	}

	if err := d.Set("display_name", floating_ip.Name); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("description", floating_ip.Description); err != nil {
		return diag.Errorf("error setting description: %+v", err)
	}

	if err := d.Set("vm_name", floating_ip.VMName); err != nil {
		return diag.Errorf("error setting vm_name: %+v", err)
	}

	if err := d.Set("vm_port_id", floating_ip.VMPortID); err != nil {
		return diag.Errorf("error setting vm_port_id: %+v", err)
	}

	if err := d.Set("vm_fixed_ip", floating_ip.VMFixedIP); err != nil {
		return diag.Errorf("error setting vm_fixed_ip: %+v", err)
	}

	return nil
}

func resourceFloatingIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if !d.HasChanges("display_name", "description", "vm_name", "vm_port_id", "vm_fixed_ip") {
		return nil
	}
	proxy := meta.(*sdk.Client).FloatingIPAddresses

	floating_ip := mapResourceDataToFloatingIp(d)
	floating_ip.FloatingIPID = strfmt.UUID(d.Id())

	_, err := proxy.Update(ctx, floating_ip)
	if err != nil {
		return diag.Errorf("error while modifying floating ip: %+v", err)
	}

	return resourceFloatingIpRead(ctx, d, meta)
}

func resourceFloatingIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FloatingIPAddresses

	err := proxy.Delete(ctx, strfmt.UUID(d.Id()))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("floating ip with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting floating ip: %+v", err)
	}

	return nil
}

func mapResourceDataToFloatingIp(d *schema.ResourceData) *models.FloatingIP {
	return &models.FloatingIP{
		Description: d.Get("description").(string),
		Name:        d.Get("display_name").(string),
		VMName:      d.Get("vm_name").(string),
		VMPortID:    strfmt.UUID(d.Get("vm_port_id").(string)),
		VMFixedIP:   d.Get("vm_fixed_ip").(string),
	}
}
