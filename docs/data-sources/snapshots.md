---
page_title: "Snapshots Data Source"
---

# Snapshots Data Source

Data Source for getting snapshots list.

## Example Usage

```hcl
data "ochk_snapshots" "snapshots" {
}
```
## Argument Reference

No additional arguments are required.

## Attribute Reference

The following attributes are exported:
* `snapshots` - Snapshots list with additional attributes per snapshot:
    * **display_name** - Snapshot name.
    * **snapshot_id** - Snapshot id.
    * **virtual_machine_id** - Id of virtual machine for which Snapshot is created.
