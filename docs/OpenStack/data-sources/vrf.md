---
page_title: "VRF Data Source"
---

# VRF Data Source

Data Source for getting VRF (Virtual Routing and Forwarding) router by display name.

## Example Usage

```hcl
data "ochk_vrf" "{{ .DataSourceName}}" {
  display_name = "vrf_name"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Display name of VRF.

## Attribute Reference

The following attributes are exported:
 * `display_name` - Router name.
 * `created_by` - Who created this resource.
 * `created_at` - When this resource was created.
 * `modified_by` - Who last modified this resource. 
 * `modified_at` - When last modification occurred. 
    
 
