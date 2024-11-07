---
page_title: "Floating IP Data Source"
---

# Floating IP Data Source

Data Source for getting allocated Public IP Address.

## Example Usage

```hcl
data "ochk_floating_address" "{{ .DataSourceName}}" {
    display_name = "floating_ip_name"
}
```

## Argument Reference

The following arguments are supported:
* `display_name` - (Required) Floating IP name.

## Attribute Reference

The following attributes are exported in addition to above arguments:
* `description` - Short description
* `public_adress` - Public IP address
* `project_id` - Project id to which floating IP is assigned.
* `vm_name` - Related virtual machine name.
* `vm_port_id` - Related virtual machine.
* `vm_fixed_ip` - Related virtual machine address ip.
* `created_by` - Who created this resource.
* `created_at` - When this resource was created.
* `modified_by` - Who last modified this resource.
* `modified_at` - When last modification occurred.

