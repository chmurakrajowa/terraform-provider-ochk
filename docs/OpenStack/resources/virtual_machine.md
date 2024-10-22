---
page_title: "Virtual machine Resource"
---

# Virtual machine Resource

Resource for provisioning virtual machine that functions as a virtual computer system with its own CPU, memory, network interface, and storage, created on a physical hardware system. 

Warning: provisioning of virtual machine can take up to 15 minutes. 

## Example Usage

### Minimal example

```hcl

data "ochk_deployment" "deploy1" {
  display_name = "deployment_name"
}

data "ochk_project" "default" {
  name = "project_name"
}

resource "ochk_virtual_network" "default" {
  display_name = "virtual_network_name"
  project_id = data.ochk_project.default.id
}

resource "ochk_virtual_machine" "{{ .ResourceName}}" {
  display_name = "virtual_machine_name"
  project_id = data.ochk_project.default.id

  deployment_id = data.ochk_deployment.deploy1.id

  initial_password = "initial-password"

  power_state = "poweredOn"
  cpu_count = 2
  memory_size_mb = 4096
  virtual_disk {
    size_mb = 40960
  }
          
  storage_policy = "STANDARD_W1"
  
  virtual_network_devices {
    virtual_network_id = data.ochk_virtual_network.default.id
  }
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Specifies display name associated with this virtual machine. 
* `deployment_id` - (Required) The unique deployment's identifier, use `ochk_deployment` data source for getting identifier by name.
* `initial_password` - (Required) Initial password. Cannot be changed after creation.
* `power_state` - (Required) Power state information for the specified virtual machine. Value is one of: `poweredOn`, `poweredOff`, `suspended`. 
* `cpu_count` - (Required) Vcpus count for virtual machine.
* `memory_size_mb` - (Required) Size of memory in megabytes for virtual machine.
* `storage_policy` - (Required) Storage Policy associated with virtual machine. The policies control which type of storage is provided for the virtual machine, how the virtual machine is placed within the storage, and which data services are offered for the virtual machine; values:
  * **`STANDARD_W1`** - virtual machine located in Data Center 1 
  * **`STANDARD_W2`** - virtual machine located in Data Center 2 
* `project_id` - (Required) Project id, use `ochk_project` data source for getting identifier by name.
* `virtual_disk` - (Required) Details of a system disk storage created by default. Disc has the following values:
  * **size_mb** - (Required) Size in megabytes of the disk.
* `folder_path` - (Optional) Folder path for virtual machine. Default `/`.
* `ssh_key` -(Optional) SSH key to set on machine. Default empty.
* `tags` (Optional) Tags for virtual machine.
* `deployment_params` - (Optional) Extra deployment parameters. Requirements: VM Tools installed and bash script on virtual machine which communicates with VM Tolls to set deployment param in OS. Each element must have the following values:
    * **param_name** - (Required) Param name which will be transferred to VM Tools.
    * **param_type** - (Required) Param type. Use `ochk_deployment_params_types` data source for getting supported types list.
    * **param_value** - (Required) Param value which will be transferred to VM Tools.
* `virtual_network_devices` - (Required) List of virtual network devices. Each element must have the following values:
    * **virtual_network_id** - (Required) The unique identifier of virtual network. Virtual network allows the virtual machine to communicate with the rest of your network, host machine, and other virtual machines. Use `ochk_virtual_network` data source for getting identifier by name.
* `additional_virtual_disks` - (Optional) List of additional virtual disks. Additional disk will be created on the same storage as the virtual machine configuration. Each element must have the following values: 
    * **controller_id** - (Required) The unique identifier of controller. The only supported value for now is "0".
    * **lun_id** - (Required) Number used to identify a logical unit. Set this to consecutive int numbers > 0. When updating, e.g. extending size, `lun_id` needs to be preserved.
    * **size_mb** - (Required) Size in megabytes of the additional disk. When updating, only disk extensions are supported (larger value).
    * **device_type** - (Optional) Type of the device. Only supported value for now: `SCSI`.
* `initial_user_name` (Optional) Only for virtual machines created from OVF to set ssh-key or ip address.

## Attribute Reference

The following attributes are exported in addition to above arguments:
* `created_by` - Who created this resource.
* `created_at` - When this resource was created.
* `modified_by` - Who last modified this resource. 
* `modified_at` - When last modification occurred.  
* `ip_address` - Default virtual machine ip address