---
page_title: "Resource Snapshot"
---

# Tag Resource

Resource for managing Snapshots. 

## Example Usage

```hcl
data "ochk_virtual_machine" "vm" {
  display_name = "vm_name"
}

resource "ochk_snapshot" "{{ .ResourceName}}" {
    display_name = "snapshot_name"
    virtual_machine_id = data.ochk_virtual_machine.vm.id
    ram = false
}
```

## Argument Reference
The following arguments are supported:
* `display_name` - (Required) The name of snapshot.
* `virtual_machine_id` - (Required) Id of virtual machine for which Snapshot is created.
* `ram` - Taking a snapshot with or without the ram
* `snapshot_description` - The short description of the snapshot

