---
page_title: "Virtual Network Data Source"
---

# Virtual Network Data Source

Data Source for getting virtual network by display name.

## Example Usage

```hcl
data "ochk_virtual_network" "{{ .DataSourceName}}" {
  display_name = "virtual_network_name"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of virtual network.

## Attribute Reference

The following attributes are exported:
* `vpc_id` - VPC id attached to this network. 
* `folder_path` - Virtual network folder path, default: `/`.
* `ipam_enabled` - The IP address management (IPAM) to discover IP address and Domain Name System (DNS) servers on the network and manage them.
* `subnet_network_cidr` - Subnet network cidr.
* `subnet_mask` - Subnet mask used to divide the IP address into network and host addresses. Set when `subnet_network_cidr` is set.
* `gateway_address` - IP address of local network gateway. Set when `subnet_network_cidr` is set.
* `subnet_gateway_address_cidr` - CIDR IP address of subnet gateway. Set when `subnet_network_cidr` is set.
* `dns_servers` - DNS server that will be transferred to the network configuration of the created virtual machine.
Each entry has following values:
* **address**: IP address of DNS server
* **id**: DNS server id
