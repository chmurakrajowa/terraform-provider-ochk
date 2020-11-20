---
page_title: "Virtual machine Resource"
---

# Virtual machine Resource

Resource for provisioning virtual machine that functions as a virtual computer system with its own CPU, memory, network interface, and storage, created on a physical hardware system. 

Warning: provisioning of virtual machine can take up to 15 minutes. 

## Example Usage

### Minimal example

```hcl

data "ochk_deployment" "centos" {
  display_name = "CentOS 7"
}

data "ochk_subtenant" "default" {
  name = "<subtenant-display-name>"
}

resource "ochk_virtual_network" "default" {
  display_name = "vnet"
  subtenants = [
    data.ochk_subtenant.default.id
  ]
}

resource "ochk_virtual_machine" "default" {
  display_name = "vm"
  deployment_id = data.ochk_deployment.centos.id
  initial_password = "<initial-password>"

  power_state = "poweredOn"
  resource_profile = "SIZE_S"
  storage_policy = "STANDARD"
  subtenant_id = data.ochk_subtenant.default.id

  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.default.id
  }
}

```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Specifies display name associated with this virtual machine. 
* `initial_password` - (Required) Initial password. Cannot be changed after creation.
* `power_state` - (Required) Power state information for the specified virtual machine. Value is one of: `poweredOn`, `poweredOff`, `suspended`. 
* `resource_profile` - (Required) The definition of the amount of resources that are allocated to virtual machines , values: 
  * **`CUSTOM`** - Custom configuration of vCPU and RAM
  * **`SIZE_XL`** - 16 vCPU, 64 GB RAM 
  * **`SIZE_L`** - 8 vCPU, 32 GB RAM 
  * **`SIZE_M`** - 4 vCPU, 16 GB RAM
  * **`SIZE_S`** - 2 vCPU, 8 GB RAM
  * **`SIZE_XS`** - 1 vCPU, 4 GB RAM
* `storage_policy` - (Required) Storage Policy associated with virtual machine. The policies control which type of storage is provided for the virtual machine, how the virtual machine is placed within the storage, and which data services are offered for the virtual machine; values: 
  * **`ENTERPRISE`** - virtual machine disks are distributed over two Data Centers
  * **`STANDARD`** - virtual machine disks are located in one Data Center 
* `deployment_id` - (Required) The unique deployment's identifier, use `ochk_deployment` data source for getting identifier by name. 
* `subtenant_id` - (Required) Business group's identifier, use `ochk_subtenant` data source for getting identifier by name.
* `virtual_network_devices` - (Required) List of virtual network devices. Each element must have the following values:
    * **virtual_network_id** - (Required) The unique identifier of virtual network. Virtual network allows the virtual machine to communicate with the rest of your network, host machine, and other virtual machines. Use `ochk_virtual_network` data source for getting identifier by name.
* `additional_virtual_disks` - (Optional) List of additional virtual disks. Additional disk will be created on the same storage as the virtual machine configuration. Each element must have the following values: 
    * **controller_id** - (Required) The unique identifier of controller.
    * **lun_id** - (Required) Number used to identify a logical unit.
    * **size_mb** - (Required) Size in megabytes of the additional disk. 
    * **device_type** - (Required) Type of the device.
* `virtual_disk` - (Optional) Configuration of system disk. Each element must have the following values:
    * **controller_id** - (Required) The unique identifier of controller.
    * **lun_id** - (Required) Number used to identify a logical unit.
    * **size_mb** - (Required) Size in megabytes of the virtual disk.
    * **device_type** - (Required) Type of the device.

## Attribute Reference

The following attributes are exported in addition to above arguments:
* `created_by` - Who created this resource.
* `created_at` - When this resource was created.
* `modified_by` - Who last modified this resource. 
* `modified_at` - When last modification occurred.  
