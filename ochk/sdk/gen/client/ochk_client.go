// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/context_profiles"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/default_services"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/firewall_rules_e_w"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/firewall_rules_s_n"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/gateway_policies"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/ip_sets"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/logical_ports"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/routers"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/security_groups"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/security_policies"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/vidm_controller"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/virtual_machines"
)

// Default ochk HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "iaas-api-proxy.ochk.pilot"
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
	cli.ContextProfiles = context_profiles.New(transport, formats)
	cli.DefaultServices = default_services.New(transport, formats)
	cli.FirewallRulesew = firewall_rules_e_w.New(transport, formats)
	cli.FirewallRulessn = firewall_rules_s_n.New(transport, formats)
	cli.GatewayPolicies = gateway_policies.New(transport, formats)
	cli.IPSets = ip_sets.New(transport, formats)
	cli.LogicalPorts = logical_ports.New(transport, formats)
	cli.Routers = routers.New(transport, formats)
	cli.SecurityGroups = security_groups.New(transport, formats)
	cli.SecurityPolicies = security_policies.New(transport, formats)
	cli.VidmController = vidm_controller.New(transport, formats)
	cli.VirtualMachines = virtual_machines.New(transport, formats)
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
	ContextProfiles context_profiles.ClientService

	DefaultServices default_services.ClientService

	FirewallRulesew firewall_rules_e_w.ClientService

	FirewallRulessn firewall_rules_s_n.ClientService

	GatewayPolicies gateway_policies.ClientService

	IPSets ip_sets.ClientService

	LogicalPorts logical_ports.ClientService

	Routers routers.ClientService

	SecurityGroups security_groups.ClientService

	SecurityPolicies security_policies.ClientService

	VidmController vidm_controller.ClientService

	VirtualMachines virtual_machines.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Ochk) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.ContextProfiles.SetTransport(transport)
	c.DefaultServices.SetTransport(transport)
	c.FirewallRulesew.SetTransport(transport)
	c.FirewallRulessn.SetTransport(transport)
	c.GatewayPolicies.SetTransport(transport)
	c.IPSets.SetTransport(transport)
	c.LogicalPorts.SetTransport(transport)
	c.Routers.SetTransport(transport)
	c.SecurityGroups.SetTransport(transport)
	c.SecurityPolicies.SetTransport(transport)
	c.VidmController.SetTransport(transport)
	c.VirtualMachines.SetTransport(transport)
}
