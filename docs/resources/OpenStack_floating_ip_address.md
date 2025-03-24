---
page_title: "Floating IP Address Resource"
---

# Floating IP Address Resource

Resource for managing accounts.

## Example Usage

```hcl
resource "ochk_floating_ip_address" "{{ .ResourceName}}" {
  display_name = "floating_ip_name"
  description = "short description"
}

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Floating IP address name.
* `description` - Short description
* `vm_name` - Related virtual machine name.
* `vm_port_id` - Related virtual machine.
* `vm_fixed_ip` - Related virtual machine address ip.
  
## Attribute Reference

The following attributes are exported in addition to above arguments:
* `public_adress` - Public IP address
