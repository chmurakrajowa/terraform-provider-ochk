// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/accounts"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/auth"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/available_public_ip"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/backups"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/billing_alarm_definition"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/billing_kms"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/billing_period"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/billing_pip"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/billing_vm"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/billing_vrf"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/custom_services"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/default_services"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/deployments"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/dfw_rule"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/firewall_rule"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/floating_ip"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/floating_ip_vms"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/folder"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/gfw_rule"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/health_check"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/identification"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/ip_collection"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/ipam_service"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/key"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/key_schedule"
	logops "github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/log"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/log_category"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/log_category_user"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/log_stats"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/log_stats_user"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/log_user"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/logout"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/nat_rule"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/port_forwarding"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/projects"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/public_ip"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/readiness"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/requests"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/router"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/security_group"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/tags"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/virtual_machine"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/virtual_machine_snapshot"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/virtual_network"
)

// Default ochk HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new ochk HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Ochk {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new ochk HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Ochk {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new ochk client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Ochk {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Ochk)
	cli.Transport = transport
	cli.Accounts = accounts.New(transport, formats)
	cli.Auth = auth.New(transport, formats)
	cli.AvailablePublicIP = available_public_ip.New(transport, formats)
	cli.Backups = backups.New(transport, formats)
	cli.BillingAlarmDefinition = billing_alarm_definition.New(transport, formats)
	cli.BillingKms = billing_kms.New(transport, formats)
	cli.BillingPeriod = billing_period.New(transport, formats)
	cli.BillingPip = billing_pip.New(transport, formats)
	cli.BillingVM = billing_vm.New(transport, formats)
	cli.BillingVrf = billing_vrf.New(transport, formats)
	cli.CustomServices = custom_services.New(transport, formats)
	cli.DefaultServices = default_services.New(transport, formats)
	cli.Deployments = deployments.New(transport, formats)
	cli.DfwRule = dfw_rule.New(transport, formats)
	cli.FirewallRule = firewall_rule.New(transport, formats)
	cli.FloatingIP = floating_ip.New(transport, formats)
	cli.FloatingIPVms = floating_ip_vms.New(transport, formats)
	cli.Folder = folder.New(transport, formats)
	cli.GfwRule = gfw_rule.New(transport, formats)
	cli.HealthCheck = health_check.New(transport, formats)
	cli.Identification = identification.New(transport, formats)
	cli.IPCollection = ip_collection.New(transport, formats)
	cli.IpamService = ipam_service.New(transport, formats)
	cli.Key = key.New(transport, formats)
	cli.KeySchedule = key_schedule.New(transport, formats)
	cli.Log = logops.New(transport, formats)
	cli.LogCategory = log_category.New(transport, formats)
	cli.LogCategoryUser = log_category_user.New(transport, formats)
	cli.LogStats = log_stats.New(transport, formats)
	cli.LogStatsUser = log_stats_user.New(transport, formats)
	cli.LogUser = log_user.New(transport, formats)
	cli.Logout = logout.New(transport, formats)
	cli.NatRule = nat_rule.New(transport, formats)
	cli.PortForwarding = port_forwarding.New(transport, formats)
	cli.Projects = projects.New(transport, formats)
	cli.PublicIP = public_ip.New(transport, formats)
	cli.Readiness = readiness.New(transport, formats)
	cli.Requests = requests.New(transport, formats)
	cli.Router = router.New(transport, formats)
	cli.SecurityGroup = security_group.New(transport, formats)
	cli.Tags = tags.New(transport, formats)
	cli.VirtualMachine = virtual_machine.New(transport, formats)
	cli.VirtualMachineSnapshot = virtual_machine_snapshot.New(transport, formats)
	cli.VirtualNetwork = virtual_network.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Ochk is a client for ochk
type Ochk struct {
	Accounts accounts.ClientService

	Auth auth.ClientService

	AvailablePublicIP available_public_ip.ClientService

	Backups backups.ClientService

	BillingAlarmDefinition billing_alarm_definition.ClientService

	BillingKms billing_kms.ClientService

	BillingPeriod billing_period.ClientService

	BillingPip billing_pip.ClientService

	BillingVM billing_vm.ClientService

	BillingVrf billing_vrf.ClientService

	CustomServices custom_services.ClientService

	DefaultServices default_services.ClientService

	Deployments deployments.ClientService

	DfwRule dfw_rule.ClientService

	FirewallRule firewall_rule.ClientService

	FloatingIP floating_ip.ClientService

	FloatingIPVms floating_ip_vms.ClientService

	Folder folder.ClientService

	GfwRule gfw_rule.ClientService

	HealthCheck health_check.ClientService

	Identification identification.ClientService

	IPCollection ip_collection.ClientService

	IpamService ipam_service.ClientService

	Key key.ClientService

	KeySchedule key_schedule.ClientService

	Log logops.ClientService

	LogCategory log_category.ClientService

	LogCategoryUser log_category_user.ClientService

	LogStats log_stats.ClientService

	LogStatsUser log_stats_user.ClientService

	LogUser log_user.ClientService

	Logout logout.ClientService

	NatRule nat_rule.ClientService

	PortForwarding port_forwarding.ClientService

	Projects projects.ClientService

	PublicIP public_ip.ClientService

	Readiness readiness.ClientService

	Requests requests.ClientService

	Router router.ClientService

	SecurityGroup security_group.ClientService

	Tags tags.ClientService

	VirtualMachine virtual_machine.ClientService

	VirtualMachineSnapshot virtual_machine_snapshot.ClientService

	VirtualNetwork virtual_network.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Ochk) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Accounts.SetTransport(transport)
	c.Auth.SetTransport(transport)
	c.AvailablePublicIP.SetTransport(transport)
	c.Backups.SetTransport(transport)
	c.BillingAlarmDefinition.SetTransport(transport)
	c.BillingKms.SetTransport(transport)
	c.BillingPeriod.SetTransport(transport)
	c.BillingPip.SetTransport(transport)
	c.BillingVM.SetTransport(transport)
	c.BillingVrf.SetTransport(transport)
	c.CustomServices.SetTransport(transport)
	c.DefaultServices.SetTransport(transport)
	c.Deployments.SetTransport(transport)
	c.DfwRule.SetTransport(transport)
	c.FirewallRule.SetTransport(transport)
	c.FloatingIP.SetTransport(transport)
	c.FloatingIPVms.SetTransport(transport)
	c.Folder.SetTransport(transport)
	c.GfwRule.SetTransport(transport)
	c.HealthCheck.SetTransport(transport)
	c.Identification.SetTransport(transport)
	c.IPCollection.SetTransport(transport)
	c.IpamService.SetTransport(transport)
	c.Key.SetTransport(transport)
	c.KeySchedule.SetTransport(transport)
	c.Log.SetTransport(transport)
	c.LogCategory.SetTransport(transport)
	c.LogCategoryUser.SetTransport(transport)
	c.LogStats.SetTransport(transport)
	c.LogStatsUser.SetTransport(transport)
	c.LogUser.SetTransport(transport)
	c.Logout.SetTransport(transport)
	c.NatRule.SetTransport(transport)
	c.PortForwarding.SetTransport(transport)
	c.Projects.SetTransport(transport)
	c.PublicIP.SetTransport(transport)
	c.Readiness.SetTransport(transport)
	c.Requests.SetTransport(transport)
	c.Router.SetTransport(transport)
	c.SecurityGroup.SetTransport(transport)
	c.Tags.SetTransport(transport)
	c.VirtualMachine.SetTransport(transport)
	c.VirtualMachineSnapshot.SetTransport(transport)
	c.VirtualNetwork.SetTransport(transport)
}
