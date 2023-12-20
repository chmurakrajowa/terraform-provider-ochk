---
page_title: "Resource Snapshot"
---

# Tag Resource

Resource for managing Snapshots. 

## Example Usage

```hcl
data "ochk_virtual_machine" "vm" {
  display_name = "example_vm"
}

resource "ochk_snapshot" "snapshot" {
    display_name = "example-snapshot"
    virtual_machine_id = data.ochk_virtual_machine.vm.id
}
```

## Argument Reference
The following arguments are supported:
* `display_name` - (Required) The name of snapshot.
* `virtual_machine_id` - (Required) Id of virtual machine for which Snapshot is created.
