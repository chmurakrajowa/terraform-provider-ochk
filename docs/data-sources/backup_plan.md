---
page_title: "Backup Plan Data Source"
---

# Backup Plan Data Source

Data Source for getting backup plan by display name.

## Example Usage

```hcl
data "ochk_backup_plan" "backup_plan" {
  display_name = "Platinium"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of backup plan. You can get full display names list from [Backup Plans](backup_plans.md) Data Source.

## Attribute Reference

The following attributes are exported:
* `display_name` - Name of backup plan.
