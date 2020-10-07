---
page_title: "Security Policy Data Source"
---

# Security Policy Data Source

Data Source for getting security policies by display name.  

## Example Usage

```hcl
data "ochk_security_policy" "sp" {
  display_name = "security-policy-display-name"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of security policy.

## Attribute Reference

The following attributes are exported:
 * `display_name` - Security policy name. 
 * `created_by` - Who created this resource.
 * `created_at` - When this resource was created.
 * `modified_by` - Who last modified this resource. 
 * `modified_at` - When last modification occurred.     
 
