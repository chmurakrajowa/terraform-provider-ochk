---
page_title: "Port Forwarding Data Resource"
---

# Port Forwarding Data Resource

Resource for managing Port Forwarding.

## Example Usage

`DNAT` example:

```hcl
data "ochk_floating_address" "f1" {
display_name = "floating_ip_name"
}

resource "ochk_port_forwarding" "{{ .ResourceName}}" {
    floating_ip_id = data.ochk_floating_ip_address.f1.id
    display_name = "port_forwarding_name"
    protocol = "tcp"
    internal_port_id = "11752b30-9749-4534-8ec1-2b0e9670b92b"
    internal_ip_address = "192.168.2.153"
    internal_port = 1219
    public_address = "203.0.113.29"
    external_port = 1221
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name for port forwarding rule.
* `floating_ip_id` - (Required) ID of floating ip.
* `description` - The short description of the port forwarding.
* `public_address` - (Required) Public IP address.
* `external_port` - External port.
* `external_port_range` - External port range (1-65535)
* `internal_ip_address` - (Required) Rule priority (0-2147483640).
* `internal_port` - Internal port.
* `internal_port_id` - ID of internal port
* `internal_port_range` - Internal port range (1-65535)
* `protocol` - (Required) Protocol type, values: `TCP`, `UPD`.

## Attribute Reference

The following attributes are exported in addition to above arguments:

* `created_by` - Who created this resource.
* `created_at` - When this resource was created.
* `modified_by` - Who last modified this resource.
* `modified_at` - When last modification occurred.

