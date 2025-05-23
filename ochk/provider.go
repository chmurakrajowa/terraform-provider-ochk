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
				DefaultFunc:      schema.EnvDefaultFunc("TF_VAR_host", nil),
				Description:      "host value",
			},
			"platform": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TF_VAR_platform", nil),
				Description: "platform value",
			},
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TF_VAR_api_key", nil),
				Description: "APi KEY value",
			},
			"platform_type": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TF_VAR_platform_type", nil),
				Description: "platform type value",
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
				Default:     "",
				Description: "path to debug log",
				Sensitive:   true,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"ochk_folders":                 dataSourceFolders(),
			"ochk_vrf":                     dataSourceVrf(),
			"ochk_vrfs":                    dataSourceVrfs(),
			"ochk_vpc":                     dataSourceVpc(),
			"ochk_vpcs":                    dataSourceVpcs(),
			"ochk_security_group":          dataSourceSecurityGroup(),
			"ochk_security_groups":         dataSourceSecurityGroups(),
			"ochk_project":                 dataSourceProject(),
			"ochk_projects":                dataSourceProjects(),
			"ochk_service":                 dataSourceService(),
			"ochk_services":                dataSourceServices(),
			"ochk_virtual_machine":         dataSourceVirtualMachine(),
			"ochk_virtual_machines":        dataSourceVirtualMachines(),
			"ochk_ip_collection":           dataSourceIPCollection(),
			"ochk_ip_collections":          dataSourceIPCollections(),
			"ochk_deployment":              dataSourceDeployment(),
			"ochk_deployments":             dataSourceDeployments(),
			"ochk_deployment_types":        dataSourceDeploymentTypes(),
			"ochk_deployment_params_types": dataSourceDeploymentParamsTypes(),
			"ochk_virtual_network":         dataSourceVirtualNetwork(),
			"ochk_virtual_networks":        dataSourceVirtualNetworks(),
			"ochk_custom_service":          dataSourceCustomService(),
			"ochk_custom_services":         dataSourceCustomServices(),
			"ochk_kms_key":                 dataSourceKMSKey(),
			"ochk_kms_keys":                dataSourceKMSKeys(),
			"ochk_backup_plan":             dataSourceBackupPlan(),
			"ochk_backup_plans":            dataSourceBackupPlans(),
			"ochk_backup_list":             dataSourceBackupList(),
			"ochk_backup_lists":            dataSourceBackupLists(),
			"ochk_tag":                     dataSourceTag(),
			"ochk_tags":                    dataSourceTags(),
			"ochk_auto_nat":                dataSourceNat(),
			"ochk_auto_nats":               dataSourceAutoNats(),
			"ochk_manual_nat":              dataSourceNat(),
			"ochk_manual_nats":             dataSourceManualNats(),
			"ochk_ports_forwarding":        dataSourcePortsForwarding(),
			"ochk_port_forwarding":         dataSourcePortForwarding(),
			"ochk_firewall_ew_rule":        dataSourceFirewallEWRule(),
			"ochk_firewall_ew_rules":       dataSourceFirewallEWRules(),
			"ochk_firewall_sn_rule":        dataSourceFirewallSNRule(),
			"ochk_firewall_sn_rules":       dataSourceFirewallSNRules(),
			"ochk_firewall_rules":          dataSourceFirewallRules(),
			"ochk_firewall_rule":           dataSourceFirewallRule(),
			"ochk_public_ip_addresses":     dataSourcePublicIPAddresses(),
			"ochk_floating_ip_address":     dataSourceFloatingIPAddress(),
			"ochk_floating_ip_addresses":   dataSourceFloatingIPAddresses(),
			"ochk_snapshot":                dataSourceSnapshot(),
			"ochk_snapshots":               dataSourceSnapshots(),
			"ochk_billing_account":         dataSourceBillingAccount(),
			"ochk_billing_accounts":        dataSourceBillingAccounts(),
			"ochk_floating_ip_vms":         dataSourceFloatingIPVms(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"ochk_firewall_ew_rule":    resourceFirewallEWRule(),
			"ochk_firewall_sn_rule":    resourceFirewallSNRule(),
			"ochk_firewall_rule":       resourceFirewallRule(),
			"ochk_vpc":                 resourceVpc(),
			"ochk_security_group":      resourceSecurityGroup(),
			"ochk_project":             resourceProject(),
			"ochk_virtual_network":     resourceVirtualNetwork(),
			"ochk_ip_collection":       resourceIPCollection(),
			"ochk_virtual_machine":     resourceVirtualMachine(),
			"ochk_custom_service":      resourceCustomService(),
			"ochk_kms_key":             resourceKMSKey(),
			"ochk_tag":                 resourceTag(),
			"ochk_auto_nat":            resourceAutoNat(),
			"ochk_manual_nat":          resourceManualNat(),
			"ochk_snapshot":            resourceSnapshot(),
			"ochk_billing_account":     resourceBillingAccount(),
			"ochk_floating_ip_address": resourceFloatingIp(),
			"ochk_port_forwarding":     resourcePortForwarding(),
		},
		ConfigureContextFunc: func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
			client, err := sdk.NewClient(
				ctx,
				d.Get("host").(string),
				d.Get("platform").(string),
				d.Get("api_key").(string),
				d.Get("insecure").(bool),
				d.Get("debug_log_file").(string),
				d.Get("platform_type").(string),
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
