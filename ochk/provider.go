package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateHost,
				Description:  "host value",
			},
			"tenant": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "tenant value",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "username value",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "password value",
				Sensitive:   true,
			},
			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "if set uses http scheme instead of https",
				Default:     false,
			},
			"debug_log_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "path to debug log",
				Sensitive:   true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ochk_security_group": resourceSecurityGroup(),
		},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			return sdk.NewClient(
				d.Get("host").(string),
				d.Get("tenant").(string),
				d.Get("username").(string),
				d.Get("password").(string),
				d.Get("insecure").(bool),
				d.Get("debug_log_file").(string),
			)
		},
	}
}

func validateHost(val interface{}, key string) (warns []string, errs []error) {
	if val == nil || val.(string) == "" {
		errs = append(errs, fmt.Errorf("%s value is not valid: %s", key, val.(string)))
	}

	return nil, errs
}
