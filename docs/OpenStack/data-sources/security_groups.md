---
page_title: "Security Groups Data Source"
---

# Security Groups Data Source

Data Source for reading security groups list.

## Example Usage

```hcl
data "ochk_security_groups" "{{ .DataSourceName}}" {
}
```

## Argument Reference

No additional arguments are required.

## Attribute Reference

The following attributes are exported:
* `security_groups` - Security groups list. Each entry has the following values:
  * **display_name** - Security group display name.
  * **security_group_id** - Security group id.
