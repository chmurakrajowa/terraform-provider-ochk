---
page_title: "Accounts Data Source"
---

# Accounts Data Source

Data Source for getting billing accounts list.

## Example Usage

```hcl
data "ochk_billing_accounts" "{{ .DataSourceName}}" {
}
```

## Argument Reference

No additional arguments are required.

## Attribute Reference

The following attributes are exported:

* `display_name` - Exact name of billing account
* `account_description` - The short description of the billing account.
* `discount` - Discount if granted
* `alarms` - Alarms if enabled
* `cost` - Total cost
* `projects` - Projects to which account is assigned.
  Each entry has following values:
    * **project_id"**: Project id.
    * **display_name**: Display name of project.
 
