---
page_title: "Tag Data Source"
---

# Tag Data Source

Data Source for getting tag by name.

## Example Usage

```hcl

data "ochk_project" "project" {
  display_name = "example_project"
}

data "ochk_tag" "os1" {
  display_name = "example_tag"
  project_id = data.ochk_project.project.id
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact name of tag. You can get all tag names from [Tags](tags.md) Data Source.
* `project_id` - (Required) Project id to which Tag is assigned.

## Attribute Reference

[The following attributes are exported:
* `display_name` -  Tag name.
* `project_id` - Project id to which Tag is assigned.

    
 
