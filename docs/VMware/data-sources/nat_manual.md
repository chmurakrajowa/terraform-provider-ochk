---
page_title: "Manual NAT Data Source"
---

# Manual NAT Data Source

Data Source for getting Manual NAT by display name.

## Example Usage

```hcl
data "ochk_manual_nat" "{{ .DataSourceName}}"{
  display_name = "maualnat_name"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name for Manual NAT.

## Attribute Reference

The following attributes are exported in addition to above arguments:

* `description` - Manual NAT description.
* `virtual_network_id` - Virtual network id.
* `vrf_id` - VRF id.
* `nat_type` - NAT type: `MANUAL` for manual nats.
* `action` - Action, e.g.: `DNAT`, `SNAT`, `NO_SNAT`.
* `enabled` - NAT enabled (True) or disabled (False).
* `priority` - Rule priority (0-2147483640).
* `service_id` - Service id.
* `source_network` - Source network CIDR or ip address
* `translated_network` - Translated network CIDR or ip address
* `destination_network` - Destination network CIDR or ip address
* `translated_ports` - Translated port for DNAT.
* `created_by` - Who created this resource.
* `created_at` - When this resource was created.
* `modified_by` - Who last modified this resource.
* `modified_at` - When last modification occurred.

