---
page_title: "Account Resource"
---

# Billing Account Resource

Resource for managing accounts.

## Example Usage

```hcl
resource "ochk_billing_account" "acct1" {
  display_name = "tf-acct-billaccount"
 
  projects {
    project_id = "63182099-5396-48f5-9474-c8031e08f25a"
  }

}
```

It is not recommended to directly use resource identifiers. To avoid that prefer to use data sources:
```hcl
data "ochk_project" "project" {
  display_name = "project-example"
}


resource "ochk_billing_account" "acct_ex1" {
  display_name = "acct-example"

  projects {
    project_id = data.ochk_project.project.id
  }
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Billing Account name.
* `account_description` - The short description of the account.
* `discount` - Discount if granted
* `alarms` - Alarms if enabled
* `projects` - Projects to which account is assigned.
  Each entry has following values:
    * **project_id"**: Project id.
    * **display_name**: Display name of project.
  
## Attribute Reference

The following attributes are exported in addition to above arguments:
* `cost` - Total cost
