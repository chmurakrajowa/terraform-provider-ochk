package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBillingAccount() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBillingAccountRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"discount": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"alarms": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"cost": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"projects": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceBillingAccountRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	proxy := meta.(*sdk.Client).Accounts

	accountName := d.Get("display_name").(string)

	accounts, err := proxy.ListAccountByName(ctx, accountName)

	if err != nil {
		return diag.Errorf("error while listing accounts: %+v", err)
	}

	if len(accounts) < 1 {
		return diag.Errorf("no account for name: %s", accountName)
	}

	if len(accounts) > 1 {
		return diag.Errorf("more than one billing account with name: %s found!", accountName)
	}

	d.SetId(fmt.Sprint(accounts[0].AccountID))

	if err := d.Set("display_name", accounts[0].AccountName); err != nil {
		return diag.Errorf("error setting account name: %+v", err)
	}
	if err := d.Set("account_description", accounts[0].AccountDescription); err != nil {
		return diag.Errorf("error setting account description: %+v", err)
	}
	if err := d.Set("discount", accounts[0].Discount); err != nil {
		return diag.Errorf("error setting discount: %+v", err)
	}
	if err := d.Set("alarms", accounts[0].Alarms); err != nil {
		return diag.Errorf("error setting alarms: %+v", err)
	}
	if err := d.Set("cost", accounts[0].Cost); err != nil {
		return diag.Errorf("error setting cost: %+v", err)
	}
	if err := d.Set("projects", flattenAccProjects(accounts[0].Projects)); err != nil {
		return diag.Errorf("error setting projects: %+v", err)
	}

	return nil
}
