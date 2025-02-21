---
page_title: "Virtual Machines Data Source"
---

# Virtual Machines Data Source

Data Source for getting virtual machines list.

## Example Usage

```hcl
data "ochk_virtual_machines" "{{ .DataSourceName}}" {
}
```

## Argument Reference

No additional arguments are required.

## Attribute Reference

The following attributes are exported:
* `virtual_machines` - Virtual machines list. Each entry has the following values:
    * **display_name** -Virtual machine display name.
    * **virtual_machine_id** - Virtual machine id.
    * **project_id** - Project id to which virtual machine is assigned.
    * **folder_path** - Virtual machine folder path, default: `/`.
 
