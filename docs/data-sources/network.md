---
page_title: "Network Data Source"
---

# Network Data Source

Data Source for network. 

## Example Usage

```hcl
data "ochk_network" "sg" {
  name = "network-name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Exact name of Network.

## Attribute Reference

The following attributes are exported:
 * `name` - network name.
 * `network` - network value.
 * `network_type` - network type.
 
    
 
