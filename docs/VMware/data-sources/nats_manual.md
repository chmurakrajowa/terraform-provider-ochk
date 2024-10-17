---
page_title: "Manual NATs Data Source"
---

# Manual NATs Data Source

Data Source for getting Auto NATs list.

## Example Usage

```hcl
data "ochk_manual_nats" "{{ .DataSourceName}}" {
}
```

## Argument Reference

No additional arguments are required.

## Attribute Reference

The following attributes are exported:
* `manual_nats` - Manual NATs list. Each entry has the following values:
  * **manual_nat_id** - Id of Auto NAT
  * **display_name** - Auto NAT display name
  * **action** - Auto NAT display name
  * **enabled** - NAT enabled (True) or disabled (False).
  * **vrf_id** - VRF id.
  * **source_network** - Source network CIDR or ip address
  * **destination_network** - Destination network CIDR or ip address

