---
page_title: "Virtual network Resource"
---

# Virtual network Resource

Resource for managing virtual networks (vNets) to enable communication between multiple virtual machines. In the onboarding process, the client declares the range of IP addresses that will be used to create the separate subnets. The new subnet must be within the range of available scopes, otherwise an error will be displayed and such vNet will not be created.

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

* `display_name` - (Required) Virtual machine's display name.
* `initial_password` - (Required) Initial password. Cannot be changed after creation.
* `power_state` - (Required) Power state, values: `poweredOn`, `poweredOff`, `suspended`. 
* `resource_profile` - (Required) Resource profile, values: `CUSTOM`, `SIZE_L`, `SIZE_M`, `SIZE_S`, `SIZE_XL`, `SIZE_XS`.
* `storage_policy` - (Required) Storage policy, value: `ENTERPRISE`, `STANDARD`, `UNKNOWN`.
* `deployment_id` - (Required) Deployment's identifier, use `ochk_deployment` data source for getting identifier by name. 
* `subtenant_id` - (Required) Subtenants' identifier, use `ochk_subtenant` data source for getting identifier by name.
* `virtual_network_devices` - (Required) List of virtual network devices. Each element must have the following values:
    * **virtual_network_id** - (Required) Identifier of virtual network. Use `ochk_virtual_network` data source for getting identifier by name.
* `additional_virtual_disks` - (Optional) List of additional virtual disks. Each element must have the following values: 
    * **controller_id** - (Required) 
    * **lun_id** - (Required)
    * **size_mb** - (Required) Size in megabytes. 
    * **device_type** - (Required)
* `virtual_disk` - (Optional) Configuration of system disk. Each element must have the following values:
    * **controller_id** - (Required)
    * **lun_id** - (Required)
    * **size_mb** - (Required)
    * **device_type** - (Required)

## Attribute Reference

The following attributes are exported in addition to above arguments:
* `created_by` - Who created this resource.
* `created_at` - When this resource was created.
* `modified_by` - Who last modified this resource. 
* `modified_at` - When last modification occurred.  
