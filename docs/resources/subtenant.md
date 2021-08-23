---
page_title: "Subtenant Resource"
---

# Subtenant Resource

Resource for managing subtenants (business groups) to collect consumers, often corresponding to a organization's line of business, department, or other organizational unit. Each subtenant has a reservation (storage and memory) that determine on which compute resources the machines that this group requested can be provisioned. It also has its own authorization groups, which determine access to functionality only in a given subtenant. Every subtenant must have at least one manager, who monitors the resource use for the group and control the users role. A user can be a member of more than one subtenant, and can have different roles in different groups.
 

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
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name for the subtenant. Updates to this attribute forces recreate.
* `email` - (Required) Email address of user that should receive capacity alert notifications.
* `description` - (Required) Description.
* `memory_reserved_size_mb` - (Required) Memory reservation size in megabytes. Should be greater than 10000 MB (10GB).
* `storage_reserved_size_gb` - (Required) Storage reservation size in gigabytes. Should be greater than 100 GB.
* `users` - (Required) List of user identifiers. After creating a subtenant, these users will be automatically added to the permission group: Permissions Manager. Use `ochk_user` data source for finding identifiers by name. Updates to this attribute forces recreate.
  
## Attribute Reference

No additional attributes are exported. 
 
