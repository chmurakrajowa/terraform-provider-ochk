
---
page_title: "Resource System Tag"
---

# System Tag Resource

Resource System Tag Resource will be available after running BAAS service for client
## Example Usage

```
resource "ochk_system_tag" "res-st-os2" {
    display_name = "${var.test-data-prefix}-system-tag-os2"
}
```


## Argument Reference
The following arguments are supported:
* `display_name` - (Required) The name of the system tag
