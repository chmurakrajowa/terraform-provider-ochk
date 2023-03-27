---
page_title: "Service Data Source"
---

# Service Data Source

Data Source for getting services by display name.

## Example Usage

```hcl
data "ochk_service" "ssh" {
  display_name = "SSH"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Display name of service.

## Attribute Reference

The following attributes are exported:
* `created_by` - Who created this resource.
* `created_at` - When this resource was created.
* `modified_by` - Who last modified this resource.
* `modified_at` - When last modification occurred.
    
 
