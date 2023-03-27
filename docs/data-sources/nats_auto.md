---
page_title: "Auto NATs Data Source"
---

# Auto NATs Data Source

Data Source for getting Auto NATs list.

## Example Usage

```hcl
data "ochk_auto_nats" "auto-nats" {
}
```

## Argument Reference

No additional arguments are required.

## Attribute Reference

The following attributes are exported:
* `auto_nats` - Auto NATs list. Each entry has the following values:
  * **aut_nat_id** - Id of Auto NAT
  * **display_name** - Auto NAT display name
  * **virtual_network_id** -  Virtual network id.
  * **enabled** - NAT enabled (True) or disabled (False).
  * **vrf_id** - VRF id.

