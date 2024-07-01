package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
	"time"
)

const (
	AcctRetryTimeout = 1 * time.Minute
)

func resourceBillingAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAccountCreate,
		UpdateContext: resourceAccountUpdate,
		ReadContext:   resourceAccountRead,
		DeleteContext: resourceAccountDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(AcctRetryTimeout),
			Update: schema.DefaultTimeout(AcctRetryTimeout),
			Delete: schema.DefaultTimeout(AcctRetryTimeout),
		},

		Importer: &schema.ResourceImporter{
			StateContext: resourceAccountImportState,
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"discount": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"alarms": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"cost": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"projects": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project_id": {
							Type:     schema.TypeString,
							Required: true,
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

func resourceAccountImportState(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	d.SetId(strings.ToLower(d.Id()))
	return []*schema.ResourceData{d}, nil
}

func resourceAccountCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Accounts

	account := mapResourceDataToAccount(d)

	created, err := proxy.Create(ctx, account)
	if err != nil {
		return diag.Errorf("error while creating account: %+v", err)
	}

	d.SetId(created.AccountID)
	return resourceAccountRead(ctx, d, meta)
}

func resourceAccountUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Accounts

	account := mapResourceDataToAccount(d)
	account.AccountID = d.Id()

	_, err := proxy.Update(ctx, account)
	if err != nil {
		return diag.Errorf("error while modifying account: %+v", err)
	}

	return resourceAccountRead(ctx, d, meta)
}

func resourceAccountRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Accounts

	account, err := proxy.Read(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("account with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading account: %+v", err)
	}

	if err := d.Set("display_name", account.AccountName); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("account_description", account.AccountDescription); err != nil {
		return diag.Errorf("error setting account_description: %+v", err)
	}

	if err := d.Set("discount", account.Discount); err != nil {
		return diag.Errorf("error setting discount: %+v", err)
	}

	if err := d.Set("alarms", account.Alarms); err != nil {
		return diag.Errorf("error setting alarms: %+v", err)
	}

	if err := d.Set("cost", account.Cost); err != nil {
		return diag.Errorf("error setting cost: %+v", err)
	}

	if err := d.Set("projects", flattenAccProjects(account.Projects)); err != nil {
		return diag.Errorf("error setting projects: %+v", err)
	}

	return nil
}

func resourceAccountDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Accounts

	err := proxy.Delete(ctx, d.Id())

	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("account with id %s not found: %+v", id, err)
		}
		return diag.Errorf("error while deleting account: %+v", err)
	}
	return nil
}

func mapResourceDataToAccount(d *schema.ResourceData) *models.AccountInstance {
	return &models.AccountInstance{
		AccountName:        d.Get("display_name").(string),
		AccountDescription: d.Get("account_description").(string),
		Projects:           expandAccountProjects(d.Get("projects").(*schema.Set).List()),
	}
}

func expandAccountProjects(in []interface{}) []*models.AccountProjectInstance {
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.AccountProjectInstance, len(in))
	for i, v := range in {
		m := v.(map[string]interface{})

		project := &models.AccountProjectInstance{
			ProjectID: m["project_id"].(string),
			Name:      m["display_name"].(string),
		}

		if displayName, ok := m["display_name"].(string); ok && displayName != "" {
			project.Name = displayName
		}

		out[i] = project
	}
	return out
}
