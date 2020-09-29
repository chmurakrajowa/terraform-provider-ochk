---
page_title: "Subtenant Data Source"
---

# Subtenant Data Source

Data Source for getting subtenants by name. 

## Example Usage

```hcl
data "ochk_subtenant" "subtenant1" {
  name = "subtenant-name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Exact name of subtenant.

## Attribute Reference

The following attributes are exported:
* `name` - Subtenant name.
* `email` - Email addresses of users that must receive capacity alert notifications.
* `description` - The short description of the subtenant
* `memory_reserved_size_mb` - Memory reservation size in megabytes.
* `storage_reserved_size_gb` - Storage reservation size in gigabytes.
* `users` - List of user identifiers assigned to the subtenat.
* `networks` - List of network identifiers assigned to the appropriate subtenant 
  
    
 
