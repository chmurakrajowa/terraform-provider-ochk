---
page_title: "Folders Data Source"
---

# Folders Data Source

Data Source for getting folders list for selected project id.

## Example Usage

```hcl
data "ochk_project" "project" {
  display_name = "project_name"
}

data "ochk_folders" "{{ .DataSourceName}}" {
  project_id = data.ochk_project.project.id
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) Project id.

## Attribute Reference

The following attributes are exported:
* `folders` - Folders list. Each entry has the following values:
    * **folder_name** - Folder name.
    * **folder_path** - Full folder path with folder name.
    * **folder_id** - Folder id.
