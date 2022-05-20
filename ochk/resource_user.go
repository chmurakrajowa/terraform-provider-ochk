package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
)

const (
	UserRetryTimeout = 1 * time.Minute
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(UserRetryTimeout),
			Update: schema.DefaultTimeout(UserRetryTimeout),
			Delete: schema.DefaultTimeout(UserRetryTimeout),
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return old == "<hidden>"
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Users

	user, err := proxy.GetUserByUserName(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("user with name %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading user: %+v", err)
	}

	if err := mapUserToResourceData(d, user); err != nil {
		return diag.Errorf("error setting user: %v", err)
	}

	return nil
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*sdk.Client)

	user := mapResourceDataToUser(d)

	_, err := client.Users.Create(ctx, user)
	if err != nil {
		return diag.Errorf("error while creating user: %+v", err)
	}

	password := d.Get("password").(string)

	_, err = client.Users.SetPassword(ctx, user, password)

	if err != nil {
		_, err = client.Users.Delete(ctx, user.SamAccountName)
		if err != nil {
			return diag.Errorf("error while set password and rollback user: %+v", err)
		}
		return diag.Errorf("error while set user password: %+v", err)
	}

	d.SetId(user.SamAccountName)
	return resourceUserRead(ctx, d, meta)
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sdkClient := meta.(*sdk.Client)

	user := mapResourceDataToUser(d)
	user.SamAccountName = d.Id()

	_, err := sdkClient.Users.Update(ctx, user)
	if err != nil {
		return diag.Errorf("error while modifying user: %+v", err)
	}

	d.SetId(user.SamAccountName)

	return resourceUserRead(ctx, d, meta)
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkClient := meta.(*sdk.Client)

	_, err := sdkClient.Users.Delete(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("user with name %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting user: %+v", err)
	}

	return nil
}

func mapResourceDataToUser(d *schema.ResourceData) *models.ADUserInstance {
	userInstance := models.ADUserInstance{
		SamAccountName: d.Get("name").(string),
		DisplayName:    d.Get("display_name").(string),
		Email:          d.Get("email_address").(string),
		Description:    d.Get("description").(string),
		FirstName:      d.Get("first_name").(string),
		LastName:       d.Get("last_name").(string),
	}
	if enabled, ok := d.GetOk("enabled"); ok && enabled.(bool) {
		userInstance.Enabled = true
	}
	return &userInstance
}
func mapUserToResourceData(d *schema.ResourceData, user *models.ADUserInstance) error {

	if err := d.Set("name", user.SamAccountName); err != nil {
		return fmt.Errorf("error setting name: %w", err)
	}
	if err := d.Set("display_name", user.DisplayName); err != nil {
		return fmt.Errorf("error setting display_name: %w", err)
	}
	if err := d.Set("email_address", user.Email); err != nil {
		return fmt.Errorf("error setting email_address: %w", err)
	}
	if err := d.Set("description", user.Description); err != nil {
		return fmt.Errorf("error setting description: %w", err)
	}
	if err := d.Set("first_name", user.FirstName); err != nil {
		return fmt.Errorf("error setting first_name: %w", err)
	}
	if err := d.Set("last_name", user.LastName); err != nil {
		return fmt.Errorf("error setting last_name: %w", err)
	}
	if err := d.Set("enabled", user.Enabled); err != nil {
		return fmt.Errorf("error setting enabled: %w", err)
	}
	if err := d.Set("password", "<hidden>"); err != nil {
		return fmt.Errorf("error setting password: %w", err)
	}
	return nil
}
