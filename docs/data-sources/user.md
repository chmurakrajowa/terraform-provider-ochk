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
 * `name` - User name. 
 * `email_address` - The user's e-mail address. 
 * `description` - The short description of the user.
 * `first_name` - The user's first name. 
 * `last_name` - The user's last name. 
 * `disabled` - Disabled user account: true if user is disabled. 
 * `locked` - Locked user account: true if user is locked. User account can be locked after a certain number of invalid login attempts. 
 * `user_principal_name` - Specifies the name of the user in the form name@domain.
 * `principal_id` - The identifier of the principal. 
 * `principal_name` - Principal name. 
 * `principal_domain` - principal domain, possible values: CUSTOM_GROUP, GROUP, SSO_GROUP, USER. 
  
    
 
