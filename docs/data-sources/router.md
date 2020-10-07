---
page_title: "Router Data Source"
---

# Router Data Source

Data Source for getting logical routers by display name. 

## Example Usage

```hcl
data "ochk_router" "rt" {
  display_name = "router-name-display-name"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of router.

## Attribute Reference

The following attributes are exported:
 * `display_name` - Logical router name. 
 * `type` - Type of the logical router e.g. TIER0, TIER1
 * `created_by` - Who created this resource.
 * `created_at` - When this resource was created.
 * `modified_by` - Who last modified this resource. 
 * `modified_at` - When last modification occurred. 
    
 
