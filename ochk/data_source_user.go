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
			"disabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"locked": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"user_principal_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"principal_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"principal_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"principal_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"principal_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Users

	name := d.Get("name").(string)

	users, err := proxy.ListByDisplayName(ctx, name)
	if err != nil {
		return diag.Errorf("error while listing users: %+v", err)
	}

	if len(users) < 1 {
		return diag.Errorf("no user found for name: %s", name)
	}

	if len(users) > 1 {
		return diag.Errorf("more than one user with name: %s found!", name)
	}

	d.SetId(users[0].UserID)

	if err := d.Set("name", users[0].Name); err != nil {
		return diag.Errorf("error setting user_type: %+v", err)
	}
	if err := d.Set("email_address", users[0].EmailAddress); err != nil {
		return diag.Errorf("error setting email_address: %+v", err)
	}
	if err := d.Set("description", users[0].Description); err != nil {
		return diag.Errorf("error setting description: %+v", err)
	}
	if err := d.Set("first_name", users[0].FirstName); err != nil {
		return diag.Errorf("error setting first_name: %+v", err)
	}
	if err := d.Set("last_name", users[0].LastName); err != nil {
		return diag.Errorf("error setting last_name: %+v", err)
	}
	if err := d.Set("disabled", users[0].Disabled); err != nil {
		return diag.Errorf("error setting disabled: %+v", err)
	}
	if err := d.Set("locked", users[0].Locked); err != nil {
		return diag.Errorf("error setting locked: %+v", err)
	}
	if err := d.Set("user_principal_name", users[0].UserPrincipalName); err != nil {
		return diag.Errorf("error setting user_principal_name: %+v", err)
	}
	if err := d.Set("principal_id", users[0].PrincipalID.ID); err != nil {
		return diag.Errorf("error setting principal_id: %+v", err)
	}
	if err := d.Set("principal_name", users[0].PrincipalID.Name); err != nil {
		return diag.Errorf("error setting principal_name: %+v", err)
	}
	if err := d.Set("principal_domain", users[0].PrincipalID.Domain); err != nil {
		return diag.Errorf("error setting principal_domain: %+v", err)
	}
	if err := d.Set("principal_type", users[0].PrincipalID.PrincipalType); err != nil {
		return diag.Errorf("error setting principal_type: %+v", err)
	}

	return nil
}
