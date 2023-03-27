---
page_title: "Virtual Networks Data Source"
---

# Virtual Networks Data Source

Data Source for getting virtual networks list.

## Example Usage

```hcl
data "ochk_virtual_networks" "vnets" {
}
```

## Argument Reference

No additional arguments are required.

## Attribute Reference

The following attributes are exported:
* `virtual_networks` - Virtual networks list. Each entry has the following values:
    * **display_name** - Virtual network display name.
    * **virtual_netwok_id** - Virtual network id.
    * **folder_path** - Virtual network folder path, default: `/`.
    * **ipam_enabled** - The IP address management (IPAM) to discover IP address and Domain Name System (DNS) servers on the network and manage them.
    * **vpc_id** - VPC id.
