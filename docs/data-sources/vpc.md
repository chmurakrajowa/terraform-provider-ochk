---
page_title: "VPC Data Source"
---

# VPC Data Source

Data Source for getting VPC (Virtual Private Cloud) router by display name and VRF id.

## Example Usage

```hcl
data "ochk_vrf" "vrf" {
  display_name = "T0"
}

data "ochk_vpc" "vpc" {
  vrf_id = data.ochk_vrf.vrf.id
  display_name = "VPC1"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Display name of VPC.
* `vrf_id` - (Required) VRF id.

## Attribute Reference

The following attributes are exported:
* `project_id` - Project id to which VPC is assigned.
* `folder_path` - VPC folder path, default: `/`.
* `created_by` - Who created this resource.
* `created_at` - When this resource was created.
* `modified_by` - Who last modified this resource.
* `modified_at` - When last modification occurred. 
    
 
