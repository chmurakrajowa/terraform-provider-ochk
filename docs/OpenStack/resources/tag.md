---
page_title: "Resource Tag"
---

# Tag Resource

Resource for managing Tags. Tags can be assigned to virtual machines for billing purposes.

## Example Usage

```hcl
data "ochk_project" "project" {
  display_name = "project_name"
}

resource "ochk_tag" "{{ .ResourceName}}" {
    display_name = "tag_name"
    project_id = data.ochk_project.project.id
}
```

## Argument Reference
The following arguments are supported:
* `display_name` - (Required) The name of tag.
* `project_id` - (Required) Project id to which Tag is assigned.
