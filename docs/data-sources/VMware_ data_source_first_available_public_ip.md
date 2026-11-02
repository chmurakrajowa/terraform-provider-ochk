---
page_title: "First available public ip Data Source"
---

# First available public ip Data Source

Data Source for getting first available public ip

## Example Usage

```hcl
data "ochk_available_public_ip" "{{ .DataSourceName}}" {
}
```

## Argument Reference

No additional arguments are required.

## Attribute Reference

The following attributes are exported:
* `public_ip_address` - Firts available public IP address with the following values:
    * `ip_address` - Available public IP address.
    * `ip_address_id` - Public IP address id.