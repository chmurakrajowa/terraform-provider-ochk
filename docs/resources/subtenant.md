---
page_title: "Subtenant Resource"
---

# Subtenant Resource

Resource for managing subtenants (business groups). 

## Example Usage

```hcl
data "ochk_user" "admin" {
  name = "admin"
}

resource "ochk_subtenant" "bg-1" {
	name = "bg-test"
	email = "email@example.com"
	description = "Description"
	memory_reserved_size_mb = 30000
	storage_reserved_size_gb = 200
	users = [data.ochk_user.admin.id]
	networks = ["bd814070-18f3-4182-b2af-edaa72a50fee"]

	lifecycle {
		ignore_changes = [description]
	}
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of subtenant. Updates to this attribute forces recreate.
* `email` - (Required) Email address.
* `description` - (Required) Description.
* `memory_reserved_size_mb` - (Required) Memory reservation size in megabytes. Should be greater than 10000 MB (10GB).
* `storage_reserved_size_gb` - (Required) Storage reservation size in gigabytes. Should be greater than 100 GB.
* `users` - (Required) List of user identifiers. Use `ochk_user` data source for finding identifiers by name. Updates to this attribute forces recreate.
* `networks` - (Required) List of network identifiers.
  
## Attribute Reference

No additional attributes are exported. 
 
