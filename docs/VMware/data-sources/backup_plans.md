---
page_title: "Backup Plans Data Source"
---

# Backup Plans Data Source

Data Source for getting backup plans list.

## Example Usage

```hcl
data "ochk_backup_plans" "{{ .DataSourceName}}" {
}
```

## Argument Reference

No additional arguments are required.

## Attribute Reference

The following attributes are exported:
* `backup_plans` - Backup plans list. Each entry has the following values:
  * **display_name** - Backup plan name.
  * **backup_plan_id** - Backup plan id.
