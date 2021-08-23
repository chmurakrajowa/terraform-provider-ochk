---
page_title: "System Tag Data Source"
---

# System Tag Data Source

Data Source for getting system tag by name.

## Example Usage

```hcl
data "ochk_system_tag" "os1" {
  display_name = var.system_tag_os
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact name of system tag.


    
 
