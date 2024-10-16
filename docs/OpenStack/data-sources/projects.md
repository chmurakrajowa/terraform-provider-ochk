---
page_title: "Projects Data Source"
---

# Projects Data Source

Data Source for getting projects list.

## Example Usage

```hcl
data "ochk_projects" "{{ .DataSourceName}}" {
}
```

## Argument Reference

No additional arguments are required.

## Attribute Reference

## Attribute Reference

The following attributes are exported:
* `projects` - Projects list. Each entry has the following values:
    * **display_name** - Project display name.
    * **project_id** - Project id.
    
 
