package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBillingAccounts() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBillingAccountsRead,
		Schema: map[string]*schema.Schema{
			"accounts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
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

func dataSourceBillingAccountsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	proxy := meta.(*sdk.Client).Accounts
	accounts, err := proxy.ListAccounts(ctx)

	if err := d.Set("accounts", flattenAccount(accounts)); err != nil {
		return diag.Errorf("error setting accounts list: %v", err)
	}

	if err != nil {
		return diag.Errorf("error while listing accounts: %+v", err)
	}

	d.SetId("accounts-list")
	return nil
}
