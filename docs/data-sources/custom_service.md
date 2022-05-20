---
page_title: "Custom Service Data Source"
---

# Custom Service Data Source

Data Source for getting custom services by display name.

## Example Usage

```hcl
data "ochk_custom_service" "web-servers" {
  display_name = "web-servers"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of custom service.

## Attribute Reference

* `ports` - Definition of communication rule between two machines.
  Each entry has the following values:
  * **protocol**: Set of rules that determine how data is transmitted between different devices in the same network. Allowed values: `TCP`, `UDP`.
  * **source**: List of source ports of a service.
  * **destination**: List of destination ports of a service.

The following attributes are exported in addition to above arguments:
* `created_by` - Who created this resource.
* `created_at` - When this resource was created.
* `modified_by` - Who last modified this resource.
* `modified_at` - When last modification occurred.
