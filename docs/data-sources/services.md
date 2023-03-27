---
page_title: "Services Data Source"
---

# Services Data Source

Data Source for getting services list.

## Example Usage

```hcl
data "ochk_services" "services" {
}
```
## Argument Reference

No additional arguments are required.

## Attribute Reference

The following attributes are exported:
* `custom_services` -  Services list. Each entry has the following values:
    * **display_name** - Service display name.
    * **custom_service_id** - Service id.
