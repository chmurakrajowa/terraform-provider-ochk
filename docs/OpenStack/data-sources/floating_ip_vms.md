---
page_title: "Floating IP VMs Data Source"
---

# Floating IP VMs Data Source

Data Source for getting floating ip vms list.

## Example Usage

```hcl
data "ochk_floating_vms" "{{ .DataSourceName}}" {
}
```

## Argument Reference

No additional arguments are required.

## Attribute Reference

The following attributes are exported:
* `floating_ip_vms` - Public IP VMs list. Each entry has the following values:
* **virtual_machine_id** - Virtual machine id..
* **virtual_machine_name** - Virtual machine display name.
* **osc_port_id** - Port id.
* **mac_address** - Mac address.
* **network_name** - Virtual network display name.
* **ip_address** - IP address.

