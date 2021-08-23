---
page_title: "Virtual network Resource"
---

# Virtual network Resource

Resource for managing virtual networks (vNets) to enable communication between multiple virtual machines. 
In the onboarding process, the client declares the range of IP addresses that will be used to create the separate subnets. 
The new subnet must be within the range of available scopes, otherwise an error will be displayed and such vNet will not be created.

Warning: provisioning of virtual networks can take up to 15 minutes. 

## Example Usage

### Minimal example
```hcl
data "ochk_subtenant" "subtenant" {
  name = "subtenant"
}

resource "ochk_virtual_network" "{{ .ResourceName}}" {
	display_name = "display name"
	ipam_enabled = false
	subtenants = ochk_subtenant.subtenant.id
}
```

### Example With IPAM, subnet and router
```hcl

data "ochk_subtenant" "subtenant" {
  name = "subtenant"
}

data "ochk_router" "router" {
  name = "router"
}

resource "ochk_virtual_network" "{{ .ResourceName}}" {
	display_name = "vnet name"
	ipam_enabled = true
	subtenants = [ochk_subtenant.subtenant.id]
	router = ochk_router.router.id
	dns_search_suffix = "test.lcl,prod.lcl"
	dns_suffix = "test.lcl"
	primary_dns_address = "192.168.1.6"
	secondary_dns_address = "192.168.1.2"
	primary_wins_address = "192.168.1.3"
	secondary_wins_address = "192.168.1.3"
	subnet_network_cidr = "10.16.1.0/24"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Virtual network name. Updates to this attribute forces recreate.
* `subtenants` - (Required) List of subtenants identifiers in which virtual network will be available. Use `ochk_subtenant` data source to find identifiers by name. 
* `router` - (Optional) Router id attached to this network. Use `ochk_router` data source to get id by display name.
* `ipam_enabled` - (Optional) The IP address management (IPAM) to discover IP address and Domain Name System (DNS) servers on the network and manage them. True for enabling IPAM. Defaults to `false`. Updates to this attribute forces recreate.
* `dns_suffix` - (Optional) The custom Domain Name System (DNS) suffix which should have your assigned domain name. Updates to this attribute forces recreate.
* `dns_search_suffix` - (Optional) Domain Name System (DNS) suffix for search domain. Updates to this attribute forces recreate.
* `primary_dns_address` - (Optional) Primary Domain Name System (DNS) server IP address. Updates to this attribute forces recreate.
* `secondary_dns_address` - (Optional) Secondary Domain Name System (DNS) server IP address. Updates to this attribute forces recreate.
* `primary_wins_address` - (Optional) Primary Windows Internet Name Service (WINS) address. Updates to this attribute forces recreate.
* `secondary_wins_address` - (Optional) Secondary Windows Internet Name Service (WINS) address. Updates to this attribute forces recreate.
* `subnet_mask` - (Optional) Subnet mask used to divide the IP address into network and host addresses. Updates to this attribute forces recreate.

## Attribute Reference

The following attributes are exported in addition to above arguments:
* `gateway_address` - gateway address, set when `subnet_network_cidr` is set.
* `subnet_mask` - subnet mask, set when `subnet_network_cidr` is set.
* `subnet_gateway_address_cidr` - subnet gateway address cidr, set when `subnet_network_cidr` is set.  
 
