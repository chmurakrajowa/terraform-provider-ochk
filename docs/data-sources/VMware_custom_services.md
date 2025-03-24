---
page_title: "Custom Services Data Source"
---

# Custom Services Data Source

Data Source for getting custom services list.

## Example Usage

```hcl
data "ochk_custom_services" "{{ .DataSourceName}}" {
}
```
## Argument Reference

No additional arguments are required.

## Attribute Reference

The following attributes are exported:
* `custom_services` - Custom services list. Each entry has the following values:
    * **display_name** - Custom service name.
    * **custom_service_id** - Custom service id.
    * **project_id** - Project id to which custom service is assigned.
