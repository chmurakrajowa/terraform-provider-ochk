---
page_title: "Manual NAT Data Resource"
---

# Manual NAT Data Resource

Resource for managing Manual NATs. There are 3 types of Manual NATs: `DNAT`, `SNAT`, `NO_SNAT`

## Example Usage

`DNAT` example:

```hcl
data "ochk_vrf" "vrf" {
  display_name = "vrf_name"
}

data "ochk_service" "http" {
  display_name = "HTTP"
}

resource "ochk_manual_nat" "{{ .ResourceName}}"  {
    display_name = "manualnat_name"
    action = "DNAT"
    enabled = true
    vrf_id = data.ochk_vrf.vrf.id
    source_network = "10.10.0.10"
    translated_network = "10.12.12.12"
    destination_network = "45.130.209.132"
    priority = 35
    service_id = data.ochk_service.http.id
    translated_ports = "80"
}
```

`SNAT` example:

```hcl
data "ochk_vrf" "vrf" {
  display_name = "vrf_name"
}

resource "ochk_manual_nat" "snat" {
    display_name = "manualnat_name"
    action = "SNAT"
    enabled = true
    vrf_id = data.ochk_vrf.vrf.id
    source_network = "10.10.0.10"
    translated_network = "45.130.209.132"
    destination_network = "10.12.12.12"
    priority = 38
}
```

`NO_SNAT` example:

```hcl
data "ochk_vrf" "vrf" {
  display_name = "vrf_name"
}

resource "ochk_manual_nat" "no_snat" {
    display_name = "manualnat_name"
    action = "NO_SNAT"
    enabled = true
    vrf_id = data.ochk_vrf.vrf.id
    source_network = "10.10.0.10"
    destination_network = "10.12.12.12"
    priority = 39
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Display name for NAT.
* `description` - (Optional) Description.
* `enabled` - (Required) NAT enabled (True) or disabled (False).
* `vrf_id` - (Required) VRF id.
* `action` - (Required) Action: `DNAT`, `SNAT` or `NO_SNAT`
* `priority` - (Optional) Rule priority (0-2147483640).
* `service_id` - (Optional) Service id.
* `destination_network` - (Required) Destination network CIDR or ip address.
* `source_network` - (Required) Source network CIDR or ip address.
* `translated_network` - (Optional) Translated network CIDR or ip address.
* `translated_ports` - (Optional) Translated port for DNAT.

## Attribute Reference

The following attributes are exported in addition to above arguments:

* `nat_type` - NAT type: `MANUAL` for manual nats.
* `created_by` - Who created this resource.
* `created_at` - When this resource was created.
* `modified_by` - Who last modified this resource.
* `modified_at` - When last modification occurred.

