---
page_title: "System Tag Data Source"
---

# System Tag Data Source

Data Source for getting system tag by name.

## Example Usage

```hcl
data "ochk_system_tag" "os1" {
  display_name = "example_system_tag"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact name of system tag.


    
 
