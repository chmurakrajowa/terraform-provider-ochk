---
page_title: "Auto NAT Data Resource"
---

# Auto NAT Data Resource

Resource for managing Auto NATs.
Auto NAT is an auto-created DNAT for internet access from selected virtual network.

## Example Usage

```hcl
data "ochk_virtual_network" "vnet" {
  display_name = "virtual_network_name"
}

resource "ochk_auto_nat" "{{ .ResourceName}}"  {
  display_name = "auto-nat"
  virtual_network_id = data.ochk_virtual_network.vnet.id
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Display name for NAT.
* `virtual_network_id` - (Required) Virtual network id.

## Attribute Reference

The following attributes are exported in addition to above arguments:

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

