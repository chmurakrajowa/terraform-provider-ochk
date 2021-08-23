---
page_title: "Billing Tag Data Source"
---

# Billing Tag Data Source

Data Source for getting billing tag by name.

## Example Usage

```hcl
data "ochk_billing_tag" "os1" {
  display_name = var.system_tag_os
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact name of billing tag.


    
 
