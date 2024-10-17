---
page_title: "VRFs Data Source"
---

# VRFs Data Source

Data Source for getting VRFs (Virtual Routing and Forwarding) list.

## Example Usage

```hcl
data "ochk_vrfs" "{{ .DataSourceName}}" {
}
```

## Argument Reference

No additional arguments are required.

## Attribute Reference

The following attributes are exported:
* `vrfs` - VRFs list. Each entry has the following values:
    * **display_name** - VRF display name.
    * **vrf_id** -VRF id.