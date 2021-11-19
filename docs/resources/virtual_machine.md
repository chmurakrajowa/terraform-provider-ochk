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
  locals {
  ssh_key = "ssh-rsa 
    ....  example@example.com"
  }
  
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
 
  backup_lists = [
     "example_backup_id"
  ]
  
  billing_tags = [
    'example_billing_tag1_id',
    'example_billing_tag2_id',
    
  ]
  system_tags = [
    "example_system_tag1_id",
    "example_system_tag2_id",
  ]
  
  os_tyoe = "LINUX/WINODWS"
  ovf_ip_configuration = false
  initial_user_name = "root"
  initial_password = "initial_password_for_vm"
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
  * **`STANDARD_W1`** - virtual machine located in Data Center 1 
  * **`STANDARD_W2`** - virtual machine located in Data Center 2 
* `deployment_id` - (Required) The unique deployment's identifier, use `ochk_deployment` data source for getting identifier by name. 
* `subtenant_id` - (Required) Business group's identifier, use `ochk_subtenant` data source for getting identifier by name.
* `virtual_network_devices` - (Required) List of virtual network devices. Each element must have the following values:
    * **virtual_network_id** - (Required) The unique identifier of virtual network. Virtual network allows the virtual machine to communicate with the rest of your network, host machine, and other virtual machines. Use `ochk_virtual_network` data source for getting identifier by name.
* `additional_virtual_disks` - (Optional) List of additional virtual disks. Additional disk will be created on the same storage as the virtual machine configuration. Each element must have the following values: 
    * **controller_id** - (Required) The unique identifier of controller. The only supported value for now is "0".
    * **lun_id** - (Required) Number used to identify a logical unit. Set this to consecutive int numbers > 0. When updating, e.g. extending size, `lun_id` needs to be preserved.
    * **size_mb** - (Required) Size in megabytes of the additional disk. When updating, only disk extensions are supported (larger value).
    * **device_type** - (Optional) Type of the device. Only supported value for now: `SCSI`. 
* `encryption` - (Optional) Enables VM encryption. Defaults to false. If this attribute is changed (performing update), encryption is either disabled (false) or VM is encrypted in place (true).  
* `encryption_key_id` - (Optional) Identifier of encryption key. If not set, encryption is managed automatically. Use `ochk_kms_key` to get key id.  
* `encryption_recrypt` - (Optional) Re-encryption operation: `NONE`, `SHALLOW`, `DEEP`. Provide `SHALLOW` or `DEEP` when enabling encryption on existing VM (when updating).                                                                                                          
* `os_type` - (Optional) Only for virtual machines created from ISO/OVF file.
* `ovf_ip_configuration` (Optional) Only for virtual machines created from ISO/OVF file.
* `initial_user_name` (Optional) Only for virtual machines created from OVF to set ssh-key or ip address.
* `backup_lists` (Optional) Backup list for virtual machine
* `billing_tags` (Optional) Billing tags for virtual machine
* `system_tags` (Optional) System tags for virtual machine

## Attribute Reference

The following attributes are exported in addition to above arguments:
* `virtual_disk` - Details of a system disk storage created by default. Additional disk will be created on the same storage as the virtual machine configuration. Each element has the following values:
    * **controller_id** - The unique identifier of controller.
    * **lun_id** - Number used to identify a logical unit.
    * **size_mb** - Size in megabytes of the disk.
    * **device_type** - Type of the device.
* `created_by` - Who created this resource.
* `created_at` - When this resource was created.
* `modified_by` - Who last modified this resource. 
* `modified_at` - When last modification occurred.  
* `ip_address` - Default virtual machine ip address
