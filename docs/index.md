# <provider> Provider

Terraform Provider for managing cloud resources in OChK. 

## Example Usage

```hcl
provider "ochk" {
    host = var.host
    tenant = var.tenant
    username = var.username
    debug_log_file = var.debug_log_file
}

data "ochk_virtual_machine" "vm" {
    display_name = "vm-display-name"
}

resource "ochk_security_group" "vm" {
    display_name = "tf-sg-vm"
    
    members {
        id = ochk_virtual_machine.vm.id
        type = "VIRTUAL_MACHINE"
    }
}
```

## Argument Reference

* `host` - API hostname
* `tenant` - tenant name
* `username` - user name
* `password` - password
* `debug_log_file` = var.debug_log_file

Arguments can be set using environment variables:
* `host` - `OCHK_HOST`
* `tenant` - `OCHK_TENANT`
* `username` - `OCHK_USERNAME`
* `password` - `OCHK_PASSWORD`
