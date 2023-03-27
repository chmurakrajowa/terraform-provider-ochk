---
page_title: "VPCs Data Source"
---

# VPCs Data Source

Data Source for getting VPCs (Virtual Private Cloud) list.

## Example Usage

```hcl
data "ochk_vrf" "vrf" {
  display_name = "T0"
}

data "ochk_vpc" "vpc" {
  vrf_id = data.ochk_vrf.vrf.id
}
```

## Argument Reference

The following arguments are supported:

* `vrf_id` - (Required) VRF id.

## Attribute Reference

The following attributes are exported:
* `backup_plans` - VPCs list. Each entry has the following values:
  * **vpc_id** - VPC id.
  * **vrf_id** - VRF id to which VPC is assigned.
  * **display_name** - Display name of VPC.
  * **project_id** - Project id to which VPC is assigned.
  * **folder_path** - VPC folder path, default: `/`.