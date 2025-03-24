---
page_title: "Security Group Data Source"
---

# Security Group Data Source

Data Source for reading security groups by display name. 

## Example Usage

```hcl
data "ochk_security_group" "{{ .DataSourceName}}" {
  display_name = "security_group_name"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of security group.

## Attribute Reference

The following attributes are exported:

* `project_id` - Project id to which security group is assigned.
* `members` - Members that are assigned to the security group.
  Each entry has following values:
   * **type**: Type of security group member, values: `IPCOLLECTION`, `VIRTUAL_MACHINE`
   * **display_name**: Display name of security group member.
   * **id**: Resource id depending on the security group type selection.
* `created_by` - Who created this resource.
* `created_at` - When this resource was created.
* `modified_by` - Who last modified this resource.
* `modified_at` - When last modification occurred.
