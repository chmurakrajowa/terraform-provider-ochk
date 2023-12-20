---
page_title: "Account Data Source"
---

# Billing Account Data Source

Data Source for getting billing account by name.

## Example Usage

```hcl

data "ochk_billing_account" "acct1" {
  display_name = "example_account"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact name of billing account. You can get all account names from [Billing Accounts](billing_accounts.md) Data Source.

## Attribute Reference

[The following attributes are exported:
* `account_description` - The short description of the billing account.
* `discount` - Discount if granted
* `alarms` - Alarms if enabled
* `cost` - Total cost
* `projects` - Projects to which account are assigned.
  Each entry has following values:
    * **project_id"**: Project id.
    * **display_name**: Display name of project.
   
    