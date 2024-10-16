---
page_title: "Virtual Machine Data Source"
---

# Virtual Machine Data Source

Data Source for getting virtual machines by display name.

## Example Usage

```hcl
data "ochk_virtual_machine" "{{ .DataSourceName}}" {
  display_name = "virtual_machine_name"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of virtual Machine.

## Attribute Reference

The following attributes are exported:
* `project_id` - Project id to which virtual machine is assigned.
* `folder_path` - Virtual machine folder path, default: `/`.
 
