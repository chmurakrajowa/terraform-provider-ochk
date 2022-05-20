package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUserRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Users

	name := d.Get("name").(string)

	user, err := proxy.GetUserByUserName(ctx, name)
	if err != nil {
		return diag.Errorf("error while listing users: %+v", err)
	}

	d.SetId(user.SamAccountName)

	if err := d.Set("display_name", user.DisplayName); err != nil {
		return diag.Errorf("error setting user display_name: %+v", err)
	}
	if err := d.Set("email_address", user.Email); err != nil {
		return diag.Errorf("error setting email_address: %+v", err)
	}
	if err := d.Set("description", user.Description); err != nil {
		return diag.Errorf("error setting description: %+v", err)
	}
	if err := d.Set("first_name", user.FirstName); err != nil {
		return diag.Errorf("error setting first_name: %+v", err)
	}
	if err := d.Set("last_name", user.LastName); err != nil {
		return diag.Errorf("error setting last_name: %+v", err)
	}
	if err := d.Set("enabled", user.Enabled); err != nil {
		return diag.Errorf("error setting enabled: %+v", err)
	}

	return nil
}
