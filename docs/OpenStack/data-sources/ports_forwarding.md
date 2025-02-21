---
page_title: "Ports Forwarding Data Source"
---

# Ports Forwarding Data Source

Data Source for getting Ports Forwarding list.

## Example Usage


```hcl
data "ochk_floating_ip_address" "fip01" {
  display_name = "fip_test01"
}

```hcl
data "ochk_port_forwarding" "{{ .DataSourceName}}" {
  floating_ip_id = data.ochk_floating_ip_address.fip01.id
}
```

## Argument Reference

The following arguments are supported:

* `floating_ip_id` - ID of floating ip.

## Attribute Reference

The following attributes are exported:
* `ports_forwarding` - Ports Forwarding list. Each entry has the following values:
  * **port_forwarding_id** - Id of Port Forwarding
  * **display_name** - Exact display name for port forwarding rule.
  * **floating_ip_id** - ID of floating ip
  * **internal_port_id** - ID of internal port.

