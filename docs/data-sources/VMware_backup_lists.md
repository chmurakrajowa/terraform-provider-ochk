---
page_title: "Backup Lists Data Source"
---

# Backup Lists Data Source

Data Source for getting all backup lists for selected backup plan.

## Example Usage

```hcl
data "ochk_backup_plan" "backup_plan" {
  display_name = "Platinium"
}

data "ochk_backup_lists" "{{ .DataSourceName}}" {
  backup_plan_id = data.ochk_backup_plan.backup_plan.id
}
```

## Argument Reference

The following arguments are supported:

* `backup_plan_id` - (Required) Backup plan id. You can get backup_plan_id from [Backup Plans](backup_plans.md) or [Backup Plan](backup_plan.md) Data Source.

## Attribute Reference

The following attributes are exported:
* `backup_lists` - Backup lists. Each entry has the following values:
    * **display_name** - Backup list name.
    * **backup_list_id** - Backup list id.
