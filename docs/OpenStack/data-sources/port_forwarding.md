---
page_title: "Port Forwarding Data Source"
---

# NAT Rule - Port Forwarding Data Source

Data Source for getting Port Forwarding rule by display name and floating ip id.

## Example Usage

```hcl
data "ochk_floating_ip_address" "fip01" {
    display_name = "floating_ip_name"
}

```hcl
data "ochk_port_forwarding" "{{ .DataSourceName}}" {
  floating_ip_id = data.ochk_floating_ip_address.fip01.id
  display_name = "port_forwarding_name"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name for port forwarding rule.
* `floating_ip_id` - ID of floating ip.

## Attribute Reference

The following attributes are exported in addition to above arguments:

* `description` - The short description of the port forwarding.
* `public_address` - Public IP address.
* `external_port` - External port.
* `external_port_range` - External port range (1-65535)
* `internal_ip_address` - Rule priority (0-2147483640).
* `internal_port` - Internal port.
* `internal_port_id` - ID of internal port
* `internal_port_range` - Internal port range (1-65535)
* `protocol` - Protocol type, values: `TCP`, `UPD`.
* `created_by` - Who created this resource.
* `created_at` - When this resource was created.
* `modified_by` - Who last modified this resource.
* `modified_at` - When last modification occurred.

