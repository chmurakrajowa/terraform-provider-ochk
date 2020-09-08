---
page_title: "User Data Source"
---

# User Data Source

Data Source for getting users by their name. 

## Example Usage

```hcl
data "ochk_user" "admin" {
  name = "admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Exact name of a user.

## Attribute Reference

The following attributes are exported:
 * `name` - name of a user. 
 * `email_address` - email address. 
 * `description` - description.
 * `first_name` - first name. 
 * `last_name` - last name. 
 * `disabled` - true if user is disabled. 
 * `locked` - true if user is locked. 
 * `user_principal_name` - principal name. 
 * `principal_id` - principal identified. 
 * `principal_name` - principal name. 
 * `principal_domain` - principal domain, possible values: CUSTOM_GROUP, GROUP, SSO_GROUP, USER. 
  
    
 
