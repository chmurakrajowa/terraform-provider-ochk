# <provider> Provider

Terraform Provider for managing cloud resources in OChK. 

## Example Usage

```hcl
provider "ochk" {
    host = var.host
    platform = var.platform
    api_key = var.api_key
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
* `platform` - platform name
* `api_key` - api key
* `debug_log_file` - output file to which all API calls will be logged. DEPRECATED: use TF_LOG=DEBUG instead.

Arguments can be set using environment variables:
* `host` - `OCHK_HOST`
* `platform` - `OCHK_PLATFORM`
* `api_key` - `OCHK_API_KEY`
    
## Debugging 

To debug API calls, set TF_LOG=DEBUG env var.