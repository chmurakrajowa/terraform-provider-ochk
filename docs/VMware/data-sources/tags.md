---
page_title: "Tags Data Source"
---

# Tags Data Source

Data Source for getting tags list.

## Example Usage

```hcl
data "ochk_tags" "{{ .DataSourceName}}" {
}
```
## Argument Reference

No additional arguments are required.

## Attribute Reference

The following attributes are exported:
* `tags` - Tags list with additional attributes per tag:
    * **display_name** - Tag name.
    * **tag_id** - Tag id.
    * **project_id** - Project id to which Tag is assigned.
