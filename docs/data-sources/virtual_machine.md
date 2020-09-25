---
page_title: "Virtual Machine Data Source"
---

# Virtual Machine Data Source

Data Source for virtual machines. 

## Example Usage

```hcl
data "ochk_virtual_machine" "sg" {
  display_name = "virtual-machine-display-name"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of virtual Machine.

## Attribute Reference

The following attributes are exported:
 * `display_name` - Display name.
 * `host_id` - Host identifier. 
    
 