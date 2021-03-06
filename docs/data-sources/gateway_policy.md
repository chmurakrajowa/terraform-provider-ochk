---
page_title: "Gateway Policy Data Source"
---

# Gateway Policy Data Source

Data Source for getting gateway policies by display name. 

## Example Usage

```hcl
data "ochk_gateway_policy" "gp" {
  display_name = "gateway-policy-name"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of gateway policy.

## Attribute Reference

The following attributes are exported:
 * `display_name` - Gateway Policy display name. 
 * `created_by` - Who created this resource.
 * `created_at` - When this resource was created.
 * `modified_by` - Who last modified this resource. 
 * `modified_at` - When last modification occurred.   
 
