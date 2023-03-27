---
page_title: "Custom service"
---

# Custom Service Resource

Resource for managing Custom Services to control the custom group of settings related to a specific network port or another network interface. Most of the typical services are defined in service inventory (`ochk_service` and `ochk_services` data sources), but we can also create custom services if we need to.

## Example Usage

```hcl
data "ochk_project" "project" {
  display_name = "project-example"
}

resource "ochk_custom_service" "default" {
  display_name = "web-servers"
  project_id = data.ochk.project.project.id

  ports {
    protocol = "TCP"
    source = ["80", "443"]
    destination = ["80", "443"]
  }

  ports {
    protocol = "TCP"
    source = ["8080-8090"]
    destination = ["1-65535"]
  }
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) The name of the custom service.
* `project_id` - (Required) Project id to which custom service is assigned.
* `ports` - (Required) Definition of communication rule between two machines. When updating ports preserve initial ordering. Always add new port definitions on the end of the list.
  Each entry must have the following values:
  * **protocol**: (Required) Set of rules that determine how data is transmitted between different devices in the same network. Allowed values: `TCP`, `UDP`.
  * **source**: (Required) List of source ports of a service. Allowed single port or port ranges, e.g. `443`, `8080-80890`. Use `1-65535` for matching all ports.
  * **destination**: (Required) List of destination ports of a service. Allowed single port or port ranges, e.g. `443`, `8080-80890`. Use `1-65535` for matching all ports.

## Attribute Reference 

The following attributes are exported in addition to above arguments: 
 * `created_by` - Who created this resource.
 * `created_at` - When this resource was created.
 * `modified_by` - Who last modified this resource. 
 * `modified_at` - When last modification occurred.
