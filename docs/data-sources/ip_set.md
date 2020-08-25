---
page_title: "IP Set Data Source"
---

# IP Set Data Source

Data Source for reading IP Sets by display name. 

## Example Usage

```hcl
data "ochk_security_group" "ips" {
  display_name = "ip-set-display-name"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of IP Set.

## Attribute Reference

The following attributes are exported:
 * `display_name` - Display name. 
 * `addresses` - List of addresses this IP Set contains. Each entry has following values:
   * **address**: address, e.g. `8.8.8.8/24`
   * **id**: address identifier
 * `created_by` - Who created this resource.
 * `created_at` - When this resource was created.
 * `modified_by` - Who last modified this resource. 
 * `modified_at` - When last modification occurred. 
     
 
