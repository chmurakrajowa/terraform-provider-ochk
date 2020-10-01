---
page_title: "Router Resource"
---

# Router Resource

Resource for managing Tier-1 logical router to reproduce routing functionality in a virtual environment. 

## Example Usage

```hcl
resource "ochk_router" "router" {
  display_name = "T1"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) The Tier-1 router name.

## Attribute Reference

No additional attributes are exported. 
