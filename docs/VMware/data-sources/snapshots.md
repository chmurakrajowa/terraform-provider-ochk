---
page_title: "Snapshots Data Source"
---

# Snapshots Data Source

Data Source for getting snapshots list.

## Example Usage

```hcl
data "ochk_virtual_machine" "vm" {
    display_name = "vm_name"
}

data "ochk_snapshots" "{{ .DataSourceName}}" {
    virtual_machine_id = data.ochk_virtual_machine.vm.id
}
```
## Argument Reference

The following arguments are supported:
* `virtual_machine_id` - (Required) Id of virtual machine from which the Snapshot is created.

## Attribute Reference

The following attributes are exported:
* `snapshots` - Snapshots list with additional attributes per snapshot:
    * **display_name** - Snapshot name.
    * **snapshot_id** - Snapshot id.
    * **virtual_machine_id** - Id of virtual machine for which Snapshot is created.
