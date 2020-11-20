---
page_title: "Virtual Network Data Source"
---

# Virtual Network Data Source

Data Source for getting virtual network by display name.

## Example Usage

```hcl
data "ochk_virtual_network" "vnet" {
  display_name = "virtual-network-display-name"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of virtual network.

## Attribute Reference

The following attributes are exported:
* `display_name` - Virtual network name.
* `subtenants` - List of subtenants identifiers in which virtual network will be available. 
* `router` - Router id attached to this network.
* `ipam_enabled` - The IP address management (IPAM) to discover IP address and Domain Name System (DNS) servers on the network and manage them.
* `dns_suffix` - The custom Domain Name System (DNS) suffix which should have your assigned domain name.
* `dns_search_suffix` - Domain Name System (DNS) suffix for search domain.
* `primary_dns_address` - Primary Domain Name System (DNS) server IP address.
* `secondary_dns_address` - Secondary Domain Name System (DNS) server IP address.
* `primary_wins_address` - Primary Windows Internet Name Service (WINS) address.
* `secondary_wins_address` - Secondary Windows Internet Name Service (WINS) address.
* `subnet_mask` - Subnet mask used to divide the IP address into network and host addresses. Set when `subnet_network_cidr` is set.
* `gateway_address` - IP address of local network gateway. Set when `subnet_network_cidr` is set.
* `subnet_gateway_address_cidr` - CIDR IP address of subnet gateway. Set when `subnet_network_cidr` is set.  
 
 
 
