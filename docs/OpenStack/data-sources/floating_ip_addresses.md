---
page_title: "Floating ip addresses Data Source"
---

# Floating ip addresses Data Source

Data Source for getting allocated Public IP Addresses list.

## Example Usage

```hcl
data "ochk_floating_ip_addresses" "{{ .DataSourceName}}" {
}
```

## Argument Reference

No additional arguments are required.

## Attribute Reference

The following attributes are exported:
* `floating_ip_addresses` - Public IP Addresses list. Each entry has the following values:
    * **display_name** - Public IP address name.
    * **public_ip** - Public IP address.
    * **floating_ip_id** - Floating IP address id.

