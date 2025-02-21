---
page_title: "Virtual network Resource"
---

# Virtual network Resource

Resource for managing virtual networks (vNets) to enable communication between multiple virtual machines. 

## Example Usage

### Minimal example
```hcl
data "ochk_project" "project" {
  display_name = "project_name"
}

resource "ochk_virtual_network" "{{ .ResourceName}}" {
	display_name = "virtual_machine_name"
	ipam_enabled = false
	project_id = data.ochk_project.project.id
}
```

### Example With IPAM, subnet and vpc
```hcl

data "ochk_project" "project" {
  display_name = "project_name"
}

data "ochk_vrf" "vrf" {
  display_name = "vrf_name"
}

data "ochk_vpc" "vpc" {
  vrf_id = data.ochk_vrf.vrf.id
  display_name = "vpc_name"
}

resource "ochk_virtual_network" "{{ .ResourceName}}" {
	display_name = "virtual_network_name"
	project_id = data.ochk_project.project.id
	ipam_enabled = true
	vpc_id = data.ochk_vpc.vpc.id
	subnet_network_cidr = "10.16.1.0/24"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Virtual network name. Updates to this attribute forces recreate.
* `project_id` - (Required) Project identifier, use `ochk_project` data source for getting identifier by name.
* `folder_path` - (Optional) Folder path for virtual network. Default `/`.
* `ipam_enabled` - (Optional) The IP address management (IPAM) to discover IP address and Domain Name System (DNS) servers on the network and manage them. True for enabling IPAM. Defaults to `false`. Updates to this attribute forces recreate.
* `subnet_mask` - (Optional) Subnet mask used to divide the IP address into network and host addresses. Updates to this attribute forces recreate.
* `vpc_id` - (Optional) VPC id attached to this network. Use `ochk_vpc` data source to get id by display name.

## Attribute Reference

The following attributes are exported in addition to above arguments:
* `gateway_address` - Gateway address, set when `subnet_network_cidr` is set.
* `subnet_mask` - Subnet mask, set when `subnet_network_cidr` is set.
* `subnet_gateway_address_cidr` - Subnet gateway address cidr, set when `subnet_network_cidr` is set.  
* `subnet_network_cidr` - Subnet network cidr.
* `dns_servers` - DNS server that will be transferred to the network configuration of the created virtual machine.
Each entry has following values:
* **address**: IP address of DNS server
* **id**: DNS server id
 
