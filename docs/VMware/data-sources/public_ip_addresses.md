---
page_title: "Backup Plans Data Source"
---

# Backup Plans Data Source

Data Source for getting allocated Public IP Addresses list.

## Example Usage

```hcl
data "ochk_public_ip_addresses" "{{ .DataSourceName}}" {
}
```

## Argument Reference

No additional arguments are required.

## Attribute Reference

The following attributes are exported:
* `public_ip_addresses` - Public IP Addresses list. Each entry has the following values:
    * **display_name** - Public IP address name.
    * **public_ip_address_id** - Public IP address id.
    * **ip_address** - Public IP address.

