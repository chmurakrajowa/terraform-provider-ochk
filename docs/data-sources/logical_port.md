---
page_title: "Logical port Data Source"
---

# Logical port Data Source

Data Source for getting logical ports by display name. 

## Example Usage

```hcl
data "ochk_logical_port" "lp" {
  display_name = "logical-port-display-name"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of logical port.

## Attribute Reference

The following attributes are exported:
 * `display_name` - Name of logical port. 
    
 
