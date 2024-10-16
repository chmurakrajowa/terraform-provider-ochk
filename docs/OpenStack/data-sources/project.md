---
page_title: "Project Data Source"
---

# Project Data Source

Data Source for getting projects by display name. 

## Example Usage

```hcl
data "ochk_project" "{{ .DataSourceName}}" {
  display_name = "project_name"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact name of project.

## Attribute Reference

The following attributes are exported:
* `display_name` - Project name.
* `vrf_id` - VRF id to which project is assigned
* `description` - The short description of the project.
* `limits_enabled` - Enabled (True) or disabled (False) project reservations.
* `memory_reserved_size_mb` - Memory reservation size in megabytes.
* `storage_reserved_size_gb` - Storage reservation size in gigabytes.
* `vcpu_reserved_quantity` -  VCPU reserved quantity. 
    
 
