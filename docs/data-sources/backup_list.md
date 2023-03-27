---
page_title: "Backup Plan Data Source"
---

# Backup List Data Source

Data Source for getting backup list by display name for selected backup plan.

## Example Usage

```hcl
data "ochk_backup_plan" "backup_plan" {
  display_name = "Platinium"
}

data "ochk_backup_list" "backup_list" {
  display_name = "example_backup_list"
  backup_plan_id = data.ochk_backup_plan.backup_plan.id
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of backup list.
* `backup_plan_id` - (Required) Backup plan id. You can get backup_plan_id from [Backup Plans](backup_plans.md) or [Backup Plan](backup_plan.md) Data Source.
