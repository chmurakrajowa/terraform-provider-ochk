---
page_title: "Backup Plan Data Source"
---

# Backup List Data Source

Data Source for getting backup list by display name.

## Example Usage

```hcl
data "ochk_backup_list" "backup_list" {
  display_name = "example_backup_list"
  backup_plan_id = "data_ochk_backup_plan.backup_plan.id"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of backup list.
* `backup_plan_id` - (Required) Backup plan id, for instance: backup_plan = "Platinium",
  backup_list = "Platinium (1 month / 8h)"
