---
page_title: "Snapshot Data Source"
---

# Snapshot Data Source

Data Source for getting snapshot by name.
## Example Usage

```hcl

data "ochk_virtual_machine" "vm" {
  display_name = "vm_name"
}

data "ochk_snapshot" "{{ .DataSourceName}}" {
  display_name = "snapshot_name"
  virtual_machine_id = data.ochk_virtual_machine.vm.id
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact name of snapshot. You can get all snapshots names from [Snapshots](snapshots.md) Data Source.
* `virtual_machine_id` - (Required) Id of virtual machine from which the Snapshot is created.

## Attribute Reference

The following attributes are exported:
* `snapshot_description` - The short description of the snapshot
* `parent_id` - Id of parent snapshot
* `child_id` - Id of child snapshot
* `power_state` - Powered state; for snapshots with RAM, the power_state is "poweredOn"
