---
page_title: "Auto NAT Data Source"
---

# Auto NAT Data Source

Data Source for getting Auto NAT by display name.
Auto NAT is an auto-created DNAT for internet access from selected virtual network.

## Example Usage

```hcl
data "ochk_auto_nat" "auto-nat" {
  display_name = "example.autonat"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name for Auto NAT.

## Attribute Reference

The following attributes are exported in addition to above arguments:

* `description` - Auto NAT description.
* `virtual_network_id` - Virtual network id.
* `vrf_id` - VRF id.
* `nat_type` - NAT type: `AUTO` for auto nats.
* `action` - Action: `DNAT`
* `enabled` - NAT enabled (True) or disabled (False).
* `priority` - Rule priority (0-2147483640).
* `service_id` - Service id.
* `source_network` - Source network CIDR or ip address.
* `translated_network` - Translated network CIDR or ip address.
* `destination_network` - Destination network CIDR or ip address.
* `translated_ports` - Translated port for DNAT. Empty for Auto NATs.
* `created_by` - Who created this resource.
* `created_at` - When this resource was created.
* `modified_by` - Who last modified this resource.
* `modified_at` - When last modification occurred.

