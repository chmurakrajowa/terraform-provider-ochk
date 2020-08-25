---
page_title: "Service Data Source"
---

# Service Data Source

Data Source for services. 

## Example Usage

```hcl
data "ochk_service" "ssh" {
  display_name = "SSH"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of service.

## Attribute Reference

The following attributes are exported:
 * `display_name` - Display name. 
    
 
