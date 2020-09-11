---
page_title: "Subtenant Data Source"
---

# Subtenant Data Source

Data Source for subtenants. 

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

* `name` - Name of subtenant.
* `email` - Email address.
* `description` - Description.
* `memory_reserved_size_mb` - Memory reservation size in megabytes.
* `storage_reserved_size_gb` - Storage reservation size in gigabytes.
* `users` - List of user identifiers.
* `networks` - List of network identifiers.
  
    
 
