package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFloatingIPAddress() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFloatingIPAddressRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_adress": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vm_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vm_port_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vm_fixed_ip": {
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

func dataSourceFloatingIPAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FloatingIPAddresses
	name := d.Get("display_name").(string)
	floatingIp, err := proxy.ListByName(ctx, name)

	if err != nil {
		return diag.Errorf("floating ip for ruleId %s not found: %+v", floatingIp, err)
	}

	if len(floatingIp) < 1 {
		return diag.Errorf("no floating ip found for display_name: %s", name)
	}

	if len(floatingIp) > 1 {
		return diag.Errorf("more than one floating ip with display_name: %s found!", name)
	}

	if err := d.Set("display_name", floatingIp[0].Name); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("description", floatingIp[0].Description); err != nil {
		return diag.Errorf("error setting description: %+v", err)
	}

	if err := d.Set("created_by", floatingIp[0].CreatedBy); err != nil {
		return diag.Errorf("error setting created_by: %+v", err)
	}

	if err := d.Set("created_at", floatingIp[0].CreationDate.String()); err != nil {
		return diag.Errorf("error setting created_at: %+v", err)
	}

	if err := d.Set("modified_by", floatingIp[0].ModifiedBy); err != nil {
		return diag.Errorf("error setting modified_by: %+v", err)
	}

	if err := d.Set("modified_at", floatingIp[0].ModificationDate.String()); err != nil {
		return diag.Errorf("error setting modified_at: %+v", err)
	}

	if err := d.Set("public_adress", floatingIp[0].PublicAddress); err != nil {
		return diag.Errorf("error setting public_adress: %+v", err)
	}

	if err := d.Set("project_id", floatingIp[0].OscProjectID); err != nil {
		return diag.Errorf("error setting project_id: %+v", err)
	}

	if err := d.Set("vm_name", floatingIp[0].VMName); err != nil {
		return diag.Errorf("error setting vm_name: %+v", err)
	}

	if err := d.Set("vm_port_id", floatingIp[0].VMPortID); err != nil {
		return diag.Errorf("error setting vm_port_id: %+v", err)
	}

	if err := d.Set("vm_fixed_ip", floatingIp[0].VMFixedIP); err != nil {
		return diag.Errorf("error setting vm_fixed_ip: %+v", err)
	}

	d.SetId(floatingIp[0].FloatingIPID.String())
	return nil
}
