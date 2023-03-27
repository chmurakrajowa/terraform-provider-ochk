---
page_title: "Router Resource"
---

# Router Resource

Resource for managing Tier-1 logical router providing routing functionality in a virtual environment. 
You can create more then one logical router to segment your network.

## Example Usage

```hcl

data "ochk_vrf" "rt" {
  display_name = "router-T0-name-display-name" 
}

data "ochk_project" "project" {
  display_name = "project_display_name"
}

resource "ochk_vpc" "router" {
  display_name = "router_name"
  vrf_id = data.ochk_router.rt.id
  project_id = data.ochk_project.project.id
  folder_path = "/"
}

```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) The Tier-1 router name (between 3 and 15 characters).
* `vrf_id` - (Required) VRF id. (default vrf name : T0)
* `project_id` - (Required) Project identifier, use `ochk_project` data source for getting identifier by name.
* `folder_path` - (Optional) Folder path for vpc. Default `/`.


## Attribute Reference


The following attributes are exported in addition to above arguments:
* `router_type` - Router type
* `created_by` - Who created this resource.
* `created_at` - When this resource was created.
* `modified_by` - Who last modified this resource.
* `modified_at` - When last modification occurred.
