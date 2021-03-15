package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validateHost,
				DefaultFunc:      schema.EnvDefaultFunc("OCHK_HOST", nil),
				Description:      "host value",
			},
			"tenant": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OCHK_TENANT", nil),
				Description: "tenant value",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OCHK_USERNAME", nil),
				Description: "username value",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OCHK_PASSWORD", nil),
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
		DataSourcesMap: map[string]*schema.Resource{
			"ochk_gateway_policy":  dataSourceGatewayPolicy(),
			"ochk_logical_port":    dataSourceLogicalPort(),
			"ochk_network":         dataSourceNetwork(),
			"ochk_router":          dataSourceRouter(),
			"ochk_security_group":  dataSourceSecurityGroup(),
			"ochk_security_policy": dataSourceSecurityPolicy(),
			"ochk_subtenant":       dataSourceSubtenant(),
			"ochk_service":         dataSourceService(),
			"ochk_user":            dataSourceUser(),
			"ochk_virtual_machine": dataSourceVirtualMachine(),
			"ochk_ip_collection":   dataSourceIPCollection(),
			"ochk_deployment":      dataSourceDeployment(),
			"ochk_virtual_network": dataSourceVirtualNetwork(),
			"ochk_custom_service":  dataSourceCustomService(),
			"ochk_kms_key":         dataSourceKMSKey(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"ochk_firewall_ew_rule": resourceFirewallEWRule(),
			"ochk_firewall_sn_rule": resourceFirewallSNRule(),
			"ochk_router":           resourceRouter(),
			"ochk_security_group":   resourceSecurityGroup(),
			"ochk_subtenant":        resourceSubtenant(),
			"ochk_virtual_network":  resourceVirtualNetwork(),
			"ochk_ip_collection":    resourceIPCollection(),
			"ochk_virtual_machine":  resourceVirtualMachine(),
			"ochk_custom_service":   resourceCustomService(),
			"ochk_kms_key":          resourceKMSKey(),
		},
		ConfigureContextFunc: func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
			client, err := sdk.NewClient(
				ctx,
				d.Get("host").(string),
				d.Get("tenant").(string),
				d.Get("username").(string),
				d.Get("password").(string),
				d.Get("insecure").(bool),
				d.Get("debug_log_file").(string),
			)
			if err != nil {
				return nil, diag.FromErr(err)
			}

			return client, nil
		},
	}
}

func validateHost(val interface{}, path cty.Path) diag.Diagnostics {
	if val == nil || val.(string) == "" {
		diag.Errorf("%s value is not valid: %s", path, val.(string))
	}

	return nil
}
