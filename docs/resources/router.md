---
page_title: "Router Resource"
---

# Router Resource

Resource for managing Tier-1 logical router providing routing functionality in a virtual environment. 
You can create more then one logical router to segment your network.

## Example Usage

```hcl

data "ochk_router" "rt" {
  display_name = "router-T0-name-display-name" 
}


resource "ochk_router" "router" {
  display_name = "router_name"
  parent_router_id = data.ochk_router.rt.id
}


```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) The Tier-1 router name (between 3 and 15 characters).
* `parent_router_id` - (Required) The Tier-0 router id. (default vrf name : T0)

## Attribute Reference


No additional attributes are exported. 
