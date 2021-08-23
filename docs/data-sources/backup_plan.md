---
page_title: "Backup Plan Data Source"
---

# Backup Plan Data Source

Data Source for getting backup plan by display name.

## Example Usage

```hcl
data "ochk_backup_plan" "backup_plan" {
  display_name = "example.backup_plan"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of backup plan.

 
